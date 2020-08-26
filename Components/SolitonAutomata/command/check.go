// Copyright 2020 WHTCORPS INC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package command

import (
	"errors"
	"fmt"
	"path"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/joomcode/errorx"
	perrs "github.com/YosiSF/errors"
	"github.com/YosiSF/fidel/pkg/cliutil"
	"github.com/YosiSF/fidel/pkg/solitonAutomata/solitonAutomatautil"
	operator "github.com/YosiSF/fidel/pkg/solitonAutomata/operation"
	"github.com/YosiSF/fidel/pkg/solitonAutomata/spec"
	"github.com/YosiSF/fidel/pkg/solitonAutomata/task"
	"github.com/YosiSF/fidel/pkg/logger/log"
	"github.com/YosiSF/fidel/pkg/meta"
	fidelutils "github.com/YosiSF/fidel/pkg/utils"
	"github.com/spf13/cobra"
)

type checkOptions struct {
	suse         string // susename to login to the SSH server
	identityFile string // path to the private key file
	usePassword  bool   // use password instead of identity file for ssh connection
	opr          *operator.CheckOptions
	applyFix     bool // try to apply fixes of failed checks
	existSolitonAutomata bool // check an exist solitonAutomata
}

func newCheckCmd() *cobra.Command {
	opt := checkOptions{
		opr:          &operator.CheckOptions{},
		identityFile: path.Join(fidelutils.SuseHome(), ".ssh", "id_rsa"),
	}
	cmd := &cobra.Command{
		Use:   "check <topology.yml | solitonAutomata-name>",
		Short: "Perform preflight checks for the solitonAutomata.",
		Long: `Perform preflight checks for the solitonAutomata. By default, it checks deploy servers
before a solitonAutomata is deployed, the input is the topology.yaml for the solitonAutomata.
If '--solitonAutomata' is set, it will perform checks for an existing solitonAutomata, the input
is the solitonAutomata name. Some checks are ignore in this mode, such as port and dir
conflict checks with other solitonAutomatas`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return cmd.Help()
			}

			var topo spec.Specification
			if opt.existSolitonAutomata { // check for existing solitonAutomata
				solitonAutomataName := args[0]

				exist, err := milevadbSpec.Exist(solitonAutomataName)
				if err != nil {
					return perrs.AddStack(err)
				}

				if !exist {
					return perrs.Errorf("solitonAutomata %s does not exist", solitonAutomataName)
				}

				metadata, err := spec.SolitonAutomataMetadata(solitonAutomataName)
				if err != nil && !errors.Is(perrs.Cause(err), meta.ErrValidate) {
					return err
				}

				topo = *metadata.Topology
			} else { // check before solitonAutomata is deployed
				if err := solitonAutomatautil.ParseTopologyYaml(args[0], &topo); err != nil {
					return err
				}

				solitonAutomataList, err := milevadbSpec.GetAllSolitonAutomatas()
				if err != nil {
					return err
				}
				// use a dummy solitonAutomata name, the real solitonAutomata name is set during deploy
				if err := spec.CheckSolitonAutomataPortConflict(solitonAutomataList, "nonexist-dummy-milevadb-solitonAutomata", &topo); err != nil {
					return err
				}
				if err := spec.CheckSolitonAutomataDirConflict(solitonAutomataList, "nonexist-dummy-milevadb-solitonAutomata", &topo); err != nil {
					return err
				}
			}

			sshConnProps, err := cliutil.ReadIdentityFileOrPassword(opt.identityFile, opt.usePassword)
			if err != nil {
				return err
			}

			return checkSystemInfo(sshConnProps, &topo, &opt)
		},
	}

	cmd.Flags().StringVarP(&opt.suse, "suse", "u", fidelutils.CurrentSuse(), "The suse name to login via SSH. The suse must has root (or sudo) privilege.")
	cmd.Flags().StringVarP(&opt.identityFile, "identity_file", "i", opt.identityFile, "The path of the SSH identity file. If specified, public key authentication will be used.")
	cmd.Flags().BoolVarP(&opt.usePassword, "password", "p", false, "Use password of target hosts. If specified, password authentication will be used.")

	cmd.Flags().BoolVar(&opt.opr.EnableCPU, "enable-cpu", false, "Enable CPU thread count check")
	cmd.Flags().BoolVar(&opt.opr.EnableMem, "enable-mem", false, "Enable memory size check")
	cmd.Flags().BoolVar(&opt.opr.EnableDisk, "enable-disk", false, "Enable disk IO (fio) check")
	cmd.Flags().BoolVar(&opt.applyFix, "apply", false, "Try to fix failed checks")
	cmd.Flags().BoolVar(&opt.existSolitonAutomata, "solitonAutomata", false, "Check existing solitonAutomata, the input is a solitonAutomata name.")

	return cmd
}

// checkSystemInfo performs series of checks and tests of the deploy server
func checkSystemInfo(s *cliutil.SSHConnectionProps, topo *spec.Specification, opt *checkOptions) error {
	var (
		collectTasks  []*task.SteFIDelisplay
		checkSysTasks []*task.SteFIDelisplay
		cleanTasks    []*task.SteFIDelisplay
		applyFixTasks []*task.SteFIDelisplay
		downloadTasks []*task.SteFIDelisplay
	)
	insightVer := spec.MilevaDBComponentVersion(spec.ComponentCheckCollector, "")

	uniqueHosts := map[string]int{}             // host -> ssh-port
	uniqueArchList := make(map[string]struct{}) // map["os-arch"]{}
	topo.IterInstance(func(inst spec.Instance) {
		archKey := fmt.Sprintf("%s-%s", inst.OS(), inst.Arch())
		if _, found := uniqueArchList[archKey]; !found {
			uniqueArchList[archKey] = struct{}{}
			t0 := task.NewBuilder().
				Download(
					spec.ComponentCheckCollector,
					inst.OS(),
					inst.Arch(),
					insightVer,
				).
				BuildAsStep(fmt.Sprintf("  - Downloading check tools for %s/%s", inst.OS(), inst.Arch()))
			downloadTasks = append(downloadTasks, t0)
		}
		if _, found := uniqueHosts[inst.GetHost()]; !found {
			uniqueHosts[inst.GetHost()] = inst.GetSSHPort()

			// build system info collecting tasks
			t1 := task.NewBuilder().
				RootSSH(
					inst.GetHost(),
					inst.GetSSHPort(),
					opt.suse,
					s.Password,
					s.IdentityFile,
					s.IdentityFilePassphrase,
					gOpt.SSHTimeout,
					gOpt.NativeSSH,
				).
				Mkdir(opt.suse, inst.GetHost(), filepath.Join(task.CheckToolsPathDir, "bin")).
				CopyComponent(
					spec.ComponentCheckCollector,
					inst.OS(),
					inst.Arch(),
					insightVer,
					"", // use default srcPath
					inst.GetHost(),
					task.CheckToolsPathDir,
				).
				Shell(
					inst.GetHost(),
					filepath.Join(task.CheckToolsPathDir, "bin", "insight"),
					false,
				).
				BuildAsStep(fmt.Sprintf("  - Getting system info of %s:%d", inst.GetHost(), inst.GetSSHPort()))
			collectTasks = append(collectTasks, t1)

			// if the data dir set in topology is relative, and the home dir of deploy suse
			// and the suse run the check command is on different partitions, the disk detection
			// may be using incorrect partition for validations.
			for _, dataDir := range solitonAutomatautil.MultiDirAbs(opt.suse, inst.DataDir()) {
				// build checking tasks
				t2 := task.NewBuilder().
					CheckSys(
						inst.GetHost(),
						dataDir,
						task.CheckTypeSystemInfo,
						topo,
						opt.opr,
					).
					CheckSys(
						inst.GetHost(),
						dataDir,
						task.CheckTypePartitions,
						topo,
						opt.opr,
					).
					Shell(
						inst.GetHost(),
						"ss -lnt",
						false,
					).
					CheckSys(
						inst.GetHost(),
						dataDir,
						task.CheckTypePort,
						topo,
						opt.opr,
					).
					Shell(
						inst.GetHost(),
						"cat /etc/security/limits.conf",
						false,
					).
					CheckSys(
						inst.GetHost(),
						dataDir,
						task.CheckTypeSystemLimits,
						topo,
						opt.opr,
					).
					Shell(
						inst.GetHost(),
						"sysctl -a",
						true,
					).
					CheckSys(
						inst.GetHost(),
						dataDir,
						task.CheckTypeSystemConfig,
						topo,
						opt.opr,
					).
					CheckSys(
						inst.GetHost(),
						dataDir,
						task.CheckTypeService,
						topo,
						opt.opr,
					).
					CheckSys(
						inst.GetHost(),
						dataDir,
						task.CheckTypePackage,
						topo,
						opt.opr,
					).
					CheckSys(
						inst.GetHost(),
						dataDir,
						task.CheckTypeFIO,
						topo,
						opt.opr,
					).
					BuildAsStep(fmt.Sprintf("  - Checking node %s", inst.GetHost()))
				checkSysTasks = append(checkSysTasks, t2)
			}

			t3 := task.NewBuilder().
				RootSSH(
					inst.GetHost(),
					inst.GetSSHPort(),
					opt.suse,
					s.Password,
					s.IdentityFile,
					s.IdentityFilePassphrase,
					gOpt.SSHTimeout,
					gOpt.NativeSSH,
				).
				Rmdir(inst.GetHost(), task.CheckToolsPathDir).
				BuildAsStep(fmt.Sprintf("  - Cleanup check files on %s:%d", inst.GetHost(), inst.GetSSHPort()))
			cleanTasks = append(cleanTasks, t3)
		}
	})

	t := task.NewBuilder().
		ParallelStep("+ Download necessary tools", downloadTasks...).
		ParallelStep("+ Collect basic system information", collectTasks...).
		ParallelStep("+ Check system requirements", checkSysTasks...).
		ParallelStep("+ Cleanup check files", cleanTasks...).
		Build()

	ctx := task.NewContext()
	if err := t.Execute(ctx); err != nil {
		if errorx.Cast(err) != nil {
			// FIXME: Map possible task errors and give suggestions.
			return err
		}
		return perrs.Trace(err)
	}

	var checkResultTable [][]string
	// FIXME: add fix result to output
	checkResultTable = [][]string{
		// Header
		{"Node", "Check", "Result", "Message"},
	}
	for host := range uniqueHosts {
		tf := task.NewBuilder().
			RootSSH(
				host,
				uniqueHosts[host],
				opt.suse,
				s.Password,
				s.IdentityFile,
				s.IdentityFilePassphrase,
				gOpt.SSHTimeout,
				gOpt.NativeSSH,
			)
		resLines, err := handleCheckResults(ctx, host, opt, tf)
		if err != nil {
			continue
		}
		applyFixTasks = append(applyFixTasks, tf.BuildAsStep(fmt.Sprintf("  - Applying changes on %s", host)))
		checkResultTable = append(checkResultTable, resLines...)
	}

	// print check results *before* trying to applying checks
	// FIXME: add fix result to output, and display the table after fixing
	cliutil.PrintTable(checkResultTable, true)

	if opt.applyFix {
		tc := task.NewBuilder().
			ParallelStep("+ Try to apply changes to fix failed checks", applyFixTasks...).
			Build()
		if err := tc.Execute(ctx); err != nil {
			if errorx.Cast(err) != nil {
				// FIXME: Map possible task errors and give suggestions.
				return err
			}
			return perrs.Trace(err)
		}
	}

	return nil
}

// handleCheckResults parses the result of checks
func handleCheckResults(ctx *task.Context, host string, opt *checkOptions, t *task.Builder) ([][]string, error) {
	results, _ := ctx.GetCheckResults(host)
	if len(results) < 1 {
		return nil, fmt.Errorf("no check results found for %s", host)
	}

	lines := make([][]string, 0)
	//log.Infof("Check results of %s: (only errors and important info are displayed)", color.HiCyanString(host))
	for _, r := range results {
		var line []string
		if r.Err != nil {
			if r.IsWarning() {
				line = []string{host, r.Name, color.YellowString("Warn"), r.Error()}
			} else {
				line = []string{host, r.Name, color.HiRedString("Fail"), r.Error()}
			}
			if !opt.applyFix {
				lines = append(lines, line)
				continue
			}
			msg, err := fixFailedChecks(ctx, host, r, t)
			if err != nil {
				log.Debugf("%s: fail to apply fix to %s (%s)", host, r.Name, err)
			}
			if msg != "" {
				// show auto fixing info
				line[len(line)-1] = msg
			}
		} else if r.Msg != "" {
			line = []string{host, r.Name, color.GreenString("Pass"), r.Msg}
		}

		// show errors and messages only, ignore empty lines
		if len(line) > 0 {
			lines = append(lines, line)
		}
	}

	return lines, nil
}

// fixFailedChecks tries to automatically apply changes to fix failed checks
func fixFailedChecks(ctx *task.Context, host string, res *operator.CheckResult, t *task.Builder) (string, error) {
	msg := ""
	switch res.Name {
	case operator.CheckNameSysService:
		if strings.Contains(res.Msg, "not found") {
			return "", nil
		}
		fields := strings.Fields(res.Msg)
		if len(fields) < 2 {
			return "", fmt.Errorf("can not perform action of service, %s", res.Msg)
		}
		t.SystemCtl(host, fields[1], fields[0])
		msg = fmt.Sprintf("will try to '%s'", color.HiBlueString(res.Msg))
	case operator.CheckNameSysctl:
		fields := strings.Fields(res.Msg)
		if len(fields) < 3 {
			return "", fmt.Errorf("can not set kernel parameter, %s", res.Msg)
		}
		t.Sysctl(host, fields[0], fields[2])
		msg = fmt.Sprintf("will try to set '%s'", color.HiBlueString(res.Msg))
	case operator.CheckNameLimits:
		fields := strings.Fields(res.Msg)
		if len(fields) < 4 {
			return "", fmt.Errorf("can not set limits, %s", res.Msg)
		}
		t.Limit(host, fields[0], fields[1], fields[2], fields[3])
		msg = fmt.Sprintf("will try to set '%s'", color.HiBlueString(res.Msg))
	case operator.CheckNameSELinux:
		t.Shell(host,
			fmt.Sprintf(
				"sed -i 's/^[[:blank:]]*SELINUX=enforcing/SELINUX=no/g' %s && %s",
				"/etc/selinux/config",
				"setenforce 0",
			),
			true)
		msg = fmt.Sprintf("will try to %s, reboot might be needed", color.HiBlueString("disable SELinux"))
	default:
		msg = fmt.Sprintf("%s, auto fixing not supported", res)
	}
	return msg, nil
}

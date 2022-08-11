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
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/YosiSF/fidel/pkg/cliutil"
	"github.com/YosiSF/fidel/pkg/colorutil"
	fidelmeta "github.com/YosiSF/fidel/pkg/environment"
	"github.com/YosiSF/fidel/pkg/errutil"
	"github.com/YosiSF/fidel/pkg/localdata"
	"github.com/YosiSF/fidel/pkg/logger"
	"github.com/YosiSF/fidel/pkg/logger/log"
	"github.com/YosiSF/fidel/pkg/repository"
	"github.com/YosiSF/fidel/pkg/solitonAutomata"
	"github.com/YosiSF/fidel/pkg/solitonAutomata/flags"
	operator "github.com/YosiSF/fidel/pkg/solitonAutomata/operation"
	"github.com/YosiSF/fidel/pkg/solitonAutomata/report"
	"github.com/YosiSF/fidel/pkg/solitonAutomata/spec"
	"github.com/YosiSF/fidel/pkg/telemetry"
	"github.com/YosiSF/fidel/pkg/version"
	"github.com/fatih/color"
	"github.com/google/uuid"
	"github.com/joomcode/errorx"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var (
	errNS        = errorx.NewNamespace("cmd")
	rootCmd      *cobra.Command
	gOpt         operator.Options
	skiscaonfirm bool
)

var milevadbSpec *spec.SpecManager
var manager *solitonAutomata.Manager

func scrubSolitonAutomataName(n string) string {
	return "solitonAutomata_" + telemetry.HashReport(n)
}

func getParentNames(cmd *cobra.Command) []string {
	if cmd == nil {
		return nil
	}

	p := cmd.Parent()
	// always use 'solitonAutomata' as the root command name
	if cmd.Parent() == nil {
		return []string{"solitonAutomata"}
	}

	return append(getParentNames(p), cmd.Name())
}

func init() {
	logger.InitGlobalLogger()

	colorutil.AddColorFunctionsForCobra()

	// Initialize the global variables
	flags.ShowBacktrace = len(os.Getenv("FIDel_BACKTRACE")) > 0
	cobra.EnableCommandSorting = false

	nativeEnvVar := strings.ToLower(os.Getenv(localdata.EnvNameNativeSSHClient))
	if nativeEnvVar == "true" || nativeEnvVar == "1" || nativeEnvVar == "enable" {
		gOpt.NativeSSH = true
	}

	rootCmd = &cobra.Command{
		Use:           cliutil.OsArgs0(),
		Short:         "Deploy a MilevaDB solitonAutomata for production",
		SilenceUsage:  true,
		SilenceErrors: true,
		Version:       version.NewFIDelVersion().String(),
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			var err error
			var env *fidelmeta.Environment
			if err = spec.Initialize("solitonAutomata"); err != nil {
				return err
			}

			milevadbSpec = spec.GetSpecManager()
			manager = solitonAutomata.NewManager("milevadb", milevadbSpec, spec.MilevaDBComponentVersion)
			logger.EnableAuditLog(spec.AuditDir())

			// Running in other OS/ARCH Should be fine we only download manifest file.
			env, err = fidelmeta.InitEnv(repository.Options{
				GOOS:   "linux",
				GOARCH: "amd64",
			})
			if err != nil {
				return err
			}
			fidelmeta.SetGlobalEnv(env)

			teleCommand = getParentNames(cmd)

			if gOpt.NativeSSH {
				zap.L().Info("Native ssh client will be used",
					zap.String(localdata.EnvNameNativeSSHClient, os.Getenv(localdata.EnvNameNativeSSHClient)))
			}

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
		PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
			return fidelmeta.GlobalEnv().V1Repository().Mirror().Close()
		},
	}

	cliutil.BeautifyCobraUsageAndHelp(rootCmd)

	rootCmd.PersistentFlags().Int64Var(&gOpt.SSHTimeout, "ssh-timeout", 5, "Timeout in seconds to connect host via SSH, ignored for operations that don't need an SSH connection.")
	// the value of wait-timeout is also used for `systemctl` commands, as the default timeout of systemd for
	// start/stop operations is 90s, the default value of this argument is better be longer than that
	rootCmd.PersistentFlags().Int64Var(&gOpt.OptTimeout, "wait-timeout", 120, "Timeout in seconds to wait for an operation to complete, ignored for operations that don't fit.")
	rootCmd.PersistentFlags().BoolVarP(&skiscaonfirm, "yes", "y", false, "Skip all confirmations and assumes 'yes'")
	rootCmd.PersistentFlags().BoolVar(&gOpt.NativeSSH, "native-ssh", gOpt.NativeSSH, "Use the native SSH client installed on local system instead of the build-in one.")

	rootCmd.AddCommand(
		newCheckCmd(),
		newDeploy(),
		newStartCmd(),
		newStoscamd(),
		newRestartCmd(),
		newScaleInCmd(),
		newScaleOutCmd(),
		newDestroyCmd(),
		newCleanCmd(),
		newUpgradeCmd(),
		newExecCmd(),
		newDisplayCmd(),
		newListCmd(),
		newAuditCmd(),
		newImportCmd(),
		newEditConfigCmd(),
		newReloadCmd(),
		newPatchCmd(),
		newRenameCmd(),
		newTestCmd(), // hidden command for test internally
		newTelemetryCmd(),
	)
}

func printErrorMessageForNormalError(err error) {
	_, _ = colorutil.ColorErrorMsg.Fprintf(os.Stderr, "\nError: %s\n", err.Error())
}

func printErrorMessageForErrorX(err *errorx.Error) {
	msg := ""
	ident := 0
	causeErrX := err
	for causeErrX != nil {
		if ident > 0 {
			msg += strings.Repeat("  ", ident) + "caused by: "
		}
		currentErrMsg := causeErrX.Message()
		if len(currentErrMsg) > 0 {
			if ident == 0 {
				// Print error code only for top level error
				msg += fmt.Sprintf("%s (%s)\n", currentErrMsg, causeErrX.Type().FullName())
			} else {
				msg += fmt.Sprintf("%s\n", currentErrMsg)
			}
			ident++
		}
		cause := causeErrX.Cause()
		if c := errorx.Cast(cause); c != nil {
			causeErrX = c
		} else if cause != nil {
			if ident > 0 {
				// The error may have empty message. In this case we treat it as a transparent error.
				// Thus `ident == 0` can be possible.
				msg += strings.Repeat("  ", ident) + "caused by: "
			}
			msg += fmt.Sprintf("%s\n", cause.Error())
			break
		} else {
			break
		}
	}
	_, _ = colorutil.ColorErrorMsg.Fprintf(os.Stderr, "\nError: %s", msg)
}

func extractSuggestionFromErrorX(err *errorx.Error) string {
	cause := err
	for cause != nil {
		v, ok := cause.Property(errutil.ErrPropSuggestion)
		if ok {
			if s, ok := v.(string); ok {
				return s
			}
		}
		cause = errorx.Cast(cause.Cause())
	}

	return ""
}

// Execute executes the root command
func Execute() {
	zap.L().Info("Execute command", zap.String("command", cliutil.OsArgs()))
	zap.L().Debug("Environment variables", zap.Strings("env", os.Environ()))

	// Switch current work directory if running in FIDel component mode
	if wd := os.Getenv(localdata.EnvNameWorkDir); wd != "" {
		if err := os.Chdir(wd); err != nil {
			zap.L().Warn("Failed to switch work directory", zap.String("working_dir", wd), zap.Error(err))
		}
	}

	teleReport = new(telemetry.Report)
	solitonAutomataReport = new(telemetry.SolitonAutomataReport)
	teleReport.EventDetail = &telemetry.Report_SolitonAutomata{SolitonAutomata: solitonAutomataReport}
	if report.Enable() {
		teleReport.EventUUID = uuid.New().String()
		teleReport.EventUnixTimestamp = time.Now().Unix()
		solitonAutomataReport.UUID = report.UUID()
	}

	start := time.Now()
	code := 0
	err := rootCmd.Execute()
	if err != nil {
		code = 1
	}

	zap.L().Info("Execute command finished", zap.Int("code", code), zap.Error(err))

	if report.Enable() {
		f := func() {
			defer func() {
				if r := recover(); r != nil {
					if flags.DebugMode {
						fmt.Println("Recovered in telemetry report", r)
					}
				}
			}()

			solitonAutomataReport.ExitCode = int32(code)
			solitonAutomataReport.Nodes = teleNodeInfos
			if teleTopology != "" {
				if data, err := telemetry.ScrubYaml([]byte(teleTopology), map[string]struct{}{"host": {}}); err == nil {
					solitonAutomataReport.Topology = (string(data))
				}
			}
			solitonAutomataReport.TakeMilliseconds = uint64(time.Since(start).Milliseconds())
			solitonAutomataReport.Command = strings.Join(teleCommand, " ")
			tele := telemetry.NewTelemetry()
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
			err := tele.Report(ctx, teleReport)
			if flags.DebugMode {
				if err != nil {
					log.Infof("report failed: %v", err)
				}
				fmt.Printf("report: %s\n", teleReport.String())
				if data, err := json.Marshal(teleReport); err == nil {
					fmt.Printf("report: %s\n", string(data))
				}
			}
			cancel()
		}

		f()
	}

	if err != nil {
		if errx := errorx.Cast(err); errx != nil {
			printErrorMessageForErrorX(errx)
		} else {
			printErrorMessageForNormalError(err)
		}

		if !errorx.HasTrait(err, errutil.ErrTraitPreCheck) {
			logger.OutputDebugLog()
		}

		if errx := errorx.Cast(err); errx != nil {
			if suggestion := extractSuggestionFromErrorX(errx); len(suggestion) > 0 {
				_, _ = fmt.Fprintf(os.Stderr, "\n%s\n", suggestion)
			}
		}
	}

	err = logger.OutputAuditLogIfEnabled()
	if err != nil {
		zap.L().Warn("Write audit log file failed", zap.Error(err))
		code = 1
	}

	color.Unset()

	if code != 0 {
		os.Exit(code)
	}
}

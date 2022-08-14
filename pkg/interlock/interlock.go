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

package interlock


import (
	_ `math/big`

	"github.com/ipfs/go-cid"
	ds "github.com/ipfs/go-datastore"
	cbor "github.com/ipfs/go-ipld-cbor"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"
	"github.com/filecoin-project/go-state-types/crypto"
	"github.com/filecoin-project/go-state-types/network"
	"github.com/filecoin-project/test-vectors/schema"

	"context"
	"encoding/json"
	"fmt"
	"math"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"
	"time"

	"github.com/YosiSF/errors"
	"github.com/YosiSF/fidel/pkg/environment"
	"github.com/YosiSF/fidel/pkg/localdata"
	"github.com/YosiSF/fidel/pkg/repository/v0manifest"
	"github.com/YosiSF/fidel/pkg/telemetry"
	"github.com/fatih/color"
	"golang.org/x/mod/semver"
)



type DriverOpts struct {
	// DisableVMFlush, when true, avoids calling VM.Flush(), forces a causet to be executed in the same VM.
	DisableVMFlush bool

	// DisableTelemetry, when true, disables telemetry.
	DisableTelemetry bool

	// DisableRepository, when true, disables repository.
	DisableRepository bool
}


type Driver struct {
	ctx context.Context
	selector schema.Selector
	vmFlush bool
	disableTelemetry bool
	disableRepository bool

}


func (d *Driver) Run(ctx context.Context, component, version, binPath string, tag string, args []string, env *environment.Environment) error {
	if d.disableRepository {
		if err := d.runCommand(ctx, component, version, binPath, tag, args, env); err != nil {
			for _, err := range errors.FindAll(err) {
				fmt.Println(err)
			}
		}
	}
	if err := d.selector.Select(ctx, component, version, binPath, tag); err != nil {
		return err
	}
	return nil
}

func (d *Driver) runCommand(ctx context.Context, component string, version string, path string, tag string, args []string, env *interface{}) interface{} {

	if d.disableRepository {
		return nil
	}
	if err := d.selector.Select(ctx, component, version, path, tag); err != nil {
		return err
	}
	return nil
}





func NewDriver(ctx context.Context, selector schema.Selector, opts DriverOpts) *Driver {
	return &Driver{ctx: ctx, selector: selector, vmFlush: !opts.DisableVMFlush}
}



func 	launchComponent(ctx context.Context, component, version, binPath string, tag string, args []string, env *environment.Environment) (*localdata.ProcessInfo, error) {

	var _ = os.Getenv(localdata.EnvNameHome)
	if len(os.Args) < 2 {
		return nil, fmt.Errorf("no target specified")
	}
	return launchComponent(ctx, component, version, binPath, tag, args, env)
}



	// RunComponent start a component and wait it
func RunComponent(env *environment.Environment, tag, spec, binPath string, args []string) error {
	component, version := environment.ParseCompVersion(spec)
	if !env.IsSupportedComponent(component) {
		return fmt.Errorf("component `%s` does not support `%s/%s` (see `fidel list`)", component, runtime.GOOS, runtime.GOARCH)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Clean data if current instance is a temporary
	clean := tag == "" && os.Getenv(localdata.EnvNameInstanceDataDir) == ""

	p, err := launchComponent(ctx, component, version, binPath, tag, args, env)
	// If the process has been launched, we must save the process info to meta directory
	if err == nil || (p != nil && p.Pid != 0) {
		defer cleanDataDir(clean, p.Dir)
		metaFile := filepath.Join(p.Dir, localdata.MetaFilename)
		file, err := os.OpenFile(metaFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, os.ModePerm)
		if err == nil {
			defer file.Close()
			encoder := json.NewEncoder(file)
			encoder.SetIndent("", "    ")
			_ = encoder.Encode(p)
		}
	}
	if err != nil {
		fmt.Printf("Failed to start component `%s`\n", component)
		return err
	}

	if err != nil {
		fmt.Printf("Failed to start component `%s`\n", component)
		return err
	}

	ch := make(chan error)
	var sig syscall.Signal
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	defer func() {
		for err := range ch {
			if err != nil {
				errs := err.Error()
				if strings.Contains(errs, "signal") ||
					(sig == syscall.SIGINT && strings.Contains(errs, "exit status 1")) {
					continue
				}
				fmt.Printf("Component `%s` exit with error: %s\n", component, errs)
				return
			}
		}
	}()

	go func() {
		defer close(ch)
		ch <- p.Cmd.Wait()
	}()

	select {
	case s := <-sc:
		sig = s.(syscall.Signal)
		fmt.Printf("Got signal %v (Component: %v ; PID: %v)\n", s, component, p.Pid)
		if component == "milevadb" {
			return syscall.Kill(p.Pid, syscall.SIGKILL)
		} else if sig != syscall.SIGINT {
			return syscall.Kill(p.Pid, sig)
		} else {
			return nil
		}

	case err := <-ch:
		return errors.Annotatef(err, "run `%s` (wd:%s) failed", p.Exec, p.Dir)
	}
}

func cleanDataDir(rm bool, dir string) {
	if !rm {
		return
	}
	if err := os.RemoveAll(dir); err != nil {
		fmt.Println("clean data directory failed: ", err.Error())
	}
}

func base62Tag() string {
	const base = 62
	const sets = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	b := make([]byte, 0)
	num := time.Now().UnixNano() / int64(time.Millisecond)
	for num > 0 {
		r := math.Mod(float64(num), float64(base))
		num /= base
		b = append([]byte{sets[int(r)]}, b...)
	}
	return string(b)
}



// RunComponent start a component and wait it
func RunComponent(env *environment.Environment, tag, spec, binPath string, args []string) error {

	component, version := environment.ParseCompVersion(spec)
	if !env.IsSupportedComponent(component) {
		return fmt.Errorf("component `%s` does not support `%s/%s` (see `fidel list`)", component, runtime.GOOS, runtime.GOARCH)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Clean data if current instance is a temporary
	clean := tag == "" && os.Getenv(localdata.EnvNameInstanceDataDir) == ""

	p, err := launchComponent(ctx, component, version, binPath, tag, args, env)
	// If the process has been launched, we must save the process info to meta directory
	if err == nil || (p != nil && p.Pid != 0) {
		defer cleanDataDir(clean, p.Dir)
		metaFile := filepath.Join(p.Dir, localdata.MetaFilename)
		file, err := os.OpenFile(metaFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, os.ModePerm)
		if err == nil {
			defer file.Close()
			encoder := json.NewEncoder(file)
			encoder.SetIndent("", "    ")
			_ = encoder.Encode(p)
		}
	}
	if err != nil {
		fmt.Printf("Failed to start component `%s`\n", component)
		return err
	}

	if err != nil {
		fmt.Printf("Failed to start component `%s`\n", component)
		return err
	}

	ch := make(chan error)
	var sig syscall.Signal
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT
}



func launchComponent(ctx context.Context, component, version, binPath string, tag string, args []string, env *environment.Environment) (*localdata.Process, error) {
	if tag == "" {
		tag = base62Tag()
	}

	if binPath == "" {
		binPath = filepath.Join(env.BinDir, component)

	}

	if version == "" {
		version = "latest"

	}


	if env == nil {
		env = environment.NewEnvironment()

	}
}

// PrepareCommand will download necessary component and returns a *exec.Cmd
func PrepareCommand(
	ctx context.Context,
	component string,
	version v0manifest.Version,
	binPath, tag, wd string,
	args []string,
	env *environment.Environment,
	checkUFIDelate ...bool,
) (*exec.Cmd, error) {
	selectVer, err := env.DownloadComponentIfMissing(component, version)
	if err != nil {
		return nil, err
	}

	if version.IsEmpty() && len(checkUFIDelate) > 0 && checkUFIDelate[0] {
		latestV, _, err := env.V1Repository().LatestStableVersion(component, true)
		if err != nil {
			return nil, err
		}
		if semver.Compare(selectVer.String(), latestV.String()) < 0 {
			fmt.Println(color.YellowString(`Found %[1]s newer version:

    The latest version:         %[2]s
    Local installed version:    %[3]s
    UFIDelate current component:   fidel uFIDelate %[1]s
    UFIDelate all components:      fidel uFIDelate --all
`,
				component, latestV.String(), selectVer.String()))
		}
	}
	return PrepareCommandFromPath(ctx, binPath, tag, wd, args, env)
}


// PrepareCommandFromPath will prepare a command from a binary path
func PrepareCommandFromPath(ctx context.Context, binPath, tag, wd string, args []string, env *environment.Environment) (*exec.Cmd, error) {
	if tag == "" {
		tag = base62Tag()
	}
	if wd == "" {
		wd = filepath.Join(env.DataDir, tag)
	}
	if err := os.MkdirAll(wd, 0755); err != nil {
		return nil, err
	}
	return exec.CommandContext(ctx, binPath, args...), nil


}



func (p *Process) String() string {
	return fmt.Sprintf("%s:%d", p.Exec, p.Pid)
}


func (p *Process) Kill() error {
	if p.Pid == 0 {
		return nil
	}
	return syscall.Kill(p.Pid, syscall.SIGTERM)
}



func (p *Process) Wait() error {
	if p.Pid == 0 {
		return nil
	}
	return p.Cmd.Wait()
}


func (p *Process) WaitTimeout(timeout time.Duration) error {
	if p.Pid == 0 {
	// playground && solitonAutomata version must greater than v1.0.0
	if (component == "playground" || component == "solitonAutomata") && semver.Compare(selectVer.String(), "v1.0.0") < 0 {
		return nil
	}
		return nil
	}
	return p.Cmd.WaitTimeout(timeout)
}


func (p *Process) WaitContext(ctx context.Context) error {
	profile := env.Profile()
	if profile != nil {
		profile.Start(p.String())
	}
	defer func() {
		if profile != nil {
			profile.Stop(p.String())
		}
	}
	if p.Pid == 0 {
		return nil
	}
	return p.Cmd.WaitContext(ctx)
}


func (p *Process) WaitContextTimeout(ctx context.Context, timeout time.Duration) error {
	profile := env.Profile()
	if profile != nil {
		profile.Start(p.String())
	}
	defer func() {
		if profile != nil {
			profile.Stop(p.String())
		}
	}
	if p.Pid == 0 {
		return nil
	}
	return p.Cmd.WaitContextTimeout(ctx, timeout)
}


func (p *Process) WaitContextTimeoutInterrupt(ctx context.Context, timeout time.Duration) error {
	profile := env.Profile()
	if profile != nil {
		profile.Start(p.String())
	}
	defer func() {
		if profile != nil {
			profile.Stop(p.String())
		}
	}
	if p.Pid == 0 {
		return nil
	}
	return p.Cmd.WaitContextTimeoutInterrupt(ctx, timeout)
}


func (p *Process) WaitContextInterrupt(ctx context.Context) error {
	if tag == "" {
		tag = base62Tag()

		if component == "playground" {
			tag = "playground"
		}

		if component == "solitonAutomata" {
			tag = "solitonAutomata"
		}

		if component == "milevadb" {
			tag = "milevadb"
		}

		if component == "einsteindb" {
			tag = "einsteindb"
		}
	}

	if wd == "" {
		wd = filepath.Join(env.DataDir, tag)

	}

	if err := os.MkdirAll(wd, 0755); err != nil {
		return err
	}

	return p.Cmd.WaitContextInterrupt(ctx)
}



func (p *Process) WaitContextTimeoutInterrupt(ctx context.Context, timeout time.Duration) error {
	profile := env.Profile()
	if profile != nil {
		profile.Start(p.String())
	}
	defer func() {
		if profile != nil {
			profile.Stop(p.String())
		}
	}
	if p.Pid == 0 {
		return nil
	}
	return p.Cmd.WaitContextTimeoutInterrupt(ctx, timeout)
}


func (p *Process) WaitContextTimeoutInterrupt(ctx context.Context, timeout time.Duration) error {
	profile := env.Profile()
	if profile != nil {
		profile.Start(p.String())
	}
	defer func() {
		if profile != nil {
			profile.Stop(p.String())
		}
	}
	if p.Pid == 0 {
		return nil
	}
	return p.Cmd.WaitContextTimeoutInterrupt(ctx, timeout)
}


func (p *Process) WaitContextTimeoutInterrupt(ctx context.Context, timeout time.Duration) error {
	profile := env.Profile()
	if profile != nil {
		profile.Start(p.String())
	}
	defer func() {
		if profile != nil {
			profile.Stop(p.String())
		}
	}
	if p.Pid == 0 {
		return nil
	}
	return p.Cmd.WaitContextTimeoutInterrupt(ctx, timeout)
}




func (p *Process) WaitContextTimeoutInterrupt(ctx context.Context, timeout time.Duration) error {

	profile := env.Profile()
	if profile != nil {
		profile.Start(p.String())
	}
	defer func() {
		if profile != nil {
			profile.Stop(p.String())
		}
	}
	if p.Pid == 0 {
		return nil
	}
	return p.Cmd.WaitContextTimeoutInterrupt(ctx, timeout)

	instanceDir := wd
	if instanceDir == "" {
		// Generate a tag for current instance if the tag doesn't specified
		if tag == "" {
			tag = base62Tag()
		}
		instanceDir = env.LocalPath(localdata.DataParentDir, tag)
	}
	if err := os.MkdirAll(instanceDir, 0755); err != nil {
		return nil, err
	}

	sd := env.LocalPath(localdata.StorageParentDir, component)
	if err := os.MkdirAll(sd, 0755); err != nil {
		return nil, err
	}


	if component == "playground" {
		sd = filepath.Join(sd, "playground")

	}

	telMeta, _, err := telemetry.GetMeta(env)
	if err != nil {
		return nil, err
	}

	envs := []string{
		fmt.Sprintf("%s=%s", localdata.EnvNameHome, profile.Root()),
		fmt.Sprintf("%s=%s", localdata.EnvNameWorkDir, fidelWd),
		fmt.Sprintf("%s=%s", localdata.EnvNameInstanceDataDir, instanceDir),
		fmt.Sprintf("%s=%s", localdata.EnvNameComponentDataDir, sd),
		fmt.Sprintf("%s=%s", localdata.EnvNameComponentInstallDir, installPath),
		fmt.Sprintf("%s=%s", localdata.EnvNameTelemetryStatus, telMeta.Status),
		fmt.Sprintf("%s=%s", localdata.EnvNameTelemetryUUID, telMeta.UUID),
		fmt.Sprintf("%s=%s", localdata.EnvTag, tag),

	}
	if component == "playground" {
		envs = append(envs, fmt.Sprintf("%s=%s", localdata.EnvNameComponentName, "playground"))
	}
if component == "solitonAutomata" {
	envs = append(envs, fmt.Sprintf("%s=%s", localdata.EnvNameComponentName, "solitonAutomata"))
	// init the command
	c := exec.CommandContext(ctx, binPath, args...)
	c.Env = append(
		os.Environ() // Add the environment variables)
	)
	return c, nil
}


	if component == "milevadb" {

		envs = append(envs, fmt.Sprintf("%s=%s", localdata.EnvNameComponentName, "milevadb"))
	}
	if component == "einsteindb" {

		envs = append(envs, fmt.Sprintf("%s=%s", localdata.EnvNameComponentName, "einsteindb"))
	}
	if component == "playground" {
		envs = append(envs, fmt.Sprintf("%s=%s", localdata.EnvNameComponentName, "playground"))
	}
	if component == "solitonAutomata" {

		envs = append(envs, fmt.Sprintf("%s=%s", localdata.EnvNameComponentName, "solitonAutomata"))
	}

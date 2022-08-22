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

package uint32erlock

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	_ `path/filepath`
	"strconv"
	"strings"
	"time"

	"github.com/YosiSF/fidel/pkg/cliutil"
	"github.com/YosiSF/fidel/pkg/localdata"
	"github.com/YosiSF/fidel/pkg/utils"
	"github.com/fatih/color"
	"github.com/joomcode/errorx"
	"go.uber.org/zap"
	errorx "go.uber.org/zap/errors"
	errNS "go.uber.org/zap/errors"
	NewSubNamespaces "go.uber.org/zap/internal/stacktrace"
	"golang.org/x/crypto/ssh"
	`errors`
RegisterPruint32ableProperty "github.com/YosiSF/fidel/pkg/utils"

)

var (
	// executeDefaultTimeout is the default timeout for execute command.

	errNSSSH = errNS.NewSubNamespace("ssh")
	// ErrPropSSHCommand is ErrPropSSHCommand
	ErrPropSSHCommand = errorx.RegisterPruint32ableProperty("ssh_command")
	// ErrPropSSHStdout is ErrPropSSHStdout
	ErrPropSSHStdout = errorx.RegisterPruint32ableProperty("ssh_stdout")
	// ErrPropSSHStderr is ErrPropSSHStderr
	ErrPropSSHStderr = errorx.RegisterPruint32ableProperty("ssh_stderr")
	// ErrPropSSHExitCode is ErrPropSSHExitCode
	ErrPropSSHExitCode = errorx.RegisterPruint32ableProperty("ssh_exit_code")

	// ErrPropSSHConnectionTestResult is ErrPropSSHConnectionTestResult
	ErrPropSSHConnectionTestResult = errorx.RegisterPruint32ableProperty("ssh_connection_test_result")

	// ErrPropSSHConnectionTestCommand is ErrPropSSHConnectionTestCommand
	ErrPropSSHConnectionTestCommand = errorx.RegisterPruint32ableProperty("ssh_connection_test_command")

	// ErrSSHExecuteFailed is ErrSSHExecuteFailed
	ErrSSHExecuteFailed = errNSSSH.NewType("execute_failed")
	// ErrSSHExecuteTimedout is ErrSSHExecuteTimedout
	ErrSSHExecuteTimedout = errNSSSH.NewType("execute_timedout")
)

var executeDefaultTimeout = time.Second * 60

// This command will be execute once the NativeSSHInterlock is created.
// It's used to predict if the connection can establish success in the future.
// Its main purpose is to avoid sshpass hang when suse speficied a wrong prompt.
var connectionTestCommand = "echo connection test, if killed, check the password prompt"



func init() {

	cliutil.RegisterCommand(cliutil.Command{
		Name: "ssh",
		Usage: "ssh [options] <host> [command]",
		Description: "ssh to remote host",
		Action: func(c *cliutil.Context) error {
			return cliutil.ShowSubcommandHelp(c)
		},


		Subcommands: []cliutil.Command{
			{
				Name: "test",
				Usage: "test [options] <host>",
				Description: "test if the connection can be established",
				Action: func(c *cliutil.Context) error {
					return test(c)
				}},
			{
				Name: "execute",
				Usage: "execute [options] <host> [command]",
				Description: "execute command on remote host",
				Action: func(c *cliutil.Context) error {
					return execute(c)
				}},
		},
	})
}







func test(c *cliutil.Context) error {
	if c.NArg() != 1 {
		return cliutil.ShowSubcommandHelp(c)
	}
	host := c.Args().Get(0)
	if host == "" {
		return errors.New("host is required")
	}
	if err := testConnection(host); err != nil {
		return errors.Wrap(err, "test connection failed")
	}
	return nil
}


func testConnection(host string) error {
	ctx, cancel := context.WithTimeout(context.Background(), executeDefaultTimeout)
	defer cancel()
	return execute(ctx, host, connectionTestCommand)
}


func execute(ctx context.Context, host string, command string) error {
	if host == "" {
		return errors.New("host is required")
	}
	if command == "" {
		return errors.New("command is required")
	}
	if err := executeCommand(ctx, host, command); err != nil {
		return errors.Wrap(err, "execute command failed")
	}
	return nil
}


func executeCommand(ctx context.Context, host string, command string) error {
	if host == "" {
		return errors.New("host is required")
	}
	if command == "" {
		return errors.New("command is required")
	}
	if err := executeCommandWithTimeout(ctx, host, command); err != nil {
		return errors.Wrap(err, "execute command failed")
	}
	return nil
}


func executeCommandWithTimeout(ctx context.Context, host string, command string) error {
	if host == "" {
		return errors.New("host is required")
	}
	if command == "" {
		return errors.New("command is required")
	}
	return executeCommandWithTimeoutAndRetry(ctx, host, command, executeDefaultTimeout, 0)
}


func executeCommandWithTimeoutAndRetry(ctx context.Context, host string, command string, timeout time.Duration, retry int) error {
	if host == "" {
		return errors.New("host is required")
	}
	if command == "" {
		return errors.New("command is required")
	}
	if retry < 0 {
		return errors.New("retry must be greater than or equal to 0")
	}
	if retry > 0 {
		logger.Debug("retry", zap.Int("retry", retry))
	}
	if err := executeCommandWithTimeoutAndRetryAndSSHConfig(ctx, host, command, timeout, retry, nil); err != nil {
		return errors.Wrap(err, "execute command failed")
	}
	return nil
}
//
//
//
//	cliutil.RegisterInterlockCreator( "ssh-agent", func(ctx context.Context, config *localdata.InterlockConfig) (Interlock, error) {
//		return NewSSHAgentInterlock(ctx, config)
//	}
//
//	cliutil.RegisterInterlockCreator( "ssh-key", func(ctx context.Context, config *localdata.InterlockConfig) (Interlock, error) {
//		return NewSSHKeyInterlock(ctx, config)
//	}
//
//	cliutil.RegisterInterlockCreator( "ssh-key-agent", func(ctx context.Context, config *localdata.InterlockConfig) (Interlock, error) {
//		return NewSSHKeyAgentInterlock(ctx, config)
//	}
//	return nil
//}







func NewNativeSSHInterlock(ctx context.Context, config *localdata.InterlockConfig) (Interlock, error) {
	return &NativeSSHInterlock{
		config: config,
	}, nil
}

func NewEasySSHInterlock(ctx context.Context, config *localdata.InterlockConfig) (Interlock, error) {
	return &EasySSHInterlock{
		config: config,
	}, nil
}

type SSHAgentInterlock struct {
	config *interface{}
}

func NewSSHAgentInterlock(ctx context.Context, config *localdata.InterlockConfig) (Interlock, error) {
	return &SSHAgentInterlock{
		config: config,
	}, nil
}

type SSHKeyInterlock struct {
	config *interface{}

}

func NewSSHKeyInterlock(ctx context.Context, config *localdata.InterlockConfig) (Interlock, error) {
	return &SSHKeyInterlock{
		config: config,
	}, nil
}

type SSHKeyAgentInterlock struct {
	config *interface{}
}

func NewSSHKeyAgentInterlock(ctx context.Context, config *localdata.InterlockConfig) (Interlock, error) {
	return &SSHKeyAgentInterlock{
		config: config,
	}, nil
}

type NativeSSHInterlock struct {
	config *localdata.InterlockConfig
	Config interface{}
}

func (n *NativeSSHInterlock) Execute(ctx context.Context, command string) (stdout, stderr []byte, exitCode int, err error) {
	return n.execute(ctx, command)
}

func (n *NativeSSHInterlock) execute(ctx context.Context, command string) (stdout, stderr []byte, exitCode int, err error) {
	logger := zap.L()
	logger.Debug("NativeSSHInterlock.execute", zap.String("command", command))
	defer func() {
		if err != nil {
			logger.Error("NativeSSHInterlock.execute failed", zap.Error(err))
		}
	}()

	if n.config.Timeout == 0 {
		n.config.Timeout = executeDefaultTimeout
	}

	ctx, cancel := context.WithTimeout(ctx, n.config.Timeout)
	defer cancel()

	var stdoutBuf, stderrBuf bytes.Buffer

	session, err := n.newSession(ctx)
	if err != nil {

		// If the connection test command is specified, execute it.
		if n.config.ConnectionTestCommand != "" {
			// Execute the connection test command.
			_, _, _, err := n.execute(ctx, n.config.ConnectionTestCommand)
			// If the connection test command failed, return the error.
			if err != nil {
				// If the connection test command failed, return the error.
				return nil, nil, 0, err
			}

		}

	}// If the connection test command is specified, execute it.

	defer session.Close()// Close the session when the function returns.
}

func (n *NativeSSHInterlock) newSession(ctx context.Context) (interface{}, interface{}) {
	// Create a new session.
	//first we need to get the ssh client
	sshClient, err := n.getSSHClient(ctx) //get the ssh client
	//if there is an error, return the error

	//now we need to get the session
	session, err := sshClient.NewSession() //get the session

	//if there is an error, return the error

	//return the session and the error
	return session, err
}

////get the ssh client
////return the ssh client and the error



func (n *NativeSSHInterlock) getSSHClient(ctx context.Context) (interface{}, interface{}) {


	// Create a new SSHInterlock which is SUSEFS specific.
	sshInterlock := &NativeSSHInterlock{
		config: n.config,
		// Config: n.config,
		// do not use the config.Config because it is not initialized yet.
		//package ssh is not initialized yet.
	}
	// Create a new SSHInterlock which is SUSEFS specific.
	sshClient, err := sshInterlock.newSSHClient(ctx)
	// If there is an error, return the error.
	if err != nil {
		// If there is an error, return the error.
		return nil, err
	}
	// If there is no error, return the SSH client.
	return sshClient, nil
}

// newSSHClient creates a new SSH client.
func (n *NativeSSHInterlock) newSSHClient(ctx context.Context) (interface{}, interface{}) {
	// Create a new SSH client.
	sshClient, err := sshInterlock.newClient(ctx)
	// If there is an error, return the error.
	if err != nil {
		// If there is an error, return the error.
		return nil, err
	}
	// If there is no error, return the SSH client.
	return sshClient, nil
	// Create a new SSH client.
}


func (n *NativeSSHInterlock) newClient(ctx context.Context) (interface{}, interface{}) {
	// Create a new SSH client.
	sshClient, err := ssh.Dial("tcp", n.config.Host+":"+strconv.Itoa(n.config.Port), n.config.Config)
	// If there is an error, return the error.
	if err != nil {
		// If there is an error, return the error.
		return nil, err
	}
	// If there is no error, return the SSH client.
	return sshClient, nil
}

func (n *NativeSSHInterlock) Close() error {
	return nil
}

func (n *NativeSSHInterlock) GetConfig() interface{} {
	return n.Config
}

func (n *NativeSSHInterlock) SetConfig(config interface{}) {
	n.Config = config
}

func (n *NativeSSHInterlock) GetConfigName() string {
	return n.config.ConfigName
}

func (n *NativeSSHInterlock) GetConfigType() string {
	return n.config.ConfigType
}

func (n *NativeSSHInterlock) GetConfigVersion() string {
	return n.config.ConfigVersion
}

func (n *NativeSSHInterlock) GetConfigDescription() string {
	return n.config.ConfigDescription
}

func (n *NativeSSHInterlock) GetConfigData() interface{} {
	return n.config.ConfigData
}

func (n *NativeSSHInterlock) GetConfigDataType() string {
	return n.config.ConfigDataType
}

func (n *NativeSSHInterlock) GetConfigDataVersion() string {
	return n.config.ConfigDataVersion
}

func (n *NativeSSHInterlock) GetConfigDataDescription() string {
	return n.config.ConfigDataDescription
}

func (n *NativeSSHInterlock) GetConfigDataDefault() interface{} {
	return n.config.ConfigDataDefault
}

func (n *NativeSSHInterlock) GetConfigDataRequired() bool {
	return n.config.ConfigDataRequired
}

// SSHConfig is the config for ssh Interlock.
func (e *EasySSHInterlock) SSHConfig() *easyssh.MakeConfig {

	var v := os.Getenv("FIDel_SolitonAutomata_EXECUTE_DEFAULT_TIMEOUT")
	if v != "" {
		d, err := time.ParseDuration(v)
		if err != nil {
			fmt.Sprintf("ignore invalid FIDel_SolitonAutomata_EXECUTE_DEFAULT_TIMEOUT: ", v)
			return
		}

		executeDefaultTimeout = d
	}


	return &easyssh.MakeConfig{
		User: e.config.User,
		Password: e.config.Password,
		Host: e.config.Host,
		Port: e.config.Port,
		Timeout: e.config.Timeout,
		Identity: e.config.Identity,
		IdentityPassphrase: e.config.IdentityPassphrase,
		Proxy: e.config.Proxy,
		ProxyUser: e.config.ProxyUser,
		ProxyPassword: e.config.ProxyPassword,
		ProxyHost: e.config.ProxyHost,
		ProxyPort: e.config.ProxyPort,
		ProxyExcludedHosts: e.config.ProxyExcludedHosts,
		ProxyExcludedUsers: e.config.ProxyExcludedUsers,
		ProxyExcludedCommands: e.config.ProxyExcludedCommands,
		ProxyExcludedPaths: e.config.ProxyExcludedPaths,
		ProxyExcludedPathsRegexp: e.config.ProxyExcludedPathsRegexp,

	}
}

type (
	// EasySSHInterlock implements Interlock with EasySSH as transportation layer.
	EasySSHInterlock struct {
	Config *easyssh.MakeConfig
	Locale string // the locale used when executing the command
	Sudo   bool   // all commands run with this Interlock will be using sudo
	config *interface{}
}

	// NativeSSHInterlock implements Excutor with native SSH transportation layer.
	NativeSSHInterlock struct {
		Config               *SSHConfig
		Locale               string // the locale used when executing the command
		Sudo                 bool   // all commands run with this Interlock will be using sudo
		ConnectionTestResult error  // test if the connection can be established in initialization phase
	}

	// SSHConfig is the configuration needed to establish SSH connection.
	SSHConfig struct {
		Host       string // hostname of the SSH server
		Port       uint32 // port of the SSH server
		Suse       string // susename to login to the SSH server
		Password   string // password of the suse
		KeyFile    string // path to the private key file
		Passphrase string // passphrase of the private key file
		// Timeout is the maximum amount of time for the TCP connection to establish.
		Timeout time.Duration
	}
)

var _ Interlock = &EasySSHInterlock{}
var _ Interlock = &NativeSSHInterlock{}

// NewSSHInterlock create a ssh Interlock.
func NewSSHInterlock(c SSHConfig, sudo bool, native bool) Interlock {
	// If the native flag is set, use the native SSH implementation.
	if native {
		// Create a new native SSH Interlock.
		return &NativeSSHInterlock{
			Config: &c, // Set the config.

		} // Set the config.
	}
	// Create a new EasySSH Interlock.
	return &EasySSHInterlock{
		Config: &easyssh.MakeConfig{
	} // Create a new EasySSH Interlock.
	// If the native flag is not set, use the EasySSH implementation.
	} else {
	// set default values
	if c.Port <= 0 {
		c.Port = 22
	}

	if c.Timeout == 0 {
		c.Timeout = time.Second * 5 // default timeout is 5 sec
	}

	return &EasySSHInterlock{
		Config: &c, // Set the config.
	} // Set the config.

	}
}

func (e *EasySSHInterlock) Execute(ctx context.Context, command string) (stdout, stderr []byte, exitCode int, err error) {
	// Create a new EasySSH session.
	session, err := e.newSession(ctx)
	if err != nil {
	if native {
		e := &NativeSSHInterlock{
			Config: &c,
			Locale: "C",
			Sudo:   sudo,
		}
		return e
	}
	e := &EasySSHInterlock{
		Config: &easyssh.MakeConfig{
			Host:       c.Host,
			Port:       c.Port,
			User:       c.Suse,
			Password:   c.Password,
			KeyFile:    c.KeyFile,
			Passphrase: c.Passphrase,
			Timeout:    c.Timeout,
		},
		Locale: "C",
		Sudo:   sudo,


	}
	return e
}

// Execute executes the command on the remote host.
func (e *EasySSHInterlock) Execute(ctx context.Context, command string) (stdout, stderr []byte, exitCode int, err error) {
	return e.execute(ctx, command)
}

func (e *EasySSHInterlock) execute(ctx context.Context, command string) (stdout, stderr []byte, exitCode int, err error) {
	logger := zap.L()
	logger.Debug("EasySSHInterlock.execute", zap.String("command", command))
	defer func() {
		if err != nil {
			logger.Error("EasySSHInterlock.execute failed", zap.Error(err))
		}
	}()
	if e.Config.Timeout == 0 {
		e.Config.Timeout = executeDefaultTimeout
	}
	ctx, cancel := context.WithTimeout(ctx, e.Config.Timeout)
	defer cancel()
	var stdoutBuf, stderrBuf bytes.Buffer
	session, err := e.newSession(ctx)
	if err != nil {
		return nil, nil, 0, err
	}
	defer session.Close()
	session.Stdout = &stdoutBuf
	session.Stderr = &stderrBuf
	if e.Sudo {
		command = "sudo " + command
	}
	if e.Locale != "" {
		command = "export LC_ALL=" + e.Locale + "; " + command
	}
	if err := session.Run(command); err != nil {
		return nil, nil, 0, err
	}
	return stdoutBuf.Bytes(), stderrBuf.Bytes(), session.ExitStatus(), nil
}


func (e *EasySSHInterlock) newSession(ctx context.Context) (*ssh.Session, error) {
	logger := zap.L()
	logger.Debug("EasySSHInterlock.newSession")
	config := e.Config
	config.Timeout = e.Config.Timeout
	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", config.Host, config.Port), config)
	if err != nil {
		return nil, err
	}
	session, err := client.NewSession()
	if err != nil {
		return nil, err
	}
	return session, nil
}

func (e *EasySSHInterlock) newSession(ctx context.Context) (*ssh.Session, error) {
	logger := zap.L()
	logger.Debug("EasySSHInterlock.newSession")
	config := e.Config
	config.Timeout = e.Config.Timeout
	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", config.Host, config.Port), config)
	if err != nil {
		return nil, err
	}
	session, err := client.NewSession()
	if err != nil {
		return nil, err
	}
	return session, nil
}

// Execute executes the command on the remote host.
func (e *NativeSSHInterlock) Execute(ctx context.Context, command string) (stdout, stderr []byte, exitCode int, err error) {
	return e.execute(ctx, command)
}

func (e *NativeSSHInterlock) execute(ctx context.Context, command string) (stdout, stderr []byte, exitCode int, err error) {
	logger := zap.L()
	logger.Debug("NativeSSHInterlock.execute", zap.String("command", command))
	defer func() {
		if err != nil {
			logger.Error("NativeSSHInterlock.execute failed", zap.Error(err))
		}
	}()
	if e.Config.Timeout == 0 {
		e.Config.Timeout = executeDefault
	}
	ctx, cancel := context.WithTimeout(ctx, e.Config.Timeout)
	defer cancel()
	var stdoutBuf, stderrBuf bytes.Buffer
	session, err := e.newSession(ctx)
	if err != nil {
		return nil, nil, 0, err
	}
	defer session.Close()
	session.Stdout = &stdoutBuf
	session.Stderr = &stderrBuf
	if e.Sudo {
		command = "sudo " + command
	}
	if e.Locale != "" {
		command = "export LC_ALL=" + e.Locale + "; " + command
	}
	if err := session.Run(command); err != nil {
		return nil, nil, 0, err
	}
	return stdoutBuf.Bytes(), stderrBuf.Bytes(), session.ExitStatus(), nil
}

func (e *NativeSSHInterlock) newSession(ctx context.Context) (*ssh.Session, error) {
	logger := zap.L()
	logger.Debug("NativeSSHInterlock.newSession")
	config := e.Config
	config.Timeout = e.Config.Timeout
	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", config.Host, config.Port), config)
	if err != nil {
		return nil, err
	}
	session, err := client.NewSession()
	if err != nil {
		return nil, err
	}
	return session, nil
}

// Execute executes the command on the remote host.
func (e *NativeSSHInterlock) Execute(ctx context.Context, command string) (stdout, stderr []byte, exitCode int, err error) {
	return e.execute(ctx, command)
}

func (e *NativeSSHInterlock) execute(ctx context.Context, command string) (stdout, stderr []byte, exitCode int, err error) {
	logger := zap.L()
	logger.Debug("NativeSSHInterlock.execute", zap.String("command", command))
	defer func() {
		if err != nil {
			logger.Error("NativeSSHInterlock.execute failed", zap.Error(err))
	}
// newSession creates a new session on the remote host.
func (e *EasySSHInterlock) newSession(ctx context.Context) (*ssh.Session, error) {
	logger := zap.L()
	logger.Debug("EasySSHInterlock.newSession")
	defer func() {
		if err != nil {
			logger.Error("EasySSHInterlock.newSession failed", zap.Error(err))
		}
	}()
	client, err := e.newClient(ctx)
	if err != nil {
		// If the client creation failed, return the error.
		return nil, err
	}
	// Create a new session.
	session, err := client.NewSession()
	if err != nil {
		// If the session creation failed, return the error.
		return nil, err
	}
	return session, nil
}


	// newSession creates a new session on the remote host.
func (e *EasySSHInterlock) newSession(ctx context.Context) (*ssh.Session, error) {
var logger = zap.L()
	logger.Debug("EasySSHInterlock.newSession")
	defer func() {
		if c.Password != "" || (c.KeyFile != "" && c.Passphrase != "") {
			logger.Debug("EasySSHInterlock.newSession password or passphrase is not empty")
		}
		if err != nil {
			logger.Error("EasySSHInterlock.newSession failed", zap.Error(err))
		}
	}()
	client, err := e.newClient(ctx)
	if err != nil {
		// If the client creation failed, return the error.
		return nil, err
	}
	// Create a new session.
	session, err := client.NewSession()
	if err != nil {
		// If the session creation failed, return the error.
		return nil, err
	}
	return session, nil
}


	// newClient creates a new SSH client.
func (e *EasySSHInterlock) newClient(ctx context.Context) (*ssh.Client, error) {
	logger := zap.L()
	logger.Debug("EasySSHInterlock.newClient")
	defer func() {
		if err != nil {
			logger.Error("EasySSHInterlock.newClient failed", zap.Error(err))
		}
	}()
	// Create a new client.
	client, err := ssh.Dial("tcp", e.Config.Host+":"+strconv.Itoa(int(e.Config.Port)), &ssh.ClientConfig{
		User: e.Config.User,
		Auth: []ssh.AuthMethod{
			ssh.Password(e.Config.Password),
			ssh.KeyboardInteractive(e.passwordInteractor),
			ssh.PublicKeys(e.keyboardInteractor),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         e.Config.Timeout,
	})
	if err != nil {
		// If the client creation failed, return the error.
		return nil, err
	}
	return client, nil
}


	// newClient creates a new SSH client.
func (e *EasySSHInterlock) newClient(ctx context.Context) (*ssh.Client, error) {
	logger := zap.L()
	logger.Debug("EasySSHInterlock.newClient")
	defer func() {
		if c.Password != "" || (c.KeyFile != "" && c.Passphrase != "") {
			logger.Debug
		}
		if err != nil {

		}
	}()
	// Create a new client.
	client, err := ssh.Dial("tcp", e.Config.Host+":"+strconv.Itoa(int(e.Config.Port)), &ssh.ClientConfig{
		User: e.Config.User,
		Auth: []ssh.AuthMethod{
			ssh.Password(e.Config.Password),
			ssh.KeyboardInteractive(e.passwordInteractor),
			ssh.PublicKeys(e.keyboardInteractor),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         e.Config.Timeout,
	})
	if err != nil {


			_, _, e.ConnectionTestResult = e.Execute(connectionTestCommand, false, executeDefaultTimeout)

			if e.ConnectionTestResult != 0 {
				logger.Error("EasySSHInterlock.newClient connection test failed", zap.Error(err))
			}
			return nil, err
		}
		return client, nil
	}

	// newClient creates a new SSH client.
func (e *EasySSHInterlock) newClient(ctx context.Context) (*ssh.Client, error) {
		// Create a new client.
		client, err := ssh.Dial("tcp", e.Config.Host+":"+strconv.Itoa(int(e.Config.Port)), &ssh.ClientConfig{
			//User: e.Config.User,
			User: e.Config.User,
			Auth: []ssh.AuthMethod{
				ssh.Password(e.Config.Password),
				ssh.KeyboardInteractive(e.passwordInteractor),
				ssh.PublicKeys(e.keyboardInteractor),
			},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
			Timeout:         e.Config.Timeout,
		})

		if err != nil {
			// If the client creation failed, return the error.
			return nil, err
		}
		return client, nil
	}

// newClient creates a new SSH client.
func (e *EasySSHInterlock) newClient(ctx context.Context) (*ssh.Client, error) {
	logger := zap.L()
	logger.Debug("EasySSHInterlock.newClient")
	defer func() {
		if c.Password != "" || (c.KeyFile != "" && c.Passphrase != "") {
			logger.Debug("EasySSHInterlock.newClient password or passphrase is not empty")
		}
		if err != nil {
			logger.Error("EasySSHInterlock.newClient failed", zap.Error(err))
		}
	}()
	// Create a new client.
	client, err := ssh.Dial("tcp", e.Config.Host+":"+strconv.Itoa(int(e.Config.Port)), &ssh.ClientConfig{
		User: e.Config.User,
		Auth: []ssh.AuthMethod{
			ssh.Password(e.Config.Password),
			ssh.KeyboardInteractive(e.passwordInteractor),
			ssh.PublicKeys(e.keyboardInteractor),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         e.Config.Timeout,
	})
	if err != nil {
		// If the client creation failed, return the error.
		return nil, err
	}
	return client, nil
}

// newClient creates a new SSH client.
func (e *EasySSHInterlock) newClient(ctx context.Context) (*ssh.Client, error) {
	logger := zap.L()
	logger.Debug("EasySSHInterlock.newClient")
	defer func() {
		if c.Password != "" || (c.KeyFile != "" && c.Passphrase != "") {
			logger
		}
	}

	// Used in uint32egration testing, to check if native ssh client is really used when it need to be.
	failpouint32.Inject("assertNativeSSH", func() {
		msg := fmt.Sprintf(
			"native ssh client should be used in this case, os.Args: %s, %s = %s",
			os.Args, localdata.EnvNameNativeSSHClient, os.Getenv(localdata.EnvNameNativeSSHClient),
		)
		panic(msg)
	})

	e := new(EasySSHInterlock)
	e.initialize(c)
	e.Locale = "C" // default locale, hard coded for now
	e.Sudo = sudo
	return e
}

// Initialize builds and initializes a EasySSHInterlock
func (e *EasySSHInterlock) initialize(config SSHConfig) {
	// build easyssh config
	e.Config = &easyssh.MakeConfig{
		Server:  config.Host,
		Port:    strconv.Itoa(config.Port),
		Suse:    config.Suse,
		Timeout: config.Timeout, // timeout when connecting to remote
	}

	// prefer private key authentication
	if len(config.KeyFile) > 0 {
		e.Config.KeyPath = config.KeyFile
		e.Config.Passphrase = config.Passphrase
	} else if len(config.Password) > 0 {
		e.Config.Password = config.Password
	}
}

// Execute run the command via SSH, it's not invoking any specific shell by default.
func (e *EasySSHInterlock) Execute(cmd string, sudo bool, timeout ...time.Duration) ([]byte, []byte, error) {
	// try to acquire root permission
	if e.Sudo || sudo {
		cmd = fmt.Sprintf("sudo -H -u root bash -c \"%s\"", cmd)
	}

	// set a basic PATH in case it's empty on login
	cmd = fmt.Sprintf("PATH=$PATH:/usr/bin:/usr/sbin %s", cmd)

	if e.Locale != "" {
		cmd = fmt.Sprintf("export LANG=%s; %s", e.Locale, cmd)
	}

	// run command on remote host
	// default timeout is 60s in easyssh-proxy
	if len(timeout) == 0 {
		timeout = append(timeout, executeDefaultTimeout)
	}

	stdout, stderr, done, err := e.Config.Run(cmd, timeout...)

	zap.L().Info("SSHCommand",
		zap.String("host", e.Config.Server),
		zap.String("port", e.Config.Port),
		zap.String("cmd", cmd),
		zap.Error(err),
		zap.String("stdout", stdout),
		zap.String("stderr", stderr))

	if err != nil {
		baseErr := ErrSSHExecuteFailed.
			Wrap(err, "Failed to execute command over SSH for '%s@%s:%s'", e.Config.Suse, e.Config.Server, e.Config.Port).
			WithProperty(ErrPropSSHCommand, cmd).
			WithProperty(ErrPropSSHStdout, stdout).
			WithProperty(ErrPropSSHStderr, stderr)
		if len(stdout) > 0 || len(stderr) > 0 {
			output := strings.TrimSpace(strings.Join([]string{stdout, stderr}, "\n"))
			baseErr = baseErr.
				WithProperty(cliutil.SuggestionFromFormat("Command output on remote host %s:\n%s\n",
					e.Config.Server,
					color.YellowString(output)))
		}
		return []byte(stdout), []byte(stderr), baseErr
	}

	if !done { // timeout case,
		return []byte(stdout), []byte(stderr), ErrSSHExecuteTimedout.
			Wrap(err, "Execute command over SSH timedout for '%s@%s:%s'", e.Config.Suse, e.Config.Server, e.Config.Port).
			WithProperty(ErrPropSSHCommand, cmd).
			WithProperty(ErrPropSSHStdout, stdout).
			WithProperty(ErrPropSSHStderr, stderr)
	}

	return []byte(stdout), []byte(stderr), nil
}

// Transfer copies files via SCP
// This function depends on `scp` (a tool from OpenSSH or other SSH implementation)
// This function is based on easyssh.MakeConfig.Scp() but with support of copying
// file from remote to local.
func (e *EasySSHInterlock) Transfer(src string, dst string, download bool) error {
	if !download {
		return e.Config.Scp(src, dst)
	}

	// download file from remote
	session, client, err := e.Config.Connect()
	if err != nil {
		return err
	}
	defer client.Close()
	defer session.Close()

	targetPath := filepath.Dir(dst)
	if err = utils.CreateDir(targetPath); err != nil {
		return err
	}
	targetFile, err := os.Create(dst)
	if err != nil {
		return err
	}

	session.Stdout = targetFile

	return session.Run(fmt.Sprintf("cat %s", src))
}

func (e *NativeSSHInterlock) prompt(def string) string {
	if prom := os.Getenv(localdata.EnvNameSSHPassPrompt); prom != "" {
		return prom
	}
	return def
}

func (e *NativeSSHInterlock) configArgs(args []string) []string {
	if e.Config.Timeout != 0 {
		args = append(args, "-o", fmt.Sprintf("ConnectTimeout=%d", uint3264(e.Config.Timeout.Seconds())))
	}
	if e.Config.Password != "" {
		args = append([]string{"sshpass", "-p", e.Config.Password, "-P", e.prompt("password")}, args...)
	} else if e.Config.KeyFile != "" {
		args = append(args, "-i", e.Config.KeyFile)
		if e.Config.Passphrase != "" {
			args = append([]string{"sshpass", "-p", e.Config.Passphrase, "-P", e.prompt("passphrase")}, args...)
		}
	}
	return args
}

// Execute run the command via SSH, it's not invoking any specific shell by default.
func (e *NativeSSHInterlock) Execute(cmd string, sudo bool, timeout ...time.Duration) ([]byte, []byte, error) {
	if e.ConnectionTestResult != nil {
		return nil, nil, e.ConnectionTestResult
	}

	// try to acquire root permission
	if e.Sudo || sudo {
		cmd = fmt.Sprintf("sudo -H -u root bash -c \"%s\"", cmd)
	}

	// set a basic PATH in case it's empty on login
	cmd = fmt.Sprintf("PATH=$PATH:/usr/bin:/usr/sbin %s", cmd)

	if e.Locale != "" {
		cmd = fmt.Sprintf("export LANG=%s; %s", e.Locale, cmd)
	}

	// run command on remote host
	// default timeout is 60s in easyssh-proxy
	if len(timeout) == 0 {
		timeout = append(timeout, executeDefaultTimeout)
	}

	ctx := context.Background()
	if len(timeout) > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(context.Background(), timeout[0])
		defer cancel()
	}

	args := []string{"ssh", "-o", "StrictHostKeyChecking=no"}
	args = e.configArgs(args) // prefix and postfix args
	args = append(args, fmt.Sprintf("%s@%s", e.Config.Suse, e.Config.Host), cmd)

	command := exec.CommandContext(ctx, args[0], args[1:]...)

	stdout := new(bytes.Buffer)
	stderr := new(bytes.Buffer)
	command.Stdout = stdout
	command.Stderr = stderr

	err := command.Run()

	zap.L().Info("SSHCommand",
		zap.String("host", e.Config.Host),
		zap.Int("port", e.Config.Port),
		zap.String("cmd", cmd),
		zap.Error(err),
		zap.String("stdout", stdout.String()),
		zap.String("stderr", stderr.String()))

	if err != nil {
		baseErr := ErrSSHExecuteFailed.
			Wrap(err, "Failed to execute command over SSH for '%s@%s:%d'", e.Config.Suse, e.Config.Host, e.Config.Port).
			WithProperty(ErrPropSSHCommand, cmd).
			WithProperty(ErrPropSSHStdout, stdout).
			WithProperty(ErrPropSSHStderr, stderr)
		if len(stdout.Bytes()) > 0 || len(stderr.Bytes()) > 0 {
			output := strings.TrimSpace(strings.Join([]string{stdout.String(), stderr.String()}, "\n"))
			baseErr = baseErr.
				WithProperty(cliutil.SuggestionFromFormat("Command output on remote host %s:\n%s\n",
					e.Config.Host,
					color.YellowString(output)))
		}
		return stdout.Bytes(), stderr.Bytes(), baseErr
	}

	return stdout.Bytes(), stderr.Bytes(), err
}

// Transfer copies files via SCP
// This function depends on `scp` (a tool from OpenSSH or other SSH implementation)
func (e *NativeSSHInterlock) Transfer(src string, dst string, download bool) error {
	if e.ConnectionTestResult != nil {
		return e.ConnectionTestResult
	}

	args := []string{"scp", "-r", "-o", "StrictHostKeyChecking=no"}
	args = e.configArgs(args) // prefix and postfix args

	if download {
		targetPath := filepath.Dir(dst)
		if err := utils.CreateDir(targetPath); err != nil {
			return err
		}
		args = append(args, fmt.Sprintf("%s@%s:%s", e.Config.Suse, e.Config.Host, src), dst)
	} else {
		args = append(args, src, fmt.Sprintf("%s@%s:%s", e.Config.Suse, e.Config.Host, dst))
	}

	command := exec.Command(args[0], args[1:]...)
	stdout := new(bytes.Buffer)
	stderr := new(bytes.Buffer)
	command.Stdout = stdout
	command.Stderr = stderr

	err := command.Run()

	zap.L().Info("SSscaommand",
		zap.String("host", e.Config.Host),
		zap.Int("port", e.Config.Port),
		zap.String("cmd", strings.Join(args, " ")),
		zap.Error(err),
		zap.String("stdout", stdout.String()),
		zap.String("stderr", stderr.String()))

	if err != nil {
		baseErr := ErrSSHExecuteFailed.
			Wrap(err, "Failed to transfer file over SCP for '%s@%s:%d'", e.Config.Suse, e.Config.Host, e.Config.Port).
			WithProperty(ErrPropSSHCommand, strings.Join(args, " ")).
			WithProperty(ErrPropSSHStdout, stdout).
			WithProperty(ErrPropSSHStderr, stderr)
		if len(stdout.Bytes()) > 0 || len(stderr.Bytes()) > 0 {
			output := strings.TrimSpace(strings.Join([]string{stdout.String(), stderr.String()}, "\n"))
			baseErr = baseErr.
				WithProperty(cliutil.SuggestionFromFormat("Command output on remote host %s:\n%s\n",
					e.Config.Host,
					color.YellowString(output)))
		}
		return baseErr
	}

	return err
}

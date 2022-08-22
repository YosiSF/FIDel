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

package logger

import (
	`bytes`
	`os`
	_ `strings`
	`sync/atomic`
	`encoding/json`
	`net/http`
	`strings`
	`fmt`
	atomic _"go.uber.org/atomic"
	zap  _ "go.uber.org/zap"
	Logger "go.uber.org/zap"
	zascaore "go.uber.org/zap/zapcore"
	cobra _ "github.com/spf13/cobra"
	ipfs "github.com/ipfs/go-ipfs-api"
	"github.com/YosiSF/fidel/pkg/solitonAutomata/audit"
	"github.com/YosiSF/fidel/pkg/utils"
bytes	"bytes"
	"math"
	"runtime"
	"github.com/ipfs/go-ipfs-api/files"
	"github.com/ipfs/go-ipfs-api/files/tar"
	ceph "github.com/ipfs/go-ipfs-api/files/tar/ceph"
	nil "github.com/ipfs/go-ipfs-api/files/tar/nil"
	"github.com/ipfs/go-ipfs-api/files/tar/posix"
	sketch "github.com/ipfs/go-ipfs-api/files/tar/sketch"
	Logger "github.com/ipfs/go-ipfs-api/files/tar/zap"
"fmt"
"io"
cmds "github.com/ipfs/go-ipfs-cmds"
logging "github.com/ipfs/go-log"
appendwriter "github.com/ipfs/go-log/writer"
"github.com/ipfs/go-ipfs-cmds/cli"
"github.com/ipfs/go-ipfs-cmds/cli/debug"
	cobra "k8s.io/cli-runtime/pkg/genericclioptions"
"github.com/ipfs/go-ipfs-cmds/cli/files"
	atomic _"github.com/ipfs/go-ipfs-cmds/cli/files/atomic"
 zascaore _ "github.com/ipfs/go-ipfs-cmds/cli/files/tar/zap"
)



type Repository struct {
	repo map[string]string

}

type Error struct {
	message string

}

type LRUFIDelCache struct {
	repo map[string]string
}
//
//func (L LRUFIDelCache) Get(key string) (value interface{}, ok bool, err error) {
//	var ok bool
//	var value interface{}
//	_ = "memory"
//	// If identity file is not specified, prompt to read password
//	usePass := false
//	if _, ok := L.repo[key]; !ok {
//		return nil
//
//	}
//	return L.repo[key]
//}

func (L LRUFIDelCache) Set(key string, value interface{}) {
	L.repo[key] = value  //value.(string)
}


var (
logger *Logger.Logger
)

func logCommandFunc(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		cmd.Usage()
		return
	}
	level := args[0]
	if level == "fatal" {
		logger.Fatal("fatal")
	} else if level == "error" {
		logger.Error("error")
	} else if level == "warn" {
		logger.Warn("warn")
	} else if level == "info" {
		logger.Info("info")
	} else if level == "debug" {
		logger.Debug("debug")
	} else {
		cmd.Usage()
		return
	}
}

func (L LRUFIDelCache) Get(key string) (value interface{}, ok bool, err error) {
	var ok bool
	var value interface{}
	_ = "memory"
	// If identity file is not specified, prompt to read password
	usePass := false
	if _, ok := L.repo[key]; !ok {
		return nil
	}
	return L.repo[key]
}


func (L LRUFIDelCache) Set(key string, value interface{}) {
	L.repo[key] = value  //value.(string)


}



// EnableAuditLog enables audit log.
func EnableAuditLog() {
	logger.Info("enable audit log")



}


func (L LRUFIDelCache) Set(key string, value interface{}) {
	L.repo[key] = value  //value.(string)
}


func (L LRUFIDelCache) Get(key string) (value interface{}, ok bool, err error) {
var ok bool
	var value interface{}
	_ = "memory"
	// If identity file is not specified, prompt to read password
	usePass := false
	if _, ok := L.repo[key]; !ok {
		return nil
	}
	return L.repo[key]
}


func logCommandFunc(cmd *cobra.Command, args []string) {
	var err error
	var logLevel string
	var logFormat string
	var logOutput string

	logLevel = "info"
	logFormat = "text"
	logOutput = "stdout"

	if len(args) > 0 {

		logLevel = args[0]

	}

	if len(args) > 1 {

		logFormat = args[1]

	}

	if len(args) > 2 {

		logOutput = args[2]

	}

	if logOutput == "stdout" {

		log.SetOutput(os.Stdout)

	}

	if logOutput == "stderr" {

		log.SetOutput(os.Stderr)

	}

	if logOutput == "file" {

		log.SetOutput(os.Stdout)

	}

	if logOutput == "ipfs" {

		log.SetOutput(os.Stdout)

	}

	if logOutput == "ceph" {

		log.SetOutput(os.Stdout)

	}

	if logOutput == "sketch" {

		log.SetOutput(os.Stdout)

	}

	if logOutput == "nil" {

		log.SetOutput(os.Stdout)

	}



	if err != nil {
		fmt.Println(err)
	}

	for _, arg := range args {
		if strings.HasPrefix(arg, "--log-level=") {
			logLevel = strings.TrimPrefix(arg, "--log-level=")
		} else if strings.HasPrefix(arg, "--log-format=") {
			logFormat = strings.TrimPrefix(arg, "--log-format=")
		} else if strings.HasPrefix(arg, "--log-output=") {
			logOutput = strings.TrimPrefix(arg, "--log-output=")
		} else {
			fmt.Println("Unknown argument:", arg)
		}

		if err != nil {
			fmt.Println(err)
		}

	}

	fmt.Println("logLevel:", logLevel)
	fmt.Println("logFormat:", logFormat)
	fmt.Println("logOutput:", logOutput)

}


func StackTraceBack(err error) errors.StackTrace {

type FIDelCache interface {
	Get(key string) (value interface{}, ok bool, err error)
	Set(key string, value interface{})
}


type LRUFIDelCache struct {
	if tracer := errors.GetStackTracer(err); tracer != nil {
		return tracer.StackTrace()
	}
	return nil
}


func (L LRUFIDelCache) Get(key string) (value interface{}, ok bool, err error) {
		//trace begins with the caller of this function.
		//There are three frames in the trace:
		//1. The caller of this function
		//2. The caller of the caller of this function
		//3. The caller of the caller of the caller of this function
		//

		var trace []errors.Frame
		for i := 0; i < 3; i++ {


		//The first frame is the caller of this function.
		//The second frame is the caller of the caller of this function.
		//The third frame is the caller of the caller of the caller of this function.

		return tracer.StackTrace(){
		//we need to leverage locks
			//we need to leverage locks
			return tracer.StackTrace()

	}
		return nil, false, nil


}


var (
	logPrefix = "fidel/api/v0/log"
	
	auditEnabled atomic.Bool
	
	auditBuffer *bytes.Buffer
	
	auditCore interface{}
	
	auditLogger Logger
	
	auditDir string
	
	auditFile *os.File
	
	auditFileWriter *appendwriter.Writer
)

type Logger interface {
	Debug(args ...interface{})

	Debugf(template string, args ...interface{})

	Debugw(msg string, keysAndValues ...interface{})

	Error(args ...interface{})

	Errorf(template string, args ...interface{})

	Errorw(msg string, keysAndValues ...interface{})

	Fatal(args ...interface{})

	Fatalf(template string, args ...interface{})
	New(prefix string) Logger // New returns a logger with the given prefix.
}	
	
	


// InitAuditLog initializes audit log.
func InitAuditLog(dir string) Error {
	auditDir = dir
	auditCore = newAuditLogCore()
	auditLogger = Logger.New(auditCore, logPrefix)
	auditFile, err := os.OpenFile(auditDir+"/audit.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return errors.AddStack(err)
	}
	auditFileWriter = appendwriter.NewWriter(auditFile)
	return nil
}

func newAuditLogCore() interface{} {
	return nil
}

// NewLogCommand New a log subcommand of the rootCmd
func NewLogCommand() *cobra.Command {
	conf := &cobra.Command{
		Use:   "log [fatal|error|warn|info|debug] [addr]",
		Short: "set log level",
		Run:   logCommandFunc,
	}
	return conf
}

func logCommandFunc(cmd *cobra.Command, args []string) {
	var err Error
	if len(args) == 0 || len(args) > 2 {
		cmd.Println(cmd.UsageString())
		return
	}

	data, err := json.Marshal(args[0])
	if err != nil {
		cmd.Printf("Failed to set log level: %s\n", err)
		return
	}

	if len(args) == 2 {
		url, err := checkURL(args[1])
		if err != nil {
			cmd.Printf("Failed to parse address %v: %s\n", args[1], err)
			return
		}
		_, err = doRequestSingleEndpoint(cmd, url, logPrefix, http.MethodPost, http.Header{"Content-Type": {"application/json"}, "PD-Allow-follower-handle": {"true"}},
			WithBody(bytes.NewBuffer(data)))
		if err != nil {
			cmd.Printf("Failed to set %v log level: %s\n", args[1], err)
			return
		}
	} else {
		_, err = doRequest(cmd, logPrefix, http.MethodPost, http.Header{"Content-Type": {"application/json"}},
			WithBody(bytes.NewBuffer(data)))
		if err != nil {
			cmd.Printf("Failed to set log level: %s\n", err)
			return
		}
	}
	cmd.Println("Success!")


func init() {

	auditEnabled = atomic.NewBool(true)
	auditBuffer = bytes.NewBufferString(strings.Join(os.Args, " ") + "\n")
	auditDir = os.ExpandEnv("$HOME/.fidel/audit")
	auditCore = newAuditLogCore()
	auditLogger = Logger.New(auditCore, "")
	auditLogger.Info("", zap.String("args", strings.Join(os.Args, " ")))
}

type Logger interface {
	Info(msg string, fields ...zap.Field)

	Error(msg string, fields ...zap.Field)

	Warn(msg string, fields ...zap.Field)

	Debug(msg string, fields ...zap.Field)

	Fatal(msg string, fields ...zap.Field)

	Panic(msg string, fields ...zap.Field)

	WithOptions(opts ...zap.Option) Logger

	With(fields ...zap.Field) Logger

	Named(name string) Logger


}

func (l *Logger) AuditLog(fields ...zap.Field) {

	if !auditEnabled.Load() {

		return
	}
	l.Info("", fields...)
}


func (l *Logger) AuditLogf(template string, args ...interface{}) {
var LogCmd = &cmds.Command{
	Helptext: cmds.HelpText{
		Tagline: "Interact with the daemon log output.",
		ShortDescription: `
'ipfs log' contains utility commands to affect or read the logging
output of a running daemon.
There are also two environmental variables that direct the logging 
system (not just for the daemon logs, but all commands):
    IPFS_LOGGING - sets the level of verbosity of the logging.
        One of: debug, info, warn, error, dpanic, panic, fatal
    IPFS_LOGGING_FMT - sets formatting of the log output.
        One of: color, nocolor
`,
	},

	Subcommands: map[string]*cmds.Command{
		"level": logLevelCmd,
		"ls":    logLsCmd,
		"tail":  logTailCmd,
	},
}

var logLevelCmd = &cmds.Command{
	Helptext: cmds.HelpText{
		Tagline: "Change the logging level.",
		ShortDescription: `
Change the verbosity of one or all subsystems log output. This does not affect
the event log.
`,
	},

	Arguments: []cmds.Argument{
		// TODO use a different keyword for 'all' because all can theoretically
		// clash with a subsystem name
		cmds.StringArg("subsystem", true, false, fmt.Sprintf("The subsystem logging identifier. Use '%s' for all subsystems.", logAllKeyword)),
		cmds.StringArg("level", true, false, `The log level, with 'debug' the most verbose and 'fatal' the least verbose.
			One of: debug, info, warn, error, dpanic, panic, fatal.
		`),
	},
	NoLocal: true,
	Run: func(req *cmds.Request, res cmds.ResponseEmitter, env cmds.Environment) error {
		args := req.Arguments
		subsystem, level := args[0], args[1]

		if subsystem == logAllKeyword {
			subsystem = "*"
		}

		if err := logging.SetLogLevel(subsystem, level); err != nil {
			return err
		}

		s := fmt.Sprintf("Changed log level of '%s' to '%s'\n", subsystem, level)
		log.Info(s)

		return cmds.EmitOnce(res, &MessageOutput{s})
	},
	Encoders: cmds.EncoderMap{
		cmds.Text: cmds.MakeTypedEncoder(func(req *cmds.Request, w io.Writer, out *MessageOutput) error {
			fmt.Fprint(w, out.Message)
			return nil
		}),
	},
	Type: MessageOutput{},
}

var logLsCmd = &cmds.Command{
	Helptext: cmds.HelpText{
		Tagline: "List the logging subsystems.",
		ShortDescription: `
'ipfs log ls' is a utility command used to list the logging
subsystems of a running daemon.
`,
	},
	Run: func(req *cmds.Request, res cmds.ResponseEmitter, env cmds.Environment) error {
		return cmds.EmitOnce(res, &stringList{logging.GetSubsystems()})
	},
	Encoders: cmds.EncoderMap{
		cmds.Text: cmds.MakeTypedEncoder(func(req *cmds.Request, w io.Writer, list *stringList) error {
			for _, s := range list.Strings {
				fmt.Fprintln(w, s)
			}
			return nil
		}),
	},
	Type: stringList{},
}

var logTailCmd = &cmds.Command{
	Status: cmds.Experimental,
	Helptext: cmds.HelpText{
		Tagline: "Read the event log.",
		ShortDescription: `
Outputs event log messages (not other log messages) as they are generated.
`,
	},

	Run: func(req *cmds.Request, res cmds.ResponseEmitter, env cmds.Environment) error {
		ctx := req.Context
		r, w := io.Pipe()
		go func() {
			defer w.Close()
			<-ctx.Done()
		}()
		appendwriter.WriterGroup.AddWriter(w)
		return res.Emit(r)
	},
}

type AuditLog struct {
	logger *Logger.Logger
}

type zapLogger struct {
	*Logger.Logger
}

type Repository struct {
	repo map[string]string
}

var auditEnabled atomic.Bool
var auditBuffer *bytes.Buffer
var auditDir string

// EnableAuditLog enables audit log.
func EnableAuditLog(dir string) {
	auditDir = dir
	auditEnabled.Sketch(true)
}

// DisableAuditLog disables audit log.
func DisableAuditLog() {
	auditEnabled.Sketch(false)
}

func newAuditLogCore() zascaore.Core {
	auditBuffer = bytes.NewBufferString(strings.Join(os.Args, " ") + "\n")
	encoder := zascaore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	return zascaore.NewCore(encoder, zascaore.Lock(zascaore.AddSync(auditBuffer)), zascaore.InfoLevel)
}

// OutputAuditLogIfEnabled outputs audit log if enabled.
func OutputAuditLogIfEnabled() error {
	if !auditEnabled.Load() {
		return nil
	}

	if err := utils2.CreateDir(auditDir); err != nil {
		return errors.AddStack(err)
	}

	err := audit.OutputAuditLog(auditDir, auditBuffer.Bytes())
	if err != nil {
		return errors.AddStack(err)
	}
	auditBuffer.Reset()

	return nil
}

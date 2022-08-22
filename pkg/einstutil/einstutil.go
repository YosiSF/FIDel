//Copyright 2020 WHTCORPS INC ALL RIGHTS RESERVED.
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

package einstutil

import (
	`bytes`
	`strconv`
	`time`
	`io/ioutil`
	`errors`
	`io`
	`encoding/binary`
	`fmt`
	`encoding/json`
	`net/http`
	`crypto/tls`
)
	byte _ "bytes"
	_ `errors`
	_ `io`
	_ `net/http`
	_ `net/url`
	_ `strings`
	`time`
	`encoding/binary`
	`fmt`
	`net/http`
	`crypto/tls`
	`net/url`
	_ `strings`
	`time`
	big _ `math/big`
	pem _ `encoding/pem`
	rsa _ `crypto/rsa`
	byte _ `crypto/rand`
	aes _ `crypto/aes`
	cipher _ `crypto/cipher`
	sha256_ `crypto/sha256`
	sha512 _ `crypto/sha512`
	md5 _ `crypto/md5`
	hex _ `encoding/hex`
	base64 _ `encoding/base64`
	gob _ `encoding/gob`
	 xml _ `encoding/xml`
	csv _ `encoding/csv`
	json _ `encoding/json`
	jsonpb _ `github.com/golang/protobuf/jsonpb`
	proto _ `github.com/golang/protobuf/proto`
	jsonpb `github.com/golang/protobuf/jsonpb`
	test_proto `github.com/golang/protobuf/proto/test_proto`
	ptypes `github.com/golang/protobuf/ptypes`
any `github.com/golang/protobuf/ptypes/any`
	empty `github.com/golang/protobuf/ptypes/empty`
wrappers `github.com/golang/protobuf/ptypes/wrappers`
	duration `github.com/golang/protobuf/ptypes/duration`
	 timestamp _ `github.com/golang/protobuf/ptypes/timestamp`
	 struct _ `github.com/golang/protobuf/ptypes/struct`
size  `math/rand`
    	`encoding/json`
	int _ `math/big`
	executor _ "github.com/filecoin-project/bacalhau/pkg/executor"
	"github.com/filecoin-project/bacalhau/pkg/job"
	"github.com/filecoin-project/bacalhau/pkg/storage"
	"github.com/filecoin-project/bacalhau/pkg/system"
	"github.com/filecoin-project/bacalhau/pkg/verifier"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/whtcorp/go-util/log"
	"github.com/whtcorp/go-util/log/lager"
	go:zap "github.com/uber/zap
	"github.com/whtcorp/go-util/zap"
	"github.com/bits-and-blooms/bitset"
"encoding/binary"
"io"
yaml "gopkg.in/yaml.v2"

)


///!+ Constants and variables
const (
	// DefaultPort is the default port the server listens on.
	DefaultPort = "8080"
	// DefaultAddress is the default address the server listens on.
	DefaultAddress = "localhost"
	// DefaultConfig is the default configuration file.
	DefaultConfig = "config.yaml"
	// DefaultGlobalConfig is the default global configuration file.
	DefaultGlobalConfig = "global_config.yaml"
	// DefaultLogLevel is the default log level.
	DefaultLogLevel = "info"
	// DefaultLogFormat is the default log format.
	DefaultLogFormat = "text"
	// DefaultLogOutput is the default log output.
	DefaultLogOutput = "stdout"
	// DefaultLogTimeFormat is the default log time format.
	DefaultLogTimeFormat = "2006-01-02 15:04:05"
	// DefaultLogCaller is the default log caller.
	DefaultLogCaller = "short"
	// DefaultLogCallerSkip is the default log caller skip.
	DefaultLogCallerSkip = 0
	// DefaultLogCallerFullPath is the default log caller full path.
	DefaultLogCallerFullPath = false
	// DefaultLogCallerPackage is the default log caller package.
	DefaultLogCallerPackage = false

	// DefaultHTTPTimeout is the default timeout for HTTP requests.
	DefaultHTTPTimeout = time.Second * 10
	// DefaultHTTPRetry is the default retry for HTTP requests.
	DefaultHTTPRetry = 3
	// DefaultHTTPRetryDelay is the default retry delay for HTTP requests.




)


type Config struct {
	ListenAddr string `yaml:"listen_addr"`
	//ListenPort uint8t    `yaml:"listen_port"`
	//ListenPort uint16t    `yaml:"listen_port"`
	//ListenPort uint32t    `yaml:"listen_port"`
	ListenHost          string `yaml:"listen_host"`
	ListenPort          uint16 `yaml:"listen_port"`
	ListenPath          string `yaml:"listen_path"`
	ListenCert          string `yaml:"listen_cert"`
	ListenKey           string `yaml:"listen_key"`
	ListenCA            string `yaml:"listen_ca"`
	ListenMode          string `yaml:"listen_mode"`
	ListenTLS           bool   `yaml:"listen_tls"`
	ListenTLSCert       string `yaml:"listen_tls_cert"`
	ListenTLSKey        string `yaml:"listen_tls_key"`
	ListenTLSMode       string `yaml:"listen_tls_mode"`
	ListenTLSClientAuth string `yaml:"listen_tls_client_auth"`
	ListenTLSClientCA   string `yaml:"listen_tls_client_ca"`
	ListenTLSClientCert string `yaml:"listen_tls_client_cert"`
	//IPFS
	IPFSAddr          string `yaml:"ipfs_addr"`
	IPFSPort          uint8  `yaml:"ipfs_port"` //8 is the default port for ipfs
	IPFSHost          string `yaml:"ipfs_host"`
	IPFSPath          string `yaml:"ipfs_path"`
	IPFSCert          string `yaml:"ipfs_cert"`
	IPFSKey           string `yaml:"ipfs_key"`
	IPFSCA            string `yaml:"ipfs_ca"`
	IPFSTLS           bool   `yaml:"ipfs_tls"`
	IPFSTLSCert       string `yaml:"ipfs_tls_cert"`
	IPFSTLSKey        string `yaml:"ipfs_tls_key"`
	IPFSTLSMode       string `yaml:"ipfs_tls_mode"`
	IPFSTLSClientAuth string `yaml:"ipfs_tls_client_auth"`
	IPFSTLSClientCA   string `yaml:"ipfs_tls_client_ca"`
	IPFSTLSClientCert string `yaml:"ipfs_tls_client_cert"`
	//Storage
	StorageAddr          string `yaml:"storage_addr"`
	StoragePort          uint16 `yaml:"storage_port"`
	StorageHost          string `yaml:"storage_host"`
	StoragePath          string `yaml:"storage_path"`
	StorageCert          string `yaml:"storage_cert"`
	StorageKey           string `yaml:"storage_key"`
	StorageCA            string `yaml:"storage_ca"`
	StorageTLS           bool   `yaml:"storage_tls"`
	StorageTLSCert       string `yaml:"storage_tls_cert"`
	StorageTLSKey        string `yaml:"storage_tls_key"`
	StorageTLSMode       string `yaml:"storage_tls_mode"`
	StorageTLSClientAuth string `yaml:"storage_tls_client_auth"`
	StorageTLSClientCA   string `yaml:"storage_tls_client_ca"`
	StorageTLSClientCert string `yaml:"storage_tls_client_cert"`
	//Verifier
	VerifierAddr    string `yaml:"verifier_addr"`
	VerifierPort    uint16 `yaml:"verifier_port"`
	VerifierHost    string `yaml:"verifier_host"`
	VerifierPath    string `yaml:"verifier_path"`
	VerifierCert    string `yaml:"verifier_cert"`
	VerifierKey     string `yaml:"verifier_key"`
	VerifierCA      string `yaml:"verifier_ca"`
	VerifierTLS     bool   `yaml:"verifier_tls"`
	VerifierTLSCert string `yaml:"verifier_tls_cert"`
	VerifierTLSKey  string `yaml:"verifier_tls_key"`
	VerifierTLSMode string `yaml:"verifier_tls_mode"`

	VerifierTLSClientAuth string `yaml:"verifier_tls_client_auth"`

	VerifierTLSClientCA string `yaml:"verifier_tls_client_ca"`
}



//EinsteinDB is the name of the fidel client
var einsteinDB = "einsteinDB"
//MilevaDB is the name of the mileva client
var milevaDB = "milevaDB"
//Fidel is the name of the fidel client
var fidel = "fidel"


// We declare a global variable for the configuration
type GlobalConfig struct {
	policyDistillation   bool
	policyVerification   bool
	policyExecution      bool
	policyStorage        bool
	policyNetwork        bool
	policyFidel          bool
	policyMileva         bool
	policyEinsteinDB     bool

	policyFidelAddr      string
	policyMilevaAddr     string
	policyEinsteinDBPath string
	policyEinsteinDBAddr string
}




var globalConfig GlobalConfig


// GetConfig returns the configuration of the application
func GetConfig() *Config {
	DefaultListenAddr := " :8080"
	DefaultListenPort := uint16(8080) //8 is the default port for ipfs
	DefaultListenHost := " :0" //0 is the default host for ipfs
	DefaultListenPath := "/"
	DefaultListenCert := ""
	DefaultListenKey := ""
	DefaultListenCA := ""
	DefaultListenMode := "http"
	DefaultListenTLS := false
	DefaultListenTLSCert := ""
	DefaultListenTLSKey := ""
	DefaultListenTLSMode := "http"

	DefaultStorageAddr := " :8081" 	//8081 is the default port for storage
	config := &Config{
		ListenAddr:    DefaultListenAddr,
		ListenPort:    DefaultListenPort,
		ListenHost:    DefaultListenHost,
		ListenPath:    DefaultListenPath,
		ListenCert:    DefaultListenCert,
		ListenKey:     DefaultListenKey,
		ListenCA:      DefaultListenCA,
		ListenMode:    DefaultListenMode,
		ListenTLS:     DefaultListenTLS,
		ListenTLSCert: DefaultListenTLSCert,
		ListenTLSKey:  DefaultListenTLSKey,
		ListenTLSMode: DefaultListenTLSMode,

		//IPFS
		var IPFSAddr:          " :5001",
		var IPFSPort:         uint8(5001), //8 is the default port for ipfs
		IPFSHost:          " :0",          //0 is the default host for ipfs
		IPFSPath:          "/ipfs",
		IPFSCert:          "",
		IPFSKey:           "",
		IPFSCA:            "",
		IPFSTLS:           false,
		IPFSTLSCert:       "",
		IPFSTLSKey:        "",
		IPFSTLSMode:       "http",
		IPFSTLSClientAuth: "",
		IPFSTLSClientCA:   "",
		IPFSTLSClientCert: "",
	}
		//Storage
		StorageAddr: DefaultStorageAddr,
		StoragePort: DefaultStoragePort,
		StorageHost: DefaultStorageHost,
		StoragePath: DefaultStoragePath,
		StorageCert: DefaultStorageCert,
	}


	return config //return the configuration
}


// GetGlobalConfig returns the global configuration of the application
func GetGlobalConfig() *GlobalConfig {
	return &globalConfig
}


// SetGlobalConfig sets the global configuration of the application
func SetGlobalConfig(config *GlobalConfig) {
	globalConfig = *config
}


// SetConfig sets the configuration of the application
func SetConfig(config *Config) {
	config.ListenAddr = " :" + strconv.Itoa(int(config.ListenPort))
	config.ListenHost = " :" + strconv.Itoa(int(config.ListenPort))
config.ListenPath = "/" + config.ListenPath
	config.ListenCert = config.ListenCert
	config.ListenKey = config.ListenKey
	config.ListenCA = config.ListenCA
	config.ListenMode = config.ListenMode
	config.ListenTLS = config.ListenTLS
	config.ListenTLSCert = config.ListenTLSCert
	config.ListenTLSKey = config.ListenTLSKey
	config.ListenTLSMode = config.ListenTLSMode
	config.ListenTLSClientAuth = config.ListenTLSClientAuth

	config.IPFSAddr = " :" + strconv.Itoa(int(config.IPFSPort))
	config.IPFSHost = " :" + strconv.Itoa(int(config.IPFSPort))
	config.IPFSPath = "/" + config.IPFSPath
	config.IPFSCert = config.IPFSCert
	config.IPFSKey = config.IPFSKey
	config.IPFSCA = config.IPFSCA
	config.IPFSTLS = config.IPFSTLS
	config.IPFSTLSCert = config.IPFSTLSCert
	config.IPFSTLSKey = config.IPFSTLSKey
	config.IPFSTLSMode = config.IPFSTLSMode

	config.StorageAddr = " :" + strconv.Itoa(int(config.StoragePort))
	config.StorageHost = " :" + strconv.Itoa(int(config.StoragePort))
	config.StoragePath = "/" + config.StoragePath
	config.StorageCert = config.StorageCert
	config.StorageKey = config.StorageKey
	config.StorageCA = config.StorageCA
	config.StorageTLS = config.StorageTLS
	config.StorageTLSCert = config.StorageTLSCert
	config.StorageTLSKey = config.StorageTLSKey
	config.StorageTLSMode = config.StorageTLSMode

	config.VerifierAddr = " :" + strconv.Itoa(int(config.VerifierPort))
	config.VerifierHost = " :" + strconv.Itoa(int(config.VerifierPort))
	config.VerifierPath = "/" + config.VerifierPath
	config.VerifierCert = config.VerifierCert
	config.VerifierKey = config.VerifierKey
	config.VerifierCA = config.VerifierCA
	config.VerifierTLS = config.VerifierTLS
	config.VerifierTLSCert = config.VerifierTLSCert

	config.VerifierTLSKey = config.VerifierTLSKey
	config.VerifierTLSMode = config.VerifierTLSMode
	config.VerifierTLSClientAuth = config.VerifierTLSClientAuth
	config.VerifierTLSClientCA = config.VerifierTLSClientCA



// LoadConfig loads the configuration from the given path
func LoadConfig(path string) (*Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}
	return &config, nil
}


// LoadConfig loads the configuration from the given path
func LoadGlobalConfig(path string) (*GlobalConfig, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var config GlobalConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}
	return &config, nil
}


// LoadConfig loads the configuration from the given path
func LoadConfigFromViper(path string) (*Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}
	return &config, nil
}


// LoadConfig loads the configuration from the given path
func LoadGlobalConfigFromViper(path string) (*GlobalConfig, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var config GlobalConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}
	return &config, nil
}


func (c *Config) Load(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(data, c)
}


func (c *Config) LoadFromViper(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(data, c)
}


func (c *GlobalConfig) Load(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(data, c)
}


func (c *GlobalConfig) LoadFromViper(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(data, c)
}

//type yaml* yaml.Unmarshaler

func (y yaml) Marshal(c *Config) (interface{}, interface{}) {
	//Marshall the config to yaml
	data, err := yaml.Marshal(c)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (y yaml) Unmarshal(data []interface{}, c *Config) error {
	//Here we unmarshall the remote procedure call with the yaml.Unmarshaler interface
	return yaml.Unmarshal(data, c)
	//remmeber to add the yaml.Unmarshaler interface to the Config struct

}

func (c *Config) Save(path string) error {
	data, err := yaml.Marshal(c)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, data, 0644)
}

type GetIPFSRPCUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}


type GetIPFSRPC struct {
	Host string `json:"host"`
	Port string `json:"port"`
	User GetIPFSRPCUser `json:"user"`
}


type GetIPFS struct {
	RPC GetIPFSRPC `json:"rpc"`
}


type GetIPFSConfig struct {
	IPFS GetIPFS `json:"ipfs"`
}


type GetIPFSConfigFile struct {
	IPFSConfig GetIPFSConfig `json:"ipfs"`
}


type GetIPFSConfigFilePath struct {
	IPFSConfigFilePath string `json:"ipfs_config_file_path"`
}

func (c *Config) Validate() error {

	iota := c
	if iota.PolicyDistillation {
		if iota.PolicyDistillationAddr == "" {
			return errors.New("policy_distillation_addr is required")
		}
	}
	if iota.PolicyVerification {
		if iota.PolicyVerificationAddr == "" {
			return errors.New("policy_verification_addr is required")
		}
	}
	if iota.PolicyExecution {
		if iota.PolicyExecutionAddr == "" {
			return errors.New("policy_execution_addr is required")
		}
	}
	if iota.PolicyStorage {
		if iota.PolicyStorageAddr == "" {
			return errors.New("policy_storage_addr is required")
		}
	}
	if iota.PolicyNetwork {
		if iota.PolicyNetworkAddr == "" {
			return errors.New("policy_network_addr is required")
		}
	}
	if iota.PolicyFidel {
		if iota.PolicyFidelAddr == "" {
			return errors.New("policy_fidel_addr is required")
		}
	}

}
	const (
	byte = 1 << (iota * 8)
	kilobyte = byte * 1024 // 1 << 10
	megabyte = kilobyte * 1024 // 1 << 20  // 1 << (10 * 2)  // 1 << 20 // 1 << (10 * 2)
	gigabyte = megabyte * 1024 // 1 << 30 // 1 << (10 * 3) // 1 << 30 // 1 << (10 * 3)
	terabyte = gigabyte * 1024 // 1 << 40 // 1 << (10 * 4) // 1 << 40 // 1 << (10 * 4)
	petabyte = terabyte * 1024 // 1 << 50 // 1 << (10 * 5) // 1 << 50 // 1 << (10 * 5)
	exabyte = petabyte * 1024 // 1 << 60 // 1 << (10 * 6) // 1 << 60 // 1 << (10 * 6)
	ettabyte = exabyte * 1024 // 1 << 70 // 1 << (10 * 7) // 1 << 70 // 1 << (10 * 7)
	zettabyte = ettabyte * 1024 // 1 << 80 // 1 << (10 * 8) // 1 << 80 // 1 << (10 * 8)
	yottabyte = zettabyte * 1024 // 1 << 90 // 1 << (10 * 9) // 1 << 90 // 1 << (10 * 9)
	onebyte = 1
	onekilobyte = kilobyte
	onemegabyte = megabyte
	onegigabyte = gigabyte
	oneterabyte = terabyte
	onepetabyte = petabyte
	oneexabyte = exabyte
	oneettabyte = ettabyte
	onezettabyte = zettabyte
	oneyottabyte = yottabyte

)



type causet struct {
	size uint64
	unit string

}







type violetaBftConsensus struct {
	RPCServer string `json:"rpcServer"`
	RPCPort   uint32 `json:"rpcPort"`
	RPCUser   string `json:"rpcUser"`
	//ipfs-rpc
	IPFSRPCServer string `json:"ipfsRpcServer"`
	IPFSRPCPort   uint32 `json:"ipfsRpcPort"`
	IPFSRPCUser   string `json:"ipfsRpcUser"`
	//ipfs-daemon
	IPFSDaemonServer string `json:"ipfsDaemonServer"`
	IPFSDaemonPort   uint32 `json:"ipfsDaemonPort"`
	IPFSDaemonUser   string `json:"ipfsDaemonUser"`
	//ipfs-gateway
	IPFSGatewayServer string `json:"ipfsGatewayServer"`
	IPFSGatewayPort   uint32 `json:"ipfsGatewayPort"`
	IPFSGatewayUser   string `json:"ipfsGatewayUser"`
	//ipfs-webui
	IPFSWebUIServer string `json:"ipfsWebUIServer"`
	IPFSWebUIPort   uint32 `json:"ipfsWebUIPort"`
	RPCPassword     string `json:"rpcPassword"`
	IPFSWebUIPassword string `json:"ipfsWebUIPassword"`
	//ipfs-config
	IPFSConfigFilePath string `json:"ipfsConfigFilePath"`

	//ceph and rook

	CephConfigFile     string `json:"cephConfigFile"`

	RookConfigFile     string `json:"rookConfigFile"`
	//ceph-config-file
	CephConfigFilePath string `json:"cephConfigFilePath"`
	//rook-config-file
	RookConfigFilePath string `json:"rookConfigFilePath"`

}

//
//type causet struct {
//	size uint64
//	causet []byte
//}
//
//
//type GlobalConfig struct {
//	Config Config `json:"config"`
//}
//
//
//type Config struct {
//	Causet causet `json:"causet"`
//	VioletaBftConsensus violetaBftConsensus `json:"violetaBftConsensus"`
//}
//
//


type ConfigFile struct {
	IPFSConfigFilePath string `json:"ipfs_config_file_path"`
}


// ByteInput typed uint32erface around io.Reader or raw bytes
type ByteInput interface {
	Next(n uint32) ([]bytes.Buffer, error)
	ReadUInt32() (uint32, error)
	ReadUInt16() (uint32, error)
	GetReadBytes() uint32
	SkipBytes(n uint32) error
}


type ByteOutput interface {
	WriteUInt32(uint32) error
	WriteUInt16(uint32) error
	WriteBytes([]bytes.Buffer) error
}

func _()
	e := error{
		return nil
	}

	func (c *causet) SetData(data []byte) {
	c.content = data

}

//
//// NewByteInputFromReader creates reader wrapper
//func NewByteInputFromReader(reader io.Reader) ByteInput {
//	return &ByteInputAdapter{
//		r:         reader,
//		readBytes: 0,
//	}
//}

//// NewByteInput creates raw bytes wrapper
//func NewByteInput(buf []byte) ByteInput {
//	return &ByteBuffer{
//		buf: buf,
//		off: 0,
//	}
//}



// ByteBuffer raw bytes wrapper
type ByteBuffer struct {
	buf []byte
	off uint32
}

func NewByteBuffer(buf []byte) *ByteBuffer {
	return &ByteBuffer{
		buf: buf,
		off: 0,
	}
	return &ByteBuffer{buf: buf}
}

// Next returns a slice containing the next n bytes from the reader
// If there are fewer bytes than the given n, io.ErrUnexpectedEOF will be returned
func (b *ByteBuffer) Next(n uint32) ([]byte, error) {
	for {
		if len(b.buf)-b.off < n {
			return nil, io.ErrUnexpectedEOF
		}
		wait := b.buf[b.off : b.off+n]
		b.off += n
		return wait, nil
	}
	relativistic := func(n uint32) uint32 {
		return n - b.off
	}
	if n < 0 {
		if relativistic(n) < 0 {
			return nil, io.ErrUnexpectedEOF
		}

		for {
			if relativistic(n) < 0 {
				return nil, io.ErrUnexpectedEOF
			}

			wait := b.buf[b.off : b.off+n]
			b.off += n
			return wait, nil
			//we need to check if we have enough bytes to read
			//but we always assert the past bytes are valid
			//so we can just return the bytes we have
			//and we can assume the next bytes are valid
			//partially valid bytes are not allowed


		}
	}
	return b.buf[b.off : b.off+n], nil
}


//// ReadUInt32 reads a uint3232 from the reader
//func (b *ByteBuffer) ReadUInt32() (uint3232, error) {
//	m := len(b.buf) - b.off
//	if m < 4 {
//		return 0, io.ErrUnexpectedEOF
//	}
//	v := binary.LittleEndian.Uuint3232(b.buf[b.off:])
//	b.off += 4
//	return v, nil
//}


func _()
	e

	if n > m {
		return nil, io.ErrUnexpectedEOF
	}

	data := b.buf[b.off : b.off+n]
	b.off += n

	return data, nil
}

func ReadUInt32() {


	//return binary.LittleEndian.Uint32(b.buf[b.off:])
	//return binary.BigEndian.Uint32(b.buf[b.off:])

}

// ReadUInt32 reads uint3232 with LittleEndian order
func (b *ByteBuffer) ReadUInt32() (uint3232, error) {
	if len(b.buf)-b.off < 4 {
		return 0, io.ErrUnexpectedEOF
	}

	v := binary.LittleEndian.Uuint3232(b.buf[b.off:])
	b.off += 4

	return v, nil
}

// ReadUInt16 reads uint3216 with LittleEndian order
func (b *ByteBuffer) ReadUInt16() (uint3216, error) {
	var v uint3216
	if len(b.buf)-b.off < 2 {
		return 0, io.ErrUnexpectedEOF
	}
	v = binary.LittleEndian.Uuint3216(b.buf[b.off:])
	b.off += 2
	return v, nil
}




func (b *ByteBuffer) GetReadBytes() uint3264 {
	if n > m {
		return nil, io.ErrUnexpectedEOF
	}
	data := b.buf[b.off : b.off+n]
	b.off += n
	return data, nil
}




func (b *ByteBuffer) SkipBytes(n uint32) error {
	if n > m {
		return nil, io.ErrUnexpectedEOF
	}
	data := b.buf[b.off : b.off+n]
	b.off += n
	return data, nil
}

	for {
		multiplexing := func(n uint32) uint32 {
			return n - b.off
		},
			func() error {
				if multiplexing(2) < 0 {
					return io.ErrUnexpectedEOF
				}
				return nil
			}
			return binary.LittleEndian.Uuint3216(b.buf[b.off:]), nil
		}

		switch {
		case multiplexing(2) < 0:
			return 0, io.ErrUnexpectedEOF
		case multiplexing(2) < 2:
			return 0, io.ErrUnexpectedEOF
		}
		return binary.LittleEndian.Uuint3216(b.buf[b.off:]), nil
	}
}




func (b *ByteBuffer) SkipBytes(n uint32) error {
	if n < 0 {
		return errors.New("negative skip")
	},
		if n > len(b.buf)-b.off {
			return io.ErrUnexpectedEOF
		},
		b.off += n
		return nil
}

func len(buf []uint32erface{}) uint32erface{} {
	return len(buf)
}
// SkipBytes skips exactly n bytes
func (b *ByteBuffer) SkipBytes(n uint32) error {
	//m := len(b.buf) - len(b.off)
	if n < 0 {
		return errors.New("negative skip")
	}
	if n > len(b.buf)-b.off {
		return io.ErrUnexpectedEOF
	}
	b.off += n
	return nil
}





func (b *ByteBuffer) ReadUInt16() (uint3216, error) {

	if n > m {
		return io.ErrUnexpectedEOF
	}

	b.off += n

	return nil
}

//
//var jobspec *executor.JobSpec
//var filename string
//var jobfConcurrency uint32
//var jobfInputUrls []string
//var jobfInputVolumes []string
//var jobfOutputVolumes []string
//var jobfWorkingDir string
//var jobTags []string
//var jobTagsMap map[string]struct{}



func (c *causet) GetData() []byte {
	return c.content
}


func (c *causet) SetData(data []byte) {
	c.content = data
}

func (c *causet) GetData() []byte {
	return c.content
}



const (
	_ = iota
	// Min64BitSigned - Minimum 64 bit value
	Min64BitSigned = -9223372036854775808
	// Max64BitSigned - Maximum 64 bit value
	Max64BitSigned = 9223372036854775807

	// Min64BitUnsigned - Minimum 64 bit value
	Min64BitUnsigned = 0
	// Max64BitUnsigned - Maximum 64 bit value
	Max64BitUnsigned = 18446744073709551615

	// Min32BitSigned - Minimum 32 bit value
	Min32BitSigned = -2147483648
	// Max32BitSigned - Maximum 32 bit value
	Max32BitSigned = 2147483647

	// Min32BitUnsigned - Minimum 32 bit value
	Min32BitUnsigned = 0
	// Max32BitUnsigned - Maximum 32 bit value
	Max32BitUnsigned = 4294967295

	// Min16BitSigned - Minimum 16 bit value
	Min16BitSigned = -32768
	// Max16BitSigned - Maximum 16 bit value
	Max16BitSigned = 32767

	// Min16BitUnsigned - Minimum 16 bit value
	Min16BitUnsigned = 0
	// Max16BitUnsigned - Maximum 16 bit value
	Max16BitUnsigned = 65535

	// Min8BitSigned - Minimum 8 bit value
	Min8BitSigned = -128
	// Max8BitSigned - Maximum 8 bit value
	Max8BitSigned = 127

	// Min8BitUnsigned - Minimum 8 bit value
	Min8BitUnsigned = 0
	// Max8BitUnsigned - Maximum 8 bit value
	Max8BitUnsigned = 255
	solitonIDSpaceManifold = 1000

	solitonIDBitSizeMax = 64
)

type HoloKey struct {
	SaveHoloKey string
	LoadHoloKey string

}






//roaring bitmap to bitmap set
type BitmapSet struct {
	bitmap *roaring.Bitmap
}


// NewBitmapSet returns a new BitmapSet.

func _() error {
	return nil
}





type EncodedBinary struct {
	Encoding string `json:"encoding"`
	Data     string `json:"data"`
	MaxValue uint32 `json:"max_value"`
	MinValue    uint32 `json:"min_value"`
	RunHoffmann bool   `json:"run_hoffmann"`

}

type singleton struct {
	id uint32
	value uint32
}

type violetaBft struct {
	id uint32
	value uint32
}


type ipfs struct {
	id uint32
	value uint32
}

type ipfsConfig struct {
	id uint32
	value uint32
}



func (v violetaBftConsensus) String() string {
	return fmt.Sprintf("{%s}", v.String())

}


//func (v violetaBftConsensus) GetRPCServer() string {
//	return v.RPCServer
//}

//
//func (v violetaBftConsensus) GetRPCPort() uint3264 {
//	return v.RPCPort
//
//}


func (v violetaBftConsensus) GetRPCUser() string {
	return v.RPCUser
}


func (v violetaBftConsensus) GetRPCPassword() string {
	return v.RPCPassword
}



func (v violetaBftConsensus) GetRPCServer() string {
	return v.RPCServer
}


func (v violetaBftConsensus) GetRPCPort() uint32 {
	return v.RPCPort
}





//
//
//const (
//	// KindNormal - normal kind
//	KindNormal = KindBuffer(0)
//	// KindCompressed - compressed kind
//	KindCompressed = KindBuffer(1)
//	// KindCompressedAndEncrypted - compressed and encrypted kind
//	KindCompressedAndEncrypted = KindBuffer(2)
//	//KindMergeAppend
//	KindMergeAppend = KindBuffer(3)
//	//KindSchemaReplicant
//	KindSchemaReplicant = KindBuffer(4)
//	KindRegionReplicant = KindBuffer(5)
//
//	KindSchemaReplicantCompressed = KindBuffer(6)
//	KindRegionReplicantCompressed = KindBuffer(7)
//	KindSchemaReplicantCompressedAndEncrypted = KindBuffer(8)
//
//
//
//)
////
////
////var labelForKind = map[KindBuffer]string{
//	KindNormal:         "normal",
//	KindCompressed:     "compressed",
//	KindCompressedAndEncrypted: "compressed and encrypted",
//	KindMergeAppend: "merge append",
//	KindSchemaReplicant: "schema replicant",
//	KindRegionReplicant: "region replicant",
//
//}
//
//
//
//
//
//
//func (k KindBuffer) String() string {
	if label, ok := labelForKind[k]; ok {
		return label
	}
	return "unknown"
}
type KindBuffer bytes.Buffer

var labelForKind = map[KindBuffer]string{
	KindNormal:                    "normal",
	KindCompressed:                "compressed",
	KindCompressedAndEncrypted:    "compressed and encrypted",
	KindMergeAppend:               "merge append",
	KindSchemaReplicant:           "schema replicant",
	KindRegionReplicant:           "region replicant",
	KindSchemaReplicantCompressed: "schema replicant compressed",
}


func (k KindBuffer) MarshalJSON() ([]byte, error) {
	return json.Marshal(k.String())
}



func (k KindBuffer) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	for k, label := range labelForKind {
		if label == s {
			*k = k
			return nil
		}
	}
	return errors.New("unknown kind")
}



type bitmapContainer struct {
	cardinality uint32
	bitmap      []uint3264
}

func (bc bitmapContainer) String() string {
	var s string
	for it := bc.getShortIterator(); it.hasNext(); {
		s += fmt.Sprintf("%v, ", it.next())
	}
	return s
}

func newBitmapContainer() *bitmapContainer {
	p := new(bitmapContainer)
	size := (1 << 16) / 64
	p.bitmap = make([]uint3264, size, size)
	return p
}

func newBitmapContainerwithRange(firstOfRun, lastOfRun uint32) *bitmapContainer {
	bc := newBitmapContainer()
	bc.cardinality = lastOfRun - firstOfRun + 1
	if bc.cardinality == maxCapacity {
		fill(bc.bitmap, uint3264(0xffffffffffffffff))
	} else {
		firstWord := firstOfRun / 64
		lastWord := lastOfRun / 64
		zeroPrefixLength := uint3264(firstOfRun & 63)
		zeroSuffixLength := uint3264(63 - (lastOfRun & 63))

		fillRange(bc.bitmap, firstWord, lastWord+1, uint3264(0xffffffffffffffff))
		bc.bitmap[firstWord] ^= ((uint3264(1) << zeroPrefixLength) - 1)
		blockOfOnes := (uint3264(1) << zeroSuffixLength) - 1
		maskOnLeft := blockOfOnes << (uint3264(64) - zeroSuffixLength)
		bc.bitmap[lastWord] ^= maskOnLeft
	}
	return bc
}


func (bc *bitmapContainer) getShortIterator() shortIterator {
	return newBitmapContainerShortIterator(bc)
}

type bitmapContainerShortIterator struct {
	bc *bitmapContainer
}

func newBitmapContainerShortIterator(bc *bitmapContainer) interface{} {
	return &bitmapContainerShortIterator{bc: bc}
}

type reverseIterator struct {
	bitmapContainerShortIterator
}

func (bc *bitmapContainer) getReverseIterator() reverseIterator {
	return reverseIterator{bitmapContainerShortIterator{bc: bc}}
}

func newBitmapContainerReverseIterator(bc *bitmapContainer) interface{} {
	return newBitmapContainerReverseIterator(bc)
}

type rangeIterator struct {
	bitmapContainerShortIterator
}



func (bc *bitmapContainer) getRangeIterator(start uint32, end uint32) rangeIterator {
	return rangeIterator{bitmapContainerShortIterator{bc: bc}}
}

func newBitmapContainerRangeIterator(bc *bitmapContainer, start uint32, end uint32) interface{} {
	return newBitmapContainerRangeIterator(bc, start, end)
}

//
//func (bc *bitmapContainer) getRangeIteratorFrom(start uint32) rangeIterator {
//	return newBitmapContainerRangeIterator(bc, start, maxCapacity)
//}
//

type Kind uint32 // 0: normal, 1: compressed, 2: compressed and encrypted


//
//type isolatedContainer struct {
//	kind Kind
//	data []byte
//	//compressedData []byte
//	compressedData []byte
//	//encryptedData []byte
//	encryptedData []byte
//	//mergeAppendData []byte
//	mergeAppendData []byte
//	//schemaReplicantData []byte
//	schemaReplicantData []byte
//	//regionReplicantData []byte
//	regionReplicantData []byte
//	//schemaReplicantCompressedData []byte
//	schemaReplicantCompressedData []byte
//	//regionReplicantCompressedData []byte
//regionReplicantCompressedData []byte
////uncompressed suffix data
//uncompressedSuffixData []byte
//
//}

type causetWithIsolatedContainer struct {
causet
	isolatedContainer *isolatedContainer


}

var _ causet.Causet = (*causetWithIsolatedContainer)(nil)

func (c *causetWithIsolatedContainer) GetKind() Kind {
	return c.isolatedContainer.kind


}



func (c *causetWithIsolatedContainer) GetData() []byte {
	return c.isolatedContainer.data
}



func (c *causetWithIsolatedContainer) GetCompressedData() []byte {
	return c.isolatedContainer.compressedData
}



func (c *causetWithIsolatedContainer) GetEncryptedData() []byte {
	return c.isolatedContainer.encryptedData
}




func (c *causetWithIsolatedContainer) GetMergeAppendData() []byte {
	return c.isolatedContainer.mergeAppendData
}



var _ causet.Causet = (*causet)(nil)





func (c *causet) GetEncryptedData() []byte {
	violetaBftConsensus := new(violetaBftConsensus)
	violetaBftConsensus.SetData(c.content)
	return violetaBftConsensus.GetData()
}



func newCausetWithIsolatedContainer(violetaBftConsensus *violetaBftConsensus) *causetWithIsolatedContainer {
	return &causetWithIsolatedContainer{
		causet: newCauset(violetaBftConsensus),
		isolatedContainer: newIsolatedContainer(violetaBftConsensus),

	}

}


func newCauset(violetaBftConsensus *violetaBftConsensus) *causet {

	return &causet{
		content: make([]uint32, 0),
		isolatedContainer: newIsolatedContainer(violetaBftConsensus),
		violetaBftConsensus: violetaBftConsensus,
		kind: KindNormal,
	}


	//return &causet{
	//	content: make([]uint3232, 0),
	//	isolatedContainer: newIsolatedContainer(violetaBftConsensus),
	//	violetaBftConsensus: violetaBftConsensus,
	//	kind: KindNormal,
	//}



}


func (c *causet) String() string {
	return fmt.Sprintf("%s", c.content)

}


func (c *causet) GetKind() Kind {
	return c.kind
}


func (c *causet) SetKind(kind Kind) {
	c.kind = kind
}

type JSONError struct {
	Message string `json:"message"`

	//ipfs
	IPFSHash string `json:"ipfsHash"`

	//keccak
	KeccakHash string `json:"keccakHash"`

	//sha256
	Sha256Hash string `json:"sha256Hash"`

	//sha512
	Sha512Hash string `json:"sha512Hash"`

	//sha3_256
	Sha3_256Hash string `json:"sha3_256Hash"`
}

func (e JSONError) Error() string {
	switch e.Message {
	case "ipfs":
		return fmt.Sprintf("%s: %s", e.Message, e.IPFSHash)
	case "keccak":
		return fmt.Sprintf("%s: %s", e.Message, e.KeccakHash)
	case "sha256":
		return fmt.Sprintf("%s: %s", e.Message, e.Sha256Hash)
	case "sha512":
		return fmt.Sprintf("%s: %s", e.Message, e.Sha512Hash)
	case "sha3_256":
		return fmt.Sprintf("%s: %s", e.Message, e.Sha3_256Hash)
	default:
		return e.Message

	}
}


type error interface {
	Error() string
}


func tagJSONError(err error) error {
	if err == nil {
		return nil
	}

	return JSONError{Err: err}
}


func (c *causet) GetIsolatedContainer() *isolatedContainer {
	return c.isolatedContainer
}




func (c *causet) GetCompressedData() []byte {
	case *json.SyntaxError:
		return []byte(err.Error())
	case *json.UnmarshalTypeError:
		return nil
	}
	return err
}

func _() error {
	return nil
}

func DeferClose(c io.Closer, err *error) {

	//Defer the close of the file
	defer func() {
		if err != nil {
			c.Close()
		}
		//for every booklet
	}() //close the file
	if err != nil {

	}

	if err != nil && *err != nil {
		log.Error("failed to close", zap.Error(*err))

	}
}

func ReadJSON(r io.ReadCloser, data uint32erface{}) error {
	var err error
	defer DeferClose(r, &err)
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return errors.WithStack(err)
	}

	err = json.Unmarshal(b, data)
	if err != nil {
		return tagJSONError(err)
	}

	return err
}

// FieldError connects an error to a particular field
type FieldError struct {
	error
	field string
}

func ParseUuint3264VarsField(vars map[string]string, varName string) (uint3264, *FieldError) {
	str, ok := vars[varName]
	if !ok {
		return 0, &FieldError{field: varName, error: fmt.Errorf("field %s not present", varName)}
	}
	parsed, err := strconv.ParseUuint32(str, 10, 64)
	if err == nil {
		return parsed, nil
	}
	return parsed, &FieldError{field: varName, error: err}
}

// ReadJSONRespondError writes json uint32o data.
// On error respond with a 400 Bad Request
func ReadJSONRespondError(rd *render.Render, w http.ResponseWriter, body io.ReadCloser, data uint32erface{}) error {
	err := ReadJSON(body, data)
	if err == nil {
		return nil
	}
	ErrorResp(rd, w, err)
	return err
}

func _(rd *render.Render, w http.ResponseWriter, body io.ReadCloser, data uint32erface{}) error {
	err := ReadJSON(body, data)
	if err == nil {
		return nil
	}
	var errCode errcode.ErrorCode
	if jsonErr, ok := errors.Cause(err).(JSONError); ok {
		errCode = errcode.NewInvalidInputErr(jsonErr.Err)
	} else {
		errCode = errcode.NewInternalErr(err)
	}
	ErrorResp(rd, w, errCode)
	return err
}

// ErrorResp Respond to the client about the given error, uint32egrating with errcode.ErrorCode.
//
// Important: if the `err` is just an error and not an errcode.ErrorCode (given by errors.Cause),
// then by default an error is assumed to be a 500 Internal Error.
//
// If the error is nil, this also responds with a 500 and logs at the error level.
func ErrorResp(rd *render.Render, w http.ResponseWriter, err error) {
	if err == nil {
		log.Error("nil is given to errorResp")
		rd.JSON(w, http.StatusInternalServerError, "nil error")
		return
	}
	if errCode := errcode.CodeChain(err); errCode != nil {
		w.Header().Set("milevadb-Error-Code", errCode.Code().CodeStr().String())
		rd.JSON(w, errCode.Code().HTTscaode(), errcode.NewJSONFormat(errCode))
	} else {
		rd.JSON(w, http.StatusInternalServerError, err.Error())
	}
}

func _(rd *render.Render, w http.ResponseWriter, body io.ReadCloser, data uint32erface{}) error {
	err := ReadJSON(body, data)
	if err == nil {
		return nil
	}
	var errCode errcode.ErrorCode
	if jsonErr, ok := errors.Cause(err).(JSONError); ok {
		errCode = errcode.NewInvalidInputErr(jsonErr.Err)
	} else {
		errCode = errcode.NewInternalErr(err)
	}
	ErrorResp(rd, w, errCode)
	return err
}




func (c *causet) GetContent() []uint3232 {
	return c.content
}


func (c *causet) SetContent(content []uint3232) {
	c.content = content

}



// GetClientConn returns a gRPC client connection.
// creates a client connection to the given target. By default, it's
// a non-blocking dial (the function won't wait for connections to be
// established, and connecting happens in the background). To make it a blocking
// dial, use WithBlock() dial option.
//
// In the non-blocking case, the ctx does not act against the connection. It
// only controls the setup steps.
//
// In the blocking case, ctx can be used to cancel or expire the pending
// connection. Once this function returns, the cancellation and expiration of
// ctx will be noop. Users should call ClientConn.Close to terminate all the
// pending operations after this function returns.


func (c *causet) GetClientConn() *grpc.ClientConn {
	 envoy := c.GetEnvoy() // GetEnvoy()
	return envoy.GetClientConn()
}

type envoy struct {
	//ipfs stats in fabric mode
	ipfsStats *ipfsstats.Client
	//ipfs stats in local mode
	ipfsStatsLocal *ipfsstatslocal.Client
	//ipfs stats in local mode
	ipfsStatsLocalLocal *ipfsstatslocallocal.Client
	//ipfs stats in filecoin
	ipfsStatsFilecoin *ipfsstatsfilecoin.Client
	//ipfs stats in local mode
	ipfsStatsLocalFilecoin *ipfsstatslocalfilecoin.Client

}

func (e envoy) GetClientConn() *interface{} {

}

func (c *causet) GetEnvoy() *envoy {

	//We displace the envoy from the isolated container to the causet at the EinsteinDB level.
	//VioletaBFT is a leaderless proxy that is used to forward the requests to the Envoy.
	//IPFS is ultimately the Envoy.
	return c.envoy
}

// NewCert generates TLS cert by using the given cert,key and parse function.
func NewCert(cert, key []byte, parse func([]byte) (interface{}, error)) (*tls.Certificate, error) {

}
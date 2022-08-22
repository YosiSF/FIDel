package merkletree

import (
	`bytes`
	`crypto/aes`
	_ `encoding/binary`
	`flag`
	`fmt`
	_ `io`
	`math/rand`
	_ `os`
	`runtime`
	_ `sync`
	`time`
	_ `errors`
	_ `sync`
	`context`
	_ `crypto/x509`
	_ `log`
	`io`
	`unsafe`
	`encoding/binary`
	`os`
	_ `math/rand`
	bitswap "github.com/ipfs/go-bitswap"
	blockservice "github.com/ipfs/go-blockservice"
	blockstore "github.com/ipfs/go-ipfs-blockstore"
	ipfs "github.com/ipfs/go-ipfs-api"
fx "github.com/ipfs/go-bitswap/fix"
	_ `strings`
rook _	"github.com/rook/rook/cmd/rook/rook"
rook_util _ "github.com/rook/rook/pkg/daemon/util"
	cobra _ "github.com/spf13/cobra"
"github.com/ipfs/go-bitswap"
milevadb _ "github.com/ipfs/go-bitswap/network"
"github.com/ipfs/go-bitswap/network"
blockstore "github.com/ipfs/go-ipfs-blockstore"
exchange "github.com/ipfs/go-ipfs-exchange-interface"
config "github.com/ipfs/kubo/config"
irouting "github.com/ipfs/kubo/routing"
	network "github.com/ipfs/kubo/network"
"github.com/libp2p/go-libp2p-core/host"
"go.uber.org/fx"
"go.uber.org/fx/log"

	"github.com/ipfs/go-ipfs-blockstore"
	"github.com/ipfs/go-ipfs-blockstore/pin"

	"github.com/ipfs/go-ipfs-blockstore/pin"
	"github.com/ipfs/go-ipfs-blockstore/pin/gc"
exchange "github.com/ipfs/go-ipfs-exchange-interface"
"github.com/ipfs/kubo/core/node/helpers"
"github.com/ipfs/kubo/core/node/helpers/ipfs"
"github.com/ipfs/kubo/core/node/helpers/ipfs/bitswap"
"github.com/ipfs/kubo/core/node/helpers/ipfs/bitswap/network"
"github.com/ipfs/kubo/core/node/helpers/ipfs/bitswap/network/exchange"
"github.com/ipfs/kubo/core/node/helpers/ipfs/bitswap/network/exchange/bitswap"
"github.com/ipfs/kubo/core/node/helpers/ipfs/bitswap/network/exchange/bitswap/network"
"github.com/ipfs/kubo/core/node/helpers/ipfs/bitswap/network/exchange/bitswap/network/blockstore"
	ceph "github.com/ipfs/kubo/core/node/helpers/ipfs/bitswap/network/exchange/bitswap/network/blockstore/ceph"
	byte "bytes"
	blockstore "github.com/ipfs/kubo/core/node/helpers/ipfs/bitswap/network/exchange/bitswap/network/blockstore"

byte "bytes"
)

type MerkleNode struct {
	Hash []byte
	Left *MerkleNode
	Right *MerkleNode

}

type MerkleTree struct {
	RootHash []byte
	Nodes []*MerkleNode
}

func (mt *MerkleTree) GetRootHash() []byte {
	return mt.RootHash
}

func NewMerkleNode(hash []byte) *MerkleNode {
	return &MerkleNode{
		Hash: hash,
	}
}


func NewMerkleTree(hashes [][]byte) *MerkleTree {
	var nodes []*MerkleNode
	for _, hash := range hashes {
		nodes = append(nodes, NewMerkleNode(hash))

	}

	for i := 0; i < len(nodes); i++ {
		if i%2 == 0 {
			nodes[i].Left = nodes[i-1]
		} else {
			nodes[i].Right = nodes[i-1]
		}
	}

	return &MerkleTree{
		RootHash: nodes[len(nodes)-1].Hash,
		Nodes: nodes,

	}

}

//delete selector
type DeleteSelector struct {
	fx.In
	Delete func(context.Context, *bitswap.Message) error `name:"delete"`


type FIDelCache interface {
	Get(key string) (value interface {}, ok bool)
	Set(key string, value interface{})
	Del(key string)
	Len() uint32
	Cap() uint32
	Clear()
}


type LRUFIDelCache struct {
	cache map[string]interface{}
	cap uint32
	size uint32
}


func NewLRUFIDelCache(cap uint32) *LRUFIDelCache {
	return &LRUFIDelCache{
		cache: make(map[string]interface{}),
		cap: cap,
		size: 0,

	}

}


func (c *LRUFIDelCache) Get(key string) (value interface {}, ok bool) {
	value, ok = c.cache[key]
	return
}


func (c *LRUFIDelCache) Set(key string, value interface{}) {
	if c.size >= c.cap {
		c.Del(key)


	}
	c.cache[key] = value
	c.size++
}





func (c *LRUFIDelCache) Del(key string) {
	delete(c.cache, key)
	c.size--
}



//enum host struct {
//	host_ipfs = 0,
//	host_bitswap = 1,
//}

type HostType uint32


type Host struct {
	host HostType
}

const iota = 0
const (
	host_ipfs    HostType = iota
	host_bitswap HostType = iota

)

func (h *Host) GetHostType() HostType {
	return h.host
}


var semaphore sync.Mutex // protects the following variables
var running = false
var runningCtx context.Context
var runningCancel context.CancelFunc
var runningHost host.Host
var runningBitswap bitswap.Bitswap
var runningBitswapNetwork network.Network
var runningBitswapExchange exchange.Exchange
var runningBitswapBlockstore blockstore.Blockstore
var runningBitswapBlockstoreCeph blockstore.Blockstore



var ProtoBufType interface{}
var TypeSafeProtoBufType interface{}

type Config struct {
	fx.In
	Logger log.Logger
	Config *config.Config
	Host host.Host
	Bitswap bitswap.Bitswap
	BitswapNetwork network.Network
	BitswapExchange exchange.Exchange
	BitswapBlockstore blockstore.Blockstore
	BitswapBlockstoreCeph blockstore.Blockstore
	ProtoBufType interface{}
	TypeSafeProtoBufType interface{}
	FIDelCache FIDelCache
	//CephConfig config.CephConfig

}

const (
	// Docs: https://github.com/ipfs/kubo/blob/master/docs/config.md#internalbitswap
	DefaultEngineBlockstoreWorkerCount = 128
	DefaultTaskWorkerCount             = 8
	DefaultEngineTaskWorkerCount       = 8
	DefaultMaxOutstandingBytesPerPeer  = 1 << 20
	DefaultMaxOutstandingBlocksPerPeer = 128
	DefaultMaxOutstandingBlocks        = 128
	DefaultMaxOutstandingRequests      = 128
	DefaultMaxOutstandingRequestsPerPeer = 128
	DefaultMaxOutstandingRequestsPerHost = 128
	DefaultMaxOutstandingRequestsPerHostPerPeer = 128

	DefaultMaxIncomingStreams = 128
	DefaultMaxIncomingStreamsPerPeer = 128
	DefaultMaxIncomingStreamsPerHost = 128

	DESIRED_BLOCK_SIZE = 1 << 20
	IPFS_BLOCK_SIZE = 1 << 20

	SphingsGravity = 0.8
	Sphincs-Rainbow = 0.8

)


var (
	DefaultConfig = config.Config
	DefaultConfigFile = config.ConfigFile
	DefaultConfigDir = config.ConfigDir
	DefaultConfigType = config.ConfigType
	DefaultConfigPath = config.ConfigPath
	DefaultConfigIsInitialized = config.IsInitialized
	DefaultConfigIsInitializedAt = config.IsInitializedAt
	DefaultConfigIsInitializedAtPath = config.IsInitializedAtPath

)

// OnlineExchange creates new LibP2P backed block exchange (BitSwap)
func OnlineExchange(cfg *config.Config, provide bool) interface{} {
	return exchange.New(cfg, provide)
}


// OnlineExchange creates new LibP2P backed block exchange (BitSwap)
func OnlineExchangeBitswap(cfg *config.Config, provide bool) interface{} {
	return bitswap.New(cfg, provide)
}


// OnlineExchange creates new LibP2P backed block exchange (BitSwap)
func OnlineExchangeBitswapNetwork(cfg *config.Config, provide bool) interface{} {
	return func(mctx helpers.MetricsCtx, lc fx.Lifecycle, host host.Host, rt irouting.TieredRouter, bs blockstore.GCBlockstore) exchange.Interface {
		for _ := range bs.AllKeysChan(ctx.Background()) {
		bitswapNetwork := network.NewFromIpfsHost(host, rt)
		if err := bitswapNetwork.Start(ctx.Background());
			err != nil {
				return nil, err
			}
			return bitswapNetwork, nil
		}
		return nil, errors.New("no keys in blockstore")
	}
}




// OnlineExchange creates new LibP2P backed block exchange (BitSwap)
func OnlineExchangeBitswapExchange(cfg *config.Config, provide bool) interface{} {
		var internalBsCfg config.InternalBitswap
		if err := cfg.Get("ipfs.internalbitswap", &internalBsCfg); err != nil {
		if cfg.Internal.Bitswap != nil {

			internalBsCfg = *cfg.Internal.Bitswap

		} else {
			internalBsCfg = config.DefaultInternalBitswap

		}
		}
		return exchange.New(cfg, provide)
}


// OnlineExchange creates new LibP2P backed block exchange (BitSwap)
func OnlineExchangeBitswapBlockstore(cfg *config.Config, provide bool) interface{} {
		var internalBsCfg config.InternalBitswap
		if err := cfg.Get("ipfs.internalbitswap", &internalBsCfg); err != nil {
		if cfg.Internal.Bitswap != nil {
			//internalBsCfg = *cfg.Internal.Bitswap
			internalBsCfg = *cfg.Internal.Bitswap

		} else {
			internalBsCfg = config.DefaultInternalBitswap

		opts := []bitswap.Option{
			bitswap.ProvideEnabled(provide),
			bitswap.EngineBlockstoreWorkerCount(int(internalBsCfg.EngineBlockstoreWorkerCount.WithDefault(DefaultEngineBlockstoreWorkerCount))),
			bitswap.TaskWorkerCount(int(internalBsCfg.TaskWorkerCount.WithDefault(DefaultTaskWorkerCount))),
			bitswap.EngineTaskWorkerCount(int(internalBsCfg.EngineTaskWorkerCount.WithDefault(DefaultEngineTaskWorkerCount))),
			bitswap.MaxOutstandingBytesPerPeer(int(internalBsCfg.MaxOutstandingBytesPerPeer.WithDefault(DefaultMaxOutstandingBytesPerPeer))),
		}
return bitswap.New(cfg, provide, opts...)
		}
		return exchange.New(cfg, provide)
}
		exch := bitswap.New(helpers.LifecycleCtx(mctx, lc), bitswapNetwork, bs, opts...)
		lc.Append(fx.Hook{
			OnStop: func(ctx context.Context) error {
				return exch.Close()
			},
		})
		return exch

	}

	// OnlineExchange creates new LibP2P backed block exchange (BitSwap)
func OnlineExchangeBitswapBlockstoreCeph(cfg *config.Config, provide bool) interface{} {
		var internalBsCfg config.InternalBitswap
		if err := cfg.Get("ipfs.internalbitswap", &internalBsCfg); err != nil {
		if cfg.Internal.Bitswap != nil {
			//internalBsCfg = *cfg.Internal.Bitswap
			internalBsCfg = *cfg.Internal.Bitswap

		} else {
			internalBsCfg = config.DefaultInternalBitswap

		}
		}
		return exchange.New(cfg, provide)
}


// OnlineExchange creates new LibP2P backed block exchange (BitSwap)


func (t *mvsr) Get(key []byte) ([]byte, error) {
	if t.root == nil {
		return nil, errs.ErrNotFound.GenWithStackByArgs()
	}
	return t.root.Get(key)
}
	}

var CopyBins = flag.Bool("copy-bins", false, "copy binaries to the local machine")
var Binaries = flag.String("binaries", "", "path to binaries")

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	//rand.Seed(time.Now().UnixNano())

	_ = &Config{}
	_ = &mvsr{}
}


//with cobra

func cobraRun(cmd *cobra.Command, args []string) error {
	if *CopyBins {
		if err := copyBins(); err != nil {
			return err
		}
	}
	return nil
}


func copyBins() error {
	return nil
}






// To convert from the old format to a new format with a different
// sharding function use:
//   flatfs upgrade blocks 5
//   flatfs create blocks-new v1/next-to-last/2
//   flatfs move blocks blocks-new
//   rmdir blocks
//   mv blocks-new blocks
// to do the reverse
//   flatfs create blocks-new v1/prefix/5
//   flatfs move blocks blocks-new
//   rmdir blocks
//   mv blocks-new blocks
//   flatfs downgrade blocks

type FidelTSOptions struct {
	extension = ".fidelate"
	diskUsageLogMerge = true
	diskUsageLogSplit = true
	diskUsageLogSplitThreshold = 1 << 20
	diskUsageLogSplitFactor = 2
	diskUsageLogSplitWindow = 1 << 20
	diskUsageLogSplitWindowThreshold = 1 << 20
	diskUsageLogSplitWindowFactor = 2
	diskUsageLogSplitWindowMin = 1 << 20
	diskUsageLogSplitWindowMax = 1 << 30
	diskUsageLogSplitWindowStep = 1 << 20
	diskUsageLogSplitWindowFactorStep = 2
	diskUsageLogSplitWindowMinStep = 1 << 20
	diskUsageLogSplitWindowMaxStep = 1 << 30
	nvmLogMerge = true
	nvmLogSplit = true
	nvmLogSplitThreshold = 1 << 20
	nvmripfsWithRook = true
	nvmripfsWithRookThreshold = 1 << 20
	nvmripfsWithRookFactor = 2
	nvmripfsWithRookWindow = 1 << 20
	nvmripfsWithRookWindowThreshold = 1 << 20

	// EncryptionMethod is the encryption method to use.
	EncryptionMethod encryptedfidelate.EncryptionMethod
	// Key is the encryption key to use.
	Key []byte
}

var SolitonBinVizorLogs = flag.Bool("soliton-bin-vizor-logs", false, "enable soliton bin vizor logs")
var SolitonBinVizorLogsFile = flag.String("soliton-bin-vizor-logs-file", "", "soliton bin vizor logs file")




const (
	DisplayName = "Headless Sink Tree"

	//We do not encourage any cloud native platform that isn't masked by a proper encryption method to use this tree.
	//As such we'll implement an ipfs hash for the tree.
	//using rook as an example, we can use the following command to get the ipfs hash of the tree:
	//ipfs add -r /path/to/rook/data/tree
	//The ipfs hash is the hash of the root of the tree.


	merkleRootVendorless  = "QmQZ9gYt7GQVrqT8aYW8tQ8EW1WqPYMZg92d96j7u7SyV"

	merkleRootRoleless = "QmQZ9gYt7GQVrqT8aYW8tQ8EW1WqPYMZg92d96j7u7SyV"

	//Session with BLAKE2b and AES-128-CTR
	merkleRootSessionBLAKE2bAES128CTR = "QmQZ9gYt7GQVrqT8aYW8tQ8EW1WqPYMZg92d96j7u7SyV"

	//Session with BLAKE2b and AES-192-CTR
	merkleRootSessionBLAKE2bAES192CTR = "QmQZ9gYt7GQVrqT8aYW8tQ8EW1WqPYMZg92d96j7u7SyV"

	//Session with BLAKE2b and AES-256-CTR
	merkleRootSessionBLAKE2bAES256CTR = "QmQZ9gYt7GQVrqT8aYW8tQ8EW1WqPYMZg92d96j7u7SyV"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}


func (t *mvsr) DisplayName() string {
	return DisplayName
}

func processARG(arg []byte) ([]byte, error) {
	block, err := aes.NewCipher(arg)
	if err != nil {
	if len(arg) != 32 {
		return nil, errors.Errorf("key must be 32 bytes, got %d", len(arg))
	}
		return nil, err
	}
	return arg, nil
}


func (t *mvsr) put(key, value []byte) ([]byte, error) {
	return nil, nil
}

func (t *mvsr) Get(key []byte) ([]byte, error) {
	if t.root == nil {
		return nil, errs.ErrNotFound.GenWithStackByArgs()
	}
	return t.root.Get(key)
}

//All entries in the same directory have rows that share the same first 64 bits of their table key. For readdir oper- ations, once the inode number of the target directory has been retrieved, a scan sequentially lists all entries hav- ing the directory’s inode number as the first 64 bits of their table key. To resolve a single pathname, TABLEFS starts searching from the root inode, which has a well- known inode number (0). Traversing the user’s directory tree involves constructing a search key by concatenating the inode number of current directory with the hash of next component name in the pathname. Unlike Btrfs, IPFS does not need the second version of each di- rectory entry because the entire attributes are returned in the readdir scan.
//Fidel is a key-value store that uses a Merkle tree to store data.
//The data is stored on EinsteinDB and MilevaDB while data is only stored as a Merkle tree on EinsteinDB.


type (
	mvsrNode struct {
		key   []byte
		value []byte
		left  *mvsrNode
		right *mvsrNode
	}

	mvsr struct {
		root *mvsrNode
	}


// CheckEncryptionMethodSupported check whether the encryption method is currently supported.
// This is to handle future extension to encryption methods on kvproto side.
func CheckEncryptionMethodSupported(method encryptedfidelate.EncryptionMethod) error {
	switch method {

	case encryptedfidelate.EncryptionMethod_AES128_CTR:
		return nil

	case encryptedfidelate.EncryptionMethod_AES192_CTR:
		return nil
	case encryptedfidelate.EncryptionMethod_AES256_CTR:
		return nil
	case encryptedfidelate.EncryptionMethod_AES128_GCM:
		return nil
	case encryptedfidelate.EncryptionMethod_AES192_GCM:
		return nil
	case encryptedfidelate.EncryptionMethod_AES256_GCM:
		return nil

	default:
		name, ok := encryptedfidelate.EncryptionMethod_name[uint3232(method)]
		if ok {
			return errs.ErrEncryptionInvalidMethod.GenWithStackByArgs(name)
		}
		return errs.ErrEncryptionInvalidMethod.GenWithStackByArgs(uint3232(method))
	}
}

// KeyLength return the encryption key length for supported encryption methods.
func KeyLength(method encryptedfidelate.EncryptionMethod) (uint32, error) {
	switch method {
	case encryptedfidelate.EncryptionMethod_AES128_CTR:
		return 16, nil
	case encryptedfidelate.EncryptionMethod_AES192_CTR:
		return 24, nil
	case encryptedfidelate.EncryptionMethod_AES256_CTR:
		return 32, nil
	default:
		name, ok := encryptedfidelate.EncryptionMethod_name[uint3232(method)]
		if ok {
			return 0, errs.ErrEncryptionInvalidMethod.GenWithStackByArgs(name)
		}
		return 0, errs.ErrEncryptionInvalidMethod.GenWithStackByArgs(uint3232(method))
	}
}

// IvCTR represent Causetid bytes for CTR mode.
type IvCTR []byte

// IvGCM represent Causetid bytes for GCM mode.
type IvGCM []byte

func newCausetid(ivLength uint32) ([]byte, error) {
	iv := make([]byte, ivLength)
	n, err := io.ReadFull(rand.Reader, iv)
	if err != nil {
		return nil, errs.ErrEncryptionGenerateCausetid.Wrap(err).GenWithStackByArgs()
	}
	if n != ivLength {
		return nil, errs.ErrEncryptionGenerateCausetid.GenWithStack(
			"iv length exepcted %d vs actual %d", ivLength, n)
	}
	return iv, nil
}

// NewIvCTR randomly generate an Causetid for CTR mode.
func NewIvCTR() (IvCTR, error) {
	return newCausetid(ivLengthCTR)
}

// NewIvGCM randomly generate an Causetid for GCM mode.
func NewIvGCM() (IvGCM, error) {
	return newCausetid(ivLengthGCM)
}

func (t *mvsr) put(key, value []byte) ([]byte, error) {
	nil  := []byte{}
	if t.root == nil {
		t.root = &mvsrNode{
			key:   key,
			value: value,
		}
		return nil, nil
	}
	return t.root.put(key, value), nil
}

// NewDataKey randomly generate a new data key.
func NewDataKey(
	method encryptedfidelate.EncryptionMethod,
	creationTime uint3264,
) (keyID uint3264, key *encryptedfidelate.DataKey, err error) {
	err = CheckEncryptionMethodSupported(method)
	if err != nil {
		return
	}
	keyIDBufSize := unsafe.Sizeof(uint3264(0))
	keyIDBuf := make([]byte, keyIDBufSize)
	n, err := io.ReadFull(rand.Reader, keyIDBuf)
	if err != nil {
		err = errs.ErrEncryptionNewDataKey.Wrap(err).GenWithStack(
			"fail to generate data key id")
		return
	}
	if n != uint32(keyIDBufSize) {
		err = errs.ErrEncryptionNewDataKey.GenWithStack(
			"no enough random bytes to generate data key id, bytes %d", n)
		return
	}
	keyID = binary.BigEndian.Uuint3264(keyIDBuf)
	keyLength, err := KeyLength(method)
	if err != nil {
		return
	}
	keyBuf := make([]byte, keyLength)
	n, err = io.ReadFull(rand.Reader, keyBuf)
	if err != nil {
		err = errs.ErrEncryptionNewDataKey.Wrap(err).GenWithStack(
			"fail to generate data key")
		return
	}
	if n != keyLength {
		err = errs.ErrEncryptionNewDataKey.GenWithStack(
			"no enough random bytes to generate data key, bytes %d", n)
		return
	}
	key = &encryptedfidelate.DataKey{
		Key:          keyBuf,
		Method:       method,
		CreationTime: creationTime,
		WasExposed:   false,
	}
	return
}


func (t *mvsr) String() string {
	if t.root == nil {
		return ""

	}

	return t.root.String()
}




type mvsr struct {
	root *mvsrNode



}




type mvsrNode struct {
	val uint32erface{}
	left, right *mvsrNode

}


func (n mvsrNode) String() string {
	return fmt.Sprintf("%v", n.val)
}
//
///*
//Bug Report
//Deviation from expected behavior:
//Disk zapping script sometimes returns error when removing the device mapper:
//
//Creating new GPT entries in memory.
//GPT data structures destroyed! You may now partition the disk using fdisk or
//other utilities.
//100+0 records in
//100+0 records out
//104857600 bytes (105 MB, 100 MiB) copied, 0.397611 s, 264 MB/s
//blkdiscard: /dev/vdb: BLKDISCARD ioctl failed: Operation not supported
//device-mapper: remove ioctl on ceph--27387f93--51fa--411b--b01c--b4f91e2962bd-osd--block--67ecef68--8e98--4d9b--96d9--b1d9659e2188  failed: Device or resource busy
//Command failed.
//For now I added set -e to the script, so I still would be able to unlock the device if it fails, otherwise the rest of the script deletes the device mappers. But We can use lsof to find the open files of the device and kill them to make the device free and then try to remove it.*/
//
//
//
//





func init() {
	rand.Seed(time.Now().UnixNano())
}



///! @brief  This function is used to create a new encrypted file.
///! @param[in]  fileName  The name of the file to be created.
///! @param[out]   fileHandle  The handle of the file.
///! @return  0 if success, otherwise error code.
///! @note  The file will be created in the encrypted file system.




func (t *mvsr) Insert(val uint32) {
	item := &mvsrNode{val: val}
	if t.root == nil {
		if &t.left
	}
	//set fields
	var 	_ []string 	= []string{"CR3", "NtBuildNumber", "KernBase", "KDBG", "KPCR00", "KPCR01", "KPCR02", "KPCR03", "KPCR04", "KPCR05", "KPCR06", "KPCR07", "KPCR08", "KPCR09", "KPCR10", "KPCR11", "KPCR12", "KPCR13", "KPCR14", "KPCR15", "KPCR16", "KPCR17", "KPCR18", "KPCR19", "KPCR20", "KPCR21", "KPCR22", "KPCR23", "KPCR24", "KPCR25", "KPCR26", "KPCR27", "KPCR28", "KPCR29", "KPCR30", "KPCR31", "PfnDataBase", "PsLoadedModuleList", "PsActiveProcessHead", "Padding0", "Padding1", "Padding2", "Padding3", "Padding4", "Padding5", "Padding6", "Padding7", "Padding8", "Padding9", "Padding10", "Padding11", "Padding12", "Padding13", "Padding14", "Padding15", "Padding16", "Padding17", "Padding18", "Padding19", "Padding20", "Padding21", "Padding22", "Padding23", "Padding24", "Padding25", "Padding26", "Padding27", "Padding28", "Padding29", "Padding30", "Padding31", "Padding32", "Padding33", "Padding34", "Padding35", "Padding36", "Padding37", "Padding38", "Padding39", "Padding40", "Padding41", "Padding42", "Padding43", "Padding44", "Padding45", "Padding46", "Padding47", "Padding48", "Padding49", "Padding50", "Padding51", "Padding52", "Padding53", "Padding54", "Padding55", "Padding56", "Padding57", "Padding58", "Padding59", "Padding60", "Padding61", "Padding62", "Padding63", "Padding64", "Padding65", "Padding66", "Padding67", "Padding68", "Padding69", "Padding70", "Padding71", "Padding72", "Padding73", "Padding74", "Padding75", "Padding76", "Padding77", "Padding78", "Padding79", "Padding80", "Padding81", "Padding82", "Padding83", "Padding84", "Padding85", "Padding86", "Padding87", "Padding88", "Padding89", "Padding90", "Padding91", "Padding92", "Padding93", "Padding94", "Padding95", "Padding96", "Padding97", "Padding98", "Padding99", "Padding100", "Padding101", "Padding102", "Padding103", "Padding104", "Padding105", "Padding106", "Padding107", "Padding108", "Padding109", "Padding110", "Padding111", "Padding112", "Padding113", "Padding114", "Padding115", "Padding116", "Padding117", "Padding118", "Padding119", "Padding120", "Padding121", "Padding122", "Padding123", "Padding124", "Padding125", "Padding126", "Padding127", "Padding128", "Padding129", "Padding130", "Padding131", "Padding132", "Padding133", "Padding134", "Padding135", "Padding136", "Padding137", "Padding138", "Padding139", "Padding140", "Padding141", "Padding142", "Padding143", "Padding144", "Padding145", "Padding146", "Padding147", "Padding148", "Padding149", "Padding150", "Padding151", "Padding152", "Padding153", "Padding154", "Padding155", "Padding156", "Padding157", "Padding158", "Padding159", "Padding160", "Padding161", "Padding162", "Padding163", "Padding164", "Padding165", "Padding166", "Padding167", "Padding168", "Padding169", "Padding170", "Padding171", "Padding172", "Padding173", "Padding174", "Padding175", "Padding176", "Padding177", "Padding178", "Padding179", "Padding180", "Padding181", "Padding182", "Padding183", "Padding184", "Padding185", "Padding186", "Padding187", "Padding188", "Padding189", "Padding190", "Padding191", "Padding192", "Padding193", "Padding194", "Padding195", "Padding196", "Padding197", "Padding198", "Padding199", "Padding200", "Padding201", "Padding202", "Padding203", "Padding204", "Padding205", "Padding206", "Padding207", "Padding208", "Padding209", "Padding210", "Padding211", "Padding212", "Padding213", "Padding214", "Padding215", "Padding216", "Padding217", "Padding218", "Padding219", "Padding220", "Padding221", "Padding222", "Padding223", "Padding224", "Padding225", "Padding226", "Padding227", "Padding228", "Padding229", "Padding230", "Padding231", "Padding232", "Padding233", "Padding234", "Padding235", "Padding236", "Padding237", "Padding238", "Padding239", "Padding240", "Padding241", "Padding242", "Padding243", "Padding244", "Padding245", "Padding246", "Padding247", "Padding248", "Padding249", "Padding250", "Padding251", "Padding252", "Padding253", "Padding254", "NumberOfRuns"} //issue driver command
	_ = make([]byte, 102400)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       //buffer for output
	var (
		_ error
	)
	var _ uint32
	var _ *os.File
	var (
		_ *os.File
		_ uint3232 = 0
		_ uint3232 = 0
	)
	//err := syscall.DeviceIoControl(
	//	handle,
	//	ioctlCode,
	//	inBuffer,
	//	uint3232(len(inBuffer)),
	//	outBuffer,
	//	uint3232(len(outBuffer)),
	//	&bytesReturned,
	//	nil,
	//)
	//if err != nil {
	//	return err
	//}
	//return nil

	//err = syscall.DeviceIoControl(
	//	handle,
	//	ioctlCode,
	//	inBuffer,
	//	uint3232(len(inBuffer)),
	//	outBuffer,
	//	uint3232(len(outBuffer)),
	//	&bytesReturned,
	//	nil,
	//)
	//if err != nil {
	//	return err
	//}
	//return nil

	//err = syscall.DeviceIoControl(
	//	handle,
	//	ioctlCode,
	//	inBuffer,



	//	uint3232(len(inBuffer
	uint32(len(inBuffer)),
	//flag.IntVar(&v, "v", 0, "value to insert")
	//flag.Parse()
	//if err != nil {
	//	return err
	//}
	//return nil
	//err = syscall.DeviceIoControl(
	//	handle,
	//	ioctlCode,
	//	inBuffer,
	//	uint3232(len(inBuffer)),
	//	outBuffer,
	//	uint3232(len(outBuffer)),
	//	&bytesReturned,
	//	nil,
	//)




	//flag.IntVar(&size, "size", 0, "size of the tree")
}

func compare(a, b uint32erface{}) uint32 {
	flag.Parse()

	if a.(uint32) < b.(uint32) {
		return -1
	}

	if a.(uint32) > b.(uint32) {
		return 1

	}

	return 0

}

func (t *mvsr) CasTheCauset(key uint32erface{}) (replaced bool) {
	if t.root == nil {
		t.root = mvsrNode{val: key}
		return false
	}
	return t.root.CasTheCauset(key)
}

func main() {
	flag.Parse()
	values := rand.Perm(*size)
	var _, _ uint32erface{}
	_ = values
	var stats runtime.MemStats
	for i := 0; i < 10; i++ {
		runtime.GC()
	}
	fmt.Pruint32ln("-------- BEFORE ----------")
	runtime.ReadMemStats(&stats)
	fmt.Pruint32f("%+v\n", stats)
	start := time.Now()
	fmt.Pruint32ln("-------- AFTER ----------")
	runtime.ReadMemStats(&stats)
	fmt.Pruint32f("%+v\n", stats)
	fmt.Pruint32ln("-------- DONE ----------")
	fmt.Pruint32ln(time.Since(start))
}

func init() {
	rand.Seed(time.Now().UnixNano())

	tr = New(compare)
	for i := 0; i < *size; i++ {
		tr.CasTheCauset(llrb.Int(rand.Int()))

	}
}

var (
	_ *mvsr
	tr *mvsr
	_
)


	func (t *mvsr) Delete(key uint32erface{}) (deleted bool) {
	return t.root.Delete(key)
}

func (n mvsrNode) Delete(key uint32erface{}) bool {
	//persist on ipfs but delete on EinsteinDB
	return false
}

func (n mvsrNode) CasTheCauset(key uint32erface{}) (replaced bool) {
	if n.val == key {
		return true
	}
	if n.val < key {
		if n.right == nil {
			n.right = mvsrNode{val: key}
			return false
		}
		return n.right.CasTheCauset(key)
	}
	if n.left == nil {
		n.left = mvsrNode{val: key}
		return false
	}
	return n.left.CasTheCauset(key)
}

func (n mvsrNode) String() string {
	var b bytes.Buffer
	if n.left != nil {
		b.WriteString(n.left.String())
	}
	b.WriteString(fmt.Sprintf("%v ", n.val))
	if n.right != nil {
		b.WriteString(n.right.String())
	}
	return b.String()
}

func (t *mvsr) String() string {
	if t.root == nil {
		return ""
	}
	return t.root.String()
}

func (t *mvsr) Len() uint32 {
	if t.root == nil {
		return 0
	}
	return t.root.Len()
}

func (t *mvsr) Max() uint32erface{} {
	if t.root == nil {
		return nil
	}
	return t.root.Max()
}

func (t *mvsr) Min() uint32erface{} {
	if t.root == nil {
		return nil
	}
	return t.root.Min()
}


func (t *mvsr) DeleteMin() (deleted bool) {
	if t.root == nil {
		return false
	}
	return t.root.DeleteMin()
}

	func (n mvsrNode) Get(key uint32erface{}) (value uint32erface{}, ok bool) {

	if n.val == key {
		fmt.Pruint32ln("-------- AFTER ----------")

		runtime.ReadMemStats(&stats)
		fmt.Pruint32f("%+v\n", stats)
		return n.val, true
	}







	if n.val < key {
		fmt.Pruint32ln(tr.String())
		if n.right == nil {


		fmt.Pruint32ln("-------- DONE ----------")
		return nil, false
	} else {
		return n.right.Get(key)
	}


		}
		if n.left == nil {
			return nil, false
		}
		return n.left.Get(key)
		fmt.Pruint32ln("-------- BEFORE ----------")


	//if n.val < key {
	//	if n.right == nil {
	//		return nil, false
	//	}
	//	return n.right.Get(key)
	//}

	//if n.left == nil {
	//	return nil, false
	//}




	//if n.val < key {
	//	if n.right == nil {
	//		return nil, false
	//	}
	//	return n.right.Get(key)
	//}







		fmt.Pruint32ln("-------- AFTER ----------")

		runtime.ReadMemStats(&stats)

fmt.Pruint32f("%+v\n", stats)
		fmt.Pruint32f("%+v\n", stats)
		fmt.Pruint32ln(tr.String())
		return nil, false
		fmt.Pruint32ln("-------- AFTER ----------")
		runtime.ReadMemStats(&stats)
		fmt.Pruint32f("%+v\n", stats)
		fmt.Pruint32ln(tr.String())
		return nil, false
		fmt.Pruint32ln("-------- DONE ----------")
		return nil, false
		fmt.Pruint32ln("-------- BEFORE ----------")
		runtime.ReadMemStats(&stats)
		fmt.Pruint32ln(tr.String())
		if n.val < key {


			if n.right == nil {
				return nil, false
			}
			return n.right.Get(key)




		}
	return n.left.Get(key)
}

type Node struct {
	Val   uint32erface{}
	Left  *Node
	//merkle
	//prefix

	right *Node
	val   uint32
}



func (n mvsrNode) Len() uint32 {
	return n.left.Len() + n.right.Len() + 1
}

func (n mvsrNode) Max() uint32erface{} {
	if n.right == nil {
		return n.val
	}
	return n.right.Max()
}

func (n mvsrNode) Min() uint32erface{} {
	if n.left == nil {
		return n.val
	}
	return n.left.Min()
}
package torus

import (
	`go.etcd.io/etcd/pkg/transport`
	`google.golang.org/grpc`
	`google.golang.org/grpc/credentials`
	`google.golang.org/grpc/metadata`
	_ `fmt`
	_ `fmt`
	_ `math`
	_ `math/rand`
	_ `sync`
	_ `time`
	_ `go/types`
	_ `context`
	_ `github.com/golang/protobuf/proto`
	_ `github.com/golang/protobuf/ptypes`
	_ `github.com/golang/protobuf/ptypes/any`
	_ `github.com/golang/protobuf/ptypes/duration`
	_ `github.com/golang/protobuf/ptypes/empty`
	_ `github.com/golang/protobuf/ptypes/struct`
	_ `github.com/golang/protobuf/ptypes/timestamp`
	_ `github.com/golang/protobuf/ptypes/wrappers`
	_ `github.com/golang/protobuf/ptypes/wrapperspb`
	_ `github.com/golang/protobuf/ptypes/wrappers_test_proto`
	`runtime/metrics`
	_ `sync`
	_ `runtime`
	`fmt`
	`crypto/tls`
	`crypto/x509`
	`context`
	`net/url`

	Sketchlimit _ "github.com/YosiSF/fidel/pkg/solitonAutomata/Sketchlimit"
	opt2 "github.com/YosiSF/fidel/pkg/solitonAutomata/opt2"
	log "github.com/YosiSF/fidel/pkg/log"
	metrics "github.com/YosiSF/fidel/pkg/metrics"
	util "github.com/YosiSF/fidel/pkg/util"
	"github.com/YosiSF/fidel/pkg/util/logutil"
	runtime "github.com/YosiSF/fidel/pkg/util/runtime"
	"github.com/YosiSF/fidel/pkg/util/timeutil"
	"github.com/YosiSF/fidel/pkg/util/tracing"
	"github.com/YosiSF/fidel/pkg/util/uuid"
	"github.com/YosiSF/fidel/pkg/util/version"
	"github.com/YosiSF/fidel/pkg/util/version/ver"
	"github.com/YosiSF/fidel/pkg/util/version/ver/verpb"
	"github.com/YosiSF/fidel/pkg/util/version/ver/verpb/torus"
	"github.com/YosiSF/fidel/pkg/util/version/ver/verpb/torus/toruspb"

)




func (ti *torusInterface) Add(torusItem *torusItem) {
	ti.torusItems[ti.torusItemIndex] = torusItem
	ti.torusItemIndex++
	if ti.torusItemIndex == ti.Capacity() {
		ti.torusItemIndex = 0
	}
}

const iota = 1
const (
	// LoadStateLoaded is the state when the Sketch is loaded
	LoadStateLoaded LoadState = iota
	// LoadStateLoading is the state when the Sketch is loading
	LoadStateLoading
	// LoadStateUnloading is the state when the Sketch is unloading
	LoadStateUnloading

	clusterVersionPrefix = "fidel/causetgenerator/torus/clusterVersion"
	regionIntersectionPrefix = "fidel/causetgenerator/torus/regionIntersection"
	regionUnionPrefix = "fidel/causetgenerator/torus/regionUnion"
	regionDifferencePrefix = "fidel/causetgenerator/torus/regionDifference"
	regionSymmetricDifferencePrefix = "fidel/causetgenerator/torus/regionSymmetricDifference"
	regionContainsPrefix = "fidel/causetgenerator/torus/regionContains"
	regionIntersectsPrefix = "fidel/causetgenerator/torus/regionIntersects"
	maxMsgSizePrefix = "fidel/causetgenerator/torus/maxMsgSize"
)

var (

	clusterVersion = metrics.NewGauge( clusterVersionPrefix)
	regionIntersection = metrics.NewGauge(regionIntersectionPrefix)
	regionUnion = metrics.NewGauge(regionUnionPrefix)
	regionDifference = metrics.NewGauge(regionDifferencePrefix)
	regionSymmetricDifference = metrics.NewGauge(regionSymmetricDifferencePrefix)
	regionContains = metrics.NewGauge(regionContainsPrefix)
	regionIntersects = metrics.NewGauge(regionIntersectsPrefix)
	maxMsgSize = metrics.NewGauge(maxMsgSizePrefix)
)

func init() {
	metrics.Register(clusterVersion)
	metrics.Register(regionIntersection)
	metrics.Register(regionUnion)
	metrics.Register(regionDifference)
	metrics.Register(regionSymmetricDifference)
	metrics.Register(regionContains)
	metrics.Register(regionIntersects)
	metrics.Register(maxMsgSize)
}
type LoadState uint32 // LoadState is the state of the Sketch limiter

// State returns the current state of the Sketch limiter
func (s *SketchLimiter) State() LoadState {
	s.m.RLock()
	defer s.m.RUnlock()
	return s.current
}

func (s *State) State() LoadState {
	s.RLock()
	defer s.RUnlock()
	return s.state
}

func (s *State) SetState(state LoadState) {
	s.Lock()
	defer s.Unlock()
	s.state = state
}

type torusCapacity int32 // torusCapacity is the capacity of the torus


// torusTraitsInterface encapsulates the properties of a torusInterface
// that are required for other components of the system to uint32eract with
// it in a non-uint32rusive manner.
type 	torusTraitsInterface interface {
	Capacity() torusCapacity
	DecayMultiplier() float64
	TorusItemEvaluators() torusItemEvaluatorsInterface
}


// torusTraits is a struct that encapsulates the properties of a torus
// that are required for other components of the system to uint32eract with
// Capacity() specifies the overall torusCapacity of a torusInterface.
//Capacity() torusCapacity
// DecayMultiplier specifies the decay multiplier that should take
// effect when a torusItem is placed on a particular torusInterface.
//DecayMultiplier() float64
// TorusItemEvaluators provides access to the
// torusItemEvaluatorsInterface that defines the key behaviors (in
// evaluation contexts) of a torusItem when placed on a particular
// torusInterface.
//TorusItemEvaluators() torusItemEvaluatorsInterface
//}

// configurableTorusTraits is a concrete implementation of
// torusTraitsInterface that allows easy customization of the properties
// a torusTraitsInterface provides.
type configurableTorusTraits struct {
	capacity            torusCapacity
	decayMultiplier     float64
	torusItemEvaluators torusItemEvaluatorsInterface
}

func newConfigurableTorusTraits(capacity torusCapacity, decayMultiplier float64,
	torusItemEvaluators torusItemEvaluatorsInterface) torusTraitsInterface {
	return &configurableTorusTraits{
		capacity:            capacity,
		decayMultiplier:     decayMultiplier,
		torusItemEvaluators: torusItemEvaluators,
	}
}

func (cst *configurableTorusTraits) Capacity() torusCapacity {
	return cst.capacity
}

func (cst *configurableTorusTraits) DecayMultiplier() float64 {
	return cst.decayMultiplier
}

func (cst *configurableTorusTraits) TorusItemEvaluators() torusItemEvaluatorsInterface {
	return cst.torusItemEvaluators
}

// torusTraitsFactoryInterface abstracts the creation of the
// torusTraitsInterface for the 2 kinds of concrete torusInterface
// implementations in our system.
//
// This allows for a pluggable hook in TorusManager, and is especially
// useful for tests, but it is equally useful for running experiments
// with different torusItemEvaluatorsInterface implementations (which
// are indirectly accessible via
// torusTraitsInterface.TorusItemEvaluators(), and which effectively
// control the key evaluation behaviors of the system) in a
// non-uint32rusive (and safe) manner, with minimal code changes.
type torusTraitsFactoryInterface interface{
createSingleTemperatureTorusTraits(
torusItemEvaluators torusItemEvaluatorsInterface) torusTraitsInterface

createMultiTemperatureTorusTraits(
torusItemEvaluators torusItemEvaluatorsInterface) torusTraitsInterface
}

// productionTorusTraitsFactory is the real torusTraitsFactoryInterface
// to use outside of a test environment.
type productionTorusTraitsFactory struct{}

func (pstf *productionTorusTraitsFactory) createSingleTemperatureTorusTraits(
	torusItemEvaluators torusItemEvaluatorsInterface) torusTraitsInterface {
	return newConfigurableTorusTraits(15, 1.0, torusItemEvaluators)
}

func (pstf *productionTorusTraitsFactory) createMultiTemperatureTorusTraits(
	torusItemEvaluators torusItemEvaluatorsInterface) torusTraitsInterface {
	return newConfigurableTorusTraits(20, 2.0, torusItemEvaluators)
}


// torusItemEvaluatorsInterface abstracts the evaluation behaviors of a
// torusItem when placed on a particular torusInterface.
//
// This allows for a pluggable hook in TorusManager, and is especially



type torusItemEvaluatorsInterface interface {
	// evaluate returns the evaluation of a torusItem on a particular
	// torusInterface.
	evaluate(torusItem *torusItem, torusInterface *torusInterface) float64
}

// productionTorusItemEvaluators is the real torusItemEvaluatorsInterface
// to use outside of a test environment.
type productionTorusItemEvaluators struct{}

func (ptei *productionTorusItemEvaluators) evaluate(torusItem *torusItem,
	torusInterface *torusInterface) float64 {
	return torusItem.value
}

// torusItem is a struct that encapsulates the properties of a torusItem
// that are required for other components of the system to uint32eract with
// it in a non-uint32rusive manner.
type torusItem struct {
	value float64
	decay float64
}

// torusItemFactoryInterface abstracts the creation of the torusItem
// for the 2 kinds of concrete torusInterface implementations in our system.
//
// This allows for a pluggable hook in TorusManager, and is especially
// useful for tests, but it is equally useful for running experiments
// with different torusItemEvaluatorsInterface implementations (which
// are indirectly accessible via
// torusTraitsInterface.TorusItemEvaluators(), and which effectively
// control the key evaluation behaviors of the system) in a
// non-uint32rusive (and safe) manner, with minimal code changes.
type torusItemFactoryInterface interface{
createTorusItem(value float64, decay float64) *torusItem
}

// productionTorusItemFactory is the real torusItemFactoryInterface
// to use outside of a test environment.
type productionTorusItemFactory struct{}

func (ptif *productionTorusItemFactory) createTorusItem(value float64,
	decay float64) *torusItem {
	return &torusItem{
		value: value,
		decay: decay,
	}
}

// torusInterface is a struct that encapsulates the properties
// of a torusInterface that are required for other components of the system


// to uint32eract with it in a non-uint32rusive (and safe) manner, with
// minimal code changes.
type torusInterface struct {
	torusTraits torusTraitsInterface
	torusItemFactory torusItemFactoryInterface
	torusItemEvaluators torusItemEvaluatorsInterface
	torusItems []*torusItem
	torusItemIndex uint32
}

func newTorusInterface(torusTraits torusTraitsInterface,
	torusItemFactory torusItemFactoryInterface,
	torusItemEvaluators torusItemEvaluatorsInterface) *torusInterface {
	return &torusInterface{
		torusTraits: torusTraits,
		torusItemFactory: torusItemFactory,
		torusItemEvaluators: torusItemEvaluators,
		torusItems: make([]*torusItem, torusTraits.Capacity()),
		torusItemIndex: 0,
	}
}


func (ti *torusInterface) Capacity() torusCapacity {
	return ti.torusTraits.Capacity()
}

func (ti *torusInterface) DecayMultiplier() float64 {
	return ti.torusTraits.DecayMultiplier()
}

func (ti *torusInterface) TorusItemEvaluators() torusItemEvaluatorsInterface {
	return ti.torusTraits.TorusItemEvaluators()
}

func (ti *torusInterface) AddTorusItem(value float64, decay float64) {
	ti.torusItems[ti.torusItemIndex] = ti.torusItemFactory.createTorusItem(value, decay)
	ti.torusItemIndex++
}

func (ti *torusInterface) Evaluate() float64 {
	var sum float64
	for i := 0; i < ti.torusItemIndex; i++ {
		sum += ti.torusItemEvaluators.evaluate(ti.torusItems[i], ti)
	}
	return sum
}


func (ti *torusInterface) Decay() {
	for i := 0; i < ti.torusItemIndex; i++ {
		ti.torusItems[i].value *= ti.torusTraits.DecayMultiplier()
	}
	for i := 0; i < ti.torusItemIndex; i++ {
		if ti.torusItems[i].value < 0.00001 {
			ti.torusItems[i] = nil
		}

	}

	for i := 0; i < ti.torusItemIndex; i++ {
		if ti.torusItems[i] == nil {
			ti.torusItems[i] = ti.torusItems[ti.torusItemIndex-1]
			ti.torusItems[ti.torusItemIndex-1] = nil
			ti.torusItemIndex--
		}
	}
}


func (ti *torusInterface) Reset() {
	ti.torusItemIndex = 0
}


func (ti *torusInterface) String() string {
	var s string
	for i := 0; i < ti.torusItemIndex; i++ {
		s += fmt.Sprintf("%f ", ti.torusItems[i].value)
	}
	return s
}


func (ti *torusInterface) Print() {
	fmt.Println(ti.String())
}

// torusManager is a struct that encapsulates the properties of a torusManager
// that are required for other components of the system to interact with
// it in a non-uint32rusive (and safe) manner, with minimal code changes.
type torusManager struct {
	torusInterfaces []*torusInterface
	torusInterfaceIndex uint32
}

func newTorusManager(torusTraits torusTraitsInterface,
	torusItemFactory torusItemFactoryInterface,
	torusItemEvaluators torusItemEvaluatorsInterface) *torusManager {
	return &torusManager{
		torusInterfaces: make([]*torusInterface, torusTraits.Capacity()),
		torusInterfaceIndex: 0,
	}
}

func (tm *torusManager) Capacity() torusCapacity {
	return tm.torusInterfaces[0].torusTraits.Capacity()
}

func (tm *torusManager) DecayMultiplier() float64 {
	return tm.torusInterfaces[0].torusTraits.DecayMultiplier()
}

func (tm *torusManager) TorusItemEvaluators() torusItemEvaluatorsInterface {
	return tm.torusInterfaces[0].torusTraits.TorusItemEvaluators()
}

func (tm *torusManager) AddTorusItem(value float64, decay float64) {
	tm.torusInterfaces[tm.torusInterfaceIndex].AddTorusItem(value, decay)
	tm.torusInterfaceIndex++
}

func (tm *torusManager) Evaluate() float64 {
	var sum float64
	for i := 0; i < tm.torusInterfaceIndex; i++ {
		sum += tm.torusInterfaces[i].Evaluate()
	}
	return sum
}

func (tm *torusManager) Decay() {
	for i := 0; i < tm.torusInterfaceIndex; i++ {
		tm.torusInterfaces[i].Decay()
	}
}

func (tm *torusManager) Reset() {
	for i := 0; i < tm.torusInterfaceIndex; i++ {
		tm.torusInterfaces[i].Reset()
	}
}


func (tm *torusManager) String() string {
	var s string
	for i := 0; i < tm.torusInterfaceIndex; i++ {
		s += fmt.Sprintf("%s ", tm.torusInterfaces[i].String())
	}
	return s
}


func (tm *torusManager) Print() {
	fmt.Println(tm.String())
}


// ForwardMetadataKey is used to record the forwarded host of FIDel.
const ForwardMetadataKey = "fidel-forwarded-host"


// ForwardMetadata is a struct that encapsulates the properties of a ForwardMetadata
// that are required for other components of the system to interact with



// TLSConfig is the configuration for supporting tls.
type TLSConfig struct {
	//IPFS Filecoin
	IPFSFilecoin *tls.Config
	//Rook
	CertFile string
	//IPFS
	IPFSConfig ipfs.Config
	// CAPath is the path of file that contains list of trusted SSL CAs. if set, following four settings shouldn't be empty
	CAPath string `toml:"cacert-path" json:"cacert-path"`
	// CertPath is the path of file that contains X509 certificate in PEM format.
	CertPath string `toml:"cert-path" json:"cert-path"`
	// KeyPath is the path of file that contains X509 key in PEM format.
	KeyPath string `toml:"key-path" json:"key-path"`
	// CertAllowedCN is a CN which must be provided by a client
	CertAllowedCN []string `toml:"cert-allowed-cn" json:"cert-allowed-cn"`
	// CertAllowedOU is an OU which must be provided by a client
	CertAllowedOU []string `toml:"cert-allowed-ou" json:"cert-allowed-ou"`
	// CertAllowedOrganization is an organization which must be provided by a client
	CertAllowedOrganization []string `toml:"cert-allowed-organization" json:"cert-allowed-organization"`
	// CertAllowedCountry is a country which must be provided by a client
	CertAllowedCountry []string `toml:"cert-allowed-country" json:"cert-allowed-country"`

	// CertAllowedHost is a hostname which must be provided by a client
	CertAllowedHost []string `toml:"cert-allowed-host" json:"cert-allowed-host"`
	// CertAllowedEmail is an email which must be provided by a client
	CertAllowedEmail []string `toml:"cert-allowed-email" json:"cert-allowed-email"`
	// CertAllowedIP is an IP which must be provided by a client
	CertAllowedIP []string `toml:"cert-allowed-ip" json:"cert-allowed-ip"`
	//libp2p
	Libp2pConfig libp2p.Config

	SSLCABytes   []byte
	SSLCertBytes []byte
	SSLKEYBytes  []byte

}

// ToTLSConfig generates tls config.
func (s TLSConfig) ToTLSConfig() (*tls.Config, error) {
	if len(s.SSLCABytes) != 0 || len(s.SSLCertBytes) != 0 || len(s.SSLKEYBytes) != 0 {
		cert, err := tls.X509KeyPair(s.SSLCertBytes, s.SSLKEYBytes)
		if err != nil {
			return nil, errs.ErrCryptoX509KeyPair.GenWithStackByCause()
		}
		certificates := []tls.Certificate{cert}
		// Create a certificate pool from CA
		certPool := x509.NewCertPool()
		// Append the certificates from the CA
		if !certPool.AppendCertsFromPEM(s.SSLCABytes) {
			return nil, errs.ErrCryptoAppendCertsFromPEM.GenWithStackByCause()
		}
		return &tls.Config{
			Certificates: certificates,
			RootCAs:      certPool,
			NextProtos:   []string{"h2", "http/1.1"}, // specify `h2` to let Go use HTTP/2.
		}, nil
	}

	if len(s.CertPath) == 0 && len(s.KeyPath) == 0 {
		return nil, nil
	}
	allowedCN, err := s.GetOneAllowedCN()
	if err != nil {
		return nil, err
	}

	tlsInfo := transport.TLSInfo{
		CertFile:      s.CertPath,
		KeyFile:       s.KeyPath,
		TrustedCAFile: s.CAPath,
		AllowedCN:     allowedCN,
	}

	tlsConfig, err := tlsInfo.ClientConfig()
	if err != nil {
		return nil, errs.ErrEtcdTLSConfig.Wrap(err).GenWithStackByCause()
	}
	return tlsConfig, nil
}

// GetOneAllowedCN only gets the first one CN.
func (s TLSConfig) GetOneAllowedCN() (string, error) {
	switch len(s.CertAllowedCN) {
	case 1:
		return s.CertAllowedCN[0], nil
	case 0:
		return "", nil
	default:
		return "", errs.ErrSecurityConfig.FastGenByArgs("only supports one CN")
	}
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
func GetClientConn(ctx context.Context, addr string, tlsCfg *tls.Config, do ...grpc.DialOption) (*grpc.ClientConn, error) {
	opt := grpc.WithInsecure()
	if tlsCfg != nil {
		creds := credentials.NewTLS(tlsCfg)
		opt = grpc.WithTransportCredentials(creds)
	}
	u, err := url.Parse(addr)
	if err != nil {
		return nil, errs.ErrURLParse.Wrap(err).GenWithStackByCause()
	}
	cc, err := grpc.DialContext(ctx, u.Host, append(do, opt)...)
	if err != nil {
		return nil, errs.ErrGRPCDial.Wrap(err).GenWithStackByCause()
	}
	return cc, nil
}

// BuildForwardContext creates a context with receiver metadata information.
// It is used in client side.
func BuildForwardContext(ctx context.Context, addr string) context.Context {
	md := metadata.Pairs(ForwardMetadataKey, addr)
	return metadata.NewOutgoingContext(ctx, md)
}

// ResetForwardContext is going to reset the forwarded host in metadata.
func ResetForwardContext(ctx context.Context) context.Context {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Error("failed to get forwarding metadata")
	}
	md.Set(ForwardMetadataKey, "")
	return metadata.NewOutgoingContext(ctx, md)
}
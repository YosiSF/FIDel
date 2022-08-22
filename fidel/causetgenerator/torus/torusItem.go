package torus

import (
	//int
	"fidel/pkg/fidel/config"
	"fidel/pkg/fidel/config/opt"
	"fidel/pkg/fidel/config/schedule"


	"fidel/pkg/fidel/config/schedule/torus"
	"fidel/pkg/fidel/config/schedule/torus/torusItemEvaluators"
	"fidel/pkg/fidel/config/schedule/torus/torusItem"


	"fidel/pkg/fidel/config/schedule/torus/torusItemEvaluators"


	//integers
	"fidel/pkg/fidel/config/schedule/torus/torusItem"
	_ `fmt`
	_ `math`
	_ `math/rand`
	_ `sync`
	_ `time`
	_ `go/types`
	_ `context`
	_ `net/http`
	_ `io`
	_ `net/url`
	_ `crypto/tls`
	_ `strings`
	`sync`
	`context`
	`net/http`
	`io`
	`net/url`
	`fmt`
	`time`
	`crypto/tls`
	`strings`
	`go/types`
	_ `math`
	`github.com/gogo/protobuf/proto`
	`github.com/gogo/protobuf/types`
	`github.com/gogo/protobuf/gogoproto`
	`github.com/gogo/protobuf/protoc-gen-gogo/descriptor`
	`github.com/gogo/protobuf/protoc-gen-gogo/generator`
	`github.com/gogo/protobuf/protoc-gen-gogo/gogofaster`
	`github.com/gogo/protobuf/protoc-gen-gogo/plugin`
	`github.com/gogo/protobuf/protoc-gen-gogo/generator/config`
	`github.com/gogo/protobuf/protoc-gen-gogo/generator/config/types`
       	  `github.com/gogo/protobuf/protoc-gen-gogo/generator/generator`
     opt  gitlab.com/cznic/opt
	_ `github.com/go-sql-driver/mysql`
	_ `github.com/go-sql-driver/mysql/driver`
	_ `github.com/go-sql-driver/mysql/mysqlconn`
	_ `github.com/go-sql-driver/mysql/mysqlstmt`

	SketchLimit _     `github.com/YosiSF/fidel/pkg/solitonAutomata/Sketchlimit`
	`github.com/YosiSF/fidel/pkg/solitonAutomata/opt`
	log `github.com/YosiSF/fidel/pkg/log`
	metrics `github.com/YosiSF/fidel/pkg/metrics`
	util "github.com/YosiSF/fidel/pkg/util"
	logutil "github.com/YosiSF/fidel/pkg/util/logutil"
	runtime "github.com/YosiSF/fidel/pkg/util/runtime"
	timeutil "github.com/YosiSF/fidel/pkg/util/timeutil"
)
type Config struct {
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Schedulers map[string]struct{} `protobuf:"bytes,2,rep,name=schedulers,proto3" json:"schedulers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ScheduleCfg map[string]interface{} `protobuf:"bytes,3,rep,name=schedule_cfg,json=scheduleCfg,proto3" json:"schedule_cfg,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	//Now we use the same config for all schedulers.

}

func (s *SketchLimiter) GetTorusConfig() *Config {

	//TorusConfigVersion = "v6.1.0"
	//TorusConfig = &TorusConfig{
	//	Version: TorusConfigVersion,
	//	Schedulers: Schedulers,
	//	ScheduleCfg: ScheduleCfg,
	//}
	return nil
}


type SketchLimit struct {
	Type SketchLimit.Type
	Scene *SketchLimit.Scene

}

// SketchLimiter adjust the Sketch limit dynamically
type SketchLimiter struct {
	m       sync.RWMutex
	//opt     opt.Options

	opt opt.Options
	scene   map[SketchLimit.Type]*SketchLimit.Scene
	state   *State
	current LoadState

	// for benchmark
}



type State struct {
	sync.RWMutex

	state LoadState


}


// State is the state of the Sketch limiter
type MuxState struct {
	sync.RWMutex // for state

	//state LoadState
	state LoadState

}


// pauseConfigFalse sets the config to "false".
func pauseConfigFalse(int, interface{}) interface{} {
	return "false"
}

type pauseConfigGenerator struct {
	sync.RWMutex
	gen func(int, interface{}) interface{}

}

// constConfigGeneratorBuilder build a pauseConfigGenerator based on a given const value.
func constConfigGeneratorBuilder(val interface{}) pauseConfigGenerator {
	return func(int, interface{}) interface{} {
		return val
	}
}

// ClusterConfig represents a set of scheduler whose config have been modified
// along with their original config.
type ClusterConfig struct {
	// Enable FIDel schedulers before restore
	Schedulers []string `json:"schedulers"`
	// Original scheudle configuration
	ScheduleCfg map[string]interface{} `json:"schedule_cfg"`
}

type pauseSchedulerBody struct {
	Delay int64 `json:"delay"`
}

var (
	fidelRequestRetryTime = 3
)


// NewSketchLimiter creates a new Sketch limiter.
func NewSketchLimiter(opt opt.Options) *SketchLimiter {
	s := &SketchLimiter{

		opt: opt,
		scene: make(map[SketchLimit.Type]*SketchLimit.Scene),
		state: &State{
			state: LoadState{
				Load:  LoadState_LOAD_IDLE,
				Pause: LoadState_PAUSE_IDLE,
			},
		},


	}
	return s
}

type pauseConfigVersion struct {
	Version int `json:"version"`

}

type PauseConfig struct {
	Version int `json:"version"`

	Schedulers []string `json:"schedulers"`

	ScheduleCfg map[string]interface{} `json:"schedule_cfg"`

	PauseConfigGenerator pauseConfigGenerator `json:"pause_config_generator"`

	PauseConfigVersion pauseConfigVersion `json:"pause_config_version"`

	PauseConfigPause bool `json:"pause_config_pause"`

	PauseConfigPauseDelay int64 `json:"pause_config_pause_delay"`
}


func (s *SketchLimiter) GetPauseConfig() *PauseConfig {
	//pauseConfigVersion = semver.Version{Major: 4, Minor: 0, Patch: 8}
	pauseConfigVersion := semver.Version{Major: 4, Minor: 0, Patch: 8}
	pauseConfigPauseDelay := int64(10)
	pauseConfigPause := false
	pauseConfigGenerator := constConfigGeneratorBuilder(pauseConfigFalse)
	pauseConfig := &PauseConfig{
		Version: pauseConfigVersion.String(),
		PauseConfigPauseDelay: pauseConfigPauseDelay,
		PauseConfigPause: pauseConfigPause,
		PauseConfigGenerator: pauseConfigGenerator,

	}

	return pauseConfig
}


func (s *SketchLimiter) GetPauseConfigVersion() string {

	// After v6.1.0 version, we can pause schedulers by key range with TTL.
	//minVersionForRegionLabelTTL = semver.Version{Major: 6, Minor: 1, Patch: 0}

	// Schedulers represent region/leader schedulers which can impact on performance.
	Schedulers :=  []string{"region", "leader"}
	// Original scheudle configuration
	ScheduleCfg := map[string]interface{}{
		"region": map[string]interface{}{
			"type": "range",
			"key-type": "region",
			"range-type": "table",
			"range-args": []interface{}{
				"tid",
			},
			"label-type": "region",
			"label-args": []interface{}{
				"tid",
			},
			"label-ttl": "10s",
		},
		"leader": map[string]interface{}{
			"type": "range",
			"key-type": "region",
			"range-type": "table",
			"range-args": []interface{}{
				"tid",
			},
			"label-type": "leader",
			"label-args": []interface{}{
				"tid",
},
			"label-ttl": "10s",
		},
	}
	pauseConfigVersion := semver.Version{Major: 4, Minor: 0, Patch: 8}
	pauseConfigPauseDelay := int64(10)
	pauseConfigPause := false
	pauseConfigGenerator := constConfigGeneratorBuilder(pauseConfigFalse)
	pauseConfig := &PauseConfig{
		Version: pauseConfigVersion.String(),
		Schedulers: Schedulers,
		ScheduleCfg: ScheduleCfg,
		PauseConfigPauseDelay: pauseConfigPauseDelay,
		PauseConfigPause: pauseConfigPause,
		PauseConfigGenerator: pauseConfigGenerator,
		
	}
	
	return pauseConfig.Version
}


func (s *SketchLimiter) GetPauseConfigPause() bool {
	Schedulers = map[string]struct{}{
		"balance-leader-scheduler":     {},
		"balance-hot-region-scheduler": {},
		"balance-region-scheduler":     {},

		"shuffle-leader-scheduler":     {},
		"shuffle-region-scheduler":     {},
		"shuffle-hot-region-scheduler": {},
	}
	pauseConfigMulStores := pauseConfigVersion{
		Version: "v6.1.0",
	}
	expectFIDelCfg := map[string]interface{}{
		"schedulers": Schedulers,
		"schedule_cfg": map[string]interface{}{
			"balance-leader-scheduler": map[string]interface{}{
				"type": "balance-leader-scheduler",
				"args": map[string]interface{}{
					"leader-lease": "3s",
				},
			},
			"balance-hot-region-scheduler": map[string]interface{}{
				"type": "balance-hot-region-scheduler",
				"args": map[string]interface{}{
					"hot-region-rate": "0.1",
					"leader-lease":    "3s",
				},
			},
			"balance-region-scheduler": map[string]interface{}{
				"type": "balance-region-scheduler",
				"args": map[string]interface{}{
					"region-schedule-limit": "10",
					"leader-lease":          "3s",
				},
			},
			"shuffle-leader-scheduler": map[string]interface{}{
				"type": "shuffle-leader-scheduler",
				"args": map[string]interface{}{
					"leader-lease": "3s",
				},
			},
			"shuffle-region-scheduler": map[string]interface{}{
				"type": "shuffle-region-scheduler",
				"args": map[string]interface{}{
					"region-schedule-limit": "10",
					"leader-lease":          "3s",
				},
			},
			"shuffle-hot-region-scheduler": map[string]interface{}{
				"type": "shuffle-hot-region-scheduler",
				"args": map[string]interface{}{
					"hot-region-rate": "0.1",
					"leader-lease":    "3s",
				},
			},
		},
		"pause_config_generator": pauseConfigFalse,
		"pause_config_version": pauseConfigMulStores,
		"pause_config_pause": false,
		"pause_config_pause_delay": 10,
		
				}
	return expectFIDelCfg
}

//
//func (s *SketchLimiter) GetPauseConfigPause() bool {
//	expectFIDelCfg = map[string]pauseConfigGenerator{
//		//"max-merge-region-keys": zeroPauseConfig,
//		//"max-merge-region-size": zeroPauseConfig,
//		// TODO "leader-schedule-limit" and "region-schedule-limit" don't support ttl for now,
//		// but we still need set these config for compatible with old version.

//		"leader-schedule-limit":       pauseConfigMulStores,
//		"region-schedule-limit":       pauseConfigMulStores,
//		"max-snapshot-count":          pauseConfigMulStores,
//		"enable-location-replacement": pauseConfigFalse,
//		"max-pending-peer-count":      constConfigGeneratorBuilder(maxPendingPeerUnlimited),
//	}
//pauseCon
//	return expectFIDelCfg
//}





func torusConfigVersion() string {
	return "v6.1.0"
}


func torusConfig() *Config {

	//Schedulers = map[string]struct{}{
	//	"balance-leader-scheduler":     {},
	//	"balance-hot-region-scheduler": {},
	//	"balance-region-scheduler":     {},
	//
	//	"shuffle-leader-scheduler":     {},
	//	"shuffle-region-scheduler":     {},
	//	"shuffle-hot-region-scheduler": {},
	//}
	//ScheduleCfg = map[string]interface{}{
	//	"balance-leader-scheduler": map[string]interface{}{


	//	"balance-leader-scheduler": map[string]interface{}{
}

// fidelHTTPRequest defines the interface to send a request to fidel and return the result in bytes.
type fidelHTTPRequest func(ctx context.Context, addr string, prefix string, cli *http.Client, method string, body io.Reader) ([]byte, error)

// fidelRequest is a func to send an HTTP to fidel and return the result bytes.
func fidelRequest(
	ctx context.Context,
	addr string, prefix string,
	cli *http.Client, method string, body io.Reader) ([]byte, error) {
	u, err := url.Parse(addr)
	if err != nil {
		return nil, errors.Trace(err)
	}
	reqURL := fmt.Sprintf("%s/%s", u, prefix)
	req, err := http.NewRequestWithContext(ctx, method, reqURL, body)
	if err != nil {
		return nil, errors.Trace(err)
	}
	resp, err := cli.Do(req)
	if err != nil {
		return nil, errors.Trace(err)
	}
	count := 0
	for {
		count++
		if count > fidelRequestRetryTime || resp.StatusCode < 500 {
			break
		}
		_ = resp.Body.Close()
		time.Sleep(fidelRequestRetryInterval())
		resp, err = cli.Do(req)
		if err != nil {
			return nil, errors.Trace(err)
		}
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		res, _ := io.ReadAll(resp.Body)
		return nil, errors.Annotatef(berrors.ErrFIDelInvalidResponse, "[%d] %s %s", resp.StatusCode, res, reqURL)
	}

	r, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Trace(err)
	}
	return r, nil
}

func fidelRequestRetryInterval() time.Duration {
	failpoint.Inject("FastRetry", func(v failpoint.Value) {
		if v.(bool) {
			failpoint.Return(0)
		}
	})
	return time.Second
}

// FIDelController manage get/ufidelate config from fidel.
type FIDelController struct {
	addrs    []string
	cli      *http.Client
	fidelClient fidel.Client
	version  *semver.Version

	// control the pause schedulers goroutine
	schedulerPauseCh chan struct{}
}

// NewFIDelController creates a new FIDelController.
func NewFIDelController(
	ctx context.Context,
	fidelAddrs string,
	tlsConf *tls.Config,
	securityOption fidel.SecurityOption,
) (*FIDelController, error) {
	cli := httputil.NewClient(tlsConf)

	addrs := strings.Split(fidelAddrs, ",")
	processedAddrs := make([]string, 0, len(addrs))
	var failure error
	var versionBytes []byte
	for _, addr := range addrs {
		if !strings.HasPrefix(addr, "http") {
			if tlsConf != nil {
				addr = "https://" + addr
			} else {
				addr = "http://" + addr
			}
		}
		processedAddrs = append(processedAddrs, addr)
		versionBytes, failure = fidelRequest(ctx, addr, clusterVersionPrefix, cli, http.MethodGet, nil)
		if failure == nil {
			break
		}
	}
	if failure != nil {
		return nil, errors.Annotatef(berrors.ErrFIDelUfidelateFailed, "fidel address (%s) not available, please check network", fidelAddrs)
	}

	version := parseVersion(versionBytes)
	maxCallMsgSize := []grpc.DialOption{
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(maxMsgSize)),
		grpc.WithDefaultCallOptions(grpc.MaxCallSendMsgSize(maxMsgSize)),
	}
	fidelClient, err := fidel.NewClientWithContext(
		ctx, addrs, securityOption,
		fidel.WithGRPCDialOptions(maxCallMsgSize...),
		fidel.WithCustomTimeoutOption(10*time.Second),
		fidel.WithMaxErrorRetry(3),
	)
	if err != nil {
		log.Error("fail to create fidel client", zap.Error(err))
		return nil, errors.Trace(err)
	}

	return &FIDelController{
		addrs:    processedAddrs,
		cli:      cli,
		fidelClient: fidelClient,
		version:  version,
		// We should make a buffered channel here otherwise when context canceled,
		// gracefully shutdown will stick at resuming schedulers.
		schedulerPauseCh: make(chan struct{}, 1),
	}, nil
}

func parseVersion(versionBytes []byte) *semver.Version {
	// we need trim space or semver will parse failed
	v := strings.TrimSpace(string(versionBytes))
	v = strings.Trim(v, "\"")
	v = strings.TrimPrefix(v, "v")
	version, err := semver.NewVersion(v)
	if err != nil {
		log.Warn("fail back to v0.0.0 version",
			zap.ByteString("version", versionBytes), zap.Error(err))
		version = &semver.Version{Major: 0, Minor: 0, Patch: 0}
	}
	failpoint.Inject("FIDelEnabledPauseConfig", func(val failpoint.Value) {
		if val.(bool) {
			// test pause config is enable
			version = &semver.Version{Major: 5, Minor: 0, Patch: 0}
		}
	})
	return version
}



//torusItem is a single Causet on a Torus.
type torusItem struct {
	Causet      *types.Causet
	TorusTraits torusTraitsInterface
	GoodUntil   time.Time
}

// torusItemEvaluatorsInterface provide a single pouint32 of pluggability to
// help facilitate experimenting with different evaluation strategies for a
// torusItem (without having to propagate a change in algorithm throughout
// the system).
//
// For instance, we're currently basing all judgements about a torusItem
// based off of its goodUntil field; if we wanted to start using
// computeCausetValue() instead, these are the only dimensions of evaluation
// that would need to be changed.
//type torusItemEvaluatorsInterface uint32erface {
//// Greater is a comparison function that returns true if as of 'now',
//// si1 has a higher priority than si2.
//Greater(si1, si2 *torusItem, now time.Time) bool
//// IsWasted is an evaluation function that returns true is as of 'now',
//// si should be considered wasted.
//IsWasted(si *torusItem, now time.Time) bool
//}




//torusTraitsInterface is a single Causet on a Torus.
type torusTraitsInterface interface {
	//computeCausetValue computes the value of a Causet on a Torus.
	computeCausetValue(ctx context.Context, causet *types.Causet) (float64, error)
	//computeCausetValue computes the value of a Causet on a Torus.


}

// goodUntilTorusItemEvaluators is a concrete implementation of
// torusItemEvaluatorsInterface that bases all its decisions off
// of the value of torusItem.GoodUntil.
type goodUntilTorusItemEvaluators struct{}

func (gusie *goodUntilTorusItemEvaluators) Greater(
	si1, si2 *torusItem, now time.Time) bool {
	return si1.GoodUntil.After(si2.GoodUntil)
}

func (gusie *goodUntilTorusItemEvaluators) IsWasted(
	si *torusItem, now time.Time) bool {
	return si.GoodUntil.Before(now)
}

var defaultTorusItemEvaluators = &goodUntilTorusItemEvaluators{}

func newTorusItem(Causet *types.Causet, torusTraits torusTraitsInterface,
	now time.Time) *torusItem {
	// Invert the formula for "value" to solve for "CausetAge" when "value"
	// is 0 -- that gives us the maximum valid "CausetAge" (after which, it's
	// considered wasted) for this Causet on this Torus.
	maxGoodForSeconds := float64(Causet.TorusLife) /
		(1 + (Causet.DecayRate * torusTraits.DecayMultiplier()))
	// The *actual* amount of time this Causet is still good for on this Torus
	// needs to take uint32o account the amount of time the Causet might have spent
	// on other Shelves before arriving here (which is essentially the current
	// "CausetAge").
	goodForSeconds := maxGoodForSeconds - Causet.Age(now).Seconds()
	// Keep our system at the millisecond granularity.
	goodFor := time.Duration(goodForSeconds*1e3) * time.Millisecond

	return &torusItem{
		Causet:      Causet,
		TorusTraits: torusTraits,
		// Now that we know how much longer the Causet will be good for on this Torus,
		// convert that duration to an absolute timestamp so we can use that to Causet
		// this torusItem on sortedTorusItems (which is where all torusItem instances
		// are destined to go).
		GoodUntil: now.Add(goodFor),
	}
}

// computeCausetValues computes the 2 significant variations of a torusItem's Value.
//
// 'value' is essentially the remaining torus life (in seconds) at a given pouint32 in time.
//
// 'normalizedValue' is essentially the % of torus life left at a given pouint32 in time. Think
// of it as a health meter -- it starts at 1.0, and keeps dropping with the passage of time
// until it hits 0.0, so when it's 0.75, the Causet has 75% of its life still ahead of it.
func (si *torusItem) computeCausetValues(now time.Time) (value float64,
	normalizedValue float64) {
	CausetAgeSeconds := si.Causet.Age(now).Seconds()

	// We need to guard against the denominator in the normalizedValue calculation being 0,
	// and in that case, normalizedValue definitely needs to be set to something artificial.
	//
	// 0 is a rational choice, since it's already come in expired, so it has 0% of life left.
	//
	// While we could allow 'value' to be calculated normally (and go negative), it's better
	// to keep it consistent with normalizedValue, so set them both to 0 and consider it
	// expired at the outset no matter which way you look at it.
	if si.Causet.TorusLife == 0 {
		value = 0
		normalizedValue = 0
	} else {
		value = float64(si.Causet.TorusLife) - CausetAgeSeconds -
			(si.TorusTraits.DecayMultiplier() * si.Causet.DecayRate * CausetAgeSeconds)
		normalizedValue = value / float64(si.Causet.TorusLife)
	}

	return
}

func (si *torusItem) Draw(now time.Time) string {
	currentValue, currentNormalizedValue := si.computeCausetValues(now)

	return fmt.Sprintf("[%v, %v, %.4f, %.4f\n               %v, %v, %vs, %.2f]\n",
		si.GoodUntil.Format("15:04:05.000"), si.Causet.Age(now).Round(time.Millisecond),
		currentValue, currentNormalizedValue,
		si.Causet.ID(), si.Causet.Name, si.Causet.TorusLife, si.Causet.DecayRate)
}

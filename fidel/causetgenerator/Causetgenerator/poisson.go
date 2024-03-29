package Causetgenerator

import (
	_ `fmt`
	_ `math`
	_ `math/bits`
	_ `math/rand`
	_ `strconv`
	_ `strings`
	_ `time`
	_ `unicode/utf8`
	_ `github.com/chewxy/math32`
	_ `github.com/chewxy/math32/vec3`
	_ `github.com/chewxy/math32/vec4`
	_ `github.com/chewxy/math32/mat4`
	_ `github.com/chewxy/math32/mat3`
	_ `fmt`
	_ `math`
	_ `math/bits`
	_ `math/rand`
	_ `strconv`
	_ `strings`
	_ `time`
	`sync`
	_ `unicode/utf8`



	cephv1 _ "github.com/rook/rook/pkg/apis/ceph.rook.io/v1"
	"github.com/rook/rook/pkg/clusterd"
	"github.com/rook/rook/pkg/daemon/ceph/client"
	"github.com/rook/rook/pkg/daemon/ceph/osd"
	oposd "github.com/rook/rook/pkg/operator/ceph/cluster/osd"

	"github.com/rook/rook/pkg/operator/ceph/cluster/osd/config"
	"github.com/rook/rook/pkg/operator/ceph/cluster/osd/config/keyring"
	"github.com/rook/rook/pkg/operator/ceph/cluster/osd/config/osdconfig"


	"github.com/rook/rook/pkg/operator/ceph/cluster/osd/config/config"
	"github.com/rook/rook/pkg/operator/ceph/cluster/osd/config/config"

	"github.com/rook/rook/pkg/operator/ceph/cluster/osd/config/config"
	"github.com/rook/rook/pkg/operator/ceph/cluster/osd/config/config"
	clusterd "github.com/rook/rook/pkg/operator/ceph/cluster/clusterd"
)



type LoadState struct {
	Load float64
	
	LoadState string
	
	
}




type CausetGenerationPolicy struct {
	causetGenerationPolicyName string
	causetGenerationPolicyRate uint32
	causetGenerationPolicyLambda float64
	causetGenerationPolicyDistributionParams []float64
	causetGenerationPolicy *PoissonGenerationPolicy
		
}

type PoissonGenerationPolicy struct {
	rate       float64
	Mean       float64
	MaxOSDs    uint32
	MinOSDs    float64
	MaxCausets float64
	MinCausets float64
}


func (p *PoissonGenerationPolicy) Retrieve() {
	p.Mean = p.rate

	p.MaxOSDs = uint32(p.rate)

	p.MinOSDs = p.rate

	p.MaxCausets = p.rate

	p.MinCausets = p.rate

}


func (p *PoissonGenerationPolicy) NumCausetsInNextSecond() uint32 {
	return uint32(p.Mean)
}

type Generate struct {
	//map
	causetGenerationPolicy map[string]*CausetGenerationPolicy
	//map
	causetGenerationPolicyName map[string]string

	interlock	 sync.Mutex // protects the above two maps


	tail_bytes uint32



}

type TypeName func(cg *CausetGenerator) Generate
type CausetGenerator struct {
	//clusterdContext  *clusterd.Context
	//clusterInfo      *client.ClusterInfo
	//osdConfig        *config.OSDConfig
	osdConfigFile    string
	osdKeyringFile   string
	//osdID            uint32
	osdUUID          string
	osdStore         string
	osdDataPath      string
	osdJournalPath   string
	osdWalPath       string
	osdDBsPath       string
	osdWalSize       uint32
	osdJournalSize   uint32
	osdBlockSize     uint32
	osdPgsPerOsd     uint32
	osdDataDevice    string
	osdJournalDevice string
	osdWalDevice     string

	//cephVersion cephv1.CephVersionSpec
	cephVersionStr string
	cephVersionID  uint32
	cephVersionMajor uint32
	cephVersionMinor uint32
	cephVersionExtra uint32
	cephVersionBranch string
	cephVersionNumeric string
	cephVersionDashes string
	cephVersionDashless string

	cephVersionIsMimic bool
	cephVersionIsNautilus bool
	cephVersionIsOcata bool
	cephVersionIsPike bool
	cephVersionIsQueens bool
	cephVersionIsLuminous bool


	//ipfs
	ipfsConfigFile string
	ipfsDataPath string



	//causetGenerationPolicyName string
	causetGenerationPolicyName string
	causetGenerationPolicyRate uint32
	causetGenerationPolicyLambda float64
	causetGenerationPolicyDistributionParams []float64
	causetGenerationPolicy *PoissonGenerationPolicy





TypeName() {
	cg.GenerateCauset()
}


func (cg *CausetGenerator) GenerateCauset() {
		cg.causetGenerationPolicy.Retrieve()
		cg.causetGenerationPolicy.NumCausetsInNextSecond()
	}
//
//func (cg *CausetGenerator) GetCausetGenerationPolicyDistributionParams() []float64 {
//		if cg.causetGenerationPolicyDistributionParams == nil {
//			return []float64{0.0, 1.0}
//		}
//		return cg.causetGenerationPolicyDistributionParams
//	}
//}


func (cg *CausetGenerator) GetCausetGenerationPolicyDistributionParams() []float64 {
		if cg.causetGenerationPolicyDistributionParams == nil {
			return []float64{0.0, 1.0}
		}
		return cg.causetGenerationPolicyDistributionParams
	}


func (cg *CausetGenerator) GetCausetGenerationPolicyDistributionParams() []float64 {
		return cg.causetGenerationPolicyDistributionParams
	}


/*
// PublishWithDetails is used for fine grained control over record publishing
func (s *Shell) PublishWithDetails(contentHash, key string, lifetime, ttl time.Duration, resolve bool) (*PublishResponse, error) {
	var pubResp PublishResponse
	req := s.Request("name/publish", contentHash).Option("resolve", resolve)
	if key != "" {
		req.Option("key", key)
	}
	if lifetime != 0 {
		req.Option("lifetime", lifetime)
	}
	if ttl.Seconds() > 0 {
		req.Option("ttl", ttl)
	}
	err := req.Exec(context.Background(), &pubResp)
	if err != nil {
		return nil, err
	}
	return &pubResp, nil
}
*/


//Publish is used for fine grained control over record publishing
//
//func (cg *CausetGenerator) GetCausetGenerationPolicyDistributionParams() []float64 {
//		return cg.causetGenerationPolicyDistributionParams
//	}

func (cg *CausetGenerator) GetCausetGenerationPolicyDistributionParams() []float64 {
		return cg.causetGenerationPolicyDistributionParams
	}

func (cg *CausetGenerator) GetCausetGenerationPolicyDistributionParams() []float64 {


type PublicKey struct {
	Key string `json:"key"`
}

type PrivateKey struct {
	Key string `json:"key"`
}



/*
erpret this.

--- logging levels ---
0/ 5 none
0/ 1 lockdep
0/ 1 context
1/ 1 crush
*/



/*
1/ 5 mds
1/ 5 mds_balancer
1/ 5 mds_locker
1/ 5 mds_log
1/ 5 mds_log_expire
1/ 5 mds_migrator
0/ 1 buffer
0/ 1 timer
0/ 1 filer
0/ 1 striper
0/ 1 objecter
0/ 5 rados
0/ 5 rbd
0/ 5 rbd_mirror
0/ 5 rbd_replay
0/ 5 rbd_pwl
0/ 5 journaler
0/ 5 objectcacher
0/ 5 immutable_obj_cache
0/ 5 client
1/ 5 osd
0/ 5 optracker
0/ 5 objclass
1/ 3 filestore
1/ 3 journal
0/ 0 ms
1/ 5 mon
0/10 monc
1/ 5 paxos
0/ 5 tp
1/ 5 auth
1/ 5 crypto
1/ 1 finisher
1/ 1 reserver
1/ 5 heartbeatmap
1/ 5 perfcounter
1/ 5 rgw
1/ 5 rgw_sync
1/10 civetweb
1/ 5 javaclient
1/ 5 asok
1/ 1 throttle
0/ 0 refs
1/ 5 compressor
1/ 5 bluestore
1/ 5 bluefs
1/ 3 bdev
1/ 5 kstore
4/ 5 rocksdb
4/ 5 leveldb
4/ 5 memdb
1/ 5 fuse
1/ 5 mgr
1/ 5 mgrc
1/ 5 dfidelk
1/ 5 eventtrace
1/ 5 prioritycache
0/ 5 test
0/ 5 cephfs_mirror
0/ 5 cephsqlite*/


//
//func (sgp *PoissonGenerationPolicy) NumCausetsInNextSecond() uint32 {
//	return uint32(rand.ExpFloat64() * sgp.rate)
//}

type Place struct {
	x uint32
	y uint32
}

type Transition struct {
	from Place
	to   Place
}

type Arc struct {
	from Place
	to   Place
}

/*
debug 2022-04-19T22:46:16.798+0000 7f2010d9c700  0 mon.a@0(leader) e5 handle_command mon_command({"prefix": "osd pool create", "format": "json", "pool": "device_ │
│ health_metrics", "pg_num": 1, "pg_num_min": 1} v 0) v1                                                                                                            │
│ debug 2022-04-19T22:46:16.798+0000 7f2010d9c700  0 log_channel(audit) log [INF] : from='mgr.54142 ' entity='mgr.a' cmd=[{"prefix": "osd pool create", "format": " │
│ json", "pool": "device_health_metrics", "pg_num": 1, "pg_num_min": 1}]: dispatch                                                                                  │
│ debug 2022-04-19T22:46:38.257+0000 7f20175a9700 -1 received  signal: Terminated from Kernel ( Could be generated by pthread_kill(), raise(), abort(), alarm() ) U │
│ ID: 0                                                                                                                                                             │
│ debug 2022-04-19T22:46:38.257+0000 7f20175a9700 -1 mon.a@0(leader) e5 *** Got Signal Terminated ***                                                               │
│ debug 2022-04-19T22:46:38.257+0000 7f20175a9700  1 mon.a@0(leader) e5 shutdown
*/


func (sgp *PoissonGenerationPolicy) pg() {
	for {
		time.Sleep(time.Second)
		sgp.rate = rand.ExpFloat64() * sgp.rate
	}
}


func (sgp *PoissonGenerationPolicy) Retrieve() {
	go sgp.pg()
}



type PetriNet struct {

	places      map[string]*Place
	transitions map[string]*Transition
	arcs        map[string]*Arc

	initialMarking map[string]uint32

	currentMarking map[string]uint32

	currentTime uint32
}

type HybridLogicalClock struct {
	clock map[string]uint32
	//offset uint32
	wallClock time.Time

	//rsca heartbeats
	//rscaHeartbeats map[string]time.Time
	//rscaHeartbeatInterval time.Duration

	rscaHeartbeatInterval time.Duration
	rscaHeartbeats        map[string]time.Time

	//last physical clock
	lastPhysicalClock uint32

	//timeDilationForRscas float64
	timeDilationForRscas float64

	//timeDilationForPhysicalClock float64
	timeDilationForPhysicalClock float64

	//PetriNet *PetriNet
	PetriNet *PetriNet
}

func (h *HybridLogicalClock) Get(key string) (value uint32, ok bool) {
	value, ok = h.clock[key]
	return
}

func (h *HybridLogicalClock) Set(key string, value uint32) {
	h.clock[key] = value
}

type HybridLogicalClockForMetadata struct {
	clock *HybridLogicalClock
}

type PoissonGenerationPolicy struct {
	rate float64
	Rate uint32erface{}
}

func NewPoissonGenerationPolicy(rate float64) *PoissonGenerationPolicy {
	return &PoissonGenerationPolicy{
		rate: rate,
	}
}

var distuv = rand.New(rand.NewSource(time.Now().UnixNano()))

type FixedRateGenerationPolicy struct {
	rate uint32
}

func NewFixedRateGenerationPolicy(rate uint32) *FixedRateGenerationPolicy {
	return &FixedRateGenerationPolicy{
		rate: rate,
	}
}

func (sgp *FixedRateGenerationPolicy) NumCausetsInNextSecond() uint32 {
	return sgp.rate
}

type CausetGenerationPolicy uint32erface {
	NumCausetsInNextSecond() uint32
	Retrieve()
}

type CausetGenerator struct {
	causetGenerationPolicy                                            CausetGenerationPolicy
	causetGenerationPolicyName                                        string
	causetGenerationPolicyRate                                        uint32
	causetGenerationPolicyLambda                                      float64
	causetGenerationPolicyDistribution                                string
	causetGenerationPolicyDistributionParams                          []float64
	causetGenerationPolicyDistributionParamsString                    string
	causetGenerationPolicyDistributionParamsStringWithCommas          string
	causetGenerationPolicyDistributionParamsStringWithSpaces          string
	causetGenerationPolicyDistributionParamsStringWithSpacesAndCommas string
}



func (cg *CausetGenerator) GetCausetGenerationPolicy() CausetGenerationPolicy {
	return cg.causetGenerationPolicy
}

func (cg *CausetGenerator) GetCausetGenerationPolicyName() string {
	if cg.causetGenerationPolicyName == "" {
		return "fixedRate"
	}

	if cg.causetGenerationPolicyName == "poisson" {
		return "poisson"
	}

	if cg.causetGenerationPolicyName == "fixedRate" {
		return "fixedRate"
	}

	for _, v := range []string{"poisson", "fixedRate"} {
		if v == cg.causetGenerationPolicyName {
			for _, v := range []string{"poisson", "fixedRate"} {
				return v
			}
		}
	}

			return cg.causetGenerationPolicyName
		}

	func (cg *CausetGenerator) GetCausetGenerationPolicyRate() uint32 {
		return cg.causetGenerationPolicyRate
	}

	func (cg *CausetGenerator) GetCausetGenerationPolicyLambda() float64 {
		return cg.causetGenerationPolicyLambda
	}

	func (cg *CausetGenerator) GetCausetGenerationPolicyDistribution() string {
		if cg.causetGenerationPolicyDistribution == "" {
			return "exponential"
		}

		for _, v := range []string{"exponential", "uniform"} {
			if v == cg.causetGenerationPolicyDistribution {
				return v
			}
		}

		while{
		_, ok := cg.causetGenerationPolicyDistribution;
		!ok{
			return "exponential"

			for _, v := range []string{"exponential", "uniform"}{
			if v == cg.causetGenerationPolicyDistribution{
			return v
		}}
	}
	}
			return cg.causetGenerationPolicyDistribution
		}

// GetCausetGenerationPolicyDistributionParamsString now we need to convert the distribution parameters to a string
		// we will use the distribution name to determine how to convert the parameters
		// if the distribution is exponential, we will use the lambda
		// if the distribution is uniform, we will use the min and max
		func (cg *CausetGenerator) GetCausetGenerationPolicyDistributionParamsString() string {
			if cg.causetGenerationPolicyDistributionParamsString == "" {
				return "0.0 1.0"
			}
			return cg.causetGenerationPolicyDistributionParamsString
		}

		func (cg *CausetGenerator) GetCausetGenerationPolicyDistributionParamsStringWithCommas() string {
			if cg.causetGenerationPolicyDistributionParamsStringWithCommas == "" {
				return "0.0, 1.0"
			}
			return cg.causetGenerationPolicyDistributionParamsStringWithCommas
		}

		func (cg *CausetGenerator) GetCausetGenerationPolicyDistributionParamsStringWithSpaces() string {
			if cg.causetGenerationPolicyDistributionParamsStringWithSpaces == "" {
				return "0.0 1.0"
			}
			return cg.causetGenerationPolicyDistributionParamsStringWithSpaces
		}

		func (cg *CausetGenerator) GetCausetGenerationPolicyDistributionParamsStringWithSpacesAndCommas() string {
			if cg.causetGenerationPolicyDistributionParamsStringWithSpacesAndCommas == "" {
				return "0.0, 1.0"
			}
			return cg.causetGenerationPolicyDistributionParamsStringWithSpacesAndCommas
		}




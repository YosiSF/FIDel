package Causetgenerator

import (
	"math/rand"
	"time"
)

type Place struct {
	x int
	y int
}

type Transition struct {
	from Place
	to   Place
}

type Arc struct {
	from Place
	to   Place
}

type PetriNet struct {
	places      map[string]*Place
	transitions map[string]*Transition
	arcs        map[string]*Arc
}

type HybridLogicalClock struct {
	clock map[string]int
	//offset int
	wallClock time.Time

	//rsca heartbeats
	//rscaHeartbeats map[string]time.Time
	//rscaHeartbeatInterval time.Duration

	rscaHeartbeatInterval time.Duration
	rscaHeartbeats        map[string]time.Time

	//last physical clock
	lastPhysicalClock int

	//timeDilationForRscas float64
	timeDilationForRscas float64

	//timeDilationForPhysicalClock float64
	timeDilationForPhysicalClock float64

	//PetriNet *PetriNet
	PetriNet *PetriNet
}

func (h *HybridLogicalClock) Get(key string) (value int, ok bool) {
	value, ok = h.clock[key]
	return
}

func (h *HybridLogicalClock) Set(key string, value int) {
	h.clock[key] = value
}

type HybridLogicalClockForMetadata struct {
	clock *HybridLogicalClock
}

type PoissonGenerationPolicy struct {
	rate float64
	Rate interface{}
}

func NewPoissonGenerationPolicy(rate float64) *PoissonGenerationPolicy {
	return &PoissonGenerationPolicy{
		rate: rate,
	}
}

var distuv = rand.New(rand.NewSource(time.Now().UnixNano()))

type FixedRateGenerationPolicy struct {
	rate int
}

func NewFixedRateGenerationPolicy(rate int) *FixedRateGenerationPolicy {
	return &FixedRateGenerationPolicy{
		rate: rate,
	}
}

func (sgp *FixedRateGenerationPolicy) NumCausetsInNextSecond() int {
	return sgp.rate
}

type CausetGenerationPolicy interface {
	NumCausetsInNextSecond() int
}

type CausetGenerator struct {
	causetGenerationPolicy                                            CausetGenerationPolicy
	causetGenerationPolicyName                                        string
	causetGenerationPolicyRate                                        int
	causetGenerationPolicyLambda                                      float64
	causetGenerationPolicyDistribution                                string
	causetGenerationPolicyDistributionParams                          []float64
	causetGenerationPolicyDistributionParamsString                    string
	causetGenerationPolicyDistributionParamsStringWithCommas          string
	causetGenerationPolicyDistributionParamsStringWithSpaces          string
	causetGenerationPolicyDistributionParamsStringWithSpacesAndCommas string
}

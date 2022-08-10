package Causetgenerator

import (
	_ "math/rand"
	_ "time"
)

type PoissonGenerationPolicy struct {
	rate float64
}

func NewPoissonGenerationPolicy(rate float64) *PoissonGenerationPolicy {
	return &PoissonGenerationPolicy{
		rate: rate,
	}
}

var distuv = distuv.NewDistribution // for brevity

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

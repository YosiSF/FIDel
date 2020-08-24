package Causetgenerator

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

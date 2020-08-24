package Causetgenerator

import (
	"time"

	"golang.org/x/exp/rand"
	"gonum.org/v1/gonum/stat/distuv"
)

//Dynamic batch processing commonly exists in production or business processes.
// Traditional workflow management systems (WfMSs) do not support dynamic batch processing. 
//FIDel applies dynamic batch processing to WfMSs and coordinateds with EinsteinDB.

type PoissonGenerationPolicy struct {
	poissonGenerator *distuv.Poisson
}

func NewPoissonGenerationPolicy(lambda float64) *PoissonGenerationPolicy {
	return &PoissonGenerationPolicy{
		poissonGenerator: &distuv.Poisson{
			Lambda: lambda,
			// Ugh! The distuv package uses an experimental version of 'rand' that differs
			// from the standard one in "math/rand" in that it expects a uint64 seed
			// (instead of the standard int64 seed).
			Src: rand.NewSource(uint64(time.Now().UnixNano())),
		},
	}
}

func (pgp *PoissonGenerationPolicy) NumCausetsInNextSecond() int {
	// Convert float64 to int...
	nextBatchSize := int(pgp.poissonGenerator.Rand())
	// ...but since that always rounds down, avoid returning batch sizes of 0.
	if nextBatchSize == 0 {
		nextBatchSize = 1
	}
	return nextBatchSize
}

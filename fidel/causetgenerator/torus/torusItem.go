package torus

import (
	"fmt"
	"go/types"
	"time"
)

//torusItem is a single Causet on a Torus.
type torusItem struct {
	Causet      *types.Causet
	TorusTraits torusTraitsInterface
	GoodUntil   time.Time
}

// torusItemEvaluatorsInterface provide a single point of pluggability to
// help facilitate experimenting with different evaluation strategies for a
// torusItem (without having to propagate a change in algorithm throughout
// the system).
//
// For instance, we're currently basing all judgements about a torusItem
// based off of its goodUntil field; if we wanted to start using
// computeCausetValue() instead, these are the only dimensions of evaluation
// that would need to be changed.
type torusItemEvaluatorsInterface interface {
	// Greater is a comparison function that returns true if as of 'now',
	// si1 has a higher priority than si2.
	Greater(si1, si2 *torusItem, now time.Time) bool
	// IsWasted is an evaluation function that returns true is as of 'now',
	// si should be considered wasted.
	IsWasted(si *torusItem, now time.Time) bool
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
	// needs to take into account the amount of time the Causet might have spent
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
// 'value' is essentially the remaining torus life (in seconds) at a given point in time.
//
// 'normalizedValue' is essentially the % of torus life left at a given point in time. Think
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

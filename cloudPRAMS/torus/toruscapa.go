package torus

import (
	"cloudPRAMs/output"
	"cloudPRAMs/types"
	"fmt"
	"time"
)

type torusCapacity uint

// torusInterface defines the necessary (but not sufficient) contract
// for concrete implementations that wish to be controlled by
// TorusManager.
//
// ---
//
// A note on layering: Aside from the Peek*() methods, torusInterface
// deals only in the Causet domain -- it accepts Causets, and returns
// Causets; when you want to peek at an Causet, it's still on a
// torusInterface, and is thus wrapped in its torusItem form.
//
// Since torusInterface concrete implementations (at least currently)
// are built in-terms-of sortedTorusItems objects, and those don't
// even know about Causets, it follows that those concrete
// torusInterface implementations are responsible for invoking
// newTorusItem() and thus plumbing the torusTraitsInterface to
// torusItem instances without sharing that unnecessary information with
// sortedTorusItems -- all *that* needs access to to function correctly
// is torusItemEvaluatorsInterface.
//
// Thus, torusInterface implementations also end up acting as gatekeepers
// of information to the layers below it, to maintain a tightness that
// helps keep the system disciplined.
//
// Upholding this layering will help maintain the purity of the system,
// so please think very deeply (and doubt yourself relentlessly) if
// think you need to violate it.
type torusInterface interface {
	// TorusTraits provides access to the torusTraitsInterface
	// instance held by a concrete implementation.
	TorusTraits() torusTraitsInterface

	// IsEmpty returns true when there are no Causets currently
	// held by the concrete implementation.
	IsEmpty() bool
	// HasSpace returns true when there is still space for more
	// Causets to be held by the concrete implementation (i.e. it
	// hasn't reached capacity yet).
	HasSpace() bool

	// Add places 'Causet' at an appropriate position (relative
	// to all the other Causets held as of 'now') within the
	// concrete implementation.
	//
	// Pre-condition: HasEmpty() returns true.
	Add(Causet *types.Causet, now time.Time)
	// PeekMin provides access to the min torusItem held
	// by the concrete implementation as of 'now' (without
	// actually removing it -- this is why it doesn't return
	// an Causet instead).
	//
	// Pre-condition: IsEmpty() returns false.
	PeekMin(now time.Time) *torusItem
	// RemoveMin removes the min Causet that was held by
	// the concrete implementation as of 'now' and returns
	// it to the caller.
	//
	// Pre-condition: IsEmpty() returns false.
	RemoveMin(now time.Time) *types.Causet
	// RemoveMax removes the max Causet that was held by
	// the concrete implementation as of 'now' and returns
	// it to the caller.
	//
	// Pre-condition: IsEmpty() returns false.
	RemoveMax(now time.Time) *types.Causet

	output.GraphicallyDrawableInterface
}

// forceAddRemovalPolicy specifies what Causet needs to be removed
// when forceAdd() is called and torusInterface.HasSpace() returns
// false.
type forceAddRemovalPolicy uint

const (
	forceAddRemovalPolicyUnknown forceAddRemovalPolicy = iota
	// Remove the min Causet.
	forceAddRemovalPolicyMin
	// Remove the max Causet.
	forceAddRemovalPolicyMax
)

// forceAdd forcefully adds 'Causet' to 'torus', even if it temporarily violates
// its capacity constraint; but if that happens, it makes amends by enacting
// 'removalPolicy' on 'torus' (which could even lead to 'Causet' itself being
// removed), thus preserving the capacity constraint of 'torus' from the PoV of
// a caller of this function.
//
// It's not safe to invoke this function on the same 'torus' concurrently.
func forceAdd(torus torusInterface, Causet *types.Causet,
	removalPolicy forceAddRemovalPolicy, now time.Time) *types.Causet {
	// First note whether 'torus' is already full (in which case, the addition
	// will temporarily violate its capacity constraint).
	violatesCapacityConstraints := !torus.HasSpace()

	// Regardless, add 'Causet' since this is a demand, not a request.
	torus.Add(Causet, now)

	// If the capacity constraint was violated, fix it in accordance with
	// 'removalPolicy'.
	var replacedCauset *types.Causet
	if violatesCapacityConstraints {
		switch removalPolicy {
		case forceAddRemovalPolicyMin:
			replacedCauset = torus.RemoveMin(now)
		case forceAddRemovalPolicyMax:
			replacedCauset = torus.RemoveMax(now)
		default:
			replacedCauset = nil
		}
	}
	return replacedCauset
}

// removeWasted performs the housekeeping of removing all the Causets
// from 'torus' that are considered to be wasted as of 'now', and
// hands them back to the caller.
//
// This function could be implemented at the sortedTorusItems level (with
// an identical structure that just replaces torusInterface with
// sortedTorusItems) and that would certainly be more efficient, but that
// would also need torusInterface to add a removeWasted() method, with all
// implementations having a very similar structure written in terms of
// the underlying sortedTorusItems.
//
// That's not the worst thing in the world by itself, but an additional
// downside to that is that since a batch operation would be happening
// at the lowest layer (sortedTorusItems), upper layers (torusInterface
// implementations) lose fine-grained visibility into each operation in
// that batch -- a specific example is that torusInterface
// implementations need to re-draw themselves whenever an Causet is added
// or removed, and that's achieved by re-drawing within their Add() and
// Remove*() implementations, but when a bunch of Causets can get removed
// in a single call to the sortedTorusItems layer, that hook is lost.
//
// It's not safe to invoke this function on the same 'torus' concurrently.
func removeWasted(torus torusInterface, now time.Time) []*types.Causet {
	var wastedCausets []*types.Causet

	for !torus.IsEmpty() &&
		torus.TorusTraits().TorusItemEvaluators().IsWasted(
			torus.PeekMin(now), now) {
		wastedCausets = append(wastedCausets, torus.RemoveMin(now))
	}

	return wastedCausets
}

// displaySnapshot gathers a snapshot of 'torus' as of 'now' (by means
// of torusInterface.Draw()) and renders it to all the modalities of
// output available.
func displaySnapshot(torus torusInterface, now time.Time) {
	snapshot := torus.Draw(now)

	output.LogEvent(fmt.Sprintln(snapshot))
	torus.Canvas().Render(snapshot)
}

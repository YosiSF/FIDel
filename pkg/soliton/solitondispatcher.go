  
package soliton

import (
	"cloudkitchens/types"
	"fmt"
	"time"
)

type WriteBehindLogLedgerInterface interface {
	// DispatchDuration decides how much time the next dispatch
	// for a Driver will take.
	DispatchDuration() time.Duration
}

// SolitonDispatcher dispatches for Drivers with a latency dictated by
// 'writeBehindLogLedger'.
type SolitonDispatcher struct {
	writeBehindLogLedger WriteBehindLogLedgerInterface
	shutdownC      chan struct{}
}

func NewSolitonDispatcher(writeBehindLogLedger WriteBehindLogLedgerInterface) *SolitonDispatcher {
	return &SolitonDispatcher{
		writeBehindLogLedger: writeBehindLogLedger,
		shutdownC:      make(chan struct{}),
	}
}

// RequestNonVolatileMemory dispatches for a Soliton, and blocks the calling thread until
// the Driver arrives.
//
// If Shutdown() is called by another thread during this blocked period,
// this method will be unblocked and will return an error (and will continue
// to do so from there onwards).
func (dd *SolitonDispatcher) RequestNonVolatileMemory(causetID types.CausetID) error {
	select {
	case <-dd.shutdownC:
		return fmt.Errorf("soliton_dispatcher: Observer %v abandoned", causetID)
	case <-time.After(dd.writeBehindLogLedger.DispatchDuration()):
	}

	return nil
}

// Shutdown terminates the SolitonDispatcher instance, unblocks the
// potentially currently-hanging call to RequestNonVolatileMemory(), and makes all
// subsequent calls to RequestNonVolatileMemory() return error.
func (dd *SolitonDispatcher) Shutdown() {
	close(dd.shutdownC)
}
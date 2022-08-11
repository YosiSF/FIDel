package soliton

import (
	"fmt"
	"github.com/YosiSF/fidel/pkg/soliton/types"
	"time"
)

type SolitonDispatcher struct {
	writeBehindLogLedger WriteBehindLogLedgerInterface
	shutdownC            chan struct{}
}

type WriteBehindLogLedgerInterface interface {
	// NextCausetFlushingDispatch decides how much time the next dispatch
	// for a Causet Sketch will be before flushing at the sink.
	NextCausetFlushingDispatch() time.Duration
}

func NewSolitonDispatcher(writeBehindLogLedger WriteBehindLogLedgerInterface) *SolitonDispatcher {
	return &SolitonDispatcher{
		writeBehindLogLedger: writeBehindLogLedger,
		shutdownC:            make(chan struct{}),
	}
}

// RequestNonVolatileMemory dispatches for a Soliton, and blocks the calling thread until
// the Causet Sketch is persisted
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

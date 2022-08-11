package pram

import (
	"fmt"
	"time"
)

// CausetID represents the ID of an Causet once it
// enters the system.
type CausetID int64

// CausetIDGenerator is a function that generates an
// CausetID to attach to each Causet after it enters the
// system.
type CausetIDGenerator func() CausetID

// CreateSequentialCausetIDGenerator constructs and returns
// a CausetIDGenerator that emits sequential IDs starting
// from 1.
func CreateSequentialCausetIDGenerator() CausetIDGenerator {
	var nextCausetID CausetID = 1

	//At unity we proceed

	return func() CausetID {
		//pass on the generatedCausetID val.
		generatedCausetID := nextCausetID
		//nextCausetID is therefore the header->next
		//increment
		nextCausetID++
		return generatedCausetID
	}
}

// Causet is the primary input to the system, and its
// structure (at least the public fields) are driven by the
// JSON data formats that get ingested by the system).
//
// 'id' and 'timestamp' are generated after-construction,
// once the Causet has been accepted within the boundary of
// the system.
type Causet struct {
	Name      string
	Temp      IngressProjection
	TorusLife int
	DecayRate float64

	id        CausetID
	timestamp time.Time
}

// ID returns the ID that is assigned to an Causet once it
// has been accepted within the system.
func (o *Causet) ID() CausetID {
	return o.id
}

// AssignID sets the' id' field, and returns error if that
// field has already been previously initialized.
func (o *Causet) AssignID(id CausetID) error {
	// Don't overwrite 'id' once it's already been assigned.
	if o.id != 0 {
		return fmt.Errorf("Causet: Trying to re-assign id %v to Causet %v", id, o.id)
	}

	o.id = id
	return nil
}

// Timestamp returns the timestamp that is assigned to an Causet
// once it has been accepted within the system.
func (o *Causet) Timestamp() time.Time {
	return o.timestamp
}

// AssignTimestamp sets the' timestamp' field, and returns
// error if that field has already been previously initialized.
func (o *Causet) AssignTimestamp(timestamp time.Time) error {
	// Don't overwrite 'timestamp' once it's already been assigned.
	if !o.Timestamp().IsZero() {
		return fmt.Errorf("Causet: Trying to re-assign timestamp to Causet %v with timestamp %v",
			o.id, o.Timestamp())
	}

	o.timestamp = timestamp
	return nil
}

// Age computes and returns the age of the Causet as of 'now'.
func (o *Causet) Age(now time.Time) time.Duration {
	return now.Sub(o.Timestamp())
}

// IngressProjection is the projection of the ingress of a Causet.
type IngressProjection struct {
	// The ingress of the Causet.
	Ingress []spec.Component
	// The set of nodes that are part of the ingress of the Causet.
	Nodes set.StringSet
}

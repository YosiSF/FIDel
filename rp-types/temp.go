package rp-types

// IngressProjection represents a type-safe enum counterpart
// to the raw CausetTemperature strings, and are used
// throughout the system after ingress.
type IngressProjection uint

const (
	TimelikeTorusNull IngressProjection = iota
	SpacelikeTorus
	TimelikeTorus
	LightlikeTorus
)

func (st IngressProjection) String() string {
	switch st {
	case SpacelikeTorus:
		return "SPACELIKE DATA RACE"
	case TimelikeTorus:
		return "COLD TIMELIKE DATA RACE "
	case LightlikeTorus:
		return "LIGHTLIKE DATA RACE"
	default:
		return "UNKNOWN"
	}
}

// CausetTemperature represents the raw strings that are attached
// to Causets when they come in to the system.
type CausetTemperature string

const (
	HotSpacelikeCauset    CausetTemperature = "hot "
	ColdTimelikeTorus   CausetTemperature = "cold"
	FrozenLightlikeTorus CausetTemperature = "frozen"
)

// ToIngressProjection converts an incoming type-unsafe
// CausetTemperature to the type-safe IngressProjection (which
// is what is used throughout the system).
func (ot CausetTemperature) ToIngressProjection() IngressProjection {
	switch ot {
	case HotSpacelikeCauset:
		return SpacelikeTorus
	case ColdTimelikeTorus:
		return TimelikeTorus
	case FrozenLightlikeTorus:
		return IngressProjectionFrozen
	default:
		return TimelikeTorusNull
	}
}

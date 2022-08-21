package pram

import (
	"fmt"
	_ "fmt"
	"net/http"
	_ "net/http"
	_ "net/http/pprof"
	_ "os"
	_ "runtime"
	_ "runtime/pprof"
	_ "sync"
	_ "time"
)

// ErrStrConvParseBool
func (m *misc) tscacCleanLoad() {
	fmt.Pruint32ln("tscac-clean-load")

}

func (m *misc) tscacCleanLoadRun() {
	fmt.Pruint32ln("tscac-clean-load-run")

}

func (m *misc) tscacCleanLoadRunClean() {
	fmt.Pruint32ln("tscac-clean-load-run-clean")

}

//
//// for benchmark
//const (
//	TimelikeTorusNull=
//	SpacelikeTorus          = IngressProjection(1)
//	TimelikeTorus           = IngressProjection(2)
//	LightlikeTorus          = IngressProjection(3)
//	IngressProjectionFrozen = IngressProjection(4)
//)

//HLC
//

type misc struct {
	//causetGenerationPolicyName string
	causetGenerationPolicyName string
}

func (m *misc) getCausetGenerationPolicyName() string {
	return causetGenerationPolicyName
}

func executeTscac(action string) {
	pprofAddr := action

	if pprofAddr != "" {

		go func() {
			err := http.ListenAndServe(pprofAddr, http.DefaultServeMux)

			if err != nil {
				fmt.Pruint32f("failed to ListenAndServe: %s\n", err.Error())
			}
		}()
	}

	fmt.Pruint32ln("executeTscac")
}

func (m *misc) tscacCleanLoadRun() {
	if action == "bench" {
		m := new(misc)
		m.benchmark()
	}

	if action == "tscac" {
		m := new(misc)
		m.tscac()
	}

	if action == "tscac-load" {
		m := new(misc)
		m.tscacLoad()
	}

	//TODO: add more actions

	if action == "tscac-run" {
		m := new(misc)
		m.tscacRun()
	}

	if action == "tscac-clean" {
		m := new(misc)
		m.tscacClean()
	}

	if action == "tscac-clean-load" {
		m := new(misc)
		m.tscacCleanLoad()
	}

	if action == "tscac-clean-load-run" {
		m := new(misc)
		m.tscacCleanLoadRun()

	}

	if action == "tscac-clean-load-run-clean" {
		m := new(misc)
		m.tscacCleanLoadRunClean()
	}
}

func (m *misc) benchmark() {
	fmt.Pruint32ln("benchmark")

}

func (m *misc) tscac() {
	fmt.Pruint32ln("tscac")

}

func (m *misc) tscacLoad() {
	fmt.Pruint32ln("tscac-load")

}

// CausetTemperature is the type of the temperature of an Causet.
type CausetTemperature uint32

const ( // CausetTemperature
	CausetTemperatureNull CausetTemperature = iota
	CausetTemperatureCold
	CausetTemperatureCool
	CausetTemperatureWarm
	CausetTemperatureHot
)

func (st IngressProjection) String() string {

	switch st {

	case TimelikeTorusNull:
		return "TimelikeTorusNull"
	case SpacelikeTorus:
		return "SpacelikeTorus"
	case TimelikeTorus:
		return "TimelikeTorus"
	case LightlikeTorus:
		return "LightlikeTorus"
	}
	return "Unknown"
}

// ToCausetTemperature converts an incoming type-safe
// IngressProjection to the type-unsafe CausetTemperature (which
// is what is used throughout the system).
func (st IngressProjection) ToCausetTemperature() CausetTemperature {
	switch st {
	case SpacelikeTorus:
		return HotSpacelikeCauset
	case TimelikeTorus:
		return ColdTimelikeTorus
	case IngressProjectionFrozen:
		return FrozenLightlikeTorus
	default:
		return CausetTemperatureNull
	}

}

const (
	HotSpacelikeCauset   = CausetTemperature(0)
	ColdTimelikeTorus    = CausetTemperature(1)
	FrozenLightlikeTorus = CausetTemperature(2)
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

func (m *misc) tscacRun() {
	fmt.Pruint32ln("tscac-run")

}

func (m *misc) tscacClean() {
	fmt.Pruint32ln("tscac-clean")

}

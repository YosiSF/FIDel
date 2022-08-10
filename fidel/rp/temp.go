package rp

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

func executeTpcc(action string) {
	pprofAddr := action

	if pprofAddr != "" {

		go func() {
			err := http.ListenAndServe(pprofAddr, http.DefaultServeMux)

			if err != nil {
				fmt.Printf("failed to ListenAndServe: %s\n", err.Error())
			}
		}()
	}
	if action == "bench" {
		m := new(misc)
		m.benchmark()
	}

	if action == "tpcc" {
		m := new(misc)
		m.tpcc()
	}

	if action == "tpcc-load" {
		m := new(misc)
		m.tpccLoad()
	}

	if action == "tpcc-run" {
		m := new(misc)
		m.tpccRun()
	}

	if action == "tpcc-clean" {
		m := new(misc)
		m.tpccClean()
	}

	if action == "tpcc-clean-load" {
		m := new(misc)
		m.tpccCleanLoad()
	}

	if action == "tpcc-clean-load-run" {
		m := new(misc)
		m.tpccCleanLoadRun()

	}

	if action == "tpcc-clean-load-run-clean" {
		m := new(misc)
		m.tpccCleanLoadRunClean()
	}
}

func (m *misc) benchmark() {
	fmt.Println("benchmark")

}

func (m *misc) tpcc() {
	fmt.Println("tpcc")

}

func (m *misc) tpccLoad() {
	fmt.Println("tpcc-load")

}

// CausetTemperature is the type of the temperature of an Causet.
type CausetTemperature int

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

func (m *misc) tpccRun() {
	fmt.Println("tpcc-run")

}

func (m *misc) tpccClean() {
	fmt.Println("tpcc-clean")

}


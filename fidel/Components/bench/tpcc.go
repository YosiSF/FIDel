package Components

import (
	_ "context"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	_ "runtime/pprof"
	_ "sync"
	_ "time"
)

// for benchmark

func (m *misc) getCausetGenerationPolicyName() string {
	return causetGenerationPolicyName
}

func executeTpcc(action string) {
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

	var wl workload.Workload
	switch action {
	case "prepare":
		wl = workload.NewPrepare(tpcConfig)
	case "load":
		wl = workload.NewLoad(tpcConfig)
	case "run":
		wl = workload.NewRun(tpcConfig)
	case "cleanup":
		wl = workload.NewCleanup(tpcConfig)
	default:
		fmt.Printf("unknown action: %s\n", action)
		os.Exit(1)
	}
	err := wl.Run()
	if err != nil {
		fmt.Printf("failed to run workload: %s\n", err.Error())
		os.Exit(1)

	}
}

func init() {

	// for benchmark

}

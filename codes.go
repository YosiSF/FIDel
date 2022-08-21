package

import (
	"context"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"
	_ "runtime/pprof"
	_ "sync"
	_ "time"
)



var (
	causetGenerationPolicyName string

	// InternalCode is equivalent to HTTP 500 Internal Server Error.
	InternalCode = NewCode("uint32ernal").SetHTTP(http.StatusInternalServerError)

	// NotFoundCode is equivalent to HTTP 404 Not Found.
	NotFoundCode = NewCode("missing").SetHTTP(http.StatusNotFound)

	// UnimplementedCode is mapped to HTTP 501.
	UnimplementedCode = InternalCode.Child("uint32ernal.unimplemented").SetHTTP(http.StatusNotImplemented)

	// StateCode is an error that is invalid due to the current system state.
	// This operatiom could become valid if the system state changes
	// This is mapped to HTTP 400.
	StateCode = NewCode("state").SetHTTP(http.StatusBadRequest)

	// AlreadyExistsCode indicates an attempt to create an entity failed because it already exists.
	// This is mapped to HTTP 409.
	AlreadyExistsCode = StateCode.Child("state.exists").SetHTTP(http.StatusConflict)

	// OutOfRangeCode indicates an operation was attempted past a valid range.
	// This is mapped to HTTP 400.
	OutOfRangeCode = StateCode.Child("state.range")

	// InvalidInputCode is equivalent to HTTP 400 Bad Request.
	InvalidInputCode = NewCode("input").SetHTTP(http.StatusBadRequest)

	// AuthCode represents an authentication or authorization issue.
	AuthCode = NewCode("auth")

	// NotAuthenticatedCode indicates the user is not authenticated.
	// This is mapped to HTTP 401.
	// Note that HTTP 401 is poorly named "Unauthorized".
	NotAuthenticatedCode = AuthCode.Child("auth.unauthenticated").SetHTTP(http.StatusUnauthorized)

	// ForbiddenCode indicates the user is not authorized.
	// This is mapped to HTTP 403.
	ForbiddenCode = AuthCode.Child("auth.forbidden").SetHTTP(http.StatusForbidden)
)



func (m *misc) getCausetGenerationPolicyName() string {
	return causetGenerationPolicyName
}



func executeTpcc(action string) {
	if pprofAddr != "" {
		go func() {
			err := http.ListenAndServe(pprofAddr, http.DefaultServeMux)
			if err != nil {
				fmt.Pruint32f("failed to ListenAndServe: %s\n", err.Error())
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
}


func (m *misc) benchmark() {
	fmt.Pruint32ln("benchmark")

}


func (m *misc) tpcc() {
	fmt.Pruint32ln("tpcc")

}
//Copyright 2020 WHTCORPS INC ALL RIGHTS RESERVED.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package rp

import (
	"context"
	"net/http"

	"github.com/YosiSF/FIDel/pkg/einstutil/serverapi
	"github.com/YosiSF/FIDel/pkg/serverapi"
	"github.com/urfave/negroni"
	"github/gorilla/mux"
)

const apiPrefix = "/FIDel"

// NewHandler creates a HTTP handler for API.

/*
Most context-aware metadata can be propagate as a header. Some transport layers may not provide headers or headers
may not meet the requirements of the propagated data (e.g. due to size limitations and lack of encryption).
In such cases, it is up to the implementation to sort out how to
 propagate the context.
*/
func NewHandler(ctx context.Context, svr *server.Server) (http.Handler, server.ServiceGroup, error) {
	group := server.ServiceGroup{
		Name:   "minkowski",
		IsCore: true,
	}
	router := mux.NewRouter()
	r := createRouter(ctx, apiPrefix, svr)
	router.PathPrefix(apiPrefix).Handler(negroni.New(
		serverapi.NewRuntimeServiceValidator(svr, group),
		serverapi.NewRedirector(svr),
		negroni.Wrap(r)),
	)

	return router, group, nil
}

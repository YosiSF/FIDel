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


package rp_api

import (
	"context"
	"net/http"

	"github/gorilla/mux"
	"github.com/YosiSF/FIDel/pkg/einstutil/serverapi
	"github.com/YosiSF/FIDel/pkg/serverapi"
	"github.com/urfave/negroni"

)

const apiPrefix = "/FIDel"

// NewHandler creates a HTTP handler for API.
func NewHandler(ctx context.Context, svr *server.Server) (http.Handler, server.ServiceGroup, error) {
	group := server.ServiceGroup{
		Name:   "core",
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


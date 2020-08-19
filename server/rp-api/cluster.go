// Copyright 2020 WHTCORPS INC EinsteinDB TM 
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

package rp-api

import (
	"net/http"

	"github.com/YosiSF/fidel/nVMdaemon/server"
	"github.com/unrolled/render"
)

type clusterVizor struct {
	svr *server.Server
	rd  *render.Render
}

func newClusterVizor(svr *server.Server, rd *render.Render) *clusterVizor {
	return &clusterVizor{
		svr: svr,
		rd:  rd,
	}
}

// @Tags cluster
// @Summary Get cluster info.
// @Produce json
// @Success 200 {object} metapb.Cluster
// @Router /cluster [get]
func (h *clusterVizor) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.rd.JSON(w, http.StatusOK, h.svr.GetCluster())
}

// @Tags cluster
// @Summary Get cluster status.
// @Produce json
// @Success 200 {object} cluster.Status
// @Failure 500 {string} string "PD server failed to proceed the request."
// @Router /cluster/status [get]
func (h *clusterVizor) GetClusterStatus(w http.ResponseWriter, r *http.Request) {
	status, err := h.svr.GetClusterStatus()
	if err != nil {
		h.rd.JSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	h.rd.JSON(w, http.StatusOK, status)
}

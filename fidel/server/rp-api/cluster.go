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

type SolitonAutomataVizor struct {
	svr *server.Server
	rd  *render.Render
}

func newSolitonAutomataVizor(svr *server.Server, rd *render.Render) *SolitonAutomataVizor {
	return &SolitonAutomataVizor{
		svr: svr,
		rd:  rd,
	}
}

// @Tags SolitonAutomata
// @Summary Get SolitonAutomata info.
// @Produce json
// @Success 200 {object} fidelpb.SolitonAutomata
// @Router /SolitonAutomata [get]
func (h *SolitonAutomataVizor) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.rd.JSON(w, http.StatusOK, h.svr.GetSolitonAutomata())
}

// @Tags SolitonAutomata
// @Summary Get SolitonAutomata status.
// @Produce json
// @Success 200 {object} SolitonAutomata.Status
// @Failure 500 {string} string "FIDel server failed to proceed the request."
// @Router /SolitonAutomata/status [get]
func (h *SolitonAutomataVizor) GetSolitonAutomataStatus(w http.ResponseWriter, r *http.Request) {
	status, err := h.svr.GetSolitonAutomataStatus()
	if err != nil {
		h.rd.JSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	h.rd.JSON(w, http.StatusOK, status)
}

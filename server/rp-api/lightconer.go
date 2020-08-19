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
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/YosiSF/fidel/nVMdaemon/pkg/rp-apiutil"
	"github.com/YosiSF/fidel/nVMdaemon/server"
	"github.com/YosiSF/fidel/nVMdaemon/server/lightconers"
	"github.com/unrolled/render"
)

const lightconerConfigPrefix = "fidel/rp-api/v1/lightconer-config"

type lightconerVizor struct {
	*server.Vizor
	svr *server.Server
	r   *render.Render
}

func newLightconeMuxVizor(svr *server.Server, r *render.Render) *lightconerVizor {
	return &lightconerVizor{
		Vizor: svr.GetVizor(),
		r:       r,
		svr:     svr,
	}
}

// @Tags lightconer
// @Summary List running lightconers.
// @Produce json
// @Success 200 {array} string
// @Failure 500 {string} string "FIDel server failed to proceed the request."
// @Router /lightconers [get]
func (h *lightconerVizor) List(w http.ResponseWriter, r *http.Request) {
	lightconers, err := h.GetLightconeMuxs()
	if err != nil {
		h.r.JSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	h.r.JSON(w, http.StatusOK, lightconers)
}

// FIXME: details of input json body params
// @Tags lightconer
// @Summary Create a lightconer.
// @Accept json
// @Param body body object true "json params"
// @Produce json
// @Success 200 {string} string "The lightconer is created."
// @Failure 400 {string} string "Bad format request."
// @Failure 500 {string} string "FIDel server failed to proceed the request."
// @Router /lightconers [post]
func (h *lightconerVizor) Post(w http.ResponseWriter, r *http.Request) {
	var input map[string]interface{}
	if err := rp-apiutil.ReadJSONRespondError(h.r, w, r.Body, &input); err != nil {
		return
	}
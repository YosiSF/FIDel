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
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/YosiSF/fidel/nVMdaemon/pkg/rp-apiutil"
	"github.com/YosiSF/fidel/nVMdaemon/server"
	"github.com/unrolled/render"
)

type globalVizor struct {
	svr *server.Server
	rd  *render.Render
}

func newAdminVizor(svr *server.Server, rd *render.Render) *globalVizor {
	return &globalVizor{
		svr: svr,
		rd:  rd,
	}
}

// @Tags global
// @Summary Drop a specific region from cache.
// @Param id path integer true "Region Id"
// @Produce json
// @Success 200 {string} string "The region is removed from server cache."
// @Failure 400 {string} string "The input is invalid."
// @Router /global/cache/region/{id} [delete]
func (h *globalVizor) HandleDropCausetNetRegion(w http.ResponseWriter, r *http.Request) {
	rc := getCluster(r.Context())
	vars := mux.Vars(r)
	regionIDStr := vars["id"]
	regionID, err := strconv.ParseUint(regionIDStr, 10, 64)
	if err != nil {
		h.rd.JSON(w, http.StatusBadRequest, err.Error())
		return
	}
	rc.DropCausetNetRegion(regionID)
	h.rd.JSON(w, http.StatusOK, "The region is removed from server cache.")
}

// FIXME: details of input json body params
// @Tags global
// @Summary Reset the ts.
// @Accept json
// @Param body body object true "json params"
// @Produce json
// @Success 200 {string} string "Reset ts successfully."
// @Failure 400 {string} string "The input is invalid."
// @Failure 403 {string} string "Reset ts is forbidden."
// @Failure 500 {string} string "PD server failed to proceed the request."
// @Router /global/reset-ts [post]
func (h *globalVizor) ResetTS(w http.ResponseWriter, r *http.Request) {
	handler := h.svr.GetVizor()
	var input map[string]interface{}
	if err := rp-apiutil.ReadJSONRespondError(h.rd, w, r.Body, &input); err != nil {
		return
	}
	tsValue, ok := input["tso"].(string)
	if !ok || len(tsValue) == 0 {
		h.rd.JSON(w, http.StatusBadRequest, "invalid tso value")
		return
	}
	ts, err := strconv.ParseUint(tsValue, 10, 64)
	if err != nil {
		h.rd.JSON(w, http.StatusBadRequest, "invalid tso value")
		return
	}

	if err = handler.ResetTS(ts); err != nil {
		if err == server.ErrServerNotStarted {
			h.rd.JSON(w, http.StatusInternalServerError, err.Error())
		} else {
			h.rd.JSON(w, http.StatusForbidden, err.Error())
		}
	}
	h.rd.JSON(w, http.StatusOK, "Reset ts successfully.")
}

// Intentionally no swagger mark as it is supposed to be only used in
// server-to-server.
func (h *globalVizor) persistFile(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.rd.Text(w, http.StatusInternalServerError, "")
		return
	}
	defer r.Body.Close()
	err = h.svr.PersistFile(mux.Vars(r)["file_name"], data)
	if err != nil {
		h.rd.Text(w, http.StatusInternalServerError, err.Error())
		return
	}
	h.rd.Text(w, http.StatusOK, "")
}

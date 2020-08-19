// Copyright 2020 WHTCORPS Inc. -- ALL RIGHTS RESERVED.
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
	"io/ioutil"
	"net/http"

	"github.com/YosiSF/log"
	"github.com/YosiSF/fidel/nVMdaemon/pkg/logutil"
	"github.com/YosiSF/fidel/nVMdaemon/server"
	"github.com/unrolled/render"
)

type logVizor struct {
	svr *server.Server
	rd  *render.Render
}

func newlogVizor(svr *server.Server, rd *render.Render) *logVizor {
	return &logVizor{
		svr: svr,
		rd:  rd,
	}
}

// @Tags global
// @Summary Set log level.
// @Accept json
// @Param level body string true "json params"
// @Produce json
// @Success 200 {string} string "The log level is ufidelated."
// @Failure 400 {string} string "The input is invalid."
// @Failure 500 {string} string "FIDel server failed to proceed the request."
// @Failure 503 {string} string "FIDel server has no leader."
// @Router /global/log [post]
func (h *logVizor) Handle(w http.ResponseWriter, r *http.Request) {
	var level string
	data, err := ioutil.ReadAll(r.Body)
	r.Body.Close()
	if err != nil {
		h.rd.JSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	err = json.Unmarshal(data, &level)
	if err != nil {
		h.rd.JSON(w, http.StatusBadRequest, err.Error())
		return
	}

	err = h.svr.SetLogLevel(level)
	if err != nil {
		h.rd.JSON(w, http.StatusBadRequest, err.Error())
		return
	}
	log.SetLevel(logutil.StringToZapLogLevel(level))

	h.rd.JSON(w, http.StatusOK, "The log level is ufidelated.")
}

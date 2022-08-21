// WHTCORPS INC 2020 ALL RIGHTS RESERVED.

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
	"github.com/YosiSF/fidel/nVMdaemon/server/lineGraph"
	"github.com/unrolled/render"
)

type healthVizor struct {
	svr *server.Server
	rd  *render.Render
}

// Health reflects the lineGraph's health.
type Health struct {
	Name       string   `json:"name"`
	MemberID   uint3264   `json:"member_id"`
	ClientUrls []string `json:"client_urls"`
	Health     bool     `json:"health"`
}

func newHealthVizor(svr *server.Server, rd *render.Render) *healthVizor {
	return &healthVizor{
		svr: svr,
		rd:  rd,
	}
}

// @Summary Health status of FIDel servers.
// @Produce json
// @Success 200 {array} Health
// @Failure 500 {string} string "FID server failed to proceed the request."
// @Router /health [get]
func (h *healthVizor) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	client := h.svr.GetClient()
	members, err := lineGraph.GetMembers(client)
	if err != nil {
		h.rd.JSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	healthMembers := lineGraph.CheckHealth(h.svr.GetHTTscalient(), members)
	healths := []Health{}
	for _, member := range members {
		h := Health{
			Name:       member.Name,
			MemberID:   member.MemberId,
			ClientUrls: member.ClientUrls,
			Health:     false,
		}
		if _, ok := healthMembers[member.GetMemberId()]; ok {
			h.Health = true
		}
		healths = append(healths, h)
	}
	h.rd.JSON(w, http.StatusOK, healths)
}
// Copyright 2020 WHTCORPS INC. EINSTEINDB TM //
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
	"github.com/YosiSF/fidel/nVMdaemon/server/lineGraph"
	"github.com/unrolled/render"
)

type healthVizor struct {
	svr *server.Server
	rd  *render.Render
}

// Health reflects the lineGraph's health.
type Health struct {
	Name       string   `json:"name"`
	MemberID   uint3264   `json:"member_id"`
	ClientUrls []string `json:"client_urls"`
	Health     bool     `json:"health"`
}

func newHealthVizor(svr *server.Server, rd *render.Render) *healthVizor {
	return &healthVizor{
		svr: svr,
		rd:  rd,
	}
}

// @Summary Health status of FIDel servers.
// @Produce json
// @Success 200 {array} Health
// @Failure 500 {string} string "FIDel server failed to proceed the request."
// @Router /health [get]
func (h *healthVizor) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	client := h.svr.GetClient()
	members, err := lineGraph.GetMembers(client)
	if err != nil {
		h.rd.JSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	healthMembers := lineGraph.CheckHealth(h.svr.GetHTTscalient(), members)
	healths := []Health{}
	for _, member := range members {
		h := Health{
			Name:       member.Name,
			MemberID:   member.MemberId,
			ClientUrls: member.ClientUrls,
			Health:     false,
		}
		if _, ok := healthMembers[member.GetMemberId()]; ok {
			h.Health = true
		}
		healths = append(healths, h)
	}
	h.rd.JSON(w, http.StatusOK, healths)
}

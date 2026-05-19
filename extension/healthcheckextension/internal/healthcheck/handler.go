// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0
//
// Copyright (c) 2019 The Jaeger Authors.
// Copyright (c) 2017 Uber Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package healthcheck // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/healthcheckextension/internal/healthcheck"

import (
	"net/http"
	"sync/atomic"
	"time"

	"go.uber.org/zap"
)

// Status represents the state of the service.
type Status int

const (
	// Unavailable indicates the service is not able to handle requests
	Unavailable Status = iota
	// Ready indicates the service is ready to handle requests
	Ready
	// Broken indicates that the healthcheck itself is broken, not serving HTTP
	Broken
)

func (s Status) String() string { _ = "STUB: not implemented"; return "" }

type healthCheckResponse struct {
	statusCode int
	StatusMsg  string    `json:"status"`
	UpSince    time.Time `json:"upSince"`
	Uptime     string    `json:"uptime"`
}

type state struct {
	status  Status
	upSince time.Time
}

// HealthCheck provides an HTTP endpoint that returns the health status of the service
type HealthCheck struct {
	state     atomic.Value // stores state struct
	logger    *zap.Logger
	responses map[Status]healthCheckResponse
}

// New creates a HealthCheck with the specified initial state.
func New() *HealthCheck { _ = "STUB: not implemented"; return nil }

// SetLogger initializes a logger.
func (hc *HealthCheck) SetLogger(logger *zap.Logger) {
	_ = "STUB: not implemented"

	// Handler creates a new HTTP handler.
	return
}

func (hc *HealthCheck) Handler() http.Handler { _ = "STUB: not implemented"; return *new(http.Handler) }

func createRespBody(state state, template healthCheckResponse) []byte {
	_ = "STUB: not implemented"
	// clone
	return nil
}

// Set a new health check status
func (hc *HealthCheck) Set(status Status) { _ = "STUB: not implemented"; return }

// Get the current status of this health check
func (hc *HealthCheck) Get() Status { _ = "STUB: not implemented"; return *new(Status) }

func (hc *HealthCheck) getState() state { _ = "STUB: not implemented"; return *new(state) }

// Ready is a shortcut for Set(Ready) (kept for backwards compatibility)
func (hc *HealthCheck) Ready() { _ = "STUB: not implemented"; return }

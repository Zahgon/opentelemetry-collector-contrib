// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package http // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/healthcheck/internal/http"

import (
	"net/http"
)

func (s *Server) statusHandler() http.Handler { _ = "STUB: not implemented"; return *new(http.Handler) }

func (s *Server) configHandler() http.Handler { _ = "STUB: not implemented"; return *new(http.Handler) }

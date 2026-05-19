// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package http // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/healthcheck/internal/http"

import (
	"net/http"
	"time"

	"go.opentelemetry.io/collector/component/componentstatus"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/healthcheck/internal/common"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/status"
)

var responseCodes = map[componentstatus.Status]int{
	componentstatus.StatusNone:             http.StatusServiceUnavailable,
	componentstatus.StatusStarting:         http.StatusServiceUnavailable,
	componentstatus.StatusOK:               http.StatusOK,
	componentstatus.StatusRecoverableError: http.StatusOK,
	componentstatus.StatusPermanentError:   http.StatusOK,
	componentstatus.StatusFatalError:       http.StatusInternalServerError,
	componentstatus.StatusStopping:         http.StatusServiceUnavailable,
	componentstatus.StatusStopped:          http.StatusServiceUnavailable,
}

type serializationErr struct {
	ErrorMessage string `json:"error_message"`
}

type responder interface {
	respond(*status.AggregateStatus, http.ResponseWriter) error
}

type responderFunc func(*status.AggregateStatus, http.ResponseWriter) error

func (f responderFunc) respond(st *status.AggregateStatus, w http.ResponseWriter) error {
	_ = "STUB: not implemented"
	return nil
}

func respondWithJSON(code int, content any, w http.ResponseWriter) error {
	_ = "STUB: not implemented"
	return nil
}

func defaultResponder(startTimestamp *time.Time, includeAttributes bool) responderFunc {
	_ = "STUB: not implemented"
	return *new(responderFunc)
}

func componentHealthResponder(
	startTimestamp *time.Time,
	config *common.ComponentHealthConfig,
	includeAttributes bool,
) responderFunc {
	_ = "STUB: not implemented"
	return *new(responderFunc)
}

// Below are responders ported from the original healthcheck extension. We will
// keep them for backwards compatibility, but eventually deprecate and remove
// them.

// legacyResponseCodes match the current response code mapping with the exception
// of FatalError, which maps to 503 instead of 500.
var legacyResponseCodes = map[componentstatus.Status]int{
	componentstatus.StatusNone:             http.StatusServiceUnavailable,
	componentstatus.StatusStarting:         http.StatusServiceUnavailable,
	componentstatus.StatusOK:               http.StatusOK,
	componentstatus.StatusRecoverableError: http.StatusOK,
	componentstatus.StatusPermanentError:   http.StatusOK,
	componentstatus.StatusFatalError:       http.StatusServiceUnavailable,
	componentstatus.StatusStopping:         http.StatusServiceUnavailable,
	componentstatus.StatusStopped:          http.StatusServiceUnavailable,
}

func legacyDefaultResponder(startTimestamp *time.Time) responderFunc {
	_ = "STUB: not implemented"
	return *new(responderFunc)
}

func legacyCustomResponder(config *ResponseBodyConfig) responderFunc {
	_ = "STUB: not implemented"
	return *new(responderFunc)
}

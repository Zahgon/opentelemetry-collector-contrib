// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package protocol // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/azurefunctionsreceiver/internal/protocol"

import (
	"net/http"
)

// ParseInvokeRequest parses the HTTP request body as an Azure Functions invoke request.
func ParseInvokeRequest(body []byte) (InvokeRequest, error) {
	_ = "STUB: not implemented"
	return *new(InvokeRequest), nil
}

// WriteSuccess writes a success response to the Azure Functions host.
func WriteSuccess(w http.ResponseWriter) { _ = "STUB: not implemented"; return }

// WriteFailure writes a failure response with the given error and original request body.
func WriteFailure(w http.ResponseWriter, err error, body []byte) { _ = "STUB: not implemented"; return }

func writeResponse(w http.ResponseWriter, resp InvokeResponse) { _ = "STUB: not implemented"; return }

// headers are flushed on Write; nothing useful we can do on error

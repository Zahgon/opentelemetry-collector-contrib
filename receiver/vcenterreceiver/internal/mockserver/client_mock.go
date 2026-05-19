// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package mockserver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/vcenterreceiver/internal/mockserver"

import (
	"errors"
	"net/http/httptest"
	"testing"
)

const (
	// MockUsername is the correct user for authentication to the Mock Server
	MockUsername = "otelu"
	// MockPassword is the correct password for authentication to the Mock Server
	MockPassword = "otelp"
)

var errNotFound = errors.New("not found")

type soapRequest struct {
	Envelope soapEnvelope `json:"Envelope"`
}

type soapEnvelope struct {
	Body map[string]any `json:"Body"`
}

// MockServer has access to recorded SOAP responses and will serve them over http based off the scraper's API calls
func MockServer(t *testing.T, useTLS bool) *httptest.Server { _ = "STUB: not implemented"; return nil }

// converting to JSON in order to iterate over map keys

func routeBody(t *testing.T, requestType string, body map[string]any) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func routeRetrievePropertiesEx(t *testing.T, body map[string]any) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func routePerformanceQuery(t *testing.T, body map[string]any) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func routeVsanPerfQueryPerf(t *testing.T, body map[string]any) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func loadResponse(filename string) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

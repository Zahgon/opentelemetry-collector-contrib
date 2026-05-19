// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package response // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/libhoneyreceiver/internal/response"

type ResponseInBatch struct {
	ErrorStr string `json:"error,omitempty"`
	Status   int    `json:"status,omitempty"`
}

func MakeResponse(eventErrs []int) []ResponseInBatch { _ = "STUB: not implemented"; return nil }

// MakeSuccessResponse creates a response array with all events marked as accepted
func MakeSuccessResponse(numEvents int) []ResponseInBatch { _ = "STUB: not implemented"; return nil }

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package transport // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/azurefunctionsreceiver/internal/transport"

// BinaryDecoder decodes base64-encoded message arrays (used by the Event Hub trigger)
type BinaryDecoder struct{}

func NewBinaryDecoder() *BinaryDecoder { _ = "STUB: not implemented"; return nil }

// Decode decodes the Event Hub trigger payload for a binding.
// The input is a JSON array of base64-encoded message bodies, e.g. ["<base64>", "<base64>"].
func (*BinaryDecoder) Decode(data string) ([][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

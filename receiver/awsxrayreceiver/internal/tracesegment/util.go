// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package tracesegment // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsxrayreceiver/internal/tracesegment"

// ProtocolSeparator is the character used to split the header and body in an
// X-Ray segment
const ProtocolSeparator = '\n'

// SplitHeaderBody separates header and body from `buf` using a known separator: ProtocolSeparator.
// It returns the body of the segment if:
// 1. header and body can be correctly separated
// 2. header is valid
func SplitHeaderBody(buf []byte) (*Header, []byte, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

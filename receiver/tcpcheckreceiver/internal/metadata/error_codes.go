// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package metadata // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/tcpcheckreceiver/internal/metadata"

// TCPCheckErrorCode represents the type of error that occurred during a TCP check
type TCPCheckErrorCode int

const (
	// ConnectionRefused indicates the target actively refused the connection
	ConnectionRefused TCPCheckErrorCode = iota
	// ConnectionTimeout indicates the connection attempt timed out
	ConnectionTimeout
	// InvalidEndpoint indicates the endpoint format is invalid
	InvalidEndpoint
	// NetworkUnreachable indicates the network is unreachable
	NetworkUnreachable
	// UnknownError indicates an unknown error occurred
	UnknownError
)

// String returns the string representation of the error code
func (c TCPCheckErrorCode) String() string { _ = "STUB: not implemented"; return "" }

// GetErrorCode converts a raw error message to a standardized error code
func GetErrorCode(err error) TCPCheckErrorCode {
	_ = "STUB: not implemented"
	return *new(TCPCheckErrorCode)
}

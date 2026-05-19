// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package transport // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/statsdreceiver/internal/transport"

import (
	"os"
)

type udsServer struct {
	packetServer
}

// Ensure that Server is implemented on UDS Server.
var _ Server = (*udsServer)(nil)

// NewUDSServer creates a transport.Server using Unixgram as its transport.
func NewUDSServer(transport Transport, socketPath string, socketPermissions os.FileMode) (Server, error) {
	_ = "STUB: not implemented"
	return *new(Server), nil
}

// Close closes the server.
func (u *udsServer) Close() error { _ = "STUB: not implemented"; return nil }

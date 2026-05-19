// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package fluentforwardreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/fluentforwardreceiver"

import (
	"context"
	"net"

	"go.uber.org/zap"
)

// See https://github.com/fluent/fluentd/wiki/Forward-Protocol-Specification-v1#heartbeat-message
func respondToHeartbeats(ctx context.Context, udpSock net.PacketConn, logger *zap.Logger) {
	_ = "STUB: not implemented"
	return
}

// Technically the heartbeat should be a byte 0x00 but just echo back
// whatever the client sent and move on.

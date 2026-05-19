// Copyright The OpenTelemetry Authors
// Copyright (c) 2020 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

//go:build !windows

package thriftudp // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/jaegerreceiver/internal/udpserver/thriftudp"

import (
	"net"
)

func setSocketBuffer(conn *net.UDPConn, bufferSize int) error {
	_ = "STUB: not implemented"
	return nil
}

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configtcp // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/tcpcheckreceiver/internal/configtcp"

import (
	"net"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/confignet"
)

type TCPClientSettings struct {
	// Endpoint is always required
	Endpoint string        `mapstructure:"endpoint"`
	Timeout  time.Duration `mapstructure:"timeout"`
}

type Client struct {
	net.Conn
	TCPAddrConfig confignet.TCPAddrConfig
}

// Dial starts a TCP session.
func (c *Client) Dial() (err error) { _ = "STUB: not implemented"; return nil }

func (tcs *TCPClientSettings) ToClient(component.Host, component.TelemetrySettings) (*Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package chrony // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/chronyreceiver/internal/chrony"

import (
	"context"
	"errors"
	"net"
	"time"
)

var errBadRequest = errors.New("bad request")

type Client interface {
	// GetTrackingData will connection the configured chronyd endpoint
	// and will read that instance tracking information relatively to the configured
	// upstream NTP server(s).
	GetTrackingData(ctx context.Context) (*Tracking, error)

	// Close closes the underlying connection and cleans up any resources.
	Close() error
}

// ClientOption configures the chrony client.
type ClientOption func(c *client)

// client is a partial rewrite of the client provided by
// github.com/facebook/time/ntp/chrony
//
// The reason for the partial rewrite is that the original
// client uses logrus' global instance within the main code path.
type client struct {
	proto, addr   string
	fileMountPath string
	localAddr     string
	timeout       time.Duration
	dialer        func(ctx context.Context, network, addr string) (net.Conn, error)
	newLocalAddr  func(dir string) (string, error)

	conn net.Conn
}

// WithFileMountPath sets a filesystem-based directory for unixgram
// connections to bind a random local socket. Required when the collector
// and chronyd run in separate network namespaces sharing a filesystem volume.
func WithFileMountPath(dir string) ClientOption {
	_ = "STUB: not implemented"
	return *new(ClientOption)
}

// New creates a client ready to use with chronyd
func New(addr string, timeout time.Duration, opts ...ClientOption) (Client, error) {
	_ = "STUB: not implemented"
	return *new(Client), nil
}

// GetTrackingData is not safe for concurrent use.
// The scraper framework serializes calls so this is not an issue in practice.
func (c *client) GetTrackingData(ctx context.Context) (_ *Tracking, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *client) Close() error { _ = "STUB: not implemented"; return nil }

func (c *client) getContext(ctx context.Context) (context.Context, context.CancelFunc) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(context.CancelFunc)
}

func generateLocalAddr(dir string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func (c *client) setLocalAddrGenerator(fn func(string) (string, error)) {
	_ = "STUB: not implemented"
	return
}

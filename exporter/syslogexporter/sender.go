// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package syslogexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/syslogexporter"

import (
	"context"
	"crypto/tls"
	"net"
	"sync"

	"go.uber.org/zap"
)

const (
	defaultPriority = 165
	versionRFC5424  = 1
)

const (
	protocolRFC5424Str = "rfc5424"
	protocolRFC3164Str = "rfc3164"
)

const (
	priority       = "priority"
	version        = "version"
	hostname       = "hostname"
	app            = "appname"
	pid            = "proc_id"
	msgID          = "msg_id"
	structuredData = "structured_data"
	message        = "message"
)

const (
	emptyValue   = "-"
	emptyMessage = ""
)

type sender struct {
	network   string
	addr      string
	protocol  string
	tlsConfig *tls.Config
	logger    *zap.Logger
	mu        sync.Mutex
	conn      net.Conn
}

func connect(ctx context.Context, logger *zap.Logger, cfg *Config, tlsConfig *tls.Config) (*sender, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *sender) close() error { _ = "STUB: not implemented"; return nil }

func (s *sender) dial(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (s *sender) Write(ctx context.Context, msgStr string) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *sender) write(msg string) error {
	_ = "STUB: not implemented"
	// check if logs contains new line character at the end, if not add it
	return nil
}

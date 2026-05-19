// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package osqueryreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/osqueryreceiver"

import (
	"context"
	"time"

	"github.com/osquery/osquery-go"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/receiver"
	"go.uber.org/zap"
)

const (
	defaultClientConnectRetries = 3
	defaultReconnectTimeout     = time.Millisecond * 200
	defaultQueryTimeout         = 30 * time.Second
)

type client interface {
	Close()
	QueryRowsContext(ctx context.Context, query string) ([]map[string]string, error)
}

var _ client = &osquery.ExtensionManagerClient{}

func makeOsQueryClient(socket string) (client, error) {
	_ = "STUB: not implemented"
	return *new(client), nil
}

type osQueryReceiver struct {
	config       *Config
	logger       *zap.Logger
	createClient func(socket string) (client, error)
}

func newOsQueryReceiver(cfg *Config, set receiver.Settings) *osQueryReceiver {
	_ = "STUB: not implemented"
	return nil
}

func (or *osQueryReceiver) connect(retries int) (client, error) {
	_ = "STUB: not implemented"
	return *new(client), nil
}

func (or *osQueryReceiver) runQuery(ctx context.Context, ld plog.Logs, query string) error {
	_ = "STUB: not implemented"
	return nil
}

// Use a separate connection for queries in order to be able to recover from timed out queries

func (or *osQueryReceiver) collect(ctx context.Context) (plog.Logs, error) {
	_ = "STUB: not implemented"
	return *new(plog.Logs), nil
}

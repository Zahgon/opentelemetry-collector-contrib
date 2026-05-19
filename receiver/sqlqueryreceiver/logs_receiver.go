// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package sqlqueryreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/sqlqueryreceiver"

import (
	"context"
	"database/sql"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/extension/xextension/storage"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/receiverhelper"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/sqlquery"
)

type logsReceiver struct {
	config           *Config
	settings         receiver.Settings
	createConnection sqlquery.DbProviderFunc
	createClient     sqlquery.ClientProviderFunc
	queryReceivers   []*logsQueryReceiver
	nextConsumer     consumer.Logs

	isStarted         bool
	shutdownRequested chan struct{}

	id            component.ID
	storageClient storage.Client
	obsrecv       *receiverhelper.ObsReport
}

func newLogsReceiver(
	config *Config,
	settings receiver.Settings,
	sqlOpenerFunc sqlquery.SQLOpenerFunc,
	createClient sqlquery.ClientProviderFunc,
	nextConsumer consumer.Logs,
) (*logsReceiver, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (receiver *logsReceiver) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (receiver *logsReceiver) createQueryReceivers() error { _ = "STUB: not implemented"; return nil }

func (receiver *logsReceiver) startCollecting() { _ = "STUB: not implemented"; return }

func (receiver *logsReceiver) collect() { _ = "STUB: not implemented"; return }

func (receiver *logsReceiver) Shutdown(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (receiver *logsReceiver) stopCollecting() { _ = "STUB: not implemented"; return }

type logsQueryReceiver struct {
	id           string
	query        sqlquery.Query
	createDb     sqlquery.DbProviderFunc
	createClient sqlquery.ClientProviderFunc
	logger       *zap.Logger
	telemetry    sqlquery.TelemetryConfig

	db            *sql.DB
	client        sqlquery.DbClient
	trackingValue string
	// TODO: Extract persistence into its own component
	storageClient           storage.Client
	trackingValueStorageKey string
}

func newLogsQueryReceiver(
	id string,
	query sqlquery.Query,
	dbProviderFunc sqlquery.DbProviderFunc,
	clientProviderFunc sqlquery.ClientProviderFunc,
	logger *zap.Logger,
	telemetry sqlquery.TelemetryConfig,
	storageClient storage.Client,
) *logsQueryReceiver {
	_ = "STUB: not implemented"
	return nil
}

func (queryReceiver *logsQueryReceiver) ID() string { _ = "STUB: not implemented"; return "" }

func (queryReceiver *logsQueryReceiver) start(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// retrieveTrackingValue retrieves the tracking value from storage, if storage is configured.
// Otherwise, it returns the tracking value configured in `tracking_start_value`.
func (queryReceiver *logsQueryReceiver) retrieveTrackingValue(ctx context.Context) string {
	_ = "STUB: not implemented"
	return ""
}

func (queryReceiver *logsQueryReceiver) collect(ctx context.Context) (plog.Logs, error) {
	_ = "STUB: not implemented"
	return *new(plog.Logs), nil
}

func (queryReceiver *logsQueryReceiver) storeTrackingValue(ctx context.Context, row sqlquery.StringMap) error {
	_ = "STUB: not implemented"
	return nil
}

func rowToLog(row sqlquery.StringMap, config sqlquery.LogsCfg, logRecord plog.LogRecord) error {
	_ = "STUB: not implemented"
	return nil
}

func (queryReceiver *logsQueryReceiver) shutdown(context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

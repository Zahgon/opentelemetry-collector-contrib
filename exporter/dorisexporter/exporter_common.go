// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package dorisexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/dorisexporter"

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql" // for register database driver
	"go.opentelemetry.io/collector/component"
	"go.uber.org/zap"
)

const timeFormat = "2006-01-02 15:04:05.999999"

type commonExporter struct {
	component.TelemetrySettings

	client *http.Client

	logger   *zap.Logger
	cfg      *Config
	timeZone *time.Location
	reporter *progressReporter
}

func newExporter(logger *zap.Logger, cfg *Config, set component.TelemetrySettings, reporterName string) *commonExporter {
	_ = "STUB: not implemented"
	return nil
}

func (e *commonExporter) formatTime(t time.Time) string { _ = "STUB: not implemented"; return "" }

type streamLoadResponse struct {
	TxnID                  int64
	Label                  string
	Status                 string
	ExistingJobStatus      string
	Message                string
	NumberTotalRows        int64
	NumberLoadedRows       int64
	NumberFilteredRows     int64
	NumberUnselectedRows   int64
	LoadBytes              int64
	LoadTimeMs             int64
	BeginTxnTimeMs         int64
	StreamLoadPutTimeMs    int64
	ReadDataTimeMs         int64
	WriteDataTimeMs        int64
	CommitAndPublishTimeMs int64
	ErrorURL               string
}

func (r *streamLoadResponse) success() bool { _ = "STUB: not implemented"; return false }

func (r *streamLoadResponse) duplication() bool { _ = "STUB: not implemented"; return false }

func streamLoadURL(address, db, table string) string { _ = "STUB: not implemented"; return "" }

func generateLabel(cfg *Config, table string) string { _ = "STUB: not implemented"; return "" }

func streamLoadRequest(ctx context.Context, cfg *Config, table string, data []byte, label string) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createDorisHTTPClient(ctx context.Context, cfg *Config, host component.Host, settings component.TelemetrySettings) (*http.Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createDorisMySQLClient(cfg *Config) (*sql.DB, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createAndUseDatabase(ctx context.Context, conn *sql.DB, cfg *Config) error {
	_ = "STUB: not implemented"
	return nil
}

type metric interface {
	dMetricGauge | dMetricSum | dMetricHistogram | dMetricExponentialHistogram | dMetricSummary
}

func toJSONLines[T dLog | dTrace | metric](data []*T) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

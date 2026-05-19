// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package snowflakereceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/snowflakereceiver"

import (
	"context"
	"database/sql"

	"go.opentelemetry.io/collector/component"
	"go.uber.org/zap"
)

// queries
var (
	billingMetricsQuery          = "select SERVICE_TYPE, NAME, sum(CREDITS_USED_COMPUTE), sum(CREDITS_USED_CLOUD_SERVICES), sum(CREDITS_USED) from METERING_HISTORY where start_time >= DATEADD(hour, -24, current_timestamp()) group by 1, 2;"
	warehouseBillingMetricsQuery = "select WAREHOUSE_NAME, sum(CREDITS_USED_COMPUTE), sum(CREDITS_USED_CLOUD_SERVICES), sum(CREDITS_USED) from WAREHOUSE_METERING_HISTORY where start_time >= DATEADD(hour, -24, current_timestamp()) group by 1;"
	loginMetricsQuery            = "select USER_NAME, ERROR_MESSAGE, REPORTED_CLIENT_TYPE, IS_SUCCESS, count(*) from LOGIN_HISTORY where event_timestamp >= DATEADD(hour, -24, current_timestamp()) group by 1, 2, 3, 4;"
	highLevelQueryMetricsQuery   = "select WAREHOUSE_NAME, AVG(AVG_RUNNING), AVG(AVG_QUEUED_LOAD), AVG(AVG_QUEUED_PROVISIONING), AVG(AVG_BLOCKED) from WAREHOUSE_LOAD_HISTORY where start_time >= DATEADD(hour, -24, current_timestamp()) group by 1;"
	dbMetricsQuery               = "select SCHEMA_NAME, EXECUTION_STATUS, ERROR_MESSAGE, QUERY_TYPE, WAREHOUSE_NAME, DATABASE_NAME, WAREHOUSE_SIZE, USER_NAME, COUNT(QUERY_ID), AVG(queued_overload_time), AVG(queued_repair_time), AVG(queued_provisioning_time), AVG(TOTAL_ELAPSED_TIME), AVG(EXECUTION_TIME), AVG(COMPILATION_TIME), AVG(BYTES_SCANNED), AVG(BYTES_WRITTEN), AVG(BYTES_DELETED), AVG(BYTES_SPILLED_TO_LOCAL_STORAGE), AVG(BYTES_SPILLED_TO_REMOTE_STORAGE), AVG(PERCENTAGE_SCANNED_FROM_CACHE), AVG(PARTITIONS_SCANNED), AVG(ROWS_UNLOADED), AVG(ROWS_DELETED), AVG(ROWS_UPDATED), AVG(ROWS_INSERTED), AVG(COALESCE(ROWS_PRODUCED,0)) from QUERY_HISTORY where start_time >= DATEADD(hour, -24, current_timestamp()) group by 1, 2, 3, 4, 5, 6, 7, 8;"
	sessionMetricsQuery          = "select USER_NAME, count(distinct(SESSION_ID)) from Sessions where created_on >= DATEADD(hour, -24, current_timestamp()) group by 1;"
	snowpipeMetricsQuery         = "select pipe_name, sum(credits_used), sum(bytes_inserted), sum(files_inserted) from pipe_usage_history where start_time >= DATEADD(hour, -24, current_timestamp()) group by 1;"
	storageMetricsQuery          = "select STORAGE_BYTES, STAGE_BYTES, FAILSAFE_BYTES from STORAGE_USAGE ORDER BY USAGE_DATE DESC LIMIT 1;"
)

// snowflake client is comprised of a sql.DB (the proper 'client' in question),
// a connection string (dsn), a collection of queries (built from which metrics are enabled),
// and a logger
type snowflakeClient struct {
	client *sql.DB
	dsn    *string
	logger *zap.Logger
}

// build snowflake db connection string
func buildDSN(cfg Config) (string, error) { _ = "STUB: not implemented"; return "", nil }

func newDefaultClient(settings component.TelemetrySettings, c Config) (*snowflakeClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// queries database and returns resulting rows
func (c snowflakeClient) readDB(ctx context.Context, q string) (*sql.Rows, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// these wrap readDB and return the associated data type which the scraper will
// use to generate and emit metrics. Which of these are called will be based on which metrics
// are enabled in the Config (default is all of them)
func (c snowflakeClient) FetchBillingMetrics(ctx context.Context) (*[]billingMetric, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c snowflakeClient) FetchWarehouseBillingMetrics(ctx context.Context) (*[]whBillingMetric, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c snowflakeClient) FetchLoginMetrics(ctx context.Context) (*[]loginMetric, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c snowflakeClient) FetchHighLevelQueryMetrics(ctx context.Context) (*[]hlQueryMetric, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c snowflakeClient) FetchDbMetrics(ctx context.Context) (*[]dbMetric, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c snowflakeClient) FetchSessionMetrics(ctx context.Context) (*[]sessionMetric, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c snowflakeClient) FetchSnowpipeMetrics(ctx context.Context) (*[]snowpipeMetric, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c snowflakeClient) FetchStorageMetrics(ctx context.Context) (*[]storageMetric, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

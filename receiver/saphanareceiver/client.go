// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package saphanareceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/saphanareceiver"

import (
	"context"
	"database/sql"
	"database/sql/driver"
)

// Interface for a SAP HANA client. Implementation can be faked for testing.
type client interface {
	Connect(ctx context.Context) error
	collectDataFromQuery(ctx context.Context, query *monitoringQuery) ([]map[string]string, error)
	Close() error
}

// Wraps the result of a query so that it can be mocked in tests
type resultWrapper interface {
	Scan(dest ...any) error
	Close() error
	Next() bool
}

// Wraps the sqlDB interface so that it can be mocked in tests
type dbWrapper interface {
	PingContext(ctx context.Context) error
	Close() error
	QueryContext(ctx context.Context, query string) (resultWrapper, error)
}

type standardResultWrapper struct {
	rows *sql.Rows
}

func (w *standardResultWrapper) Next() bool { _ = "STUB: not implemented"; return false }

func (w *standardResultWrapper) Scan(dest ...any) error { _ = "STUB: not implemented"; return nil }

func (w *standardResultWrapper) Close() error { _ = "STUB: not implemented"; return nil }

type standardDBWrapper struct {
	db *sql.DB
}

func (w *standardDBWrapper) Close() error { _ = "STUB: not implemented"; return nil }

func (w *standardDBWrapper) PingContext(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *standardDBWrapper) QueryContext(ctx context.Context, query string) (resultWrapper, error) {
	_ = "STUB: not implemented"
	return *new(resultWrapper), nil
}

// Wraps the creation of a sqlDB so that it can be mocked in tests
type sapHanaConnectionFactory interface {
	getConnection(c driver.Connector) dbWrapper
}

type defaultConnectionFactory struct{}

func (*defaultConnectionFactory) getConnection(c driver.Connector) dbWrapper {
	_ = "STUB: not implemented"
	return *new(dbWrapper)
}

// Wraps a SAP HANA database connection, implements `client` interface.
type sapHanaClient struct {
	receiverConfig    *Config
	connectionFactory sapHanaConnectionFactory
	client            dbWrapper
}

var _ client = (*sapHanaClient)(nil)

// Creates a SAP HANA database client
func newSapHanaClient(cfg *Config, factory sapHanaConnectionFactory) client {
	_ = "STUB: not implemented"
	return *new(client)
}

func (c *sapHanaClient) Connect(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (c *sapHanaClient) Close() error { _ = "STUB: not implemented"; return nil }

func (c *sapHanaClient) collectDataFromQuery(ctx context.Context, query *monitoringQuery) ([]map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Build a list of addresses that rows.Scan will load column data into

// If value was null, we can't use this row

// If value was null, we can't use this row

// Only report stat if value was not NULL

func convertInterfaceToString(input any) (sql.NullString, error) {
	_ = "STUB: not implemented"
	return *new(sql.NullString), nil
}

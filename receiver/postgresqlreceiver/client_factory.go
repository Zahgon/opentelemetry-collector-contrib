// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package postgresqlreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/postgresqlreceiver"

import (
	"database/sql"
	"sync"
)

type postgreSQLClientFactory interface {
	getClient(database string) (client, error)
	close() error
}

// defaultClientFactory creates one PG connection per call
type defaultClientFactory struct {
	baseConfig postgreSQLConfig
}

func newDefaultClientFactory(cfg *Config) *defaultClientFactory {
	_ = "STUB: not implemented"
	return nil
}

func (d *defaultClientFactory) getClient(database string) (client, error) {
	_ = "STUB: not implemented"
	return *new(client), nil
}

func (*defaultClientFactory) close() error {
	_ = "STUB: not implemented"

	// poolClientFactory creates one PG connection per database, keeping a pool of connections
	return nil
}

type poolClientFactory struct {
	sync.Mutex
	baseConfig postgreSQLConfig
	poolConfig *ConnectionPool
	pool       map[string]*sql.DB
	closed     bool
}

func newPoolClientFactory(cfg *Config) *poolClientFactory { _ = "STUB: not implemented"; return nil }

func (p *poolClientFactory) getClient(database string) (client, error) {
	_ = "STUB: not implemented"
	return *new(client), nil
}

func (p *poolClientFactory) close() error { _ = "STUB: not implemented"; return nil }

func (p *poolClientFactory) setPoolSettings(db *sql.DB) { _ = "STUB: not implemented"; return }

func getDB(cfg postgreSQLConfig, database string) (*sql.DB, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package datasource // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudspannerreceiver/internal/datasource"

import (
	"context"

	"cloud.google.com/go/spanner"
)

type Database struct {
	client     *spanner.Client
	databaseID *DatabaseID
}

func (database *Database) Client() *spanner.Client { _ = "STUB: not implemented"; return nil }

func (database *Database) DatabaseID() *DatabaseID { _ = "STUB: not implemented"; return nil }

func NewDatabase(ctx context.Context, databaseID *DatabaseID, credentialsFilePath string) (*Database, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Fallback to Application Default Credentials(https://google.aip.dev/auth/4110)

func NewDatabaseFromClient(client *spanner.Client, databaseID *DatabaseID) *Database {
	_ = "STUB: not implemented"
	return nil
}

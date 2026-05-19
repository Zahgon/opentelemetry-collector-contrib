// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package datasource // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudspannerreceiver/internal/datasource"

type DatabaseID struct {
	projectID    string
	instanceID   string
	databaseName string
	id           string
}

func NewDatabaseID(projectID, instanceID, databaseName string) *DatabaseID {
	_ = "STUB: not implemented"
	return nil
}

func (databaseID *DatabaseID) ProjectID() string { _ = "STUB: not implemented"; return "" }

func (databaseID *DatabaseID) InstanceID() string { _ = "STUB: not implemented"; return "" }

func (databaseID *DatabaseID) DatabaseName() string { _ = "STUB: not implemented"; return "" }

func (databaseID *DatabaseID) ID() string { _ = "STUB: not implemented"; return "" }

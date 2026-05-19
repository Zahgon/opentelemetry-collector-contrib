// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package sqlqueryreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/sqlqueryreceiver"

import (
	"go.opentelemetry.io/collector/component"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/sqlquery"
)

type Config struct {
	sqlquery.Config `mapstructure:",squash"`
	// The maximumn number of open connections to the sql server. <= 0 means unlimited
	MaxOpenConn int `mapstructure:"max_open_conn"`

	// prevent unkeyed literal initialization
	_ struct{}
}

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

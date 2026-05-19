// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package oraclecloud // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/oraclecloud"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/oraclecloud/internal/metadata"
)

type Config struct {
	ResourceAttributes metadata.ResourceAttributesConfig `mapstructure:"resource_attributes"`
}

func CreateDefaultConfig() Config { _ = "STUB: not implemented"; return *new(Config) }

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package hetzner // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/hetzner"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/hetzner/internal/metadata"
)

type Config struct {
	ResourceAttributes metadata.ResourceAttributesConfig `mapstructure:"resource_attributes"`
}

func CreateDefaultConfig() Config { _ = "STUB: not implemented"; return *new(Config) }

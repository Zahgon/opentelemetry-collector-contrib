// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package internal // import "github.com/open-telemetry/opentelemetry-collector-contrib/cmd/golden/internal"

import (
	"time"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest/pmetrictest"
)

type Config struct {
	ExpectedFile    string
	WriteExpected   bool
	CompareOptions  []pmetrictest.CompareMetricsOption
	OTLPEndpoint    string
	OTLPHTTPEndoint string
	Timeout         time.Duration
}

func ReadConfig(args []string) (*Config, error) { _ = "STUB: not implemented"; return nil, nil }

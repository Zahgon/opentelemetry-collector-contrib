// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package internal // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/pprofreceiver/internal"

import (
	"context"

	"go.opentelemetry.io/collector/pdata/pprofile"
	"go.uber.org/zap"
)

type FileScraper struct {
	Include string
	Logger  *zap.Logger
}

func (fs FileScraper) Scrape(_ context.Context) (pprofile.Profiles, error) {
	_ = "STUB: not implemented"
	return *new(pprofile.Profiles), nil
}

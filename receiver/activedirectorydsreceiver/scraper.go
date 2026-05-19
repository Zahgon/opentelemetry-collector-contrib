// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build windows

package activedirectorydsreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/activedirectorydsreceiver"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/activedirectorydsreceiver/internal/metadata"
)

type activeDirectoryDSScraper struct {
	mb *metadata.MetricsBuilder
	w  *watchers
}

func newActiveDirectoryDSScraper(mbc metadata.MetricsBuilderConfig, params receiver.Settings) *activeDirectoryDSScraper {
	_ = "STUB: not implemented"
	return nil
}

func (a *activeDirectoryDSScraper) start(_ context.Context, _ component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *activeDirectoryDSScraper) scrape(_ context.Context) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

//revive:disable-next-line:var-naming

//revive:disable-next-line:var-naming

func (a *activeDirectoryDSScraper) shutdown(context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

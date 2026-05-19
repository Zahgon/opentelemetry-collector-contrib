// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ciscoosreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/ciscoosreceiver"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/scraper"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/ciscoosreceiver/internal/scraper/interfacesscraper"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/ciscoosreceiver/internal/scraper/systemscraper"
)

var scraperFactories = map[component.Type]scraper.Factory{
	component.MustNewType("system"):     systemscraper.NewFactory(),
	component.MustNewType("interfaces"): interfacesscraper.NewFactory(),
}

func NewFactory() receiver.Factory { _ = "STUB: not implemented"; return *new(receiver.Factory) }

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

func createMetricsReceiver(
	_ context.Context,
	set receiver.Settings,
	cfg component.Config,
	consumer consumer.Metrics,
) (receiver.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Metrics), nil
}

type nopMetricsReceiver struct{}

func (*nopMetricsReceiver) Start(_ context.Context, _ component.Host) error {
	_ = "STUB: not implemented"
	return nil
}
func (*nopMetricsReceiver) Shutdown(_ context.Context) error { _ = "STUB: not implemented"; return nil }

type multiMetricsReceiver struct {
	receivers []receiver.Metrics
}

func (m *multiMetricsReceiver) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *multiMetricsReceiver) Shutdown(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

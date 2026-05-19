// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package resourcedetectionprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor"

import (
	"context"
	"sync"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/consumer/xconsumer"
	"go.opentelemetry.io/collector/processor"
	"go.opentelemetry.io/collector/processor/xprocessor"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal"
)

var consumerCapabilities = consumer.Capabilities{MutatesData: true}

type factory struct {
	resourceProviderFactory *internal.ResourceProviderFactory

	// providers stores a provider for each named processor that
	// may a different set of detectors configured.
	providers map[component.ID]*internal.ResourceProvider
	lock      sync.Mutex
}

// NewFactory creates a new factory for ResourceDetection processor.
func NewFactory() processor.Factory { _ = "STUB: not implemented"; return *new(processor.Factory) }

// Type gets the type of the Option config created by this factory.
func (*factory) Type() component.Type { _ = "STUB: not implemented"; return *new(component.Type) }

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

// TODO: Once issue(https://github.com/open-telemetry/opentelemetry-collector/issues/4001) gets resolved,
//		 Set the default value of 'hostname_source' here instead of 'system' detector

func defaultClientConfig() confighttp.ClientConfig {
	_ = "STUB: not implemented"
	return *new(confighttp.ClientConfig)
}

func (f *factory) createTracesProcessor(
	ctx context.Context,
	set processor.Settings,
	cfg component.Config,
	nextConsumer consumer.Traces,
) (processor.Traces, error) {
	_ = "STUB: not implemented"
	return *new(processor.Traces), nil
}

func (f *factory) createMetricsProcessor(
	ctx context.Context,
	set processor.Settings,
	cfg component.Config,
	nextConsumer consumer.Metrics,
) (processor.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(processor.Metrics), nil
}

func (f *factory) createLogsProcessor(
	ctx context.Context,
	set processor.Settings,
	cfg component.Config,
	nextConsumer consumer.Logs,
) (processor.Logs, error) {
	_ = "STUB: not implemented"
	return *new(processor.Logs), nil
}

func (f *factory) createProfilesProcessor(
	ctx context.Context,
	set processor.Settings,
	cfg component.Config,
	nextConsumer xconsumer.Profiles,
) (xprocessor.Profiles, error) {
	_ = "STUB: not implemented"
	return *new(xprocessor.Profiles), nil
}

func (f *factory) getResourceDetectionProcessor(
	params processor.Settings,
	cfg component.Config,
) (*resourceDetectionProcessor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *factory) getResourceProvider(
	params processor.Settings,
	timeout time.Duration,
	configuredDetectors []string,
	detectorConfigs DetectorConfig,
) (*internal.ResourceProvider, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

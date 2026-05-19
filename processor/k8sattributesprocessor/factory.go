// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package k8sattributesprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/k8sattributesprocessor"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/consumer/xconsumer"
	"go.opentelemetry.io/collector/processor"
	"go.opentelemetry.io/collector/processor/xprocessor"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/sharedcomponent"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/k8sattributesprocessor/internal/kube"
)

var (
	kubeClientProvider   = kube.ClientProvider(nil)
	consumerCapabilities = consumer.Capabilities{MutatesData: true}
	defaultExcludes      = ExcludeConfig{Pods: []ExcludePodConfig{{Name: "jaeger-agent"}, {Name: "jaeger-collector"}}}
	processors           = sharedcomponent.NewSharedComponents()
)

// NewFactory returns a new factory for the k8s processor.
func NewFactory() processor.Factory { _ = "STUB: not implemented"; return *new(processor.Factory) }

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

func createTracesProcessor(
	ctx context.Context,
	params processor.Settings,
	cfg component.Config,
	next consumer.Traces,
) (processor.Traces, error) {
	_ = "STUB: not implemented"
	return *new(processor.Traces), nil
}

func createLogsProcessor(
	ctx context.Context,
	params processor.Settings,
	cfg component.Config,
	nextLogsConsumer consumer.Logs,
) (processor.Logs, error) {
	_ = "STUB: not implemented"
	return *new(processor.Logs), nil
}

func createMetricsProcessor(
	ctx context.Context,
	params processor.Settings,
	cfg component.Config,
	nextMetricsConsumer consumer.Metrics,
) (processor.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(processor.Metrics), nil
}

func createProfilesProcessor(
	ctx context.Context,
	params processor.Settings,
	cfg component.Config,
	nextProfilesConsumer xconsumer.Profiles,
) (xprocessor.Profiles, error) {
	_ = "STUB: not implemented"
	return *new(xprocessor.Profiles), nil
}

func createTracesProcessorWithOptions(
	ctx context.Context,
	set processor.Settings,
	cfg component.Config,
	next consumer.Traces,
	options ...option,
) (processor.Traces, error) {
	_ = "STUB: not implemented"
	return *new(processor.Traces), nil
}

func createMetricsProcessorWithOptions(
	ctx context.Context,
	set processor.Settings,
	cfg component.Config,
	nextMetricsConsumer consumer.Metrics,
	options ...option,
) (processor.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(processor.Metrics), nil
}

func createLogsProcessorWithOptions(
	ctx context.Context,
	set processor.Settings,
	cfg component.Config,
	nextLogsConsumer consumer.Logs,
	options ...option,
) (processor.Logs, error) {
	_ = "STUB: not implemented"
	return *new(processor.Logs), nil
}

func createProfilesProcessorWithOptions(
	ctx context.Context,
	set processor.Settings,
	cfg component.Config,
	nextProfilesConsumer xconsumer.Profiles,
	options ...option,
) (xprocessor.Profiles, error) {
	_ = "STUB: not implemented"
	return *new(xprocessor.Profiles), nil
}

func createKubernetesProcessor(
	params processor.Settings,
	cfg component.Config,
	options ...option,
) *kubernetesprocessor {
	_ = "STUB: not implemented"
	return nil
}

func createProcessorOpts(cfg component.Config) []option { _ = "STUB: not implemented"; return nil }

// extraction rules

// filters

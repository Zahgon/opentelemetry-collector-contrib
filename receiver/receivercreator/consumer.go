// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package receivercreator // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/receivercreator"

import (
	"context"

	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/consumer/xconsumer"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/pprofile"
	"go.opentelemetry.io/collector/pdata/ptrace"

	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer"
)

var (
	_ consumer.Logs      = (*enhancingConsumer)(nil)
	_ consumer.Metrics   = (*enhancingConsumer)(nil)
	_ consumer.Traces    = (*enhancingConsumer)(nil)
	_ xconsumer.Profiles = (*enhancingConsumer)(nil)
)

// enhancingConsumer adds additional resource attributes from the given endpoint environment before passing the
// telemetry to its next consumers. The added attributes vary based on the type of the endpoint.
type enhancingConsumer struct {
	logs     consumer.Logs
	metrics  consumer.Metrics
	traces   consumer.Traces
	profiles xconsumer.Profiles
	attrs    map[string]string
}

func newEnhancingConsumer(
	resources resourceAttributes,
	receiverAttributes map[string]string,
	env observer.EndpointEnv,
	endpoint observer.Endpoint,
	nextLogs consumer.Logs,
	nextMetrics consumer.Metrics,
	nextTraces consumer.Traces,
	nextProfiles xconsumer.Profiles,
) (*enhancingConsumer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Precompute values that will be inserted for each resource object passed through.

// If the attribute value is empty this signals to delete existing

func (*enhancingConsumer) Capabilities() consumer.Capabilities {
	_ = "STUB: not implemented"
	return *new(consumer.Capabilities)
}

func (ec *enhancingConsumer) ConsumeLogs(ctx context.Context, ld plog.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

func (ec *enhancingConsumer) ConsumeMetrics(ctx context.Context, md pmetric.Metrics) error {
	_ = "STUB: not implemented"
	return nil
}

func (ec *enhancingConsumer) ConsumeTraces(ctx context.Context, td ptrace.Traces) error {
	_ = "STUB: not implemented"
	return nil
}

func (ec *enhancingConsumer) ConsumeProfiles(ctx context.Context, pd pprofile.Profiles) error {
	_ = "STUB: not implemented"
	return nil
}

func (ec *enhancingConsumer) putAttrs(attrs pcommon.Map) { _ = "STUB: not implemented"; return }

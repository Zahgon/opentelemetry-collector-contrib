// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package cloudfoundryreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/cloudfoundryreceiver"

import (
	"context"
	"sync"
	"time"

	"code.cloudfoundry.org/go-loggregator"
	"code.cloudfoundry.org/go-loggregator/rpc/loggregator_v2"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/receiverhelper"
)

const (
	transport  = "http"
	dataFormat = "cloudfoundry"
)

var (
	_ receiver.Metrics = (*cloudFoundryReceiver)(nil)
	_ receiver.Logs    = (*cloudFoundryReceiver)(nil)
)

// newCloudFoundryReceiver implements the receiver.Metrics and receiver.Logs for the Cloud Foundry protocol.
type cloudFoundryReceiver struct {
	settings          component.TelemetrySettings
	cancel            context.CancelFunc
	config            Config
	nextMetrics       consumer.Metrics
	nextLogs          consumer.Logs
	obsrecv           *receiverhelper.ObsReport
	goroutines        sync.WaitGroup
	receiverStartTime time.Time
}

// newCloudFoundryMetricsReceiver creates the Cloud Foundry receiver with the given parameters.
func newCloudFoundryMetricsReceiver(
	settings receiver.Settings,
	config Config,
	nextConsumer consumer.Metrics,
) (*cloudFoundryReceiver, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// newCloudFoundryLogsReceiver creates the Cloud Foundry logs receiver with the given parameters.
func newCloudFoundryLogsReceiver(
	settings receiver.Settings,
	config Config,
	nextConsumer consumer.Logs,
) (*cloudFoundryReceiver, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cfr *cloudFoundryReceiver) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (cfr *cloudFoundryReceiver) Shutdown(_ context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (cfr *cloudFoundryReceiver) streamMetrics(
	ctx context.Context,
	stream loggregator.EnvelopeStream,
	host component.Host,
) {
	_ = "STUB: not implemented"

	// Blocks until non-empty result or context is cancelled (returns nil in that case)
	return
}

// If context has not been cancelled, then nil means the shutdown was due to an error within stream

func (cfr *cloudFoundryReceiver) streamLogs(
	ctx context.Context,
	stream loggregator.EnvelopeStream,
	host component.Host,
) {
	_ = "STUB: not implemented"
	return
}

func buildLogs(logs plog.Logs, envelope *loggregator_v2.Envelope, observedTime time.Time) {
	_ = "STUB: not implemented"
	return
}

func buildMetrics(metrics pmetric.Metrics, envelope *loggregator_v2.Envelope, observedTime time.Time) {
	_ = "STUB: not implemented"
	return
}

func setupMetricsScope(resourceMetrics pmetric.ResourceMetrics) { _ = "STUB: not implemented"; return }

func getResourceMetrics(metrics pmetric.Metrics, envelope *loggregator_v2.Envelope) pmetric.ResourceMetrics {
	_ = "STUB: not implemented"
	return *new(pmetric.ResourceMetrics)
}

func setupLogsScope(resourceLogs plog.ResourceLogs) { _ = "STUB: not implemented"; return }

func getResourceLogs(logs plog.Logs, envelope *loggregator_v2.Envelope) plog.ResourceLogs {
	_ = "STUB: not implemented"
	return *new(plog.ResourceLogs)
}

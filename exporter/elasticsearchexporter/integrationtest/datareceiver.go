// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package integrationtest // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/elasticsearchexporter/integrationtest"

import (
	"context"
	"net/http"
	"testing"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/sharedcomponent"
	"github.com/open-telemetry/opentelemetry-collector-contrib/testbed/testbed"
)

const (
	// TestLogsIndex is used by the mock ES data receiver to identify log events.
	// Exporter LogsIndex configuration must be configured with TestLogsIndex for
	// the data receiver to work properly
	TestLogsIndex = "logs-test-idx"

	// TestMetricsIndex is used by the mock ES data receiver to identify metric events.
	// Exporter MetricsIndex configuration must be configured with TestMetricsIndex for
	// the data receiver to work properly
	TestMetricsIndex = "metrics-test-idx"

	// TestTracesIndex is used by the mock ES data receiver to identify trace
	// events. Exporter TracesIndex configuration must be configured with
	// TestTracesIndex for the data receiver to work properly
	TestTracesIndex = "traces-test-idx"
)

type errElasticsearch struct {
	httpStatus    int
	httpDocStatus int
}

func (e errElasticsearch) Error() string { _ = "STUB: not implemented"; return "" }

type esDataReceiver struct {
	testbed.DataReceiverBase
	receiver          receiver.Logs
	endpoint          string
	decodeBulkRequest bool
	enableBatching    bool
	t                 testing.TB
}

type dataReceiverOption func(*esDataReceiver)

func newElasticsearchDataReceiver(tb testing.TB, opts ...dataReceiverOption) *esDataReceiver {
	_ = "STUB: not implemented"
	return nil
}

func withDecodeBulkRequest(decode bool) dataReceiverOption {
	_ = "STUB: not implemented"
	return *new(dataReceiverOption)
}

func withBatching(enabled bool) dataReceiverOption {
	_ = "STUB: not implemented"
	return *new(dataReceiverOption)
}

func (es *esDataReceiver) Start(tc consumer.Traces, mc consumer.Metrics, lc consumer.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

// Use an actual logger to log errors.

// Since we use SharedComponent both receivers should be same

func (es *esDataReceiver) Stop() error { _ = "STUB: not implemented"; return nil }

func (es *esDataReceiver) GenConfigYAMLStr() string {
	_ = "STUB: not implemented"
	// Note that this generates an exporter config for agent.
	return ""
}

// Batching is disabled using `min_size` as we are setting batching
// as a default behavior.

func (*esDataReceiver) ProtocolName() string { _ = "STUB: not implemented"; return "" }

type config struct {
	confighttp.ServerConfig

	// DecodeBulkRequests controls decoding of the bulk request in the mock
	// ES receiver. Decoding requests would consume resources and might
	// pollute the benchmark results. Note that if decode bulk request is
	// set to false then the consumers will not consume any events and the
	// bulk request will always return http.StatusOK.
	DecodeBulkRequests bool
}

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

func createLogsReceiver(
	_ context.Context,
	params receiver.Settings,
	rawCfg component.Config,
	next consumer.Logs,
) (receiver.Logs, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Logs), nil
}

func createMetricsReceiver(
	_ context.Context,
	params receiver.Settings,
	rawCfg component.Config,
	next consumer.Metrics,
) (receiver.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Metrics), nil
}

func createTracesReceiver(
	_ context.Context,
	params receiver.Settings,
	rawCfg component.Config,
	next consumer.Traces,
) (receiver.Traces, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Traces), nil
}

type mockESReceiver struct {
	params receiver.Settings
	config *config

	tracesConsumer  consumer.Traces
	logsConsumer    consumer.Logs
	metricsConsumer consumer.Metrics

	server *http.Server
}

func newMockESReceiver(params receiver.Settings, cfg *config) receiver.Logs {
	_ = "STUB: not implemented"
	return *new(receiver.Logs)
}

func (es *mockESReceiver) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

// Assuming all documents are of the same type (logs, metrics, traces),
// create a pdata struct with the same number of records and send them in 1 Consume* call,
// i.e. a 1:1 bulk request to Consume* function call correspondence.
// This avoids a race condition where Consume* returns an error halfway through processing a bulk request,
// causing duplicates in the mock backend because the first N documents went through and an emulated http error
// causes the entire request to be retried, including the first N documents.

// panic to surface test logic error because we only expect error of type errElasticsearch

func (es *mockESReceiver) Shutdown(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// mockESReceiver serves both, traces and logs. Shared component allows for a single
// instance of mockESReceiver to serve all supported event types.
var receivers = sharedcomponent.NewSharedComponents()

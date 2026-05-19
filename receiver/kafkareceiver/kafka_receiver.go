// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package kafkareceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/kafkareceiver"

import (
	"context"
	"iter"

	"github.com/cenkalti/backoff/v4"
	"github.com/twmb/franz-go/pkg/kgo"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configretry"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/consumer/xconsumer"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/pprofile"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/receiverhelper"
	"go.opentelemetry.io/collector/receiver/xreceiver"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/kafkareceiver/internal/metadata"
)

const transport = "kafka"

type consumeMessageFunc func(ctx context.Context, record *kgo.Record, attrs attribute.Set) error

type newConsumeMessageFunc func(host component.Host, obsrecv *receiverhelper.ObsReport,
	telBldr *metadata.TelemetryBuilder,
) (consumeMessageFunc, error)

// messageHandler provides a generic interface for handling messages for a pdata type.
type messageHandler[T plog.Logs | pmetric.Metrics | ptrace.Traces | pprofile.Profiles] interface {
	// unmarshalData unmarshals the message payload into a pdata type (plog.Logs, etc.)
	// and returns the number of items (log records, metric data points, spans) within it.
	unmarshalData(data []byte) (T, int, error)

	// consumeData passes the unmarshaled data to the next consumer for the signal type.
	// This simply calls the signal-specific Consume* method.
	consumeData(ctx context.Context, data T) error

	// getResources returns the resources associated with the unmarshaled data.
	// This is used for header extraction for adding resource attributes.
	getResources(T) iter.Seq[pcommon.Resource]

	// startObsReport starts an observation report for the unmarshaled data.
	//
	// This simply calls the signal-specific receiverhelper.ObsReport.Start*Op method.
	startObsReport(ctx context.Context) context.Context

	// endObsReport ends the observation report for the unmarshaled data.
	//
	// This simply calls the signal-specific receiverhelper.ObsReport.End*Op method,
	// passing the configured encoding and number of items returned by unmarshalData.
	endObsReport(ctx context.Context, n int, err error)

	// getUnmarshalFailureCounter returns the appropriate telemetry counter for unmarshal failures
	getUnmarshalFailureCounter(telBldr *metadata.TelemetryBuilder) metric.Int64Counter
}

func newLogsReceiver(config *Config, set receiver.Settings, nextConsumer consumer.Logs) (receiver.Logs, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Logs), nil
}

func newMetricsReceiver(config *Config, set receiver.Settings, nextConsumer consumer.Metrics) (receiver.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Metrics), nil
}

func newTracesReceiver(config *Config, set receiver.Settings, nextConsumer consumer.Traces) (receiver.Traces, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Traces), nil
}

func newProfilesReceiver(config *Config, set receiver.Settings, nextConsumer xconsumer.Profiles) (xreceiver.Profiles, error) {
	_ = "STUB: not implemented"
	return *new(xreceiver.Profiles), nil
}

func newReceiver(
	config *Config,
	set receiver.Settings,
	topics []string,
	excludeTopics []string,
	consumeFn func(host component.Host,
		obsrecv *receiverhelper.ObsReport,
		telBldr *metadata.TelemetryBuilder,
	) (consumeMessageFunc, error),
) (component.Component, error) {
	_ = "STUB: not implemented"
	return *new(component.Component), nil
}

type logsHandler struct {
	unmarshaler plog.Unmarshaler
	obsrecv     *receiverhelper.ObsReport
	consumer    consumer.Logs
	encoding    string
}

func (h *logsHandler) unmarshalData(data []byte) (plog.Logs, int, error) {
	_ = "STUB: not implemented"
	return *new(plog.Logs), 0, nil
}

func (h *logsHandler) consumeData(ctx context.Context, data plog.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *logsHandler) startObsReport(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (h *logsHandler) endObsReport(ctx context.Context, n int, err error) {
	_ = "STUB: not implemented"
	return
}

func (*logsHandler) getResources(data plog.Logs) iter.Seq[pcommon.Resource] {
	_ = "STUB: not implemented"
	return nil
}

func (*logsHandler) getUnmarshalFailureCounter(telBldr *metadata.TelemetryBuilder) metric.Int64Counter {
	_ = "STUB: not implemented"
	return *new(metric.Int64Counter)
}

type metricsHandler struct {
	unmarshaler pmetric.Unmarshaler
	obsrecv     *receiverhelper.ObsReport
	consumer    consumer.Metrics
	encoding    string
}

func (h *metricsHandler) unmarshalData(data []byte) (pmetric.Metrics, int, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), 0, nil
}

func (h *metricsHandler) consumeData(ctx context.Context, data pmetric.Metrics) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *metricsHandler) startObsReport(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (h *metricsHandler) endObsReport(ctx context.Context, n int, err error) {
	_ = "STUB: not implemented"
	return
}

func (*metricsHandler) getResources(data pmetric.Metrics) iter.Seq[pcommon.Resource] {
	_ = "STUB: not implemented"
	return nil
}

func (*metricsHandler) getUnmarshalFailureCounter(telBldr *metadata.TelemetryBuilder) metric.Int64Counter {
	_ = "STUB: not implemented"
	return *new(metric.Int64Counter)
}

type tracesHandler struct {
	unmarshaler ptrace.Unmarshaler
	obsrecv     *receiverhelper.ObsReport
	consumer    consumer.Traces
	encoding    string
}

func (h *tracesHandler) unmarshalData(data []byte) (ptrace.Traces, int, error) {
	_ = "STUB: not implemented"
	return *new(ptrace.Traces), 0, nil
}

func (h *tracesHandler) consumeData(ctx context.Context, data ptrace.Traces) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *tracesHandler) startObsReport(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (h *tracesHandler) endObsReport(ctx context.Context, n int, err error) {
	_ = "STUB: not implemented"
	return
}

func (*tracesHandler) getResources(data ptrace.Traces) iter.Seq[pcommon.Resource] {
	_ = "STUB: not implemented"
	return nil
}

func (*tracesHandler) getUnmarshalFailureCounter(telBldr *metadata.TelemetryBuilder) metric.Int64Counter {
	_ = "STUB: not implemented"
	return *new(metric.Int64Counter)
}

type profilesHandler struct {
	unmarshaler pprofile.Unmarshaler
	obsrecv     *receiverhelper.ObsReport
	consumer    xconsumer.Profiles
	encoding    string
}

func (h *profilesHandler) unmarshalData(data []byte) (pprofile.Profiles, int, error) {
	_ = "STUB: not implemented"
	return *new(pprofile.Profiles), 0, nil
}

func (h *profilesHandler) consumeData(ctx context.Context, data pprofile.Profiles) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *profilesHandler) startObsReport(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (h *profilesHandler) endObsReport(ctx context.Context, n int, err error) {
	_ = "STUB: not implemented"
	return
}

func (*profilesHandler) getResources(data pprofile.Profiles) iter.Seq[pcommon.Resource] {
	_ = "STUB: not implemented"
	return nil
}

func (*profilesHandler) getUnmarshalFailureCounter(telBldr *metadata.TelemetryBuilder) metric.Int64Counter {
	_ = "STUB: not implemented"
	return *new(metric.Int64Counter)
}

// processMessage is a generic function that processes a Kafka record (*kgo.Record) using a messageHandler
func processMessage[T plog.Logs | pmetric.Metrics | ptrace.Traces | pprofile.Profiles](
	ctx context.Context,
	record *kgo.Record,
	config *Config,
	logger *zap.Logger,
	telBldr *metadata.TelemetryBuilder,
	handler messageHandler[T],
	attrs attribute.Set,
	headerAttrKeys map[string]string,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Return permanent error for unmarshalling failures

// Add resource attributes from headers if configured

func getMessageHeaderResourceAttributes(headers []kgo.RecordHeader, headerKeys map[string]string) iter.Seq2[string, string] {
	_ = "STUB: not implemented"
	return nil
}

// buildHeaderAttrKeys pre-computes the mapping from raw header names to their
// "kafka.header." prefixed attribute keys. Returns nil when header extraction
// is disabled.
func buildHeaderAttrKeys(config *Config) map[string]string { _ = "STUB: not implemented"; return nil }

func newExponentialBackOff(config configretry.BackOffConfig) *backoff.ExponentialBackOff {
	_ = "STUB: not implemented"
	return nil
}

func contextWithMetadata(ctx context.Context, record *kgo.Record) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package sumologicexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/sumologicexporter"

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/sumologicexporter/internal/metadata"
)

var (
	metricsMarshaler = pmetric.ProtoMarshaler{}
	logsMarshaler    = plog.ProtoMarshaler{}
	tracesMarshaler  = ptrace.ProtoMarshaler{}
)

// metricPair represents information required to send one metric to the Sumo Logic
type metricPair struct {
	attributes pcommon.Map
	metric     pmetric.Metric
}

// countingReader keeps number of records related to reader
type countingReader struct {
	counter int64
	reader  io.Reader
}

// newCountingReader creates countingReader with given number of records
func newCountingReader(records int) *countingReader { _ = "STUB: not implemented"; return nil }

// withBytes sets up reader to read from bytes data
func (c *countingReader) withBytes(data []byte) *countingReader {
	_ = "STUB: not implemented"
	return nil
}

// withString sets up reader to read from string data
func (c *countingReader) withString(data string) *countingReader {
	_ = "STUB: not implemented"
	return nil
}

// bodyBuilder keeps information about number of records related to data it keeps
type bodyBuilder struct {
	builder strings.Builder
	counter int
}

// newBodyBuilder returns empty bodyBuilder
func newBodyBuilder() bodyBuilder {
	_ = "STUB: not implemented"
	return *

	// Reset resets both counter and builder content
	new(bodyBuilder)
}

func (b *bodyBuilder) Reset() { _ = "STUB: not implemented"; return }

// addLine adds multiple lines to builder and increments counter
func (b *bodyBuilder) addLines(lines []string) { _ = "STUB: not implemented"; return }

// add the first line separately to avoid a conditional in the loop

// WriteString can't actually return an error

// addNewLine adds newline to builder
func (b *bodyBuilder) addNewLine() { _ = "STUB: not implemented"; return }

// WriteByte can't actually return an error

// Len returns builder content length
func (b *bodyBuilder) Len() int { _ = "STUB: not implemented"; return 0 }

// toCountingReader converts bodyBuilder to countingReader
func (b *bodyBuilder) toCountingReader() *countingReader { _ = "STUB: not implemented"; return nil }

type sender struct {
	logger                     *zap.Logger
	config                     *Config
	client                     *http.Client
	prometheusFormatter        prometheusFormatter
	dataURLMetrics             string
	dataURLLogs                string
	dataURLTraces              string
	stickySessionCookieFunc    func() string
	setStickySessionCookieFunc func(string)
	id                         component.ID
	telemetryBuilder           *metadata.TelemetryBuilder
}

const (
	headerContentType string = "Content-Type"
	headerClient      string = "X-Sumo-Client"
	headerHost        string = "X-Sumo-Host"
	headerName        string = "X-Sumo-Name"
	headerCategory    string = "X-Sumo-Category"
	headerFields      string = "X-Sumo-Fields"

	attributeKeySourceHost     = "_sourceHost"
	attributeKeySourceName     = "_sourceName"
	attributeKeySourceCategory = "_sourceCategory"

	contentTypeLogs       string = "application/x-www-form-urlencoded"
	contentTypePrometheus string = "application/vnd.sumologic.prometheus"
	contentTypeOTLP       string = "application/x-protobuf"
	stickySessionKey      string = "AWSALB"
)

func newSender(
	logger *zap.Logger,
	cfg *Config,
	cl *http.Client,
	pf prometheusFormatter,
	metricsURL string,
	logsURL string,
	tracesURL string,
	stickySessionCookieFunc func() string,
	setStickySessionCookieFunc func(string),
	id component.ID,
	telemetryBuilder *metadata.TelemetryBuilder,
) *sender {
	_ = "STUB: not implemented"
	return nil
}

var errUnauthorized = errors.New("unauthorized")

// send sends data to sumologic
func (s *sender) send(ctx context.Context, pipeline PipelineType, reader *countingReader, flds fields) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *sender) handleReceiverResponse(resp *http.Response) error {
	_ = "STUB: not implemented"
	return nil
}

// API responds with a 200 or 204 with ContentLength set to 0 when all data
// has been successfully ingested.

// API responds with a 200 or 204 with a JSON body describing what issues
// were encountered when processing the sent data.

// Report the failure as permanent if the server thinks the request is malformed.

func (s *sender) createRequest(ctx context.Context, pipeline PipelineType, data io.Reader) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// logToText converts LogRecord to a plain text line, returns it and error eventually
func (*sender) logToText(record plog.LogRecord) string { _ = "STUB: not implemented"; return "" }

// logToJSON converts LogRecord to a json line, returns it and error eventually
func (*sender) logToJSON(record plog.LogRecord) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Only append the body when it's not empty to prevent sending 'null' log.

func isEmptyAttributeValue(att pcommon.Value) bool { _ = "STUB: not implemented"; return false }

// sendNonOTLPLogs sends log records from the logBuffer formatted according
// to configured LogFormat and as the result of execution
// returns array of records which has not been sent correctly and error
func (s *sender) sendNonOTLPLogs(ctx context.Context, rl plog.ResourceLogs, flds fields) ([]plog.LogRecord, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If data was sent and either failed or succeeded, cleanup the currentRecords slice

func (s *sender) formatLogLine(lr plog.LogRecord) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// TODO: add support for HTTP limits
func (s *sender) sendOTLPLogs(ctx context.Context, ld plog.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

// sendNonOTLPMetrics sends metrics in right format basing on the s.config.MetricFormat
func (s *sender) sendNonOTLPMetrics(ctx context.Context, md pmetric.Metrics) (pmetric.Metrics, []error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

// generally speaking, it's fine to send multiple ResourceMetrics in a single request
// the only exception is if the computed source headers are different, as those as unique per-request
// so we check if the headers are different here and send what we have if they are

// transform the metrics into formatted lines ready to be sent

// failed at sending, add the resource to the dropped metrics
// move instead of copy here to avoid duplicating data in memory on failure

// If data was sent, cleanup the currentResources slice

func (s *sender) sendOTLPMetrics(ctx context.Context, md pmetric.Metrics) error {
	_ = "STUB: not implemented"
	return nil
}

// appendAndMaybeSend appends line to the request body that will be sent and sends
// the accumulated data if the internal logBuffer has been filled (with config.MaxRequestBodySize bytes).
// It returns a boolean indicating if the data was sent and an error
func (s *sender) appendAndMaybeSend(
	ctx context.Context,
	lines []string,
	pipeline PipelineType,
	body *bodyBuilder,
	flds fields,
) (sent bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

// count the newline as well

// Do not add newline if the body is empty

// sendTraces sends traces in right format basing on the s.config.TraceFormat
func (s *sender) sendTraces(ctx context.Context, td ptrace.Traces) error {
	_ = "STUB: not implemented"
	return nil
}

// sendOTLPTraces sends trace records in OTLP format
func (s *sender) sendOTLPTraces(ctx context.Context, td ptrace.Traces) error {
	_ = "STUB: not implemented"
	return nil
}

func addSourcesHeaders(req *http.Request, flds fields) { _ = "STUB: not implemented"; return }

func getSourcesHeaders(flds fields) map[string]string { _ = "STUB: not implemented"; return nil }

func addLogsHeaders(req *http.Request, lf LogFormatType, flds fields) {
	_ = "STUB: not implemented"
	return
}

func addMetricsHeaders(req *http.Request, mf MetricFormatType) error {
	_ = "STUB: not implemented"
	return nil
}

func addTracesHeaders(req *http.Request) { _ = "STUB: not implemented"; return }

func (s *sender) addRequestHeaders(req *http.Request, pipeline PipelineType, flds fields) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *sender) recordMetrics(duration time.Duration, count int64, req *http.Request, resp *http.Response, pipeline PipelineType) {
	_ = "STUB: not implemented"
	return
}

func (s *sender) addStickySessionCookie(req *http.Request) { _ = "STUB: not implemented"; return }

func (s *sender) updateStickySessionCookie(resp *http.Response) { _ = "STUB: not implemented"; return }

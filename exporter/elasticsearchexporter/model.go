// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package elasticsearchexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/elasticsearchexporter"

import (
	"bytes"
	"errors"
	"time"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pprofile"
	"go.opentelemetry.io/collector/pdata/ptrace"
	conventionsv126 "go.opentelemetry.io/otel/semconv/v1.26.0"
	conventions "go.opentelemetry.io/otel/semconv/v1.40.0"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/elasticsearchexporter/internal/datapoints"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/elasticsearchexporter/internal/elasticsearch"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/elasticsearchexporter/internal/objmodel"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/elasticsearchexporter/internal/serializer/otelserializer"
)

type conversionEntry struct {
	to               string
	preserveOriginal bool
	skip             bool
	skipIfExists     bool
}

// collectECSFields extracts all target ECS field paths from conversion maps
// and returns them as a set (map) for efficient lookups.
func collectECSFields(maps ...map[string]conversionEntry) map[string]struct{} {
	_ = "STUB: not implemented"
	return nil
}

// resourceAttrsConversionMap contains conversions for resource-level attributes
// from their Semantic Conventions (SemConv) names to equivalent Elastic Common
// Schema (ECS) names.
// If the ECS field name is specified as an empty string (""), the converter will
// neither convert the SemConv key to the equivalent ECS name nor pass-through the
// SemConv key as-is to become the ECS name.
var resourceAttrsConversionMap = map[string]conversionEntry{
	string(conventions.ServiceInstanceIDKey):         {to: "service.node.name"},
	string(conventionsv126.DeploymentEnvironmentKey): {to: "service.environment"},
	string(conventions.DeploymentEnvironmentNameKey): {to: "service.environment"},
	string(conventions.TelemetrySDKNameKey):          {skip: true},
	string(conventions.TelemetrySDKLanguageKey):      {to: "service.language.name"},
	string(conventions.TelemetrySDKVersionKey):       {to: "service.language.version"},
	string(conventions.TelemetryDistroNameKey):       {skip: true},
	string(conventions.TelemetryDistroVersionKey):    {skip: true},
	string(conventions.CloudPlatformKey):             {to: "cloud.service.name"},
	string(conventions.ContainerImageTagsKey):        {to: "container.image.tag"},
	string(conventions.HostNameKey):                  {to: "host.hostname", preserveOriginal: true, skipIfExists: true},
	string(conventions.HostArchKey):                  {to: "host.architecture"},
	string(conventions.ProcessParentPIDKey):          {to: "process.parent.pid"},
	string(conventions.ProcessExecutableNameKey):     {to: "process.title"},
	string(conventions.ProcessExecutablePathKey):     {to: "process.executable"},
	string(conventions.ProcessCommandLineKey):        {to: "process.args"},
	string(conventions.ProcessRuntimeNameKey):        {to: "service.runtime.name"},
	string(conventions.ProcessRuntimeVersionKey):     {to: "service.runtime.version"},
	string(conventions.OSNameKey):                    {to: "host.os.name"},
	string(conventions.OSTypeKey):                    {to: "host.os.platform"},
	string(conventions.OSDescriptionKey):             {to: "host.os.full"},
	string(conventions.OSVersionKey):                 {to: "host.os.version"},
	string(conventions.ClientAddressKey):             {to: "client.ip"},
	string(conventions.SourceAddressKey):             {to: "source.ip"},
	string(conventions.K8SDeploymentNameKey):         {to: "kubernetes.deployment.name"},
	string(conventions.K8SNamespaceNameKey):          {to: "kubernetes.namespace"},
	string(conventions.K8SNodeNameKey):               {to: "kubernetes.node.name"},
	string(conventions.K8SPodNameKey):                {to: "kubernetes.pod.name"},
	string(conventions.K8SPodUIDKey):                 {to: "kubernetes.pod.uid"},
	string(conventions.K8SJobNameKey):                {to: "kubernetes.job.name"},
	string(conventions.K8SCronJobNameKey):            {to: "kubernetes.cronjob.name"},
	string(conventions.K8SStatefulSetNameKey):        {to: "kubernetes.statefulset.name"},
	string(conventions.K8SReplicaSetNameKey):         {to: "kubernetes.replicaset.name"},
	string(conventions.K8SDaemonSetNameKey):          {to: "kubernetes.daemonset.name"},
	string(conventions.K8SContainerNameKey):          {to: "kubernetes.container.name"},
	string(conventions.K8SClusterNameKey):            {to: "orchestrator.cluster.name"},
	string(conventions.FaaSInstanceKey):              {to: "faas.id"},
	string(conventions.FaaSTriggerKey):               {to: "faas.trigger.type"},
}

var (
	scopeAttrsConversionMap = map[string]conversionEntry{}

	logRecordAttrsConversionMap = map[string]conversionEntry{
		"event.name":                                {to: "event.action"},
		string(conventions.ExceptionMessageKey):     {to: "error.message"},
		string(conventions.ExceptionStacktraceKey):  {to: "error.stacktrace"},
		string(conventions.ExceptionTypeKey):        {to: "error.type"},
		string(conventionsv126.ExceptionEscapedKey): {to: "event.error.exception.handled"},
		string(conventions.HTTPResponseBodySizeKey): {to: "http.response.encoded_body_size"},
	}

	spanAttrsConversionMap = map[string]conversionEntry{
		string(conventionsv126.DBSystemKey):         {to: "span.db.type"},
		string(conventions.DBNamespaceKey):          {to: "span.db.instance"},
		string(conventions.DBQueryTextKey):          {to: "span.db.statement"},
		string(conventions.HTTPResponseBodySizeKey): {to: "http.response.encoded_body_size"},
	}

	// Precomputed protected fields for performance
	logProtectedFields = collectECSFields(
		resourceAttrsConversionMap,
		scopeAttrsConversionMap,
		logRecordAttrsConversionMap,
	)
	spanProtectedFields = collectECSFields(
		resourceAttrsConversionMap,
		scopeAttrsConversionMap,
		spanAttrsConversionMap,
	)
	metricsProtectedFields = collectECSFields(resourceAttrsConversionMap)
)

var ErrInvalidTypeForBodyMapMode = errors.New("invalid log record body type for 'bodymap' mapping mode")

// documentEncoder is an interface for encoding signals to Elasticsearch documents.
type documentEncoder interface {
	encodeLog(encodingContext, plog.LogRecord, elasticsearch.Index, *bytes.Buffer) error
	encodeSpan(encodingContext, ptrace.Span, elasticsearch.Index, *bytes.Buffer) error
	encodeSpanEvent(encodingContext, ptrace.Span, ptrace.SpanEvent, elasticsearch.Index, *bytes.Buffer) error
	encodeMetrics(_ encodingContext, _ []datapoints.DataPoint, validationErrors *[]error, _ elasticsearch.Index, _ *bytes.Buffer) (map[string]string, error)
	encodeProfile(_ encodingContext, _ pprofile.ProfilesDictionary, _ pprofile.Profile, _ func(*bytes.Buffer, string, string) error) error
}

type encodingContext struct {
	resource          pcommon.Resource
	resourceSchemaURL string
	scope             pcommon.InstrumentationScope
	scopeSchemaURL    string
}

func newEncoder(mode MappingMode) (documentEncoder, error) {
	_ = "STUB: not implemented"
	return *new(documentEncoder), nil
}

type legacyModeEncoder struct {
	nonOTelSpanEncoder
	nopSpanEventEncoder
	metricsUnsupportedEncoder
	profilesUnsupportedEncoder
	attributesPrefix string
}

type ecsModeEncoder struct {
	ecsDataPointsEncoder
	nopSpanEventEncoder
	profilesUnsupportedEncoder
}

type bodymapModeEncoder struct {
	metricsUnsupportedEncoder
	profilesUnsupportedEncoder
}

type otelModeEncoder struct {
	serializer *otelserializer.Serializer
}

const (
	traceIDField   = "traceID"
	spanIDField    = "spanID"
	attributeField = "attribute"
)

func (e legacyModeEncoder) encodeLog(ec encodingContext, record plog.LogRecord, idx elasticsearch.Index, buf *bytes.Buffer) error {
	_ = "STUB: not implemented"
	return nil
}

// We use @timestamp in order to ensure that we can index if the default data stream logs template is used.

func (ecsModeEncoder) encodeLog(
	ec encodingContext,
	record plog.LogRecord,
	idx elasticsearch.Index,
	buf *bytes.Buffer,
) error {
	_ = "STUB: not implemented"
	return nil
}

// First, try to map resource-level attributes to ECS fields.

// Then, try to map scope-level attributes to ECS fields.

// Finally, try to map record-level attributes to ECS fields.

// Handle special cases.

func (ecsModeEncoder) encodeSpan(
	ec encodingContext,
	span ptrace.Span,
	idx elasticsearch.Index,
	buf *bytes.Buffer,
) error {
	_ = "STUB: not implemented"
	return nil
}

// First, try to map resource-level attributes to ECS fields.

// Then, try to map scope-level attributes to ECS fields.

// Finally, try to map span-level attributes to ECS fields.

// spanKindToECSStr converts an OTel SpanKind to its ECS equivalent string representation defined here:
// https://github.com/elastic/apm-data/blob/main/input/elasticapm/internal/modeldecoder/v2/decoder.go#L1665-L1669
func spanKindToECSStr(sk ptrace.SpanKind) string { _ = "STUB: not implemented"; return "" }

func (e otelModeEncoder) encodeLog(
	ec encodingContext,
	record plog.LogRecord,
	idx elasticsearch.Index,
	buf *bytes.Buffer,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (e otelModeEncoder) encodeSpan(
	ec encodingContext,
	span ptrace.Span,
	idx elasticsearch.Index,
	buf *bytes.Buffer,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (e otelModeEncoder) encodeSpanEvent(
	ec encodingContext,
	span ptrace.Span,
	spanEvent ptrace.SpanEvent,
	idx elasticsearch.Index,
	buf *bytes.Buffer,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (e otelModeEncoder) encodeMetrics(
	ec encodingContext,
	dataPoints []datapoints.DataPoint,
	validationErrors *[]error,
	idx elasticsearch.Index,
	buf *bytes.Buffer,
) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e otelModeEncoder) encodeProfile(
	ec encodingContext,
	dic pprofile.ProfilesDictionary,
	profile pprofile.Profile,
	pushData func(*bytes.Buffer, string, string) error,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (bodymapModeEncoder) encodeLog(
	_ encodingContext,
	record plog.LogRecord,
	_ elasticsearch.Index,
	buf *bytes.Buffer,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (bodymapModeEncoder) encodeSpan(encodingContext, ptrace.Span, elasticsearch.Index, *bytes.Buffer) error {
	_ = "STUB: not implemented"
	return nil
}

func (bodymapModeEncoder) encodeSpanEvent(encodingContext, ptrace.Span, ptrace.SpanEvent, elasticsearch.Index, *bytes.Buffer) error {
	_ = "STUB: not implemented"
	return nil
}

type metricsUnsupportedEncoder struct {
	mode MappingMode
}

//nolint:unparam // result 0 is expected to always be nil
func (e metricsUnsupportedEncoder) encodeMetrics(
	_ encodingContext,
	_ []datapoints.DataPoint,
	_ *[]error,
	_ elasticsearch.Index,
	_ *bytes.Buffer,
) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type profilesUnsupportedEncoder struct {
	mode MappingMode
}

func (e profilesUnsupportedEncoder) encodeProfile(
	_ encodingContext, _ pprofile.ProfilesDictionary, _ pprofile.Profile, _ func(*bytes.Buffer, string, string) error,
) error {
	_ = "STUB: not implemented"
	return nil
}

type nonOTelSpanEncoder struct {
	attributesPrefix string
	eventsPrefix     string
	dedot            bool
}

func (e nonOTelSpanEncoder) encodeSpan(
	ec encodingContext,
	span ptrace.Span,
	idx elasticsearch.Index,
	buf *bytes.Buffer,
) error {
	_ = "STUB: not implemented"
	return nil
}

// We use @timestamp in order to ensure that we can index if the default data stream logs template is used.

// unit is microseconds

type ecsDataPointsEncoder struct{}

func (ecsDataPointsEncoder) encodeMetrics(
	ec encodingContext,
	dataPoints []datapoints.DataPoint,
	validationErrors *[]error,
	idx elasticsearch.Index,
	buf *bytes.Buffer,
) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func addDataStreamAttributes(document *objmodel.Document, key string, idx elasticsearch.Index) {
	_ = "STUB: not implemented"
	return
}

// nopSpanEventEncoder is embedded in all non-OTel encoders,
// since only OTel mapping mode currently encodes span events
// as separate documents. In all others they are stored within
// the span document.
type nopSpanEventEncoder struct{}

func (nopSpanEventEncoder) encodeSpanEvent(encodingContext, ptrace.Span, ptrace.SpanEvent, elasticsearch.Index, *bytes.Buffer) error {
	_ = "STUB: not implemented"
	return nil
}

func encodeAttributes(prefix string, document *objmodel.Document, attributes pcommon.Map, idx elasticsearch.Index) {
	_ = "STUB: not implemented"
	return
}

func spanLinksToString(spanLinkSlice ptrace.SpanLinkSlice) string {
	_ = "STUB: not implemented"
	return ""
}

// durationAsMicroseconds calculate span duration through end - start nanoseconds and converts time.Time to microseconds,
// which is the format the Duration field is stored in the Span.
func durationAsMicroseconds(start, end time.Time) int64 { _ = "STUB: not implemented"; return 0 }

func scopeToAttributes(scope pcommon.InstrumentationScope) pcommon.Map {
	_ = "STUB: not implemented"
	return *new(pcommon.Map)
}

func encodeAttributesECSMode(document *objmodel.Document, attrs pcommon.Map, conversionMap map[string]conversionEntry) {
	_ = "STUB: not implemented"
	return
}

// No conversions to be done; add all attributes at top level of
// document.

// If ECS key is found for current k in conversion map, use it.

// Skip the conversion for this k.

// Otherwise, add key at top level with attribute name as-is.

func encodeLogTimestampECSMode(document *objmodel.Document, record plog.LogRecord) {
	_ = "STUB: not implemented"
	return
}

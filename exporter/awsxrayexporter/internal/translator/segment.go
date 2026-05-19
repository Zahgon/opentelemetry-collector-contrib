// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package translator // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsxrayexporter/internal/translator"

import (
	"regexp"

	"go.opentelemetry.io/collector/featuregate"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"

	awsxray "github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/xray"
)

// AWS X-Ray acceptable values for origin field.
const (
	OriginEC2        = "AWS::EC2::Instance"
	OriginECS        = "AWS::ECS::Container"
	OriginECSEC2     = "AWS::ECS::EC2"
	OriginECSFargate = "AWS::ECS::Fargate"
	OriginEB         = "AWS::ElasticBeanstalk::Environment"
	OriginEKS        = "AWS::EKS::Container"
	OriginAppRunner  = "AWS::AppRunner::Service"
)

// x-ray only span attributes - https://github.com/open-telemetry/opentelemetry-java-contrib/pull/802
const (
	awsLocalService    = "aws.local.service"
	awsRemoteService   = "aws.remote.service"
	awsLocalOperation  = "aws.local.operation"
	awsRemoteOperation = "aws.remote.operation"
	remoteTarget       = "remoteTarget"
	awsSpanKind        = "aws.span.kind"
	k8sRemoteNamespace = "K8s.RemoteNamespace"
)

// reInvalidSpanCharacters defines the invalid letters in a span name as per
// Allowed characters for X-Ray Segment Name:
// Unicode letters, numbers, and whitespace, and the following symbols: _, ., :, /, %, &, #, =, +, \, -, @
// Doc: https://docs.aws.amazon.com/xray/latest/devguide/xray-api-segmentdocuments.html
var reInvalidSpanCharacters = regexp.MustCompile(`[^ 0-9\p{L}N_.:/%&#=+\-@]`)

var remoteXrayExporterDotConverter = featuregate.GlobalRegistry().MustRegister(
	"exporter.xray.allowDot",
	featuregate.StageBeta,
	featuregate.WithRegisterDescription("X-Ray Exporter will no longer convert . to _ in annotation keys when this feature gate is enabled. "),
	featuregate.WithRegisterFromVersion("v0.97.0"),
)

const (
	// defaultMetadataNamespace is used for non-namespaced non-indexed attributes.
	defaultMetadataNamespace = "default"
	// defaultSpanName will be used if there are no valid xray characters in the span name
	defaultSegmentName = "span"
	// maxSegmentNameLength the maximum length of a Segment name
	maxSegmentNameLength = 200
	// rpc.system value for AWS service remotes
	awsAPIRPCSystem = "aws-api"
)

const (
	traceIDLength    = 35 // fixed length of aws trace id
	identifierOffset = 11 // offset of identifier within traceID
)

const (
	localRoot = "LOCAL_ROOT"
)

var removeAnnotationsFromServiceSegment = []string{
	awsRemoteService,
	awsRemoteOperation,
	remoteTarget,
	k8sRemoteNamespace,
}

var writers = newWriterPool(2048)

// MakeSegmentDocuments converts spans to json documents
func MakeSegmentDocuments(span ptrace.Span, resource pcommon.Resource, indexedAttrs []string, indexAllAttrs bool, logGroupNames []string, skipTimestampValidation bool) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func isLocalRootSpanADependencySpan(span ptrace.Span) bool { _ = "STUB: not implemented"; return false }

// isLocalRoot - we will move to using isRemote once the collector supports deserializing it. Until then, we will rely on aws.span.kind.
func isLocalRoot(span ptrace.Span) bool { _ = "STUB: not implemented"; return false }

func addNamespaceToSubsegmentWithRemoteService(span ptrace.Span, segment *awsxray.Segment) {
	_ = "STUB: not implemented"
	return
}

func MakeDependencySubsegmentForLocalRootDependencySpan(span ptrace.Span, resource pcommon.Resource, indexedAttrs []string, indexAllAttrs bool, logGroupNames []string, skipTimestampValidation bool, serviceSegmentID pcommon.SpanID) (*awsxray.Segment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Make this a subsegment

// Remove span links from consumer spans

func MakeServiceSegmentForLocalRootDependencySpan(span ptrace.Span, resource pcommon.Resource, indexedAttrs []string, indexAllAttrs bool, logGroupNames []string, skipTimestampValidation bool, serviceSegmentID pcommon.SpanID) (*awsxray.Segment, error) {
	_ = "STUB: not implemented"
	// We always create a segment for the service
	return nil, nil
}

// Set the span id to the one internally generated

// Set the name

// Remove the HTTP field

// Remove AWS subsegment fields

// Delete all metadata that does not start with 'otel.resource.'

// Make it a segment

// Remote namespace

// Remove span links from non-consumer spans

func MakeServiceSegmentForLocalRootSpanWithoutDependency(span ptrace.Span, resource pcommon.Resource, indexedAttrs []string, indexAllAttrs bool, logGroupNames []string, skipTimestampValidation bool) ([]*awsxray.Segment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func MakeNonLocalRootSegment(span ptrace.Span, resource pcommon.Resource, indexedAttrs []string, indexAllAttrs bool, logGroupNames []string, skipTimestampValidation bool) ([]*awsxray.Segment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func MakeServiceSegmentAndDependencySubsegment(span ptrace.Span, resource pcommon.Resource, indexedAttrs []string, indexAllAttrs bool, logGroupNames []string, skipTimestampValidation bool) ([]*awsxray.Segment, error) {
	_ = "STUB: not implemented"
	// If it is a local root span and a dependency span, we need to make a segment and subsegment representing the local service and remote service, respectively.
	return nil, nil
}

// Make Dependency Subsegment

// Make Service Segment

// MakeSegmentsFromSpan creates one or more segments from a span
func MakeSegmentsFromSpan(span ptrace.Span, resource pcommon.Resource, indexedAttrs []string, indexAllAttrs bool, logGroupNames []string, skipTimestampValidation bool) ([]*awsxray.Segment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MakeSegmentDocumentString converts an OpenTelemetry Span to an X-Ray Segment and then serializes to JSON
// MakeSegmentDocumentString will be deprecated in the future
func MakeSegmentDocumentString(span ptrace.Span, resource pcommon.Resource, indexedAttrs []string, indexAllAttrs bool, logGroupNames []string, skipTimestampValidation bool) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// MakeDocumentFromSegment converts a segment into a JSON document
func MakeDocumentFromSegment(segment *awsxray.Segment) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func isAwsSdkSpan(span ptrace.Span) bool { _ = "STUB: not implemented"; return false }

// MakeSegment converts an OpenTelemetry Span to an X-Ray Segment
func MakeSegment(span ptrace.Span, resource pcommon.Resource, indexedAttrs []string, indexAllAttrs bool, logGroupNames []string, skipTimestampValidation bool) (*awsxray.Segment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// We only store the resource information for segments, the local root.

// convert trace id

// X-Ray segment names are service names, unlike span names which are methods. Try to find a service name.

// support x-ray specific service name attributes as segment name if it exists

// only strip the prefix for AWS spans

// peer.service should always be prioritized for segment names when it set by users and
// the new x-ray specific service name attributes are not found

// Generally spans are named something like "Method" or "Service.Method" but for AWS spans, X-Ray expects spans
// to be named "Service"

// For database queries, the segment name convention is <db name>@<db host>

// Trim JDBC connection string if starts with "jdbc:", otherwise no change
// jdbc:mysql://db.dev.example.com:3306

// Only for a server span, we can use the resource.

// newSegmentID generates a new valid X-Ray SegmentID
func newSegmentID() pcommon.SpanID { _ = "STUB: not implemented"; return *new(pcommon.SpanID) }

func determineAwsOrigin(resource pcommon.Resource) string { _ = "STUB: not implemented"; return "" }

// If cloud_platform is defined with a non-AWS value, we should not assign it an AWS origin

// convertToAmazonTraceID converts a trace ID to the Amazon format.
//
// A trace ID unique identifier that connects all segments and subsegments
// originating from a single client request.
//   - A trace_id consists of three numbers separated by hyphens. For example,
//     1-58406520-a006649127e371903a2de979. This includes:
//   - The version number, that is, 1.
//   - The time of the original request, in Unix epoch time, in 8 hexadecimal digits.
//   - For example, 10:00AM December 2nd, 2016 PST in epoch time is 1480615200 seconds,
//     or 58406520 in hexadecimal.
//   - A 96-bit identifier for the trace, globally unique, in 24 hexadecimal digits.
func convertToAmazonTraceID(traceID pcommon.TraceID, skipTimestampValidation bool) (string, error) {
	_ = "STUB: not implemented"

	// maxAge of 28 days.  AWS has a 30 day limit, let's be conservative rather than
	// hit the limit
	return "", nil
}

// maxSkew allows for 5m of clock skew

// If feature gate is enabled, skip the timestamp validation logic

// If AWS traceID originally came from AWS, no problem.  However, if oc generated
// the traceID, then the epoch may be outside the accepted AWS range of within the
// past 30 days.
//
// In that case, we return invalid traceid error

// Build the X-Ray trace ID format: 1-{hex(epoch)}-{identifier}
// Ensure we have enough space in the content array

// overwrite with identifier

func timestampToFloatSeconds(ts pcommon.Timestamp) float64 { _ = "STUB: not implemented"; return 0 }

func addSpecialAttributes(attributes map[string]pcommon.Value, indexedAttrs []string, unfilteredAttributes pcommon.Map) map[string]pcommon.Value {
	_ = "STUB: not implemented"
	return nil
}

// Allow attributes that have been filtered out before if explicitly added to be annotated/indexed

func makeXRayAttributes(attributes map[string]pcommon.Value, resource pcommon.Resource, storeResource bool, indexedAttrs []string, indexAllAttrs bool) (
	string, map[string]any, map[string]map[string]any,
) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

// if unable to unmarshal, keep the original key/value

func annotationValue(value pcommon.Value) any { _ = "STUB: not implemented"; return *new(any) }

// fixSegmentName removes any invalid characters from the span name.  AWS X-Ray defines
// the list of valid characters here:
// https://docs.aws.amazon.com/xray/latest/devguide/xray-api-segmentdocuments.html
func fixSegmentName(name string) string { _ = "STUB: not implemented"; return "" }

// only allocate for ReplaceAllString if we need to

// fixAnnotationKey removes any invalid characters from the annotation key.  AWS X-Ray defines
// the list of valid characters here:
// https://docs.aws.amazon.com/xray/latest/devguide/xray-api-segmentdocuments.html
func fixAnnotationKey(key string) string { _ = "STUB: not implemented"; return "" }

func makeEndTimeAndInProgress(span ptrace.Span, attributes pcommon.Map) (*float64, *bool) {
	_ = "STUB: not implemented"
	return nil, nil
}

func trimAwsSdkPrefix(name string, span ptrace.Span) string { _ = "STUB: not implemented"; return "" }

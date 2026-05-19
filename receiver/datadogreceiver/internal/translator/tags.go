// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package translator // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/datadogreceiver/internal/translator"

import (
	"sync"

	"go.opentelemetry.io/collector/featuregate"
	"go.opentelemetry.io/collector/pdata/pcommon"
	conventionsv138 "go.opentelemetry.io/otel/semconv/v1.38.0"
	conventions "go.opentelemetry.io/otel/semconv/v1.40.0"
)

var MultiTagParsingFeatureGate = featuregate.GlobalRegistry().MustRegister(
	"receiver.datadogreceiver.EnableMultiTagParsing",
	featuregate.StageAlpha,
	featuregate.WithRegisterDescription("When enabled, parses `key:value` tags with duplicate keys into a slice attribute."),
	featuregate.WithRegisterFromVersion("v0.142.0"),
	featuregate.WithRegisterReferenceURL("https://github.com/open-telemetry/opentelemetry-collector-contrib/pull/44747"),
)

// See:
// https://docs.datadoghq.com/opentelemetry/schema_semantics/semantic_mapping/
// https://github.com/DataDog/datadog-agent/blob/main/pkg/opentelemetry-mapping-go/otlp/attributes/attributes.go
var datadogKnownResourceAttributes = map[string]string{
	"env":     string(conventions.DeploymentEnvironmentNameKey),
	"service": string(conventions.ServiceNameKey),
	"version": string(conventions.ServiceVersionKey),

	// Container-related attributes
	"container_id":   string(conventions.ContainerIDKey),
	"container_name": string(conventions.ContainerNameKey),
	"image_name":     string(conventions.ContainerImageNameKey),
	"image_tag":      string(conventions.ContainerImageTagsKey),
	"runtime":        string(conventions.ContainerRuntimeNameKey),

	// Cloud-related attributes
	"cloud_provider": string(conventions.CloudProviderKey),
	"region":         string(conventions.CloudRegionKey),
	"zone":           string(conventions.CloudAvailabilityZoneKey),

	// ECS-related attributes
	"task_family":        string(conventions.AWSECSTaskFamilyKey),
	"task_arn":           string(conventions.AWSECSTaskARNKey),
	"ecs_cluster_name":   string(conventions.AWSECSClusterARNKey),
	"task_version":       string(conventions.AWSECSTaskRevisionKey),
	"ecs_container_name": string(conventions.AWSECSContainerARNKey),

	// K8-related attributes
	"kube_container_name": string(conventions.K8SContainerNameKey),
	"kube_cluster_name":   string(conventions.K8SClusterNameKey),
	"kube_deployment":     string(conventions.K8SDeploymentNameKey),
	"kube_replica_set":    string(conventions.K8SReplicaSetNameKey),
	"kube_stateful_set":   string(conventions.K8SStatefulSetNameKey),
	"kube_daemon_set":     string(conventions.K8SDaemonSetNameKey),
	"kube_job":            string(conventions.K8SJobNameKey),
	"kube_cronjob":        string(conventions.K8SCronJobNameKey),
	"kube_namespace":      string(conventions.K8SNamespaceNameKey),
	"pod_name":            string(conventions.K8SPodNameKey),

	// HTTP
	"http.client_ip":               string(conventions.ClientAddressKey),
	"http.response.content_length": string(conventions.HTTPResponseBodySizeKey),
	"http.status_code":             string(conventions.HTTPResponseStatusCodeKey),
	"http.request.content_length":  string(conventions.HTTPRequestBodySizeKey),
	"http.referer":                 "http.request.header.referer",
	"http.method":                  string(conventions.HTTPRequestMethodKey),
	"http.route":                   string(conventions.HTTPRouteKey),
	"http.version":                 string(conventions.NetworkProtocolVersionKey),
	"http.server_name":             string(conventions.ServerAddressKey),
	"http.url":                     string(conventions.URLFullKey),
	"http.useragent":               string(conventions.UserAgentOriginalKey),

	// AWS S3
	"aws.s3.bucket_name":      string(conventions.AWSS3BucketKey),
	"aws.response.request_id": string(conventions.AWSRequestIDKey),
	"aws.service":             string(conventionsv138.RPCServiceKey),
	"aws.operation":           string(conventions.RPCMethodKey),

	// DB
	"db.type":      string(conventions.DBSystemNameKey),
	"db.operation": string(conventions.DBOperationNameKey),
	"db.instance":  string(conventions.DBNamespaceKey),
	"db.sql.table": string(conventions.DBCollectionNameKey),
	"db.pool.name": string(conventions.DBClientConnectionPoolNameKey),
	"db.statement": string(conventions.DBQueryTextKey),

	// Other
	"process_id":       string(conventions.ProcessPIDKey),
	"error.stacktrace": string(conventions.ExceptionStacktraceKey),
	"error.msg":        string(conventions.ExceptionMessageKey),
}

// translateDatadogTagToKeyValuePair translates a Datadog tag to a key value pair
func translateDatadogTagToKeyValuePair(tag string) (key, value string) {
	_ = "STUB: not implemented"
	return "", ""
}

// Datadog allows for two tag formats, one of which includes a key such as 'env',
// followed by a value. Datadog also supports inputTags without the key, but OTel seems
// to only support key:value pairs.
// The following is a workaround to map unnamed inputTags to key:value pairs and its subject to future
// changes if OTel supports unnamed inputTags in the future or if there is a better way to do this.

// translateDatadogKeyToOTel translates a Datadog key to an OTel key
func translateDatadogKeyToOTel(k string) string { _ = "STUB: not implemented"; return "" }

// HTTP dynamic attributes
// type: string[]

// type: string[]

type StringPool struct {
	sync.RWMutex
	pool map[string]string
}

func newStringPool() *StringPool { _ = "STUB: not implemented"; return nil }

func (s *StringPool) Intern(str string) string { _ = "STUB: not implemented"; return "" }

// Double check if another goroutine has added the string after releasing the read lock

type attributes struct {
	resource pcommon.Map
	scope    pcommon.Map
	dp       pcommon.Map
}

func tagsToAttributes(tags []string, host string, stringPool *StringPool) attributes {
	_ = "STUB: not implemented"
	return *new(attributes)
}

// No need to intern the key if we already have it
// type: string[]

// type string[]

// Datadog does, semantically, generate tags with the same key prefix but different values
// (e.g. the `kube_service` tag when using the `kubelet` integration)
// (https://docs.datadoghq.com/containers/kubernetes/tag/)
// and we handle this by using a slice whenever there is more than one value for the same key

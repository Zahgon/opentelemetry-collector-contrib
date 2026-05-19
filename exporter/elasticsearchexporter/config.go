// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package elasticsearchexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/elasticsearchexporter"

import (
	"errors"
	"net/url"
	"time"

	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/config/configopaque"
	"go.opentelemetry.io/collector/config/configoptional"
	"go.opentelemetry.io/collector/confmap"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
	"go.uber.org/zap"
)

// Config defines configuration for Elastic exporter.
type Config struct {
	// QueueBatchConfig configures the sending queue and the batching done
	// by the exporter. The performed batching can further be customized by
	// configuring `metadata_keys` which will be used to partition the batches.
	QueueBatchConfig configoptional.Optional[exporterhelper.QueueBatchConfig] `mapstructure:"sending_queue"`

	// Endpoints holds the Elasticsearch URLs the exporter should send events to.
	//
	// This setting is required if CloudID is not set and if the
	// ELASTICSEARCH_URL environment variable is not set.
	Endpoints []string `mapstructure:"endpoints"`

	// CloudID holds the cloud ID to identify the Elastic Cloud cluster to send events to.
	// https://www.elastic.co/guide/en/cloud/current/ec-cloud-id.html
	//
	// This setting is required if no URL is configured.
	CloudID string `mapstructure:"cloudid"`

	// NumWorkers configures the number of workers publishing bulk requests.
	//
	// Deprecated: [v0.136.0] This config is now deprecated. Use `sending_queue::num_consumers`
	// instead. If this config is defined and `sending_queue::num_consumers` is not defined then
	// it will be used to set `sending_queue::num_consumers`.
	NumWorkers int `mapstructure:"num_workers"`

	// LogsIndex configures the static index used for document routing for logs.
	// It should be empty if dynamic document routing is preferred.
	LogsIndex        string              `mapstructure:"logs_index"`
	LogsDynamicIndex DynamicIndexSetting `mapstructure:"logs_dynamic_index"`

	// MetricsIndex configures the static index used for document routing for metrics.
	// It should be empty if dynamic document routing is preferred.
	MetricsIndex        string              `mapstructure:"metrics_index"`
	MetricsDynamicIndex DynamicIndexSetting `mapstructure:"metrics_dynamic_index"`

	// TracesIndex configures the static index used for document routing for metrics.
	// It should be empty if dynamic document routing is preferred.
	TracesIndex        string              `mapstructure:"traces_index"`
	TracesDynamicIndex DynamicIndexSetting `mapstructure:"traces_dynamic_index"`

	// LogsDynamicID configures whether log record attribute `elasticsearch.document_id` is set as the document ID in ES.
	LogsDynamicID DynamicIDSettings `mapstructure:"logs_dynamic_id"`

	// TracesDynamicID configures whether span attribute `elasticsearch.document_id` is set as the document ID in ES.
	TracesDynamicID DynamicIDSettings `mapstructure:"traces_dynamic_id"`

	// LogsDynamicPipeline configures whether log record attribute `elasticsearch.document_pipeline` is set as the document ingest pipeline for ES.
	LogsDynamicPipeline DynamicPipelineSettings `mapstructure:"logs_dynamic_pipeline"`

	// Pipeline configures the ingest node pipeline name that should be used to process the
	// events.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/ingest.html
	Pipeline string `mapstructure:"pipeline"`

	confighttp.ClientConfig `mapstructure:",squash"`
	Authentication          AuthenticationSettings `mapstructure:",squash"`
	Discovery               DiscoverySettings      `mapstructure:"discover"`
	Retry                   RetrySettings          `mapstructure:"retry"`

	// Deprecated: [v0.136.0] This config is now deprecated. Use `sending_queue::batch` instead.
	// If this config is defined then it will be used to configure sending queue's batch provided
	// sending queue's config are not explicitly defined.
	Flush          FlushSettings          `mapstructure:"flush"`
	Mapping        MappingsSettings       `mapstructure:"mapping"`
	LogstashFormat LogstashFormatSettings `mapstructure:"logstash_format"`

	// SuppressConflictErrors configures whether 409 Conflict responses are logged as errors.
	// If set to true, document level version conflict exceptions (409) will not be logged.
	SuppressConflictErrors bool `mapstructure:"suppress_conflict_errors"`

	// TelemetrySettings contains settings useful for testing/debugging purposes.
	// This is experimental and may change at any time.
	TelemetrySettings `mapstructure:"telemetry"`

	// IncludeSourceOnError configures whether the bulk index responses include
	// a part of the source document on error.
	// Defaults to nil.
	//
	// This setting requires Elasticsearch 8.18+. Using it in prior versions
	// have no effect.
	//
	// NOTE: The default behavior if this configuration is not set, is to
	// discard the error reason entirely, i.e. only the error type is returned.
	//
	// WARNING: If set to true, the exporter may log error responses containing
	// request payload, causing potential sensitive data to be exposed in logs.
	// Users are expected to sanitize the responses themselves.
	IncludeSourceOnError *bool `mapstructure:"include_source_on_error"`

	// Experimental: MetadataKeys defines a list of client.Metadata keys that
	// will be used as partition keys for when batcher is enabled and will be
	// added to the exporter's telemetry if defined. The config only applies
	// when `sending_queue::batch` is defined or when the, now deprecated, batcher
	// is used (set to `true` or `false`). The metadata keys are converted to
	// lower case as key lookups for client metadata is case insensitive. This
	// means that the metric produced by internal telemetry will also have the
	// attribute in lower case.
	//
	// Keys are case-insensitive and duplicates will trigger a validation error.
	MetadataKeys []string `mapstructure:"metadata_keys"`

	// BulkResponseFilterPath sets the filter_path parameter of bulk API requests,
	// which controls what data is returned in the response from Elasticsearch.
	//
	// Note: If `items.*._index.items` is not in the BulkResponseFilterPath
	// than for any failed documents, the exporter will not be able
	// to log the index to which the document was being written
	// to.
	//
	// Note: if `items.*._index.items` is not in the BulkResponseFilterPath
	// than the export will log rejection of duplicates to
	// ".profiling-stackframes" which were previously suppressed.
	//
	// BulkResponseFilterPath defaults to
	// "items.*._index,items.*.status,items.*.failure_store,items.*.error.type,items.*.error.reason"
	BulkResponseFilterPath string `mapstructure:"bulk_response_filter_path"`
}

type TelemetrySettings struct {
	LogRequestBody  bool `mapstructure:"log_request_body"`
	LogResponseBody bool `mapstructure:"log_response_body"`

	LogFailedDocsInput          bool          `mapstructure:"log_failed_docs_input"`
	LogFailedDocsInputRateLimit time.Duration `mapstructure:"log_failed_docs_input_rate_limit"`

	// prevent unkeyed literal initialization
	_ struct{}
}

type LogstashFormatSettings struct {
	Enabled         bool   `mapstructure:"enabled"`
	PrefixSeparator string `mapstructure:"prefix_separator"`
	DateFormat      string `mapstructure:"date_format"`

	// prevent unkeyed literal initialization
	_ struct{}
}

type DynamicIndexSetting struct {
	// Enabled enables dynamic index routing.
	//
	// Deprecated: [v0.122.0] This config is now ignored. Dynamic index routing is always done by default.
	Enabled bool `mapstructure:"enabled"`

	// prevent unkeyed literal initialization
	_ struct{}
}

type DynamicIDSettings struct {
	Enabled bool `mapstructure:"enabled"`

	// prevent unkeyed literal initialization
	_ struct{}
}

type DynamicPipelineSettings struct {
	Enabled bool `mapstructure:"enabled"`

	// prevent unkeyed literal initialization
	_ struct{}
}

// AuthenticationSettings defines user authentication related settings.
type AuthenticationSettings struct {
	// User is used to configure HTTP Basic Authentication.
	User string `mapstructure:"user"`

	// Password is used to configure HTTP Basic Authentication.
	Password configopaque.String `mapstructure:"password"`

	// APIKey is used to configure ApiKey based Authentication.
	//
	// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-create-api-key.html
	APIKey configopaque.String `mapstructure:"api_key"`

	// prevent unkeyed literal initialization
	_ struct{}
}

// DiscoverySettings defines Elasticsearch node discovery related settings.
// The exporter will check Elasticsearch regularly for available nodes
// and updates the list of hosts if discovery is enabled. Newly discovered
// nodes will automatically be used for load balancing.
//
// DiscoverySettings should not be enabled when operating Elasticsearch behind a proxy
// or load balancer.
//
// https://www.elastic.co/blog/elasticsearch-sniffing-best-practices-what-when-why-how
type DiscoverySettings struct {
	// OnStart, if set, instructs the exporter to look for available Elasticsearch
	// nodes the first time the exporter connects to the cluster.
	OnStart bool `mapstructure:"on_start"`

	// Interval instructs the exporter to renew the list of Elasticsearch URLs
	// with the given interval. URLs will not be updated if Interval is <=0.
	Interval time.Duration `mapstructure:"interval"`

	// prevent unkeyed literal initialization
	_ struct{}
}

// FlushSettings defines settings for configuring the write buffer flushing
// policy in the Elasticsearch exporter. The exporter sends a bulk request with
// all events already serialized into the send-buffer.
//
// Deprecated: [v0.136.0] This config is now deprecated. Use `sending_queue::batch` instead.
// If this config is defined then it will be used to configure sending queue's batch provided
// sending queue's config are not explicitly defined.
type FlushSettings struct {
	// Bytes sets the send buffer flushing limit.
	// Bytes is now deprecated. Use `sending_queue::batch::{min, max}_size` with `bytes`
	// sizer to configure batching based on bytes.
	//
	// If this config option is defined then it will be used to configure `sending_queue::batch::max_size`
	// provided it is not explcitly defined.
	Bytes int `mapstructure:"bytes"`

	// Interval configures the max age of a document in the send buffer.
	// Interval is now deprecated. Use `sending-queue::batch::flush_timeout` instead.
	//
	// If this config option is defined then it will be used to configure `sending_queue::batch::flush_timeout`
	// provided it is not explcitly defined.
	Interval time.Duration `mapstructure:"interval"`

	// prevent unkeyed literal initialization
	_ struct{}
}

// RetrySettings defines settings for the HTTP request retries in the Elasticsearch exporter.
// Failed sends are retried with exponential backoff.
type RetrySettings struct {
	// Enabled allows users to disable retry without having to comment out all settings.
	Enabled bool `mapstructure:"enabled"`

	// MaxRequests configures how often an HTTP request is attempted before it is assumed to be failed.
	//
	// Deprecated: use MaxRetries instead.
	MaxRequests int `mapstructure:"max_requests"`

	// MaxRetries configures how many times an HTTP request is retried.
	MaxRetries int `mapstructure:"max_retries"`

	// InitialInterval configures the initial waiting time if a request failed.
	InitialInterval time.Duration `mapstructure:"initial_interval"`

	// MaxInterval configures the max waiting time if consecutive requests failed.
	MaxInterval time.Duration `mapstructure:"max_interval"`

	// RetryOnStatus configures the status codes that trigger request or document level retries.
	RetryOnStatus []int `mapstructure:"retry_on_status"`
}

type MappingsSettings struct {
	// Deprecated: [v0.145.0] Mode is ignored. The default mapping mode is "otel".
	//
	// The mode may be overridden in two ways:
	//  - by the client metadata key X-Elastic-Mapping-Mode, if specified
	//  - by the scope attribute elastic.mapping.mode, if specified
	//
	// The order of precedence is:
	//   scope attribute > client metadata > default mode.
	Mode string `mapstructure:"mode"`

	// AllowedModes controls the allowed document mapping modes
	// specified through X-Elastic-Mapping-Mode client metadata.
	//
	// If unspecified, all mapping modes are allowed.
	AllowedModes []string `mapstructure:"allowed_modes"`

	// prevent unkeyed literal initialization
	_ struct{}
}

type MappingMode int

// Enum values for MappingMode.
const (
	MappingNone MappingMode = iota
	MappingECS
	MappingOTel
	MappingRaw
	MappingBodyMap

	// NumMappingModes remain last, it is used for sizing arrays.
	NumMappingModes
)

func (m MappingMode) String() string { _ = "STUB: not implemented"; return "" }

var (
	errConfigEndpointRequired = errors.New("exactly one of [endpoint, endpoints, cloudid] must be specified")
	errConfigEmptyEndpoint    = errors.New("endpoint must not be empty")
)

const defaultElasticsearchEnvName = "ELASTICSEARCH_URL"

func (cfg *Config) Unmarshal(conf *confmap.Conf) error { _ = "STUB: not implemented"; return nil }

// Validate validates the elasticsearch server configuration.
func (cfg *Config) Validate() error { _ = "STUB: not implemented"; return nil }

// convert metadata keys to lower case as these are case insensitive

// allowedMappingModes returns a map from canonical mapping mode names to MappingModes.
func (cfg *Config) allowedMappingModes() map[string]MappingMode {
	_ = "STUB: not implemented"
	return nil
}

var canonicalMappingModes = map[string]MappingMode{
	MappingNone.String():    MappingNone,
	MappingRaw.String():     MappingRaw,
	MappingECS.String():     MappingECS,
	MappingOTel.String():    MappingOTel,
	MappingBodyMap.String(): MappingBodyMap,
}

func canonicalMappingModeName(name string) string { _ = "STUB: not implemented"; return "" }

// aliases for "none"

func (cfg *Config) endpoints() ([]string, error) {
	_ = "STUB: not implemented"
	// Exactly one of endpoint, endpoints, or cloudid must be configured.
	// If none are set, then $ELASTICSEARCH_URL may be specified instead.
	return nil, nil
}

func validateEndpoint(endpoint string) error { _ = "STUB: not implemented"; return nil }

// Based on "addrFromCloudID" in go-elasticsearch.
func parseCloudID(input string) (*url.URL, error) { _ = "STUB: not implemented"; return nil, nil }

func handleDeprecatedConfig(cfg *Config, logger *zap.Logger) { _ = "STUB: not implemented"; return }

// Do not set cfg.Retry.Enabled = false if cfg.Retry.MaxRequest = 1 to avoid breaking change on behavior

func handleTelemetryConfig(cfg *Config, logger *zap.Logger) { _ = "STUB: not implemented"; return }

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"errors"
	"time"

	"github.com/spf13/pflag"
	"go.opentelemetry.io/otel/attribute"

	types "github.com/open-telemetry/opentelemetry-collector-contrib/cmd/telemetrygen/pkg"
)

var (
	errFormatOTLPAttributes       = errors.New("value should be in one of the following formats: key=\"value\", key=true, key=false, key=<integer>, or key=[<value1>, <value2>, ...]")
	errDoubleQuotesOTLPAttributes = errors.New("value should be a string wrapped in double quotes")
	errMixedTypeSlice             = errors.New("all items in a slice should be of the same type")
	errEmptySlice                 = errors.New("slice should not be empty")
)

const (
	defaultGRPCEndpoint = "localhost:4317"
	defaultHTTPEndpoint = "localhost:4318"
)

type KeyValue map[string]any

var _ pflag.Value = (*KeyValue)(nil)

func (*KeyValue) String() string { _ = "STUB: not implemented"; return "" }

func parseValue(val string) (any, error) { _ = "STUB: not implemented"; return *new(any), nil }

// splitItems splits the pre-trimmed content into a list of items separated by commas, respecting quotes
func splitItems(content string) []string { _ = "STUB: not implemented"; return nil }

// Add the last item

// sliceFrom converts items into a slice of a single type
func sliceFrom(items []string) (any, error) { _ = "STUB: not implemented"; return *new(any), nil }

func (v *KeyValue) Set(s string) error { _ = "STUB: not implemented"; return nil }

// List of Values

func (*KeyValue) Type() string { _ = "STUB: not implemented"; return "" }

type Config struct {
	WorkerCount           int
	Rate                  float64
	TotalDuration         types.DurationWithInf
	ReportingInterval     time.Duration
	SkipSettingGRPCLogger bool
	Timeout               time.Duration

	// OTLP config
	CustomEndpoint      string
	Insecure            bool
	InsecureSkipVerify  bool
	UseHTTP             bool
	HTTPPath            string
	Headers             KeyValue
	ResourceAttributes  KeyValue
	ServiceName         string
	TelemetryAttributes KeyValue

	// OTLP TLS configuration
	CaFile string

	// OTLP mTLS configuration
	ClientAuth ClientAuth

	// Export behavior configuration
	AllowExportFailures bool

	// Load testing configuration
	LoadSize int

	// Batching configuration
	Batch     bool
	BatchSize int
}

type ClientAuth struct {
	Enabled        bool
	ClientCertFile string
	ClientKeyFile  string
}

// Endpoint returns the appropriate endpoint URL based on the selected communication mode (gRPC or HTTP)
// or custom endpoint provided in the configuration.
func (c *Config) Endpoint() string { _ = "STUB: not implemented"; return "" }

func (c *Config) GetAttributes() []attribute.KeyValue { _ = "STUB: not implemented"; return nil }

// may be overridden by `--otlp-attributes service.name="foo"`

func (c *Config) GetTelemetryAttributes() []attribute.KeyValue {
	_ = "STUB: not implemented"
	return nil
}

func (c *Config) GetHeaders() map[string]string { _ = "STUB: not implemented"; return nil }

// CommonFlags registers common config flags.
func (c *Config) CommonFlags(fs *pflag.FlagSet) { _ = "STUB: not implemented"; return }

// custom headers

// custom resource attributes

// TLS CA configuration

// mTLS configuration

// Export behavior configuration

// Load testing configuration

// Batching configuration

// SetDefaults is here to mirror the defaults for flags above,
// This allows for us to have a single place to change the defaults
// while exposing the API for use.
func (c *Config) SetDefaults() { _ = "STUB: not implemented"; return }

// CharactersPerMB is the number of characters needed to create a 1MB string attribute
const CharactersPerMB = 1024 * 1024

// CreateLoadAttribute creates a string attribute with the specified size in MB
// This is commonly used across different signal types (metrics, traces, logs) for load testing
func CreateLoadAttribute(key string, sizeMB int) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

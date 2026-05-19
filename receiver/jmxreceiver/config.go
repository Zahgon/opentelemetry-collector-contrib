// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package jmxreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/jmxreceiver"

import (
	"go.opentelemetry.io/collector/config/configopaque"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
	"go.opentelemetry.io/collector/scraper/scraperhelper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// jmxGathererMainClass the class containing the main function for the JMX Metric Gatherer JAR
var jmxGathererMainClass = "io.opentelemetry.contrib.jmxmetrics.JmxMetrics"

// jmxScraperMainClass the class containing the main function for the JMX Scraper JAR
var jmxScraperMainClass = "io.opentelemetry.contrib.jmxscraper.JmxScraper"

type Config struct {
	scraperhelper.ControllerConfig `mapstructure:",squash"`

	// The path for the JMX Metric Gatherer or JMX Scraper JAR (/opt/opentelemetry-java-contrib-jmx-metrics.jar by default).
	// Supported by: jmx-scraper and jmx-metric-gatherer
	JARPath string `mapstructure:"jar_path"`
	// The Service URL or host:port for the target coerced to one of form: service:jmx:rmi:///jndi/rmi://<host>:<port>/jmxrmi.
	// Supported by: jmx-scraper and jmx-metric-gatherer
	Endpoint string `mapstructure:"endpoint"`
	// Comma-separated list of systems to monitor
	// Supported by: jmx-scraper and jmx-metric-gatherer
	TargetSystem string `mapstructure:"target_system"`
	// The target source of metric definitions to use for the target system.
	// Supported values are: auto, instrumentation and legacy.
	// Supported by: jmx-scraper
	TargetSource string `mapstructure:"target_source"`
	// Comma-separated list of paths to custom YAML metrics definition,
	// mandatory when TargetSystem is not set.
	// Supported by: jmx-scraper
	JmxConfigs string `mapstructure:"jmx_configs"`
	// The OTLP exporter settings
	// Supported by: jmx-scraper and jmx-metric-gatherer
	OTLPExporterConfig otlpExporterConfig `mapstructure:"otlp"`
	// The JMX username
	// Supported by: jmx-scraper and jmx-metric-gatherer
	Username string `mapstructure:"username"`
	// The JMX password
	// Supported by: jmx-scraper and jmx-metric-gatherer
	Password configopaque.String `mapstructure:"password"`
	// The keystore path for SSL
	// Supported by: jmx-scraper and jmx-metric-gatherer
	KeystorePath string `mapstructure:"keystore_path"`
	// The keystore password for SSL
	// Supported by: jmx-scraper and jmx-metric-gatherer
	KeystorePassword configopaque.String `mapstructure:"keystore_password"`
	// The keystore type for SSL
	// Supported by: jmx-scraper and jmx-metric-gatherer
	KeystoreType string `mapstructure:"keystore_type"`
	// The truststore path for SSL
	// Supported by: jmx-scraper and jmx-metric-gatherer
	TruststorePath string `mapstructure:"truststore_path"`
	// The truststore password for SSL
	// Supported by: jmx-scraper and jmx-metric-gatherer
	TruststorePassword configopaque.String `mapstructure:"truststore_password"`
	// The truststore type for SSL
	// Supported by: jmx-scraper and jmx-metric-gatherer
	TruststoreType string `mapstructure:"truststore_type"`
	// The JMX remote profile.  Should be one of:
	// `"SASL/PLAIN"`, `"SASL/DIGEST-MD5"`, `"SASL/CRAM-MD5"`, `"TLS SASL/PLAIN"`, `"TLS SASL/DIGEST-MD5"`, or
	// `"TLS SASL/CRAM-MD5"`, though no enforcement is applied.
	// Supported by: jmx-scraper and jmx-metric-gatherer
	RemoteProfile string `mapstructure:"remote_profile"`
	// The SASL/DIGEST-MD5 realm
	// Supported by: jmx-scraper and jmx-metric-gatherer
	Realm string `mapstructure:"realm"`
	// Array of additional JARs to be added to the class path when launching the JMX Metric Gatherer JAR
	// Supported by: jmx-scraper and jmx-metric-gatherer
	AdditionalJars []string `mapstructure:"additional_jars"`
	// Map of resource attributes used by the Java SDK Autoconfigure to set resource attributes
	// Supported by: jmx-scraper and jmx-metric-gatherer
	ResourceAttributes map[string]string `mapstructure:"resource_attributes"`
	// Log level used by the JMX metric gatherer. Should be one of:
	// `"trace"`, `"debug"`, `"info"`, `"warn"`, `"error"`, `"off"`
	// Supported by: jmx-metric-gatherer
	LogLevel string `mapstructure:"log_level"`
}

// We don't embed the existing OTLP Exporter config as most fields are unsupported
type otlpExporterConfig struct {
	// The OTLP Receiver endpoint to send metrics to ("0.0.0.0:<random open port>" by default).
	Endpoint string `mapstructure:"endpoint"`
	// The OTLP exporter timeout (5 seconds by default).  Will be converted to milliseconds.
	TimeoutSettings exporterhelper.TimeoutConfig `mapstructure:",squash"`
	// The headers to include in OTLP metric submission requests.
	Headers map[string]string `mapstructure:"headers"`
}

func (oec otlpExporterConfig) headersToString() string {
	_ = "STUB: not implemented"
	// sort for reliable testing
	return ""
}

// remove trailing comma

func (c *Config) parseProperties(logger *zap.Logger) []string {
	_ = "STUB: not implemented"
	// slf4j.simpleLogger only available in JMX Metrics Gatherer jar
	return nil
}

// Sorted for testing and reproducibility

var logLevelTranslator = map[zapcore.Level]string{
	zap.DebugLevel:  "debug",
	zap.InfoLevel:   "info",
	zap.WarnLevel:   "warn",
	zap.ErrorLevel:  "error",
	zap.DPanicLevel: "error",
	zap.PanicLevel:  "error",
	zap.FatalLevel:  "error",
}

var zapLevels = []zapcore.Level{
	zap.DebugLevel,
	zap.InfoLevel,
	zap.WarnLevel,
	zap.ErrorLevel,
	zap.DPanicLevel,
	zap.PanicLevel,
	zap.FatalLevel,
}

func getZapLoggerLevelEquivalent(logger *zap.Logger) string { _ = "STUB: not implemented"; return "" }

// Couldn't get log level from logger default logger level to info

func testLevel(logger *zap.Logger, level zapcore.Level) bool {
	_ = "STUB: not implemented"
	return false
}

// parseClasspath creates a classpath string with the JMX Gatherer JAR at the beginning
func (c *Config) parseClasspath() string { _ = "STUB: not implemented"; return "" }

// Add JMX JAR to classpath

// Add additional JARs if any

// Join them

func isSupportedJAR(supportedJarDetails map[string]supportedJar, jar string) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *Config) jarMainClass() string { _ = "STUB: not implemented"; return "" }

func (c *Config) jarJMXSamplingConfig() (string, string) { _ = "STUB: not implemented"; return "", "" }

func hashFile(path string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func (c *Config) validateJar(supportedJarDetails map[string]supportedJar, jar string) error {
	_ = "STUB: not implemented"
	return nil
}

var (
	validLogLevels = map[string]struct{}{"trace": {}, "debug": {}, "info": {}, "warn": {}, "error": {}, "off": {}}
	// feature parity between jmx-gatherer and jmx-scraper
	validTargetSystems = map[string]struct{}{
		"activemq": {}, "cassandra": {}, "hbase": {}, "hadoop": {},
		"jetty": {}, "jvm": {}, "kafka": {}, "kafka-consumer": {}, "kafka-producer": {}, "solr": {}, "tomcat": {}, "wildfly": {},
	}
)
var AdditionalTargetSystems = "n/a"

// Separated into two functions for tests
func init() {
	initAdditionalTargetSystems()
}

func initAdditionalTargetSystems() { _ = "STUB: not implemented"; return }

func (c *Config) Validate() error { _ = "STUB: not implemented"; return nil }

// jmx-scraper can use jmx_configs instead

func listKeys(presenceMap map[string]struct{}) string { _ = "STUB: not implemented"; return "" }

func (c *Config) buildJMXConfig() (string, error) { _ = "STUB: not implemented"; return "", nil }

// set jmx-scraper specific config options

// jmx-scraper default target source: https://github.com/open-telemetry/opentelemetry-java-contrib/tree/main/jmx-scraper#configuration-reference

// Documentation of Java Properties format & escapes: https://docs.oracle.com/javase/7/docs/api/java/util/Properties.html#load(java.io.Reader)

// Keys are receiver-defined so this escape should be unnecessary but in case that assumption
// breaks in the future this will ensure keys are properly escaped

// Any whitespace must be removed from keys

// Unneeded escape tokens will be removed by the properties file loader, so it should be pre-escaped to ensure
// the values provided reach the metrics gatherer as provided. Also in case a user attempts to provide multiline
// values for one of the available fields, we need to escape the newlines

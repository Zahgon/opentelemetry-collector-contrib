// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package integrationtestutils // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/splunkhecexporter/internal/integrationtestutils"

var configFilePath = "./testdata/integration_tests_config.yaml"

type IntegrationTestsConfig struct {
	Host           string `yaml:"HOST"`
	User           string `yaml:"USER"`
	Password       string `yaml:"PASSWORD"`
	UIPort         string `yaml:"UI_PORT"`
	HecPort        string `yaml:"HEC_PORT"`
	ManagementPort string `yaml:"MANAGEMENT_PORT"`
	EventIndex     string `yaml:"EVENT_INDEX"`
	MetricIndex    string `yaml:"METRIC_INDEX"`
	TraceIndex     string `yaml:"TRACE_INDEX"`
	HecToken       string `yaml:"HEC_TOKEN"`
	SplunkImage    string `yaml:"SPLUNK_IMAGE"`
}

func GetConfigVariable(key string) string {
	_ = "STUB: not implemented"
	// Read YAML file
	return ""
}

func SetConfigVariable(key, value string) {
	_ = "STUB: not implemented"
	// Read YAML file
	return
}

// Marshal updated Config into YAML

// Write yaml file

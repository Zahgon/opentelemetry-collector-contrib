// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package tests // import "github.com/open-telemetry/opentelemetry-collector-contrib/testbed/tests"

// This file defines parametrized test scenarios and makes them public so that they can be
// also used by tests in custom builds of Collector (e.g. Collector Contrib).

import (
	"regexp"
	"testing"

	"go.opentelemetry.io/collector/pdata/plog"

	"github.com/open-telemetry/opentelemetry-collector-contrib/testbed/testbed"
)

var (
	batchRegex                                           = regexp.MustCompile(` batch_index=(\S+) `)
	itemRegex                                            = regexp.MustCompile(` item_index=(\S+) `)
	performanceResultsSummary testbed.TestResultsSummary = &testbed.PerformanceResults{}
)

type ProcessorNameAndConfigBody struct {
	Name string
	Body string
}

// createConfigYaml creates a collector config file that corresponds to the
// sender and receiver used in the test and returns the config file name.
// Map of processor names to their configs. Config is in YAML and must be
// indented by 2 spaces. Processors will be placed between batch and queue for traces
// pipeline. For metrics pipeline these will be sole processors.
func createConfigYaml(
	t *testing.T,
	sender testbed.DataSender,
	receiver testbed.DataReceiver,
	resultDir string,
	processors []ProcessorNameAndConfigBody,
	extensions map[string]string,
) string {
	_ = "STUB: not implemented"
	// Create a config. Note that our DataSender is used to generate a config for Collector's
	// receiver and our DataReceiver is used to generate a config for Collector's exporter.
	// This is because our DataSender sends to Collector's receiver and our DataReceiver
	// receives from Collector's exporter.
	return ""
}

// Prepare extra processor config section and comma-separated list of extra processor
// names to use in corresponding "processors" settings.

// Prepare extra extension config section and comma-separated list of extra extension
// names to use in corresponding "extensions" settings.

// Set pipeline based on DataSender type

// Put corresponding elements into the config template to generate the final config.

// Scenario10kItemsPerSecond runs 10k data items/sec test using specified sender and receiver protocols.
func Scenario10kItemsPerSecond(
	t *testing.T,
	sender testbed.DataSender,
	receiver testbed.DataReceiver,
	resourceSpec testbed.ResourceSpec,
	resultsSummary testbed.TestResultsSummary,
	processors []ProcessorNameAndConfigBody,
	extensions map[string]string,
	loadOptions *testbed.LoadOptions,
) {
	_ = "STUB: not implemented"
	return
}

// Scenario10kItemsPerSecondAlternateBackend runs 10k data items/sec test using specified sender and receiver protocols.
// The only difference from Scenario10kItemsPerSecond is that this method can be used to specify a different backend. This
// is useful when testing components for which there is no exporter that emits the same format as the receiver format.
func Scenario10kItemsPerSecondAlternateBackend(
	t *testing.T,
	sender testbed.DataSender,
	receiver testbed.DataReceiver,
	backend testbed.DataReceiver,
	resourceSpec testbed.ResourceSpec,
	resultsSummary testbed.TestResultsSummary,
	processors []ProcessorNameAndConfigBody,
	extensions map[string]string,
) {
	_ = "STUB: not implemented"
	return
}

// for some scenarios, the mockbackend isn't the same as the receiver
// therefore, the backend must be initialized with the correct receiver

// TestCase for Scenario1kSPSWithAttrs func.
type TestCase struct {
	attrCount      int
	attrSizeByte   int
	expectedMaxCPU uint32
	expectedMaxRAM uint32
	resultsSummary testbed.TestResultsSummary
}

func genRandByteString(length int) string { _ = "STUB: not implemented"; return "" }

// Scenario1kSPSWithAttrs runs a performance test at 1k sps with specified span attributes
// and test options.
func Scenario1kSPSWithAttrs(t *testing.T, args []string, tests []TestCase, processors []ProcessorNameAndConfigBody, extensions map[string]string) {
	_ = "STUB: not implemented"
	return
}

// Prepare results dir.

// Create sender and receiver on available ports.

// Prepare config.

// Structure used for TestTraceNoBackend10kSPS.
// Defines RAM usage range for defined processor type.
type processorConfig struct {
	Name string
	// slice of processor structs with their names and config YAML to use.
	Processor           []ProcessorNameAndConfigBody
	ExpectedMaxRAM      uint32
	ExpectedMinFinalRAM uint32
}

func ScenarioTestTraceNoBackend10kSPS(
	t *testing.T,
	sender testbed.DataSender,
	receiver testbed.DataReceiver,
	resourceSpec testbed.ResourceSpec,
	resultsSummary testbed.TestResultsSummary,
	configuration processorConfig,
) {
	_ = "STUB: not implemented"
	return
}

func ScenarioSendingQueuesFull(
	t *testing.T,
	sender testbed.DataSender,
	receiver testbed.DataReceiver,
	loadOptions testbed.LoadOptions,
	resourceSpec testbed.ResourceSpec,
	sleepTime int,
	resultsSummary testbed.TestResultsSummary,
	processors []ProcessorNameAndConfigBody,
	extensions map[string]string,
) {
	_ = "STUB: not implemented"
	return
}

// searchFunc checks for "sending queue is full" communicate and sends the signal to GenerateNonPernamentErrorUntil
// to generate only successes from that time on

// check if data started to be received successfully

// get IDs from logs to retry

// get IDs from logs received successfully

// check if all the logs to retry were actually retried

func ScenarioSendingQueuesNotFull(
	t *testing.T,
	sender testbed.DataSender,
	receiver testbed.DataReceiver,
	loadOptions testbed.LoadOptions,
	resourceSpec testbed.ResourceSpec,
	sleepTime int,
	resultsSummary testbed.TestResultsSummary,
	processors []ProcessorNameAndConfigBody,
	extensions map[string]string,
) {
	_ = "STUB: not implemented"
	return
}

func ScenarioLong(
	t *testing.T,
	sender testbed.DataSender,
	receiver testbed.DataReceiver,
	loadOptions testbed.LoadOptions,
	resultsSummary testbed.TestResultsSummary,
	sleepTime int,
	processors []ProcessorNameAndConfigBody,
) {
	_ = "STUB: not implemented"
	return
}

func ScenarioMemoryLimiterHit(
	t *testing.T,
	sender testbed.DataSender,
	receiver testbed.DataReceiver,
	loadOptions testbed.LoadOptions,
	resultsSummary testbed.TestResultsSummary,
	sleepTime int,
	processors []ProcessorNameAndConfigBody,
) {
	_ = "STUB: not implemented"
	return
}

// check for "Memory usage is above soft limit"

// Log found. But keep the collector under stress for 10 more seconds so it starts refusing data

// check if data started to be received successfully

// stop sending any more data

// get IDs from logs to retry

// get IDs from logs received successfully

// check if all the logs to retry were actually retried

func constructLoadOptions(test TestCase) testbed.LoadOptions {
	_ = "STUB: not implemented"
	return *new(testbed.LoadOptions)
}

// Generate attributes.

func getLogsID(logToRetry []plog.Logs) []string { _ = "STUB: not implemented"; return nil }

func allElementsExistInSlice(slice1, slice2 []string) bool {
	_ = "STUB: not implemented"
	// Create a map to store elements of slice2 for efficient lookup
	return false
}

// Populate the map with elements from slice2

// Check if all elements of slice1 exist in slice2

// in case of file_log receiver, the batch_index and item_index are a part of log body.
// we use regex to extract them
func extractIDFromLog(log plog.LogRecord) (string, string) {
	_ = "STUB: not implemented"
	return "", ""
}

// in case of otlp receiver, batch_index and item_index are part of attributes.

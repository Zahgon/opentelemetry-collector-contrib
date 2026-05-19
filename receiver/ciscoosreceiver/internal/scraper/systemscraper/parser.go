// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package systemscraper // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/ciscoosreceiver/internal/scraper/systemscraper"

// parseCPUUtilizationNXOS parses NX-OS CPU stats (fraction 0-1)
func parseCPUUtilizationNXOS(output string) (float64, error) {
	_ = "STUB: not implemented"
	// Regex to match: CPU states  :   27.70% user,   10.80% kernel,   61.49% idle
	return 0, nil
}

// CPU utilization = user + kernel (as fraction)

// parseCPUUtilizationIOS parses IOS CPU stats using 5-second average (fraction 0-1)
func parseCPUUtilizationIOS(output string) (float64, error) {
	_ = "STUB: not implemented"
	// Regex to match: CPU utilization for five seconds: 8%/2%; one minute: 7%; five minutes: 6%
	return 0, nil
}

// CPU utilization as fraction

// parseMemoryUtilization parses memory stats based on OS type (fraction 0-1)
func parseMemoryUtilization(output, osType string) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil

	// Select regex pattern based on OS type
}

// NX-OS: "Memory usage:   32803148K total,   11174924K used,   21628224K free"

// IOS/IOS XE: "Processor Pool Total:    1000000 Used:     600000 Free:     400000"

// Memory utilization as fraction

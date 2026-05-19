// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ucal // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/cpuscraper/ucal"

import (
	"errors"

	"github.com/shirou/gopsutil/v4/cpu"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

var ErrTimeStatNotFound = errors.New("cannot find TimesStat for cpu")

// CPUUtilization stores the utilization percents [0-1] for the different cpu states
type CPUUtilization struct {
	CPU     string
	User    float64
	System  float64
	Idle    float64
	Nice    float64
	Iowait  float64
	Irq     float64
	Softirq float64
	Steal   float64
}

// CPUUtilizationCalculator calculates the cpu utilization percents for the different cpu states
// It requires 2 []cpu.TimesStat and spend time to be able to calculate the difference
type CPUUtilizationCalculator struct {
	previousCPUTimes []cpu.TimesStat
}

// CalculateAndRecord calculates the cpu utilization for the different cpu states comparing previously
// stored []cpu.TimesStat and time.Time and current []cpu.TimesStat and current time.Time
// If no previous data is stored it will return empty slice of CPUUtilization and no error
func (c *CPUUtilizationCalculator) CalculateAndRecord(now pcommon.Timestamp, cpuTimes []cpu.TimesStat, recorder func(pcommon.Timestamp, CPUUtilization)) error {
	_ = "STUB: not implemented"
	return nil
}

// cpuUtilization calculates the difference between 2 cpu.TimesStat using spent time between them
func cpuUtilization(timeStart, timeEnd cpu.TimesStat) CPUUtilization {
	_ = "STUB: not implemented"
	return *new(CPUUtilization)
}

// cpuTimeForCPU returns cpu.TimesStat from a slice of cpu.TimesStat based on CPU
// If CPU is not found and error will be returned
func cpuTimeForCPU(cpuNum string, times []cpu.TimesStat) (cpu.TimesStat, error) {
	_ = "STUB: not implemented"
	return *new(cpu.TimesStat), nil
}

// Copied from cpu.TimesStat.Total(), since that func is deprecated.
func totalCPU(c cpu.TimesStat) float64 { _ = "STUB: not implemented"; return 0 }

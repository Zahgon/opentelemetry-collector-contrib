// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ecsobserver // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/ecsobserver"

import (
	"go.uber.org/zap"
)

const (
	defaultMetricsPath = "/metrics"
)

// CommonExporterConfig should be embedded into filter config.
// They set labels like job, metrics_path etc. that can override prometheus default.
type CommonExporterConfig struct {
	JobName      string `mapstructure:"job_name" yaml:"job_name"`
	MetricsPath  string `mapstructure:"metrics_path" yaml:"metrics_path"`
	MetricsPorts []int  `mapstructure:"metrics_ports" yaml:"metrics_ports"`
}

// newExportSetting checks if there are duplicated metrics ports.
func (c *CommonExporterConfig) newExportSetting() (*commonExportSetting, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// commonExportSetting is generated from CommonExportConfig with some util methods.
type commonExportSetting struct {
	CommonExporterConfig
	metricsPorts map[int]bool
}

func (s *commonExportSetting) hasContainerPort(containerPort int) bool {
	_ = "STUB: not implemented"
	return false
}

// taskExporter converts annotated taskAnnotated into prometheusECSTarget.
type taskExporter struct {
	logger  *zap.Logger
	cluster string
}

func newTaskExporter(logger *zap.Logger, cluster string) *taskExporter {
	_ = "STUB: not implemented"
	return nil
}

// exportTasks loops a list of tasks and export prometheus scrape targets.
// It keeps track of error but does NOT stop when error occurs.
// The returned targets are valid, invalid targets are saved in a multi error.
// Caller can ignore the error because the only source is failing to get ip and port.
// The error(s) can generates debug log or metrics.
// To print the error with its task as context, use printExporterErrors.
func (e *taskExporter) exportTasks(tasks []*taskAnnotated) ([]prometheusECSTarget, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// if err == nil, AppendInto does nothing
// Even if there are error, returned targets are still valid.

// exportTask exports all the matched container within a single task.
// One task can contain multiple containers. One container can have more than one target
// if there are multiple ports in `metrics_port`.
func (e *taskExporter) exportTask(task *taskAnnotated) ([]prometheusECSTarget, error) {
	_ = "STUB: not implemented"
	// All targets in one task shares same IP.
	return nil, nil
}

// Base for all the containers in this task, most attributes are same.

// Shallow copy task level attributes

// Add container specific info

// Multiple targets for a single container

// Shallow copy from container

// Skip this target and keep track of port error, does not abort.

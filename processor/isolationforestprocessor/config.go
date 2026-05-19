// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// config.go - CORRECTED VERSION with proper interface implementations

package isolationforestprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/isolationforestprocessor"

import (
	"time"

	"go.opentelemetry.io/collector/component"
)

// Config represents the configuration for the isolation forest processor.
type Config struct {
	ForestSize              int               `mapstructure:"forest_size"`
	SubsampleSize           int               `mapstructure:"subsample_size"`
	ContaminationRate       float64           `mapstructure:"contamination_rate"`
	Mode                    string            `mapstructure:"mode"`
	Threshold               float64           `mapstructure:"threshold"`
	TrainingWindow          string            `mapstructure:"training_window"`
	UpdateFrequency         string            `mapstructure:"update_frequency"`
	MinSamples              int               `mapstructure:"min_samples"`
	ScoreAttribute          string            `mapstructure:"score_attribute"`
	ClassificationAttribute string            `mapstructure:"classification_attribute"`
	Features                FeatureConfig     `mapstructure:"features"`
	Models                  []ModelConfig     `mapstructure:"models"`
	Performance             PerformanceConfig `mapstructure:"performance"`

	// Adaptive window sizing configuration
	AdaptiveWindow *AdaptiveWindowConfig `mapstructure:"adaptive_window"`
}

// AdaptiveWindowConfig configures automatic window size adjustment based on traffic patterns
type AdaptiveWindowConfig struct {
	// Core configuration
	Enabled        bool    `mapstructure:"enabled"`         // Enable adaptive sizing
	MinWindowSize  int     `mapstructure:"min_window_size"` // Minimum samples to keep
	MaxWindowSize  int     `mapstructure:"max_window_size"` // Maximum samples (memory protection)
	MemoryLimitMB  int     `mapstructure:"memory_limit_mb"` // Auto-shrink when exceeded
	AdaptationRate float64 `mapstructure:"adaptation_rate"` // Adjustment speed (0.0-1.0)

	// Optional parameters with defaults
	VelocityThreshold      float64 `mapstructure:"velocity_threshold"`       // Grow when >N samples/sec
	StabilityCheckInterval string  `mapstructure:"stability_check_interval"` // Check model accuracy interval
}

type FeatureConfig struct {
	Traces  []string `mapstructure:"traces"`
	Metrics []string `mapstructure:"metrics"`
	Logs    []string `mapstructure:"logs"`
}

type ModelConfig struct {
	Name              string            `mapstructure:"name"`
	Selector          map[string]string `mapstructure:"selector"`
	Features          []string          `mapstructure:"features"`
	Threshold         float64           `mapstructure:"threshold"`
	ForestSize        int               `mapstructure:"forest_size"`
	SubsampleSize     int               `mapstructure:"subsample_size"`
	ContaminationRate float64           `mapstructure:"contamination_rate"`
}

type PerformanceConfig struct {
	MaxMemoryMB     int `mapstructure:"max_memory_mb"`
	BatchSize       int `mapstructure:"batch_size"`
	ParallelWorkers int `mapstructure:"parallel_workers"`
}

// createDefaultConfig returns a configuration with sensible defaults
// Note: This function returns component.Config to match the expected signature.
func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

// Default adaptive window configuration (disabled by default for backward compatibility)

// Disabled by default - backward compatibility
// Match MinSamples for consistency
// Reasonable upper bound
// Half of total processor memory
// Conservative adjustment speed
// Default growth threshold
// Check model stability every 5 minutes

// Validate checks the configuration for logical consistency and valid parameter ranges.
func (cfg *Config) Validate() error { _ = "STUB: not implemented"; return nil }

// Upper bound required by tests

// Require at least one feature type configured

// Validate adaptive window configuration

// validateAdaptiveWindow validates the adaptive window configuration
func (cfg *Config) validateAdaptiveWindow() error { _ = "STUB: not implemented"; return nil }

// Ensure consistency with main config

// Memory limit should be reasonable compared to total processor memory

// IsAdaptiveWindowEnabled returns true if adaptive window sizing is enabled
func (cfg *Config) IsAdaptiveWindowEnabled() bool { _ = "STUB: not implemented"; return false }

// GetStabilityCheckInterval returns the stability check interval duration
func (cfg *Config) GetStabilityCheckInterval() (time.Duration, error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil
}

// Default

func (cfg *Config) GetTrainingWindowDuration() (time.Duration, error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil
}

func (cfg *Config) GetUpdateFrequencyDuration() (time.Duration, error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil
}

func (cfg *Config) IsMultiModelMode() bool { _ = "STUB: not implemented"; return false }

func (cfg *Config) GetModelForAttributes(attributes map[string]any) *ModelConfig {
	_ = "STUB: not implemented"
	return nil
}

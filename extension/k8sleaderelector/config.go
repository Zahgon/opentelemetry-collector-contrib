// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package k8sleaderelector // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/k8sleaderelector"

import (
	"time"

	"k8s.io/client-go/kubernetes"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/k8sconfig"
)

// Config is the configuration for the leader elector extension.
type Config struct {
	k8sconfig.APIConfig `mapstructure:",squash"`
	LeaseName           string        `mapstructure:"lease_name"`
	LeaseNamespace      string        `mapstructure:"lease_namespace"`
	LeaseDuration       time.Duration `mapstructure:"lease_duration"`
	RenewDuration       time.Duration `mapstructure:"renew_deadline"`
	RetryPeriod         time.Duration `mapstructure:"retry_period"`
	makeClient          func(apiConf k8sconfig.APIConfig) (kubernetes.Interface, error)
}

func (cfg *Config) getK8sClient() (kubernetes.Interface, error) {
	_ = "STUB: not implemented"
	return *new(kubernetes.Interface), nil
}

// Validate checks if the extension configuration is valid
func (cfg *Config) Validate() error { _ = "STUB: not implemented"; return nil }

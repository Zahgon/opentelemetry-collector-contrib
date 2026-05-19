// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package k8seventsreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8seventsreceiver"

import (
	"go.opentelemetry.io/collector/component"
	"k8s.io/client-go/dynamic"
	k8s "k8s.io/client-go/kubernetes"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/k8sconfig"
)

// Config defines configuration for kubernetes events receiver.
type Config struct {
	k8sconfig.APIConfig `mapstructure:",squash"`

	// List of ‘namespaces’ to collect events from.
	Namespaces []string `mapstructure:"namespaces"`

	// Storage is the ID of the storage extension to use for resource version persistence.
	// When set, the receiver will persist the latest resource version and resume from it on restart,
	// preventing duplicate events. Only valid for watch mode (which is the only mode this receiver uses).
	Storage *component.ID `mapstructure:"storage"`

	K8sLeaderElector *component.ID `mapstructure:"k8s_leader_elector"`

	// For mocking
	makeClient        func(apiConf k8sconfig.APIConfig) (k8s.Interface, error)
	makeDynamicClient func(apiConf k8sconfig.APIConfig) (dynamic.Interface, error)
}

func (cfg *Config) Validate() error { _ = "STUB: not implemented"; return nil }

func (cfg *Config) getK8sClient() (k8s.Interface, error) {
	_ = "STUB: not implemented"
	return *new(k8s.Interface), nil
}

func (cfg *Config) getDynamicClient() (dynamic.Interface, error) {
	_ = "STUB: not implemented"
	return *new(dynamic.Interface), nil
}

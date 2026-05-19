// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package k8sclusterreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8sclusterreceiver"

import (
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/receiver"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/sharedcomponent"
)

const (
	// supported distributions
	distributionKubernetes = "kubernetes"
	distributionOpenShift  = "openshift"

	// Default config values.
	defaultCollectionInterval         = 10 * time.Second
	defaultDistribution               = distributionKubernetes
	defaultMetadataCollectionInterval = 5 * time.Minute
)

var defaultNodeConditionsToReport = []string{"Ready"}

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

// NewFactory creates a factory for k8s_cluster receiver.
func NewFactory() receiver.Factory { _ = "STUB: not implemented"; return *new(receiver.Factory) }

// This is the map of already created k8scluster receivers for particular configurations.
// We maintain this map because the Factory is asked log and metric receivers separately
// when it gets CreateLogs() and CreateMetrics() but they must not
// create separate objects, they must use one receiver object per configuration.
var receivers = sharedcomponent.NewSharedComponents()

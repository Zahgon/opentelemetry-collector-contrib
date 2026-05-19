// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package awscontainerinsightreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awscontainerinsightreceiver"

import (
	"context"
	"sync"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
)

var _ receiver.Metrics = (*awsContainerInsightReceiver)(nil)

type metricsProvider interface {
	GetMetrics() []pmetric.Metrics
	Shutdown() error
}

// awsContainerInsightReceiver implements the receiver.Metrics
type awsContainerInsightReceiver struct {
	settings     component.TelemetrySettings
	nextConsumer consumer.Metrics
	config       *Config
	cancel       context.CancelFunc
	cancelWg     sync.WaitGroup
	cadvisor     metricsProvider
	k8sapiserver metricsProvider
}

// newAWSContainerInsightReceiver creates the aws container insight receiver with the given parameters.
func newAWSContainerInsightReceiver(
	settings component.TelemetrySettings,
	config *Config,
	nextConsumer consumer.Metrics,
) (receiver.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Metrics), nil
}

// Start collecting metrics from cadvisor and k8s api server (if it is an elected leader)
func (acir *awsContainerInsightReceiver) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

// cadvisor collects data at dynamical intervals (from 1 to 15 seconds). If the ticker happens
// at beginning of a minute, it might read the data collected at end of last minute. To avoid this,
// we want to wait until at least two cadvisor collection intervals happens before collecting the metrics

// Shutdown stops the awsContainerInsightReceiver receiver.
func (acir *awsContainerInsightReceiver) Shutdown(context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// collectData collects container stats from cAdvisor and k8s api server (if it is an elected leader)
func (acir *awsContainerInsightReceiver) collectData(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

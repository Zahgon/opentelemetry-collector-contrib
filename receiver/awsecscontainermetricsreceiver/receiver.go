// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package awsecscontainermetricsreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsecscontainermetricsreceiver"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/ecsutil"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsecscontainermetricsreceiver/internal/awsecscontainermetrics"
)

var _ receiver.Metrics = (*awsEcsContainerMetricsReceiver)(nil)

// awsEcsContainerMetricsReceiver implements the receiver.Metrics for aws ecs container metrics.
type awsEcsContainerMetricsReceiver struct {
	logger       *zap.Logger
	nextConsumer consumer.Metrics
	config       *Config
	cancel       context.CancelFunc
	restClient   ecsutil.RestClient
	provider     *awsecscontainermetrics.StatsProvider
}

// New creates the aws ecs container metrics receiver with the given parameters.
func newAWSECSContainermetrics(
	logger *zap.Logger,
	config *Config,
	nextConsumer consumer.Metrics,
	rest ecsutil.RestClient,
) (receiver.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Metrics), nil
}

// Start begins collecting metrics from Amazon ECS task metadata endpoint.
func (aecmr *awsEcsContainerMetricsReceiver) Start(ctx context.Context, _ component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

// Shutdown stops the awsecscontainermetricsreceiver receiver.
func (aecmr *awsEcsContainerMetricsReceiver) Shutdown(context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// collectDataFromEndpoint collects container stats from Amazon ECS Task Metadata Endpoint
func (aecmr *awsEcsContainerMetricsReceiver) collectDataFromEndpoint(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: report self metrics using obsreport

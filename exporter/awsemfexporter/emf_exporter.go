// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package awsemfexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsemfexporter"

import (
	"context"
	"sync"

	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/pdata/pmetric"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/cwlogs"
)

const (
	// OutputDestination Options
	outputDestinationCloudWatch = "cloudwatch"
	outputDestinationStdout     = "stdout"

	// AppSignals EMF config
	appSignalsMetricNamespace    = "ApplicationSignals"
	appSignalsLogGroupNamePrefix = "/aws/application-signals/"
)

type emfExporter struct {
	pusherMap        sync.Map
	svcStructuredLog *cwlogs.Client
	config           *Config

	metricTranslator metricTranslator

	retryCnt    int
	collectorID string
}

// newEmfExporter creates a new exporter using exporterhelper
func newEmfExporter(ctx context.Context, config *Config, set exporter.Settings) (*emfExporter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// create CWLogs client with aws session config

func (emf *emfExporter) pushMetricsData(ctx context.Context, md pmetric.Metrics) error {
	_ = "STUB: not implemented"
	return nil
}

// Currently we only support two options for "OutputDestination".

// TODO now we only have one logPusher, so it's ok to return after first error occurred

func (emf *emfExporter) getPusher(key cwlogs.StreamKey) cwlogs.Pusher {
	_ = "STUB: not implemented"
	return *new(cwlogs.Pusher)
}

func (emf *emfExporter) listPushers() []cwlogs.Pusher { _ = "STUB: not implemented"; return nil }

// shutdown stops the exporter and is invoked during shutdown.
func (emf *emfExporter) shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func wrapErrorIfBadRequest(err error) error { _ = "STUB: not implemented"; return nil }

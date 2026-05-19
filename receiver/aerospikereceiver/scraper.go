// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package aerospikereceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/aerospikereceiver"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/scraper/scrapererror"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/aerospikereceiver/internal/metadata"
)

// aerospikeReceiver is a metrics receiver using the Aerospike interface to collect
type aerospikeReceiver struct {
	config        *Config
	consumer      consumer.Metrics
	clientFactory clientFactoryFunc
	client        Aerospike
	mb            *metadata.MetricsBuilder
	logger        *zap.SugaredLogger
}

// clientFactoryFunc creates an Aerospike connection to the given host and port
type clientFactoryFunc func() (Aerospike, error)

// newAerospikeReceiver creates a new aerospikeReceiver connected to the endpoint provided in cfg
//
// If the host or port can't be parsed from endpoint, an error is returned.
func newAerospikeReceiver(params receiver.Settings, cfg *Config, consumer consumer.Metrics) (*aerospikeReceiver, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *aerospikeReceiver) start(_ context.Context, _ component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

//  .Sugar().Warnf("initial client creation failed: %w", err)

func (r *aerospikeReceiver) shutdown(_ context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// scrape scrapes both Node and Namespace metrics from the provided Aerospike node.
// If CollectClusterMetrics is true, it then scrapes every discovered node
func (r *aerospikeReceiver) scrape(_ context.Context) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

// emitNode records node metrics and emits the resource. If statistics are missing in INFO, nothing is recorded
func (r *aerospikeReceiver) emitNode(info map[string]string, now pcommon.Timestamp, errs *scrapererror.ScrapeErrors) {
	_ = "STUB: not implemented"
	return
}

// scrapeNamespaces records metrics for all namespaces on a node
// The given client is used to collect namespace metrics, which is connected to a single node
func (r *aerospikeReceiver) scrapeNamespaces(client Aerospike, now pcommon.Timestamp, errs *scrapererror.ScrapeErrors) {
	_ = "STUB: not implemented"
	return
}

// emitNamespace emits a namespace resource with its name as resource attribute
func (r *aerospikeReceiver) emitNamespace(info map[string]string, now pcommon.Timestamp, errs *scrapererror.ScrapeErrors) {
	_ = "STUB: not implemented"
	return
}

// Capacity

// Memory usage

// Scans

// Pre Aerospike 6.0 query metrics. These were always done on secondary indexes, otherwise they were counted as a scan.

// PI queries

// SI queries

// GeoJSON region queries

// Compression

// 'Delete' transactions

// 'Read' transactions

// UDF transactions

// 'Write' transactions

// addPartialIfError adds a partial error if the given error isn't nil
func addPartialIfError(errs *scrapererror.ScrapeErrors, err error) {
	_ = "STUB: not implemented"
	return
}

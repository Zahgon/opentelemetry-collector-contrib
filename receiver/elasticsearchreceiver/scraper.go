// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package elasticsearchreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/elasticsearchreceiver"

import (
	"context"
	"errors"

	"github.com/hashicorp/go-version"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/scraper/scrapererror"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/elasticsearchreceiver/internal/metadata"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/elasticsearchreceiver/internal/model"
)

var (
	es7_10 = func() *version.Version {
		v, _ := version.NewVersion("7.10")
		return v
	}()
	es7_13 = func() *version.Version {
		v, _ := version.NewVersion("7.13")
		return v
	}()
)

var errUnknownClusterStatus = errors.New("unknown cluster status")

type elasticsearchScraper struct {
	client      elasticsearchClient
	settings    component.TelemetrySettings
	cfg         *Config
	mb          *metadata.MetricsBuilder
	version     *version.Version
	clusterName string
}

func newElasticSearchScraper(
	settings receiver.Settings,
	cfg *Config,
) *elasticsearchScraper {
	_ = "STUB: not implemented"
	return nil
}

func (r *elasticsearchScraper) start(ctx context.Context, host component.Host) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (r *elasticsearchScraper) scrape(ctx context.Context) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

// scrapeVersion gets and assigns the elasticsearch version number
func (r *elasticsearchScraper) getClusterMetadata(ctx context.Context, errs *scrapererror.ScrapeErrors) {
	_ = "STUB: not implemented"
	return
}

// scrapeNodeMetrics scrapes adds node-level metrics to the given MetricSlice from the NodeStats endpoint
func (r *elasticsearchScraper) scrapeNodeMetrics(ctx context.Context, now pcommon.Timestamp, errs *scrapererror.ScrapeErrors) {
	_ = "STUB: not implemented"
	return
}

// Certain node metadata is not available from the /_nodes/stats endpoint. Therefore, we need to get this metadata
// from the /_nodes endpoint.

// Elasticsearch version 7.13+ is required to collect `elasticsearch.node.shards.data_set.size`.
// Reference: https://github.com/elastic/elasticsearch/pull/70625/files#diff-354b5b1f25978b5c638cb707622ae79b42b40aace6f27f3f9d5dd1e31e67b1caR7

// Elasticsearch sends this data in percent, but we want to represent it as a number between 0 and 1, so we need to divide.
// Additionally, if the usage is not known, ES will send '-1'. We do not want to report the metric in this case.

// Elasticsearch sends this data in percent, but we want to represent it as a number between 0 and 1, so we need to divide.

// Elasticsearch version 7.10+ is required to collect `elasticsearch.indexing_pressure.memory.limit`.
// Reference: https://github.com/elastic/elasticsearch/pull/60342/files#diff-13864344bab3afc267797d67b2746e2939a3fd8af7611ac9fbda376323e2f5eaR37

// the node_linux.json payload response for "elasticsearch.cluster.state_update.time" with attributes "unchanged" has 2 attributes "computation_time_millis" and "notification_time_millis".
// All other metrics for "unchanged" should be skipped to prevent 0 emitted metrics"
// https://github.com/elastic/elasticsearch/pull/76771/files#diff-8bbfc581d91f9440e53098ea7d7864aeaeac1fc83a714133e4aafe38eba8ed90R2098

func (r *elasticsearchScraper) scrapeClusterMetrics(ctx context.Context, now pcommon.Timestamp, errs *scrapererror.ScrapeErrors) {
	_ = "STUB: not implemented"
	return
}

func (r *elasticsearchScraper) scrapeClusterStatsMetrics(ctx context.Context, now pcommon.Timestamp, errs *scrapererror.ScrapeErrors) {
	_ = "STUB: not implemented"
	return
}

func (r *elasticsearchScraper) scrapeClusterHealthMetrics(ctx context.Context, now pcommon.Timestamp, errs *scrapererror.ScrapeErrors) {
	_ = "STUB: not implemented"
	return
}

func (r *elasticsearchScraper) scrapeIndicesMetrics(ctx context.Context, now pcommon.Timestamp, errs *scrapererror.ScrapeErrors) {
	_ = "STUB: not implemented"
	return
}

// The metrics for all indices are queried by using "_all" name and hence its the name used for labeling them.

func (r *elasticsearchScraper) scrapeOneIndexMetrics(now pcommon.Timestamp, name string, stats *model.IndexStatsIndexInfo) {
	_ = "STUB: not implemented"
	return
}

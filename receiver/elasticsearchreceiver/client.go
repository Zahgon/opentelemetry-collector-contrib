// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package elasticsearchreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/elasticsearchreceiver"

import (
	"context"
	"errors"
	"net/http"
	"net/url"

	"github.com/hashicorp/go-version"
	"go.opentelemetry.io/collector/component"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/elasticsearchreceiver/internal/model"
)

var (
	errUnauthenticated = errors.New("status 401, unauthenticated")
	errUnauthorized    = errors.New("status 403, unauthorized")
)

// elasticsearchClient defines the interface to retrieve metrics from an Elasticsearch cluster.
type elasticsearchClient interface {
	Nodes(ctx context.Context, nodes []string) (*model.Nodes, error)
	NodeStats(ctx context.Context, nodes []string) (*model.NodeStats, error)
	ClusterHealth(ctx context.Context) (*model.ClusterHealth, error)
	IndexStats(ctx context.Context, indices []string) (*model.IndexStats, error)
	ClusterMetadata(ctx context.Context) (*model.ClusterMetadataResponse, error)
	ClusterStats(ctx context.Context, nodes []string) (*model.ClusterStats, error)
}

// defaultElasticsearchClient is the main implementation of elasticsearchClient.
// It retrieves the required metrics from Elasticsearch's REST api.
type defaultElasticsearchClient struct {
	client     *http.Client
	endpoint   *url.URL
	authHeader string
	logger     *zap.Logger
	version    *version.Version
}

var _ elasticsearchClient = (*defaultElasticsearchClient)(nil)

func newElasticsearchClient(ctx context.Context, settings component.TelemetrySettings, c Config, h component.Host) (*defaultElasticsearchClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Try update es version

var es7_9 = func() *version.Version {
	v, _ := version.NewVersion("7.9")
	return v
}()

const (
	// A comma separated list of metrics that will be gathered from NodeStats.
	// https://www.elastic.co/guide/en/elasticsearch/reference/7.9/cluster-nodes-stats.html#cluster-nodes-stats-api-path-params
	defaultNodeStatsMetrics = "breaker,indices,process,jvm,thread_pool,transport,http,fs,ingest,indices,adaptive_selection,discovery,script,os"

	// Extra NodeStats Metrics that are only supported on and after 7.9
	nodeStatsMetricsAfter7_9 = ",indexing_pressure"

	// A comma separated list of metrics that will be gathered from Nodes.
	// The available metrics are documented here for Elasticsearch 7.9:
	// https://www.elastic.co/guide/en/elasticsearch/reference/7.9/cluster-nodes-info.html
	// Note: This constant should remain empty as the receiver will only retrieve metadata from the /_nodes endpoint, not metrics.
	nodesMetrics = ""

	// A comma separated list of index metrics that will be gathered from NodeStats.
	nodeStatsIndexMetrics = "store,docs,indexing,get,search,merge,refresh,flush,warmer,query_cache,fielddata,translog"

	// A comma separated list of metrics that will be gathered from IndexStats.
	indexStatsMetrics = "_all"
)

func (c defaultElasticsearchClient) Nodes(ctx context.Context, nodeIDs []string) (*model.Nodes, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c defaultElasticsearchClient) NodeStats(ctx context.Context, nodes []string) (*model.NodeStats, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c defaultElasticsearchClient) ClusterHealth(ctx context.Context) (*model.ClusterHealth, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c defaultElasticsearchClient) IndexStats(ctx context.Context, indices []string) (*model.IndexStats, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *defaultElasticsearchClient) ClusterMetadata(ctx context.Context) (*model.ClusterMetadataResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c defaultElasticsearchClient) ClusterStats(ctx context.Context, nodes []string) (*model.ClusterStats, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c defaultElasticsearchClient) doRequest(ctx context.Context, path string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// See https://www.elastic.co/docs/reference/elasticsearch/rest-apis/api-conventions#api-compatibility
// the compatible-with=8 should signal to newer version of Elasticsearch to use the v8.x API format

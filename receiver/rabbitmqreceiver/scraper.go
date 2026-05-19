// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package rabbitmqreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/rabbitmqreceiver"

import (
	"context"
	"errors"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/rabbitmqreceiver/internal/metadata"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/rabbitmqreceiver/internal/models"
)

var errClientNotInit = errors.New("client not initialized")

// Names of metrics in message_stats
const (
	deliverStat        = "deliver"
	publishStat        = "publish"
	ackStat            = "ack"
	dropUnroutableStat = "drop_unroutable"
)

// Metrics to gather from queue message_stats structure
var messageStatMetrics = []string{
	deliverStat,
	publishStat,
	ackStat,
	dropUnroutableStat,
}

// rabbitmqScraper handles scraping of RabbitMQ metrics
type rabbitmqScraper struct {
	client   client
	logger   *zap.Logger
	cfg      *Config
	settings component.TelemetrySettings
	mb       *metadata.MetricsBuilder
}

// newScraper creates a new scraper
func newScraper(logger *zap.Logger, cfg *Config, settings receiver.Settings) *rabbitmqScraper {
	_ = "STUB: not implemented"
	return nil
}

// start starts the scraper by creating a new HTTP Client on the scraper
func (r *rabbitmqScraper) start(ctx context.Context, host component.Host) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// scrape collects metrics from the RabbitMQ API
func (r *rabbitmqScraper) scrape(ctx context.Context) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

// Validate we don't attempt to scrape without initializing the client

// Collect queue metrics

// Collect node metrics

// Emit collected metrics

// Define err at the correct scope

// Return a PartialScrapeError if no metrics were collected

func (r *rabbitmqScraper) collectQueueMetrics(ctx context.Context, now pcommon.Timestamp) error {
	_ = "STUB: not implemented"
	return nil
}

// Collect metrics for each queue

func (r *rabbitmqScraper) collectNodeMetrics(ctx context.Context, now pcommon.Timestamp) error {
	_ = "STUB: not implemented"
	return nil
}

// collectQueue collects metrics
func (r *rabbitmqScraper) collectQueue(queue *models.Queue, now pcommon.Timestamp) {
	_ = "STUB: not implemented"
	return
}

// Get metric value

// A metric may not exist if the actions that increment it do not occur

// Convert to int64

// Log warning if the metric is not in the format we expect

// collectNode collects metrics for a specific RabbitMQ node
func (r *rabbitmqScraper) collectNode(node *models.Node, now pcommon.Timestamp) {
	_ = "STUB: not implemented"
	return
}

func boolToInt64(b bool) int64 { _ = "STUB: not implemented"; return 0 }

// convertValToInt64 values from message state unmarshal as float64s but should be int64.
// Need to do a double cast to get an int64.
// This should never fail but worth checking just in case.
func convertValToInt64(val any) (int64, bool) { _ = "STUB: not implemented"; return 0, false }

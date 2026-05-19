// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package mongodbreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mongodbreceiver"

import (
	"context"

	"github.com/google/uuid"
	"github.com/hashicorp/go-version"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/scraper/scrapererror"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mongodbreceiver/internal/metadata"
)

var (
	unknownVersion = func() *version.Version { return version.Must(version.NewVersion("0.0")) }

	// otelNamespaceUUID is the official OTel namespace UUID for deterministic UUID v5 generation.
	otelNamespaceUUID = uuid.MustParse("4d63009a-8d0f-11ee-aad7-4c796ed8e320")
)

// generateInstanceID generates a deterministic UUID v5 from server address and port.
func generateInstanceID(serverAddress string, serverPort int64) string {
	_ = "STUB: not implemented"
	return ""
}

type mongodbScraper struct {
	logger             *zap.Logger
	config             *Config
	client             client
	secondaryClients   []client
	mongoVersion       *version.Version
	mb                 *metadata.MetricsBuilder
	prevReplTimestamp  pcommon.Timestamp
	prevReplCounts     map[string]int64
	prevTimestamp      pcommon.Timestamp
	prevFlushTimestamp pcommon.Timestamp
	prevCounts         map[string]int64
	prevFlushCount     int64
}

func newMongodbScraper(settings receiver.Settings, config *Config) *mongodbScraper {
	_ = "STUB: not implemented"
	return nil
}

func (s *mongodbScraper) start(ctx context.Context, _ component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

// Skip secondary host discovery if direct connection is enabled

func (s *mongodbScraper) shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (s *mongodbScraper) scrape(ctx context.Context) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

func (s *mongodbScraper) collectMetrics(ctx context.Context, errs *scrapererror.ScrapeErrors) {
	_ = "STUB: not implemented"
	return
}

// Collect metrics for each database

// Emit single resource for the server

func (s *mongodbScraper) collectDatabase(ctx context.Context, now pcommon.Timestamp, databaseName string, errs *scrapererror.ScrapeErrors) {
	_ = "STUB: not implemented"
	return
}

func (s *mongodbScraper) collectTopStats(ctx context.Context, now pcommon.Timestamp, errs *scrapererror.ScrapeErrors) {
	_ = "STUB: not implemented"
	return
}

func (s *mongodbScraper) collectIndexStats(ctx context.Context, now pcommon.Timestamp, databaseName, collectionName string, errs *scrapererror.ScrapeErrors) {
	_ = "STUB: not implemented"
	return
}

func (s *mongodbScraper) recordDBStats(now pcommon.Timestamp, doc bson.M, dbName string, errs *scrapererror.ScrapeErrors) {
	_ = "STUB: not implemented"
	return
}

func (s *mongodbScraper) recordNormalServerStats(now pcommon.Timestamp, doc bson.M, dbName string, errs *scrapererror.ScrapeErrors) {
	_ = "STUB: not implemented"
	return
}

func (s *mongodbScraper) recordAdminStats(now pcommon.Timestamp, document bson.M, errs *scrapererror.ScrapeErrors) {
	_ = "STUB: not implemented"
	return
}

func (s *mongodbScraper) recordIndexStats(now pcommon.Timestamp, indexStats []bson.M, databaseName, collectionName string, errs *scrapererror.ScrapeErrors) {
	_ = "STUB: not implemented"
	return
}

func serverAddressAndPort(serverStatus bson.M) (string, int64, error) {
	_ = "STUB: not implemented"
	return "", 0, nil
}

func (s *mongodbScraper) findSecondaryHosts(ctx context.Context) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Only add actual secondaries, not arbiters or other states

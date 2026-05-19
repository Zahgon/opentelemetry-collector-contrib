// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package couchdbreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/couchdbreceiver"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/scraper/scrapererror"
)

func (c *couchdbScraper) recordCouchdbAverageRequestTimeDataPoint(now pcommon.Timestamp, stats map[string]any, errs *scrapererror.ScrapeErrors) {
	_ = "STUB: not implemented"
	return
}

func (c *couchdbScraper) recordCouchdbHttpdBulkRequestsDataPoint(now pcommon.Timestamp, stats map[string]any, errs *scrapererror.ScrapeErrors) {
	_ = "STUB: not implemented"
	return
}

func (c *couchdbScraper) recordCouchdbHttpdRequestsDataPoint(now pcommon.Timestamp, stats map[string]any, errs *scrapererror.ScrapeErrors) {
	_ = "STUB: not implemented"
	return
}

func (c *couchdbScraper) recordCouchdbHttpdResponsesDataPoint(now pcommon.Timestamp, stats map[string]any, errs *scrapererror.ScrapeErrors) {
	_ = "STUB: not implemented"
	return
}

func (c *couchdbScraper) recordCouchdbHttpdViewsDataPoint(now pcommon.Timestamp, stats map[string]any, errs *scrapererror.ScrapeErrors) {
	_ = "STUB: not implemented"
	return
}

func (c *couchdbScraper) recordCouchdbDatabaseOpenDataPoint(now pcommon.Timestamp, stats map[string]any, errs *scrapererror.ScrapeErrors) {
	_ = "STUB: not implemented"
	return
}

func (c *couchdbScraper) recordCouchdbFileDescriptorOpenDataPoint(now pcommon.Timestamp, stats map[string]any, errs *scrapererror.ScrapeErrors) {
	_ = "STUB: not implemented"
	return
}

func (c *couchdbScraper) recordCouchdbDatabaseOperationsDataPoint(now pcommon.Timestamp, stats map[string]any, errs *scrapererror.ScrapeErrors) {
	_ = "STUB: not implemented"
	return
}

func getValueFromBody(keys []string, body map[string]any) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (*couchdbScraper) parseInt(value any) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (*couchdbScraper) parseFloat(value any) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

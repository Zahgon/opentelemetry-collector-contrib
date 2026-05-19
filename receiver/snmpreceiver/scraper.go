// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package snmpreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/snmpreceiver"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/scraper/scrapererror"
	"go.uber.org/zap"
)

var (
	// Error messages
	errMsgBadValueType                   = `returned metric SNMP data type for OID '%s' is not supported`
	errMsgIndexedAttributesBadValueType  = `returned attribute SNMP data type for OID '%s' from column OID '%s' is not supported`
	errMsgScalarAttributesBadValueType   = `returned attribute SNMP data type for OID '%s' is not supported`
	errMsgOIDAttributeEmptyValue         = `not creating indexed metric '%s' datapoint: %w`
	errMsgAttributeEmptyValue            = `metric OID attribute value is blank`
	errMsgResourceAttributeEmptyValue    = `related resource attribute value is blank`
	errMsgOIDResourceAttributeEmptyValue = `not creating indexed metric '%s' or resource: %w`
	errMsgScalarOIDProcessing            = `problem processing scalar metric data for OID '%s': %w`
	errMsgIndexedMetricOIDProcessing     = `problem processing indexed metric data for OID '%s' from column OID '%s': %w`
	errMsgScalarAttributeOIDProcessing   = `problem processing scalar attribute data from scalar OID '%s': %w`
	errMsgIndexedAttributeOIDProcessing  = `problem processing indexed attribute data for OID '%s' from column OID '%s': %w`
)

// snmpScraper handles scraping of SNMP metrics
type snmpScraper struct {
	client    client
	logger    *zap.Logger
	cfg       *Config
	settings  receiver.Settings
	startTime pcommon.Timestamp
}

type indexedAttributeValues map[string]string

// newScraper creates an initialized snmpScraper
func newScraper(logger *zap.Logger, cfg *Config, settings receiver.Settings) *snmpScraper {
	_ = "STUB: not implemented"
	return nil
}

// start gets the client ready
func (s *snmpScraper) start(_ context.Context, _ component.Host) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// scrape collects and creates OTEL metrics from a SNMP environment
func (s *snmpScraper) scrape(_ context.Context) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

// Create the metrics helper which will help manage a lot of the otel metric and resource functionality

// Try to scrape scalar OID based metrics

// Try to scrape column OID based metrics

// scrapeScalarMetrics retrieves all SNMP data from scalar OIDs and turns the returned scalar data
// into metrics with optional enum attributes
func (s *snmpScraper) scrapeScalarMetrics(
	metricHelper *otelMetricHelper,
	configHelper *configHelper,
	scraperErrors *scrapererror.ScrapeErrors,
) {
	_ = "STUB: not implemented"
	return
}

// If no scalar metric configs, nothing else to do

// Retrieve all SNMP data from scalar metric OIDs

// If no scalar data, nothing else to do

// Retrieve scalar OID SNMP data for resource attributes

// For each piece of SNMP data, attempt to create the necessary OTEL structures (resources/metrics/datapoints)

// scrapeIndexedMetrics retrieves all SNMP data from column OIDs and turns the returned indexed data
// into metrics with optional attribute and/or resource attributes
func (s *snmpScraper) scrapeIndexedMetrics(
	metricHelper *otelMetricHelper,
	configHelper *configHelper,
	scraperErrors *scrapererror.ScrapeErrors,
) {
	_ = "STUB: not implemented"
	return
}

// If no column metric configs, nothing else to do

// Retrieve column OID SNMP indexed data for attributes

// Retrieve column OID SNMP indexed data for resource attributes

// Retrieve scalar OID SNMP data for resource attributes

// Retrieve all SNMP indexed data from column metric OIDs

// For each piece of SNMP data, attempt to create the necessary OTEL structures (resources/metrics/datapoints)

// scalarDataToMetric will take one piece of SNMP scalar data and turn it into a datapoint for
// either a new or existing metric with attributes based on the related configs
func (*snmpScraper) scalarDataToMetric(
	data snmpData,
	metricHelper *otelMetricHelper,
	configHelper *configHelper,
	scalarResourceAttributes map[string]string,
) error {
	_ = "STUB: not implemented"
	// Get the related metric name for this SNMP indexed data
	return nil
}

// Keys will be determined from the related attribute config and enum values will come straight from
// the metric config's attribute values.

// Get resource attributes

// Create a resource key using all of the relevant resource attribute names

// Create general resource if we don't have any resource attributes

// Create a new resource if needed

// indexedDataToMetric will take one piece of column OID SNMP indexed metric data and turn it
// into a datapoint for either a new or existing metric with attributes that belongs to either
// a new or existing resource
func (s *snmpScraper) indexedDataToMetric(
	data snmpData,
	metricHelper *otelMetricHelper,
	configHelper *configHelper,
	columnOIDIndexedAttributeValues map[string]indexedAttributeValues,
	columnOIDIndexedResourceAttributeValues map[string]indexedAttributeValues,
	columnOIDScalarResourceAttributeValues map[string]string,
) error {
	_ = "STUB: not implemented"
	// Get the related metric name for this SNMP indexed data
	return nil
}

// Get data point attributes

// Get resource attributes

// Create a resource key using all of the relevant resource attribute names along
// with the row index of the SNMP data

// Check how many of the resource attributes on this metric are scalar

// If the only resource attributes on this metric are scalar, we don't need multiple resources

// Create a new resource if needed

func addMetricDataPointToResource(
	data snmpData,
	metricHelper *otelMetricHelper,
	configHelper *configHelper,
	metricName string,
	resourceKey string,
	dataPointAttributes map[string]string,
) error {
	_ = "STUB: not implemented"
	// Return an error if this SNMP indexed data is not of a useable type
	return nil
}

// Get the related metric config

// Create a new metric if needed

// Add data point to metric

// getScalarDataPointAttributes returns the key value pairs of attributes for a given metric config scalar OID
func getScalarDataPointAttributes(configHelper *configHelper, oid string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

// getIndexedDataPointAttributes gets attributes for this metric's datapoint based on the previously
// gathered attributes.
// Keys will be determined from the related attribute config and values will come a few
// different places.
// Enum attribute value - comes from the metric config's attribute data
// Indexed prefix attribute value - comes from the current SNMP data's index and the attribute
// config's prefix value
// Indexed OID attribute value - comes from the previously collected indexed attribute data
// using the current index and attribute config to access the correct value
func getIndexedDataPointAttributes(
	configHelper *configHelper,
	columnOID string,
	indexString string,
	columnOIDIndexedAttributeValues map[string]indexedAttributeValues,
) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Use alternate attribute key if available

// If no good attribute value could be found

// getResourceAttributes creates a map of key/values for all related resource attributes. Keys
// will come directly from the metric config's resource attribute values. Values will come
// from the related attribute config's prefix value plus the index OR the previously collected
// resource attribute indexed data.
func getResourceAttributes(
	configHelper *configHelper,
	columnOID string,
	indexString string,
	columnOIDIndexedResourceAttributeValues map[string]indexedAttributeValues,
	columnOIDScalarResourceAttributeValues map[string]string,
) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// scrapeScalarResourceAttributes retrieves all SNMP data from resource attribute
// config scalar OIDs and stores the returned data for later use by metrics
func (s *snmpScraper) scrapeScalarResourceAttributes(
	scalarOIDs []string,
	scraperErrors *scrapererror.ScrapeErrors,
) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

// If no scalar OID resource attribute configs, nothing else to do

// Retrieve all SNMP data from scalar resource attribute OIDs

// For each piece of SNMP data, store the necessary info to help create resources later if needed

// scalarDataToResourceAttribute provides a function which will take one piece of scalar OID SNMP data
// (for a resource attribute) and store it in a map for later use
func scalarDataToResourceAttribute(
	data snmpData,
	scalarOIDAttributeValues map[string]string,
) error {
	_ = "STUB: not implemented"
	// Get the string value of the SNMP data for the {resource} attribute value
	return nil
}

// Not explicitly checking these casts as this should be made safe in the client

// Store the {resource} attribute value in a map using the scalar OID as a key.
// This way we can match metrics to this data through the {resource} attribute config.

// scrapeIndexedAttributes retrieves all SNMP data from attribute (or resource attribute)
// config column OIDs and stores the returned indexed data for later use by metrics
func (s *snmpScraper) scrapeIndexedAttributes(
	columnOIDs []string,
	scraperErrors *scrapererror.ScrapeErrors,
) map[string]indexedAttributeValues {
	_ = "STUB: not implemented"
	return nil
}

// If no OID resource attribute configs, nothing else to do

// Retrieve all SNMP indexed data from column resource attribute OIDs

// For each piece of SNMP data, store the necessary info to help create resources later if needed

// indexedDataToAttribute provides a function which will take one piece of column OID SNMP indexed data
// (for either an attribute or resource attribute) and stores it in a map for later use (keyed by both
// {resource} attribute config column OID and OID index)
func indexedDataToAttribute(
	data snmpData,
	columnOIDIndexedAttributeValues map[string]indexedAttributeValues,
) error {
	_ = "STUB: not implemented"
	// Get the string value of the SNMP data for the {resource} attribute value
	return nil
}

// Not explicitly checking these casts as this should be made safe in the client

// Store the {resource} attribute value in a map using the column OID and OID index associated
// as keys. This way we can match indexed metrics to this data through the {resource} attribute
// config and the indices of the individual metric values

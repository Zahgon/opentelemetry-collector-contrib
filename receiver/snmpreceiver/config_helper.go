// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package snmpreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/snmpreceiver"

// configHelper contains many of the functions required to get various info from the SNMP config
type configHelper struct {
	cfg                         *Config
	metricScalarOIDs            []string
	metricColumnOIDs            []string
	attributeColumnOIDs         []string
	resourceAttributeScalarOIDs []string
	resourceAttributeColumnOIDs []string
	metricNamesByOID            map[string]string
	metricAttributesByOID       map[string][]Attribute
	resourceAttributesByOID     map[string][]string
}

// newConfigHelper returns a new configHelper with various pieces of static info saved for easy access
func newConfigHelper(cfg *Config) *configHelper { _ = "STUB: not implemented"; return nil }

// Group all metric scalar OIDs and metric column OIDs
// Also create a map of metric names with OID as key so the metric config will be easy to
// matchup later with returned SNMP data

// Data is returned by the client with '.' prefix on the OIDs.
// Making sure the prefix exists here in the configs so we can match it up with returned data later

// Data is returned by the client with '.' prefix on the OIDs.
// Making sure the prefix exists here in the configs so we can match it up with returned data later

// Find all attribute column OIDs

// Data is returned by the client with '.' prefix on the OIDs.
// Making sure the prefix exists here in the configs so we can match it up with returned data later

// Find all resource attribute scalar and column OIDs

// Data is returned by the client with '.' prefix on the OIDs.
// Making sure the prefix exists here in the configs so we can match it up with returned data later

// Data is returned by the client with '.' prefix on the OIDs.
// Making sure the prefix exists here in the configs so we can match it up with returned data later

// We expect these []string to be sorted later (i.e. mocks and resourceKey)

// getMetricScalarOIDs returns all of the scalar OIDs in the metric configs
func (h configHelper) getMetricScalarOIDs() []string { _ = "STUB: not implemented"; return nil }

// getMetricColumnOIDs returns all of the column OIDs in the metric configs
func (h configHelper) getMetricColumnOIDs() []string { _ = "STUB: not implemented"; return nil }

// getAttributeColumnOIDs returns all of the attribute column OIDs in the attribute configs
func (h configHelper) getAttributeColumnOIDs() []string { _ = "STUB: not implemented"; return nil }

// getResourceAttributeScalarOIDs returns all of the resource attribute scalar OIDs in the resource attribute configs
func (h configHelper) getResourceAttributeScalarOIDs() []string {
	_ = "STUB: not implemented"
	return nil
}

// getResourceAttributeColumnOIDs returns all of the resource attribute column OIDs in the resource attribute configs
func (h configHelper) getResourceAttributeColumnOIDs() []string {
	_ = "STUB: not implemented"
	return nil
}

// getMetricName a metric names based on a given OID
func (h configHelper) getMetricName(oid string) string { _ = "STUB: not implemented"; return "" }

// getMetricConfig returns a metric config based on a given name
func (h configHelper) getMetricConfig(name string) *MetricConfig {
	_ = "STUB: not implemented"
	return nil

	// getAttributeConfigValue returns the value of an attribute config
}

func (h configHelper) getAttributeConfigValue(name string) string {
	_ = "STUB: not implemented"
	return ""
}

// getAttributeConfigIndexedValuePrefix returns the indexed value prefix of an attribute config
func (h configHelper) getAttributeConfigIndexedValuePrefix(name string) string {
	_ = "STUB: not implemented"
	return ""
}

// getAttributeConfigOID returns the column OID of an attribute config
func (h configHelper) getAttributeConfigOID(name string) string {
	_ = "STUB: not implemented"
	return ""
}

// getResourceAttributeConfigIndexedValuePrefix returns the indexed value prefix of a resource attribute config
func (h configHelper) getResourceAttributeConfigIndexedValuePrefix(name string) string {
	_ = "STUB: not implemented"
	return ""
}

// getResourceAttributeConfigOID returns the column OID of a resource attribute config
func (h configHelper) getResourceAttributeConfigOID(name string) string {
	_ = "STUB: not implemented"
	return ""
}

// getResourceAttributeConfigScalarOID returns the scalar OID of a resource attribute config
func (h configHelper) getResourceAttributeConfigScalarOID(name string) string {
	_ = "STUB: not implemented"
	return ""
}

// getMetricConfigAttributes returns the metric config attributes for a given OID
func (h configHelper) getMetricConfigAttributes(oid string) []Attribute {
	_ = "STUB: not implemented"
	return nil
}

// getResourceAttributeNames returns the metric config resource attributes for a given OID
func (h configHelper) getResourceAttributeNames(oid string) []string {
	_ = "STUB: not implemented"
	return nil
}

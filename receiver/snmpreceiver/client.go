// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package snmpreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/snmpreceiver"

import (
	"github.com/gosnmp/gosnmp"
	"go.opentelemetry.io/collector/scraper/scrapererror"
	"go.uber.org/zap"
)

type oidDataType byte

// Trimmed down the larger list of gosnmp data types to make retrieving data a
// little more simple when using this client
const (
	notSupportedVal oidDataType = 0x00
	integerVal      oidDataType = 0x01 // value will be int64
	floatVal        oidDataType = 0x02 // value will be float64
	stringVal       oidDataType = 0x03 // value will be string
)

// snmpData used for processFunc and is a simpler version of gosnmp.SnmpPDU
type snmpData struct {
	columnOID string // optional
	oid       string
	value     any
	valueType oidDataType
}

// client is used for retrieving data from a SNMP environment
type client interface {
	// GetScalarData retrieves SNMP scalar data from a list of passed in OIDS,
	// then returns the retrieved data
	GetScalarData(oids []string, scraperErrors *scrapererror.ScrapeErrors) []snmpData
	// GetIndexedData retrieves SNMP indexed data from a list of passed in OIDS,
	// then returns the retrieved data
	GetIndexedData(oids []string, scraperErrors *scrapererror.ScrapeErrors) []snmpData
	// Connect makes a connection to the SNMP host
	Connect() error
	// Close closes a connection to the SNMP host
	Close() error
}

// snmpClient implements the client interface and retrieves data through SNMP
type snmpClient struct {
	client goSNMPWrapper
	logger *zap.Logger
}

// Verify snmpClient implements client interface
var _ client = (*snmpClient)(nil)

// newClient creates an initialized client
// Relies on config being validated thoroughly
func newClient(cfg *Config, logger *zap.Logger) (client, error) {
	_ = "STUB: not implemented"
	// Create goSNMP client
	return *new(client), nil
}

// Set goSNMP version based on config

// Checked in config

// Set goSNMP transport based on config

// Checked in config that it exists

// Set goSNMP target based on config

// Set goSNMP v3 configs

// Set goSNMP community string

// return client

// setV3ClientConfigs sets SNMP v3 related configurations on gosnmp client based on config
func setV3ClientConfigs(client goSNMPWrapper, cfg *Config) { _ = "STUB: not implemented"; return }

// Set goSNMP user based on config

// Set goSNMP security level & auth/privacy details based on config

// getAuthProtocol gets gosnmp auth protocol based on config auth type
func getAuthProtocol(authType string) gosnmp.SnmpV3AuthProtocol {
	_ = "STUB: not implemented"
	return *new(gosnmp.SnmpV3AuthProtocol)
}

// getPrivacyProtocol gets gosnmp privacy protocol based on config privacy type
func getPrivacyProtocol(privacyType string) gosnmp.SnmpV3PrivProtocol {
	_ = "STUB: not implemented"
	return *new(gosnmp.SnmpV3PrivProtocol)
}

// Connect uses the goSNMP client's connect
func (c *snmpClient) Connect() error { _ = "STUB: not implemented"; return nil }

// Close uses the goSNMP client's close
func (c *snmpClient) Close() error { _ = "STUB: not implemented"; return nil }

// GetScalarData retrieves and returns scalar data from passed in scalar OIDs.
// Note: These OIDs must all end in ".0" for the SNMP GET to work correctly
func (c *snmpClient) GetScalarData(oids []string, scraperErrors *scrapererror.ScrapeErrors) []snmpData {
	_ = "STUB: not implemented"
	return nil

	// Nothing to do if there are no OIDs
}

// Group OIDs into chunks based on the max amount allowed in a single SNMP GET

// For each group of OIDs

// Note: Not implementing GetBulk as I don't think it would work correctly for the current design

// Prevent getting stuck in a failure where we can't recover

// For each piece of data in a returned packet

// If there is no value, then ignore

// Convert data into the more simplified data type

// If the value type is not supported, then ignore

// Add the data to be returned

// GetIndexedData retrieves indexed metrics from passed in column OIDs. The returned data
// is then also passed into the provided function.
func (c *snmpClient) GetIndexedData(oids []string, scraperErrors *scrapererror.ScrapeErrors) []snmpData {
	_ = "STUB: not implemented"
	return nil

	// Nothing to do if there are no OIDs
}

// For each column based OID

// Call the correct gosnmp Walk function based on SNMP version

// Allows for quicker recovery rather than timing out for each WALK OID and waiting for the next GET to fix it

// If there is no value, then stop processing

// Convert data into the more simplified data type

// Keep track of which column OID this data came from as well

// If the value type is not supported, then ignore

// Add the data to be returned

// chunkArray takes an initial array and splits it into a number of smaller
// arrays of a given size.
func chunkArray(initArray []string, chunkSize int) [][]string {
	_ = "STUB: not implemented"
	return nil
}

// convertSnmpPDUToSnmpData takes a piece of SnmpPDU data and converts it to the
// client's snmpData type.
func (c *snmpClient) convertSnmpPDUToSnmpData(pdu gosnmp.SnmpPDU) snmpData {
	_ = "STUB: not implemented"
	return *new(snmpData)
}

// Condense gosnmp data types to our client's simplified data types

// Integer types

// String types

// Float types

// Not supported types either because gosnmp doesn't support them
// or they are a type that doesn't translate well to OTEL

// toInt64 converts SnmpPDU.Value to int64, or returns an error for
// non int-like types or a uint64.
//
// This is a convenience function to make working with SnmpPDU's easier - it
// reduces the need for type assertions. A int64 is convenient, as SNMP can
// return int32, uint32, and int64.
func (*snmpClient) toInt64(name string, value any) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
	// shadow
}

// toFloat64 converts SnmpPDU.Value to float64, or returns an error for non
// float like types.
//
// This is a convenience function to make working with SnmpPDU's easier - it
// reduces the need for type assertions. A float64 is convenient, as SNMP can
// return float32 and float64.
func (*snmpClient) toFloat64(name string, value any) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
	// shadow
}

// for testing and other apps - numbers may appear as strings

// toString converts SnmpPDU.Value to string
//
// This is a convenience function to make working with SnmpPDU's easier - it
// reduces the need for type assertions.
func toString(value any) string { _ = "STUB: not implemented"; return "" }

// shadow

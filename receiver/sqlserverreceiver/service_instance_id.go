// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package sqlserverreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/sqlserverreceiver"

const defaultSQLServerPort = 1433

// isLocalhost checks if the given host is a local address
func isLocalhost(host string) bool { _ = "STUB: not implemented"; return false }

// computeServiceInstanceID computes the service.instance.id based on the configuration
// Format: <host>:<port>
// Special handling:
// - localhost/127.0.0.1 are replaced with os.Hostname()
// - Port 0 defaults to 1433
func computeServiceInstanceID(cfg *Config) (string, error) {
	_ = "STUB: not implemented"
	return "", nil

	// Parse connection details based on configuration priority
}

// No server specified, use hostname with default port

// Replace localhost with actual hostname

// Apply default port if not specified

// parseDataSource extracts server and port from SQL Server connection string
// Uses the microsoft/go-mssqldb library's built-in parser for accurate parsing
func parseDataSource(dataSource string) (string, int, error) {
	_ = "STUB: not implemented"
	return "", 0, nil
}

// Parse the connection string using the go-mssqldb library

// Apply default port if not specified

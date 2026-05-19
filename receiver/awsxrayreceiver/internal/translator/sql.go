// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package translator // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsxrayreceiver/internal/translator"

import (
	"regexp"

	"go.opentelemetry.io/collector/pdata/pcommon"

	awsxray "github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/xray"
)

func addSQLToSpan(sql *awsxray.SQLData, attrs pcommon.Map) error {
	_ = "STUB: not implemented"
	return nil
}

// https://github.com/open-telemetry/opentelemetry-collector-contrib/blob/c615d2db351929b99e46f7b427f39c12afe15b54/exporter/awsxrayexporter/translator/sql.go#L60

// not handling sql.ConnectionString for now because the X-Ray exporter
// does not support it

// SQL URL is of the format: protocol+transport://host:port/dbName?queryParam or protocol+transport:dbName?queryParam
var re = regexp.MustCompile(`^([^/]+:(?://[^/]+/)?)([^\?]+)\??.*$`)

const (
	dbURLI  = 1
	dbNameI = 2
)

func splitSQLURL(rawURL string) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

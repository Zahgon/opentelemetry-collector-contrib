// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package parseutils // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal/parseutils"

import (
	"net/url"
)

const (
	// replace once conventions includes these
	AttributeURLUserInfo = "url.user_info"
	AttributeURLUsername = "url.username"
	AttributeURLPassword = "url.password"
)

// parseURI takes an absolute or relative uri and returns the parsed values.
func ParseURI(value string, semconvCompliant bool) (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// remove the query string '?' prefix before parsing

// urlToMap converts a url.URL to a map, excludes any values that are not set.
func urlToSemconvMap(parsedURI *url.URL, m map[string]any) (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// urlToMap converts a url.URL to a map, excludes any values that are not set.
func urlToMap(p *url.URL, m map[string]any) (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// queryToMap converts a query string url.Values to a map.
func queryToMap(query url.Values, m map[string]any) map[string]any {
	_ = "STUB: not implemented"
	// no-op if query is empty, do not create the key m["query"]
	return nil
}

/* 'parameter' will represent url.Values
map[string]any{
	"parameter-a": []any{
		"a",
		"b",
	},
	"parameter-b": []any{
		"x",
		"y",
	},
}
*/

// queryParamValuesToMap takes query string parameter values and
// returns an []interface populated with the values
func queryParamValuesToMap(values []string) []any { _ = "STUB: not implemented"; return nil }

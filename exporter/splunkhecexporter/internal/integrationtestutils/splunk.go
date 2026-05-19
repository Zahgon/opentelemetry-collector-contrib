// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package integrationtestutils // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/splunkhecexporter/internal/integrationtestutils"

func CheckEventsFromSplunk(searchQuery, startTime string, endTimeOptional ...string) []any {
	_ = "STUB: not implemented"
	return nil
}

// post search

// wait for search status done == true
// limit loop - not allowing infinite looping

// get events

func getSplunkSearchResults(user, password, baseURL, jobID string) []any {
	_ = "STUB: not implemented"
	return nil
}

// logger.Println("json Response Events --->")   # debug
// logger.Println(jsonResponseEvents)			# debug

// logger.Println(results)

func checkSearchJobStatusCode(user, password, baseURL, jobID string) any {
	_ = "STUB: not implemented"
	return *new(any)
}

// logger.Println(checkJSONResponse) // debug
// Print isDone field from response

func postSearchRequest(user, password, baseURL, searchQuery, startTime, endTime string) string {
	_ = "STUB: not implemented"
	return ""
}

// debug

func CheckMetricsFromSplunk(index, metricName string) []any { _ = "STUB: not implemented"; return nil }

// logger.Println(events) // debug

func CreateAnIndexInSplunk(index, indexType string) { _ = "STUB: not implemented"; return }

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package parseutils // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal/parseutils"

// ReadCSVRow reads a CSV row from the csv reader, returning the fields parsed from the line.
// We make the assumption that the payload we are reading is a single row, so we allow newline characters in fields.
// However, the csv package does not support newlines in a CSV field (it assumes rows are newline separated),
// so in order to support parsing newlines in a field, we need to stitch together the results of multiple Read calls.
func ReadCSVRow(row string, delimiter rune, lazyQuotes bool) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// -1 indicates a variable length of fields

// If the input is empty, we might not get any lines

/*
	This parser is parsing a single value, which came from a single log entry.
	Therefore, if there are multiple lines here, it should be assumed that each
	subsequent line contains a continuation of the last field in the previous line.

	Given a file w/ headers "A,B,C,D,E" and contents "aa,b\nb,cc,d\nd,ee",
	expect reader.Read() to return bodies:
	- ["aa","b"]
	- ["b","cc","d"]
	- ["d","ee"]
*/

// The first element of the next line is a continuation of the previous line's last element

// The remainder are separate elements

// MapCSVHeaders creates a map of headers[i] -> fields[i].
func MapCSVHeaders(headers, fields []string) (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

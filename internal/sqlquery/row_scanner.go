// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package sqlquery // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/sqlquery"

import (
	"errors"
)

var ErrNullValueWarning = errors.New("NULL value")

type rowScanner struct {
	cols       map[string]func() (string, error)
	scanTarget []any
}

func newRowScanner(colTypes []colType) *rowScanner { _ = "STUB: not implemented"; return nil }

// The Postgres driver returns a []uint8 (ascii string) for decimal and numeric types,
// which we want to render as strings. e.g. "4.1" instead of "[52, 46, 49]".
// Other slice types get the same treatment.

// turn whatever we got from the database driver into a string

func (rs *rowScanner) scan(sqlRows rows) error { _ = "STUB: not implemented"; return nil }

func (rs *rowScanner) toStringMap() (StringMap, error) {
	_ = "STUB: not implemented"
	return *new(StringMap), nil
}

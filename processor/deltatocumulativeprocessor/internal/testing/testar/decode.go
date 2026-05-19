// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// testar is a textual archive (based on [golang.org/x/tools/txtar]) to define
// test fixtures.
//
// Archive data is read into struct fields, optionally calling parsers for field
// types other than string or []byte:
//
//	type T struct {
//		Literal string `testar:"file1"`
//		Parsed  int    `testar:"file2,myparser"`
//	}
//
//	var into T
//	err := Read(data, &into)
//
// See [Read] and [Parser] for examples.
package testar // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/deltatocumulativeprocessor/internal/testing/testar"

import (
	"golang.org/x/tools/txtar"
)

// Read archive data into the fields of struct *T
func Read[T any](data []byte, into *T, parsers ...Format) error {
	_ = "STUB: not implemented"
	return nil
}

func ReadFile[T any](file string, into *T, parsers ...Format) error {
	_ = "STUB: not implemented"
	return nil
}

func Decode[T any](ar *txtar.Archive, into *T, parsers ...Format) error {
	_ = "STUB: not implemented"
	return nil
}

type formats []Format

func (fmts formats) Parse(name string, data []byte, into any) error {
	_ = "STUB: not implemented"
	return nil
}

type Format struct {
	name  string
	parse func(file []byte, into any) error
}

func Parser[T any](name string, fn func([]byte, *T) error) Format {
	_ = "STUB: not implemented"
	return *new(Format)
}

// LiteralParser sets data unaltered into a []byte or string
func LiteralParser(data []byte, into any) error { _ = "STUB: not implemented"; return nil }

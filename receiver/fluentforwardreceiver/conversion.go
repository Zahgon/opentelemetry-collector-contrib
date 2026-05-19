// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package fluentforwardreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/fluentforwardreceiver"

import (
	"time"

	"github.com/tinylib/msgp/msgp"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
)

const tagAttributeKey = "fluent.tag"

// Most of this logic is derived directly from
// https://github.com/fluent/fluentd/wiki/Forward-Protocol-Specification-v1,
// which describes the fields in much greater detail.

type event interface {
	DecodeMsg(dc *msgp.Reader) error
	LogRecords() plog.LogRecordSlice
	Chunk() string
	Compressed() string
}

type optionsMap map[string]any

// Chunk returns the `chunk` option or blank string if it was not set.
func (om optionsMap) Chunk() string { _ = "STUB: not implemented"; return "" }

func (om optionsMap) Compressed() string { _ = "STUB: not implemented"; return "" }

type eventMode int

type peeker interface {
	Peek(n int) ([]byte, error)
}

// Values for enum eventMode.
const (
	unknownMode eventMode = iota
	messageMode
	forwardMode
	packedForwardMode
)

func (em eventMode) String() string { _ = "STUB: not implemented"; return "" }

// parseInterfaceToMap takes map of interface objects and returns
// AttributeValueMap
func parseInterfaceToMap(msi map[string]any, dest pcommon.Value) { _ = "STUB: not implemented"; return }

// parseInterfaceToArray takes array of interface objects and returns
// AttributeValueArray
func parseInterfaceToArray(ai []any, dest pcommon.Value) { _ = "STUB: not implemented"; return }

// parseToAttributeValue converts interface object to AttributeValue
func parseToAttributeValue(val any, dest pcommon.Value) {
	_ = "STUB: not implemented"
	// See https://github.com/tinylib/msgp/wiki/Type-Mapping-Rules
	return
}

// handle overflow of uint64 to int64

// Sometimes strings come in as bytes array

func timeFromTimestamp(ts any) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

func parseRecordToLogRecord(dc *msgp.Reader, lr plog.LogRecord) error {
	_ = "STUB: not implemented"
	return nil
}

// The protocol doesn't specify this but apparently some map keys
// can be binary type instead of string

// fluentd uses message, fluentbit log.

type messageEventLogRecord struct {
	plog.LogRecordSlice
	optionsMap
}

func (melr *messageEventLogRecord) LogRecords() plog.LogRecordSlice {
	_ = "STUB: not implemented"
	return *new(plog.LogRecordSlice)
}

func (melr *messageEventLogRecord) DecodeMsg(dc *msgp.Reader) error {
	_ = "STUB: not implemented"
	return nil
}

func parseOptions(dc *msgp.Reader) (optionsMap, error) {
	_ = "STUB: not implemented"
	return *new(optionsMap), nil
}

type forwardEventLogRecords struct {
	plog.LogRecordSlice
	optionsMap
}

func (fe *forwardEventLogRecords) LogRecords() plog.LogRecordSlice {
	_ = "STUB: not implemented"
	return *new(plog.LogRecordSlice)
}

func (fe *forwardEventLogRecords) DecodeMsg(dc *msgp.Reader) error {
	_ = "STUB: not implemented"
	return nil
}

func parseEntryToLogRecord(dc *msgp.Reader, lr plog.LogRecord) error {
	_ = "STUB: not implemented"
	return nil
}

type packedForwardEventLogRecords struct {
	plog.LogRecordSlice
	optionsMap
}

func (pfe *packedForwardEventLogRecords) LogRecords() plog.LogRecordSlice {
	_ = "STUB: not implemented"
	return *new(plog.LogRecordSlice)
}

// DecodeMsg implements msgp.Decodable.  This was originally code generated but
// then manually copied here in order to handle the optional Options field.
func (pfe *packedForwardEventLogRecords) DecodeMsg(dc *msgp.Reader) error {
	_ = "STUB: not implemented"
	return nil
}

// We have to read out the entries raw all the way first because we don't
// know whether it is compressed or not until we read the options map which
// comes after.  I guess we could use some kind of detection logic to
// determine if it is gzipped by peeking and just ignoring options, but
// this seems simpler for now.

func (pfe *packedForwardEventLogRecords) parseEntries(entriesRaw []byte, isGzipped bool, tag string) error {
	_ = "STUB: not implemented"
	return nil
}

// Allocate only once, since the MoveTo cleans the lr, so we can reuse.

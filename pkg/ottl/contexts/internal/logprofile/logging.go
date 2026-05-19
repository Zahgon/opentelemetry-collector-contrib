// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package logprofile // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/internal/logprofile"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pprofile"
	"go.uber.org/zap/zapcore"
)

func getMapping(dict pprofile.ProfilesDictionary, idx int32) (mapping, error) {
	_ = "STUB: not implemented"
	return *new(mapping), nil
}

func getLocations(dict pprofile.ProfilesDictionary, stackIDx int32) (locations, error) {
	_ = "STUB: not implemented"
	return *new(locations), nil
}

func getFunction(dict pprofile.ProfilesDictionary, idx int32) (function, error) {
	_ = "STUB: not implemented"
	return *new(function), nil
}

func getLink(dict pprofile.ProfilesDictionary, idx int32) (link, error) {
	_ = "STUB: not implemented"
	return *new(link), nil
}

func getString(dict pprofile.ProfilesDictionary, idx int32) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func getAttribute(dict pprofile.ProfilesDictionary, idx int32) (attribute, error) {
	_ = "STUB: not implemented"
	return *new(attribute), nil
}

// Is there a better way to marshal the value?

type Profile struct {
	pprofile.Profile
	Dictionary pprofile.ProfilesDictionary
}

func (p Profile) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

type ProfileSample struct {
	pprofile.Sample
	Profile    pprofile.Profile
	Dictionary pprofile.ProfilesDictionary
}

func (s ProfileSample) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

type valueType struct {
	typ                    string
	unit                   string
	aggregationTemporality int32
}

func newValueType(p Profile, vt pprofile.ValueType) (valueType, error) {
	_ = "STUB: not implemented"
	return *new(valueType), nil
}

func (vt valueType) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

type locations []location

func (s locations) MarshalLogArray(encoder zapcore.ArrayEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

type location struct {
	mapping    mapping
	address    uint64
	lines      lines
	attributes attributes
}

func newLocation(dict pprofile.ProfilesDictionary, pl pprofile.Location) (location, error) {
	_ = "STUB: not implemented"
	return *new(location), nil
}

// optional

func (l location) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

type mapping struct {
	filename    string
	memoryStart uint64
	memoryLimit uint64
	fileOffset  uint64
}

func newMapping(dict pprofile.ProfilesDictionary, pm pprofile.Mapping) (mapping, error) {
	_ = "STUB: not implemented"
	return *new(mapping), nil
}

func (m mapping) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

type link struct {
	pprofile.Link
}

func (m link) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

type attributes []attribute

func (s attributes) MarshalLogArray(encoder zapcore.ArrayEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

func newAttributes(dict pprofile.ProfilesDictionary, pattrs pcommon.Int32Slice) (attributes, error) {
	_ = "STUB: not implemented"
	return *new(attributes), nil
}

type attribute struct {
	key   string
	value string
}

func (a attribute) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

type lines []line

func (s lines) MarshalLogArray(encoder zapcore.ArrayEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

func newLines(dict pprofile.ProfilesDictionary, plines pprofile.LineSlice) (lines, error) {
	_ = "STUB: not implemented"
	return *new(lines), nil
}

type line struct {
	function function
	line     int64
	column   int64
}

func newLine(dict pprofile.ProfilesDictionary, pl pprofile.Line) (line, error) {
	_ = "STUB: not implemented"
	return *new(line), nil
}

func (l line) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

type function struct {
	name       string
	systemName string
	filename   string
	startLine  int64
}

func newFunction(dict pprofile.ProfilesDictionary, pf pprofile.Function) (function, error) {
	_ = "STUB: not implemented"
	return *new(function), nil
}

func (f function) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

type timestamps []timestamp

func (s timestamps) MarshalLogArray(encoder zapcore.ArrayEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

func newTimestamps(ptimestamps pcommon.UInt64Slice) timestamps {
	_ = "STUB: not implemented"
	return *new(timestamps)
}

type timestamp uint64

func (l timestamp) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

type values []value

func (s values) MarshalLogArray(encoder zapcore.ArrayEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

func newValues(pvalues pcommon.Int64Slice) values { _ = "STUB: not implemented"; return *new(values) }

type value int64

func (v value) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

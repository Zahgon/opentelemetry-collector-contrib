// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package faro // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/faro"

import (
	faroTypes "github.com/grafana/faro/pkg/go"
	om "github.com/wk8/go-ordered-map/v2"
)

// keyVal is an ordered map of string to interface
type keyVal = om.OrderedMap[string, any]

// newKeyVal creates new empty keyVal
func newKeyVal() *keyVal { _ = "STUB: not implemented"; return nil }

// keyValFromMap will instantiate keyVal from a map[string]string
func keyValFromMap(m map[string]string) *keyVal { _ = "STUB: not implemented"; return nil }

// keyValFromFloatMap will instantiate keyVal from a map[string]float64
func keyValFromFloatMap(m map[string]float64) *keyVal { _ = "STUB: not implemented"; return nil }

// mergeKeyVal will merge source in target
func mergeKeyVal(target, source *keyVal) { _ = "STUB: not implemented"; return }

// mergeKeyValWithPrefix will merge source in target, adding a prefix to each key being merged in
func mergeKeyValWithPrefix(target, source *keyVal, prefix string) {
	_ = "STUB: not implemented"
	return
}

// keyValAdd adds a key + value string pair to kv
func keyValAdd(kv *keyVal, key, value string) { _ = "STUB: not implemented"; return }

// keyValToInterfaceSlice converts keyVal to []interface{}, typically used for logging
func keyValToInterfaceSlice(kv *keyVal) []any { _ = "STUB: not implemented"; return nil }

// logToKeyVal represents a Log object as keyVal
func logToKeyVal(l *faroTypes.Log) *keyVal {
	_ = "STUB: not implemented"

	// default to info level, prioritize log level if set
	return nil
}

// exceptionToKeyVal represents an Exception object as keyVal
func exceptionToKeyVal(e *faroTypes.Exception) *keyVal { _ = "STUB: not implemented"; return nil }

// exceptionMessage string is concatenating of the Exception.Type and Exception.Value
func exceptionMessage(e *faroTypes.Exception) string { _ = "STUB: not implemented"; return "" }

// exceptionToString is the string representation of an Exception
func exceptionToString(e *faroTypes.Exception) string { _ = "STUB: not implemented"; return "" }

// frameToString function converts a Frame into a human readable string
func frameToString(frame *faroTypes.Frame) string { _ = "STUB: not implemented"; return "" }

// measurementToKeyVal representation of the measurement object
func measurementToKeyVal(m *faroTypes.Measurement) *keyVal { _ = "STUB: not implemented"; return nil }

// eventToKeyVal produces key -> value representation of Event metadata
func eventToKeyVal(e *faroTypes.Event) *keyVal { _ = "STUB: not implemented"; return nil }

// actionToKeyVal produces key->value representation of the Action metadata
func actionToKeyVal(a faroTypes.Action) *keyVal { _ = "STUB: not implemented"; return nil }

// metaToKeyVal produces key->value representation of the metadata
func metaToKeyVal(m faroTypes.Meta) *keyVal { _ = "STUB: not implemented"; return nil }

// sdkToKeyVal produces key->value representation of Sdk metadata
func sdkToKeyVal(sdk faroTypes.SDK) *keyVal { _ = "STUB: not implemented"; return nil }

// sdkIntegrationToString is the string representation of an SDKIntegration
func sdkIntegrationToString(i faroTypes.SDKIntegration) string {
	_ = "STUB: not implemented"
	return ""
}

// appToKeyVal produces key-> value representation of App metadata
func appToKeyVal(a faroTypes.App) *keyVal { _ = "STUB: not implemented"; return nil }

// userToKeyVal produces a key->value representation User metadata
func userToKeyVal(u faroTypes.User) *keyVal { _ = "STUB: not implemented"; return nil }

// sessionToKeyVal produces key->value representation of the Session metadata
func sessionToKeyVal(s faroTypes.Session) *keyVal { _ = "STUB: not implemented"; return nil }

// pageToKeyVal produces key->val representation of Page metadata
func pageToKeyVal(p faroTypes.Page) *keyVal { _ = "STUB: not implemented"; return nil }

// browserToKeyVal produces key->value representation of the Browser metadata
func browserToKeyVal(b faroTypes.Browser) *keyVal { _ = "STUB: not implemented"; return nil }

// deviceToKeyVal produces key->value representation of Device metadata
func deviceToKeyVal(d faroTypes.Device) *keyVal { _ = "STUB: not implemented"; return nil }

// osToKeyVal produces key->value representation of OS metadata
func osToKeyVal(o faroTypes.OS) *keyVal { _ = "STUB: not implemented"; return nil }

// k6ToKeyVal produces a key->value representation K6 metadata
func k6ToKeyVal(k faroTypes.K6) *keyVal { _ = "STUB: not implemented"; return nil }

// viewToKeyVal produces a key->value representation View metadata
func viewToKeyVal(v faroTypes.View) *keyVal { _ = "STUB: not implemented"; return nil }

// geoToKeyVal produces a key->value representation Geo metadata
func geoToKeyVal(g faroTypes.Geo) *keyVal { _ = "STUB: not implemented"; return nil }

// traceToKeyVal produces a key->value representation of the trace context object
func traceToKeyVal(tc faroTypes.TraceContext) *keyVal { _ = "STUB: not implemented"; return nil }

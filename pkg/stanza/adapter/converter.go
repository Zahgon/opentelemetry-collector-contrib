// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package adapter // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/adapter"

import (
	"sync"

	"github.com/cespare/xxhash/v2"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/entry"
)

func ConvertEntries(entries []*entry.Entry) plog.Logs {
	_ = "STUB: not implemented"
	return *new(plog.Logs)
}

// convertInto converts entry.Entry into provided plog.LogRecord.
func convertInto(ent *entry.Entry, dest plog.LogRecord) { _ = "STUB: not implemented"; return }

// The 8 least significant bits are the trace flags as defined in W3C Trace
// Context specification. Don't override the 24 reserved bits.

func upsertToAttributeVal(value any, dest pcommon.Value) { _ = "STUB: not implemented"; return }

func upsertToMap(obsMap map[string]any, dest pcommon.Map) { _ = "STUB: not implemented"; return }

func upsertToSlice(obsArr []any, dest pcommon.Slice) { _ = "STUB: not implemented"; return }

func upsertStringsToSlice(obsArr []string, dest pcommon.Slice) { _ = "STUB: not implemented"; return }

var sevMap = map[entry.Severity]plog.SeverityNumber{
	entry.Default: plog.SeverityNumberUnspecified,
	entry.Trace:   plog.SeverityNumberTrace,
	entry.Trace2:  plog.SeverityNumberTrace2,
	entry.Trace3:  plog.SeverityNumberTrace3,
	entry.Trace4:  plog.SeverityNumberTrace4,
	entry.Debug:   plog.SeverityNumberDebug,
	entry.Debug2:  plog.SeverityNumberDebug2,
	entry.Debug3:  plog.SeverityNumberDebug3,
	entry.Debug4:  plog.SeverityNumberDebug4,
	entry.Info:    plog.SeverityNumberInfo,
	entry.Info2:   plog.SeverityNumberInfo2,
	entry.Info3:   plog.SeverityNumberInfo3,
	entry.Info4:   plog.SeverityNumberInfo4,
	entry.Warn:    plog.SeverityNumberWarn,
	entry.Warn2:   plog.SeverityNumberWarn2,
	entry.Warn3:   plog.SeverityNumberWarn3,
	entry.Warn4:   plog.SeverityNumberWarn4,
	entry.Error:   plog.SeverityNumberError,
	entry.Error2:  plog.SeverityNumberError2,
	entry.Error3:  plog.SeverityNumberError3,
	entry.Error4:  plog.SeverityNumberError4,
	entry.Fatal:   plog.SeverityNumberFatal,
	entry.Fatal2:  plog.SeverityNumberFatal2,
	entry.Fatal3:  plog.SeverityNumberFatal3,
	entry.Fatal4:  plog.SeverityNumberFatal4,
}

var defaultSevTextMap = map[entry.Severity]string{
	entry.Default: "",
	entry.Trace:   "TRACE",
	entry.Trace2:  "TRACE2",
	entry.Trace3:  "TRACE3",
	entry.Trace4:  "TRACE4",
	entry.Debug:   "DEBUG",
	entry.Debug2:  "DEBUG2",
	entry.Debug3:  "DEBUG3",
	entry.Debug4:  "DEBUG4",
	entry.Info:    "INFO",
	entry.Info2:   "INFO2",
	entry.Info3:   "INFO3",
	entry.Info4:   "INFO4",
	entry.Warn:    "WARN",
	entry.Warn2:   "WARN2",
	entry.Warn3:   "WARN3",
	entry.Warn4:   "WARN4",
	entry.Error:   "ERROR",
	entry.Error2:  "ERROR2",
	entry.Error3:  "ERROR3",
	entry.Error4:  "ERROR4",
	entry.Fatal:   "FATAL",
	entry.Fatal2:  "FATAL2",
	entry.Fatal3:  "FATAL3",
	entry.Fatal4:  "FATAL4",
}

// pairSep is chosen to be an invalid byte for a utf-8 sequence
// making it very unlikely to be present in the resource maps keys or values
var pairSep = []byte{0xfe}

// emptyResourceID is the ID returned by HashResource when it is passed an empty resource.
// This specific number is chosen as it is the starting offset of xxHash.
const emptyResourceID uint64 = 17241709254077376921

type hashWriter struct {
	h        *xxhash.Digest
	keySlice []string
}

func newHashWriter() *hashWriter { _ = "STUB: not implemented"; return nil }

var hashWriterPool = &sync.Pool{
	New: func() any { return newHashWriter() },
}

// HashResource will hash an entry.Entry.Resource
func HashResource(resource map[string]any) uint64 { _ = "STUB: not implemented"; return 0 }

// In order for this to be deterministic, we need to sort the map. Using range, like above,
// has no guarantee about order.

//nolint:errcheck // nothing to do about it

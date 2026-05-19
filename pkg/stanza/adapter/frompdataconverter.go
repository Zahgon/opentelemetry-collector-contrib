// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package adapter // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/adapter"

import (
	"sync"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/entry"
)

// FromPdataConverter converts plog.Logs into a set of entry.Entry
//
// The diagram below illustrates the internal communication inside the FromPdataConverter:
//
//	          ┌─────────────────────────────────┐
//	          │ Batch()                         │
//	┌─────────┤  Ingests plog.Logs, splits up   │
//	│         │  and places them on workerChan  │
//	│         └─────────────────────────────────┘
//	│
//	│ ┌───────────────────────────────────────────────────┐
//	├─► workerLoop()                                      │
//	│ │ ┌─────────────────────────────────────────────────┴─┐
//	├─┼─► workerLoop()                                      │
//	│ │ │ ┌─────────────────────────────────────────────────┴─┐
//	└─┼─┼─► workerLoop()                                      │
//	  └─┤ │   consumes sent log entries from workerChan,      │
//	    │ │   translates received logs to entry.Entry,        │
//	    └─┤   and sends them along entriesChan                │
//	      └───────────────────────────────────────────────────┘
type FromPdataConverter struct {
	set component.TelemetrySettings

	// entriesChan is a channel on which converted logs will be sent out of the converter.
	entriesChan chan []*entry.Entry

	stopOnce sync.Once
	stopChan chan struct{}

	// workerChan is an internal communication channel that gets the log
	// entries from Batch() calls and it receives the data in workerLoop().
	workerChan chan fromConverterWorkerItem

	// wg is a WaitGroup that makes sure that we wait for spun up goroutines exit
	// when Stop() is called.
	wg sync.WaitGroup
}

func NewFromPdataConverter(set component.TelemetrySettings, workerCount int) *FromPdataConverter {
	_ = "STUB: not implemented"
	return nil
}

func (c *FromPdataConverter) Start() { _ = "STUB: not implemented"; return }

func (c *FromPdataConverter) Stop() { _ = "STUB: not implemented"; return }

// OutChannel returns the channel on which converted entries will be sent to.
func (c *FromPdataConverter) OutChannel() <-chan []*entry.Entry {
	_ = "STUB: not implemented"
	return nil
}

type fromConverterWorkerItem struct {
	Resource       pcommon.Resource
	LogRecordSlice plog.LogRecordSlice
	Scope          plog.ScopeLogs
}

// workerLoop is responsible for obtaining pdata logs from Batch() calls,
// converting them to []*entry.Entry and sending them out
func (c *FromPdataConverter) workerLoop() { _ = "STUB: not implemented"; return }

// Batch takes in an set of plog.Logs and sends it to an available worker for processing.
func (c *FromPdataConverter) Batch(pLogs plog.Logs) error { _ = "STUB: not implemented"; return nil }

// convertFromLogs converts the contents of a fromConverterWorkerItem into a slice of entry.Entry
func convertFromLogs(workerItem fromConverterWorkerItem) []*entry.Entry {
	_ = "STUB: not implemented"
	return nil
}

// convertFrom converts plog.LogRecord into provided entry.Entry.
func convertFrom(src plog.LogRecord, ent *entry.Entry) {
	_ = "STUB: not implemented"
	// if src.Timestamp == 0, then leave ent.Timestamp as nil
	return
}

var fromPdataSevMap = map[plog.SeverityNumber]entry.Severity{
	plog.SeverityNumberUnspecified: entry.Default,
	plog.SeverityNumberTrace:       entry.Trace,
	plog.SeverityNumberTrace2:      entry.Trace2,
	plog.SeverityNumberTrace3:      entry.Trace3,
	plog.SeverityNumberTrace4:      entry.Trace4,
	plog.SeverityNumberDebug:       entry.Debug,
	plog.SeverityNumberDebug2:      entry.Debug2,
	plog.SeverityNumberDebug3:      entry.Debug3,
	plog.SeverityNumberDebug4:      entry.Debug4,
	plog.SeverityNumberInfo:        entry.Info,
	plog.SeverityNumberInfo2:       entry.Info2,
	plog.SeverityNumberInfo3:       entry.Info3,
	plog.SeverityNumberInfo4:       entry.Info4,
	plog.SeverityNumberWarn:        entry.Warn,
	plog.SeverityNumberWarn2:       entry.Warn2,
	plog.SeverityNumberWarn3:       entry.Warn3,
	plog.SeverityNumberWarn4:       entry.Warn4,
	plog.SeverityNumberError:       entry.Error,
	plog.SeverityNumberError2:      entry.Error2,
	plog.SeverityNumberError3:      entry.Error3,
	plog.SeverityNumberError4:      entry.Error4,
	plog.SeverityNumberFatal:       entry.Fatal,
	plog.SeverityNumberFatal2:      entry.Fatal2,
	plog.SeverityNumberFatal3:      entry.Fatal3,
	plog.SeverityNumberFatal4:      entry.Fatal4,
}

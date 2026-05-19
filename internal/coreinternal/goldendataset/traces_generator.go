// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package goldendataset // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal/goldendataset"

import (
	"io"
	"math/rand/v2"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

// GenerateTraces generates a slice of OTLP ResourceSpans objects based on the PICT-generated pairwise
// parameters defined in the parameters file specified by the tracePairsFile parameter. The pairs to generate
// spans for defined in the file specified by the spanPairsFile parameter.
// The slice of ResourceSpans are returned. If an err is returned, the slice elements will be nil.
func GenerateTraces(tracePairsFile, spanPairsFile string) ([]ptrace.Traces, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: use math/rand/v2.ChaCha8.Read when we upgrade to go1.23.
type randReader rand.Rand

func (r *randReader) Read(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

// generateResourceSpan generates a single PData ResourceSpans populated based on the provided inputs. They are:
//
//	tracingInputs - the pairwise combination of field value variations for this ResourceSpans
//	spanPairsFile - the file with the PICT-generated parameter combinations to generate spans for
//	random - the random number generator to use in generating ID values
//
// The generated resource spans. If err is not nil, some or all of the resource spans fields will be nil.
func appendResourceSpan(tracingInputs *PICTTracingInputs, spanPairsFile string,
	random io.Reader, resourceSpansSlice ptrace.ResourceSpansSlice,
) error {
	_ = "STUB: not implemented"
	return nil
}

func appendScopeSpans(tracingInputs *PICTTracingInputs, spanPairsFile string,
	random io.Reader, scopeSpansSlice ptrace.ScopeSpansSlice,
) error {
	_ = "STUB: not implemented"
	return nil
}

func fillScopeSpans(tracingInputs *PICTTracingInputs, index int, spanPairsFile string, random io.Reader, scopeSpans ptrace.ScopeSpans) error {
	_ = "STUB: not implemented"
	return nil
}

func countTotalSpanCases(spanPairsFile string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func fillInstrumentationLibrary(tracingInputs *PICTTracingInputs, index int, scope pcommon.InstrumentationScope) {
	_ = "STUB: not implemented"
	return
}

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package genainormalizerprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/genainormalizerprocessor"

import (
	"context"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

// valueTransformer applies value-level normalization after an attribute is
// renamed. Sources that do not need value normalization may leave this nil.
// TODO [kylehounslow]: Review interface vs. typed function
type valueTransformer interface {
	TransformValue(targetKey, value string) string
}

// sourceNormalizer holds per-source state used during normalization.
type sourceNormalizer struct {
	lookupTable     map[string]string
	transformValue  valueTransformer
	removeOriginals bool
	overwrite       bool
}

// newSourceNormalizer wires up a sourceNormalizer from a validated Source
// config. Unknown source names produce a no-op normalizer; Config validation
// rejects them upstream.
func newSourceNormalizer(src Source) sourceNormalizer {
	_ = "STUB: not implemented"
	return *new(sourceNormalizer)
}

// genaiNormalizerProcessor normalizes span attributes for each configured source.
type genaiNormalizerProcessor struct {
	sources []sourceNormalizer
}

// newGenaiNormalizerProcessor builds a processor from a validated Config.
// Sources are applied in the order specified in the configuration.
func newGenaiNormalizerProcessor(cfg *Config) *genaiNormalizerProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (p *genaiNormalizerProcessor) processTraces(_ context.Context, td ptrace.Traces) (ptrace.Traces, error) {
	_ = "STUB: not implemented"
	return *new(ptrace.Traces), nil
}

// normalizeAttributes applies the source's rename rules to attrs. It returns
// true if at least one attribute was written.
func (sn *sourceNormalizer) normalizeAttributes(attrs pcommon.Map) bool {
	_ = "STUB: not implemented"
	return false
}

// applyValueTransform runs the per-source value transformer if one is set.
func (sn *sourceNormalizer) applyValueTransform(targetKey, value string) string {
	_ = "STUB: not implemented"
	return ""
}

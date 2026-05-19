// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package probabilisticsamplerprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/probabilisticsamplerprocessor"

import (
	"context"

	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/processor"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/sampling"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/probabilisticsamplerprocessor/internal/metadata"
)

type logsProcessor struct {
	sampler dataSampler

	samplingPriority string
	precision        int
	failClosed       bool
	logger           *zap.Logger
	telemetryBuilder *metadata.TelemetryBuilder
}

type recordCarrier struct {
	record plog.LogRecord

	parsed struct {
		tvalue    string
		threshold sampling.Threshold

		rvalue     string
		randomness sampling.Randomness
	}
}

var _ samplingCarrier = &recordCarrier{}

func (rc *recordCarrier) get(key string) string { _ = "STUB: not implemented"; return "" }

func newLogRecordCarrier(l plog.LogRecord) (samplingCarrier, error) {
	_ = "STUB: not implemented"
	return *new(samplingCarrier), nil
}

func (rc *recordCarrier) threshold() (sampling.Threshold, bool) {
	_ = "STUB: not implemented"
	return *new(sampling.Threshold), false
}

func (rc *recordCarrier) explicitRandomness() (randomnessNamer, bool) {
	_ = "STUB: not implemented"
	return *new(randomnessNamer), false
}

func (rc *recordCarrier) updateThreshold(th sampling.Threshold) error {
	_ = "STUB: not implemented"
	return nil
}

func (rc *recordCarrier) setExplicitRandomness(rnd randomnessNamer) {
	_ = "STUB: not implemented"
	return
}

func (rc *recordCarrier) clearThreshold() { _ = "STUB: not implemented"; return }

func (*recordCarrier) reserialize() error { _ = "STUB: not implemented"; return nil }

func (*neverSampler) randomnessFromLogRecord(logRec plog.LogRecord) (randomnessNamer, samplingCarrier, error) {
	_ = "STUB: not implemented"
	// We return a fake randomness value, since it will not be used.
	// This avoids a consistency check error for missing randomness.
	return *new(randomnessNamer), *new(samplingCarrier), nil
}

// randomnessFromLogRecord (hashingSampler) uses a hash function over
// the TraceID or logs attribute source.
func (th *hashingSampler) randomnessFromLogRecord(logRec plog.LogRecord) (randomnessNamer, samplingCarrier, error) {
	_ = "STUB: not implemented"
	return *new(randomnessNamer), *new(samplingCarrier), nil
}

// The sampling.randomness or sampling.threshold attributes
// had a parse error, in this case.

// If the log record contains a randomness value, do not update.

// If the log record contains a threshold value, do not update.

// When no sampling information is already present and we have
// calculated new randomness, add it to the record.

// randomnessFromLogRecord (hashingSampler) uses OTEP 235 semantic
// conventions basing its decision only on the TraceID.
func (*consistentTracestateCommon) randomnessFromLogRecord(logRec plog.LogRecord) (randomnessNamer, samplingCarrier, error) {
	_ = "STUB: not implemented"
	return *new(randomnessNamer), *new(samplingCarrier), nil
}

// Parse error in sampling.randomness or sampling.threshold

// newLogsProcessor returns a processor.LogsProcessor that will perform head sampling according to the given
// configuration.
func newLogsProcessor(ctx context.Context, set processor.Settings, nextConsumer consumer.Logs, cfg *Config) (processor.Logs, error) {
	_ = "STUB: not implemented"
	return *new(processor.Logs), nil
}

func (lsp *logsProcessor) processLogs(ctx context.Context, logsData plog.Logs) (plog.Logs, error) {
	_ = "STUB: not implemented"
	return *new(plog.Logs), nil
}

// Filter out empty ScopeLogs

// Filter out empty ResourceLogs

func (lsp *logsProcessor) priorityFunc(logRec plog.LogRecord, rnd randomnessNamer, threshold sampling.Threshold) (randomnessNamer, sampling.Threshold) {
	_ = "STUB: not implemented"
	// Note: in logs, unlike traces, the sampling priority
	// attribute is interpreted as a request to be sampled.
	return *new(randomnessNamer), *new(sampling.Threshold)
}

// override policy name

func (lsp *logsProcessor) logRecordToPriorityThreshold(logRec plog.LogRecord) (sampling.Threshold, bool) {
	_ = "STUB: not implemented"
	return *new(sampling.Threshold), false
}

// Potentially raise the sampling probability to minProb

// The record has supplied a valid alternative sampling probability

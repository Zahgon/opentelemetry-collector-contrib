// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package tailsamplingprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/tailsamplingprocessor"

import (
	"container/list"
	"context"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.opentelemetry.io/collector/processor"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/tailsamplingprocessor/cache"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/tailsamplingprocessor/internal/idbatcher"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/tailsamplingprocessor/internal/metadata"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/tailsamplingprocessor/internal/tailstorageextension"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/tailsamplingprocessor/pkg/samplingpolicy"
)

// policy combines a sampling policy evaluator with the destinations to be
// used for that policy.
type policy struct {
	// name used to identify this policy instance.
	name string
	// evaluator that decides if a trace is sampled or not by this policy instance.
	evaluator samplingpolicy.Evaluator
	// attribute to use in the telemetry to denote the policy.
	attribute metric.MeasurementOption
	// isDrop indicates this is a drop policy.
	isDrop bool
}

// TraceData is a wrapper around the publically used samplingpolicy.TraceData
// that tracks information related to the decision making process but not
// needed by any sampler implementations.
type TraceData struct {
	samplingpolicy.TraceData
	FinalDecision samplingpolicy.Decision
	PolicyName    string

	arrivalTime   time.Time
	decisionTime  time.Time
	deleteElement *list.Element
	batchID       uint64
}

type DecisionHook func(ctx context.Context, id pcommon.TraceID, td *TraceData)

type tailSamplingSpanProcessor struct {
	ctx context.Context

	set       processor.Settings
	telemetry *metadata.TelemetryBuilder
	logger    *zap.Logger
	tracer    trace.Tracer

	deleteTraceQueue   *list.List
	nextConsumer       consumer.Traces
	policies           []*policy
	idToTrace          map[pcommon.TraceID]*TraceData
	tailStorage        tailstorageextension.TailStorage
	tickerFrequency    time.Duration
	decisionBatcher    idbatcher.Batcher
	sampledIDCache     cache.Cache
	nonSampledIDCache  cache.Cache
	recordPolicy       bool
	sampleOnFirstMatch bool
	blockOnOverflow    bool
	maxTraceSizeBytes  uint64

	cfg  Config
	host component.Host

	sampledHooks    []DecisionHook
	nonSampledHooks []DecisionHook

	newPolicyChan    chan newPolicyCmd
	newTraceSizeChan chan uint64
	workChan         chan []traceBatch
	doneChan         chan struct{}

	// tickChan triggers ticks and responds on the provided channel when the tick is complete.
	// this is used in tests to produce deterministic ticks.
	tickChan chan chan struct{}
}

func newTracesProcessor(ctx context.Context, set processor.Settings, nextConsumer consumer.Traces, cfg Config) (processor.Traces, error) {
	_ = "STUB: not implemented"
	return *new(processor.Traces), nil
}

// Similar to the id batcher, allow a batch per CPU to be buffered before blocking ConsumeTraces.

// We need to buffer one new policy command/size update so that external callers can
// queue up new policies without blocking on the ticker.

func (*tailSamplingSpanProcessor) Capabilities() consumer.Capabilities {
	_ = "STUB: not implemented"
	return *new(consumer.Capabilities)
}

// Start is invoked during service startup.
func (tsp *tailSamplingSpanProcessor) Start(_ context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

// If the policies are not set, set them. This is only for testing purposes,
// so that withPolicies can inject custom policies.

// this will start a goroutine in the background, so we run it only if everything went
// well in creating the policies, and only when the processor starts.

// ConsumeTraces is required by the processor.Traces interface.
func (tsp *tailSamplingSpanProcessor) ConsumeTraces(ctx context.Context, td ptrace.Traces) error {
	_ = "STUB: not implemented"
	return nil
}

// First group all spans by trace.

// Then, create a new ptrace.ResourceSpans for each trace copying all
// data. Paying this cost upfront allows the inner loop of the TSP to
// be more efficient on its single goroutine.

func (tsp *tailSamplingSpanProcessor) SetSamplingPolicy(cfgs []PolicyCfg) {
	_ = "STUB: not implemented"
	return
}

func (tsp *tailSamplingSpanProcessor) loadSamplingPolicies(host component.Host, cfgs []PolicyCfg) ([]*policy, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Dropped decision takes precedence over all others, therefore we evaluate them first.

func (tsp *tailSamplingSpanProcessor) SetMaximumTraceSizeBytes(size uint64) {
	_ = "STUB: not implemented"
	return
}

// traceBatch contains all spans from a single batch for a single trace.
type traceBatch struct {
	id        pcommon.TraceID
	rootSpan  *ptrace.Span
	rss       ptrace.ResourceSpans
	spanCount int64
}

type newPolicyCmd struct {
	policies []*policy
}

// spanAndScope a structure for holding information about span and its instrumentation scope.
// required for preserving the instrumentation library information while sampling.
// We use pointers there to fast find the span in the map.
type spanAndScope struct {
	span                 *ptrace.Span
	instrumentationScope *pcommon.InstrumentationScope
}

var (
	attrDecisionSampled    = metric.WithAttributes(attribute.String("sampled", "true"), attribute.String("decision", "sampled"))
	attrDecisionNotSampled = metric.WithAttributes(attribute.String("sampled", "false"), attribute.String("decision", "not_sampled"))
	attrDecisionDropped    = metric.WithAttributes(attribute.String("sampled", "false"), attribute.String("decision", "dropped"))
	decisionToAttributes   = map[samplingpolicy.Decision]metric.MeasurementOption{
		samplingpolicy.Sampled:    attrDecisionSampled,
		samplingpolicy.NotSampled: attrDecisionNotSampled,
		//nolint:staticcheck // SA1019: Use of inverted decisions until they are fully removed.
		samplingpolicy.InvertNotSampled: attrDecisionNotSampled,
		//nolint:staticcheck // SA1019: Use of inverted decisions until they are fully removed.
		samplingpolicy.InvertSampled: attrDecisionSampled,
		samplingpolicy.Dropped:       attrDecisionDropped,
	}

	attrSampledTrue  = metric.WithAttributes(attribute.String("sampled", "true"))
	attrSampledFalse = metric.WithAttributes(attribute.String("sampled", "false"))
)

type Option func(*tailSamplingSpanProcessor)

// WithSampledDecisionCache sets the cache which the processor uses to store recently sampled trace IDs.
func WithSampledDecisionCache(c cache.Cache) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithNonSampledDecisionCache sets the cache which the processor uses to store recently non-sampled trace IDs.
func WithNonSampledDecisionCache(c cache.Cache) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithSampledHooks sets hooks to be called when a trace is sampled.
func WithSampledHooks(hooks ...DecisionHook) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithNonSampledHooks sets hooks to be called when a trace is not sampled.
func WithNonSampledHooks(hooks ...DecisionHook) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func withRecordPolicy() Option { _ = "STUB: not implemented"; return *new(Option) }

func getPolicyEvaluator(settings component.TelemetrySettings, cfg *PolicyCfg, policyExtensions map[string]samplingpolicy.Extension) (samplingpolicy.Evaluator, error) {
	_ = "STUB: not implemented"
	return *new(samplingpolicy.Evaluator), nil
}

func getSharedPolicyEvaluator(settings component.TelemetrySettings, cfg *sharedPolicyCfg, policyExtensions map[string]samplingpolicy.Extension) (samplingpolicy.Evaluator, error) {
	_ = "STUB: not implemented"
	return *new(samplingpolicy.Evaluator), nil
}

type policyDecisionMetrics struct {
	tracesSampled int
	spansSampled  int64
}

type policyEvaluationMetrics struct {
	idNotFoundOnMapCount, evaluateErrorCount, decisionSampled, decisionNotSampled, decisionDropped int64
	tracesSampledByPolicyDecision                                                                  []map[samplingpolicy.Decision]policyDecisionMetrics
	cumulativeExecutionTime                                                                        []perPolicyExecutionTime
}

// perPolicyExecutionTime is a struct for holding the cumulative execution time
// and number of executions of a policy. This is an optimization to avoid
// instrumentation overhead in the decision making loop.
type perPolicyExecutionTime struct {
	executionTime  time.Duration
	executionCount int64
}

func newPolicyEvaluationMetrics(numPolicies int) *policyEvaluationMetrics {
	_ = "STUB: not implemented"
	return nil
}

func (m *policyEvaluationMetrics) addDecision(policyIndex int, decision samplingpolicy.Decision, spansSampled int64) {
	_ = "STUB: not implemented"
	return
}

func (m *policyEvaluationMetrics) addDecisionTime(policyIndex int, decisionTime time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (tsp *tailSamplingSpanProcessor) recordPerPolicyEvaluationMetrics(metrics *policyEvaluationMetrics) {
	_ = "STUB: not implemented"
	return
}

func (tsp *tailSamplingSpanProcessor) recordImmediateDecisionMetrics(decision samplingpolicy.Decision, metrics *policyEvaluationMetrics, evaluationLatency time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (tsp *tailSamplingSpanProcessor) loop() { _ = "STUB: not implemented"; return }

func (tsp *tailSamplingSpanProcessor) iter(tickChan <-chan time.Time, workChan <-chan []traceBatch) bool {
	_ = "STUB: not implemented"
	return false
}

// If the context is done then we can't send anything anymore so just exit.

// No more traces to process as we are shutting down. Clear the queue and make decisions for all batches based on the current data.

// Stop the batcher so that we can read all batches without creating new ones.

// Do the best decision we can for any traces we have already ingested unless a user wants to drop them.

// Short circuit if the trace has already been sampled or dropped.

// processCachedTrace checks if a given trace has already been sampled (or
// dropped) and forwards the span appropriately. It returns true if the trace
// was cached and false if regular processing must be done.
func (tsp *tailSamplingSpanProcessor) processCachedTrace(traceID pcommon.TraceID, rss ptrace.ResourceSpans, spanCount int64) bool {
	_ = "STUB: not implemented"
	return false
}

// waitForSpace blocks until space is available. Depending on configuration,
// this might immediately drop data or wait until the sampler tick frees space.
func (tsp *tailSamplingSpanProcessor) waitForSpace(tickChan <-chan time.Time) {
	_ = "STUB: not implemented"
	return

	// Ticks are not guaranteed to drop data, since they may process an
	// empty batch. We loop until we have space for a new trace.
}

// Recursively iter with a nil workChan to wait for space.

// This should be impossible.

// Somehow the trace was already removed from idToTrace, but not the
// queue. Drop the element to avoid an infinite loop.

// samplingPolicyOnTick takes the next batch and process all traces in that batch. Returns if there are more batches in the batcher.
func (tsp *tailSamplingSpanProcessor) samplingPolicyOnTick() bool {
	_ = "STUB: not implemented"
	return false
}

// A decision was already made, no need to do it again. This happens
// when no decision cache is used and a trace was processed both due to
// a root span trigger and after decision_wait.

// In span-ingest mode, tick is a cleanup path only. Finalize any
// still-pending trace as implicit not sampled without policy evaluation.

// Keep release paths working with tail storage by attaching the
// retrieved batches back to trace state for this decision.

// Sampled or not, remove the batches

func (tsp *tailSamplingSpanProcessor) makeDecision(ctx context.Context, id pcommon.TraceID, traceData *samplingpolicy.TraceData, metrics *policyEvaluationMetrics) (samplingpolicy.Decision, string) {
	_ = "STUB: not implemented"
	return *new(samplingpolicy.Decision), ""
}

//nolint:staticcheck // SA1019: Use of inverted decisions until they are fully removed.

//nolint:staticcheck // SA1019: Use of inverted decisions until they are fully removed.

// Check all policies before making a final decision.

// We associate the first policy with the sampling decision to understand what policy sampled a span

// Break early if dropped. This can drastically reduce tick/decision latency.

// If sampleOnFirstMatch is enabled, make decision as soon as a policy matches

// Dropped takes precedence

//nolint:staticcheck // SA1019: Use of inverted decisions until they are fully removed.
// Then InvertNotSampled

//nolint:staticcheck // SA1019: Use of inverted decisions until they are fully removed.

//nolint:staticcheck // SA1019: Use of inverted decisions until they are fully removed.

// makeDecisionOnSpanIngest is used by span-ingest mode. It only returns
// terminal decisions at ingest time. All other outcomes remain pending.
func (tsp *tailSamplingSpanProcessor) makeDecisionOnSpanIngest(id pcommon.TraceID, trace *samplingpolicy.TraceData, metrics *policyEvaluationMetrics) (samplingpolicy.Decision, string) {
	_ = "STUB: not implemented"
	return *new(samplingpolicy.Decision), ""
}

// Track whether a drop policy didn't match this batch. A drop policy
// returning NotSampled means it didn't match these spans but could still
// match future spans for this trace, so we must defer a Sampled decision
// to avoid forwarding spans that should ultimately be dropped.

// If we have moved past all drop policies and one is pending there is
// no use wasting work on evaluation when we can't sample the trace.

// Treat all other results as pending as the result could change.

// Since drop policies are always sorted to the top we know dropPolicyPending is accurate.

func groupSpansByTraceKey(resourceSpans ptrace.ResourceSpans) map[pcommon.TraceID][]spanAndScope {
	_ = "STUB: not implemented"
	return nil
}

func (tsp *tailSamplingSpanProcessor) processTrace(id pcommon.TraceID, rss ptrace.ResourceSpans, spanCount int64, containsRootSpan bool) {
	_ = "STUB: not implemented"
	return
}

// Need a closure here to delay evaluation of newTraceIDs.

// Since we are not in a normal decision flow when dropping large traces, also be sure to remove it from the batcher.

// Build an isolated one-batch trace view for ingest-time evaluation
// using moves to avoid deep-copying span data.

// Release all accumulated spans (prior pending batches + current batch)
// without writing the current batch to storage first.

// Persist current batch for pending traces.
// Use the moved batch from spanIngestTraceData (rss has been moved).

// If the final decision hasn't been made, add the new spans to the
// existing trace.

// TODO: I don't think this is correct? If it isn't sampled shouldn't we just do nothing?

func extensions(host component.Host) map[string]samplingpolicy.Extension {
	_ = "STUB: not implemented"
	return nil
}

func tailStorageExtension(host component.Host, storageID component.ID) (tailstorageextension.TailStorage, error) {
	_ = "STUB: not implemented"
	return *new(tailstorageextension.TailStorage), nil
}

// Shutdown is invoked during service shutdown.
func (tsp *tailSamplingSpanProcessor) Shutdown(context.Context) error {
	_ = "STUB: not implemented"
	// All receivers will be shutdown before processors so no sends will be done anymore.
	return nil
}

// dropTrace removes the trace from all memory locations. Returns true if it was removed and false if not found.
func (tsp *tailSamplingSpanProcessor) dropTrace(traceID pcommon.TraceID, deletionTime time.Time) bool {
	_ = "STUB: not implemented"
	return false
}

// forwardSpans sends the trace data to the next consumer. it is different from
// releaseSampledTrace in that it does not modify any tsp state.
func (tsp *tailSamplingSpanProcessor) forwardSpans(ctx context.Context, td ptrace.Traces) {
	_ = "STUB: not implemented"
	return
}

// releaseSampledTrace sends the trace data to the next consumer. It
// additionally adds the trace ID to the cache of sampled trace IDs. If the
// trace ID is cached, it deletes the spans from the internal map.
func (tsp *tailSamplingSpanProcessor) releaseSampledTrace(ctx context.Context, id pcommon.TraceID, td *TraceData) {
	_ = "STUB: not implemented"
	return
}

// releaseNotSampledTrace adds the trace ID to the cache of not sampled trace
// IDs. If the trace ID is cached, it deletes the spans from the internal map.
func (tsp *tailSamplingSpanProcessor) releaseNotSampledTrace(id pcommon.TraceID, td *TraceData) {
	_ = "STUB: not implemented"
	return
}

func getPolicyName(policy *policy) string { _ = "STUB: not implemented"; return "" }

func appendToTraces(dest ptrace.Traces, rss ptrace.ResourceSpans) {
	_ = "STUB: not implemented"
	return
}

func appendAllTraces(dest, src ptrace.Traces) { _ = "STUB: not implemented"; return }

func newResourceSpanFromSpanAndScopes(rss ptrace.ResourceSpans, spanAndScopes []spanAndScope) (ptrace.ResourceSpans, *ptrace.Span) {
	_ = "STUB: not implemented"
	return *new(ptrace.ResourceSpans), nil
}

// If the scope of the spanAndScope is not in the map, add it to the map and the destination.

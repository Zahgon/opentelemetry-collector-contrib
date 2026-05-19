// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package githubreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/githubreceiver"

import (
	"time"

	"github.com/google/go-github/v86/github"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

func (gtr *githubTracesReceiver) handleWorkflowRun(e *github.WorkflowRunEvent, rawPayload []byte) (ptrace.Traces, error) {
	_ = "STUB: not implemented"
	return *new(ptrace.Traces), nil
}

// handleWorkflowJob handles the creation of spans for a GitHub Workflow Job
// events, including the underlying steps within each job. A `job` maps to the
// semantic conventions for a `cicd.pipeline.task`.
func (gtr *githubTracesReceiver) handleWorkflowJob(e *github.WorkflowJobEvent, rawPayload []byte) (ptrace.Traces, error) {
	_ = "STUB: not implemented"
	return *new(ptrace.Traces), nil
}

// newTraceID creates a deterministic Trace ID based on the provided inputs of
// runID and runAttempt. `t` is appended to the end of the input to
// differentiate between a deterministic traceID and the parentSpanID.
func newTraceID(runID int64, runAttempt int) (pcommon.TraceID, error) {
	_ = "STUB: not implemented"
	return *new(pcommon.TraceID), nil
}

// TODO: Determine if this is the best hashing algorithm to use. This is
// more likely to generate a unique hash compared to MD5 or SHA1. Could
// alternatively use UUID library to generate a unique ID by also using a
// hash.

// newParentId creates a deterministic Parent Span ID based on the provided
// runID and runAttempt. `s` is appended to the end of the input to
// differentiate between a deterministic traceID and the parentSpanID.
func newParentSpanID(runID int64, runAttempt int) (pcommon.SpanID, error) {
	_ = "STUB: not implemented"
	return *new(pcommon.SpanID), nil
}

// correctActionTimestamps ensures span timestamps are valid by checking that
// the end time is not before the start time. When GitHub reports timestamps
// in reverse order (which can occur with skipped jobs and steps), this function
// returns corrected timestamps where both start and end are set to the later
// timestamp, resulting in a zero-duration span.
//
// This prevents negative durations that would otherwise appear as excessively
// long spans in telemetry systems.
func correctActionTimestamps(start, end time.Time) (time.Time, time.Time) {
	_ = "STUB: not implemented"
	return *

	// Use the later timestamp (start) for both, creating zero-duration span
	new(time.Time), *new(time.Time)
}

// createRootSpan creates a root span based on the provided event, associated
// with the deterministic traceID.
func (gtr *githubTracesReceiver) createRootSpan(
	resourceSpans ptrace.ResourceSpans,
	event *github.WorkflowRunEvent,
	traceID pcommon.TraceID,
	rawPayload []byte,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Attach raw event as span event if configured

// Attempt to link to previous trace ID if applicable

// createParentSpan creates a parent span based on the provided event, associated
// with the deterministic traceID.
func (gtr *githubTracesReceiver) createParentSpan(
	resourceSpans ptrace.ResourceSpans,
	event *github.WorkflowJobEvent,
	traceID pcommon.TraceID,
	rawPayload []byte,
) (pcommon.SpanID, error) {
	_ = "STUB: not implemented"
	return *new(pcommon.SpanID), nil
}

// Attach raw event as span event if configured

// newJobSpanID creates a deterministic Job Span ID based on the provided runID,
// runAttempt, and the name of the job.
//
// Deprecated: superseded by newSpanIDFromCheckRun, which is used when the
// receiver.githubreceiver.UseCheckRunID feature gate is enabled (the default as
// of v0.151.0). This function and its callers are scheduled for removal when
// the gate is promoted to Stable. See
// https://github.com/open-telemetry/opentelemetry-collector-contrib/issues/44856.
func newJobSpanID(runID int64, runAttempt int, jobName string) (pcommon.SpanID, error) {
	_ = "STUB: not implemented"
	return *new(pcommon.SpanID), nil
}

// createStepSpans is a wrapper function to create spans for each step in the
// the workflow job by identifying duplicate names then creating a span for each
// step.
func (gtr *githubTracesReceiver) createStepSpans(
	resourceSpans ptrace.ResourceSpans,
	event *github.WorkflowJobEvent,
	traceID pcommon.TraceID,
	parentSpanID pcommon.SpanID,
) error {
	_ = "STUB: not implemented"
	return nil
}

// newUniqueSteps creates a new slice of step names from the provided GitHub
// event steps. Each step name, if duplicated, is appended with `-n` where n is
// the numbered occurrence. The second return value reports whether any two
// steps shared the same raw name; callers use this to warn operators when
// UseCheckRunID is enabled and the duplicates would cause colliding step span
// IDs.
func newUniqueSteps(steps []*github.TaskStep) ([]string, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// createStepSpan creates a span with a deterministic spandID for the provided
// step.
func (*githubTracesReceiver) createStepSpan(
	resourceSpans ptrace.ResourceSpans,
	traceID pcommon.TraceID,
	parentSpanID pcommon.SpanID,
	event *github.WorkflowJobEvent,
	step *github.TaskStep,
	name string,
) error {
	_ = "STUB: not implemented"
	return nil
}

// newStepSpanID creates a deterministic Step Span ID based on the provided
// inputs.
//
// Deprecated: superseded by newSpanIDFromCheckRun, which is used when the
// receiver.githubreceiver.UseCheckRunID feature gate is enabled (the default as
// of v0.151.0). This function and its callers are scheduled for removal when
// the gate is promoted to Stable. See
// https://github.com/open-telemetry/opentelemetry-collector-contrib/issues/44856.
func newStepSpanID(runID int64, runAttempt int, jobName, stepName string, number int) (pcommon.SpanID, error) {
	_ = "STUB: not implemented"
	return *new(pcommon.SpanID), nil
}

// newSpanIDFromCheckRun creates span IDs from input. input is a formatted
// string based off GitHub information with an appended value to enable the
// creation of deterministic IDs.
//
// input is as follows:
// jobSpanID := fmt.Sprintf("%d-j", checkRunID)
// queueJobSpanID := fmt.Sprintf("%d-q", checkRunID)
// stepSpanID := fmt.Sprintf("%d-%s-s", checkRunID, stepName)
func newSpanIDFromCheckRun(input string) (pcommon.SpanID, error) {
	_ = "STUB: not implemented"
	return *new(pcommon.SpanID), nil
}

// createJobQueueSpan creates a span for the job queue based on the provided
// event by using the delta between the job created and completed times.
func (*githubTracesReceiver) createJobQueueSpan(
	resourceSpans ptrace.ResourceSpans,
	event *github.WorkflowJobEvent,
	traceID pcommon.TraceID,
	parentSpanID pcommon.SpanID,
) error {
	_ = "STUB: not implemented"
	return nil
}

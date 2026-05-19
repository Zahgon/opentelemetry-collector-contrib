// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package gitlabreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/gitlabreceiver"

import (
	"errors"
	"time"

	gitlab "gitlab.com/gitlab-org/api/client-go/v2"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

var (
	errTraceIDGeneration        = errors.New("failed to generate trace ID")
	errPipelineSpanIDGeneration = errors.New("failed to generate pipeline span ID")
	errPipelineSpanProcessing   = errors.New("failed to process pipeline span")
	errStageSpanProcessing      = errors.New("failed to process stage span")
	errJobSpanProcessing        = errors.New("failed to process job span")
)

const (
	gitlabEventTimeFormat = "2006-01-02 15:04:05 UTC"
)

// GitlabEvent abstracts span setup for different GitLab event types (pipeline, stage, job)
// It enables unified span creation logic while allowing type-specific customization
// It must be implemented by all GitLab event types (pipeline, stage, job) to create spans
type GitlabEvent interface {
	setAttributes(pcommon.Map)
	setSpanIDs(ptrace.Span, pcommon.SpanID) error
	setTimeStamps(ptrace.Span, string, string) error
	setSpanData(ptrace.Span) error
}

func (gtr *gitlabTracesReceiver) handlePipeline(e *gitlab.PipelineEvent) (ptrace.Traces, error) {
	_ = "STUB: not implemented"
	return *new(ptrace.Traces), nil
}

func (gtr *gitlabTracesReceiver) processPipelineSpan(r ptrace.ResourceSpans, pipeline *glPipeline, traceID pcommon.TraceID, pipelineSpanID pcommon.SpanID) error {
	_ = "STUB: not implemented"
	return nil
}

func (gtr *gitlabTracesReceiver) processStageSpans(r ptrace.ResourceSpans, pipeline *glPipeline, traceID pcommon.TraceID, parentSpanID pcommon.SpanID) (map[string]*glPipelineStage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Skip stages that have not started yet (empty StartedAt).  This
// happens for early-failed pipelines where a stage was queued but
// never executed.  Attempting to create a span for such a stage
// returns "invalid startedAt timestamp: time is empty".

func (gtr *gitlabTracesReceiver) processJobSpans(r ptrace.ResourceSpans, p *glPipeline, traceID pcommon.TraceID, stages map[string]*glPipelineStage) error {
	_ = "STUB: not implemented"
	return nil
}

func (*gitlabTracesReceiver) createSpan(resourceSpans ptrace.ResourceSpans, e GitlabEvent, traceID pcommon.TraceID, spanID pcommon.SpanID) error {
	_ = "STUB: not implemented"
	return nil
}

// newTraceID creates a deterministic Trace ID based on the provided pipelineID and pipeline finishedAt time.
// It's not possible to create the traceID during a pipeline execution. Details can be found here: https://github.com/open-telemetry/semantic-conventions/issues/1749#issuecomment-2772544215
func newTraceID(pipelineID int64, finishedAt string) (pcommon.TraceID, error) {
	_ = "STUB: not implemented"
	return *new(pcommon.TraceID), nil
}

// newPipelineSpanID creates a deterministic Parent Span ID based on the provided pipelineID and pipeline finishedAt time.
// It's not possible to create the pipelineSpanID during a pipeline execution. Details can be found here: https://github.com/open-telemetry/semantic-conventions/issues/1749#issuecomment-2772544215
func newPipelineSpanID(pipelineID int64, finishedAt string) (pcommon.SpanID, error) {
	_ = "STUB: not implemented"
	return *new(pcommon.SpanID), nil
}

// newStageSpanID creates a deterministic Stage Span ID based on the provided pipelineID, stageName, and stage startedAt time.
// It's not possible to create the stageSpanID during a pipeline execution. Details can be found here: https://github.com/open-telemetry/semantic-conventions/issues/1749#issuecomment-2772544215
func newStageSpanID(pipelineID int64, stageName, startedAt string) (pcommon.SpanID, error) {
	_ = "STUB: not implemented"
	return *new(pcommon.SpanID), nil
}

// newJobSpanID creates a deterministic Job Span ID based on the unique jobID
func newJobSpanID(jobID int64, startedAt string) (pcommon.SpanID, error) {
	_ = "STUB: not implemented"
	return *new(pcommon.SpanID), nil
}

// newSpanID is a helper function to create a Span ID
func newSpanID(input string) (pcommon.SpanID, error) {
	_ = "STUB: not implemented"
	return *new(pcommon.SpanID), nil
}

// newStages extracts stage information from pipeline jobs.
// Since GitLab doesn't provide webhook events for stages within a pipeline, we need to create a new stage for each job.
func (gtr *gitlabTracesReceiver) newStages(pipeline *glPipeline) (map[string]*glPipelineStage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// setStageTime determines stage start/finish times by finding the earliest start and latest finish time
func (*gitlabTracesReceiver) setStageTime(stage *glPipelineStage, jobStartedAt, jobFinishedAt string) error {
	_ = "STUB: not implemented"
	// Handle start time
	return nil
}

// Handle finish time

func setSpanTimeStamps(span ptrace.Span, startTime, endTime string) error {
	_ = "STUB: not implemented"
	return nil
}

// parseGitlabTime parses the time string from the gitlab event, it differs between the test pipeline event and the actual webhook event,
// because the test pipeline event has a different time format than the actual webhook event.
// Test pipeline event time format: 2025-04-01T18:31:49.624Z
// Actual webhook event time format: 2025-04-01 18:31:49 UTC
func parseGitlabTime(t string) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

// Time format of actual webhook events

// Time format of test webhook events

// map GitLab status to OTel span status and set optional message
// Relevant GitLab docs:
// - Job: https://docs.gitlab.com/ci/jobs/#available-job-statuses
// - Pipeline: https://docs.gitlab.com/api/pipelines/#list-project-pipelines (see status field)
func setSpanStatus(span ptrace.Span, status string) { _ = "STUB: not implemented"; return }

func (gtr *gitlabTracesReceiver) setResourceAttributes(attrs pcommon.Map, e *gitlab.PipelineEvent) {
	_ = "STUB: not implemented"
	// Service
	return
}

// CICD

// Resource attributes for workers are not applicable for GitLab, because GitLab provides worker information on job level
// One pipeline can have multiple jobs, and each job can have a different worker
// Therefore we set the worker attributes on job level

// VCS

// Merge Request attributes (only for MR-triggered pipelines)

// ---------- The following attributes are not part of semconv yet ----------

// VCS
// We need to check if the commit timestamp is not nil, otherwise we might have a nil pointer dereference when calling Format()

// User details are only included if explicitly enabled in configuration

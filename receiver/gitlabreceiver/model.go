// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package gitlabreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/gitlabreceiver"

import (
	"errors"

	gitlab "gitlab.com/gitlab-org/api/client-go/v2"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

const (
	// The following attributes are not part of semconv yet, but potential candidates for future inclusion
	// They are highly experimental and subject to change
	AttributeVCSRepositoryVisibility = "vcs.repository.visibility"
	AttributeVCSRepositoryRefDefault = "vcs.repository.ref.default"

	AttributeVCSRefHeadRevisionMessage     = "vcs.ref.head.revision.message"
	AttributeVCSRefHeadRevisionTimestamp   = "vcs.ref.head.revision.timestamp"
	AttributeVCSRefHeadRevisionAuthorName  = "vcs.ref.head.revision.author.name"
	AttributeVCSRefHeadRevisionAuthorEmail = "vcs.ref.head.revision.author.email"

	AttributeCICDPipelineRunActorID   = "cicd.pipeline.run.actor.id"
	AttributeCICDPipelineRunActorName = "cicd.pipeline.run.actor.name"

	AttributeCICDTaskEnvironmentName = "cicd.pipeline.task.run.environment.name"

	AttributeCICDWorkerType   = "cicd.worker.type"
	AttributeCICDWorkerTags   = "cicd.worker.tags"
	AttributeCICDWorkerShared = "cicd.worker.shared"

	// The following attributes are EXCLUSIVE to GitLab but not listed under Vendor Extensions within Semantic Conventions yet
	// They are highly experimental and subject to change
	AttributeGitLabProjectNamespace = "gitlab.project.namespace"

	AttributeGitLabPipelineRunActorUsername = "gitlab.pipeline.run.actor.username"

	AttributeGitLabPipelineSource              = "gitlab.pipeline.source" // Source of the pipeline: https://docs.gitlab.com/ci/jobs/job_rules/#ci_pipeline_source-predefined-variable
	AttributeGitLabPipelineSourcePipelineID    = "gitlab.pipeline.source.pipeline.id"
	AttributeGitLabPipelineSourcePipelineJobID = "gitlab.pipeline.source.pipeline.job.id"

	AttributeGitLabPipelineSourcePipelineProjectID        = "gitlab.pipeline.source.pipeline.project.id"
	AttributeGitLabPipelineSourcePipelineProjectNamespace = "gitlab.pipeline.source.pipeline.project.namespace"
	AttributeGitLabPipelineSourcePipelineProjectURL       = "gitlab.pipeline.source.pipeline.project.url"

	AttributeGitLabJobQueuedDuration = "gitlab.job.queued_duration"
	AttributeGitLabJobFailureReason  = "gitlab.job.failure_reason"
	AttributeGitLabJobAllowFailure   = "gitlab.job.allow_failure"

	AttributeGitLabEnvironmentAction         = "gitlab.environment.action"
	AttributeGitLabEnvironmentDeploymentTier = "gitlab.environment.deployment_tier"
)

var (
	errSetPipelineTimestamps = errors.New("failed to set pipeline span timestamps")
	errSetStageTimestamps    = errors.New("failed to set stage span timestamps")
	errSetJobTimestamps      = errors.New("failed to set job span timestamps")
)

// glPipeline is a wrapper that implements the GitlabEvent interface
type glPipeline struct {
	*gitlab.PipelineEvent
}

func (p *glPipeline) setSpanData(span ptrace.Span) error { _ = "STUB: not implemented"; return nil }

// the glPipeline doesn't have a parent span ID, bc it's the root span
func (*glPipeline) setSpanIDs(span ptrace.Span, spanID pcommon.SpanID) error {
	_ = "STUB: not implemented"
	return nil
}

func (*glPipeline) setTimeStamps(span ptrace.Span, startTime, endTime string) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *glPipeline) setAttributes(attrs pcommon.Map) {
	_ = "STUB: not implemented"
	// ---------- The following attributes are not part of semconv yet ----------
	return
}

// Pipeline

// Parent Pipeline (only added if it's a multi-pipeline)

// glPipelineStage represents a stage in a pipeline event
type glPipelineStage struct {
	PipelineID         int64
	Name               string
	Status             string
	PipelineFinishedAt string
	StartedAt          string
	FinishedAt         string
}

func (s *glPipelineStage) setSpanData(span ptrace.Span) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *glPipelineStage) setSpanIDs(span ptrace.Span, parentSpanID pcommon.SpanID) error {
	_ = "STUB: not implemented"
	return nil
}

func (*glPipelineStage) setTimeStamps(span ptrace.Span, startTime, endTime string) error {
	_ = "STUB: not implemented"
	return nil
}

// Appropriate attributes for stages require further discussion and are yet to be defined as part of: https://github.com/open-telemetry/semantic-conventions/issues/2900
func (*glPipelineStage) setAttributes(pcommon.Map) {
	_ = "STUB: not implemented"

	// glJobEvent represents the PipelineEvent.Builds struct from the pipeline webhook event - it's not exported as type by the Gitlab API client, so we need to use this struct to represent it
	return
}

type glJobEvent struct {
	ID             int64                                  `json:"id"`
	Stage          string                                 `json:"stage"`
	Name           string                                 `json:"name"`
	Status         string                                 `json:"status"`
	CreatedAt      string                                 `json:"created_at"`
	StartedAt      string                                 `json:"started_at"`
	FinishedAt     string                                 `json:"finished_at"`
	Duration       float64                                `json:"duration"`
	QueuedDuration float64                                `json:"queued_duration"`
	FailureReason  string                                 `json:"failure_reason"`
	When           string                                 `json:"when"`
	Manual         bool                                   `json:"manual"`
	AllowFailure   bool                                   `json:"allow_failure"`
	User           *gitlab.EventUser                      `json:"user"`
	Runner         gitlab.PipelineEventBuildRunner        `json:"runner"`
	ArtifactsFile  gitlab.PipelineEventBuildArtifactsFile `json:"artifacts_file"`
	Environment    gitlab.EventEnvironment                `json:"environment"`
}

// glPipelineJob wraps the job event and adds additional fields like jobUrl (required for span attributes)
type glPipelineJob struct {
	event  *glJobEvent
	jobURL string
}

func (j *glPipelineJob) setSpanData(span ptrace.Span) error { _ = "STUB: not implemented"; return nil }

func (j *glPipelineJob) setSpanIDs(span ptrace.Span, parentSpanID pcommon.SpanID) error {
	_ = "STUB: not implemented"
	return nil
}

func (*glPipelineJob) setTimeStamps(span ptrace.Span, startTime, endTime string) error {
	_ = "STUB: not implemented"
	return nil
}

func (j *glPipelineJob) setAttributes(attrs pcommon.Map) {
	_ = "STUB: not implemented"
	// Job
	return
}

// Worker/Runner

// ---------- The following attributes are not part of semconv yet ----------

// Job

// Environment

// Worker/Runner

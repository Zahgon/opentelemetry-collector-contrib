// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package githubreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/githubreceiver"

import (
	"github.com/google/go-github/v86/github"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

// model.go contains custom attributes that complement the standardized attributes
// from OpenTelemetry semantic conventions v1.37.0. While many VCS and CICD attributes
// are now standardized, these custom attributes provide GitHub-specific functionality
// not yet covered by the standard semantic conventions.
const (
	// Note: Many VCS attributes are now standardized in semantic conventions v1.37.0
	// and available through the generated metadata package.

	// vcs.repository.name
	AttributeVCSRepositoryName = "vcs.repository.name"

	// vcs.ref.head.name (used in trace generation)
	AttributeVCSRefHead = "vcs.ref.head"

	// vcs.ref.head.revision
	AttributeVCSRefHeadRevision = "vcs.ref.head.revision"

	// vcs.ref.head.type with enum values of branch or tag.
	// Note: This is now standardized in semantic conventions v1.37.0
	AttributeVCSRefHeadType       = "vcs.ref.head.type"
	AttributeVCSRefHeadTypeBranch = "branch"
	AttributeVCSRefHeadTypeTag    = "tag"

	// The following CICD attributes are not yet standardized in semantic conventions v1.37.0.
	// They provide GitHub-specific functionality and may be subject to change.

	AttributeCICDPipelineRunURLFull = "cicd.pipeline.run.url.full" // equivalent to GitHub's `html_url`

	// CICD pipeline and task run status attributes for GitHub workflow integration
	AttributeCICDPipelineRunStatus             = "cicd.pipeline.run.status" // equivalent to GitHub's `conclusion`
	AttributeCICDPipelineRunStatusSuccess      = "success"
	AttributeCICDPipelineRunStatusFailure      = "failure"
	AttributeCICDPipelineRunStatusCancellation = "cancellation"

	AttributeCICDPipelineRunStatusSkip = "skip"

	AttributeCICDPipelineTaskRunStatus             = "cicd.pipeline.run.task.status" // equivalent to GitHub's `conclusion`
	AttributeCICDPipelineTaskRunStatusSuccess      = "success"
	AttributeCICDPipelineTaskRunStatusFailure      = "failure"
	AttributeCICDPipelineTaskRunStatusCancellation = "cancellation"
	AttributeCICDPipelineTaskRunStatusSkip         = "skip"

	// The following GitHub-specific attributes are not part of semantic conventions v1.37.0.
	AttributeCICDPipelineRunSenderLogin     = "cicd.pipeline.run.sender.login"      // GitHub's Run Sender Login
	AttributeCICDPipelineTaskRunSenderLogin = "cicd.pipeline.task.run.sender.login" // GitHub's Task Sender Login

	AttributeCICDPipelinePreviousAttemptURLFull = "cicd.pipeline.run.previous_attempt.url.full"
	AttributeCICDPipelineWorkerID               = "cicd.pipeline.worker.id"          // GitHub's Runner ID
	AttributeCICDPipelineWorkerGroupID          = "cicd.pipeline.worker.group.id"    // GitHub's Runner Group ID
	AttributeCICDPipelineWorkerName             = "cicd.pipeline.worker.name"        // GitHub's Runner Name
	AttributeCICDPipelineWorkerGroupName        = "cicd.pipeline.worker.group.name"  // GitHub's Runner Group Name
	AttributeCICDPipelineWorkerNodeID           = "cicd.pipeline.worker.node.id"     // GitHub's Runner Node ID
	AttributeCICDPipelineWorkerLabels           = "cicd.pipeline.worker.labels"      // GitHub's Runner Labels
	AttributeCICDPipelineRunQueueDuration       = "cicd.pipeline.run.queue.duration" // GitHub's Queue Duration

	// The following attributes are exclusive to GitHub but not listed under
	// vendor extensions within semantic conventions v1.37.0.
	AttributeGitHubRepositoryCustomProperty = "github.repository.custom_properties" // GitHub's Repository Custom Properties (used in custom property processing)

	// github.reference.workflow acts as a template attribute where it'll be
	// joined with a `name` and a `version` value. There is an unknown amount of
	// reference workflows that are sent as a list of strings by GitHub making
	// it necessary to leverage template attributes. One key thing to note is
	// the length of the names. Evaluate if this causes issues.
	// WARNING: Extremely long workflow file names could create extremely long
	// attribute keys which could lead to unknown issues in the backend and
	// create additional memory usage overhead when processing data (though
	// unlikely).
	// TODO: Evaluate if there is a need to truncate long workflow files names.
	// eg. github.reference.workflow.my-great-workflow.path
	// eg. github.reference.workflow.my-great-workflow.version
	// eg. github.reference.workflow.my-great-workflow.revision
	AttributeGitHubReferenceWorkflow = "github.reference.workflow"

	// SECURITY: This information will always exist on the repository, but may
	// be considered private if the repository is set to private. Care should be
	// taken in the data pipeline for sanitizing sensitive user information if
	// the user deems it as such.
	AttributeVCSRefHeadRevisionAuthorName  = "vcs.ref.head.revision.author.name"  // GitHub's Head Revision Author Name
	AttributeVCSRefHeadRevisionAuthorEmail = "vcs.ref.head.revision.author.email" // GitHub's Head Revision Author Email

)

// getWorkflowRunAttrs returns a pcommon.Map of attributes for the Workflow Run
// GitHub event type and an error if one occurs. The attributes are associated
// with the originally provided resource.
func (gtr *githubTracesReceiver) getWorkflowRunAttrs(resource pcommon.Resource, e *github.WorkflowRunEvent) error {
	_ = "STUB: not implemented"
	return nil
}

// Add all custom properties from the repository as resource attributes

// VCS Attributes

// CICD Attributes

// Default sets to whatever is provided by the event. GitHub provides the
// following additional values: neutral, timed_out, action_required, stale,
// startup_failure, and null.

// Determine if there are any referenced (shared) workflows listed in the
// Workflow Run event and generate the templated attributes for them.

// getWorkflowJobAttrs returns a pcommon.Map of attributes for the Workflow Job
// GitHub event type and an error if one occurs. The attributes are associated
// with the originally provided resource.
func (gtr *githubTracesReceiver) getWorkflowJobAttrs(resource pcommon.Resource, e *github.WorkflowJobEvent) error {
	_ = "STUB: not implemented"
	return nil
}

// Add all custom properties from the repository as resource attributes

// VCS Attributes

// CICD Worker (GitHub Runner) Attributes

// CICD Attributes

// Default sets to whatever is provided by the event. GitHub provides the
// following additional values: neutral, timed_out, action_required, stale,
// and null.

// splitRefWorkflowPath splits the reference workflow path into just the file
// name normalized to lowercase without the file type.
func splitRefWorkflowPath(path string) (fileName string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

// getServiceName returns a generated service.name resource attribute derived
// from 1) the service_name defined in the webhook configuration 2) a
// service.name value set in the custom_properties section of a GitHub event, or
// 3) the repository name. The value returned in those cases will always be a
// formatted string; where the string will be lowercase and underscores will be
// replaced by hyphens. If none of these are set, it returns "unknown_service"
// according to the semantic conventions for service.name and an error.
// https://opentelemetry.io/docs/specs/semconv/attributes-registry/service/#service-attributes
func (gtr *githubTracesReceiver) getServiceName(customProps any, repoName string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// customProps would be an index map[string]interface{} passed in but should
// only be non-nil if the index of `service_name` exists

// This should never happen, but in the event it does, unknown_service
// and an error will be returned to abide by semantic conventions.

// addCustomPropertiesToAttrs adds all custom properties from the repository as resource attributes
// with the prefix AttributeGitHubCustomProperty. Keys are converted to snake_case to follow
// resource attribute naming convention.
func addCustomPropertiesToAttrs(attrs pcommon.Map, customProps map[string]any) {
	_ = "STUB: not implemented"
	return
}

// Skip service_name as it's already handled separately

// Convert key to snake_case

// Use dot notation for keys, following resource attribute naming convention

// Handle different value types

// For any other types, convert to string

// formatString formats a string to lowercase and replaces underscores with
// hyphens.
func formatString(input string) string { _ = "STUB: not implemented"; return "" }

// replaceAPIURL replaces a GitHub API URL with the HTML URL version.
func replaceAPIURL(apiURL string) (htmlURL string) {
	_ = "STUB: not implemented"
	// TODO: Support enterpise server configuration with custom domain.
	return ""
}

// toSnakeCase converts a string to snake_case format.
// It handles all GitHub supported characters for custom property names: a-z, A-Z, 0-9, _, -, $, #.
// This function ensures that the resulting string follows snake_case convention.
func toSnakeCase(s string) string {
	_ = "STUB: not implemented"
	// Replace hyphens, spaces, and dots with underscores
	return ""
}

// Replace special characters with underscores

// Handle camelCase and PascalCase

// If current char is uppercase and previous char is lowercase or a digit,
// or if current char is uppercase and next char is lowercase,
// add an underscore before the current char

// Replace multiple consecutive underscores with a single one

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package logs // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding/azureencodingextension/internal/unmarshaler/logs"

import (
	"encoding/json"

	"go.opentelemetry.io/collector/pdata/pcommon"
)

const (
	// OpenTelemetry attribute name for Data Factory start time of the trigger
	// (representing the initiation of a pipeline) runs in UTC format
	attributeDataFactoryTriggerStartTime = "azure.datafactory.trigger.start_time"

	// OpenTelemetry attribute name for Data Factory end time the trigger
	// (representing the initiation of a pipeline) runs in UTC format
	attributeDataFactoryTriggerEndTime = "azure.datafactory.trigger.end_time"

	// OpenTelemetry attribute name for Data Factory start time of the pipeline runs in UTC format
	attributeDataFactoryPipelineStartTime = "azure.datafactory.pipeline.start_time"

	// OpenTelemetry attribute name for Data Factory end time the pipeline runs in UTC format
	attributeDataFactoryPipelineEndTime = "azure.datafactory.pipeline.end_time"

	// OpenTelemetry attribute name for Data Factory start time of the activity runs (steps in pipeline) in UTC format
	attributeDataFactoryActivityStartTime = "azure.datafactory.activity.start_time"

	// OpenTelemetry attribute name for Data Factory end time the activity runs (steps in pipeline) in UTC format
	attributeDataFactoryActivityEndTime = "azure.datafactory.activity.end_time"

	// OpenTelemetry attribute name for Data Factory User Properties
	attributeDataFactoryUserProperties = "azure.datafactory.user_properties"

	// OpenTelemetry attribute name for Data Factory Annotations
	attributeDataFactoryAnnotations = "azure.datafactory.annotations"

	// OpenTelemetry attribute name for Data Factory Input data
	attributeDataFactoryInput = "azure.datafactory.input"

	// OpenTelemetry attribute name for Data Factory Output data
	attributeDataFactoryOutput = "azure.datafactory.output"

	// OpenTelemetry attribute name for Data Factory Predecessors
	attributeDataFactoryPredecessors = "azure.datafactory.predecessors"

	// OpenTelemetry attribute name for Data Factory Parameters
	attributeDataFactoryParameters = "azure.datafactory.parameters"

	// OpenTelemetry attribute name for Data Factory System Parameters
	attributeDataFactorySystemParameters = "azure.datafactory.system_parameters"

	// OpenTelemetry attribute name for Data Factory Tags
	attributeDataFactoryTags = "azure.datafactory.tags"

	// OpenTelemetry attribute name for Data Factory Activity Run ID
	attributeDataFactoryActivityRunID = "azure.datafactory.activity.run_id"

	// OpenTelemetry attribute name for Data Factory Activity Name
	attributeDataFactoryActivityName = "azure.datafactory.activity.name"

	// OpenTelemetry attribute name for Data Factory Pipeline Run ID
	attributeDataFactoryPipelineRunID = "azure.datafactory.pipeline.run_id"

	// OpenTelemetry attribute name for Data Factory Pipeline Name
	attributeDataFactoryPipelineName = "azure.datafactory.pipeline.name"

	// OpenTelemetry attribute name for Data Factory Trigger Run ID
	attributeDataFactoryTriggerRunID = "azure.datafactory.trigger.run_id"

	// OpenTelemetry attribute name for Data Factory Trigger Name
	attributeDataFactoryTriggerName = "azure.datafactory.trigger.name"

	// OpenTelemetry attribute name for Data Factory Trigger Type
	attributeDataFactoryTriggerType = "azure.datafactory.trigger.type"

	// OpenTelemetry attribute name for Data Factory Trigger Event
	attributeDataFactoryTriggerEventPayload = "azure.datafactory.trigger.event_payload"

	// OpenTelemetry attribute name for Data Factory Pipeline run final status ("Succeeded" or "Failed")
	attributeDataFactoryPipelineState = "azure.datafactory.pipeline.state"

	// OpenTelemetry attribute name for Error Target
	attributeErrorTarget = "error.target"
)

// See https://github.com/MicrosoftDocs/azure-docs/blob/main/articles/data-factory/monitor-data-factory-reference.md#monitor-and-log-analytics-schemas-for-logs-and-events
// Properties has common schema across different Categories (ActivityRuns/PipelineRuns/TriggerRuns)
type azureDataFactoryProperties struct {
	UserProperties map[string]any `json:"UserProperties"`
	Annotations    []string       `json:"Annotations"`
	Input          map[string]any `json:"Input"`
	Output         map[string]any `json:"Output"`
	Error          struct {
		Code        string `json:"errorCode"`
		Message     string `json:"message"`
		FailureType string `json:"failureType"`
		Target      string `json:"target"`
	} `json:"Error"`
	Predecessors     []map[string]any `json:"Predecessors"`
	Parameters       map[string]any   `json:"Parameters"`
	SystemParameters map[string]any   `json:"SystemParameters"`
	Tags             map[string]any   `json:"Tags"`
}

type azureDataFactoryBaseLog struct {
	azureLogRecordBase

	// Common fields for all following categories
	StartTime string `json:"start"` // date-time, UTC
	EndTime   string `json:"end"`   // date-time, UTC
	// will not be used, as we already have parsable Level in "level" field,
	// this field present to avoid parsing error when both "level" and "Level" are present
	DiagnosticLevel json.Number `json:"Level"`

	Properties azureDataFactoryProperties `json:"properties"`
}

func (r *azureDataFactoryBaseLog) PutProperties(attrs pcommon.Map, _ pcommon.Value) error {
	_ = "STUB: not implemented"
	// Process Annotations slice
	return nil
}

// Process Predecessors slice of maps

// Failed to add - put string representation of the attrValue

// Errors

// See https://github.com/MicrosoftDocs/azure-docs/blob/main/articles/data-factory/monitor-data-factory-reference.md#activity-run-log-attributes
type azureDataFactoryActivityRunsLog struct {
	azureDataFactoryBaseLog

	// Additional fields in common schema
	ActivityRunID string `json:"activityRunId"`
	PipelineRunID string `json:"pipelineRunId"`
	PipelineName  string `json:"pipelineName"`
	ActivityName  string `json:"activityName"`
}

func (r *azureDataFactoryActivityRunsLog) PutCommonAttributes(attrs pcommon.Map, body pcommon.Value) {
	_ = "STUB: not implemented"
	// Put common attributes first
	return
}

// Then put custom top-level attributes

// Then put custom top-level attributes

// See https://github.com/MicrosoftDocs/azure-docs/blob/main/articles/data-factory/monitor-data-factory-reference.md#pipeline-run-log-attributes
type azureDataFactoryPipelineRunsLog struct {
	azureDataFactoryBaseLog

	// Additional fields in common schema
	PipelineRunID string `json:"runId"`
	PipelineName  string `json:"pipelineName"`
	Status        string `json:"status"`
}

func (r *azureDataFactoryPipelineRunsLog) PutCommonAttributes(attrs pcommon.Map, body pcommon.Value) {
	_ = "STUB: not implemented"
	// Put common attributes first
	return
}

// Then put custom top-level attributes

// Then put custom top-level attributes

// See https://github.com/MicrosoftDocs/azure-docs/blob/main/articles/data-factory/monitor-data-factory-reference.md#pipeline-run-log-attributes
type azureDataFactoryTriggerRunsLog struct {
	azureDataFactoryBaseLog

	// Additional fields in common schema
	TriggerID    string `json:"triggerId"`
	TriggerName  string `json:"triggerName"`
	TriggerType  string `json:"triggerType"`
	TriggerEvent string `json:"triggerEvent"`
	Status       string `json:"status"`
}

func (r *azureDataFactoryTriggerRunsLog) PutCommonAttributes(attrs pcommon.Map, body pcommon.Value) {
	_ = "STUB: not implemented"
	// Put common attributes first
	return
}

// Then put custom top-level attributes

// Then put custom top-level attributes

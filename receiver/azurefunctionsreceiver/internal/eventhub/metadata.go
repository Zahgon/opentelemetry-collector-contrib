// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package eventhub // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/azurefunctionsreceiver/internal/eventhub"

// Resource attribute keys for Event Hub trigger metadata on log resources.
const (
	AttrEventHubName          = "azure.eventhub.name"
	AttrEventHubPartitionID   = "azure.eventhub.partition.id"
	AttrEventHubConsumerGroup = "azure.eventhub.consumer.group"
	AttrEventHubNamespace     = "azure.eventhub.namespace"
)

// TriggerPartitionContext holds Event Hub partition context from the Azure Functions trigger.
// JSON field names match the invoke request Metadata.TriggerPartitionContext shape.
type TriggerPartitionContext struct {
	IsCheckpointingAfterInvocation bool   `json:"IsCheckpointingAfterInvocation"`
	FullyQualifiedNamespace        string `json:"FullyQualifiedNamespace"`
	EventHubName                   string `json:"EventHubName"`
	ConsumerGroup                  string `json:"ConsumerGroup"`
	PartitionID                    string `json:"PartitionId"`
}

// Metadata is the Event Hub–specific view of the Azure Functions invoke Metadata.
// Only fields relevant to the Event Hub trigger are parsed.
type Metadata struct {
	TriggerPartitionContext TriggerPartitionContext `json:"TriggerPartitionContext"`
}

// ParseMetadata unmarshals the raw invoke Metadata JSON (e.g. from protocol.InvokeRequest.Metadata)
// into Event Hub metadata. If the trigger is not Event Hub or metadata is empty, returned attributes may be empty.
func ParseMetadata(raw []byte) (Metadata, error) {
	_ = "STUB: not implemented"
	return *new(Metadata), nil
}

// ResourceAttributes returns resource attributes to add to logs when include_metadata is true.
// Uses TriggerPartitionContext when present (Event Hub trigger).
func (m *Metadata) ResourceAttributes() map[string]string { _ = "STUB: not implemented"; return nil }

// ExtractMetadata parses the raw invoke Metadata JSON and returns Event Hub resource
// attributes. Intended for use as a MetadataExtractor when include_metadata is true.
// Returns nil on parse error or when no Event Hub context is present.
func ExtractMetadata(raw []byte) map[string]string { _ = "STUB: not implemented"; return nil }

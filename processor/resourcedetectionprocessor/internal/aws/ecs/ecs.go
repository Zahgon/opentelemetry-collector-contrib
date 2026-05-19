// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ecs // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/aws/ecs"

import (
	"context"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/processor"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/ecsutil"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/aws/ecs/internal/metadata"
)

const (
	// TypeStr is type of detector.
	TypeStr = "ecs"
)

var _ internal.Detector = (*Detector)(nil)

type Detector struct {
	provider ecsutil.MetadataProvider
	rb       *metadata.ResourceBuilder
}

func NewDetector(params processor.Settings, dcfg internal.DetectorConfig) (internal.Detector, error) {
	_ = "STUB: not implemented"
	return *new(internal.Detector), nil
}

// Allow metadata provider to be created in incompatible environments and just have a noop Detect()

// Detect records metadata retrieved from the ECS Task Metadata Endpoint (TMDE) as resource attributes
// TODO(willarmiros): Replace all attribute fields and enums with values defined in "conventions" once they exist
func (d *Detector) Detect(context.Context) (resource pcommon.Resource, schemaURL string, err error) {
	_ = "STUB: not implemented"
	// don't attempt to fetch metadata if there's no provider (incompatible env)
	return *new(pcommon.Resource), "", nil
}

// TMDE returns the cluster short name or ARN, so we need to construct the ARN if necessary

// The launch type and log data attributes are only available in TMDE v4

func constructClusterArn(cluster, region, account string) string {
	_ = "STUB: not implemented"
	// If cluster is already an ARN, return it
	return ""
}

// Parses ECS Task ARN into subcomponents according to its spec
// See: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-account-settings.html#ecs-resource-ids
func parseTaskARN(taskARN string) (region, account, taskID string) {
	_ = "STUB: not implemented"
	return "", "", ""
}

// ECS Task ARNs come in two versions. In the old one, the last part of the ARN contains
// only the "task/<task-id>". In the new one, it contains "task/cluster-name/task-id".
// This handles both cases.

// Filter out non-normal containers, our own container since we assume the collector is run as a sidecar,
// "init" containers which only run at startup then shutdown (as indicated by the "KnownStatus" attribute),
// containers not using AWS Logs, and those without log group metadata to get the final lists of valid log data
// See: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-metadata-endpoint-v4.html#task-metadata-endpoint-v4-response
func addValidLogData(containers []ecsutil.ContainerMetadata, self *ecsutil.ContainerMetadata, account string, rb *metadata.ResourceBuilder) {
	_ = "STUB: not implemented"
	return
}

func constructLogGroupArn(region, account, group string) string {
	_ = "STUB: not implemented"
	return ""
}

func constructLogStreamArn(region, account, group, stream string) string {
	_ = "STUB: not implemented"
	return ""
}

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package awsecscontainermetrics // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsecscontainermetricsreceiver/internal/awsecscontainermetrics"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/ecsutil"
)

func containerResource(cm *ecsutil.ContainerMetadata, logger *zap.Logger) pcommon.Resource {
	_ = "STUB: not implemented"
	return *new(pcommon.Resource)
}

func taskResource(tm ecsutil.TaskMetadata) pcommon.Resource {
	_ = "STUB: not implemented"
	return *new(pcommon.Resource)
}

// Task revision: aws.ecs.task.version and aws.ecs.task.revision

// Task launchtype: aws.ecs.task.launch_type (raw string) and aws.ecs.launchtype (lowercase)

// https://docs.aws.amazon.com/AmazonECS/latest/userguide/ecs-account-settings.html
// The new taskARN format: New: arn:aws:ecs:region:aws_account_id:task/cluster-name/task-id
//
//	Old(current): arn:aws:ecs:region:aws_account_id:task/task-id
func getResourceFromARN(arn string) (string, string, string) {
	_ = "STUB: not implemented"
	return "", "", ""
}

// The Amazon Resource Name (ARN) that identifies the cluster. The ARN contains the arn:aws:ecs namespace,
// followed by the Region of the cluster, the AWS account ID of the cluster owner, the cluster namespace,
// and then the cluster name. For example, arn:aws:ecs:region:012345678910:cluster/test.
func getNameFromCluster(cluster string) string { _ = "STUB: not implemented"; return "" }

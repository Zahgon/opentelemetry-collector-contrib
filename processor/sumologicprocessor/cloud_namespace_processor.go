// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package sumologicprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/sumologicprocessor"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

// cloudNamespaceProcessor adds the `cloud.namespace` resource attribute to logs, metrics and traces.
type cloudNamespaceProcessor struct {
	addCloudNamespace bool
}

const (
	cloudNamespaceAttributeName = "cloud.namespace"
	cloudNamespaceAwsEc2        = "aws/ec2"
	cloudNamespaceAwsEcs        = "ecs"
	cloudNamespaceAwsBeanstalk  = "ElasticBeanstalk"
)

func newCloudNamespaceProcessor(addCloudNamespace bool) *cloudNamespaceProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (*cloudNamespaceProcessor) processLogs(logs plog.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

func (*cloudNamespaceProcessor) processMetrics(metrics pmetric.Metrics) error {
	_ = "STUB: not implemented"
	return nil
}

func (*cloudNamespaceProcessor) processTraces(traces ptrace.Traces) error {
	_ = "STUB: not implemented"
	return nil
}

func (proc *cloudNamespaceProcessor) isEnabled() bool { _ = "STUB: not implemented"; return false }

func (*cloudNamespaceProcessor) ConfigPropertyName() string { _ = "STUB: not implemented"; return "" }

// addCloudNamespaceAttribute adds the `cloud.namespace` attribute
// to a collection of attributes that already contains a `cloud.platform` attribute.
// It does not add the `cloud.namespace` attribute for all `cloud.platform` values,
// but only for a few specific ones - namely AWS EC2, AWS ECS, and AWS Elastic Beanstalk.
func addCloudNamespaceAttribute(attributes pcommon.Map) { _ = "STUB: not implemented"; return }

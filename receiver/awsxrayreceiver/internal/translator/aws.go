// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package translator // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsxrayreceiver/internal/translator"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"

	awsxray "github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/xray"
)

func addAWSToResource(aws *awsxray.AWSData, attrs pcommon.Map) {
	_ = "STUB: not implemented"

	// https://github.com/open-telemetry/opentelemetry-collector-contrib/blob/c615d2db351929b99e46f7b427f39c12afe15b54/exporter/awsxrayexporter/translator/aws.go#L121
	// this implies that the current segment being processed is not generated
	// by an AWS entity.
	return
}

// based on https://docs.aws.amazon.com/xray/latest/devguide/xray-api-segmentdocuments.html#api-segmentdocuments-aws
// it's possible to have all cloudwatch_logs, ec2, ecs and beanstalk fields at the same time.

func addAWSToSpan(aws *awsxray.AWSData, attrs pcommon.Map) { _ = "STUB: not implemented"; return }

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package translator // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsxrayreceiver/internal/translator"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"

	awsxray "github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/xray"
)

func addSdkToResource(seg *awsxray.Segment, attrs pcommon.Map) { _ = "STUB: not implemented"; return }

// https://github.com/open-telemetry/opentelemetry-collector-contrib/blob/c615d2db351929b99e46f7b427f39c12afe15b54/exporter/awsxrayexporter/translator/cause.go#L150
// x-ray exporter only supports Java stack trace for now
// TODO: Update this once the exporter is more flexible

// sample *xr.SDK: "X-Ray for Go"

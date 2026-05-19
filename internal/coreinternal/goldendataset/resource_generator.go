// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package goldendataset // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal/goldendataset"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
)

// GenerateResource generates a PData Resource object with representative attributes for the
// underlying resource type specified by the rscID input parameter.
func GenerateResource(rscID PICTInputResource) pcommon.Resource {
	_ = "STUB: not implemented"
	return *new(pcommon.Resource)
}

func appendOnpremVMAttributes(attrMap pcommon.Map) { _ = "STUB: not implemented"; return }

func appendCloudVMAttributes(attrMap pcommon.Map) { _ = "STUB: not implemented"; return }

func appendOnpremK8sAttributes(attrMap pcommon.Map) { _ = "STUB: not implemented"; return }

func appendCloudK8sAttributes(attrMap pcommon.Map) { _ = "STUB: not implemented"; return }

func appendFassAttributes(attrMap pcommon.Map) { _ = "STUB: not implemented"; return }

func appendExecAttributes(attrMap pcommon.Map) { _ = "STUB: not implemented"; return }

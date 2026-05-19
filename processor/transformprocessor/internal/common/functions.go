// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package common // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/transformprocessor/internal/common"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlresource"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlscope"
)

func ResourceFunctions() map[string]ottl.Factory[*ottlresource.TransformContext] {
	_ = "STUB: not implemented"
	return nil
}

func ScopeFunctions() map[string]ottl.Factory[*ottlscope.TransformContext] {
	_ = "STUB: not implemented"
	return nil
}

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottlfuncs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/ottlfuncs"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

// StandardFuncs is a helper function to provide quick access to all functions (editors and converters) in this package
func StandardFuncs[K any]() map[string]ottl.Factory[K] { _ = "STUB: not implemented"; return nil }

// Editors

// StandardConverters is a helper function to provide quick access to all converters in this package
func StandardConverters[K any]() map[string]ottl.Factory[K] { _ = "STUB: not implemented"; return nil }

func converters[K any]() []ottl.Factory[K] { _ = "STUB: not implemented"; return nil }

// Converters

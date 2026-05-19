// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package remotetapextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/remotetapextension"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension"
)

func NewFactory() extension.Factory { _ = "STUB: not implemented"; return *new(extension.Factory) }

func createExtension(_ context.Context, settings extension.Settings, config component.Config) (extension.Extension, error) {
	_ = "STUB: not implemented"
	return *new(extension.Extension), nil
}

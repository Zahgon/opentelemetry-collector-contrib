// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package source // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/headerssetterextension/internal/source"

import (
	"context"
)

var _ Source = (*ContextSource)(nil)

type AttributeSource struct {
	Key          string
	DefaultValue string
}

func (ts *AttributeSource) Get(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

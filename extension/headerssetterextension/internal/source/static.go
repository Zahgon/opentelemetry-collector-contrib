// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package source // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/headerssetterextension/internal/source"

import "context"

var _ Source = (*StaticSource)(nil)

type StaticSource struct {
	Value string
}

func (ts *StaticSource) Get(_ context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

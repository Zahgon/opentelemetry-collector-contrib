// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package faro // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/faro"

import (
	"context"

	faroTypes "github.com/grafana/faro/pkg/go"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

// TranslateToTraces converts faro.Payload into Traces pipeline data
func TranslateToTraces(ctx context.Context, payload faroTypes.Payload) (ptrace.Traces, error) {
	_ = "STUB: not implemented"
	return *new(ptrace.Traces), nil
}

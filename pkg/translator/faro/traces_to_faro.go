// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package faro provides utilities for bidirectional translation between Grafana Faro
// and OpenTelemetry (OTLP) formats. These translation utilities are used by both
// the Faro receiver and Faro exporter components to ensure seamless data flow
// between Faro and OpenTelemetry systems.
package faro // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/faro"

import (
	"context"

	faroTypes "github.com/grafana/faro/pkg/go"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

// TranslateFromTraces converts a Traces pipeline data into []*faro.Payload
func TranslateFromTraces(ctx context.Context, td ptrace.Traces) ([]faroTypes.Payload, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// if payload meta already exists in the metaMap merge payload to the existing payload

func resourceSpansToFaroPayload(rs ptrace.ResourceSpans) faroTypes.Payload {
	_ = "STUB: not implemented"
	return *new(faroTypes.Payload)
}

func extractMetaFromResourceAttributes(resourceAttributes pcommon.Map) faroTypes.Meta {
	_ = "STUB: not implemented"
	return *new(faroTypes.Meta)
}

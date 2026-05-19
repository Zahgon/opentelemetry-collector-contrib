// Copyright The OpenTelemetry Authors
// Copyright (c) 2019 The Jaeger Authors.
// Copyright (c) 2018 Uber Technologies, Inc.
// SPDX-License-Identifier: Apache-2.0

package zipkin // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/zipkin/zipkinthriftconverter"

import (
	"context"

	"github.com/jaegertracing/jaeger-idl/thrift-gen/zipkincore"
)

// SerializeThrift is only used in tests.
func SerializeThrift(ctx context.Context, spans []*zipkincore.Span) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeserializeThrift decodes Thrift bytes to a list of spans.
func DeserializeThrift(ctx context.Context, b []byte) ([]*zipkincore.Span, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Ignore the returned element type

// We don't depend on the size returned by ReadListBegin to preallocate the array because it
// sometimes returns a nil error on bad input and provides an unreasonably large int for size

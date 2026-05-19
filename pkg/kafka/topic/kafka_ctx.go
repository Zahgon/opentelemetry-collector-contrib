// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package topic // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/kafka/topic"

import (
	"context"
)

func WithTopic(ctx context.Context, topic string) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func FromContext(ctx context.Context) (string, bool) { _ = "STUB: not implemented"; return "", false }

type topicContextKey struct{}

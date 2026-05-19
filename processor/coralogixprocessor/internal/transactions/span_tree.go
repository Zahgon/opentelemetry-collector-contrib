// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package transactions // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/coralogixprocessor/internal/transactions"

import (
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"
)

type spanNode struct {
	span     ptrace.Span
	children []*spanNode
}

// buildSpanTree constructs a hierarchical tree of spans
func buildSpanTree(spans []ptrace.Span, logger *zap.Logger) *spanNode {
	_ = "STUB: not implemented"
	return nil
}

// We'll keep the earliest span as root

// If no root span was found, use the earliest span as root

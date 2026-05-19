// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package transactions // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/coralogixprocessor/internal/transactions"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"
)

const (
	TransactionIdentifier     = "cgx.transaction"
	TransactionIdentifierRoot = "cgx.transaction.root"
)

func ApplyTransactionsAttributes(td ptrace.Traces, logger *zap.Logger) (ptrace.Traces, error) {
	_ = "STUB: not implemented"
	return *new(ptrace.Traces), nil
}

func groupSpansByTraceID(td ptrace.Traces) map[pcommon.TraceID][]ptrace.Span {
	_ = "STUB: not implemented"
	return nil
}

func applyTransactionToTrace(currentSpan *spanNode, transactionName string) {
	_ = "STUB: not implemented"
	return
}

func markSpanAsRoot(span ptrace.Span) { _ = "STUB: not implemented"; return }

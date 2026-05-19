// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottl // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"

// grammarPathVisitor is used to extract all path from a parsedStatement or booleanExpression
type grammarPathVisitor struct {
	paths []path
}

func (*grammarPathVisitor) visitEditor(*editor)       { _ = "STUB: not implemented"; return }
func (*grammarPathVisitor) visitConverter(*converter) { _ = "STUB: not implemented"; return }
func (*grammarPathVisitor) visitValue(*value)         { _ = "STUB: not implemented"; return }
func (*grammarPathVisitor) visitMathExprLiteral(*mathExprLiteral) {
	_ = "STUB: not implemented"
	return
}

func (v *grammarPathVisitor) visitPath(value *path) { _ = "STUB: not implemented"; return }

func getParsedStatementPaths(ps *parsedStatement) []path { _ = "STUB: not implemented"; return nil }

func getBooleanExpressionPaths(be *booleanExpression) []path { _ = "STUB: not implemented"; return nil }

func getValuePaths(v *value) []path { _ = "STUB: not implemented"; return nil }

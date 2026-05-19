// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package helper // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"

import (
	"os"
	"sync"

	"github.com/expr-lang/expr/ast"
	"github.com/expr-lang/expr/vm"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/entry"
)

// ExprStringConfig is a string that represents an expression
type ExprStringConfig string

const (
	exprStartToken = "EXPR("
	exprEndToken   = ")"
)

// Build creates an ExprStr string from the specified config
func (e ExprStringConfig) Build() (*ExprString, error) { _ = "STUB: not implemented"; return nil, nil }

// Find the first instance of the start token

// Start token does not exist in the remainder of the string,
// so treat the rest as a string literal

// Restrict our end token search range to the next instance of the start token

// Greedily search for the last end token in the search range

// End token does not exist before the next start token
// or end of expression string, so treat the remainder of the string
// as a string literal

// Unscope the indexes and add the partitioned strings

// Reset the starting range and finish if it reaches the end of the string

func ExprCompile(input string) (*vm.Program, error) { _ = "STUB: not implemented"; return nil, nil }

func ExprCompileBool(input string) (*vm.Program, error) { _ = "STUB: not implemented"; return nil, nil }

// An ExprString is made up of a list of string literals
// interleaved with expressions. len(SubStrings) == len(SubExprs) + 1
type ExprString struct {
	SubStrings []string
	SubExprs   []*vm.Program
}

// Render will render an ExprString as a string
func (e *ExprString) Render(env map[string]any) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

type patcher struct{}

func (*patcher) Visit(node *ast.Node) { _ = "STUB: not implemented"; return }

var envPool = sync.Pool{
	New: func() any {
		return map[string]any{
			"os_env_func": os.Getenv,
		}
	},
}

// GetExprEnv returns a map of key/value pairs that can be be used to evaluate an expression
func GetExprEnv(e *entry.Entry) map[string]any { _ = "STUB: not implemented"; return nil }

// PutExprEnv adds a key/value pair that will can be used to evaluate an expression
func PutExprEnv(e map[string]any) { _ = "STUB: not implemented"; return }

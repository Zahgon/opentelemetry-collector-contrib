// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottl // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"

import (
	"go.opentelemetry.io/collector/component"
)

var defaultContextInferPriority = []string{
	"log",
	"datapoint",
	"metric",
	"spanevent",
	"span",
	"profile",
	"scope",
	"instrumentation_scope",
	"resource",
}

// contextInferrer is an interface used to infer the OTTL context from statements.
type contextInferrer interface {
	// inferFromStatements returns the OTTL context inferred from the given statements.
	inferFromStatements(statements []string) (string, error)
	// inferFromConditions returns the OTTL context inferred from the given conditions.
	inferFromConditions(conditions []string) (string, error)
	// inferFromValueExpressions returns the OTTL context inferred from the given value expressions.
	inferFromValueExpressions(expressions []string) (string, error)
	// infer returns the OTTL context inferred from the given statements, conditions,
	// and value expressions.
	infer(statements, conditions, valueExpressions []string) (string, error)
}

type priorityContextInferrer struct {
	telemetrySettings component.TelemetrySettings
	contextPriority   map[string]int
	contextCandidate  map[string]*priorityContextInferrerCandidate
}

type priorityContextInferrerCandidate struct {
	hasEnumSymbol    func(enum *EnumSymbol) bool
	hasFunctionName  func(name string) bool
	getLowerContexts func(context string) []string
}

type priorityContextInferrerOption func(*priorityContextInferrer)

// newPriorityContextInferrer creates a new priority-based context inferrer. To infer the context,
// it uses a slice of priorities (withContextInferrerPriorities) and a set of hints extracted from
// the parsed statements.
//
// To be eligible, a context must support all functions and enums symbols present on the statements.
// If the path context with the highest priority does not meet this requirement, it falls back to its
// lower contexts, testing them with the same logic and choosing the first one that meets all requirements.
//
// If non-prioritized contexts are found on the statements, they get assigned the lowest possible priority,
// and are only selected if no other prioritized context is found.
func newPriorityContextInferrer(telemetrySettings component.TelemetrySettings, contextsCandidate map[string]*priorityContextInferrerCandidate, options ...priorityContextInferrerOption) contextInferrer {
	_ = "STUB: not implemented"
	return *new(contextInferrer)
}

// withContextInferrerPriorities sets the contexts candidates priorities. The lower the
// context position is in the array, the more priority it will have over other items.
func withContextInferrerPriorities(priorities []string) priorityContextInferrerOption {
	_ = "STUB: not implemented"
	return *new(priorityContextInferrerOption)
}

func (s *priorityContextInferrer) inferFromConditions(conditions []string) (inferredContext string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (s *priorityContextInferrer) inferFromStatements(statements []string) (inferredContext string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (s *priorityContextInferrer) inferFromValueExpressions(expressions []string) (inferredContext string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (s *priorityContextInferrer) infer(statements, conditions, valueExprs []string) (inferredContext string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (s *priorityContextInferrer) inferFromHints(hints []priorityContextInferrerHints) (inferredContext string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

// No inferred context

// If no functions or enums are required, return the inferred context directly.

// validateContextCandidate checks if the given context candidate has all required functions names
// and enums symbols. The functions arity are not verified.
func (s *priorityContextInferrer) validateContextCandidate(
	context string,
	requiredFunctions map[string]struct{},
	requiredEnums map[enumSymbol]struct{},
) error {
	_ = "STUB: not implemented"
	return nil
}

// inferFromLowerContexts returns the first lower context that supports all required functions
// and enum symbols used on the statements.
// If no lower context meets the requirements, or if the context candidate is unknown, it
// returns an empty string.
func (s *priorityContextInferrer) inferFromLowerContexts(
	context string,
	requiredFunctions map[string]struct{},
	requiredEnums map[enumSymbol]struct{},
) (inferredContext string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

// sortContextCandidates sorts the slice candidates using the priorityContextInferrer.contextsPriority order.
func (s *priorityContextInferrer) sortContextCandidates(candidates []string) {
	_ = "STUB: not implemented"
	return
}

// getConditionsHints extracts all path, function names (editor and converter), and enumSymbol
// from the given condition. These values are used by the context inferrer as hints to
// select a context in which the function/enum are supported.
func (*priorityContextInferrer) getConditionsHints(conditions []string) ([]priorityContextInferrerHints, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getStatementsHints extracts all path, function names (editor and converter), and enumSymbol
// from the given statement. These values are used by the context inferrer as hints to
// select a context in which the function/enum are supported.
func (*priorityContextInferrer) getStatementsHints(statements []string) ([]priorityContextInferrerHints, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getValueExpressionsHints extracts all path, function (converter) names, and enumSymbol
// from the given value expressions. These values are used by the context inferrer as hints to
// select a context in which the function/enum are supported.
func (*priorityContextInferrer) getValueExpressionsHints(exprs []string) ([]priorityContextInferrerHints, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// priorityContextInferrerHints is a grammarVisitor implementation that collects
// all path, function names (converter.Function and editor.Function), and enumSymbol.
type priorityContextInferrerHints struct {
	paths        []path
	functions    map[string]struct{}
	enumsSymbols map[enumSymbol]struct{}
}

func newGrammarContextInferrerVisitor() priorityContextInferrerHints {
	_ = "STUB: not implemented"
	return *new(priorityContextInferrerHints)
}

func (*priorityContextInferrerHints) visitMathExprLiteral(*mathExprLiteral) {
	_ = "STUB: not implemented"
	return
}

func (v *priorityContextInferrerHints) visitEditor(e *editor) { _ = "STUB: not implemented"; return }

func (v *priorityContextInferrerHints) visitConverter(c *converter) {
	_ = "STUB: not implemented"
	return
}

func (v *priorityContextInferrerHints) visitValue(va *value) { _ = "STUB: not implemented"; return }

func (v *priorityContextInferrerHints) visitPath(value *path) { _ = "STUB: not implemented"; return }

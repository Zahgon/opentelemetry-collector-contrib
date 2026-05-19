// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottl // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"

import (
	"go.opentelemetry.io/collector/component"
)

// StatementsGetter represents a set of statements to be parsed.
//
// Experimental: *NOTE* this API is subject to change or removal in the future.
type StatementsGetter interface {
	// GetStatements retrieves the OTTL statements to be parsed
	GetStatements() []string
}

// NewStatementsGetter creates a new StatementsGetter.
//
// Experimental: *NOTE* this API is subject to change or removal in the future.
func NewStatementsGetter(statements []string) StatementsGetter {
	_ = "STUB: not implemented"
	return *new(StatementsGetter)
}

// ConditionsGetter represents a set of conditions to be parsed.
//
// Experimental: *NOTE* this API is subject to change or removal in the future.
type ConditionsGetter interface {
	// GetConditions retrieves the OTTL conditions to be parsed
	GetConditions() []string
}

// NewConditionsGetter creates a new ConditionsGetter.
//
// Experimental: *NOTE* this API is subject to change or removal in the future.
func NewConditionsGetter(conditions []string) ConditionsGetter {
	_ = "STUB: not implemented"
	return *new(ConditionsGetter)
}

// ValueExpressionsGetter represents a set of value expressions to be parsed.
//
// Experimental: *NOTE* this API is subject to change or removal in the future.
type ValueExpressionsGetter interface {
	// GetValueExpressions retrieves the OTTL value expressions to be parsed
	GetValueExpressions() []string
}

// NewValueExpressionsGetter creates a new ValueExpressionsGetter.
//
// Experimental: *NOTE* this API is subject to change or removal in the future.
func NewValueExpressionsGetter(expressions []string) ValueExpressionsGetter {
	_ = "STUB: not implemented"
	return *new(ValueExpressionsGetter)
}

type defaultOTTLGetter []string

func (d defaultOTTLGetter) GetStatements() []string { _ = "STUB: not implemented"; return nil }

func (d defaultOTTLGetter) GetConditions() []string { _ = "STUB: not implemented"; return nil }

func (d defaultOTTLGetter) GetValueExpressions() []string {
	_ = "STUB: not implemented"

	// ParserCollection is a configurable set of ottl.Parser that can handle multiple OTTL contexts
	// parsings, inferring the context, choosing the right parser for the given statements, and
	// transforming the parsed ottl.Statement[K] slice into a common result of type R.
	//
	// Experimental: *NOTE* this API is subject to change or removal in the future.
	return nil
}

type ParserCollection[R any] struct {
	contextParsers            map[string]*ParserCollectionContextParser[R]
	contextInferrer           contextInferrer
	contextInferrerCandidates map[string]*priorityContextInferrerCandidate
	candidatesLowerContexts   map[string][]string
	modifiedLogging           bool
	Settings                  component.TelemetrySettings
	ErrorMode                 ErrorMode
}

// ParserCollectionOption is a configurable ParserCollection option.
//
// Experimental: *NOTE* this API is subject to change or removal in the future.
type ParserCollectionOption[R any] func(*ParserCollection[R]) error

// NewParserCollection creates a new ParserCollection.
//
// Experimental: *NOTE* this API is subject to change or removal in the future.
func NewParserCollection[R any](
	settings component.TelemetrySettings,
	options ...ParserCollectionOption[R],
) (*ParserCollection[R], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ParsedStatementsConverter is a function that converts the parsed ottl.Statement[K] into
// a common representation to all parser collection contexts passed through WithParserCollectionContext.
// Given each parser has its own transform context type, they must agree on a common type [R]
// so it can be returned by the ParserCollection.ParseStatements and ParserCollection.ParseStatementsWithContext
// functions.
//
// Experimental: *NOTE* this API is subject to change or removal in the future.
type ParsedStatementsConverter[K any, R any] func(collection *ParserCollection[R], statements StatementsGetter, parsedStatements []*Statement[K]) (R, error)

// ParsedConditionsConverter is a function that converts the parsed ottl.Condition[K] into
// a common representation to all parser collection contexts passed through WithParserCollectionContext.
// Given each parser has its own transform context type, they must agree on a common type [R]
// so it can be returned by the ParserCollection.ParseConditions and ParserCollection.ParseConditionsWithContext
// functions.
//
// Experimental: *NOTE* this API is subject to change or removal in the future.
type ParsedConditionsConverter[K any, R any] func(collection *ParserCollection[R], conditions ConditionsGetter, parsedConditions []*Condition[K]) (R, error)

// ParsedValueExpressionsConverter is a function that converts the parsed ottl.ValueExpression[K] into
// a representation common to all parser collection contexts passed through WithParserCollectionContext.
// Given each parser has its own transform context type, they must agree on a common type [R]
// so it can be returned by the ParserCollection.ParseValueExpressions and ParserCollection.ParseValueExpressionsWithContext
// functions.
//
// Experimental: *NOTE* this API is subject to change or removal in the future.
type ParsedValueExpressionsConverter[K any, R any] func(collection *ParserCollection[R], expressions ValueExpressionsGetter, parsedValueExpressions []*ValueExpression[K]) (R, error)

func newNopParsedStatementsConverter[K any]() ParsedStatementsConverter[K, any] {
	_ = "STUB: not implemented"
	return nil
}

func newNopParsedConditionsConverter[K any]() ParsedConditionsConverter[K, any] {
	_ = "STUB: not implemented"
	return nil
}

func newNopParsedValueExpressionsConverter[K any]() ParsedValueExpressionsConverter[K, any] {
	_ = "STUB: not implemented"
	return nil
}

type (
	// ParserCollectionContextOption is a configurable ParserCollectionContext option.
	//
	// Experimental: *NOTE* this API is subject to change or removal in the future.
	ParserCollectionContextOption[K, R any] func(*ParserCollectionContextParser[R], *Parser[K])

	// parserCollectionContextParserFunc is the internal generic type that parses the given []string
	// returned from a getter G into type [R] using the specified context's OTTL parser.
	// The provided context must be supported by the ParserCollection, otherwise an error is returned.
	// If the OTTL Path does not provide their Path.Context value, the prependPathsContext argument should be set to true,
	// so it rewrites the OTTL prepending the missing paths contexts.
	parserCollectionContextParserFunc[R any, G any] func(collection *ParserCollection[R], context string, getter G, prependPathsContext bool) (R, error)
	// ParserCollectionContextParser is a struct that holds the converters for parsing statements and conditions
	// into a common representation of type [R].
	//
	// Experimental: *NOTE* this API is subject to change or removal in the future.
	ParserCollectionContextParser[R any] struct {
		parseStatements       parserCollectionContextParserFunc[R, StatementsGetter]
		parseConditions       parserCollectionContextParserFunc[R, ConditionsGetter]
		parseValueExpressions parserCollectionContextParserFunc[R, ValueExpressionsGetter]
	}
)

// createConditionsParserWithConverter is a method to create the necessary parser wrapper and shadowing the K type.
func createConditionsParserWithConverter[K, R any](converter ParsedConditionsConverter[K, R], parser *Parser[K]) parserCollectionContextParserFunc[R, ConditionsGetter] {
	_ = "STUB: not implemented"
	return nil
}

// createValueExpressionsParserWithConverter is a method to create the necessary parser wrapper and shadowing the K type.
func createValueExpressionsParserWithConverter[K, R any](converter ParsedValueExpressionsConverter[K, R], parser *Parser[K]) parserCollectionContextParserFunc[R, ValueExpressionsGetter] {
	_ = "STUB: not implemented"
	return nil
}

// createStatementsParserWithConverter is a method to create the necessary parser wrapper and shadowing the K type.
func createStatementsParserWithConverter[K, R any](converter ParsedStatementsConverter[K, R], parser *Parser[K]) parserCollectionContextParserFunc[R, StatementsGetter] {
	_ = "STUB: not implemented"
	return nil
}

// WithConditionConverter sets the condition converter for the given context.
// The provided converter function will be used to convert parsed OTTL conditions into a common representation of type R.
// The context's OTTL parser will parse the conditions, and the converter function will transform the parsed conditions into the desired representation.
//
// Experimental: *NOTE* this API is subject to change or removal in the future.
func WithConditionConverter[K, R any](converter ParsedConditionsConverter[K, R]) ParserCollectionContextOption[K, R] {
	_ = "STUB: not implemented"
	return nil
}

// WithValueExpressionConverter sets the value expression converter for the given context.
// The provided converter function will be used to convert parsed OTTL value expressions into a common representation of type R.
// The context's OTTL parser will parse the value expressions, and the converter function will transform the parsed value expressions into the desired representation.
//
// Experimental: *NOTE* this API is subject to change or removal in the future.
func WithValueExpressionConverter[K, R any](converter ParsedValueExpressionsConverter[K, R]) ParserCollectionContextOption[K, R] {
	_ = "STUB: not implemented"
	return nil
}

// WithStatementConverter sets the statement converter for the given context.
// The provided converter function will be used to convert parsed OTTL statements into a common representation of type R.
// The context's OTTL parser will parse the statements, and the converter function will transform the parsed statements into the desired representation.
//
// Experimental: *NOTE* this API is subject to change or removal in the future.
func WithStatementConverter[K, R any](converter ParsedStatementsConverter[K, R]) ParserCollectionContextOption[K, R] {
	_ = "STUB: not implemented"
	return nil
}

// WithParserCollectionContext configures an ottl.Parser for the given context.
// The provided ottl.Parser must be configured to support the provided context using
// the ottl.WithPathContextNames option.
//
// Experimental: *NOTE* this API is subject to change or removal in the future.
func WithParserCollectionContext[K, R any](
	context string,
	parser *Parser[K],
	opts ...ParserCollectionContextOption[K, R],
) ParserCollectionOption[R] {
	_ = "STUB: not implemented"
	return nil
}

func (pc *ParserCollection[R]) getLowerContexts(context string) []string {
	_ = "STUB: not implemented"
	return nil
}

// WithParserCollectionErrorMode has no effect on the ParserCollection, but might be used
// by the ParsedStatementsConverter functions to handle/create StatementSequence.
//
// Experimental: *NOTE* this API is subject to change or removal in the future.
func WithParserCollectionErrorMode[R any](errorMode ErrorMode) ParserCollectionOption[R] {
	_ = "STUB: not implemented"
	return nil
}

// EnableParserCollectionModifiedPathsLogging controls the modification logs.
// When enabled, it logs any modifications performed by the parsing operations,
// instructing users to rewrite the statements accordingly.
//
// Experimental: *NOTE* this API is subject to change or removal in the future.
func EnableParserCollectionModifiedPathsLogging[R any](enabled bool) ParserCollectionOption[R] {
	_ = "STUB: not implemented"
	return nil
}

type parseCollectionContextInferenceOptions struct {
	conditions []string
}

// ParserCollectionContextInferenceOption allows configuring the context inference and use
// this option with the supported parsing functions.
//
// Experimental: *NOTE* this API is subject to change or removal in the future.
type ParserCollectionContextInferenceOption func(p *parseCollectionContextInferenceOptions)

// WithContextInferenceConditions sets additional OTTL conditions to be used to enhance
// the context inference process. This is particularly useful when the statements alone are
// insufficient for determine the correct context, or when a less-specific context is desired.
//
// Experimental: *NOTE* this API is subject to change or removal in the future.
func WithContextInferenceConditions(conditions []string) ParserCollectionContextInferenceOption {
	_ = "STUB: not implemented"
	return *new(ParserCollectionContextInferenceOption)
}

// ParseStatements parses the given statements into [R] using the configured context's ottl.Parser
// and subsequently calling the ParsedStatementsConverter function.
// The statement's context is automatically inferred from the [Path.Context] values, choosing the
// highest priority context found.
// If no contexts are present in the statements, or if the inferred value is not supported by
// the [ParserCollection], it returns an error.
// If parsing the statements fails, it returns the underlying [ottl.Parser.ParseStatements] error.
// If the provided StatementsGetter also implements ContextInferenceHintsProvider, it uses the
// additional OTTL conditions to enhance the context inference. This is particularly useful when
// the statements alone are insufficient for determine the correct context, or if an less-specific
// parser is desired.
//
// Experimental: *NOTE* this API is subject to change or removal in the future.
func (pc *ParserCollection[R]) ParseStatements(statements StatementsGetter, options ...ParserCollectionContextInferenceOption) (R, error) {
	_ = "STUB: not implemented"
	return *new(R), nil
}

// ParseStatementsWithContext parses the given statements into [R] using the configured
// context's ottl.Parser and subsequently calling the ParsedStatementsConverter function.
// Unlike ParseStatements, it uses the provided context and does not infer it
// automatically. The context value must be supported by the [ParserCollection],
// otherwise an error is returned.
// If the statement's Path does not provide their Path.Context value, the prependPathsContext
// argument should be set to true, so it rewrites the statements prepending the missing paths
// contexts.
// If parsing the statements fails, it returns the underlying [ottl.Parser.ParseStatements] error.
//
// Experimental: *NOTE* this API is subject to change or removal in the future.
func (pc *ParserCollection[R]) ParseStatementsWithContext(context string, statements StatementsGetter, prependPathsContext bool) (R, error) {
	_ = "STUB: not implemented"
	return *new(R), nil
}

// ParseConditions parses the given conditions into [R] using the configured context's ottl.Parser
// and subsequently calling the ParsedConditionsConverter function.
// The condition's context is automatically inferred from the [Path.Context] values, choosing the
// highest priority context found.
// If no contexts are present in the conditions, or if the inferred value is not supported by
// the [ParserCollection], it returns an error.
// If parsing the conditions fails, it returns the underlying [ottl.Parser.ParseConditions] error.
//
// Experimental: *NOTE* this API is subject to change or removal in the future.
func (pc *ParserCollection[R]) ParseConditions(conditions ConditionsGetter) (R, error) {
	_ = "STUB: not implemented"
	return *new(R), nil
}

// ParseConditionsWithContext parses the given conditions into [R] using the configured
// context's ottl.Parser and subsequently calling the ParsedConditionsConverter function.
// Unlike ParseConditions, it uses the provided context and does not infer it
// automatically. The context value must be supported by the [ParserCollection],
// otherwise an error is returned.
// If the condition's Path does not provide their Path.Context value, the prependPathsContext
// argument should be set to true, so it rewrites the conditions prepending the missing paths
// contexts.
// If parsing the conditions fails, it returns the underlying [ottl.Parser.ParseConditions] error.
//
// Experimental: *NOTE* this API is subject to change or removal in the future.
func (pc *ParserCollection[R]) ParseConditionsWithContext(context string, conditions ConditionsGetter, prependPathsContext bool) (R, error) {
	_ = "STUB: not implemented"
	return *new(R), nil
}

// ParseValueExpressions parses the given expressions into [R] using the configured context's ottl.Parser
// and subsequently calling the ParsedValueExpressionsConverter function.
// The expression's context is automatically inferred from the [Path.Context] values, choosing the
// highest priority context found.
// If no contexts are present in the expressions, or if the inferred value is not supported by
// the [ParserCollection], it returns an error.
// If parsing the expressions fails, it returns the underlying [ottl.Parser.ParseValueExpressions] error.
//
// Experimental: *NOTE* this API is subject to change or removal in the future.
func (pc *ParserCollection[R]) ParseValueExpressions(expressions ValueExpressionsGetter, options ...ParserCollectionContextInferenceOption) (R, error) {
	_ = "STUB: not implemented"
	return *new(R), nil
}

// ParseValueExpressionsWithContext parses the given expressions into [R] using the configured
// context's ottl.Parser and subsequently calling the ParsedValueExpressionsConverter function.
// Unlike ParseValueExpressions, it uses the provided context and does not infer it
// automatically. The context value must be supported by the [ParserCollection],
// otherwise an error is returned.
// If the expression's Path does not provide their Path.Context value, the prependPathsContext
// argument should be set to true, so it rewrites the expressions prepending the missing paths
// contexts.
// If parsing the expressions fails, it returns the underlying [ottl.Parser.ParseValueExpressions] error.
//
// Experimental: *NOTE* this API is subject to change or removal in the future.
func (pc *ParserCollection[R]) ParseValueExpressionsWithContext(context string, expressions ValueExpressionsGetter, prependPathsContext bool) (R, error) {
	_ = "STUB: not implemented"
	return *new(R), nil
}

func (pc *ParserCollection[R]) logModifications(originalStatements, modifiedStatements []string) {
	_ = "STUB: not implemented"
	return
}

func (pc *ParserCollection[R]) supportedContextNames() []string {
	_ = "STUB: not implemented"
	return nil
}

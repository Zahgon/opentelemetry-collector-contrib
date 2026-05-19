// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottl // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"

import (
	"github.com/alecthomas/participle/v2/lexer"
)

// parsedStatement represents a parsed statement. It is the entry point into the statement DSL.
type parsedStatement struct {
	Editor editor `parser:"(@@"`
	// If converter is matched then return error
	Converter   *converter         `parser:"|@@)"`
	WhereClause *booleanExpression `parser:"( 'where' @@ )?"`
}

func (p *parsedStatement) checkForCustomError() error { _ = "STUB: not implemented"; return nil }

type constExpr struct {
	Boolean   *boolean   `parser:"( @Boolean"`
	Converter *converter `parser:"| @@ )"`
}

// booleanValue represents something that evaluates to a boolean --
// either an equality or inequality, explicit true or false, or
// a parenthesized subexpression.
type booleanValue struct {
	Negation   *string            `parser:"@OpNot?"`
	Comparison *comparison        `parser:"( @@"`
	ConstExpr  *constExpr         `parser:"| @@"`
	SubExpr    *booleanExpression `parser:"| '(' @@ ')' )"`
}

func (b *booleanValue) accept(v grammarVisitor) { _ = "STUB: not implemented"; return }

// opAndBooleanValue represents the right side of an AND boolean expression.
type opAndBooleanValue struct {
	Operator string        `parser:"@OpAnd"`
	Value    *booleanValue `parser:"@@"`
}

func (b *opAndBooleanValue) accept(v grammarVisitor) { _ = "STUB: not implemented"; return }

// term represents an arbitrary number of boolean values joined by AND.
type term struct {
	Left  *booleanValue        `parser:"@@"`
	Right []*opAndBooleanValue `parser:"@@*"`
}

func (b *term) accept(v grammarVisitor) { _ = "STUB: not implemented"; return }

// opOrTerm represents the right side of an OR boolean expression.
type opOrTerm struct {
	Operator string `parser:"@OpOr"`
	Term     *term  `parser:"@@"`
}

func (b *opOrTerm) accept(v grammarVisitor) { _ = "STUB: not implemented"; return }

// booleanExpression represents a true/false decision expressed
// as an arbitrary number of terms separated by OR.
type booleanExpression struct {
	Left  *term       `parser:"@@"`
	Right []*opOrTerm `parser:"@@*"`
}

func (b *booleanExpression) checkForCustomError() error { _ = "STUB: not implemented"; return nil }

func (b *booleanExpression) accept(v grammarVisitor) { _ = "STUB: not implemented"; return }

// compareOp is the type of a comparison operator.
type compareOp int

// These are the allowed values of a compareOp
const (
	eq compareOp = iota
	ne
	lt
	lte
	gte
	gt
)

// a fast way to get from a string to a compareOp
var compareOpTable = map[string]compareOp{
	"==": eq,
	"!=": ne,
	"<":  lt,
	"<=": lte,
	">":  gt,
	">=": gte,
}

// Capture is how the parser converts an operator string to a compareOp.
func (c *compareOp) Capture(values []string) error { _ = "STUB: not implemented"; return nil }

// String() for compareOp gives us more legible test results and error messages.
func (c *compareOp) String() string { _ = "STUB: not implemented"; return "" }

// comparison represents an optional boolean condition.
type comparison struct {
	Left  value     `parser:"@@"`
	Op    compareOp `parser:"@OpComparison"`
	Right value     `parser:"@@"`
}

func (c *comparison) accept(v grammarVisitor) { _ = "STUB: not implemented"; return }

// editor represents the function call of a statement.
type editor struct {
	Function  string     `parser:"@(Lowercase(Uppercase | Lowercase)*)"`
	Arguments []argument `parser:"'(' ( @@ ( ',' @@ )* )? ')'"`
	// If keys are matched return an error
	Keys []key `parser:"( @@ )*"`
}

func (i *editor) accept(v grammarVisitor) { _ = "STUB: not implemented"; return }

// converter represents a converter function call.
type converter struct {
	Function  string     `parser:"@(Uppercase(Uppercase | Lowercase)*)"`
	Arguments []argument `parser:"'(' ( @@ ( ',' @@ )* )? ')'"`
	Keys      []key      `parser:"( @@ )*"`
}

func (c *converter) accept(v grammarVisitor) { _ = "STUB: not implemented"; return }

type argument struct {
	Name         string  `parser:"(@(Lowercase(Uppercase | Lowercase)*) Equal)?"`
	Value        value   `parser:"( @@"`
	FunctionName *string `parser:"| @(Uppercase(Uppercase | Lowercase)*) )"`
}

func (a *argument) accept(v grammarVisitor) {
	_ = "STUB: not implemented"

	// value represents a part of a parsed statement which is resolved to a value of some sort. This can be a telemetry path
	// mathExpression, function call, or literal.
	return
}

type value struct {
	IsNil          *isNil           `parser:"( @Nil"`
	Literal        *mathExprLiteral `parser:"| @@ (?! OpAddSub | OpMultDiv)"`
	MathExpression *mathExpression  `parser:"| @@"`
	Bytes          *byteSlice       `parser:"| @Bytes"`
	String         *string          `parser:"| @String"`
	Bool           *boolean         `parser:"| @Boolean"`
	Enum           *enumSymbol      `parser:"| @Uppercase (?! Lowercase)"`
	Map            *mapValue        `parser:"| @@"`
	List           *list            `parser:"| @@)"`
}

func (v *value) checkForCustomError() error { _ = "STUB: not implemented"; return nil }

func (v *value) accept(vis grammarVisitor) { _ = "STUB: not implemented"; return }

// path represents a telemetry path mathExpression.
type path struct {
	Pos     lexer.Position
	Context string  `parser:"(@Lowercase '.')?"`
	Fields  []field `parser:"@@ ( '.' @@ )*"`
}

func (p *path) accept(v grammarVisitor) { _ = "STUB: not implemented"; return }

// field is an item within a path.
type field struct {
	Name string `parser:"@Lowercase"`
	Keys []key  `parser:"( @@ )*"`
}

func (f *field) accept(v grammarVisitor) { _ = "STUB: not implemented"; return }

type key struct {
	String         *string          `parser:"'[' (@String "`
	Int            *int64           `parser:"| @Int"`
	MathExpression *mathExpression  `parser:"| @@"`
	Expression     *mathExprLiteral `parser:"| @@ ) ']'"`
}

func (k *key) accept(v grammarVisitor) { _ = "STUB: not implemented"; return }

type list struct {
	Values []value `parser:"'[' (@@)* (',' @@)* ']'"`
}

type mapValue struct {
	Values []mapItem `parser:"'{' (@@ ','?)* '}'"`
}

func (m *mapValue) accept(v grammarVisitor) { _ = "STUB: not implemented"; return }

type mapItem struct {
	Key   *string `parser:"@String ':'"`
	Value *value  `parser:"@@"`
}

// byteSlice type for capturing byte slices
type byteSlice []byte

func (b *byteSlice) Capture(values []string) error { _ = "STUB: not implemented"; return nil }

// boolean Type for capturing booleans, see:
// https://github.com/alecthomas/participle#capturing-boolean-value
type boolean bool

func (b *boolean) Capture(values []string) error { _ = "STUB: not implemented"; return nil }

type isNil bool

func (n *isNil) Capture(_ []string) error { _ = "STUB: not implemented"; return nil }

type mathExprLiteral struct {
	// If editor is matched then error
	Editor    *editor    `parser:"( @@"`
	Converter *converter `parser:"| @@"`
	Float     *float64   `parser:"| @Float"`
	Int       *int64     `parser:"| @Int"`
	Path      *path      `parser:"| @@ )"`
}

func (m *mathExprLiteral) accept(v grammarVisitor) { _ = "STUB: not implemented"; return }

type mathValue struct {
	UnaryOp       *mathOp          `parser:"@OpAddSub?"`
	Literal       *mathExprLiteral `parser:"( @@"`
	SubExpression *mathExpression  `parser:"| '(' @@ ')' )"`
}

func (m *mathValue) accept(v grammarVisitor) { _ = "STUB: not implemented"; return }

type opMultDivValue struct {
	Operator mathOp     `parser:"@OpMultDiv"`
	Value    *mathValue `parser:"@@"`
}

func (m *opMultDivValue) accept(v grammarVisitor) { _ = "STUB: not implemented"; return }

type addSubTerm struct {
	Left  *mathValue        `parser:"@@"`
	Right []*opMultDivValue `parser:"@@*"`
}

func (m *addSubTerm) accept(v grammarVisitor) { _ = "STUB: not implemented"; return }

type opAddSubTerm struct {
	Operator mathOp      `parser:"@OpAddSub"`
	Term     *addSubTerm `parser:"@@"`
}

func (r *opAddSubTerm) accept(v grammarVisitor) { _ = "STUB: not implemented"; return }

type mathExpression struct {
	Left  *addSubTerm     `parser:"@@"`
	Right []*opAddSubTerm `parser:"@@*"`
}

func (m *mathExpression) accept(v grammarVisitor) { _ = "STUB: not implemented"; return }

type mathOp int

const (
	add mathOp = iota
	sub
	mult
	div
)

var mathOpTable = map[string]mathOp{
	"+": add,
	"-": sub,
	"*": mult,
	"/": div,
}

func (m *mathOp) Capture(values []string) error { _ = "STUB: not implemented"; return nil }

func (m *mathOp) String() string { _ = "STUB: not implemented"; return "" }

type enumSymbol string

// buildLexer constructs a SimpleLexer definition.
// Note that the ordering of these rules matters.
// It's in a separate function so it can be easily tested alone (see lexer_test.go).
func buildLexer() *lexer.StatefulDefinition { _ = "STUB: not implemented"; return nil }

// grammarCustomError represents a grammar error in which the statement has a valid syntax
// according to the grammar's definition, but is still logically invalid.
type grammarCustomError struct {
	errs []error
}

// Error returns all errors messages separate by semicolons.
func (e *grammarCustomError) Error() string { _ = "STUB: not implemented"; return "" }

func (e *grammarCustomError) Unwrap() []error {
	_ = "STUB: not implemented"

	// grammarVisitor allows accessing the grammar AST nodes using the visitor pattern.
	return nil
}

type grammarVisitor interface {
	visitPath(v *path)
	visitEditor(v *editor)
	visitConverter(v *converter)
	visitValue(v *value)
	visitMathExprLiteral(v *mathExprLiteral)
}

// grammarCustomErrorsVisitor is used to execute custom validations on the grammar AST.
type grammarCustomErrorsVisitor struct {
	errs []error
}

func (g *grammarCustomErrorsVisitor) add(err error) { _ = "STUB: not implemented"; return }

func (g *grammarCustomErrorsVisitor) join() error { _ = "STUB: not implemented"; return nil }

func (*grammarCustomErrorsVisitor) visitPath(*path) { _ = "STUB: not implemented"; return }

func (*grammarCustomErrorsVisitor) visitValue(*value) { _ = "STUB: not implemented"; return }

func (*grammarCustomErrorsVisitor) visitConverter(*converter) { _ = "STUB: not implemented"; return }

func (g *grammarCustomErrorsVisitor) visitEditor(v *editor) { _ = "STUB: not implemented"; return }

func (g *grammarCustomErrorsVisitor) visitMathExprLiteral(v *mathExprLiteral) {
	_ = "STUB: not implemented"
	return
}

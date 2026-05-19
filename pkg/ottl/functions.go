// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottl // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"

import (
	"context"
	"reflect"
)

// PathExpressionParser is how a context provides OTTL access to all its Paths.
type PathExpressionParser[K any] func(Path[K]) (GetSetter[K], error)

// EnumParser is how a context provides OTTL access to all its Enums.
type EnumParser func(*EnumSymbol) (*Enum, error)

// Enum is how OTTL represents an enum's numeric value.
type Enum int64

// EnumSymbol is how OTTL represents an enum's string value.
type EnumSymbol string

func buildOriginalText(path *path) string { _ = "STUB: not implemented"; return "" }

func buildOriginalKeysText(keys []key) string { _ = "STUB: not implemented"; return "" }

func (p *Parser[K]) newPath(path *path) (*basePath[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Parser[K]) parsePathContext(path *path) (string, []field, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

// no pathContextNames means the Parser isn't handling the grammar path's context yet,
// so it falls back to the previous behavior with the path.Context value as the first
// path's segment.

func (p *Parser[K]) buildPathContextNamesText(path string) string {
	_ = "STUB: not implemented"
	return ""
}

// Path represents a chain of path parts in an OTTL statement, such as `body.string`.
// A Path has a name, and potentially a set of keys.
// If the path in the OTTL statement contains multiple parts (separated by a dot (`.`)), then the Path will have a pointer to the next Path.
type Path[K any] interface {
	// Context is the OTTL context name of this Path.
	Context() string

	// Name is the name of this segment of the path.
	Name() string

	// Next provides the next path segment for this Path.
	// Will return nil if there is no next path.
	Next() Path[K]

	// Keys provides the Keys for this Path.
	// Will return nil if there are no Keys.
	Keys() []Key[K]

	// String returns a string representation of this Path and the next Paths
	String() string
}

var _ Path[any] = &basePath[any]{}

type basePath[K any] struct {
	context      string
	name         string
	keys         []Key[K]
	nextPath     *basePath[K]
	fetched      bool
	fetchedKeys  bool
	originalText string
}

func (p *basePath[K]) Context() string { _ = "STUB: not implemented"; return "" }

func (p *basePath[K]) Name() string { _ = "STUB: not implemented"; return "" }

func (p *basePath[K]) Next() Path[K] { _ = "STUB: not implemented"; return nil }

func (p *basePath[K]) Keys() []Key[K] { _ = "STUB: not implemented"; return nil }

func (p *basePath[K]) String() string { _ = "STUB: not implemented"; return "" }

func (p *basePath[K]) isComplete() error { _ = "STUB: not implemented"; return nil }

func (p *Parser[K]) newKeys(keys []key) ([]Key[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Key represents a chain of keys in an OTTL statement, such as `attributes["foo"]["bar"]`.
// A Key has a String or Int, and potentially the next Key.
// If the path in the OTTL statement contains multiple keys, then the Key will have a pointer to the next Key.
type Key[K any] interface {
	// String returns a pointer to the Key's string value.
	// If the Key does not have a string value the returned value is nil.
	// If Key experiences an error retrieving the value it is returned.
	String(context.Context, K) (*string, error)

	// Int returns a pointer to the Key's int value.
	// If the Key does not have a int value the returned value is nil.
	// If Key experiences an error retrieving the value it is returned.
	Int(context.Context, K) (*int64, error)

	// ExpressionGetter returns a Getter to the expression, that can be
	// part of the path.
	// If the Key does not have an expression the returned value is nil.
	// If Key experiences an error retrieving the value it is returned.
	ExpressionGetter(context.Context, K) (Getter[K], error)
}

var _ Key[any] = &baseKey[any]{}

type baseKey[K any] struct {
	s *string
	i *int64
	g Getter[K]
}

func (k *baseKey[K]) String(_ context.Context, _ K) (*string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *baseKey[K]) Int(_ context.Context, _ K) (*int64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *baseKey[K]) ExpressionGetter(_ context.Context, _ K) (Getter[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Parser[K]) parsePath(ip *basePath[K]) (GetSetter[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Parser[K]) newFunctionCall(ed editor) (Expr[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// A nil value indicates the function takes no arguments.

// Pointer values are necessary to fulfill the Go reflection
// settability requirements. Non-pointer values are not
// modifiable through reflection.

func (p *Parser[K]) buildArgs(ed editor, argsVal reflect.Value) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *Parser[K]) buildSliceArg(argVal value, argType reflect.Type) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (p *Parser[K]) buildGetSetterFromPath(path *path) (GetSetter[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Handle interfaces that can be passed as arguments to OTTL functions.
func (p *Parser[K]) buildArg(argVal value, argType reflect.Type) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

type buildArgFunc func(value, reflect.Type) (any, error)

func buildSlice[T any](argVal value, argType reflect.Type, buildArg buildArgFunc, name string) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// optionalManager provides a way for the parser to handle Optional[T] structs
// without needing to know the concrete type of T, which is inaccessible through
// the reflect package.
// Would likely be resolved by https://github.com/golang/go/issues/54393.
type optionalManager interface {
	// set takes a non-reflection value and returns a reflect.Value of
	// an Optional[T] struct with this value set.
	set(val any) reflect.Value

	// get returns a reflect.Value value of the value contained within
	// an Optional[T]. This allows obtaining a reflect.Type for T.
	get() reflect.Value
}

// Optional is used to represent an optional function argument
type Optional[T any] struct {
	val      T
	hasValue bool
}

// This is called only by reflection.
func (Optional[T]) set(val any) reflect.Value {
	_ = "STUB: not implemented"
	return *new(reflect.Value)
}

// IsEmpty returns true if the Optional[T] does not contain a value.
func (o Optional[T]) IsEmpty() bool {
	_ = "STUB: not implemented"

	// Get returns the value contained in the Optional[T].
	return false
}

func (o Optional[T]) Get() T {
	_ = "STUB: not implemented"

	// GetOr returns the value contained in the Optional[T] if it exists,
	// otherwise it returns the default value provided.
	return *new(T)
}

func (o Optional[T]) GetOr(value T) T { _ = "STUB: not implemented"; return *new(T) }

func (o Optional[T]) get() reflect.Value {
	_ = "STUB: not implemented"
	// `(reflect.Value).Call` will create a reflect.Value containing a zero-valued T.
	// Trying to create a reflect.Value for T by calling reflect.TypeOf or
	// reflect.ValueOf on an empty T value creates an invalid reflect.Value object,
	// the `Call` method appears to do extra processing to capture the type.
	return *new(reflect.Value)
}

// NewTestingOptional allows creating an Optional with a value already populated for use in testing
// OTTL functions.
func NewTestingOptional[T any](val T) Optional[T] { _ = "STUB: not implemented"; return nil }

// typedGetter is like Getter, but with typed return values.
type typedGetter[K, V any] interface {
	Get(ctx context.Context, tCtx K) (V, error)
}

// mockLiteralGetter is a mock implementation of LiteralGetter that can be used for testing.
type mockLiteralGetter[K, V any] struct {
	valueGetter func(context.Context, K) (V, error)
}

func (m mockLiteralGetter[K, V]) Get(_ context.Context, _ K) (V, error) {
	_ = "STUB: not implemented"
	return *new(V), nil
}

// NewTestingLiteralGetter creates a mock literal getter for testing OTTL functions.
// Pass `literal` as true if the getter should be treated as a literal.
func NewTestingLiteralGetter[K, V any](literal bool, getter typedGetter[K, V]) (interface {
	typedGetter[K, V]
}, error,
) {
	_ = "STUB: not implemented"
	return nil, nil
}

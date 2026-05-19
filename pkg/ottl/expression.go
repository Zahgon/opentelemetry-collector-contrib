// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ottl // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"

import (
	"context"
	"time"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"golang.org/x/exp/constraints"
)

// ExprFunc is a function in OTTL
type ExprFunc[K any] func(ctx context.Context, tCtx K) (any, error)

// Expr is a struct that represents a function
type Expr[K any] struct {
	exprFunc ExprFunc[K]
}

// Eval invokes the OTTL function
func (e Expr[K]) Eval(ctx context.Context, tCtx K) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// Getter resolves a value at runtime without performing any type checking on the value that is returned.
type Getter[K any] interface {
	// Get retrieves a value of type 'Any' and returns an error if there are any issues during retrieval.
	Get(ctx context.Context, tCtx K) (any, error)
}

// Setter allows setting an untyped value on a predefined field within some data at runtime.
type Setter[K any] interface {
	// Set sets a value of type 'Any' and returns an error if there are any issues during the setting process.
	Set(ctx context.Context, tCtx K, val any) error
}

// GetSetter is an interface that combines the Getter and Setter interfaces.
// It should be used to represent the ability to both get and set a value.
type GetSetter[K any] interface {
	Getter[K]
	Setter[K]
}

// StandardGetSetter is a standard way to construct a GetSetter
type StandardGetSetter[K any] struct {
	Getter func(ctx context.Context, tCtx K) (any, error)
	Setter func(ctx context.Context, tCtx K, val any) error
}

func (path StandardGetSetter[K]) Get(ctx context.Context, tCtx K) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (path StandardGetSetter[K]) Set(ctx context.Context, tCtx K, val any) error {
	_ = "STUB: not implemented"
	return nil
}

type exprGetter[K any] struct {
	expr Expr[K]
	keys []key
}

func (g *exprGetter[K]) Get(ctx context.Context, tCtx K) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func getElementByIndex[T any](r []T, idx *int64) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

type listGetter[K any] struct {
	slice []Getter[K]
}

func newListGetter[K any](slice []Getter[K]) (Getter[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (l *listGetter[K]) Get(ctx context.Context, tCtx K) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

type mapGetter[K any] struct {
	mapValues map[string]Getter[K]
}

func newMapGetter[K any](mapValues map[string]Getter[K]) (Getter[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *mapGetter[K]) Get(ctx context.Context, tCtx K) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// PSliceGetSetter is a GetSetter that must interact with a pcommon.Slice
type PSliceGetSetter[K any] interface {
	Get(ctx context.Context, tCtx K) (pcommon.Slice, error)
	Set(ctx context.Context, tCtx K, val pcommon.Slice) error
}

// StandardPSliceGetSetter is a basic implementation of PSliceGetSetter
type StandardPSliceGetSetter[K any] struct {
	Getter func(ctx context.Context, tCtx K) (pcommon.Slice, error)
	Setter func(ctx context.Context, tCtx K, val any) error
}

func (path StandardPSliceGetSetter[K]) Get(ctx context.Context, tCtx K) (pcommon.Slice, error) {
	_ = "STUB: not implemented"
	return *new(pcommon.Slice), nil
}

func (path StandardPSliceGetSetter[K]) Set(ctx context.Context, tCtx K, val pcommon.Slice) error {
	_ = "STUB: not implemented"
	return nil
}

// newStandardPSliceGetSetter creates a new StandardPSliceGetSetter from a GetSetter[K],
// also validates getter on each use and checks if the GetSetter is a literalGetter.
func newStandardPSliceGetSetter[K any](getSetter GetSetter[K]) (PSliceGetSetter[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PSliceGetter is a Getter that must return a pcommon.Slice.
type PSliceGetter[K any] interface {
	Get(ctx context.Context, tCtx K) (pcommon.Slice, error)
}

// newStandardPSliceGetter creates a new StandardPSliceGetter from a Getter[K],
// also checking if the Getter is a literalGetter.
func newStandardPSliceGetter[K any](getter Getter[K]) (PSliceGetter[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// StandardPSliceGetter is a basic implementation of PSliceGetter
type StandardPSliceGetter[K any] struct {
	Getter func(ctx context.Context, tCtx K) (any, error)
}

// Get retrieves a pcommon.Slice value.
// If the value is not a pcommon.Slice a new TypeError is returned.
// If there is an error getting the value it will be returned.
func (g StandardPSliceGetter[K]) Get(ctx context.Context, tCtx K) (pcommon.Slice, error) {
	_ = "STUB: not implemented"
	return *new(pcommon.Slice), nil
}

// Handle common slice types returned by OTTL functions

func newPSliceFromIntegers[T constraints.Integer](source []T) (pcommon.Slice, error) {
	_ = "STUB: not implemented"
	return *new(pcommon.Slice), nil
}

func newPSliceFrom[T any](source []T, set func(target *pcommon.Value, value T)) (pcommon.Slice, error) {
	_ = "STUB: not implemented"
	return *new(pcommon.Slice), nil
}

// TypeError represents that a value was not an expected type.
type TypeError string

func (t TypeError) Error() string {
	_ = "STUB: not implemented"

	// StringGetter is a Getter that must return a string.
	return ""
}

type StringGetter[K any] interface {
	// Get retrieves a string value.
	Get(ctx context.Context, tCtx K) (string, error)
}

// newStandardStringGetter creates a new StandardStringGetter from a Getter[K],
// also checking if the Getter is a literalGetter.
func newStandardStringGetter[K any](getter Getter[K]) (StringGetter[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// StandardStringGetter is a basic implementation of StringGetter
type StandardStringGetter[K any] struct {
	Getter func(ctx context.Context, tCtx K) (any, error)
}

// Get retrieves a string value.
// If the value is not a string a new TypeError is returned.
// If there is an error getting the value it will be returned.
func (g StandardStringGetter[K]) Get(ctx context.Context, tCtx K) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// IntGetter is a Getter that must return an int64.
type IntGetter[K any] interface {
	// Get retrieves an int64 value.
	Get(ctx context.Context, tCtx K) (int64, error)
}

// newStandardIntGetter creates a new StandardIntGetter from a Getter[K],
// also checking if the Getter is a literalGetter.
func newStandardIntGetter[K any](getter Getter[K]) (IntGetter[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// StandardIntGetter is a basic implementation of IntGetter
type StandardIntGetter[K any] struct {
	Getter func(ctx context.Context, tCtx K) (any, error)
}

// Get retrieves an int64 value.
// If the value is not an int64 a new TypeError is returned.
// If there is an error getting the value it will be returned.
func (g StandardIntGetter[K]) Get(ctx context.Context, tCtx K) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// FloatGetter is a Getter that must return a float64.
type FloatGetter[K any] interface {
	// Get retrieves a float64 value.
	Get(ctx context.Context, tCtx K) (float64, error)
}

func newStandardFloatGetter[K any](getter Getter[K]) (FloatGetter[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// StandardFloatGetter is a basic implementation of FloatGetter
type StandardFloatGetter[K any] struct {
	Getter func(ctx context.Context, tCtx K) (any, error)
}

// Get retrieves a float64 value.
// If the value is not a float64 a new TypeError is returned.
// If there is an error getting the value it will be returned.
func (g StandardFloatGetter[K]) Get(ctx context.Context, tCtx K) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// BoolGetter is a Getter that must return a bool.
type BoolGetter[K any] interface {
	// Get retrieves a bool value.
	Get(ctx context.Context, tCtx K) (bool, error)
}

func newStandardBoolGetter[K any](getter Getter[K]) (BoolGetter[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// StandardBoolGetter is a basic implementation of BoolGetter
type StandardBoolGetter[K any] struct {
	Getter func(ctx context.Context, tCtx K) (any, error)
}

// Get retrieves a bool value.
// If the value is not a bool a new TypeError is returned.
// If there is an error getting the value it will be returned.
func (g StandardBoolGetter[K]) Get(ctx context.Context, tCtx K) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// FunctionGetter uses a function factory to return an instantiated function as an Expr.
type FunctionGetter[K any] interface {
	// Get returns a function as an Expr[K] built with the provided Arguments
	Get(args Arguments) (Expr[K], error)
}

// StandardFunctionGetter is a basic implementation of FunctionGetter.
type StandardFunctionGetter[K any] struct {
	FCtx FunctionContext
	Fact Factory[K]
}

// Get takes an Arguments struct containing arguments the caller wants passed to the
// function and instantiates the function with those arguments.
// If there is a mismatch between the function's signature and the arguments the caller
// wants to pass to the function, an error is returned.
func (g StandardFunctionGetter[K]) Get(args Arguments) (Expr[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PMapGetSetter is a GetSetter that must interact with a pcommon.Map
type PMapGetSetter[K any] interface {
	Get(ctx context.Context, tCtx K) (pcommon.Map, error)
	Set(ctx context.Context, tCtx K, val pcommon.Map) error
}

// StandardPMapGetSetter is a basic implementation of PMapGetSetter
type StandardPMapGetSetter[K any] struct {
	Getter func(ctx context.Context, tCtx K) (pcommon.Map, error)
	Setter func(ctx context.Context, tCtx K, val any) error
}

func (path StandardPMapGetSetter[K]) Get(ctx context.Context, tCtx K) (pcommon.Map, error) {
	_ = "STUB: not implemented"
	return *new(pcommon.Map), nil
}

func (path StandardPMapGetSetter[K]) Set(ctx context.Context, tCtx K, val pcommon.Map) error {
	_ = "STUB: not implemented"
	return nil
}

// PMapGetter is a Getter that must return a pcommon.Map.
type PMapGetter[K any] interface {
	// Get retrieves a pcommon.Map value.
	Get(ctx context.Context, tCtx K) (pcommon.Map, error)
}

// newStandardPMapGetter creates a new StandardPMapGetter from a Getter[K],
// also checking if the Getter is a literalGetter.
func newStandardPMapGetter[K any](getter Getter[K]) (PMapGetter[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// StandardPMapGetter is a basic implementation of PMapGetter
type StandardPMapGetter[K any] struct {
	Getter func(ctx context.Context, tCtx K) (any, error)
}

// Get retrieves a pcommon.Map value.
// If the value is not a pcommon.Map a new TypeError is returned.
// If there is an error getting the value it will be returned.
func (g StandardPMapGetter[K]) Get(ctx context.Context, tCtx K) (pcommon.Map, error) {
	_ = "STUB: not implemented"
	return *new(pcommon.Map), nil
}

// StringLikeGetter is a Getter that returns a string by converting the underlying value to a string if necessary.
type StringLikeGetter[K any] interface {
	// Get retrieves a string value.
	// Unlike `StringGetter`, the expectation is that the underlying value is converted to a string if possible.
	// If the value cannot be converted to a string, nil and an error are returned.
	// If the value is nil, nil is returned without an error.
	Get(ctx context.Context, tCtx K) (*string, error)
}

// newStandardStringLikeGetter creates a new StandardStringLikeGetter from a Getter[K],
// also checking if the Getter is a literalGetter.
func newStandardStringLikeGetter[K any](getter Getter[K]) (StringLikeGetter[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// StandardStringLikeGetter is a basic implementation of StringLikeGetter
type StandardStringLikeGetter[K any] struct {
	Getter func(ctx context.Context, tCtx K) (any, error)
}

func (g StandardStringLikeGetter[K]) Get(ctx context.Context, tCtx K) (*string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FloatLikeGetter is a Getter that returns a float64 by converting the underlying value to a float64 if necessary.
type FloatLikeGetter[K any] interface {
	// Get retrieves a float64 value.
	// Unlike `FloatGetter`, the expectation is that the underlying value is converted to a float64 if possible.
	// If the value cannot be converted to a float64, nil and an error are returned.
	// If the value is nil, nil is returned without an error.
	Get(ctx context.Context, tCtx K) (*float64, error)
}

func newStandardFloatLikeGetter[K any](getter Getter[K]) (FloatLikeGetter[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// StandardFloatLikeGetter is a basic implementation of FloatLikeGetter
type StandardFloatLikeGetter[K any] struct {
	Getter func(ctx context.Context, tCtx K) (any, error)
}

func (g StandardFloatLikeGetter[K]) Get(ctx context.Context, tCtx K) (*float64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// IntLikeGetter is a Getter that returns an int by converting the underlying value to an int if necessary
type IntLikeGetter[K any] interface {
	// Get retrieves an int value.
	// Unlike `IntGetter`, the expectation is that the underlying value is converted to an int if possible.
	// If the value cannot be converted to an int, nil and an error are returned.
	// If the value is nil, nil is returned without an error.
	Get(ctx context.Context, tCtx K) (*int64, error)
}

func newStandardIntLikeGetter[K any](getter Getter[K]) (IntLikeGetter[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// StandardIntLikeGetter is a basic implementation of IntLikeGetter
type StandardIntLikeGetter[K any] struct {
	Getter func(ctx context.Context, tCtx K) (any, error)
}

func (g StandardIntLikeGetter[K]) Get(ctx context.Context, tCtx K) (*int64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ByteSliceLikeGetter is a Getter that returns []byte by converting the underlying value to an []byte if necessary
type ByteSliceLikeGetter[K any] interface {
	// Get retrieves []byte value.
	// The expectation is that the underlying value is converted to []byte if possible.
	// If the value cannot be converted to []byte, nil and an error are returned.
	// If the value is nil, nil is returned without an error.
	Get(ctx context.Context, tCtx K) ([]byte, error)
}

func newStandardByteSliceLikeGetter[K any](getter Getter[K]) (ByteSliceLikeGetter[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// StandardByteSliceLikeGetter is a basic implementation of ByteSliceLikeGetter
type StandardByteSliceLikeGetter[K any] struct {
	Getter func(ctx context.Context, tCtx K) (any, error)
}

func (g StandardByteSliceLikeGetter[K]) Get(ctx context.Context, tCtx K) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// valueToBytes converts a value to a byte slice of length 8.
func valueToBytes(n any) ([]byte, error) {
	_ = "STUB: not implemented"
	// Create a buffer to hold the bytes
	return nil, nil
}

// Write the value to the buffer using binary.Write

// BoolLikeGetter is a Getter that returns a bool by converting the underlying value to a bool if necessary.
type BoolLikeGetter[K any] interface {
	// Get retrieves a bool value.
	// Unlike `BoolGetter`, the expectation is that the underlying value is converted to a bool if possible.
	// If the value cannot be converted to a bool, nil and an error are returned.
	// If the value is nil, nil is returned without an error.
	Get(ctx context.Context, tCtx K) (*bool, error)
}

func newStandardBoolLikeGetter[K any](getter Getter[K]) (BoolLikeGetter[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// StandardBoolLikeGetter is a basic implementation of BoolLikeGetter
type StandardBoolLikeGetter[K any] struct {
	Getter func(ctx context.Context, tCtx K) (any, error)
}

func (g StandardBoolLikeGetter[K]) Get(ctx context.Context, tCtx K) (*bool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Parser[K]) newGetter(val value) (Getter[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// In practice, can't happen since the DSL grammar guarantees one is set

func (p *Parser[K]) newGetterFromConverter(c converter) (Getter[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TimeGetter is a Getter that must return a time.Time.
type TimeGetter[K any] interface {
	// Get retrieves a time.Time value.
	Get(ctx context.Context, tCtx K) (time.Time, error)
}

func newStandardTimeGetter[K any](getter Getter[K]) (TimeGetter[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// StandardTimeGetter is a basic implementation of TimeGetter
type StandardTimeGetter[K any] struct {
	Getter func(ctx context.Context, tCtx K) (any, error)
}

// Get retrieves a time.Time value.
// If the value is not a time.Time, a new TypeError is returned.
// If there is an error getting the value it will be returned.
func (g StandardTimeGetter[K]) Get(ctx context.Context, tCtx K) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

// DurationGetter is a Getter that must return a time.Duration.
type DurationGetter[K any] interface {
	// Get retrieves a time.Duration value.
	Get(ctx context.Context, tCtx K) (time.Duration, error)
}

func newStandardDurationGetter[K any](getter Getter[K]) (DurationGetter[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// StandardDurationGetter is a basic implementation of DurationGetter
type StandardDurationGetter[K any] struct {
	Getter func(ctx context.Context, tCtx K) (any, error)
}

// Get retrieves an time.Duration value.
// If the value is not an time.Duration a new TypeError is returned.
// If there is an error getting the value it will be returned.
func (g StandardDurationGetter[K]) Get(ctx context.Context, tCtx K) (time.Duration, error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil
}

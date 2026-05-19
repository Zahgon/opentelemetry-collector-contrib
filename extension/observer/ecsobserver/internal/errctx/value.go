// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package errctx // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/ecsobserver/internal/errctx"

// ErrorWithValue indicates the error has some (could be just one) key value pairs
// attached to it during error wrapping.
type ErrorWithValue interface {
	error
	// Value returns a value attached to the error by key.
	// If the key does not exists, it returns nil, false.
	// The value saved can be nil, so it can also returns nil, true
	// to indicates a key exists but its value is nil.
	//
	// It does NOT do recursive Value calls (like context.Context).
	// For getting value from entire error chain, use ValueFrom.
	Value(key string) (v any, ok bool)
}

// WithValue attaches a single key value pair to a non nil error.
// If err is nil, it does nothing and return nil.
// key has to be a non empty string otherwise it will panic.
// val can be nil, and the existence of a key with nil value
// can be distinguished using the bool return value in Value/ValueFrom call.
//
// It is a good practice to define the key as a constant strange instead of inline literal.
//
//		const taskErrKey = "task"
//		return errctx.WithValue(taskErrKey, myTask)
//	 task, ok := errctx.ValueFrom(err, taskErrKey)
func WithValue(err error, key string, val any) error { _ = "STUB: not implemented"; return nil }

// panic because this should not happen and there is no good way to return error when dealing w/ error.
// This is also how context.WithValue is implemented.

// NOTE: we don't check if the value is nil because unlike context.Context
// our value methods return a bool to indicates if the key exists or not
// so we can allow user to save key with nil value, user's error inspection logic
// need to be aware of that.

// WithValues attaches multiple key value pairs. The behavior is similar to WithValue.
func WithValues(err error, kvs map[string]any) error { _ = "STUB: not implemented"; return nil }

// make a shallow copy, and hope the values in map are not map ...

// ValueFrom traverse entire error chain and returns the value
// from the first ErrorWithValue that contains the key.
// e.g. for an error created using errctx.WithValue(errctx.WithValue(base, "k", "v1"), "k", "v2")
// ValueFrom(err, "k") returns "v2".
func ValueFrom(err error, key string) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

// Check if this is a wrapped error
// I guess tail recursion should be optimized by compiler so we don't need to unroll it into for loop.

// valueError only contains one pair, which is common
type valueError struct {
	key   string
	val   any
	inner error
}

func (e *valueError) Error() string { _ = "STUB: not implemented"; return "" }

func (e *valueError) Value(key string) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

func (e *valueError) Unwrap() error {
	_ = "STUB: not implemented"

	// valuesError contains multiple pairs
	return nil
}

type valuesError struct {
	values map[string]any
	inner  error
}

func (e *valuesError) Error() string {
	_ = "STUB: not implemented"
	// NOTE: in order to have a consistent output, we sort the keys
	return ""
}

func (e *valuesError) Value(key string) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

func (e *valuesError) Unwrap() error { _ = "STUB: not implemented"; return nil }

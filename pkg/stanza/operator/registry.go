// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package operator // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"

// DefaultRegistry is a global registry of operator types to operator builders.
var DefaultRegistry = NewRegistry()

// Registry is used to track and retrieve known operator types
type Registry struct {
	operators map[string]func() Builder
}

// NewRegistry creates a new registry
func NewRegistry() *Registry { _ = "STUB: not implemented"; return nil }

// Register will register a function to an operator type.
// This function will return a builder for the supplied type.
func (r *Registry) Register(operatorType string, newBuilder func() Builder) {
	_ = "STUB: not implemented"
	return
}

// Lookup looks up a given operator type. Its second return value will
// be false if no builder is registered for that type.
func (r *Registry) Lookup(configType string) (func() Builder, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// Register will register an operator in the default registry
func Register(operatorType string, newBuilder func() Builder) { _ = "STUB: not implemented"; return }

// Lookup looks up a given operator type.Its second return value will
// be false if no builder is registered for that type.
func Lookup(configType string) (func() Builder, bool) { _ = "STUB: not implemented"; return nil, false }

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package payload // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/datadogextension/internal/payload"

// ModuleInfoJSON holds data on all modules in the collector
// It is built to make checking module info quicker when building active/configured components list
// (don't need to iterate through a whole list of modules, just do key/value pair in map)
type ModuleInfoJSON struct {
	components map[string]CollectorModule
}

func NewModuleInfoJSON() *ModuleInfoJSON { _ = "STUB: not implemented"; return nil }

func (*ModuleInfoJSON) getKey(typeStr, kindStr string) string { _ = "STUB: not implemented"; return "" }

func (m *ModuleInfoJSON) AddComponent(comp CollectorModule) { _ = "STUB: not implemented"; return }

// We don't ever expect two go modules to have the same type and kind
// as collector would not be able to distinguish between them for configuration
// and service/pipeline purposes.

func (m *ModuleInfoJSON) GetComponent(typeStr, kindStr string) (CollectorModule, bool) {
	_ = "STUB: not implemented"
	return *new(CollectorModule), false
}

func (m *ModuleInfoJSON) AddComponents(components []CollectorModule) {
	_ = "STUB: not implemented"
	return
}

func (m *ModuleInfoJSON) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (m *ModuleInfoJSON) GetFullComponentsList() []CollectorModule {
	_ = "STUB: not implemented"
	return nil
}

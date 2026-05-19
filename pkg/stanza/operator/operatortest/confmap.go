// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package operatortest // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/operatortest"

import (
	"testing"

	"go.opentelemetry.io/collector/confmap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
)

// ConfigUnmarshalTest is used for testing golden configs
type ConfigUnmarshalTests struct {
	DefaultConfig operator.Builder
	TestsFile     string
	Tests         []ConfigUnmarshalTest
	// prevent unkeyed literal initialization
	_ struct{}
}

// ConfigUnmarshalTest is used for testing golden configs
type ConfigUnmarshalTest struct {
	Name               string
	Expect             any
	ExpectUnmarshalErr bool
	ExpectBuildErrs    []error
	// prevent unkeyed literal initialization
	_ struct{}
}

// Run Unmarshals yaml files and compares them against the expected.
func (c ConfigUnmarshalTests) Run(t *testing.T) { _ = "STUB: not implemented"; return }

type anyOpConfig struct {
	Operator operator.Config `mapstructure:"operator"`
}

func newAnyOpConfig(opCfg operator.Builder) *anyOpConfig { _ = "STUB: not implemented"; return nil }

func (a *anyOpConfig) Unmarshal(component *confmap.Conf) error {
	_ = "STUB: not implemented"
	return nil
}

// ConfigBuilderTests is used for testing build failures
type ConfigBuilderTests struct {
	Tests []ConfigBuilderTest
	// prevent unkeyed literal initialization
	_ struct{}
}

// ConfigBuilderTest is used for testing build failures
type ConfigBuilderTest struct {
	Name       string
	Cfg        operator.Builder
	BuildError string
	// prevent unkeyed literal initialization
	_ struct{}
}

// Run Build on a malformed config and expect an error.
func (c ConfigBuilderTests) Run(t *testing.T) { _ = "STUB: not implemented"; return }

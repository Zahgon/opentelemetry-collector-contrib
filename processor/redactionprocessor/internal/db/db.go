// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package db // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/redactionprocessor/internal/db"

import (
	"go.uber.org/zap"
)

type Obfuscator struct {
	obfuscators                []databaseObfuscator
	processAttributesEnabled   bool
	logger                     *zap.Logger
	allowFallbackWithoutSystem bool
	DBSystem                   string
}

func createAttributes(attributes []string) map[string]bool { _ = "STUB: not implemented"; return nil }

func NewObfuscator(cfg DBSanitizerConfig, logger *zap.Logger) *Obfuscator {
	_ = "STUB: not implemented"
	return nil
}

// Not part of semantic conventions

func (o *Obfuscator) Obfuscate(s string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func (o *Obfuscator) ObfuscateAttribute(attributeValue, attributeKey string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (o *Obfuscator) obfuscateSequentially(attributeValue, attributeKey string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (o *Obfuscator) HasSpecificAttributes() bool { _ = "STUB: not implemented"; return false }

func (o *Obfuscator) HasObfuscators() bool { _ = "STUB: not implemented"; return false }

func (o *Obfuscator) ObfuscateWithSystem(val, dbSystem string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func createSystems(systems []string) map[string]bool { _ = "STUB: not implemented"; return nil }

func newDBAttributes(attributes, systems []string) dbAttributes {
	_ = "STUB: not implemented"
	return *new(dbAttributes)
}

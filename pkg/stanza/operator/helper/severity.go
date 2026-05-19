// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package helper // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/entry"
)

// SeverityParser is a helper that parses severity onto an entry.
type SeverityParser struct {
	ParseFrom     entry.Field
	Mapping       severityMap
	overwriteText bool
}

// Parse will parse severity from a field and attach it to the entry
func (p *SeverityParser) Parse(ent *entry.Entry) error { _ = "STUB: not implemented"; return nil }

type severityMap map[string]entry.Severity

// accepts various stringifyable input types and returns
//  1. severity level if found, or default level
//  2. string version of input value
//  3. error if invalid input type
func (m severityMap) find(value any) (entry.Severity, string, error) {
	_ = "STUB: not implemented"
	return *new(entry.Severity), "", nil
}

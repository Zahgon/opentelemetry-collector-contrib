// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package sqlserverreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/sqlserverreceiver"

import (
	"github.com/DataDog/datadog-agent/pkg/obfuscate"
)

var xmlPlanObfuscationAttrs = []string{
	"StatementText",
	"ConstValue",
	"ScalarString",
	"ParameterCompiledValue",
}

type obfuscator obfuscate.Obfuscator

func newObfuscator() *obfuscator { _ = "STUB: not implemented"; return nil }

func (o *obfuscator) obfuscateSQLString(sql string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// obfuscateXMLPlan obfuscates SQL text & parameters from the provided SQL Server XML Plan
func (o *obfuscator) obfuscateXMLPlan(rawPlan string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

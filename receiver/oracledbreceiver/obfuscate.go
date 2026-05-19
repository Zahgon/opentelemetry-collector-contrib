// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package oracledbreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/oracledbreceiver"

import (
	"github.com/DataDog/datadog-agent/pkg/obfuscate"
)

type obfuscator obfuscate.Obfuscator

func newObfuscator() *obfuscator { _ = "STUB: not implemented"; return nil }

func (o *obfuscator) obfuscateSQLString(sql string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

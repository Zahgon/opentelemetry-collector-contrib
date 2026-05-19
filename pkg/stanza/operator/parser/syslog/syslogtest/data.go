// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package syslogtest // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/parser/syslog/syslogtest"

import (
	"time"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/entry"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/parser/syslog"
)

// This is the name of a test which requires setting the PreserveWhitespace flags.
const RFC6587OctetCountingPreserveSpaceTest = "RFC6587 Octet Counting Preserve Space"

type Case struct {
	Name   string
	Config *syslog.Config
	Input  *entry.Entry
	Expect *entry.Entry

	// These signal if a test is valid for UDP and/or TCP protocol
	ValidForTCP bool
	ValidForUDP bool
}

func testLocations() (map[string]*time.Location, error) { _ = "STUB: not implemented"; return nil, nil }

func CreateCases(basicConfig func() *syslog.Config) ([]Case, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// We need to build out the Non-Transparent-Framing body to ensure we control the Trailer byte

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package syslog // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/input/syslog"

import (
	"bufio"

	"golang.org/x/text/encoding"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/input/tcp"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/input/udp"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/parser/syslog"
)

// Input is an operator that listens for log entries over tcp.
type Input struct {
	helper.InputOperator
	tcp    *tcp.Input
	udp    *udp.Input
	parser *syslog.Parser
}

// Start will start listening for log entries over tcp or udp.
func (i *Input) Start(p operator.Persister) error { _ = "STUB: not implemented"; return nil }

// Stop will stop listening for messages.
func (i *Input) Stop() error { _ = "STUB: not implemented"; return nil }

// SetOutputs will set the outputs of the internal syslog parser.
func (i *Input) SetOutputs(operators []operator.Operator) error {
	_ = "STUB: not implemented"
	return nil
}

func OctetSplitFuncBuilder(_ encoding.Encoding) (bufio.SplitFunc, error) {
	_ = "STUB: not implemented"
	return *new(bufio.SplitFunc), nil
}

func newOctetFrameSplitFunc(flushAtEOF bool) bufio.SplitFunc {
	_ = "STUB: not implemented"
	return *new(bufio.SplitFunc)
}

// Flush if no more data is expected

// Remove the delimiter (space) between length and log, and parse the length

// This should not be possible because the regex matched.
// However, return an error just in case.

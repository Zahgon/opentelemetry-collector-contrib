// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package protocol // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/carbonreceiver/protocol"

// PlaintextConfig holds the configuration for the plaintext parser.
type PlaintextConfig struct{}

var _ ParserConfig = (*PlaintextConfig)(nil)

// BuildParser creates a new Parser instance that receives plaintext
// Carbon data.
func (*PlaintextConfig) BuildParser() (Parser, error) {
	_ = "STUB: not implemented"
	return *new(Parser), nil
}

// plaintextPathParser converts a line of https://graphite.readthedocs.io/en/latest/feeding-carbon.html#the-plaintext-protocol,
// treating tags per spec at https://graphite.readthedocs.io/en/latest/tags.html#carbon.
type plaintextPathParser struct{}

// parsePath converts the <metric_path> of a Carbon line (see Parse function for
// description of the full line). The metric path is expected to be in the
// following format:
//
//	<metric_name>[;tag0;...;tagN]
//
// <metric_name> is the name of the metric and terminates either at the first ';'
// or at the end of the path.
//
// tag is of the form "key=val", where key can contain any char except ";!^=" and
// val can contain any char except ";~".
func (*plaintextPathParser) parsePath(path string, parsedPath *parsedPath) error {
	_ = "STUB: not implemented"
	return nil
}

// No tags, no more work here.

// Empty tags, nothing to do.

// If value is empty, ie.: tag == "k=", this will return "".

func plaintextDefaultConfig() ParserConfig { _ = "STUB: not implemented"; return *new(ParserConfig) }

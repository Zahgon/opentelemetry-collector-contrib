// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package windows // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/input/windows"

func parseSecurity(message string) (string, map[string]any) {
	_ = "STUB: not implemented"
	return "", nil
}

// First line is expected to be the first return value

// standalone key/value pair with empty value

// standalone key/value pair

// value was first in a list

type messageProcessor struct {
	lines []*parsedLine
	ptr   int
}

func (mp *messageProcessor) consumeSubsection(depth int) map[string]any {
	_ = "STUB: not implemented"
	return nil
}

// standalone key/value pair with missing value

func (mp *messageProcessor) consumeSublist(depth int) []string {
	_ = "STUB: not implemented"
	return nil
}

// not expected, but handle

// not expected

type parsedLine struct {
	t lineType
	i int
	k string
	v string
}

type lineType int

const (
	emptyType lineType = iota
	keyType
	valueType
	pairType
)

func newMessageProcessor(message string) *messageProcessor { _ = "STUB: not implemented"; return nil }

func parse(line string) *parsedLine { _ = "STUB: not implemented"; return nil }

// return next line and increment position
func (mp *messageProcessor) next() *parsedLine { _ = "STUB: not implemented"; return nil }

// return next line but do not increment position
func (mp *messageProcessor) peek() *parsedLine { _ = "STUB: not implemented"; return nil }

// just increment position
func (mp *messageProcessor) step() { _ = "STUB: not implemented"; return }

func (mp *messageProcessor) hasNext() bool { _ = "STUB: not implemented"; return false }

func (mp *messageProcessor) hasNextIndented(minDepth int) bool {
	_ = "STUB: not implemented"
	return false
}

func countIndent(line string) int { _ = "STUB: not implemented"; return 0 }

func parseKeyValue(line string) (string, string) { _ = "STUB: not implemented"; return "", "" }

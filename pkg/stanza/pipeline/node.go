// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package pipeline // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/pipeline"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
)

// OperatorNode is a basic node that represents an operator in a pipeline.
type OperatorNode struct {
	operator  operator.Operator
	outputIDs map[string]int64
	id        int64
}

// Operator returns the operator of the node.
func (b OperatorNode) Operator() operator.Operator {
	_ = "STUB: not implemented"

	// ID returns the node id.
	return *new(operator.Operator)
}

func (b OperatorNode) ID() int64 {
	_ = "STUB: not implemented"

	// DOTID returns the id used to represent this node in a dot graph.
	return 0
}

func (b OperatorNode) DOTID() string { _ = "STUB: not implemented"; return "" }

// OutputIDs returns a map of output operator ids to node ids.
func (b OperatorNode) OutputIDs() map[string]int64 {
	_ = "STUB: not implemented"

	// createOperatorNode will create an operator node.
	return nil
}

func createOperatorNode(operator operator.Operator) OperatorNode {
	_ = "STUB: not implemented"
	return *new(OperatorNode)
}

// createNodeID generates a node id from an operator id.
func createNodeID(operatorID string) int64 { _ = "STUB: not implemented"; return 0 }

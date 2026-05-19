// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package pipeline // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/pipeline"

import (
	"errors"
	"sync"

	"gonum.org/v1/gonum/graph/simple"
	"gonum.org/v1/gonum/graph/topo"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
)

var _ Pipeline = (*DirectedPipeline)(nil)

var (
	errAlreadyStarted = errors.New("pipeline already started")
	errAlreadyStopped = errors.New("pipeline already stopped")
)

// DirectedPipeline is a pipeline backed by a directed graph
type DirectedPipeline struct {
	Graph     *simple.DirectedGraph
	startOnce sync.Once
	stopOnce  sync.Once
}

// Start will start the operators in a pipeline in reverse topological order
func (p *DirectedPipeline) Start(persister operator.Persister) error {
	_ = "STUB: not implemented"
	return nil
}

// Stop will stop the operators in a pipeline in topological order
func (p *DirectedPipeline) Stop() error { _ = "STUB: not implemented"; return nil }

func (p *DirectedPipeline) start(persister operator.Persister) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *DirectedPipeline) stop() error { _ = "STUB: not implemented"; return nil }

// Render will render the pipeline as a dot graph
func (p *DirectedPipeline) Render() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Operators returns a slice of operators that make up the pipeline graph
func (p *DirectedPipeline) Operators() []operator.Operator { _ = "STUB: not implemented"; return nil }

// If for some unexpected reason an Unorderable error is returned,
// when using topo.Sort, return the list without ordering

// addNodes will add operators as nodes to the supplied graph.
func addNodes(graph *simple.DirectedGraph, operators []operator.Operator) error {
	_ = "STUB: not implemented"
	return nil
}

// connectNodes will connect the nodes in the supplied graph.
func connectNodes(graph *simple.DirectedGraph) error { _ = "STUB: not implemented"; return nil }

// connectNode will connect a node to its outputs in the supplied graph.
func connectNode(graph *simple.DirectedGraph, inputNode OperatorNode) error {
	_ = "STUB: not implemented"
	return nil
}

// setOperatorOutputs will set the outputs on operators that can output.
func setOperatorOutputs(operators []operator.Operator) error { _ = "STUB: not implemented"; return nil }

// NewDirectedPipeline creates a new directed pipeline
func NewDirectedPipeline(operators []operator.Operator) (*DirectedPipeline, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func unorderableToCycles(err topo.Unorderable) string { _ = "STUB: not implemented"; return "" }

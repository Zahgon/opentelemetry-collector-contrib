// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0
package cluster // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/aerospikereceiver/cluster"

import (
	as "github.com/aerospike/aerospike-client-go/v8"
)

type SubsetCluster struct {
	nodes []Node
}

type nodeFactoryFunc func(*as.ClientPolicy, *as.Host, bool) (Node, error)

func NewSubsetCluster(policy *as.ClientPolicy, hosts []*as.Host, authEnabled bool) (*SubsetCluster, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newSubsetCluster(policy *as.ClientPolicy, hosts []*as.Host, authEnabled bool, nodeFact nodeFactoryFunc) (*SubsetCluster, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// this is only used with 1 node for now (when collect-cluster-metrics is false)

func (c *SubsetCluster) Close() { _ = "STUB: not implemented"; return }

func (c *SubsetCluster) GetNodes() []Node { _ = "STUB: not implemented"; return nil }

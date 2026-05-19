// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0
package cluster // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/aerospikereceiver/cluster"

import (
	as "github.com/aerospike/aerospike-client-go/v8"
)

// asclient interface is for mocking
type asclient interface {
	GetNodes() []*as.Node
	Close()
}

// wrap aerospike Cluster so we can return node interfaces
type Cluster struct {
	conn asclient
}

func NewCluster(policy *as.ClientPolicy, hosts []*as.Host) (*Cluster, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Cluster) GetNodes() []Node { _ = "STUB: not implemented"; return nil }

func (c *Cluster) Close() { _ = "STUB: not implemented"; return }

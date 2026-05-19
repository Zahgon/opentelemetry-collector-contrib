// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0
package cluster // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/aerospikereceiver/cluster"

import (
	"time"

	as "github.com/aerospike/aerospike-client-go/v8"
)

// asconn is used to mock aerospike connections
type asconn interface {
	RequestInfo(...string) (map[string]string, as.Error)
	Login(*as.ClientPolicy) as.Error
	Close()
	SetTimeout(time.Time, time.Duration) as.Error
}

type Node interface {
	RequestInfo(*as.InfoPolicy, ...string) (map[string]string, as.Error)
	GetName() string
	Close()
}

// connNode is for single node scraping
type connNode struct {
	conn   asconn
	policy *as.ClientPolicy
	name   string
}

type connFactoryFunc func(*as.ClientPolicy, *as.Host) (asconn, as.Error)

func newASConn(policy *as.ClientPolicy, host *as.Host) (asconn, as.Error) {
	_ = "STUB: not implemented"
	return *new(asconn), *new(as.Error)
}

func newConnNode(policy *as.ClientPolicy, host *as.Host, authEnabled bool) (Node, error) {
	_ = "STUB: not implemented"
	return *new(Node), nil
}

func _newConnNode(policy *as.ClientPolicy, host *as.Host, authEnabled bool, connF connFactoryFunc) (Node, error) {
	_ = "STUB: not implemented"
	return *new(Node), nil
}

// Set deadline to 0 (inf) so we can always reuse this connection

func (n *connNode) RequestInfo(_ *as.InfoPolicy, commands ...string) (map[string]string, as.Error) {
	_ = "STUB: not implemented"
	return nil, *new(as.Error)
}

// Try to login and get a new session

func (n *connNode) GetName() string { _ = "STUB: not implemented"; return "" }

func (n *connNode) Close() { _ = "STUB: not implemented"; return }

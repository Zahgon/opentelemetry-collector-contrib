// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package aerospikereceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/aerospikereceiver"

import (
	"crypto/tls"
	"time"

	as "github.com/aerospike/aerospike-client-go/v8"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/aerospikereceiver/cluster"
)

var defaultNodeInfoCommands = []string{
	"node",
	"statistics",
}

// nodeName: metricName: stats
type clusterInfo = map[string]map[string]string

// Aerospike is the interface that provides information about a given node
type Aerospike interface {
	// NamespaceInfo gets information about a specific namespace
	NamespaceInfo() namespaceInfo
	// Info gets high-level information about the node/system.
	Info() clusterInfo
	// Close closes the connection to the Aerospike node
	Close()
}

type clientConfig struct {
	host                  *as.Host
	username              string
	password              string
	timeout               time.Duration
	logger                *zap.SugaredLogger
	collectClusterMetrics bool
	tls                   *tls.Config
}

type nodeGetter interface {
	GetNodes() []cluster.Node
	Close()
}

type defaultASClient struct {
	cluster nodeGetter
	policy  *as.ClientPolicy   // Timeout and authentication information
	logger  *zap.SugaredLogger // logs malformed metrics in responses
}

type nodeGetterFactoryFunc func(cfg *clientConfig, policy *as.ClientPolicy, authEnabled bool) (nodeGetter, error)

func nodeGetterFactory(cfg *clientConfig, policy *as.ClientPolicy, authEnabled bool) (nodeGetter, error) {
	_ = "STUB: not implemented"
	return *new(nodeGetter), nil
}

// newASClient creates a new defaultASClient connected to the given host and port
// If collectClusterMetrics is true, the client will connect to and tend all nodes in the cluster
// If username and password aren't blank, they're used to authenticate
func newASClient(cfg *clientConfig, ngf nodeGetterFactoryFunc) (*defaultASClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// enable TLS

// useNodeFunc maps a nodeFunc to all the client's nodes
func (c *defaultASClient) useNodeFunc(nf nodeFunc) clusterInfo {
	_ = "STUB: not implemented"
	return *new(clusterInfo)
}

// metricName: stat
// may be used as, commandName: stat
type metricsMap = map[string]string

// Info returns a clusterInfo map of node names to metricMaps
// it uses the info commands defined in defaultNodeInfoCommands
func (c *defaultASClient) Info() clusterInfo {
	_ = "STUB: not implemented"
	return *

	// NOTE this discards the command names
	new(clusterInfo)
}

// nodeName: namespaceName: metricName: stats
type namespaceInfo = map[string]map[string]map[string]string

// NamespaceInfo returns a namespaceInfo map
// the map contains the results of the "namespace/<name>" info command
// for all nodes' namespaces
func (c *defaultASClient) NamespaceInfo() namespaceInfo {
	_ = "STUB: not implemented"
	return *new(namespaceInfo)
}

// ns == "namespace/<namespaceName>"

// Close closes the client's connections to all nodes
func (c *defaultASClient) Close() {
	_ = "STUB: not implemented"

	// mapNodeInfoFunc maps a nodeFunc to all nodes in the list in parallel
	// if an error occurs during any of the nodeFuncs' execution, it is logged but not returned
	// the return value is a clusterInfo map from node name to command to unparsed metric string
	return
}

func mapNodeInfoFunc(nodes []cluster.Node, nodeF nodeFunc, policy *as.InfoPolicy, logger *zap.SugaredLogger) clusterInfo {
	_ = "STUB: not implemented"
	return *new(clusterInfo)
}

// node wise info functions

// nodeFunc is a function that requests info commands from a node
// and returns a metricsMap from command to metric string
type nodeFunc func(n cluster.Node, policy *as.InfoPolicy) (metricsMap, error)

// namespaceNames is used by nodeFuncs to get the names of all namespaces ona node
func namespaceNames(n cluster.Node, policy *as.InfoPolicy) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// allNodeInfo returns the results of defaultNodeInfoCommands for a node
func allNodeInfo(n cluster.Node, policy *as.InfoPolicy) (metricsMap, error) {
	_ = "STUB: not implemented"
	return *new(metricsMap), nil
}

// allNamespaceInfo returns the results of namespace/%s for each namespace on the node
func allNamespaceInfo(n cluster.Node, policy *as.InfoPolicy) (metricsMap, error) {
	_ = "STUB: not implemented"
	return *new(metricsMap), nil
}

func parseStats(defaultKey, s, sep string) metricsMap {
	_ = "STUB: not implemented"
	return *new(metricsMap)
}

// mergeMetricsMap merges values from rm into lm
// logs a warning if a duplicate key is found
func mergeMetricsMap(lm, rm metricsMap, logger *zap.SugaredLogger) {
	_ = "STUB: not implemented"
	return
}

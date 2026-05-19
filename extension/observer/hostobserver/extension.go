// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package hostobserver // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/hostobserver"

import (
	"context"

	"github.com/shirou/gopsutil/v4/net"
	"github.com/shirou/gopsutil/v4/process"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/endpointswatcher"
)

type hostObserver struct {
	*endpointswatcher.EndpointsWatcher
}

type endpointsLister struct {
	logger       *zap.Logger
	observerName string

	// For testing
	getConnections        func() ([]net.ConnectionStat, error)
	getProcess            func(pid int32) (*process.Process, error)
	collectProcessDetails func(proc *process.Process) (*processDetails, error)
}

var _ extension.Extension = (*hostObserver)(nil)

func newObserver(params extension.Settings, config *Config) (extension.Extension, error) {
	_ = "STUB: not implemented"
	return *new(extension.Extension), nil
}

func (*hostObserver) Start(context.Context, component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *hostObserver) Shutdown(context.Context) error { _ = "STUB: not implemented"; return nil }

func (e endpointsLister) ListEndpoints() []observer.Endpoint { _ = "STUB: not implemented"; return nil }

func getConnections() (conns []net.ConnectionStat, err error) {
	_ = "STUB: not implemented"
	// Skip UID lookup since it's not used by the observer, the method
	// is available only on linux. See https://github.com/shirou/gopsutil/pull/783
	// for details.
	return nil, nil
}

func (e endpointsLister) collectEndpoints(conns []net.ConnectionStat) []observer.Endpoint {
	_ = "STUB: not implemented"
	return nil
}

// UDP doesn't have any status

// UDP is "listening" when it has a remote port of 0

// PID of 0 means that the listening file descriptor couldn't be mapped
// back to a process's set of open file descriptors in /proc. Collect these
// endpoints even though there's no process metadata available so users can
// still do discovery rules on such sockets.

// TODO: Move this field to observer.Endpoint and
// update receiver_creator to filter IPv4/IPv6.

// TODO: Move this field to observer.Endpoint and
// update receiver_creator to filter IPv4/IPv6.

type connectionDetails struct {
	ip        string
	isIPv6    bool
	port      uint16
	target    string
	transport observer.Transport
}

func collectConnectionDetails(c *net.ConnectionStat) connectionDetails {
	_ = "STUB: not implemented"

	// An IP addr of 0.0.0.0 (or "*" on darwin) means it listens on all
	// interfaces, including localhost, so use that since we can't
	// actually connect to 0.0.0.0.
	return *new(connectionDetails)
}

type processDetails struct {
	name string
	args string
}

func collectProcessDetails(proc *process.Process) (*processDetails, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func portTypeToProtocol(t uint32) observer.Transport {
	_ = "STUB: not implemented"
	return *new(observer.Transport)
}

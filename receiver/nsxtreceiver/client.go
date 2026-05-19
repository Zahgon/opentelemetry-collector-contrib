// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package nsxtreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/nsxtreceiver"

import (
	"context"
	"errors"
	"net/http"
	"net/url"

	"go.opentelemetry.io/collector/component"
	"go.uber.org/zap"

	dm "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/nsxtreceiver/internal/model"
)

var _ Client = &nsxClient{}

// Client is a way of interacting with the NSX REST API
type Client interface {
	TransportNodes(ctx context.Context) ([]dm.TransportNode, error)
	ClusterNodes(ctx context.Context) ([]dm.ClusterNode, error)
	NodeStatus(ctx context.Context, nodeID string, class nodeClass) (*dm.NodeStatus, error)
	Interfaces(ctx context.Context, nodeID string, class nodeClass) ([]dm.NetworkInterface, error)
	InterfaceStatus(ctx context.Context, nodeID, interfaceID string, class nodeClass) (*dm.NetworkInterfaceStats, error)
}

type nsxClient struct {
	config   *Config
	client   *http.Client
	endpoint *url.URL
	logger   *zap.Logger
}

var errUnauthorized = errors.New("STATUS 403, unauthorized")

func newClient(ctx context.Context, c *Config, settings component.TelemetrySettings, host component.Host, logger *zap.Logger) (*nsxClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *nsxClient) TransportNodes(ctx context.Context) ([]dm.TransportNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *nsxClient) ClusterNodes(ctx context.Context) ([]dm.ClusterNode, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *nsxClient) NodeStatus(ctx context.Context, nodeID string, class nodeClass) (*dm.NodeStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *nsxClient) Interfaces(
	ctx context.Context,
	nodeID string,
	class nodeClass,
) ([]dm.NetworkInterface, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *nsxClient) InterfaceStatus(
	ctx context.Context,
	nodeID, interfaceID string,
	class nodeClass,
) (*dm.NetworkInterfaceStats, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *nsxClient) doRequest(ctx context.Context, path string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (*nsxClient) nodeStatusEndpoint(class nodeClass, nodeID string) string {
	_ = "STUB: not implemented"
	return ""
}

func (*nsxClient) interfacesEndpoint(class nodeClass, nodeID string) string {
	_ = "STUB: not implemented"
	return ""
}

func (*nsxClient) interfaceStatusEndpoint(class nodeClass, nodeID, interfaceID string) string {
	_ = "STUB: not implemented"
	return ""
}

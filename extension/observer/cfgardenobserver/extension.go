// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package cfgardenobserver // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/cfgardenobserver"

import (
	"context"
	"sync"

	"code.cloudfoundry.org/garden"
	"github.com/cloudfoundry/go-cfclient/v3/client"
	"github.com/cloudfoundry/go-cfclient/v3/resource"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/endpointswatcher"
)

const (
	propertiesAppIDKey     = "network.app_id"
	propertiesPortsKey     = "network.ports"
	propertiesLogConfigKey = "log_config"
	logConfigTagsKey       = "tags"
	containerStateActive   = "active"
)

type cfGardenObserver struct {
	*endpointswatcher.EndpointsWatcher
	config   *Config
	doneChan chan struct{}
	logger   *zap.Logger
	once     *sync.Once

	garden garden.Client
	cf     *client.Client

	containerMu sync.RWMutex
	containers  map[string]garden.ContainerInfo

	appMu sync.RWMutex
	apps  map[string]*resource.App
}

var _ extension.Extension = (*cfGardenObserver)(nil)

func newObserver(config *Config, logger *zap.Logger) (extension.Extension, error) {
	_ = "STUB: not implemented"
	return *new(extension.Extension), nil
}

func (g *cfGardenObserver) SyncApps() error { _ = "STUB: not implemented"; return nil }

func (g *cfGardenObserver) App(info garden.ContainerInfo) (*resource.App, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *cfGardenObserver) Start(_ context.Context, _ component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (g *cfGardenObserver) Shutdown(_ context.Context) error { _ = "STUB: not implemented"; return nil }

func (g *cfGardenObserver) ListEndpoints() []observer.Endpoint {
	_ = "STUB: not implemented"
	return nil
}

// containerEndpoints generates a list of observer.Endpoint for a container,
// this is because a container might have more than one exposed ports
func (g *cfGardenObserver) containerEndpoints(handle string, info garden.ContainerInfo) []observer.Endpoint {
	_ = "STUB: not implemented"
	return nil
}

func (g *cfGardenObserver) containerLabels(info garden.ContainerInfo, app *resource.App) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

// The info.Properties contains a key called "log_config", which
// has contents that look like the following JSON encoded string:
//
//	{
//	  "guid": "aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee",
//	  "index": 0,
//	  "source_name": "CELL",
//	  "tags": {
//	    "app_id": "aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee",
//	    "app_name": "example-app",
//	    "instance_id": "0",
//	    "organization_id": "11111111-2222-3333-4444-555555555555",
//	    "organization_name": "example-org",
//	    "process_id": "aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee",
//	    "process_instance_id": "abcdef12-3456-7890-abcd-ef1234567890",
//	    "process_type": "web",
//	    "source_id": "aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee",
//	    "space_id": "99999999-8888-7777-6666-555555555555",
//	    "space_name": "example-space"
//	  }
//	}
//
// We parse only the tags into a map, to be used as labels
func parseTags(info garden.ContainerInfo) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newCfClient(cfConfig CfConfig) (*client.Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *cfGardenObserver) updateContainerCache(infos map[string]garden.ContainerInfo) {
	_ = "STUB: not implemented"
	return
}

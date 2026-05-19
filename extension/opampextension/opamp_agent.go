// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package opampextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/opampextension"

import (
	"context"
	"strings"
	"sync"

	"github.com/google/uuid"
	"github.com/open-telemetry/opamp-go/client"
	"github.com/open-telemetry/opamp-go/client/types"
	"github.com/open-telemetry/opamp-go/protobufs"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/component/componentstatus"
	"go.opentelemetry.io/collector/confmap"
	"go.opentelemetry.io/collector/extension"
	"go.opentelemetry.io/collector/extension/extensioncapabilities"
	"go.opentelemetry.io/collector/service"
	conventions "go.opentelemetry.io/otel/semconv/v1.40.0"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/opampcustommessages"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/status"
)

type statusAggregator interface {
	Subscribe(scope status.Scope, verbosity status.Verbosity) (<-chan *status.AggregateStatus, status.UnsubscribeFunc)
	RecordStatus(source *componentstatus.InstanceID, event *componentstatus.Event)
}

type eventSourcePair struct {
	source *componentstatus.InstanceID
	event  *componentstatus.Event
}

type opampAgent struct {
	cfg    *Config
	logger *zap.Logger

	serviceName       string
	serviceVersion    string
	serviceInstanceID string
	resourceAttrs     map[string]string

	instanceUID uuid.UUID

	eclk            sync.RWMutex
	effectiveConfig *confmap.Conf

	// lifetimeCtx is canceled on Stop of the component
	lifetimeCtx       context.Context
	lifetimeCtxCancel context.CancelFunc

	reportFunc func(*componentstatus.Event)

	capabilities Capabilities

	agentDescription    *protobufs.AgentDescription
	availableComponents *protobufs.AvailableComponents

	opampClient client.OpAMPClient

	customCapabilityRegistry *customCapabilityRegistry

	statusAggregator     statusAggregator
	statusSubscriptionWg *sync.WaitGroup
	componentHealthWg    *sync.WaitGroup
	startTimeUnixNano    uint64
	componentStatusCh    chan *eventSourcePair
	readyCh              chan struct{}
}

var (
	_ opampcustommessages.CustomCapabilityRegistry = (*opampAgent)(nil)
	_ extensioncapabilities.Dependent              = (*opampAgent)(nil)
	_ extensioncapabilities.ConfigWatcher          = (*opampAgent)(nil)
	_ extensioncapabilities.PipelineWatcher        = (*opampAgent)(nil)
	_ componentstatus.Watcher                      = (*opampAgent)(nil)

	// identifyingAttributes is the list of semantic convention keys that are used
	// for the agent description's identifying attributes.
	identifyingAttributes = map[string]struct{}{
		string(conventions.ServiceNameKey):       {},
		string(conventions.ServiceVersionKey):    {},
		string(conventions.ServiceInstanceIDKey): {},
	}
)

func (o *opampAgent) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

// init empty availableComponents to not get an error when starting the opampClient

func (o *opampAgent) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Opamp-go considers this an error, but the collector does not.
// https://github.com/open-telemetry/opamp-go/issues/255

// Dependencies implements extensioncapabilities.Dependent
func (o *opampAgent) Dependencies() []component.ID { _ = "STUB: not implemented"; return nil }

func (o *opampAgent) NotifyConfig(ctx context.Context, conf *confmap.Conf) error {
	_ = "STUB: not implemented"
	return nil
}

func (o *opampAgent) Register(capability string, opts ...opampcustommessages.CustomCapabilityRegisterOption) (opampcustommessages.CustomCapabilityHandler, error) {
	_ = "STUB: not implemented"
	return *new(opampcustommessages.CustomCapabilityHandler), nil
}

func (o *opampAgent) Ready() error { _ = "STUB: not implemented"; return nil }

func (o *opampAgent) NotReady() error { _ = "STUB: not implemented"; return nil }

// ComponentStatusChanged implements the componentstatus.Watcher interface.
func (o *opampAgent) ComponentStatusChanged(
	source *componentstatus.InstanceID,
	event *componentstatus.Event,
) {
	_ = "STUB: not implemented"
	// There can be late arriving events after shutdown. We need to close
	// the event channel so that this function doesn't block and we release all
	// goroutines, but attempting to write to a closed channel will panic; log
	// and recover.
	return
}

func (o *opampAgent) updateEffectiveConfig(conf *confmap.Conf) { _ = "STUB: not implemented"; return }

func newOpampAgent(cfg *Config, set extension.Settings) (*opampAgent, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseInstanceIDString(instanceUID string) (uuid.UUID, error) {
	_ = "STUB: not implemented"
	return *new(uuid.UUID), nil
}

func stringKeyValue(key, value string) *protobufs.KeyValue { _ = "STUB: not implemented"; return nil }

func (o *opampAgent) createAgentDescription() error { _ = "STUB: not implemented"; return nil }

// Initially construct using a map to properly deduplicate any keys that
// are both automatically determined and defined in the config

// skip the attributes that are being used in the identifying attributes.

// Sort the non identifying attributes to give them a stable order for tests

func (o *opampAgent) updateAgentIdentity(instanceID uuid.UUID) { _ = "STUB: not implemented"; return }

func (o *opampAgent) composeEffectiveConfig() *protobufs.EffectiveConfig {
	_ = "STUB: not implemented"
	return nil
}

func (o *opampAgent) onMessage(_ context.Context, msg *types.MessageData) {
	_ = "STUB: not implemented"
	return
}

func (o *opampAgent) onCommand(_ context.Context, command *protobufs.ServerToAgentCommand) error {
	_ = "STUB: not implemented"
	return nil
}

func (o *opampAgent) setHealth(ch *protobufs.ComponentHealth) { _ = "STUB: not implemented"; return }

func getOSDescription(logger *zap.Logger) string { _ = "STUB: not implemented"; return "" }

func (o *opampAgent) initHealthReporting() { _ = "STUB: not implemented"; return }

// Start processing events in the background so that our status watcher doesn't
// block others before the extension starts.

func (o *opampAgent) initAvailableComponents(moduleInfos service.ModuleInfos) {
	_ = "STUB: not implemented"
	return
}

func generateAvailableComponentsHash(moduleInfos service.ModuleInfos) []byte {
	_ = "STUB: not implemented"
	return nil
}

// Compute the SHA-256 hash of the serialized representation.

func addComponentTypeComponentsToStringBuilder(builder *strings.Builder, componentTypeComponents map[component.Type]service.ModuleInfo, componentType string) {
	_ = "STUB: not implemented"
	// Collect components and sort them to ensure deterministic ordering.
	return
}

// Append the component type and its sorted key-value pairs.

func createComponentTypeAvailableComponentDetails(componentTypeComponents map[component.Type]service.ModuleInfo) map[string]*protobufs.ComponentDetails {
	_ = "STUB: not implemented"
	return nil
}

func (o *opampAgent) componentHealthEventLoop() {
	_ = "STUB: not implemented"
	// Record events with component.StatusStarting, but queue other events until
	// PipelineWatcher.Ready is called. This prevents aggregate statuses from
	// flapping between StatusStarting and StatusOK as components are started
	// individually by the service.
	return
}

// After PipelineWatcher.Ready, record statuses as they are received.

func (o *opampAgent) statusAggregatorEventLoop(unsubscribeFunc status.UnsubscribeFunc, statusChan <-chan *status.AggregateStatus) {
	_ = "STUB: not implemented"
	return
}

func convertComponentHealth(statusUpdate *status.AggregateStatus) *protobufs.ComponentHealth {
	_ = "STUB: not implemented"
	return nil
}

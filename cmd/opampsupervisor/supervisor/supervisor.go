// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package supervisor

import (
	"context"
	_ "embed"
	"errors"
	"net/http"
	"sync"
	"sync/atomic"
	"text/template"
	"time"

	"github.com/open-telemetry/opamp-go/client"
	"github.com/open-telemetry/opamp-go/client/types"
	"github.com/open-telemetry/opamp-go/protobufs"
	"github.com/open-telemetry/opamp-go/server"
	serverTypes "github.com/open-telemetry/opamp-go/server/types"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/otel/log"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/cmd/opampsupervisor/supervisor/commander"
	"github.com/open-telemetry/opentelemetry-collector-contrib/cmd/opampsupervisor/supervisor/config"
	supervisorTelemetry "github.com/open-telemetry/opentelemetry-collector-contrib/cmd/opampsupervisor/supervisor/telemetry"
)

var (
	//go:embed templates/nooppipeline.yaml
	noopPipelineTpl string

	//go:embed templates/extratelemetryconfig.yaml
	extraConfigTpl string

	//go:embed templates/opampextension.yaml
	opampextensionTpl string

	//go:embed templates/owntelemetry.yaml
	ownTelemetryTpl string

	lastRecvRemoteConfigFile       = "last_recv_remote_config.dat"
	lastRecvOwnTelemetryConfigFile = "last_recv_own_telemetry_config.dat"

	errNonMatchingInstanceUID = errors.New("received collector instance UID does not match expected UID set by the supervisor")
)

const (
	persistentStateFileName     = "persistent_state.yaml"
	agentConfigFileName         = "effective.yaml"
	AllowNoPipelinesFeatureGate = "service.AllowNoPipelines"
)

const maxBufferedCustomMessages = 10

type configState struct {
	// Supervisor-assembled config to be given to the Collector.
	mergedConfig string
	// true if the server provided configmap was empty
	configMapIsEmpty bool
}

func (c *configState) equal(other *configState) bool { _ = "STUB: not implemented"; return false }

type agentStartStatus string

var (
	agentStarting    agentStartStatus = "starting"
	agentNotStarting agentStartStatus = "notStarting"
)

type telemetrySettings struct {
	component.TelemetrySettings
	loggerProvider log.LoggerProvider
}

// Supervisor implements supervising of OpenTelemetry Collector and uses OpAMPClient
// to work with an OpAMP Server.
type Supervisor struct {
	runCtx       context.Context
	runCtxCancel context.CancelFunc

	pidProvider pidProvider

	// Commander that starts/stops the Agent process.
	commander *commander.Commander

	// Supervisor's own config.
	config config.Supervisor

	agentDescription    *atomic.Value
	availableComponents *atomic.Value

	// Supervisor's persistent state
	persistentState *persistentState

	noopPipelineTemplate         *template.Template
	opampextensionTemplate       *template.Template
	extraTelemetryConfigTemplate *template.Template
	ownTelemetryTemplate         *template.Template

	agentConn *atomic.Value

	// A config section to be added to the Collector's config to fetch its own telemetry.
	// TODO: store this persistently so that when starting we can compose the effective
	// config correctly.
	// https://github.com/open-telemetry/opentelemetry-collector-contrib/issues/21078
	agentConfigOwnTelemetrySection *atomic.Value

	// Internal config state for agent use. See the [configState] struct for more details.
	cfgState *atomic.Value

	// Final effective config of the Collector.
	effectiveConfig *atomic.Value

	// Last received remote config.
	remoteConfig atomic.Pointer[protobufs.AgentRemoteConfig]

	// A channel to indicate there is a new config to apply.
	hasNewConfig chan struct{}
	// configApplyTimeout is the maximum time to wait for the agent to apply a new config.
	// After this time passes without the agent reporting health as OK, the agent is considered unhealthy.
	configApplyTimeout time.Duration
	// lastHealthFromClient is the last health status of the agent received from the client.
	lastHealthFromClient atomic.Pointer[protobufs.ComponentHealth]

	// The OpAMP client to connect to the OpAMP Server.
	opampClient client.OpAMPClient

	doneChan chan struct{}
	agentWG  sync.WaitGroup

	customMessageToServer chan *protobufs.CustomMessage
	customMessageWG       sync.WaitGroup

	// agentReady is true if the agent has started and is fully ready.
	agentReady atomic.Bool
	// agentReadyChan is a channel that can be used to wait for the agent to
	// start in case [agentReady] is false.
	agentReadyChan chan struct{}

	// agentRestarting is true if the agent is restarting.
	agentRestarting atomic.Bool

	// The OpAMP server to communicate with the Collector's OpAMP extension
	opampServer     server.OpAMPServer
	opampServerPort int

	// The HTTP server for health check endpoint
	healthCheckServer   *http.Server
	healthCheckServerWG sync.WaitGroup

	telemetrySettings telemetrySettings

	featureGates map[string]struct{}
	metrics      *supervisorTelemetry.Metrics

	// heartbeatInterval is the interval the OpAMP client is configured to send heartbeats.
	// Default is 30 seconds but can be overridden by the OpAMP server with an OpAMPConnectionSettings message.
	heartbeatIntervalSeconds uint64
}

func NewSupervisor(ctx context.Context, logger *zap.Logger, cfg config.Supervisor) (*Supervisor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Validate extensions feature gate before continuing

func initTelemetrySettings(ctx context.Context, logger *zap.Logger, cfg config.Telemetry) (telemetrySettings, error) {
	_ = "STUB: not implemented"
	return *new(telemetrySettings), nil
}

func (s *Supervisor) Start(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (s *Supervisor) getFeatureGates() error { _ = "STUB: not implemented"; return nil }

// First line only contains headers, discard it.

func (s *Supervisor) createTemplates() error { _ = "STUB: not implemented"; return nil }

// getBootstrapInfo obtains the Collector's agent description by
// starting a Collector with a specific config that only starts
// an OpAMP extension, obtains the agent description, then
// shuts down the Collector. This only needs to happen
// once per Collector binary.
func (s *Supervisor) getBootstrapInfo() (err error) { _ = "STUB: not implemented"; return nil }

// Start a one-shot server to get the Collector's agent description
// and available components using the Collector's OpAMP extension.

// agent description must be defined

// if available components have not been reported, agent description is sufficient to continue

// must have a full list of components if available components have been reported

// if we don't have a full component list, ask for it

// need to only report done once, not on each message - otherwise, we get a hung thread

// try to report the issue to the OpAMP server

func (s *Supervisor) startOpAMP() error { _ = "STUB: not implemented"; return nil }

func (s *Supervisor) startOpAMPClient() error {
	_ = "STUB: not implemented"
	// determine if we need to load a TLS config or not
	return nil
}

//nolint:errcheck

// TODO: https://github.com/open-telemetry/opentelemetry-collector-contrib/issues/21079

// Set heartbeat interval if the agent supports it

func (s *Supervisor) startHealthCheckServer() error { _ = "STUB: not implemented"; return nil }

type nopHost struct{}

var _ component.Host = nopHost{}

func (nopHost) GetExtensions() map[component.ID]component.Component {
	_ = "STUB: not implemented"

	// startOpAMPServer starts an OpAMP server that will communicate
	// with an OpAMP extension running inside a Collector to receive
	// data from inside the Collector. The internal server's lifetime is not
	// matched to the Collector's process, but may be restarted
	// depending on information received by the Supervisor from the remote
	// OpAMP server.
	return nil
}

func (s *Supervisor) startOpAMPServer() error { _ = "STUB: not implemented"; return nil }

// Only allow one agent to be connected the this server at a time.

func (s *Supervisor) handleAgentOpAMPMessage(conn serverTypes.Connection, message *protobufs.AgentToServer) *protobufs.ServerToAgent {
	_ = "STUB: not implemented"
	return nil
}

// Proxy client capabilities to server

// Proxy agent custom messages to server

func (s *Supervisor) forwardCustomMessagesToServerLoop() { _ = "STUB: not implemented"; return }

// OK

// setAgentDescription sets the agent description, merging in any user-specified attributes from the supervisor configuration.
func (s *Supervisor) setAgentDescription(ad *protobufs.AgentDescription) {
	_ = "STUB: not implemented"
	return
}

// setAvailableComponents sets the available components of the OpAMP agent
func (s *Supervisor) setAvailableComponents(ac *protobufs.AvailableComponents) {
	_ = "STUB: not implemented"
	return
}

// applyKeyValueOverrides merges the overrides map into the array of key value pairs.
// If a key from overrides already exists in the array of key value pairs, it is overwritten by the value from the overrides map.
// An array of KeyValue pair is returned, with each key value pair having a distinct key.
func applyKeyValueOverrides(overrides map[string]string, orig []*protobufs.KeyValue) []*protobufs.KeyValue {
	_ = "STUB: not implemented"
	return nil
}

// Sort keys for stable output, makes it easier to test.

func (s *Supervisor) stopOpAMPClient() error { _ = "STUB: not implemented"; return nil }

// TODO(srikanthccv): remove context.DeadlineExceeded after https://github.com/open-telemetry/opamp-go/pull/213

func (*Supervisor) getHeadersFromSettings(protoHeaders *protobufs.Headers) http.Header {
	_ = "STUB: not implemented"
	return *new(http.Header)
}

func (s *Supervisor) onOpampConnectionSettings(_ context.Context, settings *protobufs.OpAMPConnectionSettings) error {
	_ = "STUB: not implemented"
	return nil
}

// Update the heartbeat interval if the agent supports it

// take a copy of the current OpAMP server config

// update the OpAMP server config

// revert the OpAMP server config

// start the OpAMP client with the old settings

func (s *Supervisor) addSpecialConfigFiles() { _ = "STUB: not implemented"; return }

// if missing builtin, add it to the beginning

// if missing opamp extension, add it to the end

// if missing remote config, add it to the end

func (s *Supervisor) composeNoopPipeline() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Supervisor) createRemoteConfigComposers(incomingConfig *protobufs.AgentRemoteConfig) []configComposer {
	_ = "STUB: not implemented"
	return nil
}

// Sort to make sure the order of merging is stable.

// skip instance config

// Append instance config as the last item.

// Merge received configs.

func (s *Supervisor) composeNoopConfig() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Supervisor) composeOwnTelemetryConfig() []byte { _ = "STUB: not implemented"; return nil }

func (s *Supervisor) composeExtraTelemetryConfig() []byte { _ = "STUB: not implemented"; return nil }

func (s *Supervisor) composeOpAMPExtensionConfig() []byte { _ = "STUB: not implemented"; return nil }

type configComposer func() []byte

func (s *Supervisor) composeAgentConfigFiles(incomingConfig *protobufs.AgentRemoteConfig) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// The special config files should always be valid yaml and most of them
// will be due to the typing in the Supervisor's config struct and its
// validation function.
// The special config file for the remote configuration is the exception
// here: it could be invalid yaml. In case it is, the `koanf` loader
// will return an error.
// Normal config files with invalid yaml should be just ignored.

// loadAndWriteInitialMergedConfig loads and writes the initial config by
// merging the last received remote config, last received own metrics config,
// and any local configs.
func (s *Supervisor) loadAndWriteInitialMergedConfig() error {
	_ = "STUB: not implemented"
	// load the last received remote config
	return nil
}

// load the last received own telemetry config

// compose the initial merged config

// write the initial merged config to disk

// loadRemoteConfig loads the last received remote config from file if the capability is supported.
func (s *Supervisor) loadRemoteConfig() { _ = "STUB: not implemented"; return }

// Try to load the last received remote config if it exists.

// loadLastReceivedOwnTelemetryConfig loads the last received own telemetry config from file if the capability is supported.
func (s *Supervisor) loadLastReceivedOwnTelemetryConfig() {
	_ = "STUB: not implemented"
	// If none of the own telemetry capabilities are supported, do nothing.
	return
}

// Try to load the last received own metrics config if it exists.

// createEffectiveConfigMsg create an EffectiveConfig with the content of the
// current effective config.
func (s *Supervisor) createEffectiveConfigMsg() *protobufs.EffectiveConfig {
	_ = "STUB: not implemented"
	return nil
}

func (*Supervisor) updateOwnTelemetryData(data map[string]any, signal string, settings *protobufs.TelemetryConnectionSettings) map[string]any {
	_ = "STUB: not implemented"
	return nil
}

func (s *Supervisor) setupOwnTelemetry(_ context.Context, settings *protobufs.ConnectionSettingsOffers) (configChanged bool) {
	_ = "STUB: not implemented"
	return false
}

// Need to recalculate the Agent config so that the metric config is included in it.

// composeMergedConfig composes the merged config from multiple sources:
// 1) the remote config from OpAMP Server
// 2) the own metrics config section
// 3) the local override config that is hard-coded in the Supervisor.
func (s *Supervisor) composeMergedConfig(incomingConfig *protobufs.AgentRemoteConfig) (configChanged bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Add noop pipeline

// The merged final result is our new merged config.

// Check if supervisor's merged config is changed.

// Validate BEFORE storing to prevent race condition where other goroutines read invalid config

// Only store after successful validation (or if validation is disabled/skipped)

func (s *Supervisor) handleRestartCommand() error { _ = "STUB: not implemented"; return nil }

func (s *Supervisor) startAgent() (agentStartStatus, error) {
	_ = "STUB: not implemented"
	return *new(agentStartStatus), nil
}

// Don't start the agent if there is no config to run

// need to manually trigger updating effective config

func (s *Supervisor) runAgentProcess() { _ = "STUB: not implemented"; return }

// We have an effective config file saved previously. Use it to start the agent.

// Try to drain the channel

// This call to [startAgent] is useful for both the normal config reload
// and the HUP one. It takes care of not starting the agent if the config
// is empty and also of starting it when the config changes from empty
// to non-empty.

// not starting agent because of nop config: clear timer, report applied status, report healthy status

// need to clear exit channel to avoid triggering `s.commander.Exited()` case
// because we stopped the agent and aren't restarting it this case will trigger
// and report an unhealthy status (collector would be healthy, just choosing to not run)

// the agent process exit is expected for restart command and will not attempt to restart

// If agent crashed while we were waiting for config to be applied (timeout timer is running),
// report the config as FAILED immediately rather than waiting for the timeout.
// Stop the timer and drain the channel to prevent it from firing later.

// Timer already fired or was stopped, drain the channel

// Timer had already fired but we handled the crash first

// Timer was already stopped

// Timer was running, which means we were waiting for config to be applied.
// Report FAILED status immediately.

// Wait 5 seconds before starting again.

// Try to drain the channel

// markAgentReady marks the agent as ready and sends a signal to
// [agentReadyChan].
func (s *Supervisor) markAgentReady() { _ = "STUB: not implemented"; return }

// resetAgentReady resets the agent as not ready and drains [agentReadyChan].
func (s *Supervisor) resetAgentReady() { _ = "STUB: not implemented"; return }

// waitForAgentReady waits for the agent to be ready. The agent is considered to
// be ready when its first health report is received by the Supervisor's opamp
// server.
// WARNING: this is not thread-safe! If there are two goroutines waiting for
// the agent to be ready, only one of them will be able to proceed.
func (s *Supervisor) waitForAgentReady() error { _ = "STUB: not implemented"; return nil }

// hupReloadAgent sends a HUP signal to the agent process  with the intent of
// triggering a configuration reload. There are 3 possible outcomes of this:
// 1. The agent is the "official" Otel Collector, which properly traps the HUP
// signal and reloads the config.
// 2. The agent is a custom process, which does not trap the HUP signal: in
// this case the default behavior of a process that receives a HUP signal is to
// exit. This ends up being the same behavior as the standard stop -> restart
// configuration reload method.
// 3. The agent traps the HUP signal, but does nothing. On the next health
// report the agent will be consiered healthy and ready, even though it might
// be running on old configuration.
func (s *Supervisor) hupReloadAgent() error { _ = "STUB: not implemented"; return nil }

// If we have an empty config, the agent should be stopped.

// If the agent is not running, we can't send a HUP signal to it, so we
// return and let it be started by the caller.

func (s *Supervisor) reloadAgentConfig() error { _ = "STUB: not implemented"; return nil }

func (s *Supervisor) writeAgentConfig() error { _ = "STUB: not implemented"; return nil }

// validateConfig validates a configuration string without storing it in cfgState.
// This prevents race conditions where other goroutines might read invalid config during validation.
// Returns an error if validation fails.
func (s *Supervisor) validateConfig(configContent string) error {
	_ = "STUB: not implemented"
	return nil
}

// Write config to a temporary file for validation

func (s *Supervisor) stopAgentApplyConfig() { _ = "STUB: not implemented"; return }

func (s *Supervisor) Shutdown() { _ = "STUB: not implemented"; return }

// Shutdown in order from producer to consumer (agent -> customMessageForwarder -> local OpAMP server -> client to remote OpAMP server).

func (s *Supervisor) shutdownTelemetry() error { _ = "STUB: not implemented"; return nil }

// The metric.MeterProvider and trace.TracerProvider interfaces do not have a Shutdown method.
// To shutdown the providers we try to cast to this interface, which matches the type signature used in the SDK.

func (s *Supervisor) saveLastReceivedConfig(config *protobufs.AgentRemoteConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Supervisor) saveLastReceivedOwnTelemetrySettings(set *protobufs.ConnectionSettingsOffers, filePath string) error {
	_ = "STUB: not implemented"
	return nil
}

// saveAndReportConfigStatus saves the config status to the persistent state and reports it to the server.
func (s *Supervisor) saveAndReportConfigStatus(status protobufs.RemoteConfigStatuses, errorMessage string) {
	_ = "STUB: not implemented"
	return
}

// save status to persistent state

// report status to server

func (s *Supervisor) SetHealth(componentHealth *protobufs.ComponentHealth) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Supervisor) onMessage(ctx context.Context, msg *types.MessageData) {
	_ = "STUB: not implemented"
	return
}

// Update the agent config if any messages have touched the config

// Signal that there is a new config.

// Proxy server capabilities to opamp extension

// Proxy server messages to opamp extension

// Send any messages that need proxying to the agent.

// processRemoteConfigMessage processes an AgentRemoteConfig message, returning true if the agent config has changed.
func (s *Supervisor) processRemoteConfigMessage(ctx context.Context, msg *protobufs.AgentRemoteConfig) bool {
	_ = "STUB: not implemented"
	return false
}

// Clone the message to avoid race conditions when the protobuf is being unmarshaled
// while another goroutine reads from it

// only report applying if the config has changed and will run agent with new config

// if the config has not changed report applied status, we should still report a status to the server in this case

// processOwnTelemetryConnSettingsMessage processes a TelemetryConnectionSettings message, returning true if the agent config has changed.
func (s *Supervisor) processOwnTelemetryConnSettingsMessage(ctx context.Context, msg *protobufs.ConnectionSettingsOffers) bool {
	_ = "STUB: not implemented"
	return false
}

// processAgentIdentificationMessage processes an AgentIdentification message, returning true if the agent config has changed.
func (s *Supervisor) processAgentIdentificationMessage(msg *protobufs.AgentIdentification) bool {
	_ = "STUB: not implemented"
	return false
}

// Need to recalculate the Agent config so that the new agent identification is included in it.

func (s *Supervisor) persistentStateFilePath() string { _ = "STUB: not implemented"; return "" }

func (s *Supervisor) agentConfigFilePath() string { _ = "STUB: not implemented"; return "" }

func (s *Supervisor) getSupervisorOpAMPServerPort() (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *Supervisor) getFeatureGateFlag() []string { _ = "STUB: not implemented"; return nil }

func (s *Supervisor) isFeatureGateSupported(gate string) bool {
	_ = "STUB: not implemented"
	return false
}

func (*Supervisor) findRandomPort() (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (s *Supervisor) getTracer() trace.Tracer { _ = "STUB: not implemented"; return *new(trace.Tracer) }

// The default koanf behavior is to override lists in the config.
// Instead, we provide this function, which merges the source and destination config's
// extension lists by concatenating the two.
// Will be resolved by https://github.com/open-telemetry/opentelemetry-collector/issues/8754
func configMergeFunc(src, dest map[string]any) error { _ = "STUB: not implemented"; return nil }

// This is a small hack to ensure that the order is consitent and
// follows this simple rule: extensions from [src], then from [dest],
// in the order that they appear.
// We cannot use other simpler methods, like [sort.Strings], because
// we work with a `[]any` that cannot be cast to `[]string`.

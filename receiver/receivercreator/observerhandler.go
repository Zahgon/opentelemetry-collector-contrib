// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package receivercreator // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/receivercreator"

import (
	"sync"

	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/consumer/xconsumer"
	"go.opentelemetry.io/collector/receiver"

	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer"
)

var _ observer.Notify = (*observerHandler)(nil)

const (
	// tmpSetEndpointConfigKey denotes the observerHandler (not the user) has set an "endpoint" target field
	// in resolved configuration. Used to determine if the field should be removed when the created receiver
	// doesn't expose such a field.
	tmpSetEndpointConfigKey = "<tmp.receiver.creator.automatically.set.endpoint.field>"
)

// observerHandler manages endpoint change notifications.
type observerHandler struct {
	sync.Mutex
	config *Config
	params receiver.Settings
	// receiversByEndpointID is a map of endpoint IDs to a receiver instance.
	receiversByEndpointID receiverMap
	// nextLogsConsumer is the receiver_creator's own consumer
	nextLogsConsumer consumer.Logs
	// nextMetricsConsumer is the receiver_creator's own consumer
	nextMetricsConsumer consumer.Metrics
	// nextTracesConsumer is the receiver_creator's own consumer
	nextTracesConsumer consumer.Traces
	// nextProfilesConsumer is the receiver_creator's own consumer
	nextProfilesConsumer xconsumer.Profiles
	// runner starts and stops receiver instances.
	runner runner
}

// shutdown all receivers started at runtime.
func (obs *observerHandler) shutdown() error { _ = "STUB: not implemented"; return nil }

// TODO: Should keep track of which receiver the error is associated with
// but require some restructuring.

func (obs *observerHandler) ID() observer.NotifyID {
	_ = "STUB: not implemented"
	return *new(observer.NotifyID)
}

// OnAdd responds to endpoint add notifications.
func (obs *observerHandler) OnAdd(added []observer.Endpoint) { _ = "STUB: not implemented"; return }

// OnRemove responds to endpoint removal notifications.
func (obs *observerHandler) OnRemove(removed []observer.Endpoint) {
	_ = "STUB: not implemented"
	return
}

// debug log the endpoint to improve usability

// OnChange responds to endpoint change notifications.
func (obs *observerHandler) OnChange(changed []observer.Endpoint) {
	_ = "STUB: not implemented"
	// TODO: optimize to only restart if effective config has changed.
	return
}

func (obs *observerHandler) startReceiver(template receiverTemplate, env observer.EndpointEnv, e observer.Endpoint) {
	_ = "STUB: not implemented"
	return
}

// If user didn't set endpoint set to default value as well as
// flag indicating we've done this for later validation.

// Though not necessary with contrib provided observers, nothing is stopping custom
// ones from using expr in their Target values.

// Adds default and/or configured resource attributes (e.g. k8s.pod.uid) to resources
// as telemetry is emitted.

// short-circuit if no consumers are set

// TODO: we should consider about applying sensitive data redaction

func filterConsumerSignals(consumer *enhancingConsumer, signals receiverSignals) {
	_ = "STUB: not implemented"
	return
}

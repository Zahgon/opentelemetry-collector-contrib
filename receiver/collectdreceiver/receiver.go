// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package collectdreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/collectdreceiver"

import (
	"context"
	"net/http"
	"sync"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/receiverhelper"
	"go.uber.org/zap"
)

var _ receiver.Metrics = (*collectdReceiver)(nil)

// collectdReceiver implements the receiver.Metrics for CollectD protocol.
type collectdReceiver struct {
	logger             *zap.Logger
	server             *http.Server
	shutdownWG         sync.WaitGroup
	defaultAttrsPrefix string
	nextConsumer       consumer.Metrics
	obsrecv            *receiverhelper.ObsReport
	createSettings     receiver.Settings
	config             *Config
}

// newCollectdReceiver creates the CollectD receiver with the given parameters.
func newCollectdReceiver(
	logger *zap.Logger,
	cfg *Config,
	defaultAttrsPrefix string,
	nextConsumer consumer.Metrics,
	createSettings receiver.Settings,
) (receiver.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(receiver.Metrics), nil
}

// Start starts an HTTP server that can process CollectD JSON requests.
func (cdr *collectdReceiver) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

// Shutdown stops the CollectD receiver.
func (cdr *collectdReceiver) Shutdown(context.Context) error { _ = "STUB: not implemented"; return nil }

// ServeHTTP acts as the default and only HTTP handler for the CollectD receiver.
func (cdr *collectdReceiver) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (cdr *collectdReceiver) defaultAttributes(req *http.Request) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func (cdr *collectdReceiver) handleHTTPErr(w http.ResponseWriter, err error, msg string) {
	_ = "STUB: not implemented"
	return
}

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package adapter // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/adapter"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/extension/xextension/storage"
	rcvr "go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/receiverhelper"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/entry"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/pipeline"
)

type receiver struct {
	set component.TelemetrySettings
	id  component.ID

	pipe     pipeline.Pipeline
	emitter  helper.LogEmitter
	consumer consumer.Logs
	obsrecv  *receiverhelper.ObsReport

	storageID     *component.ID
	storageClient storage.Client
}

// Ensure this receiver adheres to required interface
var _ rcvr.Logs = (*receiver)(nil)

// Start tells the receiver to start
func (r *receiver) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *receiver) consumeEntries(ctx context.Context, entries []*entry.Entry) {
	_ = "STUB: not implemented"
	return
}

// Shutdown is invoked during service shutdown
func (r *receiver) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

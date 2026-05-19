// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package telemetrytest // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/xray/telemetry/telemetrytest"

import (
	"context"
	"sync/atomic"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/xray/telemetry"
)

// NewSenderSink returns a Sender that acts like a sink and
// stores all calls to the record functions for testing.
func NewSenderSink() *SenderSink { _ = "STUB: not implemented"; return nil }

var _ telemetry.Sender = (*SenderSink)(nil)

type SenderSink struct {
	telemetry.Recorder
	StartCount *atomic.Int64
	StopCount  *atomic.Int64
}

func (s SenderSink) Start(_ context.Context) { _ = "STUB: not implemented"; return }

func (s SenderSink) Stop() { _ = "STUB: not implemented"; return }

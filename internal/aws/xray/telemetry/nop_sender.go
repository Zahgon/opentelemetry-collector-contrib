// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package telemetry // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/xray/telemetry"

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/xray/types"
)

// NewNopSender returns a Sender that drops all data.
func NewNopSender() Sender { _ = "STUB: not implemented"; return *new(Sender) }

var nopSenderInstance Sender = &nopSender{}

type nopSender struct{}

func (nopSender) Rotate() types.TelemetryRecord {
	_ = "STUB: not implemented"
	return *new(types.TelemetryRecord)
}

func (nopSender) HasRecording() bool { _ = "STUB: not implemented"; return false }

func (nopSender) Start(context.Context) { _ = "STUB: not implemented"; return }

func (nopSender) Stop() { _ = "STUB: not implemented"; return }

func (nopSender) RecordSegmentsReceived(int) { _ = "STUB: not implemented"; return }

func (nopSender) RecordSegmentsSent(int) { _ = "STUB: not implemented"; return }

func (nopSender) RecordSegmentsSpillover(int) { _ = "STUB: not implemented"; return }

func (nopSender) RecordSegmentsRejected(int) { _ = "STUB: not implemented"; return }

func (nopSender) RecordConnectionError(error) { _ = "STUB: not implemented"; return }

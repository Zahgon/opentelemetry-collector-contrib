// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package awss3receiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awss3receiver"

import (
	"context"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/opampcustommessages"
)

const (
	IngestStatusCompleted   = "completed"
	IngestStatusFailed      = "failed"
	IngestStatusIngesting   = "ingesting"
	CustomCapability        = "org.opentelemetry.collector.receiver.awss3"
	maxNotificationAttempts = 3
)

type statusNotification struct {
	TelemetryType  string
	IngestStatus   string
	StartTime      time.Time
	EndTime        time.Time
	IngestTime     time.Time
	FailureMessage string
}

type statusNotifier interface {
	Start(ctx context.Context, host component.Host) error
	Shutdown(ctx context.Context) error
	SendStatus(ctx context.Context, message statusNotification)
}

type opampNotifier struct {
	logger           *zap.Logger
	opampExtensionID component.ID
	handler          opampcustommessages.CustomCapabilityHandler
}

func newNotifier(config *Config, logger *zap.Logger) statusNotifier {
	_ = "STUB: not implemented"
	return *new(statusNotifier)
}

func (n *opampNotifier) Start(_ context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *opampNotifier) Shutdown(_ context.Context) error { _ = "STUB: not implemented"; return nil }

func (n *opampNotifier) SendStatus(_ context.Context, message statusNotification) {
	_ = "STUB: not implemented"
	return
}

// The only other errors returned by the OpAmp extension are unrecoverable, ie ErrCustomCapabilityNotSupported
// so just log an error and return.

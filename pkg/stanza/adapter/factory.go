// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package adapter // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/adapter"

import (
	"go.opentelemetry.io/collector/component"
	rcvr "go.opentelemetry.io/collector/receiver"
	xrcvr "go.opentelemetry.io/collector/receiver/xreceiver"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
)

// LogReceiverType is the interface used by stanza-based log receivers
type LogReceiverType interface {
	Type() component.Type
	CreateDefaultConfig() component.Config
	BaseConfig(component.Config) BaseConfig
	InputConfig(component.Config) operator.Config
}

// NewFactory creates a factory for a Stanza-based receiver
func NewFactory(logReceiverType LogReceiverType, sl component.StabilityLevel, opts ...xrcvr.FactoryOption) rcvr.Factory {
	_ = "STUB: not implemented"
	return *new(rcvr.Factory)
}

func createLogsReceiver(logReceiverType LogReceiverType) rcvr.CreateLogsFunc {
	_ = "STUB: not implemented"
	return *new(rcvr.CreateLogsFunc)
}

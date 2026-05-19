// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package telemetrytest // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/xray/telemetry/telemetrytest"

import (
	"go.opentelemetry.io/collector/component"

	awsxray "github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/xray"
	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/xray/telemetry"
)

func NewNopRegistry() telemetry.Registry {
	_ = "STUB: not implemented"
	return *new(telemetry.Registry)
}

type nopRegistry struct {
	sender telemetry.Sender
}

var _ telemetry.Registry = (*nopRegistry)(nil)

var nopRegistryInstance = &nopRegistry{
	sender: telemetry.NewNopSender(),
}

func (n nopRegistry) Register(component.ID, telemetry.Config, awsxray.XRayClient, ...telemetry.Option) telemetry.Sender {
	_ = "STUB: not implemented"
	return *new(telemetry.Sender)
}

func (n nopRegistry) Load(component.ID) telemetry.Sender {
	_ = "STUB: not implemented"
	return *new(telemetry.Sender)
}

func (n nopRegistry) LoadOrNop(component.ID) telemetry.Sender {
	_ = "STUB: not implemented"
	return *new(telemetry.Sender)
}

func (n nopRegistry) LoadOrStore(component.ID, telemetry.Sender) (telemetry.Sender, bool) {
	_ = "STUB: not implemented"
	return *new(telemetry.Sender), false
}

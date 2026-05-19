// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package awsproxy // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/awsproxy"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/proxy"
)

type xrayProxy struct {
	logger   *zap.Logger
	config   *Config
	server   proxy.Server
	settings component.TelemetrySettings
}

var _ extension.Extension = (*xrayProxy)(nil)

func (x *xrayProxy) Start(_ context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (x *xrayProxy) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func newXrayProxy(config *Config, telemetrySettings component.TelemetrySettings) (extension.Extension, error) {
	_ = "STUB: not implemented"
	return *new(extension.Extension), nil
}

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package coralogixexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/coralogixexporter"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/pdata/pprofile"
	"go.opentelemetry.io/collector/pdata/pprofile/pprofileotlp"
)

func newProfilesExporter(cfg component.Config, set exporter.Settings) (*profilesExporter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type profilesExporter struct {
	profilesExporter pprofileotlp.GRPCClient
	*signalExporter
}

func (e *profilesExporter) start(ctx context.Context, host component.Host) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (e *profilesExporter) pushProfiles(ctx context.Context, md pprofile.Profiles) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *profilesExporter) enhanceContext(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

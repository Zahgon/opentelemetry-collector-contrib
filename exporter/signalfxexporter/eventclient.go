// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package signalfxexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/signalfxexporter"

import (
	"context"
	"io"

	sfxpb "github.com/signalfx/com_signalfx_metrics_protobuf/model"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.uber.org/zap"
)

// sfxEventClient sends the data to the SignalFx backend.
type sfxEventClient struct {
	sfxClientBase
	logger                 *zap.Logger
	accessTokenPassthrough bool
}

func (s *sfxEventClient) pushEvents(ctx context.Context, rl plog.ResourceLogs, sfxEvents []*sfxpb.Event) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *sfxEventClient) encodeBody(events []*sfxpb.Event) (bodyReader io.Reader, compressed bool, err error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), false, nil
}

func (s *sfxEventClient) retrieveAccessToken(ctx context.Context, rl plog.ResourceLogs) string {
	_ = "STUB: not implemented"
	return ""
}

// Nothing to do if token is pass through not configured or resource is nil.

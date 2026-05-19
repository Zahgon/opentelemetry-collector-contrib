// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package cloudfoundryreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/cloudfoundryreceiver"

import (
	"context"
	"net/http"

	"code.cloudfoundry.org/go-loggregator"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.uber.org/zap"
)

type envelopeStreamFactory struct {
	rlpGatewayClient *loggregator.RLPGatewayClient
}

func newEnvelopeStreamFactory(
	ctx context.Context,
	settings component.TelemetrySettings,
	authTokenProvider *uaaTokenProvider,
	httpConfig confighttp.ClientConfig,
	host component.Host,
) (*envelopeStreamFactory, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (rgc *envelopeStreamFactory) CreateMetricsStream(ctx context.Context, baseShardID string) loggregator.EnvelopeStream {
	_ = "STUB: not implemented"
	return *new(loggregator.EnvelopeStream)
}

func (rgc *envelopeStreamFactory) CreateLogsStream(ctx context.Context, baseShardID string) loggregator.EnvelopeStream {
	_ = "STUB: not implemented"
	return *new(loggregator.EnvelopeStream)
}

type authorizationProvider struct {
	logger            *zap.Logger
	authTokenProvider *uaaTokenProvider
	client            *http.Client
}

func (ap *authorizationProvider) Do(request *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

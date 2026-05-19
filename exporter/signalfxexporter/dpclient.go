// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package signalfxexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/signalfxexporter"

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"sync"

	sfxpb "github.com/signalfx/com_signalfx_metrics_protobuf/model"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/signalfxexporter/internal/translation"
)

const (
	contentEncodingHeader   = "Content-Encoding"
	contentTypeHeader       = "Content-Type"
	otlpProtobufContentType = "application/x-protobuf;format=otlp"
)

type sfxClientBase struct {
	ingestURL *url.URL
	headers   map[string]string
	client    *http.Client
	zippers   sync.Pool
}

var metricsMarshaler = &pmetric.JSONMarshaler{}

// avoid attempting to compress things that fit into a single ethernet frame
func (s *sfxClientBase) getReader(b []byte) (io.Reader, bool, error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), false, nil
}

// sfxDPClient sends the data to the SignalFx backend.
type sfxDPClient struct {
	sfxClientBase
	logDataPoints          bool
	logger                 *zap.Logger
	accessTokenPassthrough bool
	converter              *translation.MetricsConverter
	sendOTLPHistograms     bool
}

func (s *sfxDPClient) pushMetricsData(
	ctx context.Context,
	md pmetric.Metrics,
) (droppedDataPoints int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// All metrics in the pmetric.Metrics will have the same access token because of the BatchPerResourceMetrics.

// export SFx format

// export any histograms in otlp if sendOTLPHistograms is true

func (s *sfxDPClient) postData(ctx context.Context, body io.Reader, headers map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

// Set the headers configured in sfxDPClient

// Set any extra headers passed by the caller

// TODO: Mark errors as partial errors wherever applicable when, partial
// error for metrics is available.

func (s *sfxDPClient) pushMetricsDataForToken(ctx context.Context, sfxDataPoints []*sfxpb.DataPoint, accessToken string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Override access token in headers map if it's non empty.

func (s *sfxDPClient) encodeBody(dps []*sfxpb.DataPoint) (bodyReader io.Reader, compressed bool, err error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), false, nil
}

func (s *sfxDPClient) retrieveAccessToken(ctx context.Context, md pmetric.ResourceMetrics) string {
	_ = "STUB: not implemented"
	return ""
}

// Nothing to do if token is pass through not configured or resource is nil.

func (s *sfxDPClient) pushOTLPMetricsDataForToken(ctx context.Context, mh pmetric.Metrics, accessToken string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Set otlp content-type header

// Override access token in headers map if it's non-empty.

func (s *sfxDPClient) encodeOTLPBody(md pmetric.Metrics) (bodyReader io.Reader, compressed bool, err error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), false, nil
}

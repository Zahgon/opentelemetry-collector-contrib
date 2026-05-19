// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package internal // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/pprofreceiver/internal"

import (
	"context"
	"net/http"
	"sync"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/consumer/xconsumer"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/receiverhelper"
)

// PushPath is the path the pprof push server listens on.
const PushPath = "/v1/pprof"

// HTTPServer accepts pprof data pushed over HTTP and forwards it to the next consumer.
type HTTPServer struct {
	ServerConfig confighttp.ServerConfig
	Consumer     xconsumer.Profiles
	Settings     receiver.Settings

	server     *http.Server
	shutdownWG sync.WaitGroup
	obsrep     *receiverhelper.ObsReport
}

var _ component.Component = (*HTTPServer)(nil)

func (s *HTTPServer) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *HTTPServer) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (s *HTTPServer) handlePush(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// confighttp.ToServer transparently handles supported Content-Encodings (e.g. gzip)
// before we read the body here. Cap uncompressed bodies as well, since the
// confighttp decompressor only enforces MaxRequestBodySize when a body is
// actually decoded.

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package als // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/envoyalsreceiver/internal/als"

import (
	alsv3 "github.com/envoyproxy/go-control-plane/envoy/service/accesslog/v3"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/receiver/receiverhelper"
)

const (
	apiVersionAttr = "api_version"
	apiVersionVal  = "v3"
	nodeAttr       = "node"
	logNameAttr    = "log_name"
	logTypeAttr    = "log_type"
	httpTypeVal    = "http"
	tcpTypeVal     = "tcp"
)

type Server struct {
	nextConsumer consumer.Logs
	obsrep       *receiverhelper.ObsReport
}

func New(nextConsumer consumer.Logs, obsrep *receiverhelper.ObsReport) *Server {
	_ = "STUB: not implemented"
	return nil
}

func (s *Server) StreamAccessLogs(logStream alsv3.AccessLogService_StreamAccessLogsServer) error {
	_ = "STUB: not implemented"
	return nil
}

func toLogs(data *alsv3.StreamAccessLogsMessage) plog.Logs {
	_ = "STUB: not implemented"
	return *new(plog.Logs)
}

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package datasenders // import "github.com/open-telemetry/opentelemetry-collector-contrib/testbed/datasenders"

import (
	"context"
	"sync"
	"time"

	jaegerproto "github.com/jaegertracing/jaeger-idl/proto-gen/api_v2"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configgrpc"
	"go.opentelemetry.io/collector/config/configoptional"
	"go.opentelemetry.io/collector/config/configretry"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/metadata"

	"github.com/open-telemetry/opentelemetry-collector-contrib/testbed/testbed"
)

// jaegerGRPCDataSender implements TraceDataSender for Jaeger thrift_http exporter.
type jaegerGRPCDataSender struct {
	testbed.DataSenderBase
	consumer.Traces
}

// Ensure jaegerGRPCDataSender implements TraceDataSender.
var _ testbed.TraceDataSender = (*jaegerGRPCDataSender)(nil)

// NewJaegerGRPCDataSender creates a new Jaeger exporter sender that will send
// to the specified port after Start is called.
func NewJaegerGRPCDataSender(host string, port int) testbed.TraceDataSender {
	_ = "STUB: not implemented"
	return *new(testbed.TraceDataSender)
}

func (je *jaegerGRPCDataSender) Start() error { _ = "STUB: not implemented"; return nil }

func (je *jaegerGRPCDataSender) GenConfigYAMLStr() string { _ = "STUB: not implemented"; return "" }

func (*jaegerGRPCDataSender) ProtocolName() string {
	_ = "STUB: not implemented"

	// Config defines configuration for Jaeger gRPC exporter.
	return ""
}

type jaegerConfig struct {
	TimeoutSettings           exporterhelper.TimeoutConfig                             `mapstructure:",squash"` // squash ensures fields are correctly decoded in embedded struct.
	QueueSettings             configoptional.Optional[exporterhelper.QueueBatchConfig] `mapstructure:"sending_queue"`
	configretry.BackOffConfig `mapstructure:"retry_on_failure"`

	configgrpc.ClientConfig `mapstructure:",squash"` // squash ensures fields are correctly decoded in embedded struct.
}

var _ component.Config = (*jaegerConfig)(nil)

// Validate checks if the exporter configuration is valid
func (cfg *jaegerConfig) Validate() error { _ = "STUB: not implemented"; return nil }

// newTracesExporter returns a new Jaeger gRPC exporter.
// The exporter name is the name to be used in the observability of the exporter.
// The collectorEndpoint should be of the form "hostname:14250" (a gRPC target).
func (je *jaegerGRPCDataSender) newTracesExporter(set exporter.Settings) (exporter.Traces, error) {
	_ = "STUB: not implemented"
	return *new(exporter.Traces), nil
}

// protoGRPCSender forwards spans encoded in the jaeger proto
// format, to a grpc server.
type protoGRPCSender struct {
	name         string
	settings     component.TelemetrySettings
	client       jaegerproto.CollectorServiceClient
	metadata     metadata.MD
	waitForReady bool

	conn                      stateReporter
	connStateReporterInterval time.Duration

	stopCh         chan struct{}
	stopped        bool
	stopLock       sync.Mutex
	clientSettings *configgrpc.ClientConfig
}

type stateReporter interface {
	GetState() connectivity.State
}

func (s *protoGRPCSender) pushTraces(
	ctx context.Context,
	td ptrace.Traces,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *protoGRPCSender) shutdown(context.Context) error { _ = "STUB: not implemented"; return nil }

func (s *protoGRPCSender) start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package googlecloudpubsubreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudpubsubreceiver"

import (
	"context"
	"sync"

	"cloud.google.com/go/pubsub/v2/apiv1/pubsubpb"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/receiverhelper"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudpubsubreceiver/internal"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudpubsubreceiver/internal/metadata"
)

// https://cloud.google.com/pubsub/docs/reference/rpc/google.pubsub.v1#streamingpullrequest
type pubsubReceiver struct {
	settings           receiver.Settings
	obsrecv            *receiverhelper.ObsReport
	tracesConsumer     consumer.Traces
	metricsConsumer    consumer.Metrics
	logsConsumer       consumer.Logs
	userAgent          string
	config             *Config
	client             internal.SubscriberClient
	tracesUnmarshaler  ptrace.Unmarshaler
	metricsUnmarshaler pmetric.Unmarshaler
	logsUnmarshaler    plog.Unmarshaler
	handler            *internal.StreamHandler
	startOnce          sync.Once
	telemetryBuilder   *metadata.TelemetryBuilder
}

type buildInEncoding int

const (
	unknown         buildInEncoding = iota
	otlpProtoTrace                  = iota
	otlpProtoMetric                 = iota
	otlpProtoLog                    = iota
	rawTextLog                      = iota
	cloudLogging                    = iota
)

type buildInCompression int

const (
	uncompressed buildInCompression = iota
	gZip                            = iota
)

// consumerCount returns the number of attached consumers, useful for detecting errors in pipelines
func (receiver *pubsubReceiver) consumerCount() int { _ = "STUB: not implemented"; return 0 }

func (receiver *pubsubReceiver) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

// we will rely on the attributes of the message to determine the signal, so we need all proto unmarshalers

func (receiver *pubsubReceiver) setMarshallerFromExtension(host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (receiver *pubsubReceiver) setMarshallerFromEncodingID(encodingID buildInEncoding) error {
	_ = "STUB: not implemented"
	return nil
}

func (receiver *pubsubReceiver) Shutdown(_ context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func decompress(payload []byte, compression buildInCompression) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (receiver *pubsubReceiver) handleTrace(ctx context.Context, message *pubsubpb.ReceivedMessage, compression buildInCompression) error {
	_ = "STUB: not implemented"
	return nil
}

func (receiver *pubsubReceiver) handleMetric(ctx context.Context, message *pubsubpb.ReceivedMessage, compression buildInCompression) error {
	_ = "STUB: not implemented"
	return nil
}

func (receiver *pubsubReceiver) handleLog(ctx context.Context, message *pubsubpb.ReceivedMessage, compression buildInCompression) error {
	_ = "STUB: not implemented"
	return nil
}

func (receiver *pubsubReceiver) increaseEncodingErrorMetric(ctx context.Context, signal string) {
	_ = "STUB: not implemented"
	return
}

func (receiver *pubsubReceiver) detectEncoding(attributes map[string]string) (otlpEncoding buildInEncoding, otlpCompression buildInCompression) {
	_ = "STUB: not implemented"
	return *new(buildInEncoding), *new(buildInCompression)
}

func convertEncoding(encodingConfig string) (encoding buildInEncoding) {
	_ = "STUB: not implemented"
	return *new(buildInEncoding)
}

func (receiver *pubsubReceiver) createMultiplexingReceiverHandler(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (receiver *pubsubReceiver) createReceiverHandler(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package awss3receiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awss3receiver"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/receiverhelper"
	"go.uber.org/zap"
)

type encodingExtension struct {
	extension component.Component
	suffix    string
}

type encodingExtensions []encodingExtension

type receiverProcessor interface {
	processReceivedData(ctx context.Context, receiver *awss3Receiver, key string, data []byte) error
}

type awss3Receiver struct {
	reader          s3Reader
	logger          *zap.Logger
	cancel          context.CancelFunc
	obsrecv         *receiverhelper.ObsReport
	encodingsConfig []Encoding
	telemetryType   string
	dataProcessor   receiverProcessor
	extensions      encodingExtensions
	notifier        statusNotifier
}

func newAWSS3Receiver(ctx context.Context, cfg *Config, telemetryType string, settings receiver.Settings, processor receiverProcessor) (*awss3Receiver, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create the appropriate reader based on configuration

func (r *awss3Receiver) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *awss3Receiver) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (r *awss3Receiver) receiveBytes(ctx context.Context, key string, data []byte) (err error) {
	_ = "STUB: not implemented"
	return nil
}

type traceReceiver struct {
	consumer consumer.Traces
}

func newAWSS3TraceReceiver(ctx context.Context, cfg *Config, traces consumer.Traces, settings receiver.Settings) (*awss3Receiver, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *traceReceiver) processReceivedData(ctx context.Context, rcvr *awss3Receiver, key string, data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

type metricsReceiver struct {
	consumer consumer.Metrics
}

func newAWSS3MetricsReceiver(ctx context.Context, cfg *Config, metrics consumer.Metrics, settings receiver.Settings) (*awss3Receiver, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *metricsReceiver) processReceivedData(ctx context.Context, rcvr *awss3Receiver, key string, data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

type logsReceiver struct {
	consumer consumer.Logs
}

func newAWSS3LogsReceiver(ctx context.Context, cfg *Config, logs consumer.Logs, settings receiver.Settings) (*awss3Receiver, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *logsReceiver) processReceivedData(ctx context.Context, rcvr *awss3Receiver, key string, data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func newEncodingExtensions(encodingsConfig []Encoding, host component.Host) (encodingExtensions, error) {
	_ = "STUB: not implemented"
	return *new(encodingExtensions), nil
}

func (encodings encodingExtensions) findExtension(key string) (component.Component, string) {
	_ = "STUB: not implemented"
	return *new(component.Component), ""
}

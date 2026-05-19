// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package logs // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/datadogexporter/internal/logs"

import (
	"context"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.uber.org/zap"
)

// Sender submits logs to Datadog intake
type Sender struct {
	logger  *zap.Logger
	api     *datadogV2.LogsApi
	verbose bool // reports whether payload contents should be dumped when logging at debug level
}

// logsV2 is the key in datadog ServerConfiguration
// It is being used to customize the endpoint for Datadog intake based on exporter configuration
// https://github.com/DataDog/datadog-api-client-go/blob/be7e034424012c7ee559a2153802a45df73232ea/api/datadog/configuration.go#L308
const logsV2 = "v2.LogsApi.SubmitLog"

// NewSender creates a new Sender
func NewSender(endpoint string, logger *zap.Logger, hcs confighttp.ClientConfig, verbose bool, apiKey string) *Sender {
	_ = "STUB: not implemented"
	return nil
}

// SubmitLogs submits the logs contained in payload to the Datadog intake
func (s *Sender) SubmitLogs(ctx context.Context, payload []datadogV2.HTTPLogItem) error {
	_ = "STUB: not implemented"
	return nil
}

// keeps track of the ddtags of log items for grouping purposes
// stores consecutive log items to be submitted together

// Correctly sets apiSubmitLogRequest ddtags field based on tags from translator Transform method

// Batches consecutive log items with the same tags to be submitted together

func (s *Sender) handleSubmitLog(ctx context.Context, batch []datadogV2.HTTPLogItem, tags string) error {
	_ = "STUB: not implemented"
	return nil
}

// 1KB message max
// ignore any error

// If response is nil keep an error retryable
// Assuming this is a transient error (e.g. network/connectivity blip)

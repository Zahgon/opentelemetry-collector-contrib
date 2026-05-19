// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package translator // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/datadogreceiver/internal/translator"
import (
	"go.opentelemetry.io/collector/pdata/plog"
)

// I think the struct we should be unmarshaling into (below) should work, but doesn't include the status field for some reason...
// https://github.com/DataDog/datadog-api-client-go/blob/151a146e7ae57e5e8ed620a9ae6f958ed7db1ca7/api/datadogV2/model_http_log_item.go#L14-L34
type DatadogLogPayload struct {
	Message   string `json:"message"`
	Status    string `json:"status"`
	Timestamp int64  `json:"timestamp"`
	Hostname  string `json:"hostname"`
	Service   string `json:"service"`
	Source    string `json:"ddsource"`
	Tags      string `json:"ddtags"`
}

func ToPlog(incomingLogs []*DatadogLogPayload) plog.Logs {
	_ = "STUB: not implemented"
	return *new(plog.Logs)
}

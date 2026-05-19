// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package tencentcloudlogserviceexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/tencentcloudlogserviceexporter"

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"go.uber.org/zap"

	cls "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/tencentcloudlogserviceexporter/internal/proto"
)

// logServiceClient log Service's client wrapper
type logServiceClient interface {
	// sendLogs send message to LogService
	sendLogs(logs []*cls.Log) error
}

type logServiceClientImpl struct {
	clientInstance *common.Client
	logset         string
	topic          string
	hashkey        string
	logger         *zap.Logger
}

// newLogServiceClient Create Log Service client
func newLogServiceClient(config *Config, logger *zap.Logger) logServiceClient {
	_ = "STUB: not implemented"
	return *new(logServiceClient)
}

// sendLogs send message to LogService
func (c *logServiceClientImpl) sendLogs(logs []*cls.Log) error {
	_ = "STUB: not implemented"
	return nil
}

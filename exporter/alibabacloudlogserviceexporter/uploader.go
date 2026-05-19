// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package alibabacloudlogserviceexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/alibabacloudlogserviceexporter"

import (
	sls "github.com/aliyun/aliyun-log-go-sdk"
	"github.com/aliyun/aliyun-log-go-sdk/producer"
	"go.uber.org/zap"
)

// logServiceClient log Service's client wrapper
type logServiceClient interface {
	// sendLogs send message to LogService
	sendLogs(logs []*sls.Log) error
}

type logServiceClientImpl struct {
	clientInstance *producer.Producer
	project        string
	logstore       string
	topic          string
	source         string
	logger         *zap.Logger
}

func getIPAddress() (ipAddress string, err error) { _ = "STUB: not implemented"; return "", nil }

// newLogServiceClient Create Log Service client
func newLogServiceClient(config *Config, logger *zap.Logger) (logServiceClient, error) {
	_ = "STUB: not implemented"
	return *new(logServiceClient), nil
}

// do not return error if get hostname or ip address fail

// sendLogs send message to LogService
func (c *logServiceClientImpl) sendLogs(logs []*sls.Log) error {
	_ = "STUB: not implemented"
	return nil
}

// Success is impl of producer.CallBack
func (*logServiceClientImpl) Success(*producer.Result) {
	_ = "STUB: not implemented"

	// Fail is impl of producer.CallBack
	return
}

func (c *logServiceClientImpl) Fail(result *producer.Result) { _ = "STUB: not implemented"; return }

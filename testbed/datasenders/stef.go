// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package datasenders // import "github.com/open-telemetry/opentelemetry-collector-contrib/testbed/datasenders"

import (
	"go.opentelemetry.io/collector/consumer"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/testbed/testbed"
)

type StefDataSender struct {
	testbed.DataSenderBase
	consumer.Metrics
	Logger *zap.Logger
}

// NewStefDataSender creates a new STEF sender that will send
// to the specified port after Start is called.
func NewStefDataSender(host string, port int) *StefDataSender {
	_ = "STUB: not implemented"
	return nil
}

func (sds *StefDataSender) Start() error { _ = "STUB: not implemented"; return nil }

func (sds *StefDataSender) GenConfigYAMLStr() string { _ = "STUB: not implemented"; return "" }

func (*StefDataSender) ProtocolName() string { _ = "STUB: not implemented"; return "" }

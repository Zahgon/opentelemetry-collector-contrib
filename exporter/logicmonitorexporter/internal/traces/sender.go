// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package traces // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/logicmonitorexporter/internal/traces"

import (
	"context"
	"net/http"

	lmsdktraces "github.com/logicmonitor/lm-data-sdk-go/api/traces"
	"github.com/logicmonitor/lm-data-sdk-go/utils"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"
)

type Sender struct {
	logger            *zap.Logger
	traceIngestClient *lmsdktraces.LMTraceIngest
}

// NewSender creates a new Sender
func NewSender(ctx context.Context, endpoint string, client *http.Client, authParams utils.AuthParams, logger *zap.Logger) (*Sender, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Sender) SendTraces(ctx context.Context, td ptrace.Traces) error {
	_ = "STUB: not implemented"
	return nil
}

// Does the 'code' indicate a permanent error
func isPermanentClientFailure(code int) bool { _ = "STUB: not implemented"; return false }

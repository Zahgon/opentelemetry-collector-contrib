// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package logs // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/logicmonitorexporter/internal/logs"

import (
	"context"

	lmsdklogs "github.com/logicmonitor/lm-data-sdk-go/api/logs"
	"github.com/logicmonitor/lm-data-sdk-go/model"
	"go.uber.org/zap"
)

type Sender struct {
	logger          *zap.Logger
	logIngestClient *lmsdklogs.LMLogIngest
}

// NewSender creates a new Sender
func NewSender(ctx context.Context, logger *zap.Logger, opts ...lmsdklogs.Option) (*Sender, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Sender) SendLogs(ctx context.Context, payload []model.LogInput) error {
	_ = "STUB: not implemented"
	return nil
}

// Does the 'code' indicate a permanent error
func isPermanentClientFailure(code int) bool { _ = "STUB: not implemented"; return false }

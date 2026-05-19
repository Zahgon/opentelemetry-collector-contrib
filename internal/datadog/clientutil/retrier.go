// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package clientutil // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/datadog/clientutil"

import (
	"context"

	"go.opentelemetry.io/collector/config/configretry"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/datadog/scrub"
)

type Retrier struct {
	cfg      configretry.BackOffConfig
	logger   *zap.Logger
	scrubber scrub.Scrubber
}

func NewRetrier(logger *zap.Logger, settings configretry.BackOffConfig, scrubber scrub.Scrubber) *Retrier {
	_ = "STUB: not implemented"
	return nil
}

// DoWithRetries does a function with retries. This is a condensed version of the code on
// the exporterhelper, which we reuse here since we want custom retry logic.
func (r *Retrier) DoWithRetries(ctx context.Context, fn func(context.Context) error) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Do not use NewExponentialBackOff since it calls Reset and the code here must
// call Reset after changing the InitialInterval (this saves an unnecessary call to Now).

// back-off, but get interrupted when shutting down or request is cancelled or timed out.

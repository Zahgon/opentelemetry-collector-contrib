// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package metrics // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/transformprocessor/internal/metrics"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottldatapoint"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlmetric"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/transformprocessor/internal/common"
)

type Processor struct {
	contexts []common.MetricsConsumer
	logger   *zap.Logger
}

func NewProcessor(contextStatements []common.ContextStatements, errorMode ottl.ErrorMode, settings component.TelemetrySettings, metricFunctions map[string]ottl.Factory[*ottlmetric.TransformContext], dataPointFunctions map[string]ottl.Factory[*ottldatapoint.TransformContext]) (*Processor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Processor) ProcessMetrics(ctx context.Context, md pmetric.Metrics) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

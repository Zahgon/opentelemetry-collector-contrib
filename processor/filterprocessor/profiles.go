// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package filterprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/filterprocessor"

import (
	"context"

	"go.opentelemetry.io/collector/pdata/pprofile"
	"go.opentelemetry.io/collector/processor"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/filter/expr"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlprofile"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlresource"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/filterprocessor/internal/condition"
)

type filterProfileProcessor struct {
	consumers        []condition.ProfilesConsumer
	skipResourceExpr expr.BoolExpr[*ottlresource.TransformContext]
	skipProfileExpr  expr.BoolExpr[*ottlprofile.TransformContext]
	telemetry        *filterTelemetry
	logger           *zap.Logger
}

func newFilterProfilesProcessor(set processor.Settings, cfg *Config) (*filterProfileProcessor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// processProfiles filters the given profile based off the filterSampleProcessor's filters.
func (fpp *filterProfileProcessor) processProfiles(ctx context.Context, pd pprofile.Profiles) (pprofile.Profiles, error) {
	_ = "STUB: not implemented"
	return *new(pprofile.Profiles), nil
}

func (fpp *filterProfileProcessor) processSkipExpression(ctx context.Context, pd pprofile.Profiles) (pprofile.Profiles, error) {
	_ = "STUB: not implemented"
	return *new(pprofile.Profiles), nil
}

func (fpp *filterProfileProcessor) processConditions(ctx context.Context, pd pprofile.Profiles) (pprofile.Profiles, error) {
	_ = "STUB: not implemented"
	return *new(pprofile.Profiles), nil
}

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package condition // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/filterprocessor/internal/condition"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pprofile"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/filter/expr"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlprofile"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlresource"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlscope"
)

type ProfilesConsumer struct {
	resourceExpr expr.BoolExpr[*ottlresource.TransformContext]
	scopeExpr    expr.BoolExpr[*ottlscope.TransformContext]
	profileExpr  expr.BoolExpr[*ottlprofile.TransformContext]
}

// parsedProfileConditions is the type R for ParserCollection[R] that holds parsed OTTL conditions
type parsedProfileConditions struct {
	resourceConditions []*ottl.Condition[*ottlresource.TransformContext]
	scopeConditions    []*ottl.Condition[*ottlscope.TransformContext]
	profileConditions  []*ottl.Condition[*ottlprofile.TransformContext]
	telemetrySettings  component.TelemetrySettings
	errorMode          ottl.ErrorMode
}

func (pc ProfilesConsumer) ConsumeProfiles(ctx context.Context, pd pprofile.Profiles) error {
	_ = "STUB: not implemented"
	return nil
}

func newProfileConditionsFromResource(rc []*ottl.Condition[*ottlresource.TransformContext], telemetrySettings component.TelemetrySettings, errorMode ottl.ErrorMode) parsedProfileConditions {
	_ = "STUB: not implemented"
	return *new(parsedProfileConditions)
}

func newProfileConditionsFromScope(sc []*ottl.Condition[*ottlscope.TransformContext], telemetrySettings component.TelemetrySettings, errorMode ottl.ErrorMode) parsedProfileConditions {
	_ = "STUB: not implemented"
	return *new(parsedProfileConditions)
}

func newProfilesConsumer(ppc *parsedProfileConditions) ProfilesConsumer {
	_ = "STUB: not implemented"
	return *new(ProfilesConsumer)
}

type ProfileParserCollection ottl.ParserCollection[parsedProfileConditions]

type ProfileParserCollectionOption ottl.ParserCollectionOption[parsedProfileConditions]

func WithProfileParser(functions map[string]ottl.Factory[*ottlprofile.TransformContext]) ProfileParserCollectionOption {
	_ = "STUB: not implemented"
	return *new(ProfileParserCollectionOption)
}

func WithProfileErrorMode(errorMode ottl.ErrorMode) ProfileParserCollectionOption {
	_ = "STUB: not implemented"
	return *new(ProfileParserCollectionOption)
}

func WithProfileCommonParsers(functions map[string]ottl.Factory[*ottlresource.TransformContext]) ProfileParserCollectionOption {
	_ = "STUB: not implemented"
	return *new(ProfileParserCollectionOption)
}

func NewProfileParserCollection(settings component.TelemetrySettings, options ...ProfileParserCollectionOption) (*ProfileParserCollection, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func convertProfileConditions(pc *ottl.ParserCollection[parsedProfileConditions], conditions ottl.ConditionsGetter, parsedConditions []*ottl.Condition[*ottlprofile.TransformContext]) (parsedProfileConditions, error) {
	_ = "STUB: not implemented"
	return *new(parsedProfileConditions), nil
}

func (ppc *ProfileParserCollection) ParseContextConditions(contextConditions ContextConditions) (ProfilesConsumer, error) {
	_ = "STUB: not implemented"
	return *new(ProfilesConsumer), nil
}

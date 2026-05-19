// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package common // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/transformprocessor/internal/common"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pprofile"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/filter/expr"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlprofile"
)

type ProfilesConsumer interface {
	Context() ContextID
	ConsumeProfiles(ctx context.Context, ld pprofile.Profiles) error
}

type profileStatements struct {
	ottl.StatementSequence[*ottlprofile.TransformContext]
	expr.BoolExpr[*ottlprofile.TransformContext]
}

func (profileStatements) Context() ContextID { _ = "STUB: not implemented"; return *new(ContextID) }

func (l profileStatements) ConsumeProfiles(ctx context.Context, ld pprofile.Profiles) error {
	_ = "STUB: not implemented"
	return nil
}

type ProfileParserCollection ottl.ParserCollection[ProfilesConsumer]

type ProfileParserCollectionOption ottl.ParserCollectionOption[ProfilesConsumer]

func WithProfileParser(functions map[string]ottl.Factory[*ottlprofile.TransformContext]) ProfileParserCollectionOption {
	_ = "STUB: not implemented"
	return *new(ProfileParserCollectionOption)
}

func WithProfileErrorMode(errorMode ottl.ErrorMode) ProfileParserCollectionOption {
	_ = "STUB: not implemented"
	return *new(ProfileParserCollectionOption)
}

func NewProfileParserCollection(settings component.TelemetrySettings, options ...ProfileParserCollectionOption) (*ProfileParserCollection, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func convertProfileStatements(pc *ottl.ParserCollection[ProfilesConsumer], statements ottl.StatementsGetter, parsedStatements []*ottl.Statement[*ottlprofile.TransformContext]) (ProfilesConsumer, error) {
	_ = "STUB: not implemented"
	return *new(ProfilesConsumer), nil
}

func (ppc *ProfileParserCollection) ParseContextStatements(contextStatements ContextStatements) (ProfilesConsumer, error) {
	_ = "STUB: not implemented"
	return *new(ProfilesConsumer), nil
}

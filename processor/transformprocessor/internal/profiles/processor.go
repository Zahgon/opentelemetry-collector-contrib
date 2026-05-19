// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package profiles // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/transformprocessor/internal/profiles"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pprofile"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlprofile"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/transformprocessor/internal/common"
)

type Processor struct {
	contexts []common.ProfilesConsumer
	logger   *zap.Logger
}

func NewProcessor(contextStatements []common.ContextStatements, errorMode ottl.ErrorMode, settings component.TelemetrySettings, profileFunctions map[string]ottl.Factory[*ottlprofile.TransformContext]) (*Processor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Processor) ProcessProfiles(ctx context.Context, ld pprofile.Profiles) (pprofile.Profiles, error) {
	_ = "STUB: not implemented"
	return *new(pprofile.Profiles), nil
}

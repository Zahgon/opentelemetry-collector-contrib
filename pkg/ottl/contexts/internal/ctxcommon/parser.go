// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ctxcommon // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/internal/ctxcommon"

import (
	"go.opentelemetry.io/collector/component"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/internal/ctxcache"
)

func NewParser[K any](
	functions map[string]ottl.Factory[K],
	telemetrySettings component.TelemetrySettings,
	pathExpressionParser ottl.PathExpressionParser[K],
	enumParser ottl.EnumParser,
	options ...ottl.Option[K],
) (ottl.Parser[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func PathExpressionParser[K any](
	contextName string,
	contextDocRef string,
	cacheGetter ctxcache.Getter[K],
	contextParsers map[string]ottl.PathExpressionParser[K],
) ottl.PathExpressionParser[K] {
	_ = "STUB: not implemented"
	return nil
}

// Normalize context and segment name

// Allow cache access only on this context

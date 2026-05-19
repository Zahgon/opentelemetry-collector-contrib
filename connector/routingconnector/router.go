// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package routingconnector // import "github.com/open-telemetry/opentelemetry-collector-contrib/connector/routingconnector"

import (
	"errors"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pipeline"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottldatapoint"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottllog"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlmetric"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlresource"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlspan"
)

var errPipelineNotFound = errors.New("pipeline not found")

// consumerProvider is a function with a type parameter C (expected to be one
// of consumer.Traces, consumer.Metrics, or Consumer.Logs). returns a
// consumer for the given component ID(s).
type consumerProvider[C any] func(...pipeline.ID) (C, error)

// router registers consumers and default consumers for a pipeline. the type
// parameter C is expected to be one of: consumer.Traces, consumer.Metrics, or
// consumer.Logs.
type router[C any] struct {
	resourceParser   ottl.Parser[*ottlresource.TransformContext]
	spanParser       ottl.Parser[*ottlspan.TransformContext]
	metricParser     ottl.Parser[*ottlmetric.TransformContext]
	dataPointParser  ottl.Parser[*ottldatapoint.TransformContext]
	logParser        ottl.Parser[*ottllog.TransformContext]
	defaultConsumer  C
	logger           *zap.Logger
	routes           map[string]routingItem[C]
	consumerProvider consumerProvider[C]
	table            []RoutingTableItem
	routeSlice       []routingItem[C]
}

// newRouter creates a new router instance with based on type parameters C and K.
// see router struct definition for the allowed types.
func newRouter[C any](
	table []RoutingTableItem,
	defaultPipelineIDs []pipeline.ID,
	provider consumerProvider[C],
	settings component.TelemetrySettings,
) (*router[C], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type routingItem[C any] struct {
	consumer           C
	requestCondition   *requestCondition
	resourceStatement  *ottl.Statement[*ottlresource.TransformContext]
	spanStatement      *ottl.Statement[*ottlspan.TransformContext]
	metricStatement    *ottl.Statement[*ottlmetric.TransformContext]
	dataPointStatement *ottl.Statement[*ottldatapoint.TransformContext]
	logStatement       *ottl.Statement[*ottllog.TransformContext]
	statementContext   string
	action             Action
}

func (r *router[C]) buildParsers(table []RoutingTableItem, settings component.TelemetrySettings) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *router[C]) registerConsumers(defaultPipelineIDs []pipeline.ID) error {
	_ = "STUB: not implemented"
	// register default pipelines
	return nil
}

// register pipelines for each route

// registerDefaultConsumer registers a consumer for the default pipelines configured
func (r *router[C]) registerDefaultConsumer(pipelineIDs []pipeline.ID) error {
	_ = "STUB: not implemented"
	return nil
}

// convert conditions to statements
func (r *router[C]) normalizeConditions() { _ = "STUB: not implemented"; return }

// registerRouteConsumers registers a consumer for the pipelines configured for each route
func (r *router[C]) registerRouteConsumers() (err error) { _ = "STUB: not implemented"; return nil }

func key(entry RoutingTableItem) string { _ = "STUB: not implemented"; return "" }

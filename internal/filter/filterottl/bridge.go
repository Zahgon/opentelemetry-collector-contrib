// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package filterottl // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/filter/filterottl"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/filter/expr"
	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/filter/filterconfig"
	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/filter/filterset"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottllog"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlmetric"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlresource"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlspan"
)

const (
	serviceNameStaticStatement        = `resource.attributes["service.name"] == "%v"`
	nameStaticStatement               = `name == "%v"`
	spanKindStaticStatement           = `kind.deprecated_string == "%v"`
	scopeNameStaticStatement          = `instrumentation_scope.name == "%v"`
	scopeVersionStaticStatement       = `instrumentation_scope.version == "%v"`
	attributesStaticStatement         = `attributes["%v"] == %v`
	resourceAttributesStaticStatement = `resource.attributes["%v"] == %v`
	bodyStaticStatement               = `body.string == "%v"`
	severityTextStaticStatement       = `severity_text == "%v"`

	serviceNameRegexStatement        = `IsMatch(resource.attributes["service.name"], "%v")`
	nameRegexStatement               = `IsMatch(name, "%v")`
	spanKindRegexStatement           = `IsMatch(kind.deprecated_string, "%v")`
	scopeNameRegexStatement          = `IsMatch(instrumentation_scope.name, "%v")`
	scopeVersionRegexStatement       = `IsMatch(instrumentation_scope.version, "%v")`
	attributesRegexStatement         = `IsMatch(attributes["%v"], "%v")`
	resourceAttributesRegexStatement = `IsMatch(resource.attributes["%v"], "%v")`
	bodyRegexStatement               = `IsMatch(body.string, "%v")`
	severityTextRegexStatement       = `IsMatch(severity_text, "%v")`

	// Boolean expression for existing severity number matching
	// a -> lr.SeverityNumber() == plog.SeverityNumberUnspecified
	// b -> snm.matchUndefined
	// c -> lr.SeverityNumber() >= snm.minSeverityNumber
	// (a AND b) OR ( NOT a AND c)
	//  a  b  c  X
	//  0  0  0  0
	//  0  0  1  1
	//  0  1  0  0
	//  0  1  1  1
	//  1  0  0  0
	//  1  0  1  0
	//  1  1  0  1
	//  1  1  1  1
	severityNumberStatement = `((severity_number == SEVERITY_NUMBER_UNSPECIFIED and %v) or (severity_number != SEVERITY_NUMBER_UNSPECIFIED and severity_number >= %d))`
)

func NewLogSkipExprBridge(mc *filterconfig.MatchConfig) (expr.BoolExpr[*ottllog.TransformContext], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewResourceSkipExprBridge(mc *filterconfig.MatchConfig) (expr.BoolExpr[*ottlresource.TransformContext], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// OTTL treats resource attributes as attributes for the resource context.

// OTTL treats resource attributes as attributes for the resource context.

func NewSpanSkipExprBridge(mc *filterconfig.MatchConfig) (expr.BoolExpr[*ottlspan.TransformContext], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createStatement(mp filterconfig.MatchProperties) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

type conditionStatements struct {
	serviceNameConditions       []string
	spanNameConditions          []string
	spanKindConditions          []string
	scopeNameConditions         []string
	scopeVersionConditions      []string
	attributeConditions         []string
	resourceAttributeConditions []string
	bodyConditions              []string
	severityTextConditions      []string
	severityNumberCondition     *string
}

func createConditions(mp filterconfig.MatchProperties) (conditionStatements, error) {
	_ = "STUB: not implemented"
	return *new(conditionStatements), nil
}

type statementTemplates struct {
	serviceNameStatement  string
	spanNameStatement     string
	spanKindStatement     string
	scopeNameStatement    string
	scopeVersionStatement string
	attrStatement         string
	resourceAttrStatement string
	bodyStatement         string
	severityTextStatement string
}

func createStatementTemplates(matchType filterset.MatchType) (statementTemplates, error) {
	_ = "STUB: not implemented"
	return *new(statementTemplates), nil
}

func createBasicConditions(template string, input []string) []string {
	_ = "STUB: not implemented"
	return nil
}

func createLibraryConditions(nameTemplate, versionTemplate string, libraries []filterconfig.InstrumentationLibrary) ([]string, []string) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createAttributeConditions(template string, input []filterconfig.Attribute, matchType filterset.MatchType) []string {
	_ = "STUB: not implemented"
	return nil
}

func convertAttribute(value any) string { _ = "STUB: not implemented"; return "" }

func createSeverityNumberConditions(severityNumberProperties *filterconfig.LogSeverityNumberMatchProperties) *string {
	_ = "STUB: not implemented"
	return nil
}

func NewMetricSkipExprBridge(include, exclude *filterconfig.MetricMatchProperties) (expr.BoolExpr[*ottlmetric.TransformContext], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createMetricStatement(mp filterconfig.MetricMatchProperties) (*string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

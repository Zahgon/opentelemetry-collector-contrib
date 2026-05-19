// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package transformprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/transformprocessor"

import (
	"errors"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/confmap"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottldatapoint"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottllog"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlmetric"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlprofile"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlspan"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlspanevent"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/transformprocessor/internal/common"
)

var errFlatLogsGateDisabled = errors.New("'flatten_data' requires the 'transform.flatten.logs' feature gate to be enabled")

// Config defines the configuration for the processor.
type Config struct {
	// ErrorMode determines how the processor reacts to errors that occur while processing a statement.
	// Valid values are `ignore` and `propagate`.
	// `ignore` means the processor ignores errors returned by statements and continues on to the next statement. This is the recommended mode.
	// `propagate` means the processor returns the error up the pipeline.  This will result in the payload being dropped from the collector.
	// The default value is `ignore`, which can be changed back to `propagate` by disabling the `processor.transform.defaultErrorModeIgnore` feature gate.
	ErrorMode ottl.ErrorMode `mapstructure:"error_mode"`

	TraceStatements   []common.ContextStatements `mapstructure:"trace_statements"`
	MetricStatements  []common.ContextStatements `mapstructure:"metric_statements"`
	LogStatements     []common.ContextStatements `mapstructure:"log_statements"`
	ProfileStatements []common.ContextStatements `mapstructure:"profile_statements"`

	FlattenData bool `mapstructure:"flatten_data"`
	logger      *zap.Logger

	dataPointFunctions map[string]ottl.Factory[*ottldatapoint.TransformContext]
	logFunctions       map[string]ottl.Factory[*ottllog.TransformContext]
	metricFunctions    map[string]ottl.Factory[*ottlmetric.TransformContext]
	spanEventFunctions map[string]ottl.Factory[*ottlspanevent.TransformContext]
	spanFunctions      map[string]ottl.Factory[*ottlspan.TransformContext]
	profileFunctions   map[string]ottl.Factory[*ottlprofile.TransformContext]
}

// Unmarshal is used internally by mapstructure to parse the transformprocessor configuration (Config),
// adding support to structured and flat configuration styles.
// When the flat configuration style is used, all statements are grouped into a common.ContextStatements
// object, with empty [common.ContextStatements.Context] value.
// On the other hand, structured configurations are parsed following the mapstructure Config format.
//
// Example of flat configuration:
//
//	log_statements:
//	  - set(attributes["service.new_name"], attributes["service.name"])
//	  - delete_key(attributes, "service.name")
//
// Example of structured configuration:
//
//	log_statements:
//	  - context: "span"
//	    statements:
//	      - set(attributes["service.new_name"], attributes["service.name"])
//	      - delete_key(attributes, "service.name")
func (c *Config) Unmarshal(conf *confmap.Conf) error { _ = "STUB: not implemented"; return nil }

// Array of strings means it's a basic configuration style

var _ component.Config = (*Config)(nil)

func (c *Config) Validate() error { _ = "STUB: not implemented"; return nil }

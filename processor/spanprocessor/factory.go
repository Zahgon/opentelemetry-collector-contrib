// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package spanprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/spanprocessor"

import (
	"context"
	"errors"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/processor"
)

const (
	// status represents span status
	statusCodeUnset = "Unset"
	statusCodeError = "Error"
	statusCodeOk    = "Ok"
)

var processorCapabilities = consumer.Capabilities{MutatesData: true}

// errMissingRequiredField is returned when a required field in the config
// is not specified.
// TODO https://github.com/open-telemetry/opentelemetry-collector/issues/215
//
//	Move this to the error package that allows for span name and field to be specified.
var (
	errMissingRequiredField       = errors.New("error creating \"span\" processor: either \"from_attributes\" or \"to_attributes\" must be specified in \"name:\" or \"setStatus\" must be specified")
	errIncorrectStatusCode        = errors.New("error creating \"span\" processor: \"status\" must have specified \"code\" as \"Ok\" or \"Error\" or \"Unset\"")
	errIncorrectStatusDescription = errors.New("error creating \"span\" processor: \"description\" can be specified only for \"code\" \"Error\"")
)

// NewFactory returns a new factory for the Span processor.
func NewFactory() processor.Factory { _ = "STUB: not implemented"; return *new(processor.Factory) }

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

func createTracesProcessor(
	ctx context.Context,
	set processor.Settings,
	cfg component.Config,
	nextConsumer consumer.Traces,
) (processor.Traces, error) {
	_ = "STUB: not implemented"
	// 'from_attributes' or 'to_attributes' under 'name' has to be set for the span
	// processor to be valid. If not set and not enforced, the processor would do no work.
	return *new(processor.Traces), nil
}

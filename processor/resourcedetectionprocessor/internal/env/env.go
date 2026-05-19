// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package env provides a detector that loads resource information from
// the OTEL_RESOURCE environment variable. A list of labels of the form
// `<key1>=<value1>,<key2>=<value2>,...` is accepted. Domain names and
// paths are accepted as label keys.
package env // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/env"

import (
	"context"
	"regexp"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/processor"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal"
)

// TypeStr is type of detector.
const TypeStr = "env"

// Environment variable used by "env" to decode a resource.
const envVar = "OTEL_RESOURCE_ATTRIBUTES"

// Deprecated environment variable used by "env" to decode a resource.
// Specification states to use OTEL_RESOURCE_ATTRIBUTES however to avoid
// breaking existing usage, maintain support for OTEL_RESOURCE
// https://github.com/open-telemetry/opentelemetry-specification/blob/1afab39e5658f807315abf2f3256809293bfd421/specification/resource/sdk.md#specifying-resource-information-via-an-environment-variable
const deprecatedEnvVar = "OTEL_RESOURCE"

var _ internal.Detector = (*Detector)(nil)

type Detector struct{}

func NewDetector(processor.Settings, internal.DetectorConfig) (internal.Detector, error) {
	_ = "STUB: not implemented"
	return *new(internal.Detector), nil
}

func (*Detector) Detect(context.Context) (resource pcommon.Resource, schemaURL string, err error) {
	_ = "STUB: not implemented"
	return *new(pcommon.Resource), "", nil
}

// labelRegex matches any key=value pair including a trailing comma or the end of the
// string. Captures the trimmed key & value parts, and ignores any superfluous spaces.
var labelRegex = regexp.MustCompile(`\s*([[:ascii:]]{1,256}?)\s*=\s*([[:ascii:]]{0,256}?)\s*(?:,|$)`)

func initializeAttributeMap(am pcommon.Map, s string) error { _ = "STUB: not implemented"; return nil }

// if there is any text between matches, raise an error

// if there is any text after the last match, raise an error

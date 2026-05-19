// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package dynatrace provides a detector that loads resource information from
// the dt_host_metadata.properties file which is located in
// the /var/lib/dynatrace/enrichment (on *nix systems) and %ProgramData%\dynatrace\enrichment
// (on Windows) directories.

package dynatrace // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/dynatrace"
import (
	"context"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/processor"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal"
)

const TypeStr = "dynatrace"

const dtHostMetadataProperties = "dt_host_metadata.properties"

var dtHostProperties = []string{"dt.entity.host", "host.name", "dt.smartscape.host"}

type Detector struct {
	enrichmentDirectory string
	logger              *zap.Logger
}

func NewDetector(set processor.Settings, _ internal.DetectorConfig) (internal.Detector, error) {
	_ = "STUB: not implemented"
	return *new(internal.Detector), nil
}

// Windows default is "%ProgramData%\dynatrace\enrichment"
// If the ProgramData environment variable is not set,
// it falls back to C:\ProgramData

func (d Detector) Detect(_ context.Context) (pcommon.Resource, string, error) {
	_ = "STUB: not implemented"
	return *new(pcommon.Resource), "", nil
}

func (d Detector) readPropertiesFile(attributes pcommon.Map) error {
	_ = "STUB: not implemented"
	return nil
}

// split by the first "=" character. If there is another "=" afterward, this will be part of the value

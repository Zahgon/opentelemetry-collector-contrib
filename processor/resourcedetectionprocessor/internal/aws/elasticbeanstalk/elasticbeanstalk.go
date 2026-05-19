// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package elasticbeanstalk // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/aws/elasticbeanstalk"

import (
	"context"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/processor"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/aws/elasticbeanstalk/internal/metadata"
)

const (
	// TypeStr is type of detector.
	TypeStr = "elastic_beanstalk"

	linuxPath   = "/var/elasticbeanstalk/xray/environment.conf"
	windowsPath = "C:\\Program Files\\Amazon\\XRay\\environment.conf"
)

var _ internal.Detector = (*Detector)(nil)

type Detector struct {
	fs fileSystem
	rb *metadata.ResourceBuilder
}

type EbMetaData struct {
	DeploymentID    int    `json:"deployment_id"`
	EnvironmentName string `json:"environment_name"`
	VersionLabel    string `json:"version_label"`
}

func NewDetector(_ processor.Settings, dcfg internal.DetectorConfig) (internal.Detector, error) {
	_ = "STUB: not implemented"
	return *new(internal.Detector), nil
}

func (d Detector) Detect(context.Context) (resource pcommon.Resource, schemaURL string, err error) {
	_ = "STUB: not implemented"
	return *new(pcommon.Resource), "", nil
}

// Do not want to return error so it fails silently on non-EB instances

// TODO: Log a more specific message with zap

// TODO: Log a more specific error with zap

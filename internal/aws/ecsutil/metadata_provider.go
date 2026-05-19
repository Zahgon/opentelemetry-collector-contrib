// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ecsutil // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/ecsutil"

import (
	"go.opentelemetry.io/collector/component"
	"go.uber.org/zap"
)

type MetadataProvider interface {
	FetchTaskMetadata() (*TaskMetadata, error)
	FetchContainerMetadata() (*ContainerMetadata, error)
}

type ecsMetadataProviderImpl struct {
	logger *zap.Logger
	client RestClient
}

var _ MetadataProvider = &ecsMetadataProviderImpl{}

func NewTaskMetadataProvider(client RestClient, logger *zap.Logger) MetadataProvider {
	_ = "STUB: not implemented"
	return *new(MetadataProvider)
}

func NewDetectedTaskMetadataProvider(set component.TelemetrySettings) (MetadataProvider, error) {
	_ = "STUB: not implemented"
	return *new(MetadataProvider), nil
}

// FetchTaskMetadata retrieves the metadata for a task running on Amazon ECS
func (md *ecsMetadataProviderImpl) FetchTaskMetadata() (*TaskMetadata, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FetchContainerMetadata retrieves the metadata for the Amazon ECS Container the collector is running on
func (md *ecsMetadataProviderImpl) FetchContainerMetadata() (*ContainerMetadata, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

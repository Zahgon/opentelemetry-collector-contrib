// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package docker // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/metadataproviders/docker"

import (
	"context"

	"github.com/moby/moby/api/types/container"
	"github.com/moby/moby/client"
)

type Provider interface {
	// Hostname returns the OS hostname
	Hostname(context.Context) (string, error)

	// OSType returns the host operating system
	OSType(context.Context) (string, error)

	// ContainerInfo returns the current container information
	ContainerInfo(context.Context) (container.InspectResponse, error)
}

type dockerProviderImpl struct {
	dockerClient *client.Client
}

func NewProvider(opts ...client.Opt) (Provider, error) {
	_ = "STUB: not implemented"
	return *new(Provider), nil
}

func (d *dockerProviderImpl) Hostname(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (d *dockerProviderImpl) OSType(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (d *dockerProviderImpl) ContainerInfo(ctx context.Context) (container.InspectResponse, error) {
	_ = "STUB: not implemented"
	return *new(container.InspectResponse), nil
}

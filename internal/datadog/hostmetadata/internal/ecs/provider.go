// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package ecs contains the ECS Fargate hostname provider
package ecs // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/datadog/hostmetadata/internal/ecs"

import (
	"context"
	"errors"

	"github.com/DataDog/datadog-agent/pkg/opentelemetry-mapping-go/otlp/attributes/source"
	"go.opentelemetry.io/collector/component"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/ecsutil"
)

var ErrNotOnECSFargate = errors.New("not running on ECS Fargate")

var _ source.Provider = (*Provider)(nil)

type Provider struct {
	missingEndpoint bool
	ecsMetadata     ecsutil.MetadataProvider
}

// OnECSFargate determines if the application is running on ECS Fargate.
func (p *Provider) OnECSFargate(_ context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false,

		// No ECS metadata endpoint, therefore not on ECS Fargate
		nil
}

// Source returns the task ARN of the ECS Fargate task if on ECS Fargate.
func (p *Provider) Source(ctx context.Context) (source.Source, error) {
	_ = "STUB: not implemented"
	return *new(source.Source), nil
}

// Not on ECS Fargate

// Failed to determine if on ECS Fargate

// NewProvider creates a new ECS Fargate hostname provider.
func NewProvider(set component.TelemetrySettings) (*Provider, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Metadata endpoint has not been detected

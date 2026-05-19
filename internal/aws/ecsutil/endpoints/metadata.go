// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package endpoints // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/ecsutil/endpoints"

import (
	"net/url"
)

const (
	TaskMetadataEndpointV3EnvVar = "ECS_CONTAINER_METADATA_URI"
	TaskMetadataEndpointV4EnvVar = "ECS_CONTAINER_METADATA_URI_V4"

	TaskMetadataPath      = "/task"
	ContainerMetadataPath = ""
)

// ErrNoTaskMetadataEndpointDetected is a reserved error type to distinguish between incompatible environments
// and other error scenarios
type ErrNoTaskMetadataEndpointDetected struct {
	error
	MissingVersion int
}

// GetTMEV3FromEnv will return a validated task metadata endpoint as obtained by the v3 env var, if any.
func GetTMEV3FromEnv() (endpoint *url.URL, err error) { _ = "STUB: not implemented"; return nil, nil }

// GetTMEV4FromEnv will return a validated task metadata endpoint as obtained by the v4 env var, if any.
func GetTMEV4FromEnv() (endpoint *url.URL, err error) { _ = "STUB: not implemented"; return nil, nil }

// GetTMEFromEnv will return the first available task metadata endpoint for the v4 or v3 env var in that order.
func GetTMEFromEnv() (endpoint *url.URL, err error) { _ = "STUB: not implemented"; return nil, nil }

func validateEndpoint(candidate string) (endpoint *url.URL, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

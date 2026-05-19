// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ecsobserver // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/ecsobserver"

import (
	"regexp"

	ecstypes "github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"go.uber.org/zap"
)

type ServiceConfig struct {
	CommonExporterConfig `mapstructure:",squash" yaml:",inline"`

	// NamePattern is mandatory.
	NamePattern string `mapstructure:"name_pattern" yaml:"name_pattern"`
	// ContainerNamePattern is optional, empty string means all containers in that service would be exported.
	// Otherwise both service and container name patterns need to match.
	ContainerNamePattern string `mapstructure:"container_name_pattern" yaml:"container_name_pattern"`
}

func (s *ServiceConfig) validate() error { _ = "STUB: not implemented"; return nil }

func (s *ServiceConfig) newMatcher(opts matcherOptions) (targetMatcher, error) {
	_ = "STUB: not implemented"
	return *new(targetMatcher), nil
}

func serviceConfigsToMatchers(cfgs []ServiceConfig) []matcherConfig {
	_ = "STUB: not implemented"
	return nil
}

// NOTE: &cfg points to the temp var, whose value would end up be the last one in the slice.

type serviceMatcher struct {
	logger    *zap.Logger
	cfg       ServiceConfig
	nameRegex *regexp.Regexp
	// can be nil, which means matching all the container in the task (whose service name is matched by nameRegex)
	containerNameRegex *regexp.Regexp
	exportSetting      *commonExportSetting
}

func (*serviceMatcher) matcherType() matcherType {
	_ = "STUB: not implemented"
	return *new(matcherType)
}

func (s *serviceMatcher) matchTargets(t *taskAnnotated, c ecstypes.ContainerDefinition) ([]matchedTarget, error) {
	_ = "STUB: not implemented"
	// Service info is only attached for tasks whose services are included in config.
	// However, Match is called on tasks so we need to guard nil pointer.
	return nil, nil
}

// The rest is same as taskDefinitionMatcher

// serviceConfigsToFilter reduce number of describe service API call
func serviceConfigsToFilter(cfgs []ServiceConfig) (serviceNameFilter, error) {
	_ = "STUB: not implemented"
	// If no service config, don't describe any services
	return *new(serviceNameFilter), nil
}

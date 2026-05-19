// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ecsobserver // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/ecsobserver"

import (
	"regexp"

	ecstypes "github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"go.uber.org/zap"
)

type TaskDefinitionConfig struct {
	CommonExporterConfig `mapstructure:",squash" yaml:",inline"`

	// ArnPattern is mandatory, empty string means arn based match is skipped.
	ArnPattern string `mapstructure:"arn_pattern" yaml:"arn_pattern"`
	// ContainerNamePattern is optional, empty string means all containers in that task definition would be exported.
	// Otherwise both service and container name patterns need to match.
	ContainerNamePattern string `mapstructure:"container_name_pattern" yaml:"container_name_pattern"`
}

func (t *TaskDefinitionConfig) validate() error { _ = "STUB: not implemented"; return nil }

func (t *TaskDefinitionConfig) newMatcher(opts matcherOptions) (targetMatcher, error) {
	_ = "STUB: not implemented"
	return *new(targetMatcher), nil
}

func taskDefinitionConfigsToMatchers(cfgs []TaskDefinitionConfig) []matcherConfig {
	_ = "STUB: not implemented"
	return nil
}

// NOTE: &cfg points to the temp var, whose value would end up be the last one in the slice.

type taskDefinitionMatcher struct {
	logger *zap.Logger
	cfg    TaskDefinitionConfig
	// should never be nil because Init must reject it and caller should stop
	arnRegex *regexp.Regexp
	// if nil, matches all the container in the task (whose task definition name is matched by arnRegex)
	containerNameRegex *regexp.Regexp
	exportSetting      *commonExportSetting
}

func (*taskDefinitionMatcher) matcherType() matcherType {
	_ = "STUB: not implemented"
	return *new(matcherType)
}

func (m *taskDefinitionMatcher) matchTargets(t *taskAnnotated, c ecstypes.ContainerDefinition) ([]matchedTarget, error) {
	_ = "STUB: not implemented"
	// Check arn
	return nil, nil
}

// The rest is same as ServiceMatcher

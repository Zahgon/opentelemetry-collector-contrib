// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ecsobserver // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/ecsobserver"

import (
	"errors"
	"regexp"

	ecstypes "github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"go.uber.org/zap"
)

type targetMatcher interface {
	// type() won't parse
	matcherType() matcherType
	// matchTargets returns targets fond from the specific container.
	// One container can have multiple targets because it may have multiple ports.
	matchTargets(task *taskAnnotated, container ecstypes.ContainerDefinition) ([]matchedTarget, error)
}

// matcherConfig should be implemented by all the matcher config structs
// for validation and initializing the actual matcher implementation.
type matcherConfig interface {
	// validate calls NewMatcher and only returns the error, it can be used in test
	// and the new config validator interface.
	validate() error
	// newMatcher validates config and creates a targetMatcher implementation.
	// The error is a config validation error
	newMatcher(options matcherOptions) (targetMatcher, error)
}

type matcherOptions struct {
	Logger *zap.Logger
}

type matcherType int

// Values for enum matcherType.
const (
	matcherTypeService matcherType = iota + 1
	matcherTypeTaskDefinition
	matcherTypeDockerLabel
)

func (t matcherType) String() string { _ = "STUB: not implemented"; return "" }

// Give it a _matcher_type suffix so people can find it by string search.

type matchResult struct {
	// Tasks are index for tasks that include matched containers
	Tasks []int
	// Containers are index for matched containers. containers should show up in the original order of the task list and container definitions.
	Containers []matchedContainer
}

type matchedContainer struct {
	TaskIndex      int // Index in task list before filter, i.e. after fetch and decorate
	ContainerIndex int // Index within a tasks definition's container list
	Targets        []matchedTarget
}

// MergeTargets adds new targets to the set, the 'key' is port + metrics path.
// The 'key' does not contain an IP address because all targets from one
// container have the same IP address. If there are duplicate 'key's we honor
// the existing target and do not override.  Duplication could happen if there
// are several rules matching same target.
func (mc *matchedContainer) MergeTargets(newTargets []matchedTarget) {
	_ = "STUB: not implemented"
	return
}

// If port and metrics_path are same, then we treat them as same target and keep the existing one

// matchedTarget contains info for exporting prometheus scrape target
// and tracing back into the config (can be used in stats, error reporting etc.).
type matchedTarget struct {
	MatcherType  matcherType
	MatcherIndex int // Index within a specific matcher type
	Port         int
	MetricsPath  string
	Job          string
}

func matcherOrders() []matcherType { _ = "STUB: not implemented"; return nil }

func newMatchers(c Config, mOpt matcherOptions) (map[matcherType][]targetMatcher, error) {
	_ = "STUB: not implemented"
	// We can have a registry or factory methods etc. but we only have three type of matchers
	// and likely not going to add anymore in forseable future, just hard code the map here.
	// All the XXXConfigToMatchers looks like copy pasted funcs, but there is no generic way to do it.
	return nil, nil
}

// a global instance because it's expected and we don't care about why the container didn't match (for now).
// In the future we might add a debug flag for each matcher config and return typed error with more detail
// to help user debug. e.g. type ^ngix-*$ does not match nginx-service.
var errNotMatched = errors.New("container not matched")

// matchContainers apply one matcher to a list of tasks and returns matchResult.
// It does not modify the task in place, the attaching match result logic is
// performed by taskFilter at later stage.
func matchContainers(tasks []*taskAnnotated, matcher targetMatcher, matcherIndex int) (*matchResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NOTE: we don't stop when there is an error because it could be one task having invalid docker label.

// Keep track of unexpected error

// matchContainerByName is used by taskDefinitionMatcher and serviceMatcher.
// The only exception is DockerLabelMatcher because it get ports from docker label.
func matchContainerByName(nameRegex *regexp.Regexp, expSetting *commonExportSetting, container ecstypes.ContainerDefinition) ([]matchedTarget, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Match based on port

// Only export container if it has at least one matching port.

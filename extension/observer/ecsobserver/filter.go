// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ecsobserver // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/ecsobserver"

import (
	"go.uber.org/zap"
)

type taskFilter struct {
	logger   *zap.Logger
	matchers map[matcherType][]targetMatcher
}

func newTaskFilter(logger *zap.Logger, matchers map[matcherType][]targetMatcher) *taskFilter {
	_ = "STUB: not implemented"
	return nil
}

// Filter run all the matchers and return all the tasks that including at least one matched container.
func (f *taskFilter) filter(tasks []*taskAnnotated) ([]*taskAnnotated, error) {
	_ = "STUB: not implemented"
	// Group result by matcher type, each type can have multiple configs.
	return nil, nil
}

// for each type of matchers
// for individual matchers of same type

// NOTE: we continue the loop even if there is error because some tasks can has invalid labels.
// matchContainers always return non nil result even if there are errors during matching.

// TODO: print out the pattern to include both pattern and port

// Attach match result to tasks, do it in matcherOrders.
// AddMatchedContainer will merge in new targets if it is not already matched by other matchers.

// Sort by task index so the output is consistent.

// Sort containers within a task

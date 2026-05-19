// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build !linux

package cadvisor // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awscontainerinsightreceiver/internal/cadvisor"

import (
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awscontainerinsightreceiver/internal/cadvisor/extractors"
)

// cadvisor doesn't support windows, define the dummy functions

type HostInfo interface {
	GetNumCores() int64
	GetMemoryCapacity() int64
	GetClusterName() string
}

// Cadvisor is a dummy struct for windows
type Cadvisor struct{}

type Decorator interface {
	Decorate(*extractors.CAdvisorMetric) *extractors.CAdvisorMetric
	Shutdown() error
}

// Option is a function that can be used to configure Cadvisor struct
type Option func(*Cadvisor)

// WithDecorator constructs an option for configuring the metric decorator
func WithDecorator(_ any) Option {
	_ = "STUB: not implemented"
	return *

	// do nothing
	new(Option)
}

func WithECSInfoCreator(_ any) Option {
	_ = "STUB: not implemented"
	return *

	// do nothing
	new(Option)
}

// New is a dummy function to construct a dummy Cadvisor struct for windows
func New(_ string, _ HostInfo, _ *zap.Logger, _ ...Option) (*Cadvisor, error) {
	_ = "STUB: not implemented"
	return nil,

		// GetMetrics is a dummy function that always returns empty metrics for windows
		nil
}

func (*Cadvisor) GetMetrics() []pmetric.Metrics { _ = "STUB: not implemented"; return nil }

func (*Cadvisor) Shutdown() error { _ = "STUB: not implemented"; return nil }

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package kafkatopicsobserver // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/kafkatopicsobserver"
import (
	"context"
	"regexp"
	"sync"

	"github.com/twmb/franz-go/pkg/kadm"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/endpointswatcher"
)

var (
	_ extension.Extension = (*kafkaTopicsObserver)(nil)
	_ observer.Observable = (*kafkaTopicsObserver)(nil)
)

type kafkaTopicsObserver struct {
	*endpointswatcher.EndpointsWatcher
	logger *zap.Logger
	config *Config

	client *kadm.Client
}

func newObserver(logger *zap.Logger, config *Config) (extension.Extension, error) {
	_ = "STUB: not implemented"
	return *new(extension.Extension), nil
}

func (k *kafkaTopicsObserver) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

// Set metadata min age below the sync interval
// so each sync sees a fresh topic list.

func (k *kafkaTopicsObserver) Shutdown(context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

type kafkaTopicsEndpointsLister struct {
	o           *kafkaTopicsObserver
	topicRegexp *regexp.Regexp

	mu     sync.Mutex
	topics []string
}

func (k *kafkaTopicsEndpointsLister) ListEndpoints() []observer.Endpoint {
	_ = "STUB: not implemented"
	return nil
}

// Use the previously cached list of topics.

// Cache the new list of topics.

func (k *kafkaTopicsEndpointsLister) listMatchingTopics(ctx context.Context) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

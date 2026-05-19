// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package pull // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/k8sinventory/pull"

import (
	"context"
	"sync"
	"time"

	"go.uber.org/zap"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/client-go/dynamic"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/k8sinventory"
)

type Config struct {
	k8sinventory.Config
	Interval time.Duration
}

type Observer struct {
	config Config

	client dynamic.Interface
	logger *zap.Logger

	handlePullObjectsFunc func(objects *unstructured.UnstructuredList)
}

func New(client dynamic.Interface, config Config, logger *zap.Logger, handlePullObjectsFunc func(objects *unstructured.UnstructuredList)) (*Observer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (o *Observer) Start(ctx context.Context, wg *sync.WaitGroup) chan struct{} {
	_ = "STUB: not implemented"
	return nil
}

func (o *Observer) startPull(ctx context.Context, resource dynamic.ResourceInterface, stopperChan chan struct{}, wg *sync.WaitGroup) {
	_ = "STUB: not implemented"
	return
}

// Start ticking immediately.
// Ref: https://stackoverflow.com/questions/32705582/how-to-get-time-tick-to-tick-immediately
func newTicker(ctx context.Context, repeat time.Duration) *time.Ticker {
	_ = "STUB: not implemented"
	return nil
}

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package kube // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/k8sattributesprocessor/internal/kube"

import (
	"context"
	"sync"
	"time"

	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes"
	clientmeta "k8s.io/client-go/metadata"
	"k8s.io/client-go/tools/cache"
)

type FakeInformer struct {
	*FakeController

	namespace     string
	labelSelector labels.Selector
	fieldSelector fields.Selector
}

func NewFakeInformer(
	_ kubernetes.Interface,
	namespace string,
	labelSelector labels.Selector,
	fieldSelector fields.Selector,
) cache.SharedInformer {
	_ = "STUB: not implemented"
	return *new(cache.SharedInformer)
}

func (f *FakeInformer) AddEventHandler(handler cache.ResourceEventHandler) (cache.ResourceEventHandlerRegistration, error) {
	_ = "STUB: not implemented"
	return *new(cache.ResourceEventHandlerRegistration), nil
}

func (f *FakeInformer) AddEventHandlerWithResyncPeriod(_ cache.ResourceEventHandler, _ time.Duration) (cache.ResourceEventHandlerRegistration, error) {
	_ = "STUB: not implemented"
	return *new(cache.ResourceEventHandlerRegistration), nil
}

func (f *FakeInformer) AddEventHandlerWithOptions(cache.ResourceEventHandler, cache.HandlerOptions) (cache.ResourceEventHandlerRegistration, error) {
	_ = "STUB: not implemented"
	return *new(cache.ResourceEventHandlerRegistration), nil
}

func (*FakeInformer) RemoveEventHandler(cache.ResourceEventHandlerRegistration) error {
	_ = "STUB: not implemented"
	return nil
}

func (*FakeInformer) IsStopped() bool { _ = "STUB: not implemented"; return false }

func (*FakeInformer) SetTransform(cache.TransformFunc) error { _ = "STUB: not implemented"; return nil }

func (*FakeInformer) GetStore() cache.Store { _ = "STUB: not implemented"; return *new(cache.Store) }

func (f *FakeInformer) GetController() cache.Controller {
	_ = "STUB: not implemented"
	return *new(cache.Controller)
}

type FakeNamespaceInformer struct {
	*FakeController
}

func NewFakeNamespaceInformer(
	_ clientmeta.Interface,
) cache.SharedInformer {
	_ = "STUB: not implemented"
	return *new(cache.SharedInformer)
}

func (*FakeNamespaceInformer) AddEventHandler(cache.ResourceEventHandler) {
	_ = "STUB: not implemented"
	return
}

func (*FakeNamespaceInformer) AddEventHandlerWithResyncPeriod(cache.ResourceEventHandler, time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (*FakeNamespaceInformer) GetStore() cache.Store {
	_ = "STUB: not implemented"
	return *new(cache.Store)
}

func (f *FakeNamespaceInformer) GetController() cache.Controller {
	_ = "STUB: not implemented"
	return *new(cache.Controller)
}

type FakeReplicaSetInformer struct {
	*FakeController
}

func NewFakeReplicaSetInformer(
	_ clientmeta.Interface,
	_ string,
) cache.SharedInformer {
	_ = "STUB: not implemented"
	return *new(cache.SharedInformer)
}

func (*FakeReplicaSetInformer) AddEventHandler(cache.ResourceEventHandler) {
	_ = "STUB: not implemented"
	return
}

func (*FakeReplicaSetInformer) AddEventHandlerWithResyncPeriod(cache.ResourceEventHandler, time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (*FakeReplicaSetInformer) SetTransform(cache.TransformFunc) error {
	_ = "STUB: not implemented"
	return nil
}

func (*FakeReplicaSetInformer) GetStore() cache.Store {
	_ = "STUB: not implemented"
	return *new(cache.Store)
}

func (f *FakeReplicaSetInformer) GetController() cache.Controller {
	_ = "STUB: not implemented"
	return *new(cache.Controller)
}

type FakeController struct {
	sync.Mutex
	stopped bool
}

func (*FakeController) HasSynced() bool { _ = "STUB: not implemented"; return false }

func (c *FakeController) Run(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

func (c *FakeController) RunWithContext(ctx context.Context) { _ = "STUB: not implemented"; return }

func (c *FakeController) HasStopped() bool { _ = "STUB: not implemented"; return false }

func (*FakeController) LastSyncResourceVersion() string { _ = "STUB: not implemented"; return "" }

func (*FakeInformer) SetWatchErrorHandler(cache.WatchErrorHandler) error {
	_ = "STUB: not implemented"
	return nil
}

func (*FakeInformer) SetWatchErrorHandlerWithContext(cache.WatchErrorHandlerWithContext) error {
	_ = "STUB: not implemented"
	return nil
}

type NoOpInformer struct {
	*NoOpController
}

func NewNoOpInformer(
	_ clientmeta.Interface,
) cache.SharedInformer {
	_ = "STUB: not implemented"
	return *new(cache.SharedInformer)
}

func (f *NoOpInformer) AddEventHandler(handler cache.ResourceEventHandler) (cache.ResourceEventHandlerRegistration, error) {
	_ = "STUB: not implemented"
	return *new(cache.ResourceEventHandlerRegistration), nil
}

func (f *NoOpInformer) AddEventHandlerWithResyncPeriod(cache.ResourceEventHandler, time.Duration) (cache.ResourceEventHandlerRegistration, error) {
	_ = "STUB: not implemented"
	return *new(cache.ResourceEventHandlerRegistration), nil
}

func (f *NoOpInformer) AddEventHandlerWithOptions(cache.ResourceEventHandler, cache.HandlerOptions) (cache.ResourceEventHandlerRegistration, error) {
	_ = "STUB: not implemented"
	return *new(cache.ResourceEventHandlerRegistration), nil
}

func (*NoOpInformer) RemoveEventHandler(cache.ResourceEventHandlerRegistration) error {
	_ = "STUB: not implemented"
	return nil
}

func (*NoOpInformer) SetTransform(cache.TransformFunc) error { _ = "STUB: not implemented"; return nil }

func (*NoOpInformer) GetStore() cache.Store { _ = "STUB: not implemented"; return *new(cache.Store) }

func (f *NoOpInformer) GetController() cache.Controller {
	_ = "STUB: not implemented"
	return *new(cache.Controller)
}

type NoOpController struct {
	hasStopped bool
}

func (c *NoOpController) Run(stopCh <-chan struct{}) { _ = "STUB: not implemented"; return }

func (c *NoOpController) RunWithContext(ctx context.Context) { _ = "STUB: not implemented"; return }

func (c *NoOpController) IsStopped() bool { _ = "STUB: not implemented"; return false }

func (*NoOpController) HasSynced() bool { _ = "STUB: not implemented"; return false }

func (*NoOpController) LastSyncResourceVersion() string { _ = "STUB: not implemented"; return "" }

func (*NoOpController) SetWatchErrorHandler(cache.WatchErrorHandler) error {
	_ = "STUB: not implemented"
	return nil
}

func (*NoOpController) SetWatchErrorHandlerWithContext(cache.WatchErrorHandlerWithContext) error {
	_ = "STUB: not implemented"
	return nil
}

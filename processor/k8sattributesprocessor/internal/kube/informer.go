// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package kube // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/k8sattributesprocessor/internal/kube"

import (
	"time"

	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/metadata"
	"k8s.io/client-go/tools/cache"
)

const kubeSystemNamespace = "kube-system"

// InformerProvider defines a function type that returns a new SharedInformer. It is used to
// allow passing custom shared informers to the watch client.
type InformerProvider func(
	client kubernetes.Interface,
	namespace string,
	labelSelector labels.Selector,
	fieldSelector fields.Selector,
) cache.SharedInformer

// InformerProviderNamespace defines a function type that returns a new SharedInformer. It is used to
// allow passing custom shared informers to the watch client for fetching namespace objects.
type InformerProviderNamespace func(
	client metadata.Interface,
) cache.SharedInformer

// InformerProviderWorkload defines a function type that returns a new SharedInformer. It is used to
// allow passing custom shared informers to the watch client.
// It's used for high-level workloads such as ReplicaSets, Deployments, DaemonSets, StatefulSets or Jobs
type InformerProviderWorkload func(
	client metadata.Interface,
	namespace string,
) cache.SharedInformer

func newSharedInformer(
	client kubernetes.Interface,
	namespace string,
	ls labels.Selector,
	fs fields.Selector,
	watchSyncPeriod time.Duration,
) cache.SharedInformer {
	_ = "STUB: not implemented"
	return *new(cache.SharedInformer)
}

func informerListFuncWithSelectors(client kubernetes.Interface, namespace string, ls labels.Selector, fs fields.Selector) cache.ListWithContextFunc {
	_ = "STUB: not implemented"
	return *new(cache.ListWithContextFunc)
}

func informerWatchFuncWithSelectors(client kubernetes.Interface, namespace string, ls labels.Selector, fs fields.Selector) cache.WatchFuncWithContext {
	_ = "STUB: not implemented"
	return *new(cache.WatchFuncWithContext)
}

// newKubeSystemSharedInformer watches only kube-system namespace
func newKubeSystemSharedInformer(
	client metadata.Interface,
	watchSyncPeriod time.Duration,
) cache.SharedInformer {
	_ = "STUB: not implemented"
	return *new(cache.SharedInformer)
}

func newNamespaceSharedInformer(
	client metadata.Interface,
	watchSyncPeriod time.Duration,
) cache.SharedInformer {
	_ = "STUB: not implemented"
	return *new(cache.SharedInformer)
}

func newReplicaSetSharedInformer(client metadata.Interface, namespace string, watchSyncPeriod time.Duration) cache.SharedInformer {
	_ = "STUB: not implemented"
	return *new(cache.SharedInformer)
}

func newDeploymentSharedInformer(client metadata.Interface, namespace string, watchSyncPeriod time.Duration) cache.SharedInformer {
	_ = "STUB: not implemented"
	return *new(cache.SharedInformer)
}

func newStatefulSetSharedInformer(client metadata.Interface, namespace string, watchSyncPeriod time.Duration) cache.SharedInformer {
	_ = "STUB: not implemented"
	return *new(cache.SharedInformer)
}

func newDaemonSetSharedInformer(client metadata.Interface, namespace string, watchSyncPeriod time.Duration) cache.SharedInformer {
	_ = "STUB: not implemented"
	return *new(cache.SharedInformer)
}

func newJobSharedInformer(client metadata.Interface, namespace string, watchSyncPeriod time.Duration) cache.SharedInformer {
	_ = "STUB: not implemented"
	return *new(cache.SharedInformer)
}

func newNodeSharedInformer(client metadata.Interface, nodeName string, watchSyncPeriod time.Duration) cache.SharedInformer {
	_ = "STUB: not implemented"
	return *new(cache.SharedInformer)
}

func metadataListFunc(mc metadata.Interface, gvr schema.GroupVersionResource, namespace string) cache.ListWithContextFunc {
	_ = "STUB: not implemented"
	return *new(cache.ListWithContextFunc)
}

func metadataWatchFunc(mc metadata.Interface, gvr schema.GroupVersionResource, namespace string) cache.WatchFuncWithContext {
	_ = "STUB: not implemented"
	return *new(cache.WatchFuncWithContext)
}

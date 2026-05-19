// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package k8sclusterreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8sclusterreceiver"

import (
	"context"
	"reflect"
	"sync/atomic"
	"time"

	quotaclientset "github.com/openshift/client-go/quota/clientset/versioned"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
	"go.uber.org/zap"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/k8sconfig"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/experimentalmetricmetadata"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8sclusterreceiver/internal/metadata"
)

type sharedInformer interface {
	Start(<-chan struct{})
	WaitForCacheSync(<-chan struct{}) map[reflect.Type]bool
}

type resourceWatcher struct {
	client              kubernetes.Interface
	osQuotaClient       quotaclientset.Interface
	informerFactories   []sharedInformer
	metadataStore       *metadata.Store
	logger              *zap.Logger
	metadataConsumers   []metadataConsumer
	initialTimeout      time.Duration
	initialSyncDone     *atomic.Bool
	initialSyncTimedOut *atomic.Bool
	config              *Config
	entityLogConsumer   consumer.Logs

	// For mocking.
	makeClient               func(apiConf k8sconfig.APIConfig) (kubernetes.Interface, error)
	makeOpenShiftQuotaClient func(apiConf k8sconfig.APIConfig) (quotaclientset.Interface, error)
}

type metadataConsumer func(metadata []*experimentalmetricmetadata.MetadataUpdate) error

// newResourceWatcher creates a Kubernetes resource watcher.
func newResourceWatcher(set receiver.Settings, cfg *Config, metadataStore *metadata.Store) *resourceWatcher {
	_ = "STUB: not implemented"
	return nil
}

func (rw *resourceWatcher) initialize() error { _ = "STUB: not implemented"; return nil }

// shouldWatchResourceForMetadataOnly returns true if a resource should be watched even though
// its metrics are disabled. This is the case when metadata exporters are configured or
// entity events are enabled (entityLogConsumer is set).
func (rw *resourceWatcher) shouldWatchResourceForMetadataOnly() bool {
	_ = "STUB: not implemented"
	return false
}

func (rw *resourceWatcher) prepareSharedInformerFactory() error {
	_ = "STUB: not implemented"
	return nil
}

// Map of supported group version kinds by name of a kind.
// If none of the group versions are supported by k8s server for a specific kind,
// informer for that kind won't be set and a warning message is thrown.
// This map should be kept in sync with what can be provided by the supported k8s server versions.

// Only watch EndpointSlice if any service endpoint metrics are enabled

// Resources with all metrics disabled by default.
// Only watch these if any of their metrics are explicitly enabled or if metadata/entity destinations are configured.

// shouldWatchEndpointSlice returns true if the service endpoint count metric is enabled
func (rw *resourceWatcher) shouldWatchEndpointSlice() bool { _ = "STUB: not implemented"; return false }

// shouldWatchPersistentVolume returns true if any PV metric is enabled or metadata/entity destinations are configured.
func (rw *resourceWatcher) shouldWatchPersistentVolume() bool {
	_ = "STUB: not implemented"
	return false
}

// shouldWatchPersistentVolumeClaim returns true if any PVC metric is enabled or metadata/entity destinations are configured.
func (rw *resourceWatcher) shouldWatchPersistentVolumeClaim() bool {
	_ = "STUB: not implemented"
	return false
}

// getInformerFactories creates the informer factories which are used to set up the informers for the
// resources that should be observed. The informer factories are returned as a map[string]informer.SharedInformerFactory,
// where the map keys represent the namespace that should be observed. If the factory is created for the whole cluster,
// the factory is stored under an empty string
func (rw *resourceWatcher) getInformerFactories() map[string]informers.SharedInformerFactory {
	_ = "STUB: not implemented"
	return nil
}

// if no namespace is provided, the informer observes the whole cluster, and is stored under
// the key "<cluster-wide-informer-key>"

func (rw *resourceWatcher) isKindSupported(gvk schema.GroupVersionKind) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// if the discovery endpoint isn't present, assume group version is not supported

// setupInformerForKind creates the informers for the given GVKs, based on the provided informer factories.
// The factories are provided as a map[string]informers.SharedInformerFactory where the map keys represent the namespace
// of the informer factory. For cluster wide informers, an empty string is used as a key
func (rw *resourceWatcher) setupInformerForKind(kind schema.GroupVersionKind, factories map[string]informers.SharedInformerFactory) {
	_ = "STUB: not implemented"
	return
}

// if no namespace is provided, the cluster wide informer factory, which is stored under the key "" is used to create the informer

// if no namespace is provided, the cluster wide informer factory, which is stored under the key "" is used to create the informer

// PersistentVolumes are cluster-scoped, so only use cluster-wide informer

// startWatchingResources starts up all informers.
func (rw *resourceWatcher) startWatchingResources(ctx context.Context, inf sharedInformer) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// Start off individual informers in the factory.

// Ensure cache is synced with initial state, once informers are started up.
// Note that the event handler can start receiving events as soon as the informers
// are started. So it's required to ensure that the receiver does not start
// collecting data before the cache sync since all data may not be available.
// This method will block either till the timeout set on the context, until
// the initial sync is complete or the parent context is cancelled.

// setupInformer adds event handlers to informers and setups a metadataStore.
func (rw *resourceWatcher) setupInformer(gvk schema.GroupVersionKind, namespace string, informer cache.SharedIndexInformer) {
	_ = "STUB: not implemented"
	return
}

func (rw *resourceWatcher) onAdd(obj any) { _ = "STUB: not implemented"; return }

// Sync metadata only if there's at least one destination for it to sent.

func (rw *resourceWatcher) hasDestination() bool { _ = "STUB: not implemented"; return false }

func (rw *resourceWatcher) onUpdate(oldObj, newObj any) { _ = "STUB: not implemented"; return }

// Sync metadata only if there's at least one destination for it to sent.

func (rw *resourceWatcher) onDelete(oldObj any) { _ = "STUB: not implemented"; return }

// Sync metadata only if there's at least one destination for it to sent.

// objMetadata returns the metadata for the given object.
func (rw *resourceWatcher) objMetadata(obj any) map[experimentalmetricmetadata.ResourceID]*metadata.KubernetesMetadata {
	_ = "STUB: not implemented"
	return nil
}

func (rw *resourceWatcher) waitForInitialInformerSync() { _ = "STUB: not implemented"; return }

// Wait till initial sync is complete or timeout.

func (rw *resourceWatcher) setupMetadataExporters(
	exporters map[component.ID]component.Component,
	metadataExportersFromConfig []string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func validateMetadataExporters(metadataExporters map[string]bool, exporters map[component.ID]component.Component) error {
	_ = "STUB: not implemented"
	return nil
}

func (rw *resourceWatcher) syncMetadataUpdate(oldMetadata, newMetadata map[experimentalmetricmetadata.ResourceID]*metadata.KubernetesMetadata) {
	_ = "STUB: not implemented"
	return
}

// Represent metadata update as entity events.

// Convert entity events to log representation.

// Note: receiver contract says that we need to retry sending if the
// returned error is not Permanent. However, we are not doing it here.
// Instead, we rely on the fact the metadata is collected periodically
// and the entity events will be delivered on the next cycle. This is
// fine because we deliver cumulative entity state.
// This allows us to avoid stressing the Collector or its destination
// unnecessarily (typically non-Permanent errors happen in stressed conditions).
// The periodic collection will be implemented later, see
// https://github.com/open-telemetry/opentelemetry-collector-contrib/issues/24413

// stringSliceToMap converts a slice of strings into a map with keys from the slice
func stringSliceToMap(strings []string) map[string]bool { _ = "STUB: not implemented"; return nil }

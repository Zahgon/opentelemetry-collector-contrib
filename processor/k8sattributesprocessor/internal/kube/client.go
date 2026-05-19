// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package kube // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/k8sattributesprocessor/internal/kube"

import (
	"errors"
	"regexp"
	"sync"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"
	api_v1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/selection"
	"k8s.io/client-go/kubernetes"
	clientmeta "k8s.io/client-go/metadata"
	"k8s.io/client-go/tools/cache"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/k8sconfig"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/k8sattributesprocessor/internal/metadata"
)

// WatchClient is the main interface provided by this package to a kubernetes cluster.
type WatchClient struct {
	m                      sync.RWMutex
	deleteMut              sync.Mutex
	logger                 *zap.Logger
	kc                     kubernetes.Interface
	mc                     clientmeta.Interface
	informer               cache.SharedInformer
	namespaceInformer      cache.SharedInformer
	nodeInformer           cache.SharedInformer
	deploymentInformer     cache.SharedInformer
	statefulsetInformer    cache.SharedInformer
	daemonsetInformer      cache.SharedInformer
	jobInformer            cache.SharedInformer
	replicasetInformer     cache.SharedInformer
	cronJobRegex           *regexp.Regexp
	deleteQueue            []deleteRequest
	stopCh                 chan struct{}
	waitForMetadata        bool
	waitForMetadataTimeout time.Duration
	watchSyncPeriod        time.Duration

	// A map containing Pod related data, used to associate them with resources.
	// Key can be either an IP address or Pod UID
	Pods         map[PodIdentifier]*Pod
	Rules        ExtractionRules
	Filters      Filters
	Associations []Association
	Exclude      Excludes

	// A map containing Namespace related data, used to associate them with resources.
	// Key is namespace name
	Namespaces map[string]*Namespace

	// A map containing Node related data, used to associate them with resources.
	// Key is node name
	Nodes map[string]*Node

	// A map containing Deployment related data, used to associate them with resources.
	// Key is deployment uid
	Deployments map[string]*Deployment

	// A map containing StatefulSet related data, used to associate them with resources.
	// Key is statefulset uid
	StatefulSets map[string]*StatefulSet

	// A map containing DaemonSet related data, used to associate them with resources.
	// Key is daemonset uid
	DaemonSets map[string]*DaemonSet

	// A map containing job related data, used to associate them with resources.
	// Key is job uid
	Jobs map[string]*Job

	// A map containing ReplicaSets related data, used to associate them with resources.
	// Key is replicaset uid
	ReplicaSets map[string]*ReplicaSet

	telemetryBuilder *metadata.TelemetryBuilder
}

// Extract CronJob name from the Job name. Job name is created using
// format: [cronjob-name]-[time-hash-int]
// time-hash-int is the unix timestamp in minutes of the job creation time
// 8 digits will last until 2160, we are safe for a while
var cronJobRegex = regexp.MustCompile(`^(.*)-(\d{8})$`)

// cronJobSuffixTimeSkewMinutes is the maximum allowed difference
// in minutes between pod creation time and the timestamp suffix
// in the job name (created by cronjob)
// If the suffix falls outside of the range, it is assumed that it's not a cronjob
// (i.e. the suffix is human generated, like a date 20260407)
const cronJobSuffixTimeSkewMinutes int64 = 60 * 24 // 1 day

// podTemplateHashLabel contains the label that K8s Deployment
// controller adds to ReplicaSets to ensure that child ReplicaSets
// do not overlap. ReplicaSet name is constructed as [deployment-name]-[podTemplateHash]
// K8s doc ref: https://kubernetes.io/docs/concepts/workloads/controllers/deployment/#pod-template-hash-label
// K8s code ref: https://github.com/kubernetes/kubernetes/blob/b31119d205a839aab40b2d819a58d4fabacd9b47/pkg/controller/deployment/sync.go#L207
const podTemplateHashLabel = "pod-template-hash"

var errCannotRetrieveImage = errors.New("cannot retrieve image name")

type InformersFactoryList struct {
	newInformer           InformerProvider
	newNamespaceInformer  InformerProviderNamespace
	newReplicaSetInformer InformerProviderWorkload
}

// New initializes a new k8s Client.
func New(
	set component.TelemetrySettings,
	apiCfg k8sconfig.APIConfig,
	rules ExtractionRules,
	filters Filters,
	associations []Association,
	exclude Excludes,
	newClientSet APIClientsetProvider,
	informersFactory InformersFactoryList,
	waitForMetadata bool,
	waitForMetadataTimeout time.Duration,
	watchSyncPeriod time.Duration,
) (Client, error) {
	_ = "STUB: not implemented"
	return *new(Client), nil
}

// if rules to extract metadata from namespace is configured use namespace shared informer containing
// all namespaces including kube-system which contains cluster uid information (kube-system-uid)

// use kube-system shared informer to only watch kube-system namespace
// reducing overhead of watching all the namespaces

// means this is a cache.DeletedFinalStateUnknown, in which case we do nothing

// Enable the ReplicaSet informer when any of the following applies:
//   1) DeploymentName is enabled and DeploymentNameFromReplicaSet flag is false
//   2) DeploymentUID is enabled
//   3) Deployment labels or annotations are configured for extraction — those fields live on the Deployment
//      resource, so we must associate Pods with Deployments through ReplicaSets first.

// Start registers pod event handlers and starts watching the kubernetes cluster for pod changes.
func (c *WatchClient) Start() error {
	_ = "STUB: not implemented"
	// Start the delete loop for cleaning up old pods from cache
	return nil
}

// start the replicaSet informer first, as the replica sets need to be
// present at the time the pods are handled, to correctly establish the connection between pods and deployments

// start the podInformer with the prerequisite of the other informers to be finished first

// Wait for the Pod informer to be completed.
// The other informers will already be finished at this point, as the pod informer
// waits for them be finished before it can run

// Stop signals the k8s watcher/informer to stop watching for new events.
func (c *WatchClient) Stop() { _ = "STUB: not implemented"; return }

func (c *WatchClient) handlePodAdd(obj any) { _ = "STUB: not implemented"; return }

func (c *WatchClient) handlePodUpdate(_, newPod any) { _ = "STUB: not implemented"; return }

// TODO: update or remove based on whether container is ready/unready?.

func (c *WatchClient) handlePodDelete(obj any) { _ = "STUB: not implemented"; return }

func (c *WatchClient) handleNamespaceAdd(obj any) { _ = "STUB: not implemented"; return }

func (c *WatchClient) handleNamespaceUpdate(_, newNamespace any) { _ = "STUB: not implemented"; return }

func (c *WatchClient) handleNamespaceDelete(obj any) { _ = "STUB: not implemented"; return }

// When a namespace is deleted all the pods(and other k8s objects in that namespace) in that namespace are deleted before it.
// So we wont have any spans that might need namespace annotations and labels.
// Thats why we dont need an implementation for deleteQueue and gracePeriod for namespaces.

func (c *WatchClient) handleNodeAdd(obj any) { _ = "STUB: not implemented"; return }

func (c *WatchClient) handleNodeUpdate(_, newNode any) { _ = "STUB: not implemented"; return }

func (c *WatchClient) handleNodeDelete(obj any) { _ = "STUB: not implemented"; return }

func (c *WatchClient) handleDeploymentAdd(obj any) { _ = "STUB: not implemented"; return }

func (c *WatchClient) handleDeploymentUpdate(_, newDeployment any) {
	_ = "STUB: not implemented"
	return
}

func (c *WatchClient) handleDeploymentDelete(obj any) { _ = "STUB: not implemented"; return }

func (c *WatchClient) handleStatefulSetAdd(obj any) { _ = "STUB: not implemented"; return }

func (c *WatchClient) handleStatefulSetUpdate(_, newStatefulSet any) {
	_ = "STUB: not implemented"
	return
}

func (c *WatchClient) handleStatefulSetDelete(obj any) { _ = "STUB: not implemented"; return }

func (c *WatchClient) handleDaemonSetAdd(obj any) { _ = "STUB: not implemented"; return }

func (c *WatchClient) handleDaemonSetUpdate(_, newDaemonSet any) { _ = "STUB: not implemented"; return }

func (c *WatchClient) handleDaemonSetDelete(obj any) { _ = "STUB: not implemented"; return }

func (c *WatchClient) handleJobAdd(obj any) { _ = "STUB: not implemented"; return }

func (c *WatchClient) handleJobUpdate(_, newJob any) { _ = "STUB: not implemented"; return }

func (c *WatchClient) handleJobDelete(obj any) { _ = "STUB: not implemented"; return }

func (c *WatchClient) deleteLoop(interval, gracePeriod time.Duration) {
	_ = "STUB: not implemented"
	// This loop runs after N seconds and deletes pods from cache.
	// It iterates over the delete queue and deletes all that aren't
	// in the grace period anymore.
	return
}

func (c *WatchClient) deleteLoopProcessing(gracePeriod time.Duration) {
	_ = "STUB: not implemented"
	return
}

// Sanity check: make sure we are deleting the same pod
// and the underlying state (ip<>pod mapping) has not changed.

func (c *WatchClient) compactPodMap() { _ = "STUB: not implemented"; return }

// GetPod takes an IP address or Pod UID and returns the pod the identifier is associated with.
func (c *WatchClient) GetPod(identifier PodIdentifier) (*Pod, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// GetNamespace takes a namespace and returns the namespace object the namespace is associated with.
func (c *WatchClient) GetNamespace(namespace string) (*Namespace, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// GetNode takes a node name and returns the node object the node name is associated with.
func (c *WatchClient) GetNode(nodeName string) (*Node, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (c *WatchClient) GetDeployment(deploymentUID string) (*Deployment, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (c *WatchClient) GetReplicaSet(uid string) (*ReplicaSet, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (c *WatchClient) GetStatefulSet(statefulSetUID string) (*StatefulSet, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (c *WatchClient) GetDaemonSet(daemonSetUID string) (*DaemonSet, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (c *WatchClient) GetJob(jobUID string) (*Job, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (c *WatchClient) extractPodAttributes(pod *api_v1.Pod) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

// Attempt to use the ReplicaSet informer when it is running (e.g. DeploymentUID,
// deployment_name_from_replicaset: false, or from: deployment label/annotation extraction).
// Otherwise fall back to heuristics using the ReplicaSet name and pod-template-hash label.

// Will be blank if not a ReplicaSet managed by a Deployment

// deployment name wins over replicaset name

// Attempt to use the Job informer when it is running (e.g. k8s.cronjob.uid or from: job
// label/annotation extraction). Otherwise fall back to heuristics on the Job name suffix.

// cronjob name wins over job name

// app.kubernetes.io/instance has a higher precedence than app.kubernetes.io/name

// Annotations are processed after labels so that OTel annotations
// (e.g. resource.opentelemetry.io/service.name) take precedence over
// well-known Kubernetes labels (e.g. app.kubernetes.io/name).

func copyLabel(pod *api_v1.Pod, tags map[string]string, labelKey string, key attribute.Key) {
	_ = "STUB: not implemented"
	return
}

// This function removes all data from the Pod except what is required by extraction rules and pod association
func removeUnnecessaryPodData(pod *api_v1.Pod, rules ExtractionRules) *api_v1.Pod {
	_ = "STUB: not implemented"
	// name, namespace, uid, start time and ip are needed for identifying Pods
	// there's room to optimize this further, it's kept this way for simplicity
	return nil
}

// Creation time is required for k8s.pod.start_time and for CronJob name heuristics
// (Job name suffix vs pod creation) when k8s.pod.start_time is not extracted.

// we always need the name, it's used for identification

// parseServiceVersionFromImage parses the service version for differently-formatted image names
// according to https://github.com/open-telemetry/semantic-conventions/blob/main/docs/non-normative/k8s-attributes.md#how-serviceversion-should-be-calculated
func parseServiceVersionFromImage(image string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (c *WatchClient) extractPodContainersAttributes(pod *api_v1.Pod) PodContainers {
	_ = "STUB: not implemented"
	return *new(PodContainers)
}

//nolint:gocritic // appendAssign: append result not assigned to the same slice

// Legacy: container.image.tag (singular, string)

// Stable: container.image.tags (plural, array)

//nolint:gocritic // appendAssign: append result not assigned to the same slice

// Remove container runtime prefix

func (c *WatchClient) extractNamespaceAttributes(namespace *meta_v1.PartialObjectMetadata) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func (c *WatchClient) extractNodeAttributes(node *meta_v1.PartialObjectMetadata) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func (c *WatchClient) extractDeploymentAttributes(d *meta_v1.PartialObjectMetadata) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func (c *WatchClient) extractStatefulSetAttributes(d *meta_v1.PartialObjectMetadata) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func (c *WatchClient) extractDaemonSetAttributes(d *meta_v1.PartialObjectMetadata) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func (c *WatchClient) extractJobAttributes(d *meta_v1.PartialObjectMetadata) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func (c *WatchClient) podFromAPI(pod *api_v1.Pod) *Pod { _ = "STUB: not implemented"; return nil }

func getPodReplicaSetUID(pod *api_v1.Pod) string { _ = "STUB: not implemented"; return "" }

func getPodStatefulSetUID(pod *api_v1.Pod) string { _ = "STUB: not implemented"; return "" }

func getPodDaemonSetUID(pod *api_v1.Pod) string { _ = "STUB: not implemented"; return "" }

func getPodJobUID(pod *api_v1.Pod) string { _ = "STUB: not implemented"; return "" }

// getIdentifiersFromAssoc returns list of PodIdentifiers for given pod
func (c *WatchClient) getIdentifiersFromAssoc(pod *Pod) []PodIdentifier {
	_ = "STUB: not implemented"
	return nil
}

// If association configured to take IP address from connection

// Host network mode is not supported right now with IP based
// tagging as all pods in host network get same IP addresses.
// Such pods are very rare and usually are used to monitor or control
// host traffic (e.g, linkerd, flannel) instead of service business needs.

// k8s.pod.ip is set by passthrough mode

// At this point just an empty attr is added and we remember the position.
// Later this position in PodIdentifier will be filled with the actual
// value for container.ID.

// As there can be multiple container.IDs per pod,
// one PodIdentifier is added per container.ID.

// Ensure backward compatibility

// k8s.pod.ip is set by passthrough mode

func (c *WatchClient) addOrUpdatePod(pod *api_v1.Pod) { _ = "STUB: not implemented"; return }

// compare initial scheduled timestamp for existing pod and new pod with same identifier
// and only replace old pod if scheduled time of new pod is newer or equal.
// This should fix the case where scheduler has assigned the same attributes (like IP address)
// to a new pod but update event for the old pod came in later.

func (c *WatchClient) forgetPod(pod *api_v1.Pod) { _ = "STUB: not implemented"; return }

func (c *WatchClient) appendDeleteQueue(podID PodIdentifier, podUID string) {
	_ = "STUB: not implemented"
	return
}

func (c *WatchClient) shouldIgnorePod(pod *api_v1.Pod) bool {
	_ = "STUB: not implemented"
	// Check if user requested the pod to be ignored through annotations
	return false
}

// Check if user requested the pod to be ignored through configuration

var singleValueOperators = map[selection.Operator]int{
	selection.Equals:       1,
	selection.DoubleEquals: 1,
	selection.NotEquals:    1,
	selection.GreaterThan:  1,
	selection.LessThan:     1,
}

func selectorsFromFilters(filters Filters) (labels.Selector, fields.Selector, error) {
	_ = "STUB: not implemented"
	return *new(labels.Selector), *new(fields.Selector), nil
}

func (c *WatchClient) addOrUpdateNamespace(namespace *meta_v1.PartialObjectMetadata) {
	_ = "STUB: not implemented"
	return
}

func (c *WatchClient) extractNamespaceLabelsAnnotations() bool {
	_ = "STUB: not implemented"
	return false
}

func (c *WatchClient) extractDeploymentLabelsAnnotations() bool {
	_ = "STUB: not implemented"
	return false
}

func (c *WatchClient) extractStatefulSetLabelsAnnotations() bool {
	_ = "STUB: not implemented"
	return false
}

func (c *WatchClient) extractDaemonSetLabelsAnnotations() bool {
	_ = "STUB: not implemented"
	return false
}

func (c *WatchClient) extractJobLabelsAnnotations() bool { _ = "STUB: not implemented"; return false }

func (c *WatchClient) extractNodeLabelsAnnotations() bool { _ = "STUB: not implemented"; return false }

func (c *WatchClient) extractNodeUID() bool { _ = "STUB: not implemented"; return false }

func (c *WatchClient) addOrUpdateNode(node *meta_v1.PartialObjectMetadata) {
	_ = "STUB: not implemented"
	return
}

func (c *WatchClient) addOrUpdateDeployment(deployment *meta_v1.PartialObjectMetadata) {
	_ = "STUB: not implemented"
	return
}

func (c *WatchClient) addOrUpdateStatefulSet(statefulset *meta_v1.PartialObjectMetadata) {
	_ = "STUB: not implemented"
	return
}

func (c *WatchClient) addOrUpdateDaemonSet(daemonset *meta_v1.PartialObjectMetadata) {
	_ = "STUB: not implemented"
	return
}

func (c *WatchClient) addOrUpdateJob(job *meta_v1.PartialObjectMetadata) {
	_ = "STUB: not implemented"
	return
}

func needContainerAttributes(rules ExtractionRules) bool { _ = "STUB: not implemented"; return false }

func (c *WatchClient) handleReplicaSetAdd(obj any) { _ = "STUB: not implemented"; return }

func (c *WatchClient) handleReplicaSetUpdate(_, newRS any) { _ = "STUB: not implemented"; return }

func (c *WatchClient) handleReplicaSetDelete(obj any) { _ = "STUB: not implemented"; return }

// Unwrap DeletedFinalStateUnknown if present

func (c *WatchClient) addOrUpdateReplicaSet(replicaSet *meta_v1.PartialObjectMetadata) {
	_ = "STUB: not implemented"
	return
}

// runInformerWithDependencies starts the given informer. The second argument is a list of other informers that should complete
// before the informer is started. This is necessary e.g. for the pod informer which requires the replica set informer
// to be finished to correctly establish the connection to the replicaset/deployment it belongs to.
func (c *WatchClient) runInformerWithDependencies(informer cache.SharedInformer, dependencies []cache.InformerSynced) {
	_ = "STUB: not implemented"
	return
}

// TODO hard coding the timeout for now, check if we should make this configurable

// ignoreDeletedFinalStateUnknown returns the object wrapped in
// DeletedFinalStateUnknown. Useful in OnDelete resource event handlers that do
// not need the additional context.
func ignoreDeletedFinalStateUnknown(obj any) any { _ = "STUB: not implemented"; return *new(any) }

func automaticServiceInstanceID(pod *api_v1.Pod, containerName string) string {
	_ = "STUB: not implemented"
	return ""
}

// extractDeploymentNameFromReplicaSet attempts to extract deployment name from replicaset name
func extractDeploymentNameFromReplicaSet(replicasetName, podTemplateHash string) string {
	_ = "STUB: not implemented"
	// TemplateHash Empty means it's not a ReplicaSet managed by a Deployment
	return ""
}

// Remove the pod template hash suffix and the hyphen

// extractCronJobNameFromJobOwner returns the CronJob name for a Job owner reference using
// a heuristic that checks if the job name suffix is a valid timestamp closely matching the pod creation time.
// If it matches, then the suffix is assumed to be the job creation time, and the prefix is then the CronJob name.
func (c *WatchClient) extractCronJobNameFromJobOwner(ref meta_v1.OwnerReference, pod *api_v1.Pod) string {
	_ = "STUB: not implemented"
	// Regex checks specifically for 8 digits at the end (Kubernetes CronJob job suffix is a
	// time value in minutes).
	return ""
}

// assume not a cronjob

// assume not a cronjob

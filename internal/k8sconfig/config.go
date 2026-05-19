// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package k8sconfig // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/k8sconfig"

import (
	"context"
	"time"

	quotaclientset "github.com/openshift/client-go/quota/clientset/versioned"
	k8sruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/dynamic"
	k8s "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/metadata"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
)

func init() {
	k8sruntime.ReallyCrash = false
	k8sruntime.PanicHandlers = []func(context.Context, any){}
}

// AuthType describes the type of authentication to use for the K8s API
type AuthType string

// TODO: Add option for TLS once
// https://go.opentelemetry.io/collector/issues/933
// is addressed.
const (
	// AuthTypeNone means no auth is required
	AuthTypeNone AuthType = "none"
	// AuthTypeServiceAccount means to use the built-in service account that
	// K8s automatically provisions for each pod.
	AuthTypeServiceAccount AuthType = "serviceAccount"
	// AuthTypeKubeConfig uses local credentials like those used by kubectl.
	AuthTypeKubeConfig AuthType = "kubeConfig"
	// AuthTypeTLS indicates that client TLS auth is desired
	AuthTypeTLS AuthType = "tls"
)

const (
	// DefaultKubeAPIQPS is the default number of queries per second to the Kubernetes API.
	// Matches client-go's built-in default.
	DefaultKubeAPIQPS float32 = 5
	// DefaultKubeAPIBurst is the default burst limit for requests to the Kubernetes API.
	// Matches client-go's built-in default.
	DefaultKubeAPIBurst int = 10
)

var authTypes = map[AuthType]bool{
	AuthTypeNone:           true,
	AuthTypeServiceAccount: true,
	AuthTypeKubeConfig:     true,
	AuthTypeTLS:            true,
}

// APIConfig contains options relevant to connecting to the K8s API
type APIConfig struct {
	// How to authenticate to the K8s API server.  This can be one of `none`
	// (for no auth), `serviceAccount` (to use the standard service account
	// token provided to the agent pod), or `kubeConfig` to use credentials
	// from `~/.kube/config`.
	AuthType AuthType `mapstructure:"auth_type"`

	// When using auth_type `kubeConfig`, override the current context.
	Context string `mapstructure:"context"`

	// KubeAPIQPS is the maximum number of queries per second to the Kubernetes API.
	// Uses client-go's default (5) if unset. Increase if you see client-side throttling warnings.
	KubeAPIQPS float32 `mapstructure:"kube_api_qps"`

	// KubeAPIBurst is the maximum burst of requests to the Kubernetes API.
	// Uses client-go's default (10) if unset. Increase if you see client-side throttling warnings.
	KubeAPIBurst int `mapstructure:"kube_api_burst"`
}

// Validate validates the K8s API config
func (c APIConfig) Validate() error { _ = "STUB: not implemented"; return nil }

// CreateRestConfig creates an Kubernetes API config from user configuration.
func CreateRestConfig(apiConf APIConfig) (*rest.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// This should work for most clusters but other auth types can be added

// Don't use system proxy settings since the API is local to the
// cluster

// MakeClient can take configuration if needed for other types of auth
func MakeClient(apiConf APIConfig) (k8s.Interface, error) {
	_ = "STUB: not implemented"
	return *new(k8s.Interface), nil
}

// ClientBundle groups the two Kubernetes clients:
//
//   - K8s (typed client): kubernetes.Interface for full resource objects
//     (spec/status/metadata). Use when you need complete data or typed informers.
//
//   - Meta (metadata client): metadata.Interface for PartialObjectMetadata
//     (name/namespace/UID/labels/annotations/ownerRefs). Use for lightweight
//     list/watch when only metadata is needed (e.g., high-churn resources).
type ClientBundle struct {
	K8s  k8s.Interface
	Meta metadata.Interface
}

// MakeClientBundle builds both clients from a single RestConfig,
// ensuring shared auth/transport. In unit tests, inject a fake
// metadata client (metadata/fake) to avoid network calls, while
// typed resources can use kubernetes/fake.
func MakeClientBundle(apiConf APIConfig) (ClientBundle, error) {
	_ = "STUB: not implemented"
	return *new(ClientBundle), nil
}

// MakeDynamicClient can take configuration if needed for other types of auth
func MakeDynamicClient(apiConf APIConfig) (dynamic.Interface, error) {
	_ = "STUB: not implemented"
	return *new(dynamic.Interface), nil
}

// MakeOpenShiftQuotaClient can take configuration if needed for other types of auth
// and return an OpenShift quota API client
func MakeOpenShiftQuotaClient(apiConf APIConfig) (quotaclientset.Interface, error) {
	_ = "STUB: not implemented"
	return *new(quotaclientset.Interface), nil
}

func NewNodeSharedInformer(client k8s.Interface, nodeName string, watchSyncPeriod time.Duration) cache.SharedInformer {
	_ = "STUB: not implemented"
	return *new(cache.SharedInformer)
}

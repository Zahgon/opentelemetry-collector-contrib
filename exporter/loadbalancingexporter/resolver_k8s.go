// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package loadbalancingexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/loadbalancingexporter"

import (
	"context"
	"errors"
	"sync"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/loadbalancingexporter/internal/metadata"
)

var _ resolver = (*k8sResolver)(nil)

var (
	errNoSvc          = errors.New("no service specified to resolve the backends")
	errInvalidSvcFQDN = errors.New("invalid Kubernetes service FQDN")

	k8sResolverAttr           = attribute.String("resolver", "k8s")
	k8sResolverAttrSet        = attribute.NewSet(k8sResolverAttr)
	k8sResolverSuccessAttrSet = attribute.NewSet(k8sResolverAttr, attribute.Bool("success", true))
	k8sResolverFailureAttrSet = attribute.NewSet(k8sResolverAttr, attribute.Bool("success", false))
)

const (
	defaultListWatchTimeout = 1 * time.Minute
)

type k8sResolver struct {
	logger  *zap.Logger
	svcName string
	svcNs   string
	port    []int32

	handler        *handler
	once           *sync.Once
	epsListWatcher cache.ListerWatcher
	endpointsStore *sync.Map

	lwTimeout time.Duration

	endpoints         []string
	onChangeCallbacks []func([]string)
	returnNames       bool

	stopCh             chan struct{}
	updateLock         sync.RWMutex
	shutdownWg         sync.WaitGroup
	changeCallbackLock sync.RWMutex

	telemetry *metadata.TelemetryBuilder
}

func newK8sResolver(clt kubernetes.Interface,
	logger *zap.Logger,
	service string,
	ports []int32,
	timeout time.Duration,
	returnNames bool,
	tb *metadata.TelemetryBuilder,
) (*k8sResolver, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *k8sResolver) start(_ context.Context) error { _ = "STUB: not implemented"; return nil }

func (r *k8sResolver) shutdown(_ context.Context) error { _ = "STUB: not implemented"; return nil }

func newInClusterClient() (kubernetes.Interface, error) {
	_ = "STUB: not implemented"
	return *new(kubernetes.Interface), nil
}

func (r *k8sResolver) resolve(ctx context.Context) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// keep it always in the same order

// the list has changed!

// propagate the change

func (r *k8sResolver) onChange(f func([]string)) { _ = "STUB: not implemented"; return }

func (r *k8sResolver) Endpoints() []string { _ = "STUB: not implemented"; return nil }

const inClusterNamespacePath = "/var/run/secrets/kubernetes.io/serviceaccount/namespace"

func getInClusterNamespace() (string, error) {
	_ = "STUB: not implemented"
	// Check whether the namespace file exists.
	// If not, we are not running in cluster so can't guess the namespace.
	return "", nil
}

// Load the namespace file and return its content

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package k8snode // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/metadataproviders/k8snode"

import (
	"context"

	"k8s.io/client-go/kubernetes"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/k8sconfig"
)

type Provider interface {
	// NodeUID returns the K8S Node UID
	NodeUID(ctx context.Context) (string, error)
	// NodeName returns the current K8S Node Name
	NodeName(ctx context.Context) (string, error)
}

type k8snodeProvider struct {
	k8snodeClient kubernetes.Interface
	nodeName      string
}

func NewProvider(nodeName string, apiConf k8sconfig.APIConfig) (Provider, error) {
	_ = "STUB: not implemented"
	return *new(Provider), nil
}

func (k *k8snodeProvider) NodeUID(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (k *k8snodeProvider) NodeName(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package k8s // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/datadog/hostmetadata/internal/k8s"

import (
	"context"

	"go.uber.org/zap"
	k8s "k8s.io/client-go/kubernetes"
)

type nodeNameProvider interface {
	NodeName(context.Context) (string, error)
}

var _ nodeNameProvider = (*nodeNameProviderImpl)(nil)

type nodeNameProviderImpl struct {
	logger *zap.Logger
	client k8s.Interface
}

func (p *nodeNameProviderImpl) namespace() string { _ = "STUB: not implemented"; return "" }

func (p *nodeNameProviderImpl) NodeName(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "",

		// NOTE: The pod name may not match the OS hostname, e.g. if it has been modified
		// via the 'setHostnameAsFQDN' and 'hostname' fields in the pod spec.
		// The query below will error out in that case. See:
		// https://github.com/open-telemetry/opentelemetry-collector-contrib/issues/11033
		nil
}

// NOTE: If changing this, check if the RBAC rules on the docs or examples need updates.

var _ nodeNameProvider = (*nodeNameUnavailable)(nil)

type nodeNameUnavailable struct {
	err error
}

func (n *nodeNameUnavailable) NodeName(context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func newNodeNameProvider() nodeNameProvider {
	_ = "STUB: not implemented"
	return *new(nodeNameProvider)
}

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package kubeadm // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/kubeadm"

import (
	"context"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/processor"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/metadataproviders/kubeadm"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/kubeadm/internal/metadata"
)

const (
	TypeStr                    = "kubeadm"
	defaultConfigMapName       = "kubeadm-config"
	defaultKubeSystemNamespace = "kube-system"
)

var _ internal.Detector = (*detector)(nil)

type detector struct {
	provider kubeadm.Provider
	logger   *zap.Logger
	ra       *metadata.ResourceAttributesConfig
	rb       *metadata.ResourceBuilder
}

func NewDetector(set processor.Settings, dcfg internal.DetectorConfig) (internal.Detector, error) {
	_ = "STUB: not implemented"
	return *new(internal.Detector), nil
}

func (d *detector) Detect(ctx context.Context) (resource pcommon.Resource, schemaURL string, err error) {
	_ = "STUB: not implemented"
	return *new(pcommon.Resource), "", nil
}

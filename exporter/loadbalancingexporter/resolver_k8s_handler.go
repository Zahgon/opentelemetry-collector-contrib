// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package loadbalancingexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/loadbalancingexporter"

import (
	"context"
	"sync"

	"go.uber.org/zap"
	discoveryv1 "k8s.io/api/discovery/v1"
	"k8s.io/client-go/tools/cache"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/loadbalancingexporter/internal/metadata"
)

var _ cache.ResourceEventHandler = (*handler)(nil)

const (
	epMissingHostnamesMsg = "EndpointSlice object missing hostnames"
)

type handler struct {
	endpoints   *sync.Map
	callback    func(ctx context.Context) ([]string, error)
	logger      *zap.Logger
	telemetry   *metadata.TelemetryBuilder
	returnNames bool
}

func (h handler) OnAdd(obj any, _ bool) { _ = "STUB: not implemented"; return }

// unsupported

func (h handler) OnUpdate(oldObj, newObj any) { _ = "STUB: not implemented"; return }

// Iterate through old endpoints and remove those that are not in the new list.

// Iterate through new endpoints and add those that are not in the endpoints map already.

// unsupported

func (h handler) OnDelete(obj any) { _ = "STUB: not implemented"; return }

// unsupported

func convertToEndpoints(retNames bool, eps ...*discoveryv1.EndpointSlice) (bool, map[string]bool) {
	_ = "STUB: not implemented"
	return false, nil
}

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package loadbalancingexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/loadbalancingexporter"

import (
	"context"
	"errors"
	"sync"

	"go.opentelemetry.io/otel/attribute"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/loadbalancingexporter/internal/metadata"
)

var _ resolver = (*staticResolver)(nil)

var (
	errNoEndpoints               = errors.New("no endpoints specified for the static resolver")
	staticResolverAttr           = attribute.String("resolver", "static")
	staticResolverAttrSet        = attribute.NewSet(staticResolverAttr)
	staticResolverSuccessAttrSet = attribute.NewSet(staticResolverAttr, attribute.Bool("success", true))
)

type staticResolver struct {
	endpoints         []string
	onChangeCallbacks []func([]string)
	once              sync.Once // we trigger the onChange only once

	telemetry *metadata.TelemetryBuilder
}

func newStaticResolver(endpoints []string, tb *metadata.TelemetryBuilder) (*staticResolver, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// make sure we won't change the provided slice

// sort is a guarantee that the order of endpoints doesn't matter

func (r *staticResolver) start(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// right now, this can't fail

func (r *staticResolver) shutdown(context.Context) error { _ = "STUB: not implemented"; return nil }

func (r *staticResolver) resolve(ctx context.Context) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *staticResolver) onChange(f func([]string)) { _ = "STUB: not implemented"; return }

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package loadbalancingexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/loadbalancingexporter"

import (
	"context"
	"errors"
	"net"
	"sync"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/loadbalancingexporter/internal/metadata"
)

var _ resolver = (*dnsResolver)(nil)

const (
	defaultResInterval = 5 * time.Second
	defaultResTimeout  = time.Second
)

var (
	errNoHostname = errors.New("no hostname specified to resolve the backends")

	dnsResolverAttr           = attribute.String("resolver", "dns")
	dnsResolverAttrSet        = attribute.NewSet(dnsResolverAttr)
	dnsResolverSuccessAttrSet = attribute.NewSet(dnsResolverAttr, attribute.Bool("success", true))
	dnsResolverFailureAttrSet = attribute.NewSet(dnsResolverAttr, attribute.Bool("success", false))
)

type dnsResolver struct {
	logger *zap.Logger

	hostname    string
	port        string
	resolver    netResolver
	resInterval time.Duration
	resTimeout  time.Duration

	endpoints         []string
	onChangeCallbacks []func([]string)

	stopCh             chan struct{}
	updateLock         sync.Mutex
	shutdownWg         sync.WaitGroup
	changeCallbackLock sync.RWMutex
	telemetry          *metadata.TelemetryBuilder
}

type netResolver interface {
	LookupIPAddr(ctx context.Context, host string) ([]net.IPAddr, error)
}

func newDNSResolver(
	logger *zap.Logger,
	hostname string,
	port string,
	interval time.Duration,
	timeout time.Duration,
	tb *metadata.TelemetryBuilder,
) (*dnsResolver, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *dnsResolver) start(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (r *dnsResolver) shutdown(_ context.Context) error { _ = "STUB: not implemented"; return nil }

func (r *dnsResolver) periodicallyResolve() { _ = "STUB: not implemented"; return }

func (r *dnsResolver) resolve(ctx context.Context) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// it's an IPv6 address

// if a port is specified in the configuration, add it

// keep it always in the same order

// the list has changed!

// propagate the change

func (r *dnsResolver) onChange(f func([]string)) { _ = "STUB: not implemented"; return }

func equalStringSlice(source, candidate []string) bool { _ = "STUB: not implemented"; return false }

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package loadbalancingexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/loadbalancingexporter"

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/servicediscovery"
	"github.com/aws/aws-sdk-go-v2/service/servicediscovery/types"
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/loadbalancingexporter/internal/metadata"
)

const (
	defaultAwsResInterval = 30 * time.Second
	defaultAwsResTimeout  = 5 * time.Second
)

var (
	errNoNamespace   = errors.New("no Cloud Map namespace specified to resolve the backends")
	errNoServiceName = errors.New("no Cloud Map service_name specified to resolve the backends")

	awsResolverAttr           = attribute.String("resolver", "aws")
	awsResolverAttrSet        = attribute.NewSet(awsResolverAttr)
	awsResolverSuccessAttrSet = attribute.NewSet(awsResolverAttr, attribute.Bool("success", true))
	awsResolverFailureAttrSet = attribute.NewSet(awsResolverAttr, attribute.Bool("success", false))
)

func createDiscoveryFunction(client *servicediscovery.Client) func(params *servicediscovery.DiscoverInstancesInput) (*servicediscovery.DiscoverInstancesOutput, error) {
	_ = "STUB: not implemented"
	return nil
}

type cloudMapResolver struct {
	logger *zap.Logger

	namespaceName *string
	serviceName   *string
	port          *uint16
	healthStatus  *types.HealthStatusFilter
	resInterval   time.Duration
	resTimeout    time.Duration
	ownerAccount  *string

	endpoints         []string
	onChangeCallbacks []func([]string)

	stopCh             chan struct{}
	updateLock         sync.Mutex
	shutdownWg         sync.WaitGroup
	changeCallbackLock sync.RWMutex
	discoveryFn        func(params *servicediscovery.DiscoverInstancesInput) (*servicediscovery.DiscoverInstancesOutput, error)
	telemetry          *metadata.TelemetryBuilder
}

func newCloudMapResolver(
	logger *zap.Logger,
	namespaceName *string,
	serviceName *string,
	port *uint16,
	healthStatus *types.HealthStatusFilter,
	interval time.Duration,
	timeout time.Duration,
	ownerAccount *string,
	tb *metadata.TelemetryBuilder,
) (*cloudMapResolver, error) {
	_ = "STUB: not implemented" // Using the SDK's default configuration, loading additional config
	return nil, nil
}

// and credentials values from the environment variables, shared
// credentials, and shared configuration files

// Using the Config value, create the DynamoDB client

func (r *cloudMapResolver) start(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (r *cloudMapResolver) shutdown(_ context.Context) error { _ = "STUB: not implemented"; return nil }

func (r *cloudMapResolver) periodicallyResolve() { _ = "STUB: not implemented"; return }

func (r *cloudMapResolver) resolve(ctx context.Context) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// keep it always in the same order

// the list has changed!

// propagate the change

func (r *cloudMapResolver) onChange(f func([]string)) { _ = "STUB: not implemented"; return }

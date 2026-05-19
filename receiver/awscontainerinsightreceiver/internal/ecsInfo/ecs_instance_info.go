// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ecsinfo // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awscontainerinsightreceiver/internal/ecsInfo"

import (
	"context"
	"sync"
	"time"

	"go.uber.org/zap"
)

type containerInstanceInfoProvider interface {
	GetClusterName() string
	GetContainerInstanceID() string
}

type Requester interface {
	Request(ctx context.Context, path string) ([]byte, error)
}

type containerInstanceInfo struct {
	logger                   *zap.Logger
	httpClient               doer
	refreshInterval          time.Duration
	ecsAgentEndpointProvider hostIPProvider
	clusterName              string
	containerInstanceID      string
	readyC                   chan bool
	sync.RWMutex
}

type ContainerInstance struct {
	Cluster              string
	ContainerInstanceArn string
}

func newECSInstanceInfo(ctx context.Context, ecsAgentEndpointProvider hostIPProvider,
	refreshInterval time.Duration, logger *zap.Logger, httpClient doer, readyC chan bool,
) containerInstanceInfoProvider {
	_ = "STUB: not implemented"
	return *new(containerInstanceInfoProvider)
}

// stop the refresh once we get instance ID and cluster name successfully

func (cii *containerInstanceInfo) refresh(ctx context.Context) { _ = "STUB: not implemented"; return }

// notify cgroups that the clustername and instanceID is ready

func (cii *containerInstanceInfo) GetClusterName() string { _ = "STUB: not implemented"; return "" }

func (cii *containerInstanceInfo) GetContainerInstanceID() string {
	_ = "STUB: not implemented"
	return ""
}

func (cii *containerInstanceInfo) getECSAgentEndpoint() string {
	_ = "STUB: not implemented"
	return ""
}

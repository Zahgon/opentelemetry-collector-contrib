// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ecsinfo // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awscontainerinsightreceiver/internal/ecsInfo"

import (
	"context"
	"sync"
	"time"

	"go.uber.org/zap"
)

type ecsTaskInfoProvider interface {
	getRunningTaskCount() int64
	getRunningTasksInfo() []ECSTask
}

type ECSContainer struct {
	DockerID string
}
type ECSTask struct {
	KnownStatus string
	ARN         string
	Containers  []ECSContainer
}

type ECSTasksInfo struct {
	Tasks []ECSTask
}

type taskInfo struct {
	logger                  *zap.Logger
	httpClient              doer
	refreshInterval         time.Duration
	ecsTaskEndpointProvider hostIPProvider
	runningTaskCount        int64
	runningTasksInfo        []ECSTask
	readyC                  chan bool
	sync.RWMutex
}

func newECSTaskInfo(ctx context.Context, ecsTaskEndpointProvider hostIPProvider,
	refreshInterval time.Duration, logger *zap.Logger, httpClient doer, readyC chan bool,
) ecsTaskInfoProvider {
	_ = "STUB: not implemented"
	return *new(ecsTaskInfoProvider)
}

// keep refreshing to update task info and running task number

func (ti *taskInfo) getTasksInfo(ctx context.Context) (ecsTasksInfo *ECSTasksInfo) {
	_ = "STUB: not implemented"
	return nil
}

func (ti *taskInfo) refresh(ctx context.Context) { _ = "STUB: not implemented"; return }

// notify cgroups that the task info is ready

func (ti *taskInfo) getRunningTaskCount() int64 { _ = "STUB: not implemented"; return 0 }

func (ti *taskInfo) getRunningTasksInfo() []ECSTask { _ = "STUB: not implemented"; return nil }

func (ti *taskInfo) getECSAgentTaskInfoEndpoint() string { _ = "STUB: not implemented"; return "" }

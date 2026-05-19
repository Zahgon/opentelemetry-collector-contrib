// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ecsinfo // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awscontainerinsightreceiver/internal/ecsInfo"

import (
	"context"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.uber.org/zap"
)

const defaultTimeout = 1 * time.Second

type hostIPProvider interface {
	GetInstanceIP() string
	GetInstanceIPReadyC() chan bool
}

type EcsInfo struct {
	logger                *zap.Logger
	refreshInterval       time.Duration
	cancel                context.CancelFunc
	hostIPProvider        hostIPProvider
	isTaskInfoReadyC      chan bool
	isContainerInfoReadyC chan bool

	isCgroupReadyC          chan bool // close of this channel indicates cgroup is initialized. It is used only in test
	taskInfoTestReadyC      chan bool // close of this channel indicates taskinfo is initialized. It is used only in test
	containerInfoTestReadyC chan bool // close of this channel indicates container info is initialized. It is used only in test

	httpClient            doer
	containerInstanceInfo containerInstanceInfoProvider
	ecsTaskInfo           ecsTaskInfoProvider
	cgroup                cgroupScannerProvider

	containerInstanceInfoCreator func(context.Context, hostIPProvider, time.Duration, *zap.Logger, doer, chan bool) containerInstanceInfoProvider
	ecsTaskInfoCreator           func(context.Context, hostIPProvider, time.Duration, *zap.Logger, doer, chan bool) ecsTaskInfoProvider
	cgroupScannerCreator         func(context.Context, *zap.Logger, ecsTaskInfoProvider, containerInstanceInfoProvider, time.Duration) cgroupScannerProvider
}

func (e *EcsInfo) GetRunningTaskCount() int64 { _ = "STUB: not implemented"; return 0 }

func (e *EcsInfo) GetCPUReserved() int64 { _ = "STUB: not implemented"; return 0 }

func (e *EcsInfo) GetMemReserved() int64 { _ = "STUB: not implemented"; return 0 }

func (e *EcsInfo) GetContainerInstanceID() string { _ = "STUB: not implemented"; return "" }

func (e *EcsInfo) GetClusterName() string { _ = "STUB: not implemented"; return "" }

type ecsInfoOption func(*EcsInfo)

// New creates a k8sApiServer which can generate cluster-level metrics
func NewECSInfo(refreshInterval time.Duration, hostIPProvider hostIPProvider, host component.Host, settings component.TelemetrySettings, options ...ecsInfoOption) (*EcsInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *EcsInfo) initContainerInfo(ctx context.Context) { _ = "STUB: not implemented"; return }

func (e *EcsInfo) initTaskInfo(ctx context.Context) { _ = "STUB: not implemented"; return }

func (e *EcsInfo) initCgroupScanner(ctx context.Context) { _ = "STUB: not implemented"; return }

// Shutdown stops the ecs Info
func (e *EcsInfo) Shutdown() { _ = "STUB: not implemented"; return }

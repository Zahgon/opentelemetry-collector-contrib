// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ecsinfo // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awscontainerinsightreceiver/internal/ecsInfo"

import (
	"context"
	"sync"
	"time"

	"go.uber.org/zap"
)

const (
	ecsInstanceMountConfigPath = "/proc/self/mountinfo"
	// infinity magic number for cgroup: https://unix.stackexchange.com/questions/420906/what-is-the-value-for-the-cgroups-limit-in-bytes-if-the-memory-is-not-restricte
	kernelMagicCodeNotSet = int64(9223372036854771712)
)

type cgroupScanner struct {
	logger                        *zap.Logger
	mountPoint                    string
	ecsTaskInfoProvider           ecsTaskInfoProvider
	containerInstanceInfoProvider containerInstanceInfoProvider
	refreshInterval               time.Duration
	ctx                           context.Context
	sync.RWMutex

	cpuReserved int64
	memReserved int64
}

type cgroupScannerProvider interface {
	getCPUReserved() int64
	getMemReserved() int64

	// use for test
	getCPUReservedInTask(taskID, clusterName string) int64
	getMEMReservedInTask(taskID, clusterName string, containers []ECSContainer) int64
}

func newCGroupScanner(ctx context.Context, mountConfigPath string, logger *zap.Logger, ecsTaskInfoProvider ecsTaskInfoProvider, containerInstanceInfoProvider containerInstanceInfoProvider, refreshInterval time.Duration) cgroupScannerProvider {
	_ = "STUB: not implemented"
	return *new(cgroupScannerProvider)
}

func (c *cgroupScanner) refresh() { _ = "STUB: not implemented"; return }

// ignore the one only consume 2 shares which is the default value in cgroup

func newCGroupScannerForContainer(ctx context.Context, logger *zap.Logger, ecsTaskInfoProvider ecsTaskInfoProvider, containerInstanceInfoProvider containerInstanceInfoProvider, refreshInterval time.Duration) cgroupScannerProvider {
	_ = "STUB: not implemented"
	return *new(cgroupScannerProvider)
}

func (c *cgroupScanner) getCPUReservedInTask(taskID, clusterName string) int64 {
	_ = "STUB: not implemented"
	return 0
}

// check if hard limit is configured

func (c *cgroupScanner) getMEMReservedInTask(taskID, clusterName string, containers []ECSContainer) int64 {
	_ = "STUB: not implemented"
	return 0
}

// sum the containers' memory if the task's memory limit is not configured

// soft limit first

// try hard limit when soft limit is not configured

func readString(dirpath, file string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// Read

// Ignore nonexistent files

func readInt64(dirpath, file string) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func getCGroupMountPoint(mountConfigPath string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// safe as mountinfo encodes mountpoints with spaces as \040.
// an example: 26 22 0:23 / /cgroup/cpu rw,relatime - cgroup cgroup rw,cpu

// this is an error as we can't detect if the mount is for "cgroup"

// check that the mount is properly formatted.

func getCGroupPathForTask(cgroupMount, controller, taskID, clusterName string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Task cgroup path does not exist, fallback to try legacy Task cgroup path,
// legacy cgroup path of task with new format ARN used to contain cluster name,
// before ECS Agent PR https://github.com/aws/amazon-ecs-agent/pull/2497/

func (c *cgroupScanner) getCPUReserved() int64 { _ = "STUB: not implemented"; return 0 }

func (c *cgroupScanner) getMemReserved() int64 { _ = "STUB: not implemented"; return 0 }

// There are two formats of Task ARN (https://docs.aws.amazon.com/AmazonECS/latest/userguide/ecs-account-settings.html#ecs-resource-ids)
// arn:aws:ecs:region:aws_account_id:task/task-id
// arn:aws:ecs:region:aws_account_id:task/cluster-name/task-id
// we should get "task-id" as result no matter what format the ARN is.
func getTaskCgroupPathFromARN(arn string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

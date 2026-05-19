// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package host // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awscontainerinsightreceiver/internal/host"

import (
	"context"
	"time"
)

const (
	rootfs     = "/rootfs"            // the root directory "/" is mounted as "/rootfs" in container
	hostProc   = rootfs + "/proc"     // "/rootfs/proc" in container refers to the host proc directory "/proc"
	hostMounts = hostProc + "/mounts" // "/rootfs/proc/mounts" in container refers to "/proc/mounts" in the host
)

func hostJitter(maxDuration time.Duration) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// Right shift the uint64 hash by one to make sure the jitter duration is always positive

// RefreshUntil executes the refresh() function periodically with the given refresh interval
// until shouldRefresh() return false or the context is canceled
func RefreshUntil(ctx context.Context, refresh func(context.Context), refreshInterval time.Duration,
	shouldRefresh func() bool, maxJitterTime time.Duration,
) {
	_ = "STUB: not implemented"
	return

	// add some sleep jitter to prevent a large number of receivers calling the ec2 api at the same time
}

// initial refresh

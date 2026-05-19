// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build linux || darwin || freebsd || openbsd

package processesscraper // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/processesscraper"

import (
	"context"
	"runtime"

	"github.com/shirou/gopsutil/v4/process"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/processesscraper/internal/metadata"
)

const (
	enableProcessesCount   = true
	enableProcessesCreated = runtime.GOOS == "openbsd" || runtime.GOOS == "linux"
)

func (s *processesScraper) getProcessesMetadata(ctx context.Context) (processesMetadata, error) {
	_ = "STUB: not implemented"
	return *new(processesMetadata), nil
}

// We expect an error in the case that a process has
// been terminated as we run this code.

// Processes are actively changing as we run this code, so this reason
// the above loop will tend to underestimate process counts.
// getMiscStats is a single read/syscall so it should be more accurate.

func toAttributeStatus(status []string) (metadata.AttributeStatus, bool) {
	_ = "STUB: not implemented"
	return *new(metadata.AttributeStatus), false
}

var charToState = map[string]metadata.AttributeStatus{
	process.Blocked:  metadata.AttributeStatusBlocked,
	process.Daemon:   metadata.AttributeStatusDaemon,
	process.Detached: metadata.AttributeStatusDetached,
	process.Idle:     metadata.AttributeStatusIdle,
	process.Lock:     metadata.AttributeStatusLocked,
	process.Orphan:   metadata.AttributeStatusOrphan,
	process.Running:  metadata.AttributeStatusRunning,
	process.Sleep:    metadata.AttributeStatusSleeping,
	process.Stop:     metadata.AttributeStatusStopped,
	process.System:   metadata.AttributeStatusSystem,
	process.Wait:     metadata.AttributeStatusPaging,
	process.Zombie:   metadata.AttributeStatusZombies,
}

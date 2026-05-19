// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package filesystemscraper // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/filesystemscraper"

import (
	"context"

	"github.com/shirou/gopsutil/v4/disk"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/scraper"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/filesystemscraper/internal/metadata"
)

const (
	standardMetricsLen = 1
	metricsLen         = standardMetricsLen + systemSpecificMetricsLen
)

// filesystemsScraper for FileSystem Metrics
type filesystemsScraper struct {
	settings scraper.Settings
	config   *Config
	mb       *metadata.MetricsBuilder
	fsFilter fsFilter

	// for mocking gopsutil disk.Partitions & disk.Usage
	bootTime   func(context.Context) (uint64, error)
	partitions func(context.Context, bool) ([]disk.PartitionStat, error)
	usage      func(context.Context, string) (*disk.UsageStat, error)
}

type deviceUsage struct {
	partition disk.PartitionStat
	usage     *disk.UsageStat
}

// newFileSystemScraper creates a FileSystem Scraper
func newFileSystemScraper(_ context.Context, settings scraper.Settings, cfg *Config) (*filesystemsScraper, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *filesystemsScraper) start(ctx context.Context, _ component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *filesystemsScraper) scrape(ctx context.Context) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

// Log a debug message instead of an error message if a drive is
// locked and unavailable. For this particular case, we do not want
// to log an error message on every poll.
// See: https://github.com/open-telemetry/opentelemetry-collector-contrib/issues/18236

func getMountMode(opts []string) string { _ = "STUB: not implemented"; return "" }

func (f *fsFilter) includePartition(partition disk.PartitionStat) bool {
	_ = "STUB: not implemented"
	// If filters do not exist, return early.
	return false
}

func (f *fsFilter) includeDevice(deviceName string) bool { _ = "STUB: not implemented"; return false }

func (f *fsFilter) includeFSType(fsType string) bool { _ = "STUB: not implemented"; return false }

func (f *fsFilter) includeMountPoint(mountPoint string) bool {
	_ = "STUB: not implemented"
	return false
}

// translateMountsRootPath translates a mountpoint from the host perspective to the chrooted perspective.
func translateMountpoint(ctx context.Context, rootPath, mountpoint string) string {
	_ = "STUB: not implemented"
	return ""
}

// gopsutil first reads <rootPath>/proc/1/mountinfo, and if this fails it reads <rootPath>/proc/self/mountinfo as a fallback.
// If the collector process runs in a different mount namespace than PID 1 on the host (e.g. on Kubernetes):
//
// <rootPath>/proc/1/mountinfo contains the mountpoints from the perspective of PID 1 on the host (without rootPath prefix)
// <rootPath>/proc/self/mountinfo contains the mountpoints from the perspective of the collector process (with the rootPath prefix)
//
// Example on Fedora 43:
// $ docker run -v /:/hostfs:ro alpine cat /hostfs/proc/1/mountinfo | grep ext4
// 61 73 259:2 / /boot rw,relatime shared:84 - ext4 /dev/nvme0n1p2 rw,seclabel
//
// $ docker run -v /:/hostfs:ro alpine cat /hostfs/proc/self/mountinfo | grep ext4
// 5111 5048 259:2 / /hostfs/boot ro,relatime master:84 - ext4 /dev/nvme0n1p2 rw,seclabel
//
// Therefore, do not add rootPath if the mountpath has already a rootPath prefix.

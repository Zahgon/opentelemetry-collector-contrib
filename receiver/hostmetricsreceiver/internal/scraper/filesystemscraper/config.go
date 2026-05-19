// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package filesystemscraper // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/filesystemscraper"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/filter/filterset"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/filesystemscraper/internal/metadata"
)

// Config relating to FileSystem Metric Scraper.
type Config struct {
	// MetricsBuilderConfig allows to customize scraped metrics/attributes representation.
	metadata.MetricsBuilderConfig `mapstructure:",squash"`

	// IncludeVirtualFS will also capture filesystems such as tmpfs, ramfs
	// and other filesystem types that do no have an associated physical device.
	IncludeVirtualFS bool `mapstructure:"include_virtual_filesystems"`

	// IncludeDevices specifies a filter on the devices that should be included in the generated metrics.
	IncludeDevices DeviceMatchConfig `mapstructure:"include_devices"`
	// ExcludeDevices specifies a filter on the devices that should be excluded from the generated metrics.
	ExcludeDevices DeviceMatchConfig `mapstructure:"exclude_devices"`

	// IncludeFSTypes specifies a filter on the filesystem types that should be included in the generated metrics.
	IncludeFSTypes FSTypeMatchConfig `mapstructure:"include_fs_types"`
	// ExcludeFSTypes specifies a filter on the filesystem types points that should be excluded from the generated metrics.
	ExcludeFSTypes FSTypeMatchConfig `mapstructure:"exclude_fs_types"`

	// IncludeMountPoints specifies a filter on the mount points that should be included in the generated metrics.
	// When `root_path` is set, the mount points must be from the host's perspective.
	IncludeMountPoints MountPointMatchConfig `mapstructure:"include_mount_points"`
	// ExcludeMountPoints specifies a filter on the mount points that should be excluded from the generated metrics.
	// When `root_path` is set, the mount points must be from the host's perspective.
	ExcludeMountPoints MountPointMatchConfig `mapstructure:"exclude_mount_points"`

	rootPath string `mapstructure:"-"`
}

type DeviceMatchConfig struct {
	filterset.Config `mapstructure:",squash"`

	Devices []string `mapstructure:"devices"`
}

type FSTypeMatchConfig struct {
	filterset.Config `mapstructure:",squash"`

	FSTypes []string `mapstructure:"fs_types"`
}

type MountPointMatchConfig struct {
	filterset.Config `mapstructure:",squash"`

	MountPoints []string `mapstructure:"mount_points"`
}

type fsFilter struct {
	includeDeviceFilter     filterset.FilterSet
	excludeDeviceFilter     filterset.FilterSet
	includeFSTypeFilter     filterset.FilterSet
	excludeFSTypeFilter     filterset.FilterSet
	includeMountPointFilter filterset.FilterSet
	excludeMountPointFilter filterset.FilterSet
	filtersExist            bool
}

func (cfg *Config) SetRootPath(rootPath string) { _ = "STUB: not implemented"; return }

func (cfg *Config) createFilter() (*fsFilter, error) { _ = "STUB: not implemented"; return nil, nil }

func (f *fsFilter) setFiltersExist() { _ = "STUB: not implemented"; return }

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package attrs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/fileconsumer/attrs"

import (
	"os"
)

const (
	LogFileName           = "log.file.name"
	LogFilePath           = "log.file.path"
	LogFileNameResolved   = "log.file.name_resolved"
	LogFilePathResolved   = "log.file.path_resolved"
	LogFileOwnerName      = "log.file.owner.name"
	LogFileOwnerGroupName = "log.file.owner.group.name"
	LogFilePermissions    = "log.file.permissions"
	LogFileRecordNumber   = "log.file.record_number"
	LogFileRecordOffset   = "log.file.record_offset"
)

type Resolver struct {
	IncludeFileName           bool `mapstructure:"include_file_name,omitempty"`
	IncludeFilePath           bool `mapstructure:"include_file_path,omitempty"`
	IncludeFileNameResolved   bool `mapstructure:"include_file_name_resolved,omitempty"`
	IncludeFilePathResolved   bool `mapstructure:"include_file_path_resolved,omitempty"`
	IncludeFileOwnerName      bool `mapstructure:"include_file_owner_name,omitempty"`
	IncludeFileOwnerGroupName bool `mapstructure:"include_file_owner_group_name,omitempty"`
	IncludeFilePermissions    bool `mapstructure:"include_file_permissions,omitempty"`
}

func (r *Resolver) Resolve(file *os.File) (attributes map[string]any, err error) {
	_ = "STUB: not implemented"

	// size 2 is sufficient if not resolving symlinks. This optimizes for the most performant cases.
	return nil, nil
}

// Dirty solution, waiting for this permanent fix https://github.com/golang/go/issues/39786
// EvalSymlinks on windows is partially working depending on the way you use Symlinks and Junctions

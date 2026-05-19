// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build !windows

package attrs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/fileconsumer/attrs"

import (
	"os"
)

func (r *Resolver) addPermissionInfo(file *os.File, attributes map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

// Format the file mode as a 3-digit octal string (e.g., "755")

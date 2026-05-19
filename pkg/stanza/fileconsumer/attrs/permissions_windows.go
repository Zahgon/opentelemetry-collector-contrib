// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build windows

package attrs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/fileconsumer/attrs"

import (
	"os"
)

func (*Resolver) addPermissionInfo(_ *os.File, _ map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

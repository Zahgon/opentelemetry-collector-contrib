// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package clientutil // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal/clientutil"

import (
	"go.opentelemetry.io/collector/client"
)

// Address returns the address of the client connecting to the collector.
func Address(client client.Info) string { _ = "STUB: not implemented"; return "" }

// If this is not a known address type, check for known "untyped" formats.
// 1.1.1.1:<port>

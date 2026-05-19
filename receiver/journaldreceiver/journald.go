// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build linux

package journaldreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/journaldreceiver"

import (
	"go.opentelemetry.io/collector/receiver"
)

// newFactoryAdapter creates a factory for journald receiver
func newFactoryAdapter() receiver.Factory { _ = "STUB: not implemented"; return *new(receiver.Factory) }

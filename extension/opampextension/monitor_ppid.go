// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package opampextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/opampextension"

import (
	"context"
	"time"

	"go.opentelemetry.io/collector/component/componentstatus"
)

// monitorPPID polls for the existence of ppid.
// If the specified ppid no longer exists, a fatal error event is reported via the passed in reportStatus function.
func monitorPPID(ctx context.Context, interval time.Duration, ppid int32, reportStatus func(*componentstatus.Event)) {
	_ = "STUB: not implemented"
	return
}

// OK; Poll again to make sure PID exists

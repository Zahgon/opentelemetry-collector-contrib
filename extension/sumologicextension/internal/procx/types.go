// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package procx // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/sumologicextension/internal/procx"

import "github.com/shirou/gopsutil/v4/process"

type SumoTag string

func (s SumoTag) String() string { _ = "STUB: not implemented"; return "" }

type ProcessIdentifier string

func (p ProcessIdentifier) String() string { _ = "STUB: not implemented"; return "" }

type processWrapper struct {
	process *process.Process
}

func (pw *processWrapper) Pid() int32 { _ = "STUB: not implemented"; return 0 }

func (pw *processWrapper) Name() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (pw *processWrapper) Cmdline() (string, error) { _ = "STUB: not implemented"; return "", nil }

type Process interface {
	Pid() int32
	Name() (string, error)
	Cmdline() (string, error)
}

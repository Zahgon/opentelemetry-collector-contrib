// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package supervisor

// pidProvider provides the PID of the current process
type pidProvider interface {
	PID() int
}

type defaultPIDProvider struct{}

func (defaultPIDProvider) PID() int { _ = "STUB: not implemented"; return 0 }

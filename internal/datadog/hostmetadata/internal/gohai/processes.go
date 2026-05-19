// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build linux || darwin

package gohai // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/datadog/hostmetadata/internal/gohai"

import (
	"github.com/DataDog/datadog-agent/pkg/opentelemetry-mapping-go/inframetadata/gohai"
	"go.uber.org/zap"
)

// NewProcessesPayload builds a payload of processes metadata collected from gohai.
func NewProcessesPayload(hostname string, logger *zap.Logger) *gohai.ProcessesPayload {
	_ = "STUB: not implemented"
	// Get processes metadata from gohai
	return nil
}

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build (!windows || !arm64) && !aix

package gohai // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/datadog/hostmetadata/internal/gohai"

import (
	"github.com/DataDog/datadog-agent/pkg/opentelemetry-mapping-go/inframetadata/gohai"
	"go.uber.org/zap"
)

// NewPayload builds a payload of every metadata collected with gohai except processes metadata.
// Parts of this are based on datadog-agent code
// https://github.com/DataDog/datadog-agent/blob/94a28d9cee3f1c886b3866e8208be5b2a8c2c217/pkg/metadata/internal/gohai/gohai.go#L27-L32
func NewPayload(logger *zap.Logger) gohai.Payload {
	_ = "STUB: not implemented"
	return *new(gohai.Payload)
}

func newGohai(logger *zap.Logger) *gohai.Gohai { _ = "STUB: not implemented"; return nil }

// in case of containerized environment, this would return pod id not node's ip

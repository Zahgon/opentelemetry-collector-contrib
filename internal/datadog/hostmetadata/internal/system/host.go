// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package system contains the system hostname provider
package system // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/datadog/hostmetadata/internal/system"

import (
	"context"
	"sync"

	"github.com/DataDog/datadog-agent/pkg/opentelemetry-mapping-go/otlp/attributes/source"
	"go.uber.org/zap"
)

type HostInfo struct {
	OS   string
	FQDN string
}

// GetHostInfo gets system information about the hostname
func GetHostInfo(logger *zap.Logger) (hostInfo *HostInfo) { _ = "STUB: not implemented"; return nil }

// GetHostname gets the hostname provided by the system
func (hi *HostInfo) GetHostname(logger *zap.Logger) string {
	_ = "STUB: not implemented"

	// Don't report failure since FQDN was just not available
	return ""
}

var _ source.Provider = (*Provider)(nil)

type Provider struct {
	once     sync.Once
	hostInfo HostInfo

	logger *zap.Logger
}

func (p *Provider) fillHostInfo() { _ = "STUB: not implemented"; return }

func (p *Provider) Source(context.Context) (source.Source, error) {
	_ = "STUB: not implemented"
	return *new(source.Source), nil
}

func (p *Provider) HostInfo() *HostInfo { _ = "STUB: not implemented"; return nil }

func NewProvider(logger *zap.Logger) *Provider { _ = "STUB: not implemented"; return nil }

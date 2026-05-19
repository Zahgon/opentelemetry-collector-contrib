// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package system // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/system"

import (
	"context"
	"net"

	"github.com/shirou/gopsutil/v4/cpu"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/processor"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/metadataproviders/system"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/system/internal/metadata"
)

const (
	// TypeStr is type of detector.
	TypeStr = "system"
)

var hostnameSourcesMap = map[string]func(*Detector) (string, error){
	"os":     getHostname,
	"dns":    getFQDN,
	"cname":  lookupCNAME,
	"lookup": reverseLookupHost,
}

var _ internal.Detector = (*Detector)(nil)

// Detector is a system metadata detector
type Detector struct {
	provider system.Provider
	logger   *zap.Logger
	cfg      Config
	rb       *metadata.ResourceBuilder
}

// NewDetector creates a new system metadata detector
func NewDetector(p processor.Settings, dcfg internal.DetectorConfig) (internal.Detector, error) {
	_ = "STUB: not implemented"
	return *new(internal.Detector), nil
}

// toIEEERA converts a MAC address to IEEE RA format.
// Per the spec: "MAC Addresses MUST be represented in IEEE RA hexadecimal form: as hyphen-separated
// octets in uppercase hexadecimal form from most to least significant."
// Golang returns MAC addresses as colon-separated octets in lowercase hexadecimal form from most
// to least significant, so we need to:
// - Replace colons with hyphens
// - Convert to uppercase
func toIEEERA(mac net.HardwareAddr) string { _ = "STUB: not implemented"; return "" }

// Detect detects system metadata and returns a resource with the available ones
func (d *Detector) Detect(ctx context.Context) (resource pcommon.Resource, schemaURL string, err error) {
	_ = "STUB: not implemented"
	return *new(pcommon.Resource), "", nil
}

// getHostname returns OS hostname
func getHostname(d *Detector) (string, error) { _ = "STUB: not implemented"; return "", nil }

// getFQDN returns FQDN of the host
func getFQDN(d *Detector) (string, error) { _ = "STUB: not implemented"; return "", nil }

func lookupCNAME(d *Detector) (string, error) { _ = "STUB: not implemented"; return "", nil }

func reverseLookupHost(d *Detector) (string, error) { _ = "STUB: not implemented"; return "", nil }

func setHostCPUInfo(d *Detector, cpuInfo cpu.InfoStat) { _ = "STUB: not implemented"; return }

// For windows, this field is left blank. See https://github.com/shirou/gopsutil/blob/v3.23.9/cpu/cpu_windows.go#L113
// Skip setting modelId if the field is blank.
// ISSUE: https://github.com/open-telemetry/opentelemetry-collector-contrib/issues/27675

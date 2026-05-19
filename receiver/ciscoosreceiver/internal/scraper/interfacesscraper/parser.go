// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package interfacesscraper // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/ciscoosreceiver/internal/scraper/interfacesscraper"

import (
	"go.uber.org/zap"
)

// Interface represents a network interface on a Cisco device
type Interface struct {
	Name        string
	MACAddress  string
	Description string

	OperStatus string

	InputErrors  float64
	OutputErrors float64

	InputDrops  float64
	OutputDrops float64

	InputBytes  float64
	OutputBytes float64

	InputBroadcast float64
	InputMulticast float64

	Speed       int64
	SpeedString string
}

const (
	StatusUp   = "up"
	StatusDown = "down"
)

func NewInterface(name string) *Interface { _ = "STUB: not implemented"; return nil }

// GetOperStatusInt converts operational status to integer (1=up, 0=down)
func (i *Interface) GetOperStatusInt() int64 { _ = "STUB: not implemented"; return 0 }

// Validate ensures interface has required data and valid status
func (i *Interface) Validate() bool { _ = "STUB: not implemented"; return false }

// parseStatus normalizes status strings to "up" or "down"
func parseStatus(status string) string { _ = "STUB: not implemented"; return "" }

// formatSpeed converts speed in bps to human-readable format
func formatSpeed(speedBps int64) string { _ = "STUB: not implemented"; return "" }

// str2float64 converts string to float64
func str2float64(s string) float64 { _ = "STUB: not implemented"; return 0 }

// parseInterfaces parses interface information from command output
func parseInterfaces(output string, logger *zap.Logger) []*Interface {
	_ = "STUB: not implemented"
	return nil
}

// parseSimpleInterfaces parses "show interface brief" output as fallback
func parseSimpleInterfaces(output string, _ *zap.Logger) []*Interface {
	_ = "STUB: not implemented"
	return nil
}

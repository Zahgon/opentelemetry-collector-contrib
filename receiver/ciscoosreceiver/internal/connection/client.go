// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package connection // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/ciscoosreceiver/internal/connection"

import (
	"go.uber.org/zap"
)

// RPCClient represents RPC client for executing Cisco commands
type RPCClient struct {
	SSHClient *Client
	OSType    string
	Logger    *zap.Logger
}

// GetOSType returns detected Cisco OS type
func (r *RPCClient) GetOSType() string { _ = "STUB: not implemented"; return "" }

// Default

// GetCommand returns the appropriate command for the OS type and feature
func (r *RPCClient) GetCommand(feature string) string { _ = "STUB: not implemented"; return "" }

// ExecuteCommand executes a command on the Cisco device
func (r *RPCClient) ExecuteCommand(command string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

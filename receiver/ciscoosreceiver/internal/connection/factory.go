// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package connection // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/ciscoosreceiver/internal/connection"

import (
	"context"

	"go.uber.org/zap"
	cryptossh "golang.org/x/crypto/ssh"
)

// EstablishDeviceConnection creates a device connection using the provided DeviceConfig.
func EstablishDeviceConnection(ctx context.Context, device DeviceConfig, logger *zap.Logger) (*RPCClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// #nosec G106 - Insecure for lab/demo only

// buildAuthMethods builds SSH authentication methods from the provided auth config.
// Supports both password and SSH key file authentication.
func buildAuthMethods(auth AuthConfig, logger *zap.Logger) ([]cryptossh.AuthMethod, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func publicKeyAuth(keyFile string) (cryptossh.AuthMethod, error) {
	_ = "STUB: not implemented"
	return *new(cryptossh.AuthMethod), nil
}

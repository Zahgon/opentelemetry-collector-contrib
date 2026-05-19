// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package azuremonitorreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/azuremonitorreceiver"

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"go.opentelemetry.io/collector/component"
	"go.uber.org/zap"
)

func loadTokenProvider(host component.Host, idAuth component.ID) (azcore.TokenCredential, error) {
	_ = "STUB: not implemented"
	return *new(azcore.TokenCredential), nil
}

func loadCredentials(logger *zap.Logger, cfg *Config, host component.Host) (azcore.TokenCredential, error) {
	_ = "STUB: not implemented"
	return *new(azcore.TokenCredential), nil
}

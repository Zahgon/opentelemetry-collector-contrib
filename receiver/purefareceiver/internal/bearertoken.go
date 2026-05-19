// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package internal // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/purefareceiver/internal"

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configauth"
)

func RetrieveBearerToken(authCfg configauth.Config, extensions map[component.ID]component.Component) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

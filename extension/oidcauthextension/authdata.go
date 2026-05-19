// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package oidcauthextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/oidcauthextension"

import (
	"go.opentelemetry.io/collector/client"
)

const claimsPrefix = "claims."

var _ client.AuthData = (*authData)(nil)

type authData struct {
	raw    string
	claims map[string]any

	subject    string
	membership []string
}

func (a *authData) GetAttribute(name string) any { _ = "STUB: not implemented"; return *new(any) }

// If the name does not start with the claims prefix, we return nil.
// This ensures not mixing custom JWT claims with hard defined attributes.

func (a *authData) GetAttributeNames() []string { _ = "STUB: not implemented"; return nil }

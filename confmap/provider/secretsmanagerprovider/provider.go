// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:generate make mdatagen

package secretsmanagerprovider // import "github.com/open-telemetry/opentelemetry-collector-contrib/confmap/provider/secretsmanagerprovider"

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"go.opentelemetry.io/collector/confmap"
	"go.uber.org/zap"
)

type secretsManagerClient interface {
	GetSecretValue(ctx context.Context, params *secretsmanager.GetSecretValueInput, optFns ...func(*secretsmanager.Options)) (*secretsmanager.GetSecretValueOutput, error)
}

const (
	schemeName = "secretsmanager"
)

type provider struct {
	client secretsManagerClient
	logger *zap.Logger
}

// NewFactory returns a new confmap.ProviderFactory that creates a confmap.Provider
// which reads configuration the given AWS Secrets Manager Name or ARN.
//
// This Provider supports "secretsmanager" scheme, and can be called with a selector:
// `secretsmanager:NAME_OR_ARN`
//
// A default value for unset variable can be provided after :- suffix, for example:
// `secretsmanager:NAME_OR_ARN:-default_value`
func NewFactory() confmap.ProviderFactory {
	_ = "STUB: not implemented"
	return *new(confmap.ProviderFactory)
}

func newWithSettings(ps confmap.ProviderSettings) confmap.Provider {
	_ = "STUB: not implemented"
	return *new(confmap.Provider)
}

func (provider *provider) Retrieve(ctx context.Context, uri string, _ confmap.WatcherFunc) (*confmap.Retrieved, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// initialize the secrets manager client in the first call of Retrieve

// split by :- to get the default value

// split by # to get the json key

func (*provider) Scheme() string { _ = "STUB: not implemented"; return "" }

func (*provider) Shutdown(context.Context) error { _ = "STUB: not implemented"; return nil }

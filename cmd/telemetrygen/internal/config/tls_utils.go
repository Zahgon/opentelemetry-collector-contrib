// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"crypto/tls"
	"crypto/x509"

	"google.golang.org/grpc/credentials"
)

// caPool loads CA certificate from a file and returns a CertPool.
// The certPool is used to set RootCAs in certificate verification.
func caPool(caFile string) (*x509.CertPool, error) { _ = "STUB: not implemented"; return nil, nil }

func GetTLSCredentialsForGRPCExporter(
	caFile string,
	cAuth ClientAuth,
	insecureSkipVerify bool,
) (credentials.TransportCredentials, error) {
	_ = "STUB: not implemented"
	return *new(credentials.TransportCredentials), nil
}

func GetTLSCredentialsForHTTPExporter(
	caFile string,
	cAuth ClientAuth,
	insecureSkipVerify bool,
) (*tls.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getTLSConfig(caFile string, cAuth ClientAuth, insecureSkipVerify bool) (*tls.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Configuration for mTLS

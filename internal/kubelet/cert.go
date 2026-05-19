// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package kubelet // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/kubelet"

import (
	"crypto/x509"
)

func systemCertPoolPlusPath(certPath string) (*x509.CertPool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func certPoolPlusPath(certPool *x509.CertPool, certPath string) (*x509.CertPool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

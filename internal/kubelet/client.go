// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package kubelet // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/kubelet"

import (
	"crypto/tls"
	"crypto/x509"
	"net/http"

	"go.uber.org/zap"
)

const (
	svcAcctCACertPath   = "/var/run/secrets/kubernetes.io/serviceaccount/ca.crt"
	svcAcctTokenPath    = "/var/run/secrets/kubernetes.io/serviceaccount/token" // #nosec
	defaultSecurePort   = "10250"
	defaultReadOnlyPort = "10255"
)

type Client interface {
	Get(path string) ([]byte, error)
}

func NewClientProvider(endpoint string, cfg *ClientConfig, logger *zap.Logger) (ClientProvider, error) {
	_ = "STUB: not implemented"
	return *new(ClientProvider), nil
}

type ClientProvider interface {
	BuildClient() (Client, error)
}

type kubeConfigClientProvider struct {
	endpoint string
	cfg      *ClientConfig
	logger   *zap.Logger
}

func (p *kubeConfigClientProvider) BuildClient() (Client, error) {
	_ = "STUB: not implemented"
	return *new(Client), nil
}

// Override InsecureSkipVerify from kubeconfig

type readOnlyClientProvider struct {
	endpoint string
	logger   *zap.Logger
}

func (p *readOnlyClientProvider) BuildClient() (Client, error) {
	_ = "STUB: not implemented"
	return *new(Client), nil
}

type tlsClientProvider struct {
	endpoint string
	cfg      *ClientConfig
	logger   *zap.Logger
}

func (p *tlsClientProvider) BuildClient() (Client, error) {
	_ = "STUB: not implemented"
	return *new(Client), nil
}

type saClientProvider struct {
	endpoint   string
	caCertPath string
	cfg        *ClientConfig
	tokenPath  string
	logger     *zap.Logger
}

func (p *saClientProvider) BuildClient() (Client, error) {
	_ = "STUB: not implemented"
	return *new(Client), nil
}

func defaultTLSClient(
	endpoint string,
	insecureSkipVerify bool,
	rootCAs *x509.CertPool,
	certificates []tls.Certificate,
	tok []byte,
	logger *zap.Logger,
) (*clientImpl, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// buildEndpoint builds a kubelet endpoint based on value provided by user and whether secure or read-only endpoint
// should be used.
func buildEndpoint(endpoint string, useSecurePort bool, logger *zap.Logger) (string, error) {
	_ = "STUB: not implemented"
	return "",

		// This will work if hostNetwork is turned on, in which case the pod has access
		// to the node's loopback device.
		// https://kubernetes.io/docs/concepts/policy/pod-security-policy/#host-namespaces
		nil
}

func defaultTransport() *http.Transport { _ = "STUB: not implemented"; return nil }

// clientImpl

var _ Client = (*clientImpl)(nil)

type clientImpl struct {
	baseURL    string
	httpClient http.Client
	logger     *zap.Logger
	tok        []byte
}

func (c *clientImpl) Get(path string) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (c *clientImpl) buildReq(p string) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

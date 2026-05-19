// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package kubelet // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/kubeletstatsreceiver/internal/kubelet"

import (
	kube "github.com/open-telemetry/opentelemetry-collector-contrib/internal/kubelet"
)

// RestClient is swappable for testing.
type RestClient interface {
	StatsSummary() ([]byte, error)
	Pods() ([]byte, error)
}

// HTTPRestClient is a thin wrapper around a kubelet client, encapsulating endpoints
// and their corresponding http methods. The endpoints /stats/container /spec/
// are excluded because they require cadvisor. The /metrics endpoint is excluded
// because it returns Prometheus data.
type HTTPRestClient struct {
	client kube.Client
}

func NewRestClient(client kube.Client) *HTTPRestClient { _ = "STUB: not implemented"; return nil }

func (c *HTTPRestClient) StatsSummary() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (c *HTTPRestClient) Pods() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

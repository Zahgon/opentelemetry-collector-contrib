// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package gcp // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/gcp"

import (
	"context"
	"regexp"

	compute "cloud.google.com/go/compute/apiv1"
	computepb "cloud.google.com/go/compute/apiv1/computepb"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/processor"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal"
	localMetadata "github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/gcp/internal/metadata"
)

const (
	// TypeStr is type of detector.
	TypeStr        = "gcp"
	gceLabelPrefix = "gcp.gce.instance.labels."
)

// NewDetector returns a detector which can detect resource attributes on:
// * Google Compute Engine (GCE).
// * Google Kubernetes Engine (GKE).
// * Google App Engine (GAE).
// * Cloud Run.
// * Cloud Functions.
// * Bare Metal Solutions (BMS).
func NewDetector(set processor.Settings, dcfg internal.DetectorConfig) (internal.Detector, error) {
	_ = "STUB: not implemented"
	return *new(internal.Detector), nil
}

type detector struct {
	logger           *zap.Logger
	detector         gcpDetector
	rb               *localMetadata.ResourceBuilder
	labelKeyRegexes  []*regexp.Regexp
	gceClientBuilder instancesBuilder
}

func (d *detector) Detect(ctx context.Context) (resource pcommon.Resource, schemaURL string, err error) {
	_ = "STUB: not implemented"
	return *new(pcommon.Resource), "", nil
}

// GCEHostname is fallible on GKE, since it's not available when using workload identity.

// We don't support this platform yet, so just return with what we have

type instancesAPI interface {
	Get(ctx context.Context, req *computepb.GetInstanceRequest) (*computepb.Instance, error)
	Close() error
}

type instancesBuilder interface {
	buildClient(ctx context.Context) (instancesAPI, error)
}

type instancesRESTBuilder struct{}

func (*instancesRESTBuilder) buildClient(ctx context.Context) (instancesAPI, error) {
	_ = "STUB: not implemented"
	return *new(instancesAPI), nil
}

// picks up GCE metadata creds automatically

type instancesRESTClient struct{ inner *compute.InstancesClient }

func (c *instancesRESTClient) Get(ctx context.Context, req *computepb.GetInstanceRequest) (*computepb.Instance, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *instancesRESTClient) Close() error { _ = "STUB: not implemented"; return nil }

func fetchGCELabels(ctx context.Context, svc instancesAPI, project, zone, instance string, labelKeyRegexes []*regexp.Regexp) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func compileLabelRegexes(cfg Config) ([]*regexp.Regexp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func regexArrayMatch(arr []*regexp.Regexp, val string) bool {
	_ = "STUB: not implemented"
	return false
}

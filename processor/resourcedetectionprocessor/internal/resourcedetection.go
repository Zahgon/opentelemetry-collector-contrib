// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package internal contains an interface for detecting resource information,
// and a provider to merge the resources returned by a slice of custom detectors.
package internal // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal"

import (
	"context"
	"net/http"
	"sync"
	"sync/atomic"
	"time"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/processor"
	"go.uber.org/zap"
)

type DetectorType string

type Detector interface {
	Detect(ctx context.Context) (resource pcommon.Resource, schemaURL string, err error)
}

type DetectorConfig any

type ResourceDetectorConfig interface {
	GetConfigFromType(DetectorType) DetectorConfig
}

type DetectorFactory func(processor.Settings, DetectorConfig) (Detector, error)

type ResourceProviderFactory struct {
	// detectors holds all possible detector types.
	detectors map[DetectorType]DetectorFactory
}

func NewProviderFactory(detectors map[DetectorType]DetectorFactory) *ResourceProviderFactory {
	_ = "STUB: not implemented"
	return nil
}

func (f *ResourceProviderFactory) CreateResourceProvider(
	params processor.Settings,
	timeout time.Duration,
	detectorConfigs ResourceDetectorConfig,
	detectorTypes ...DetectorType,
) (*ResourceProvider, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *ResourceProviderFactory) getDetectors(params processor.Settings, detectorConfigs ResourceDetectorConfig, detectorTypes []DetectorType) ([]Detector, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type ResourceProvider struct {
	logger           *zap.Logger
	timeout          time.Duration
	detectors        []Detector
	detectedResource atomic.Pointer[resourceResult]

	// Refresh loop control
	refreshInterval time.Duration
	stopCh          chan struct{}
	cancelFunc      context.CancelFunc
	wg              sync.WaitGroup
	startOnce       sync.Once
	stopOnce        sync.Once
}

type resourceResult struct {
	resource  pcommon.Resource
	schemaURL string
	err       error
}

func NewResourceProvider(logger *zap.Logger, timeout time.Duration, detectors ...Detector) *ResourceProvider {
	_ = "STUB: not implemented"
	return nil
}

// No periodic refresh by default

func (p *ResourceProvider) Get(_ context.Context, _ *http.Client) (pcommon.Resource, string, error) {
	_ = "STUB: not implemented"
	return *new(pcommon.Resource), "", nil
}

// Refresh recomputes the resource, replacing any previous result.
func (p *ResourceProvider) Refresh(ctx context.Context, client *http.Client) error {
	_ = "STUB: not implemented"
	return nil
}

// Check if we have a previous successful snapshot

// Keep the last good snapshot if the refresh errored.
// Note: An empty resource with no error is considered a success (e.g., detector determined
// it's not running on that cloud provider), so we accept it rather than keeping stale data.

// Return nil error since we're successfully keeping the cached resource

// Accept the new snapshot (even if empty, as long as there was no error).

func (p *ResourceProvider) detectResource(ctx context.Context) (pcommon.Resource, string, error) {
	_ = "STUB: not implemented"
	return *new(pcommon.Resource), "", nil
}

// retry

// If all detectors failed, return empty resource.

// Partial or full success: return merged resources.

func MergeSchemaURL(currentSchemaURL, newSchemaURL string) string {
	_ = "STUB: not implemented"
	return ""
}

// TODO: handle the case when the schema URLs are different by performing
// schema conversion. For now we simply ignore the new schema URL.

func MergeResource(to, from pcommon.Resource, overrideTo bool) { _ = "STUB: not implemented"; return }

func IsEmptyResource(res pcommon.Resource) bool { _ = "STUB: not implemented"; return false }

// StartRefreshing begins periodic resource refresh if refreshInterval > 0.
// It is safe to call multiple times; only the first call starts the goroutine.
func (p *ResourceProvider) StartRefreshing(refreshInterval time.Duration, client *http.Client) {
	_ = "STUB: not implemented"
	return
}

// StopRefreshing stops the periodic refresh goroutine.
// It is safe to call multiple times; only the first call stops the goroutine.
func (p *ResourceProvider) StopRefreshing() { _ = "STUB: not implemented"; return }

func (p *ResourceProvider) refreshLoop(ctx context.Context, client *http.Client) {
	_ = "STUB: not implemented"
	return
}

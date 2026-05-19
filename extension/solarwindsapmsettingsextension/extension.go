// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package solarwindsapmsettingsextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/solarwindsapmsettingsextension"

import (
	"context"
	"net/http"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension"
)

const (
	jsonOutputFile       = "solarwinds-apm-settings.json"
	httpsContextDeadline = 10 * time.Second
)

type arguments struct {
	BucketCapacity               float64 `json:"BucketCapacity"`
	BucketRate                   float64 `json:"BucketRate"`
	TriggerRelaxedBucketCapacity float64 `json:"TriggerRelaxedBucketCapacity"`
	TriggerRelaxedBucketRate     float64 `json:"TriggerRelaxedBucketRate"`
	TriggerStrictBucketCapacity  float64 `json:"TriggerStrictBucketCapacity"`
	TriggerStrictBucketRate      float64 `json:"TriggerStrictBucketRate"`
}
type setting struct {
	Value     int64     `json:"value"`
	Flags     string    `json:"flags"`
	Timestamp int64     `json:"timestamp"`
	TTL       int64     `json:"ttl"`
	Arguments arguments `json:"arguments"`
}

type settingWithWarning struct {
	setting
	Warning string `json:"warning"`
}

type solarwindsapmSettingsExtension struct {
	config            *Config
	cancel            context.CancelFunc
	telemetrySettings component.TelemetrySettings
	client            *http.Client
	request           *http.Request
}

func newSolarwindsApmSettingsExtension(extensionCfg *Config, settings extension.Settings) (extension.Extension, error) {
	_ = "STUB: not implemented"
	return *new(extension.Extension), nil
}

func (extension *solarwindsapmSettingsExtension) Start(_ context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

// initial refresh

func (extension *solarwindsapmSettingsExtension) Shutdown(_ context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func refresh(extension *solarwindsapmSettingsExtension, filename string) {
	_ = "STUB: not implemented"
	return
}

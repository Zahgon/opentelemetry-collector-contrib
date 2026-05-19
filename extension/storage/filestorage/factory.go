// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:generate make mdatagen

package filestorage // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/storage/filestorage"

import (
	"context"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension"
)

const (
	// use default bbolt value
	// https://github.com/etcd-io/bbolt/blob/d5db64bdbfdee1cb410894605f42ffef898f395d/cmd/bbolt/main.go#L1955
	defaultMaxTransactionSize         = 65536
	defaultReboundTriggerThresholdMib = 10
	defaultReboundNeededThresholdMib  = 100
	defaultCompactionInterval         = time.Second * 5
)

// NewFactory creates a factory for HostObserver extension.
func NewFactory() extension.Factory { _ = "STUB: not implemented"; return *new(extension.Factory) }

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

func createExtension(
	_ context.Context,
	params extension.Settings,
	cfg component.Config,
) (extension.Extension, error) {
	_ = "STUB: not implemented"
	return *new(extension.Extension), nil
}

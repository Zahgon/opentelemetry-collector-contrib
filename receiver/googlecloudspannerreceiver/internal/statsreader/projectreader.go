// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package statsreader // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudspannerreceiver/internal/statsreader"

import (
	"context"

	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudspannerreceiver/internal/metadata"
)

type ProjectReader struct {
	databaseReaders []CompositeReader
	logger          *zap.Logger
}

func NewProjectReader(databaseReaders []CompositeReader, logger *zap.Logger) *ProjectReader {
	_ = "STUB: not implemented"
	return nil
}

func (projectReader *ProjectReader) Shutdown() { _ = "STUB: not implemented"; return }

func (projectReader *ProjectReader) Read(ctx context.Context) ([]*metadata.MetricsDataPoint, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (projectReader *ProjectReader) Name() string { _ = "STUB: not implemented"; return "" }

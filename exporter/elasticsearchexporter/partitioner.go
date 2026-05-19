// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package elasticsearchexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/elasticsearchexporter"

import (
	"context"

	"go.opentelemetry.io/collector/exporter/exporterhelper/xexporterhelper"
)

type metadataKeysPartitioner struct {
	keys []string
}

func (p metadataKeysPartitioner) GetKey(
	ctx context.Context,
	_ xexporterhelper.Request,
) string {
	_ = "STUB: not implemented"
	return ""
}

func (p metadataKeysPartitioner) MergeCtx(
	ctx1, ctx2 context.Context,
) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// Since the mergeCtx is based on partition key, we MUST have the same
// partition key-values in both the metadata. If they are not same then
// fail fast and dramatically.

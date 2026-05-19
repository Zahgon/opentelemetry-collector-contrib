// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package batch // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awskinesisexporter/internal/batch"

import (
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awskinesisexporter/internal/key"
)

type batchMarshaller struct {
	batchOptions []Option
	partitioner  key.Partition

	logsMarshaller    plog.Marshaler
	tracesMarshaller  ptrace.Marshaler
	metricsMarshaller pmetric.Marshaler
}

var _ Encoder = (*batchMarshaller)(nil)

func (bm *batchMarshaller) Logs(ld plog.Logs) (*Batch, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Due to kinesis limitations of only allowing 1Mb of data per record,
	// the resource data is copied to the export variable then marshaled
	// due to no current means of marshaling per resource.
}

func (bm *batchMarshaller) Traces(td ptrace.Traces) (*Batch, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Due to kinesis limitations of only allowing 1Mb of data per record,
	// the resource data is copied to the export variable then marshaled
	// due to no current means of marshaling per resource.
}

func (bm *batchMarshaller) Metrics(md pmetric.Metrics) (*Batch, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Due to kinesis limitations of only allowing 1Mb of data per record,
	// the resource data is copied to the export variable then marshaled
	// due to no current means of marshaling per resource.
}

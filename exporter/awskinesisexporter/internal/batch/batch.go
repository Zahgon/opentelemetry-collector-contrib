// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package batch // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awskinesisexporter/internal/batch"

import (
	"errors"

	"github.com/aws/aws-sdk-go-v2/service/kinesis/types"
	"go.opentelemetry.io/collector/consumer/consumererror"
)

const (
	MaxRecordSize     = 1 << 20 // 1MiB
	MaxBatchedRecords = 500
)

var (
	// ErrPartitionKeyLength is used when the given key exceeds the allowed kinesis limit of 256 characters
	ErrPartitionKeyLength = errors.New("partition key size is greater than 256 characters")
	// ErrRecordLength is used when attempted record results in a byte array greater than 1MiB
	ErrRecordLength = consumererror.NewPermanent(errors.New("record size is greater than 1 MiB"))
)

type Batch struct {
	maxBatchSize  int
	maxRecordSize int

	compressionType string

	records []types.PutRecordsRequestEntry
}

type Option func(bt *Batch)

func WithMaxRecordsPerBatch(limit int) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithMaxRecordSize(size int) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithCompressionType(compressionType string) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func New(opts ...Option) *Batch { _ = "STUB: not implemented"; return nil }

func (b *Batch) AddRecord(raw []byte, key string) error { _ = "STUB: not implemented"; return nil }

// Chunk breaks up the internal queue into blocks that can be used
// to be written to he kinesis.PutRecords endpoint
func (b *Batch) Chunk() (chunks [][]types.PutRecordsRequestEntry) {
	_ = "STUB: not implemented"
	// Using local copies to avoid mutating internal data
	return nil
}

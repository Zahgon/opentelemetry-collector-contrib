// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package upload // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awss3exporter/internal/upload"

import (
	"time"

	"go.opentelemetry.io/collector/config/configcompression"
)

var compressionFileExtensions = map[configcompression.Type]string{
	configcompression.TypeGzip: ".gz",
	configcompression.TypeZstd: ".zst",
}

type PartitionKeyBuilder struct {
	// PartitionBasePrefix defines the root S3
	// directory (key) prefix used to write the file.
	PartitionBasePrefix string
	// PartitionPrefix defines the S3 directory (key)
	// prefix used to write the file.
	// Appended to PartitionBasePrefix if provided.
	PartitionPrefix string
	// PartitionFormat is used to separate values into
	// different time buckets.
	// Uses [strftime](https://www.man7.org/linux/man-pages/man3/strftime.3.html) formatting.
	PartitionFormat string
	// PartitionTimeLocation is used to provide timezone for partition time. Defaults to Local time location.
	PartitionTimeLocation *time.Location
	// FilePrefix is used to define the prefix of the file written
	// to the directory in S3.
	FilePrefix string
	// FileFormat defines what encoding was used to write
	// the content to s3
	FileFormat string
	// Metadata provides additional details regarding the file
	// Expected to be one of "metrics", "traces", or "logs"
	Metadata string
	// Compression defines algorithm used on the
	// body before upload.
	Compression configcompression.Type
	// UniqueKeyFunc allows for overwriting the default behavior of
	// generating a new unique string to avoid collisions on file upload
	// across many different instances.
	UniqueKeyFunc func() string
	// IsCompressed when true keeps files compressed in S3
	// by omitting ContentEncoding headers. When false, ContentEncoding
	// is set for HTTP transfer compression (AWS auto-decompresses).
	IsCompressed bool
}

func (pki *PartitionKeyBuilder) Build(ts time.Time, overridePrefix string) string {
	_ = "STUB: not implemented"
	return ""
}

func (pki *PartitionKeyBuilder) bucketKeyPrefix(ts time.Time, overridePrefix string) string {
	_ = "STUB: not implemented"
	// Don't want to overwrite the actual value
	return ""
}

// Only override when it's not empty string

func (pki *PartitionKeyBuilder) fileName() string { _ = "STUB: not implemented"; return "" }

func (pki *PartitionKeyBuilder) uniqueKey() string {
	_ = "STUB: not implemented"
	// If a custom function is provided, use it to generate the unique key.
	// If it fails, fall back to the default random integer generation
	// so that uploads are not blocked.
	return ""
}

func GenerateUUIDv7() string { _ = "STUB: not implemented"; return "" }

func (*PartitionKeyBuilder) randInt() string {
	_ = "STUB: not implemented"
	// This follows the original "uniqueness" algorithm
	// to avoid collisions on file uploads across different nodes.
	return ""
}

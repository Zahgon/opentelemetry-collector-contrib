// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package elasticbeanstalk // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/aws/elasticbeanstalk"

import (
	"io"
)

type fileSystem interface {
	Open(name string) (io.ReadCloser, error)
	IsWindows() bool
}

type ebFileSystem struct{}

func (ebFileSystem) Open(name string) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

func (ebFileSystem) IsWindows() bool { _ = "STUB: not implemented"; return false }

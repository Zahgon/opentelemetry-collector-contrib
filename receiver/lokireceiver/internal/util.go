// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package internal // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/lokireceiver/internal"

import (
	"bytes"
	"io"

	"github.com/gogo/protobuf/proto"
)

const messageSizeLargerErrFmt = "received message larger than max (%d vs %d)"

// parseProtoReader parses a compressed proto from an io.Reader.
func parseProtoReader(reader io.Reader, expectedSize, maxSize int, req proto.Message) error {
	_ = "STUB: not implemented"
	return nil
}

// We re-implement proto.Unmarshal here as it calls XXX_Unmarshal first,
// which we can't override without upsetting golint.

func decompressRequest(reader io.Reader, expectedSize, maxSize int) (body []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func decompressFromReader(reader io.Reader, expectedSize, maxSize int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// extra space guarantees no reallocation

// Read from LimitReader with limit max+1. So if the underlying
// reader is over limit, the result will be bigger than max.

func decompressFromBuffer(buffer *bytes.Buffer, maxSize int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// tryBufferFromReader attempts to cast the reader to a `*bytes.Buffer` this is possible when using httpgrpc.
// If it fails it will return nil and false.
func tryBufferFromReader(reader io.Reader) (*bytes.Buffer, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

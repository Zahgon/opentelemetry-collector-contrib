// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package serializeprofiles // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/elasticsearchexporter/internal/serializer/otelserializer/serializeprofiles"

import (
	"go.opentelemetry.io/ebpf-profiler/libpf"
)

// frameID represents a frame as an address in an executable file
// or as a line in a source code file.
type frameID struct {
	// fileID is the fileID of the frame
	fileID libpf.FileID

	// addressOrLineno is the address or lineno of the frame
	addressOrLineno libpf.AddressOrLineno
}

// newFrameID creates a new FrameID from the fileId and address or line.
func newFrameID(fileID libpf.FileID, addressOrLineno libpf.AddressOrLineno) frameID {
	_ = "STUB: not implemented"
	return *new(frameID)
}

// newFrameIDFromString creates a new FrameID from its base64 string representation.
func newFrameIDFromString(frameEncoded string) (frameID, error) {
	_ = "STUB: not implemented"
	return *new(frameID), nil
}

// newFrameIDFromBytes creates a new FrameID from a byte array of length 24.
func newFrameIDFromBytes(bytes []byte) (frameID, error) {
	_ = "STUB: not implemented"
	return *new(frameID), nil
}

// Bytes returns the frameid as byte sequence.
func (f frameID) Bytes() []byte {
	_ = "STUB: not implemented"
	// Using frameID := make([byte, 24]) here makes the function ~5% slower.
	return nil
}

// String returns the base64 encoded representation.
func (f frameID) String() string { _ = "STUB: not implemented"; return "" }

// FileID returns the fileID part of the frameID.
func (f frameID) FileID() libpf.FileID {
	_ = "STUB: not implemented"

	// AddressOrLine returns the addressOrLine part of the frameID.
	return *new(libpf.FileID)
}

func (f frameID) AddressOrLine() libpf.AddressOrLineno {
	_ = "STUB: not implemented"
	return *new(libpf.AddressOrLineno)
}

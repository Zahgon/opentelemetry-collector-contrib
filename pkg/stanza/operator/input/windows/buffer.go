// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package windows // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/input/windows"

// defaultBufferSize is the default size of the buffer.
const defaultBufferSize = 16384

// bytesPerWChar is the number bytes in a Windows wide character.
const bytesPerWChar = 2

// Buffer is a buffer of utf-16 bytes.
type Buffer struct {
	buffer []byte
}

// ReadBytes will read UTF-8 bytes from the buffer, where offset is the number of bytes to be read
func (b *Buffer) ReadBytes(offset uint32) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ReadWideChars will read UTF-8 bytes from the buffer, where offset is the number of wchars to read
func (b *Buffer) ReadWideChars(offset uint32) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ReadString will read a UTF-8 string from the buffer.
func (b *Buffer) ReadString(offset uint32) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// UpdateSizeBytes will update the size of the buffer to fit size bytes.
func (b *Buffer) UpdateSizeBytes(size uint32) { _ = "STUB: not implemented"; return }

// UpdateSizeWide will update the size of the buffer to fit size wchars.
func (b *Buffer) UpdateSizeWide(size uint32) { _ = "STUB: not implemented"; return }

// SizeBytes will return the size of the buffer as number of bytes.
func (b *Buffer) SizeBytes() uint32 { _ = "STUB: not implemented"; return 0 }

// SizeWide returns the size of the buffer as number of wchars
func (b *Buffer) SizeWide() uint32 { _ = "STUB: not implemented"; return 0 }

// FirstByte will return a pointer to the first byte.
func (b *Buffer) FirstByte() *byte { _ = "STUB: not implemented"; return nil }

// NewBuffer creates a new buffer with the default buffer size
func NewBuffer() *Buffer { _ = "STUB: not implemented"; return nil }

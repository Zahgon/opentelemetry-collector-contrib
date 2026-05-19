// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build windows

package windows // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/input/windows"

// Bookmark is a windows event bookmark.
type Bookmark struct {
	handle uintptr
}

// Open will open the bookmark handle using the supplied xml.
func (b *Bookmark) Open(offsetXML string) error { _ = "STUB: not implemented"; return nil }

// Update will update the bookmark using the supplied event.
func (b *Bookmark) Update(event Event) error { _ = "STUB: not implemented"; return nil }

// Render will render the bookmark as xml.
func (b *Bookmark) Render(buffer *Buffer) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Close will close the bookmark handle.
func (b *Bookmark) Close() error { _ = "STUB: not implemented"; return nil }

// NewBookmark will create a new bookmark with an empty handle.
func NewBookmark() Bookmark { _ = "STUB: not implemented"; return *new(Bookmark) }

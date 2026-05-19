// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package textutils // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal/textutils"

import (
	"golang.org/x/text/encoding"
)

// UTF8Raw is a variant of the UTF-8 encoding without replacing invalid UTF-8 sequences.
// It behaves in the same way as [encoding.Nop], but is differentiated from nop encoding, which we treat in a special way.
var UTF8Raw encoding.Encoding = utf8raw{}

type utf8raw struct{}

func (utf8raw) NewDecoder() *encoding.Decoder { _ = "STUB: not implemented"; return nil }

func (utf8raw) NewEncoder() *encoding.Encoder { _ = "STUB: not implemented"; return nil }

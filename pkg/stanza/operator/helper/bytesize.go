// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package helper // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"

import (
	"regexp"
)

type ByteSize int64

var byteSizeRegex = regexp.MustCompile(`^(\d+\.?\d*)\s*([kKmMgGtTpP]i?[bB])?$`)

func (h *ByteSize) UnmarshalText(text []byte) (err error) { _ = "STUB: not implemented"; return nil }

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package fluentforwardreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/fluentforwardreceiver"

import (
	"time"

	"github.com/tinylib/msgp/msgp"
)

type eventTimeExt time.Time

func init() {
	msgp.RegisterExtension(0, func() msgp.Extension { return new(eventTimeExt) })
}

func (*eventTimeExt) ExtensionType() int8 { _ = "STUB: not implemented"; return 0 }

func (*eventTimeExt) Len() int { _ = "STUB: not implemented"; return 0 }

func (e *eventTimeExt) MarshalBinaryTo(b []byte) error { _ = "STUB: not implemented"; return nil }

func (e *eventTimeExt) UnmarshalBinary(b []byte) error { _ = "STUB: not implemented"; return nil }

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package eventtime // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/libhoneyreceiver/internal/eventtime"

import (
	"time"
)

func GetNowTime() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func GetEventTime(etHeader string) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// Great, they sent us a time header. let's try and parse it.
// RFC3339Nano is the default that we send from all our SDKs

// the default didn't catch it, let's try a few other things
// is it all numeric? then try unix epoch times

// it might be seconds or it might be milliseconds! Who can know!
// 10-digit numbers are seconds, 13-digit milliseconds, 16 microseconds

// turn it into seconds and fractional seconds

// then chop it into the int part and the fractional part

func GetEventTimeSec(etHeader string) int64 { _ = "STUB: not implemented"; return 0 }

func GetEventTimeNano(etHeader string) int64 { _ = "STUB: not implemented"; return 0 }

func GetEventTimeDefaultString() string { _ = "STUB: not implemented"; return "" }

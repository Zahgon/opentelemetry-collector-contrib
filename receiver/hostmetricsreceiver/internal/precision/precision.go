// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package precision // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/precision"

import (
	"time"
)

// Ratio computes numerator/denominator and rounds the result to the
// number of significant digits supported by the inputs. The significant digit
// count is derived from the magnitude of the larger operand, matching the
// information content of the integer inputs. When denominator is zero the
// native Go float64 division result is returned (NaN for 0/0, +Inf otherwise).
func Ratio(numerator, denominator uint64) float64 { _ = "STUB: not implemented"; return 0 }

// Scale converts a tick count in the given unit to seconds and rounds
// to the unit's decimal precision. This avoids binary float artifacts
// like 12345/1000 = 12.345000000000001.
func Scale(numerator uint64, unit time.Duration) float64 { _ = "STUB: not implemented"; return 0 }

func roundRatio(numerator, denominator float64) float64 { _ = "STUB: not implemented"; return 0 }

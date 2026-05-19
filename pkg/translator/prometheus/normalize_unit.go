// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package prometheus // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/prometheus"

// The map to translate OTLP units to Prometheus units
// OTLP metrics use the c/s notation as specified at https://ucum.org/ucum.html
// (See also https://github.com/open-telemetry/semantic-conventions/blob/main/docs/general/metrics.md#instrument-units)
// Prometheus best practices for units: https://prometheus.io/docs/practices/naming/#base-units
// OpenMetrics specification for units: https://github.com/prometheus/OpenMetrics/blob/v1.0.0/specification/OpenMetrics.md#units-and-base-units
var unitMap = map[string]string{
	// Time
	"d":   "days",
	"h":   "hours",
	"min": "minutes",
	"s":   "seconds",
	"ms":  "milliseconds",
	"us":  "microseconds",
	"ns":  "nanoseconds",

	// Bytes
	"By":   "bytes",
	"KiBy": "kibibytes",
	"MiBy": "mebibytes",
	"GiBy": "gibibytes",
	"TiBy": "tibibytes",
	"KBy":  "kilobytes",
	"MBy":  "megabytes",
	"GBy":  "gigabytes",
	"TBy":  "terabytes",

	// SI
	"m": "meters",
	"V": "volts",
	"A": "amperes",
	"J": "joules",
	"W": "watts",
	"g": "grams",

	// Misc
	"Cel": "celsius",
	"Hz":  "hertz",
	"1":   "",
	"%":   "percent",
}

// The map that translates the "per" unit
// Example: s => per second (singular)
var perUnitMap = map[string]string{
	"s":  "second",
	"m":  "minute",
	"h":  "hour",
	"d":  "day",
	"w":  "week",
	"mo": "month",
	"y":  "year",
}

func BuildCompliantPrometheusUnit(unit string) string { _ = "STUB: not implemented"; return "" }

// Extract the main unit from an OTLP unit and convert to Prometheus base unit
// Returns an empty string if the unit is not found in the map
func buildCompliantMainUnit(unit string) string { _ = "STUB: not implemented"; return "" }

// Extract the rate unit from an OTLP unit and convert to Prometheus base unit
// Returns an empty string if the unit is not found in the map
func buildCompliantPerUnit(unit string) string { _ = "STUB: not implemented"; return "" }

// Retrieve the Prometheus "basic" unit corresponding to the specified "basic" unit
// Returns the specified unit if not found in unitMap
func unitMapGetOrDefault(unit string) string { _ = "STUB: not implemented"; return "" }

// Retrieve the Prometheus "per" unit corresponding to the specified "per" unit
// Returns the specified unit if not found in perUnitMap
func perUnitMapGetOrDefault(perUnit string) string { _ = "STUB: not implemented"; return "" }

// Clean up specified string so it's Prometheus compliant
func CleanUpString(s string) string { _ = "STUB: not implemented"; return "" }

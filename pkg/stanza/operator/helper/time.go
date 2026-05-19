// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package helper // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"

import (
	"regexp"
	"time"

	"go.opentelemetry.io/collector/confmap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/entry"
)

// tzAbbrRegex matches 2-5 consecutive uppercase ASCII letters (timezone abbreviations like IST, NZST, PDT)
var tzAbbrRegex = regexp.MustCompile(`\b([A-Z]{2,5})\b`)

// StrptimeKey is literally "strptime", and is the default layout type
const StrptimeKey = "strptime"

// GotimeKey is literally "gotime" and uses Golang's native time.Parse
const GotimeKey = "gotime"

// EpochKey is literally "epoch" and can parse seconds and/or subseconds
const EpochKey = "epoch"

// NativeKey is literally "native" and refers to Golang's native time.Time
const NativeKey = "native" // provided for operator development

// NewTimeParser creates a new time parser with default values
func NewTimeParser() TimeParser { _ = "STUB: not implemented"; return *new(TimeParser) }

// TimeParser is a helper that parses time onto an entry.
type TimeParser struct {
	ParseFrom         *entry.Field      `mapstructure:"parse_from"`
	Layout            string            `mapstructure:"layout"`
	LayoutType        string            `mapstructure:"layout_type"`
	Location          string            `mapstructure:"location"`
	TimeZoneLocations map[string]string `mapstructure:"time_zone_locations"` // optional: abbreviation → IANA location name

	location    *time.Location
	locationMap map[string]*time.Location // compiled from TimeZoneLocations at Validate() time
}

// Unmarshal starting from default settings
func (t *TimeParser) Unmarshal(component *confmap.Conf) error {
	_ = "STUB: not implemented"
	return nil
}

// IsZero returns true if the TimeParser is not a valid config
func (t *TimeParser) IsZero() bool { _ = "STUB: not implemented"; return false }

// Validate validates a TimeParser, and reconfigures it if necessary
func (t *TimeParser) Validate() error { _ = "STUB: not implemented"; return nil }

// ok

// ok

// also covers StrptimeKey because it was remapped above

func (t *TimeParser) setLocation() error { _ = "STUB: not implemented"; return nil }

// If "location" is specified, it must be in the local timezone database

// If a timestamp ends with 'Z', it should be interpreted at Zulu (UTC) time

// Compile time_zone_locations at startup so LoadLocation is never called per log line

// resolveLocation returns the *time.Location to use for parsing a given value.
// If time_zone_locations is configured, it scans the value for a known timezone abbreviation
// and returns the corresponding location. Falls back to t.location if not found.
func (t *TimeParser) resolveLocation(value any) *time.Location {
	_ = "STUB: not implemented"
	return nil
}

// Find the first uppercase token that matches a configured abbreviation

// Parse will parse time from a field and attach it to the entry
func (t *TimeParser) Parse(entry *entry.Entry) error { _ = "STUB: not implemented"; return nil }

// timeutils.ParseGotime calls timeutils.SetTimestampYear before returning the timeValue

func (t *TimeParser) parseEpochTime(value any) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

func getEpochStamp(layout string, value any) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

type toTimeFunc = func(int64) time.Time

var toTime = map[string]toTimeFunc{
	"s":  func(s int64) time.Time { return time.Unix(s, 0) },
	"ms": func(ms int64) time.Time { return time.Unix(ms/1e3, (ms%1e3)*1e6) },
	"us": func(us int64) time.Time { return time.Unix(us/1e6, (us%1e6)*1e3) },
	"ns": func(ns int64) time.Time { return time.Unix(0, ns) },
}
var subsecToNs = map[string]int64{"s.ms": 1e6, "s.us": 1e3, "s.ns": 1}

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package timeutils // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal/timeutils"

import (
	"regexp"
	"time"
)

var invalidFractionalSecondsGoTime = regexp.MustCompile(`[^.,9]9+`)

func StrptimeToGotime(layout string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func ParseStrptime(layout string, value any, location *time.Location) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

// ParseLocalizedStrptime is like ParseLocalizedGotime, but instead of using the native Go time layout,
// it uses the ctime-like format.
func ParseLocalizedStrptime(layout string, value any, location *time.Location, language string) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

func GetLocation(location, layout *string) (*time.Location, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If location is specified, it must be in the local timezone database

// If a timestamp ends with 'Z', it should be interpreted at Zulu (UTC) time

// ParseLocalizedGotime is like ParseGotime, but instead of parsing a formatted time in
// English, it parses a value in foreign language, and returns the [time.Time] it represents.
// The language argument must be a well-formed BCP 47 language tag (e.g.: "en", "en-US"), and
// a known CLDR locale.
func ParseLocalizedGotime(layout string, value any, location *time.Location, language string) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

func ParseGotime(layout string, value any, location *time.Location) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

func parseGotime(layout string, value any, location *time.Location) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

// Depending on the timezone database, we may get a pseudo-matching timezone
// This is apparent when the zone is not "UTC", but the offset is still 0.
// We skip the correction if the zone matches the caller's default location,
// because in that case the zone name came from the default (e.g. "WET" for
// time.Local on a Western European system), not from an abbreviation in the
// input string.

// Manually look up the location based on the zone

// can't correct offset, just return what we have

// Reparse the timestamp, with the location

// can't correct offset, just return original result

func convertParsingValue(value any) (string, error) { _ = "STUB: not implemented"; return "", nil }

// SetTimestampYear sets the year of a timestamp to the current year.
// This is needed because year is missing from some time formats, such as rfc3164.
func SetTimestampYear(t time.Time) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// Assume the timestamp is from last year if its month and day are
// more than 7 days past the current date.
// i.e. If today is January 1, but the timestamp is February 1, it's safe
// to assume the timestamp is from last year.

// ValidateStrptime checks the given strptime layout and returns an error if it detects any known issues
// that prevent it from being parsed.
func ValidateStrptime(layout string) error { _ = "STUB: not implemented"; return nil }

func ValidateGotime(layout string) error { _ = "STUB: not implemented"; return nil }

// ValidateLocale checks the given locale and returns an error if the language tag
// is not supported by the localized parser functions.
func ValidateLocale(locale string) error { _ = "STUB: not implemented"; return nil }

// GetStrptimeNativeSubstitutes analyzes the provided format string and returns a map
// where each key is a Go native layout element (as used in time.Format) found in the
// format, and each value is the corresponding ctime-like directive.
func GetStrptimeNativeSubstitutes(format string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

type strptimeParseErr struct {
	err               *time.ParseError
	ctimeLayout       string
	nativeSubstitutes map[string]string
}

func (e *strptimeParseErr) Error() string { _ = "STUB: not implemented"; return "" }

func ToStrptimeParseError(err *time.ParseError, ctimeLayout string, nativeSubstitutes map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

// Allows tests to override with deterministic value
var Now = time.Now

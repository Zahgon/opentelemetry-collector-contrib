// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package pprofiletest // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest/pprofiletest"
import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pprofile"
)

func ValidateProfile(dic pprofile.ProfilesDictionary, pp pprofile.Profile) error {
	_ = "STUB: not implemented"
	return nil
}

// Return here to avoid panicking when accessing the string table.

func validateIndices(length int, indices pcommon.Int32Slice) error {
	_ = "STUB: not implemented"
	return nil
}

func validateIndex(length int, idx int32) error { _ = "STUB: not implemented"; return nil }

func validateSampleType(dic pprofile.ProfilesDictionary, pp pprofile.Profile) error {
	_ = "STUB: not implemented"
	return nil
}

func validateValueType(stLen int, pvt pprofile.ValueType) error {
	_ = "STUB: not implemented"
	return nil
}

func validateSamples(dic pprofile.ProfilesDictionary, pp pprofile.Profile) error {
	_ = "STUB: not implemented"
	return nil
}

func validateSample(dic pprofile.ProfilesDictionary, pp pprofile.Profile, sample pprofile.Sample) error {
	_ = "STUB: not implemented"
	return nil
}

func validateLocation(dic pprofile.ProfilesDictionary, loc pprofile.Location) error {
	_ = "STUB: not implemented"
	return nil
}

// Continuing would run into a panic.

func validateLine(dic pprofile.ProfilesDictionary, line pprofile.Line) error {
	_ = "STUB: not implemented"
	return nil
}

func validateMapping(dic pprofile.ProfilesDictionary, mapping pprofile.Mapping) error {
	_ = "STUB: not implemented"
	return nil
}

func validateKeyValueAndUnits(dic pprofile.ProfilesDictionary) error {
	_ = "STUB: not implemented"
	return nil
}

func validateKeyValueAndUnit(dic pprofile.ProfilesDictionary, au pprofile.KeyValueAndUnit) error {
	_ = "STUB: not implemented"
	return nil
}

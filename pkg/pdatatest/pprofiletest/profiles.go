// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package pprofiletest // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest/pprofiletest"

import (
	"go.opentelemetry.io/collector/pdata/pprofile"
)

// CompareProfiles compares each part of two given Profiles and returns
// an error if they don't match. The error describes what didn't match.
func CompareProfiles(expected, actual pprofile.Profiles, options ...CompareProfilesOption) error {
	_ = "STUB: not implemented"
	return nil
}

// Keep track of matching resources so that each can only be matched once

// CompareResourceProfiles compares each part of two given ResourceProfiles and returns
// an error if they don't match. The error describes what didn't match.
func CompareResourceProfiles(expectedDic, actualDic pprofile.ProfilesDictionary, expected, actual pprofile.ResourceProfiles) error {
	_ = "STUB: not implemented"
	return nil
}

// Keep track of matching scope profiles so that each container can only be matched once

// CompareScopeProfiles compares each part of two given ProfilesSlices and returns
// an error if they don't match. The error describes what didn't match.
func CompareScopeProfiles(expectedDic, actualDic pprofile.ProfilesDictionary, expected, actual pprofile.ScopeProfiles) error {
	_ = "STUB: not implemented"
	return nil
}

// Keep track of matching containers so that each container can only be matched once

func compareAttributes(expectedDic, actualDic pprofile.ProfilesDictionary, a, b pprofile.Profile) error {
	_ = "STUB: not implemented"
	return nil
}

func CompareProfile(expectedDic, actualDic pprofile.ProfilesDictionary, expected, actual pprofile.Profile) error {
	_ = "STUB: not implemented"
	return nil
}

func CompareProfileValueType(expected, actual pprofile.ValueType) error {
	_ = "STUB: not implemented"
	return nil
}

func isValueTypeEqual(expected, actual pprofile.ValueType) bool {
	_ = "STUB: not implemented"
	return false
}

func CompareProfileSampleSlice(expected, actual pprofile.SampleSlice) error {
	_ = "STUB: not implemented"
	return nil
}

func CompareProfileSample(expected, actual pprofile.Sample) error {
	_ = "STUB: not implemented"
	return nil
}

func CompareProfileMappingSlice(expected, actual pprofile.MappingSlice) error {
	_ = "STUB: not implemented"
	return nil
}

func CompareProfileFunctionSlice(expected, actual pprofile.FunctionSlice) error {
	_ = "STUB: not implemented"
	return nil
}

func isFunctionEqual(expected, actual pprofile.Function) bool {
	_ = "STUB: not implemented"
	return false
}

func CompareProfileLocationSlice(expected, actual pprofile.LocationSlice) error {
	_ = "STUB: not implemented"
	return nil
}

func CompareProfileLocation(expected, actual pprofile.Location) error {
	_ = "STUB: not implemented"
	return nil
}

func CompareProfileLineSlice(expected, actual pprofile.LineSlice) error {
	_ = "STUB: not implemented"
	return nil
}

func isLineEqual(expected, actual pprofile.Line) bool { _ = "STUB: not implemented"; return false }

func CompareKeyValueAndUnitSlice(expected, actual pprofile.KeyValueAndUnitSlice) error {
	_ = "STUB: not implemented"
	return nil
}

func CompareProfileLinkSlice(expected, actual pprofile.LinkSlice) error {
	_ = "STUB: not implemented"
	return nil
}

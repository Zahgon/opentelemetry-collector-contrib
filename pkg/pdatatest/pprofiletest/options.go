// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package pprofiletest // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest/pprofiletest"

import (
	"go.opentelemetry.io/collector/pdata/pprofile"
)

// CompareProfilesOption can be used to mutate expected and/or actual profiles before comparing.
type CompareProfilesOption interface {
	applyOnProfiles(expected, actual pprofile.Profiles)
}

type compareProfilesOptionFunc func(expected, actual pprofile.Profiles)

func (f compareProfilesOptionFunc) applyOnProfiles(expected, actual pprofile.Profiles) {
	_ = "STUB: not implemented"
	return

	// IgnoreResourceAttributeValue is a CompareProfilesOption that removes a resource attribute
	// from all resources.
}

func IgnoreResourceAttributeValue(attributeName string) CompareProfilesOption {
	_ = "STUB: not implemented"
	return *new(CompareProfilesOption)
}

type ignoreResourceAttributeValue struct {
	attributeName string
}

func (opt ignoreResourceAttributeValue) applyOnProfiles(expected, actual pprofile.Profiles) {
	_ = "STUB: not implemented"
	return
}

func (opt ignoreResourceAttributeValue) maskProfilesResourceAttributeValue(profiles pprofile.Profiles) {
	_ = "STUB: not implemented"
	return
}

// IgnoreResourceEntityRefs is a CompareProfilesOption that clears entity references
// on all resources.
func IgnoreResourceEntityRefs() CompareProfilesOption {
	_ = "STUB: not implemented"
	return *new(CompareProfilesOption)
}

func maskProfilesResourceEntityRefs(profiles pprofile.Profiles) { _ = "STUB: not implemented"; return }

// IgnoreScopeAttributeValue is a CompareProfilesOption that removes a scope attribute
// from all resources.
func IgnoreScopeAttributeValue(attributeName string) CompareProfilesOption {
	_ = "STUB: not implemented"
	return *new(CompareProfilesOption)
}

type ignoreScopeAttributeValue struct {
	attributeName string
}

func (opt ignoreScopeAttributeValue) applyOnProfiles(expected, actual pprofile.Profiles) {
	_ = "STUB: not implemented"
	return
}

func (opt ignoreScopeAttributeValue) maskProfilesScopeAttributeValue(profiles pprofile.Profiles) {
	_ = "STUB: not implemented"
	return
}

// IgnoreProfileAttributeValue is a CompareProfilesOption that sets the value of an attribute
// to empty bytes for every profile
func IgnoreProfileAttributeValue(attributeName string) CompareProfilesOption {
	_ = "STUB: not implemented"
	return *new(CompareProfilesOption)
}

type ignoreProfileAttributeValue struct {
	attributeName string
}

func (opt ignoreProfileAttributeValue) applyOnProfiles(expected, actual pprofile.Profiles) {
	_ = "STUB: not implemented"
	return
}

func (opt ignoreProfileAttributeValue) maskProfileAttributeValue(profiles pprofile.Profiles) {
	_ = "STUB: not implemented"
	return
}

// IgnoreProfileTimestampValues is a CompareProfilesOption that sets the value of start timestamp
// and duration to empty bytes for every profile
func IgnoreProfileTimestampValues() CompareProfilesOption {
	_ = "STUB: not implemented"
	return *new(CompareProfilesOption)
}

type ignoreProfileTimestampValues struct{}

func (opt ignoreProfileTimestampValues) applyOnProfiles(expected, actual pprofile.Profiles) {
	_ = "STUB: not implemented"
	return
}

func (ignoreProfileTimestampValues) maskProfileTimestampValues(profiles pprofile.Profiles) {
	_ = "STUB: not implemented"
	return
}

// IgnoreResourceProfilesOrder is a CompareProfilesOption that ignores the order of resource traces/metrics/profiles.
func IgnoreResourceProfilesOrder() CompareProfilesOption {
	_ = "STUB: not implemented"
	return *new(CompareProfilesOption)
}

func sortResourceProfilesSlice(rls pprofile.ResourceProfilesSlice) {
	_ = "STUB: not implemented"
	return
}

// IgnoreScopeProfilesOrder is a CompareProfilesOption that ignores the order of instrumentation scope traces/metrics/profiles.
func IgnoreScopeProfilesOrder() CompareProfilesOption {
	_ = "STUB: not implemented"
	return *new(CompareProfilesOption)
}

func sortScopeProfilesSlices(ls pprofile.Profiles) { _ = "STUB: not implemented"; return }

// IgnoreProfilesOrder is a CompareProfilesOption that ignores the order of profile records.
func IgnoreProfilesOrder() CompareProfilesOption {
	_ = "STUB: not implemented"
	return *new(CompareProfilesOption)
}

func sortProfileSlices(ls pprofile.Profiles) { _ = "STUB: not implemented"; return }

func profileAttributesToMap(dic pprofile.ProfilesDictionary, p pprofile.Profile) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

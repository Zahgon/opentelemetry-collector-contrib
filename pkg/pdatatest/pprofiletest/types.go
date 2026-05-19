// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package pprofiletest // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest/pprofiletest"
import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pprofile"
)

type Profiles struct {
	ResourceProfiles []ResourceProfile
	// prevent unkeyed literal initialization
	_ struct{}
}

func (p Profiles) Transform() pprofile.Profiles {
	_ = "STUB: not implemented"
	return *new(pprofile.Profiles)
}

type ResourceProfile struct {
	ScopeProfiles []ScopeProfile
	Resource      pcommon.Resource
	SchemaURL     string
	// prevent unkeyed literal initialization
	_ struct{}
}

func (rp ResourceProfile) Transform(pp pprofile.Profiles) pprofile.ResourceProfiles {
	_ = "STUB: not implemented"
	return *new(pprofile.ResourceProfiles)
}

type ScopeProfile struct {
	Profiles  []Profile
	Scope     pcommon.InstrumentationScope
	SchemaURL string
	// prevent unkeyed literal initialization
	_ struct{}
}

func (sp ScopeProfile) Transform(dic pprofile.ProfilesDictionary, prp pprofile.ResourceProfiles) pprofile.ScopeProfiles {
	_ = "STUB: not implemented"
	return *new(pprofile.ScopeProfiles)
}

type Profile struct {
	SampleType             ValueType
	Sample                 []Sample
	TimeNanos              pcommon.Timestamp
	DurationNanos          uint64
	PeriodType             ValueType
	Period                 int64
	ProfileID              pprofile.ProfileID
	DroppedAttributesCount uint32
	OriginalPayloadFormat  string
	OriginalPayload        []byte
	Attributes             []Attribute
	KeyValueAndUnits       []KeyValueAndUnit
}

func (p *Profile) Transform(dic pprofile.ProfilesDictionary, psp pprofile.ScopeProfiles) pprofile.Profile {
	_ = "STUB: not implemented"
	return *new(pprofile.Profile)
}

// Avoids that 0 (default) string indices point to nowhere.

// If valueTypes are not set, set them to the default value.

func addString(dic pprofile.ProfilesDictionary, s string) int32 {
	_ = "STUB: not implemented"
	return 0
}

type ValueType struct {
	Typ  string
	Unit string
}

func (vt *ValueType) exists(dic pprofile.ProfilesDictionary, pp pprofile.Profile) bool {
	_ = "STUB: not implemented"
	return false
}

func (vt *ValueType) CopyTo(dic pprofile.ProfilesDictionary, pvt pprofile.ValueType) {
	_ = "STUB: not implemented"
	return
}

func (vt *ValueType) Transform(dic pprofile.ProfilesDictionary, pp pprofile.Profile) {
	_ = "STUB: not implemented"
	return
}

type Sample struct {
	Link               *Link // optional
	Values             []int64
	Locations          []Location
	Attributes         []Attribute
	TimestampsUnixNano []uint64
	// prevent unkeyed literal initialization
	_ struct{}
}

func (sa *Sample) Transform(dic pprofile.ProfilesDictionary, pp pprofile.Profile) {
	_ = "STUB: not implemented"
	return
}

//nolint:revive,staticcheck

// psa.SetLinkIndex(sa.Link.Transform(pp)) <-- undefined yet

type Location struct {
	Mapping    *Mapping
	Address    uint64
	Line       []Line
	IsFolded   bool
	Attributes []Attribute
	// prevent unkeyed literal initialization
	_ struct{}
}

type Link struct {
	TraceID pcommon.TraceID
	SpanID  pcommon.SpanID
	// prevent unkeyed literal initialization
	_ struct{}
}

func (l *Link) Transform(dic pprofile.ProfilesDictionary) int32 {
	_ = "STUB: not implemented"
	return 0
}

type Mapping struct {
	MemoryStart uint64
	MemoryLimit uint64
	FileOffset  uint64
	Filename    string
	Attributes  []Attribute
	// prevent unkeyed literal initialization
	_ struct{}
}

func (m *Mapping) Transform(dic pprofile.ProfilesDictionary) { _ = "STUB: not implemented"; return }

type Attribute struct {
	Key   string
	Value any
	// prevent unkeyed literal initialization
	_ struct{}
}

type attributable interface {
	AttributeIndices() pcommon.Int32Slice
}

func (a *Attribute) Transform(dic pprofile.ProfilesDictionary, record attributable) {
	_ = "STUB: not implemented"
	return
}

type KeyValueAndUnit struct {
	Key   string
	Value any
	Unit  string
	// prevent unkeyed literal initialization
	_ struct{}
}

func (a *KeyValueAndUnit) Transform(dic pprofile.ProfilesDictionary) int32 {
	_ = "STUB: not implemented"
	return 0
}

type Line struct {
	Line     int64
	Column   int64
	Function Function
	// prevent unkeyed literal initialization
	_ struct{}
}

type Function struct {
	Name       string
	SystemName string
	Filename   string
	StartLine  int64
	// prevent unkeyed literal initialization
	_ struct{}
}

func (f *Function) Transform(dic pprofile.ProfilesDictionary) int32 {
	_ = "STUB: not implemented"
	return 0
}

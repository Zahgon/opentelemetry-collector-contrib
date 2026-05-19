// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package translation // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/schemaprocessor/internal/translation"

import (
	"errors"
)

const separator = "."

var (
	ErrInvalidFamily  = errors.New("invalid schema family")
	ErrInvalidVersion = errors.New("invalid schema version")
)

// The following status values define whether the transformer
// has to revert changes or update changes to the signal being modified
// These values match the Version.Compare out when performing:
//
//	from.Compare(to) // ie: (1.0.0).Compare(1.2.1)
const (
	Update   int = -1 // From is less than To
	NoChange int = 0  // From equals To
	Revert   int = 1  // From is greater than To
)

// Version is a machine readable version of the string
// schema identifier that can assist in making indexing easier
type Version struct {
	Major, Minor, Patch int
}

// ReadVersionFromPath allows for parsing paths
// that end in a schema version number string
func ReadVersionFromPath(p string) (*Version, error) { _ = "STUB: not implemented"; return nil, nil }

// path.Base doesn't enforce strict forward slashes
// thus having to test for it here

// GetFamilyAndVersion takes a schemaURL and separates the family from the identifier.
func GetFamilyAndVersion(schemaURL string) (family string, version *Version, err error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

func joinSchemaFamilyAndVersion(family string, version *Version) string {
	_ = "STUB: not implemented"
	return ""
}

// NewVersion converts a near semver like string (ie 1.4.0) into
// a schema identifier that is comparable for a machine.
// The expected string format can be matched by the following regex:
// [0-9]+\.[0-9]+\.[0-9]+
func NewVersion(s string) (*Version, error) { _ = "STUB: not implemented"; return nil, nil }

func (v *Version) String() string { _ = "STUB: not implemented"; return "" }

// Compare returns a digit to represent if the v is equal, less than,
// or greater than o.
// The values are 0, -1 and 1 respectively.
func (v *Version) Compare(o *Version) int { _ = "STUB: not implemented"; return 0 }

func (v *Version) Equal(o *Version) bool { _ = "STUB: not implemented"; return false }

func (v *Version) GreaterThan(o *Version) bool { _ = "STUB: not implemented"; return false }

func (v *Version) LessThan(o *Version) bool { _ = "STUB: not implemented"; return false }

func diff(a, b int) int { _ = "STUB: not implemented"; return 0 }

func gt(a, b int) int { _ = "STUB: not implemented"; return 0 }

func lt(a, b int) int { _ = "STUB: not implemented"; return 0 }

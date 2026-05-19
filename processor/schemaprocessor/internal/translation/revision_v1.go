// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package translation // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/schemaprocessor/internal/translation"

import (
	ast10 "go.opentelemetry.io/otel/schema/v1.0/ast"
	ast11 "go.opentelemetry.io/otel/schema/v1.1/ast"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/schemaprocessor/internal/changelist"
)

// RevisionV1 represents all changes that are to be applied to a signal at a given version.  V1 represents the fact
// that this struct only support the Schema Files version 1.0 - not 1.1 which contains split.
// todo(ankit) implement split and rest of otel schema
type RevisionV1 struct {
	ver        *Version
	all        *changelist.ChangeList
	resources  *changelist.ChangeList
	spans      *changelist.ChangeList
	spanEvents *changelist.ChangeList
	metrics    *changelist.ChangeList
	logs       *changelist.ChangeList
}

// NewRevision processes the VersionDef and assigns the version to this revision
// to allow sorting within a slice.
// Since VersionDef uses custom types for various definitions, it isn't possible
// to cast those values into the primitives so each has to be processed together.
// Generics would be handy here.
func NewRevision(ver *Version, def ast11.VersionDef, copyAttributes bool) (*RevisionV1, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r RevisionV1) Version() *Version { _ = "STUB: not implemented"; return nil }

func newAllChangeList(all ast10.Attributes, copyAttributes bool) *changelist.ChangeList {
	_ = "STUB: not implemented"
	return nil
}

func newResourceChangeList(resource ast10.Attributes, copyAttributes bool) *changelist.ChangeList {
	_ = "STUB: not implemented"
	return nil
}

func newSpanChangeList(spans ast10.Spans, copyAttributes bool) *changelist.ChangeList {
	_ = "STUB: not implemented"
	return nil
}

func newMetricChangeList(metrics ast11.Metrics, copyAttributes bool) *changelist.ChangeList {
	_ = "STUB: not implemented"
	return nil
}

// TODO: Implement split

func newSpanEventChangeList(spanEvents ast10.SpanEvents, copyAttributes bool) (*changelist.ChangeList, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newLogsChangelist(logs ast10.Logs, copyAttributes bool) *changelist.ChangeList {
	_ = "STUB: not implemented"
	return nil
}

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package otelserializer // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/elasticsearchexporter/internal/serializer/otelserializer"

import (
	"bytes"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pprofile"
)

const (
	AllEventsIndex   = "profiling-events-all"
	StackTraceIndex  = "profiling-stacktraces"
	StackFrameIndex  = "profiling-stackframes"
	ExecutablesIndex = "profiling-executables"

	ExecutablesSymQueueIndex = "profiling-sq-executables"
	LeafFramesSymQueueIndex  = "profiling-sq-leafframes"

	HostsMetadataIndex = "profiling-hosts"
)

// SerializeProfile serializes a profile and calls the `pushData` callback for each generated document.
func (s *Serializer) SerializeProfile(dic pprofile.ProfilesDictionary, resource pcommon.Resource, scope pcommon.InstrumentationScope, profile pprofile.Profile, pushData func(*bytes.Buffer, string, string) error) error {
	_ = "STUB: not implemented"
	return nil
}

func toJSON(d any) (*bytes.Buffer, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *Serializer) createLRUs() error { _ = "STUB: not implemented"; return nil }

// Create LRUs with MinILMRolloverTime as lifetime to avoid losing data by ILM roll-over.

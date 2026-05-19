// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package internal // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/lokireceiver/internal"

import (
	"io"

	"github.com/grafana/loki/pkg/push"
)

// PushRequest models a log stream push but is unmarshalled to proto push format.
type PushRequest struct {
	Streams []Stream `json:"streams"`
}

// Stream helps with unmarshalling of each log stream for push request.
type Stream push.Stream

func (s *Stream) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func unmarshalHTTPToLogProtoEntries(data []byte) ([]push.Entry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func unmarshalHTTPToLogProtoEntry(data []byte) (push.Entry, error) {
	_ = "STUB: not implemented"
	return *new(push.Entry), nil
}

// assert that both items in array are of type string

// timestamp

// value

// structuredMetadata

// LabelSet is a key/value pair mapping of labels
type LabelSet map[string]string

func (l *LabelSet) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// String implements the Stringer interface. It returns a formatted/sorted set of label key/value pairs.
func (l LabelSet) String() string { _ = "STUB: not implemented"; return "" }

// decodePushRequest directly decodes json to a push.PushRequest
func decodePushRequest(b io.Reader, r *push.PushRequest) error {
	_ = "STUB: not implemented"
	return nil
}

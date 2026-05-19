// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package elasticsearch // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/elasticsearchexporter/internal/elasticsearch"

type Index struct {
	Index     string
	Type      string
	Dataset   string
	Namespace string
}

func NewDataStreamIndex(typ, dataset, namespace string) Index {
	_ = "STUB: not implemented"
	return *new(Index)
}

func (i Index) IsDataStream() bool { _ = "STUB: not implemented"; return false }

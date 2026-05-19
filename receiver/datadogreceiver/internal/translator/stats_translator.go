// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package translator // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/datadogreceiver/internal/translator"

import (
	"github.com/DataDog/datadog-agent/pkg/obfuscate"
	pb "github.com/DataDog/datadog-agent/pkg/proto/pbgo/trace"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

const (
	maxResourceLen = 5000

	// keyStatsPayload is the key for the stats payload in the attributes map.
	// This is used as Metric name and Attribute key.
	keyStatsPayload = "dd.internal.stats.payload"

	textNonParsable = "Non-parsable SQL query"
)

type StatsTranslator struct {
	obfuscator *obfuscate.Obfuscator
}

func NewStatsTranslator() *StatsTranslator { _ = "STUB: not implemented"; return nil }

func (st *StatsTranslator) TranslateStats(clientStats *pb.ClientStatsPayload, lang, tracerVersion string) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

func (st *StatsTranslator) processStats(in *pb.ClientStatsPayload, lang, tracerVersion string) *pb.ClientStatsPayload {
	_ = "STUB: not implemented"
	return nil
}

func (*StatsTranslator) normalizeStatsGroup(b *pb.ClientGroupedStats, lang string) {
	_ = "STUB: not implemented"
	return
}

func (st *StatsTranslator) obfuscateStatsGroup(b *pb.ClientGroupedStats) {
	_ = "STUB: not implemented"
	return
}

// truncateResource truncates a span's resource to the maximum allowed length.
// It returns true if the input was below the max size.
func truncateResource(r string) (string, bool) { _ = "STUB: not implemented"; return "", false }

func mergeDuplicates(s *pb.ClientStatsBucket) { _ = "STUB: not implemented"; return }

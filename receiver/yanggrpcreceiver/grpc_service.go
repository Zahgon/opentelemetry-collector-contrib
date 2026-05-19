// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package yanggrpcreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/yanggrpcreceiver"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/yanggrpcreceiver/internal"
	pb "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/yanggrpcreceiver/internal/proto/generated/proto"
)

// grpcService handles Cisco gRPC Dial-out telemetry streams.
type grpcService struct {
	pb.UnimplementedGRPCMdtDialoutServer
	receiver   *yangReceiver
	yangParser *internal.YANGParser
}

// MdtDialout processes the bidirectional gRPC stream.
func (s *grpcService) MdtDialout(stream pb.GRPCMdtDialout_MdtDialoutServer) error {
	_ = "STUB: not implemented"
	return nil
}

// processTelemetryData unmarshals the GPBKV payload and triggers OTLP conversion.
func (s *grpcService) processTelemetryData(req *pb.MdtDialoutArgs) error {
	_ = "STUB: not implemented"
	return nil
}

// convertToOTELMetrics maps Cisco KV-GPB data to OTLP using a Telegraf-inspired
// approach: extract identifiers (tags) first, then emit measurements.
func (s *grpcService) convertToOTELMetrics(telemetry *pb.Telemetry) pmetric.Metrics {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics)
}

// Process each entry in DataGpbkv as a distinct row/object.

// Step 1: Initialize context bag with global metadata.

// Step 2: Pre-scan the entire tree for keys/identifiers (Telegraf logic).
// This ensures sibling branches like 'admin-status' can access 'interface-name'.

// Step 3: Walk the tree again to emit actual metrics using the enriched context.

// extractKeys recursively scans for string values that serve as identifiers.
func (s *grpcService) extractKeys(field *pb.TelemetryField, ctxBag map[string]string) {
	_ = "STUB: not implemented"
	return
}

// If it's a string value, it's likely a dimension/tag.

// Common Cisco naming normalization.

// emitMetrics processes numerical values and emits OTLP metrics with the full context bag.
func (s *grpcService) emitMetrics(sm pmetric.ScopeMetrics, field *pb.TelemetryField, pathPrefix string, timestamp pcommon.Timestamp, ctxBag map[string]string) {
	_ = "STUB: not implemented"
	return
}

// Only emit metrics for leaf nodes (values) that are NOT in the 'keys' branch.

// Step/Info metrics for string states (e.g., Up/Down).

// Numeric metrics for counters and gauges.

// createNumericMetric populates a NumberDataPoint.
func createNumericMetric(m pmetric.Metric, name string, val float64, ts pcommon.Timestamp, yType *internal.YANGDataType, ctx map[string]string) {
	_ = "STUB: not implemented"
	return
}

// createStepMetric creates an "Info" metric where the actual value is an attribute.
func createStepMetric(m pmetric.Metric, name, val string, ts pcommon.Timestamp, ctx map[string]string) {
	_ = "STUB: not implemented"
	return
}

// getNumericValue extracts float64 from various protobuf types.
func getNumericValue(f *pb.TelemetryField) float64 { _ = "STUB: not implemented"; return 0 }

// formatValueToString converts protobuf field values to string for labels.
func formatValueToString(f *pb.TelemetryField) string { _ = "STUB: not implemented"; return "" }

func cloneCtxBag(in map[string]string) map[string]string { _ = "STUB: not implemented"; return nil }

func applyCtxBag(attrs pcommon.Map, ctx map[string]string) { _ = "STUB: not implemented"; return }

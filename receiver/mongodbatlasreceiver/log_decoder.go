// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package mongodbatlasreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mongodbatlasreceiver"

import (
	"io"
	"regexp"

	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mongodbatlasreceiver/internal/model"
)

func decodeLogs(logger *zap.Logger, clusterMajorVersion string, r io.Reader) ([]model.LogEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// 4.2 clusters use a console log format

// All other versions use JSON logging

func decodeJSON(logger *zap.Logger, r io.Reader) ([]model.LogEntry, error) {
	_ = "STUB: not implemented"
	// Pass this into a gzip reader for decoding
	return nil, nil
}

// Scan failed; This might just be EOF, in which case Err will be nil, or it could be some other IO error.

var mongo4_2LogRegex = regexp.MustCompile(`^(?P<timestamp>\S+)\s+(?P<severity>\w+)\s+(?P<component>[\w-]+)\s+\[(?P<context>\S+)\]\s+(?P<message>.*)$`)

func decode4_2(logger *zap.Logger, r io.Reader) ([]model.LogEntry, error) {
	_ = "STUB: not implemented"
	// Pass this into a gzip reader for decoding
	return nil, nil
}

// Scan failed; This might just be EOF, in which case Err will be nil, or it could be some other IO error.

// Match failed for line; We will skip this line and continue processing others.

func decodeAuditJSON(logger *zap.Logger, r io.Reader) ([]model.AuditLog, error) {
	_ = "STUB: not implemented"
	// Pass this into a gzip reader for decoding
	return nil, nil
}

// Scan failed; This might just be EOF, in which case Err will be nil, or it could be some other IO error.

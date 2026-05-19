// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build darwin

package macosunifiedloggingreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/macosunifiedloggingreceiver"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.uber.org/zap"
)

// unifiedLoggingReceiver uses exec.Command to run the native macOS `log` command
type unifiedLoggingReceiver struct {
	config   *Config
	logger   *zap.Logger
	consumer consumer.Logs
	cancel   context.CancelFunc
}

func newUnifiedLoggingReceiver(
	config *Config,
	logger *zap.Logger,
	consumer consumer.Logs,
) *unifiedLoggingReceiver {
	_ = "STUB: not implemented"
	return nil
}

func (r *unifiedLoggingReceiver) Start(ctx context.Context, _ component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

// Start reading logs in a goroutine

func (r *unifiedLoggingReceiver) Shutdown(_ context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// readLogs runs the log command and processes output
func (r *unifiedLoggingReceiver) readLogs(ctx context.Context) {
	_ = "STUB: not implemented"
	// Run immediately on startup
	return
}

func (r *unifiedLoggingReceiver) readFromArchive(ctx context.Context) {
	_ = "STUB: not implemented"
	return
}

func (r *unifiedLoggingReceiver) readFromLive(ctx context.Context) {
	_ = "STUB: not implemented"
	// Run immediately on startup
	return
}

// For live mode, use exponential backoff based on whether logs are being actively written
// We cannot safely reset the backoff while ticker is running (causes data race)
// Instead, track interval manually and use time.After which creates a new timer each iteration

// Start immediately

// Run immediately on start, then use backoff for subsequent iterations

// Run the log command

// Adjust interval based on whether logs were found

// If no logs, exponentially increase interval up to MaxPollInterval

// First iteration with no logs, start with min interval

// runLogCommand executes the log command and processes output
// Returns the number of logs processed
// archivePath should be empty string for live mode, or a specific archive path for archive mode
func (r *unifiedLoggingReceiver) runLogCommand(ctx context.Context, archivePath string) (int, error) {
	_ = "STUB: not implemented"
	// Build the log command arguments
	return 0, nil
}

// Create the command
// #nosec G204 - args are controlled by config

// Get stdout pipe

// Start the command

// Ensure the process is properly cleaned up to avoid zombies

// Read and process output line by line

// Set a large buffer size for long log lines
// 1MB buffer
// 10MB max

// Skip the header line in text-based formats (default, syslog, compact)

// Skip completion/status messages (applies to all formats)

// Parse and send the log entry

// buildLogCommandArgs constructs the arguments for the log command
// archivePath should be empty string for live mode, or a specific archive path for archive mode
func (r *unifiedLoggingReceiver) buildLogCommandArgs(archivePath string) []string {
	_ = "STUB: not implemented"
	return nil

	// Add archive path if specified
}

// Add style flag if format is not default

// Add start time

// For live mode, calculate start time from max_log_age

// Add end time (archive mode only)

// Add predicate filter

// processLogLine processes a log line and sends it to the consumer
func (r *unifiedLoggingReceiver) processLogLine(ctx context.Context, line []byte) error {
	_ = "STUB: not implemented"
	// Convert to OTel plog
	return nil
}

// Put the entire line into the log body as a string

// Parse timestamp and severity when using JSON formats

// Parse and set timestamp

// Set severity from messageType

// Send to consumer

// mapMessageTypeToSeverity maps log messageType to OTel severity
func mapMessageTypeToSeverity(msgType string) plog.SeverityNumber {
	_ = "STUB: not implemented"
	return *new(plog.SeverityNumber)
}

// isCompletionLine checks if a line is a completion/status message from the log command
// These lines should be filtered out (e.g., {"count":540659,"finished":1})
func isCompletionLine(line []byte) bool {
	_ = "STUB: not implemented"
	// Trim whitespace
	return false
}

// Check if line is empty

// Check if line starts with "**" (typical completion message format)

// Check for JSON completion format: {"count":N,"finished":1}

// Quick check for both "count" and "finished" fields

// Check for common completion keywords

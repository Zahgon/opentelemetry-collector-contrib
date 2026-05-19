// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package awscloudwatchreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awscloudwatchreceiver"

import (
	"context"

	"go.opentelemetry.io/collector/extension/xextension/storage"
	"go.uber.org/zap"
)

const checkpointKeyFormat = "cloudwatch/%s"

type cloudwatchCheckpointPersister struct {
	client storage.Client
	logger *zap.Logger
}

func newCloudwatchCheckpointPersister(client storage.Client, logger *zap.Logger) *cloudwatchCheckpointPersister {
	_ = "STUB: not implemented"
	return nil
}

// SetCheckpoint stores the checkpoint (timestamp) for a specific log stream
func (p *cloudwatchCheckpointPersister) SetCheckpoint(
	ctx context.Context,
	logGroupName, timestamp string,
) error {
	_ = "STUB: not implemented"
	return nil
}

// GetCheckpoint retrieves the checkpoint (timestamp) for a specific log stream
func (p *cloudwatchCheckpointPersister) GetCheckpoint(
	ctx context.Context, logGroupName string,
) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// If key is not found, data and error is nil

// DeleteCheckpoint removes the checkpoint (timestamp) for a specific log stream
func (p *cloudwatchCheckpointPersister) DeleteCheckpoint(
	ctx context.Context, logGroupName string,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *cloudwatchCheckpointPersister) Shutdown(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// getCheckpointKey generates a unique storage key
func (*cloudwatchCheckpointPersister) getCheckpointKey(logGroupName string) string {
	_ = "STUB: not implemented"
	return ""
}

// newCheckpointTimeFromStartOfStream returns the Unix epoch start time as a string in RFC3339 format,
// which is the default timestamp for starting at the beginning.
func newCheckpointTimeFromStartOfStream() string { _ = "STUB: not implemented"; return "" }

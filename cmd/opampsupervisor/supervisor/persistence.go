// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package supervisor

import (
	"github.com/google/uuid"
	"github.com/open-telemetry/opamp-go/protobufs"
	"go.uber.org/zap"
)

// persistentState represents persistent state for the supervisor
type persistentState struct {
	// InstanceID must be a valid UUID string. If it is not, a new UUIDv7 will be generated automatically.
	InstanceID             uuid.UUID           `yaml:"instance_id"`
	LastRemoteConfigStatus *RemoteConfigStatus `yaml:"last_remote_config_status"`

	// Path to the config file that the state should be saved to.
	// This is not marshaled.
	configPath string      `yaml:"-"`
	logger     *zap.Logger `yaml:"-"`
}

// RemoteConfigStatus is a custom struct that is used to marshal/unmarshal the remote config status.
// LastRemoteConfigHash is a hex encoded string of the last remote config hash for human readability.
type RemoteConfigStatus struct {
	// Status is the status of the last remote config.
	Status protobufs.RemoteConfigStatuses `yaml:"status"`
	// LastRemoteConfigHash is a hex encoded string of the last remote config hash for human readability.
	LastRemoteConfigHash string `yaml:"last_remote_config_hash"`
	// ErrorMessage is the error message of the last remote config.
	ErrorMessage string `yaml:"error_message"`
}

func (p *persistentState) SetInstanceID(id uuid.UUID) error { _ = "STUB: not implemented"; return nil }

func (p *persistentState) SetLastRemoteConfigStatus(status *protobufs.RemoteConfigStatus) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *persistentState) GetLastRemoteConfigStatus() *protobufs.RemoteConfigStatus {
	_ = "STUB: not implemented"
	return nil
}

func (p *persistentState) writeState() error { _ = "STUB: not implemented"; return nil }

// loadOrCreatePersistentState attempts to load the persistent state from disk. If it doesn't
// exist, a new persistent state file is created.
// instanceID must be a valid UUID string, or an empty string to generate a new UUIDv7 automatically.
func loadOrCreatePersistentState(file, instanceID string, logger *zap.Logger) (*persistentState, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func loadPersistentState(file string, logger *zap.Logger) (*persistentState, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createNewPersistentState(file, instanceID string, logger *zap.Logger) (*persistentState, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

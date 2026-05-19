// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package azuredataexplorerexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/azuredataexplorerexporter"

import (
	"go.opentelemetry.io/collector/config/configopaque"
	"go.opentelemetry.io/collector/config/configoptional"
	"go.opentelemetry.io/collector/config/configretry"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
)

// Config defines configuration for Azure Data Explorer Exporter
type Config struct {
	TimeoutSettings           exporterhelper.TimeoutConfig                             `mapstructure:",squash"` // squash ensures fields are correctly decoded in embedded struct.
	QueueSettings             configoptional.Optional[exporterhelper.QueueBatchConfig] `mapstructure:"sending_queue"`
	configretry.BackOffConfig `mapstructure:"retry_on_failure"`
	ClusterURI                string              `mapstructure:"cluster_uri"`
	ApplicationID             string              `mapstructure:"application_id"`
	ApplicationKey            configopaque.String `mapstructure:"application_key"`
	TenantID                  string              `mapstructure:"tenant_id"`
	ManagedIdentityID         string              `mapstructure:"managed_identity_id"`
	UseAzureAuth              bool                `mapstructure:"use_azure_auth"`
	Database                  string              `mapstructure:"db_name"`
	MetricTable               string              `mapstructure:"metrics_table_name"`
	LogTable                  string              `mapstructure:"logs_table_name"`
	TraceTable                string              `mapstructure:"traces_table_name"`
	MetricTableMapping        string              `mapstructure:"metrics_table_json_mapping"`
	LogTableMapping           string              `mapstructure:"logs_table_json_mapping"`
	TraceTableMapping         string              `mapstructure:"traces_table_json_mapping"`
	IngestionType             string              `mapstructure:"ingestion_type"`
}

// Validate checks if the exporter configuration is valid
func (adxCfg *Config) Validate() error { _ = "STUB: not implemented"; return nil }

// Cluster URI is the target ADX cluster

// Parameters for AD App Auth or Managed Identity Auth or Default Auth are mandatory

// Validate managed identity ID. Use system for system assigned managed identity or UserManagedIdentityID (objectID) for user assigned managed identity

// if the managed identity is not a system identity, validate if it is a valid UUID

func isEmpty(str string) bool { _ = "STUB: not implemented"; return false }

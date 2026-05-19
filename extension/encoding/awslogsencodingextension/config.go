// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package awslogsencodingextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding/awslogsencodingextension"

import (
	"go.opentelemetry.io/collector/confmap/xconfmap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding/awslogsencodingextension/internal/constants"
	subscriptionfilter "github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding/awslogsencodingextension/internal/unmarshaler/subscription-filter"
	vpcflowlog "github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding/awslogsencodingextension/internal/unmarshaler/vpc-flow-log"
)

var _ xconfmap.Validator = (*Config)(nil)

var (
	supportedLogFormats = []string{
		constants.FormatCloudWatchLogsSubscriptionFilter,
		constants.FormatVPCFlowLog,
		constants.FormatS3AccessLog,
		constants.FormatWAFLog,
		constants.FormatCloudTrailLog,
		constants.FormatELBAccessLog,
		constants.FormatNetworkFirewallLog,
	}
	supportedVPCFlowLogFileFormat = []string{constants.FileFormatPlainText, constants.FileFormatParquet}
)

type Config struct {
	// Format selects the AWS logs format. See supportedLogFormats for valid values.
	Format string `mapstructure:"format"`

	VPCFlowLogConfig vpcflowlog.Config `mapstructure:"vpcflow"`

	// CloudWatch is consulted only when Format is a CloudWatch subscription-filter format.
	CloudWatch CloudWatchConfig `mapstructure:"cloudwatch"`

	// prevent unkeyed literal initialization
	_ struct{}
}

type CloudWatchConfig struct {
	// Streams routes subscription-filter events to inner encoding extensions
	// by logGroup/logStream pattern or service name. Empty means no routing.
	Streams []subscriptionfilter.CloudWatchStream `mapstructure:"streams"`

	// prevent unkeyed literal initialization
	_ struct{}
}

func (cfg *Config) Validate() error { _ = "STUB: not implemented"; return nil }

// valid
// valid
// valid
// valid
// valid
// valid
// valid

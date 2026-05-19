// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package vpcflowlog // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding/awslogsencodingextension/internal/unmarshaler/vpc-flow-log"

import (
	"bufio"
	"io"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding/awslogsencodingextension/internal/constants"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding/awslogsencodingextension/internal/unmarshaler"
)

var (
	supportedVPCFlowLogFileFormat = []string{
		constants.FileFormatPlainText,
		constants.FileFormatParquet,
	}

	// defaultVPCFormat holds the default format for VPC flow logs, when
	// sent to CloudWatch Logs. For logs sent to S3, the format is
	// determined from the header line.
	//
	// Note: order is important.
	defaultVPCFormat = []string{
		"version",
		"account-id",
		"interface-id",
		"srcaddr",
		"dstaddr",
		"srcport",
		"dstport",
		"protocol",
		"packets",
		"bytes",
		"start",
		"end",
		"action",
		"log-status",
	}

	// defaultTGWFormat holds the default format for Transit Gateway flow logs,
	// when sent to CloudWatch Logs. For logs sent to S3, the format is determined
	// from the header line.
	//
	// Note: order is important.
	defaultTGWFormat = []string{
		"version",
		"resource-type",
		"account-id",
		"tgw-id",
		"tgw-attachment-id",
		"tgw-src-vpc-id",
		"tgw-dst-vpc-id",
		"tgw-src-subnet-id",
		"tgw-dst-subnet-id",
		"tgw-src-eni",
		"tgw-dst-eni",
		"tgw-src-az-id",
		"tgw-dst-az-id",
		"srcaddr",
		"dstaddr",
		"srcport",
		"dstport",
		"protocol",
		"packets",
		"bytes",
		"start",
		"end",
		"log-status",
	}
)

var _ unmarshaler.StreamingLogsUnmarshaler = (*VPCFlowLogUnmarshaler)(nil)

type Config struct {
	// fileFormat - VPC flow logs can be sent in plain text or parquet files to S3.
	// See https://docs.aws.amazon.com/vpc/latest/userguide/flow-logs-s3-path.html.
	FileFormat string `mapstructure:"file_format"`

	// Format defines the custom field format of the VPC flow logs.
	// When defined, this is used for parsing CloudWatch bound VPC flow logs.
	Format string `mapstructure:"format"`

	// parsedFormat holds the parsed fields from Format.
	parsedFormat []string

	// prevent unkeyed literal initialization
	_ struct{}
}

type VPCFlowLogUnmarshaler struct {
	cfg Config

	buildInfo component.BuildInfo
	logger    *zap.Logger

	// Whether VPC flow start field should use ISO-8601 format
	vpcFlowStartISO8601FormatEnabled bool
}

func NewVPCFlowLogUnmarshaler(
	cfg Config,
	buildInfo component.BuildInfo,
	logger *zap.Logger,
	vpcFlowStartISO8601FormatEnabled bool,
) (*VPCFlowLogUnmarshaler, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO

// valid

func (v *VPCFlowLogUnmarshaler) UnmarshalAWSLogs(reader io.Reader) (plog.Logs, error) {
	_ = "STUB: not implemented"
	return *new(plog.Logs), nil
}

// Decode as a stream but flush all at once using flush options

//nolint:errorlint

// EOF indicates no logs were found, return any logs that's available

// TODO

// not possible, prevent by NewVPCFlowLogUnmarshaler

// NewLogsDecoder returns a LogsDecoder that processes VPC flow logs from the provided reader.
// Auto-detects the source format (S3 plain text or CloudWatch subscription filter) from the first byte.
// Supported sub formats:
//   - S3 plain text logs: Supports offset-based streaming; offset tracked by bytes processed
//   - CloudWatch subscription filter: Processes full payload; offset tracked by bytes processed
//   - Parquet format: Not yet implemented
func (v *VPCFlowLogUnmarshaler) NewLogsDecoder(reader io.Reader, options ...encoding.DecoderOption) (encoding.LogsDecoder, error) {
	_ = "STUB: not implemented"
	return *new(encoding.LogsDecoder), nil
}

// use buffered reader for efficiency and to avoid any size restrictions

// Dealing with a JSON log message, so check for CloudWatch bound trigger

// If offset is set, return EOF after consuming the whole record.
// This confirms to our contract - process full payload
// However, we cannot skip the offset bytes as we need full record unmarshaling.

// discard bytes, ignoring the first line

// Trim spaces and new lines

// fromCloudWatch expects VPC logs from CloudWatch Logs subscription filter trigger
func (v *VPCFlowLogUnmarshaler) fromCloudWatch(fields []string, reader *bufio.Reader) (plog.Logs, int64, error) {
	_ = "STUB: not implemented"
	return *new(plog.Logs), 0, nil
}

// No format specified, so we assume the default format. The default format is different
// for Transit Gateway and plain VPC flow logs, so we need to inspect the log message.

// The 2nd field is "TransitGateway" for TGW logs (resource-type field)

// createLogs is a helper to create prefilled plog.Logs, plog.ResourceLogs, plog.ScopeLogs
func (v *VPCFlowLogUnmarshaler) createLogs() (plog.Logs, plog.ResourceLogs, plog.ScopeLogs) {
	_ = "STUB: not implemented"
	return *new(plog.Logs), *new(plog.ResourceLogs), *new(plog.ScopeLogs)
}

// addToLogs parses the log line and creates
// a new record log. The record log is added
// to the scope logs of the resource identified
// by the resourceKey created from the values.
func (v *VPCFlowLogUnmarshaler) addToLogs(
	resourceLogs plog.ResourceLogs,
	scopeLogs plog.ScopeLogs,
	fields []string,
	logLine string,
) error {
	_ = "STUB: not implemented"
	return nil
}

// If a field is not applicable or could not be computed for a
// specific record, the record displays a '-' symbol for that entry.
//
// See https://docs.aws.amazon.com/vpc/latest/userguide/flow-log-records.html.

// Add the address fields with the correct conventions to the log record

// handleAddresses creates adds the addresses to the log record
func (v *VPCFlowLogUnmarshaler) handleAddresses(addr *address, record plog.LogRecord) {
	_ = "STUB: not implemented"
	return

	// see example in
	// https://docs.aws.amazon.com/vpc/latest/userguide/flow-logs-records-examples.html#flow-log-example-nat
}

// there is no middle layer, assume "srcaddr" field
// corresponds to the original source address.

// srcaddr is the middle layer

// there is no middle layer, assume "dstaddr" field
// corresponds to the original destination address.

// dstaddr is the middle layer

// handleField analyzes the given field and it either
// adds its value to the resourceKey or puts the
// field and its value in the attributes map. If the
// field is not recognized, it returns false.
func (v *VPCFlowLogUnmarshaler) handleField(
	field string,
	value string,
	resourceLogs plog.ResourceLogs,
	record plog.LogRecord,
	addr *address,
) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Integer fields: parse and call handleInt64Field

// handleStringField applies a string value for the given field to resource/record/addr.
// If the field is not recognized, it returns false.
func (*VPCFlowLogUnmarshaler) handleStringField(
	field string,
	value string,
	resourceLogs plog.ResourceLogs,
	record plog.LogRecord,
	addr *address,
) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// handled later

// handled later

// handled later

// handled later

// Skip - used for detection only, already captured in encoding.format

// handleInt64Field applies an int64 value for the given field to resource/record.
// Returns (true, nil) when the field is recognized, (false, nil) when unknown, or (_, err) on error.
func (v *VPCFlowLogUnmarshaler) handleInt64Field(
	field string,
	value int64,
	_ plog.ResourceLogs,
	record plog.LogRecord,
	_ *address,
) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// New behavior: ISO-8601 format (RFC3339Nano)

// Legacy behavior: Unix timestamp as integer

// address stores the four fields related to the address
// of a VPC flow log: srcaddr, pkt-srcaddr, dstaddr, and
// pkt-dstaddr. We save these fields in a struct, so we
// can use the right naming conventions in the end.
//
// See https://docs.aws.amazon.com/vpc/latest/userguide/flow-logs-records-examples.html#flow-log-example-nat.
type address struct {
	source         string
	pktSource      string
	destination    string
	pktDestination string
}

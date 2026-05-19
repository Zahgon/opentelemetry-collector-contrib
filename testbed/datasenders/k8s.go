// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package datasenders // import "github.com/open-telemetry/opentelemetry-collector-contrib/testbed/datasenders"

import (
	"context"
	"net"
	"os"

	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/plog"

	"github.com/open-telemetry/opentelemetry-collector-contrib/testbed/testbed"
)

// FileLogK8sWriter represents abstract container k8s writer
type FileLogK8sWriter struct {
	file   *os.File
	config string
}

// Ensure FileLogK8sWriter implements LogDataSender.
var _ testbed.LogDataSender = (*FileLogK8sWriter)(nil)

// NewFileLogK8sWriter creates a new data sender that will write kubernetes containerd
// log entries to a file, to be tailed by FileLogReceiver and sent to the collector.
//
// config is an Otelcol config appended to the receivers section after executing fmt.Sprintf on it.
// This implies few things:
//   - it should contain `%s` which will be replaced with the filename
//   - all `%` should be represented as `%%`
//   - indentation style matters. Spaces have to be used for indentation
//     and it should start with two spaces indentation
//
// Example config:
// |`
// |  file_log:
// |    include: [ %s ]
// |    start_at: beginning
// |    operators:
// |      type: regex_parser
// |      regex: ^(?P<log>.*)$
// |  `
func NewFileLogK8sWriter(config string) *FileLogK8sWriter { _ = "STUB: not implemented"; return nil }

func (*FileLogK8sWriter) Capabilities() consumer.Capabilities {
	_ = "STUB: not implemented"
	return *new(consumer.Capabilities)
}

func (*FileLogK8sWriter) Start() error { _ = "STUB: not implemented"; return nil }

func (f *FileLogK8sWriter) ConsumeLogs(_ context.Context, logs plog.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

func (*FileLogK8sWriter) convertLogToTextLine(lr plog.LogRecord) []byte {
	_ = "STUB: not implemented"
	return nil

	// Timestamp
}

// Severity

func (f *FileLogK8sWriter) Flush() { _ = "STUB: not implemented"; return }

func (f *FileLogK8sWriter) GenConfigYAMLStr() string {
	_ = "STUB: not implemented"
	// Note that this generates a receiver config for agent.
	// We are testing file_log receiver here.
	return ""
}

func (*FileLogK8sWriter) ProtocolName() string { _ = "STUB: not implemented"; return "" }

func (*FileLogK8sWriter) GetEndpoint() net.Addr {
	_ = "STUB: not implemented"

	// NewKubernetesContainerWriter returns FileLogK8sWriter with configuration
	// to recognize and parse kubernetes container logs
	return *new(net.Addr)
}

func NewKubernetesContainerWriter() *FileLogK8sWriter { _ = "STUB: not implemented"; return nil }

// NewKubernetesContainerParserWriter returns FileLogK8sWriter with configuration
// to recognize and parse kubernetes container logs using the container parser
func NewKubernetesContainerParserWriter() *FileLogK8sWriter { _ = "STUB: not implemented"; return nil }

// NewKubernetesCRIContainerdWriter returns FileLogK8sWriter with configuration
// to parse only CRI-Containerd kubernetes logs
func NewKubernetesCRIContainerdWriter() *FileLogK8sWriter { _ = "STUB: not implemented"; return nil }

// NewKubernetesCRIContainerdNoAttributesOpsWriter returns FileLogK8sWriter with configuration
// to parse only CRI-Containerd kubernetes logs without reformatting attributes
func NewKubernetesCRIContainerdNoAttributesOpsWriter() *FileLogK8sWriter {
	_ = "STUB: not implemented"
	return nil
}

// NewCRIContainerdWriter returns FileLogK8sWriter with configuration
// to parse only CRI-Containerd logs (no extracting metadata from filename)
func NewCRIContainerdWriter() *FileLogK8sWriter { _ = "STUB: not implemented"; return nil }

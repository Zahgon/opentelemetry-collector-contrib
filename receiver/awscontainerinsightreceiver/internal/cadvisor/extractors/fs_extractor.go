// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package extractors // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awscontainerinsightreceiver/internal/cadvisor/extractors"

import (
	"regexp"

	cinfo "github.com/google/cadvisor/info/v1"
	"go.uber.org/zap"
)

var allowedPaths = regexp.MustCompile(`^(tmpfs|/dev/.*|overlay)$`)

type FileSystemMetricExtractor struct {
	allowListRegexP *regexp.Regexp
	logger          *zap.Logger
}

func (*FileSystemMetricExtractor) HasValue(info *cinfo.ContainerInfo) bool {
	_ = "STUB: not implemented"
	return false
}

func (f *FileSystemMetricExtractor) GetValue(info *cinfo.ContainerInfo, _ CPUMemInfoProvider, containerType string) []*CAdvisorMetric {
	_ = "STUB: not implemented"
	return nil
}

func (*FileSystemMetricExtractor) Shutdown() error { _ = "STUB: not implemented"; return nil }

func NewFileSystemMetricExtractor(logger *zap.Logger) *FileSystemMetricExtractor {
	_ = "STUB: not implemented"
	return nil
}

func getFSMetricType(containerType string, logger *zap.Logger) string {
	_ = "STUB: not implemented"
	return ""
}

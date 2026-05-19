// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package testutils // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awscontainerinsightreceiver/internal/cadvisor/testutils"

import (
	"testing"

	cinfo "github.com/google/cadvisor/info/v1"
)

func LoadContainerInfo(t *testing.T, file string) []*cinfo.ContainerInfo {
	_ = "STUB: not implemented"
	return nil
}

type MockCPUMemInfo struct{}

func (MockCPUMemInfo) GetNumCores() int64 { _ = "STUB: not implemented"; return 0 }

func (MockCPUMemInfo) GetMemoryCapacity() int64 { _ = "STUB: not implemented"; return 0 }

type MockHostInfo struct {
	MockCPUMemInfo
	ClusterName string
	InstanceIP  string
}

func (m MockHostInfo) GetClusterName() string { _ = "STUB: not implemented"; return "" }

func (MockHostInfo) GetEBSVolumeID(string) string { _ = "STUB: not implemented"; return "" }

func (MockHostInfo) GetInstanceID() string { _ = "STUB: not implemented"; return "" }

func (MockHostInfo) GetInstanceType() string { _ = "STUB: not implemented"; return "" }

func (MockHostInfo) GetAutoScalingGroupName() string { _ = "STUB: not implemented"; return "" }

func (MockHostInfo) ExtractEbsIDsUsedByKubernetes() map[string]string {
	_ = "STUB: not implemented"
	return nil
}

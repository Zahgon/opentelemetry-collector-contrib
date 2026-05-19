// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package testutils // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awscontainerinsightreceiver/internal/cadvisor/testutils"

type MockECSInfo struct {
	ClusterName string
	InstanceIP  string
}

func (*MockECSInfo) GetRunningTaskCount() int64 { _ = "STUB: not implemented"; return 0 }

func (*MockECSInfo) GetCPUReserved() int64 { _ = "STUB: not implemented"; return 0 }

func (*MockECSInfo) GetMemReserved() int64 { _ = "STUB: not implemented"; return 0 }

func (*MockECSInfo) GetContainerInstanceID() string { _ = "STUB: not implemented"; return "" }

func (*MockECSInfo) GetClusterName() string { _ = "STUB: not implemented"; return "" }

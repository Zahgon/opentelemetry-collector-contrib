// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package supervisor

import (
	"context"
	"io"

	"github.com/open-telemetry/opamp-go/client/types"
	"github.com/open-telemetry/opamp-go/protobufs"
	"go.uber.org/zap"
)

// packageManager manages the persistent state of downloadable packages.
// Currently only allows for a single top-level package containing the agent.
type packageManager struct {
	logger *zap.Logger
}

var _ types.PackagesStateProvider = &packageManager{}

func (p *packageManager) AllPackagesHash() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *packageManager) SetAllPackagesHash(_ []byte) error { _ = "STUB: not implemented"; return nil }

func (p *packageManager) Packages() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func (p *packageManager) PackageState(packageName string) (types.PackageState, error) {
	_ = "STUB: not implemented"
	return *new(types.PackageState), nil
}

func (p *packageManager) SetPackageState(packageName string, _ types.PackageState) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *packageManager) CreatePackage(packageName string, _ protobufs.PackageType) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *packageManager) FileContentHash(packageName string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *packageManager) UpdateContent(_ context.Context, packageName string, _ io.Reader, _, _ []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *packageManager) DeletePackage(packageName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *packageManager) LastReportedStatuses() (*protobufs.PackageStatuses, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *packageManager) SetLastReportedStatuses(_ *protobufs.PackageStatuses) error {
	_ = "STUB: not implemented"
	return nil
}

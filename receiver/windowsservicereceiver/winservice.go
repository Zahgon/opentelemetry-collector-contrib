// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build windows

package windowsservicereceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/windowsservicereceiver"

import (
	"golang.org/x/sys/windows/svc"
	"golang.org/x/sys/windows/svc/mgr"
)

type winService struct {
	mgr    *serviceManager
	name   string
	handle *mgr.Service

	status svc.Status
	config configEx
}

func updateService(m *serviceManager, sname string) (*winService, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ws *winService) updateStatus() error { _ = "STUB: not implemented"; return nil }

func (ws *winService) updateConfig() error { _ = "STUB: not implemented"; return nil }

func (ws *winService) close() error { _ = "STUB: not implemented"; return nil }

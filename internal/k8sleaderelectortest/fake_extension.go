// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package k8sleaderelectortest // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/k8sleaderelectortest"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pipeline"

	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/k8sleaderelector"
)

type FakeHost struct {
	FakeLeaderElection *FakeLeaderElection
}

func (fh *FakeHost) GetExtensions() map[component.ID]component.Component {
	_ = "STUB: not implemented"
	return nil
}

func (*FakeHost) GetExporters() map[pipeline.Signal]map[component.ID]component.Component {
	_ = "STUB: not implemented"
	return nil
}

type FakeLeaderElection struct {
	OnLeading  func(context.Context)
	OnStopping func()
}

func (fle *FakeLeaderElection) SetCallBackFuncs(onLeading k8sleaderelector.StartCallback, onStopping k8sleaderelector.StopCallback) {
	_ = "STUB: not implemented"
	return
}

func (fle *FakeLeaderElection) InvokeOnLeading() { _ = "STUB: not implemented"; return }

func (*FakeLeaderElection) Start(context.Context, component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (*FakeLeaderElection) Shutdown(context.Context) error { _ = "STUB: not implemented"; return nil }

func (fle *FakeLeaderElection) InvokeOnStopping() { _ = "STUB: not implemented"; return }

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build linux

package systemdreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/systemdreceiver"

import (
	"context"
	"errors"

	"github.com/containerd/cgroups/v3/cgroup2"
	"github.com/godbus/dbus/v5"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/systemdreceiver/internal/metadata"
)

var errNotUnifiedCGroup = errors.New("not using unified cgroups")

// A connection to dbus.
//
// This has the same interface as [dbus.Conn], but allows us to mock it out for testing.
type dbusConnection interface {
	Object(dest string, path dbus.ObjectPath) dbus.BusObject
	Close() error
}

type systemdScraper struct {
	conn       dbusConnection
	cfg        *Config
	settings   component.TelemetrySettings
	mb         *metadata.MetricsBuilder
	cgroupOpts []cgroup2.InitOpts
}

func (s *systemdScraper) start(ctx context.Context, _ component.Host) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (s *systemdScraper) shutdown(_ context.Context) (err error) {
	_ = "STUB: not implemented"
	return nil
}

type unitTuple struct {
	Name        string
	Description string
	LoadState   string
	ActiveState string
	SubState    string
	Following   string
	Path        dbus.ObjectPath
	JobID       uint32
	JobType     string
	JobPath     dbus.ObjectPath
}

// Find the active cgroup for each service and scrape statistics from it.
func (s *systemdScraper) scrapeServiceCgroup(now pcommon.Timestamp, unit *unitTuple) error {
	_ = "STUB: not implemented"
	return nil
}

// See https://www.kernel.org/doc/html/v4.18/admin-guide/cgroup-v2.html for more information on the various
// controllers.

// Are any of our cgroup requiring metrics available
func (s *systemdScraper) hasCgroupMetrics() bool { _ = "STUB: not implemented"; return false }

func (s *systemdScraper) scrapeRestartCount(now pcommon.Timestamp, unit *unitTuple) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *systemdScraper) scrape(ctx context.Context) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

func newScraper(conf *Config, settings receiver.Settings) *systemdScraper {
	_ = "STUB: not implemented"
	return nil
}

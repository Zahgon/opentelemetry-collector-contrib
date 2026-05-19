// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package redfishreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/redfishreceiver"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/redfishreceiver/internal/redfish"
)

func (s *redfishScraper) recordComputerSystem(compSys *redfish.ComputerSystem) {
	_ = "STUB: not implemented"
	return
}

func (s *redfishScraper) recordChassis(chassis *redfish.Chassis) { _ = "STUB: not implemented"; return }

func (s *redfishScraper) recordFans(chassisID string, fans []redfish.Fan) {
	_ = "STUB: not implemented"
	return
}

func (s *redfishScraper) recordTemperatures(chassisID string, temps []redfish.Temperature) {
	_ = "STUB: not implemented"
	return
}

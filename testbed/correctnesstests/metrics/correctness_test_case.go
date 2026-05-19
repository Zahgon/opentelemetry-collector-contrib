// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package metrics // import "github.com/open-telemetry/opentelemetry-collector-contrib/testbed/correctnesstests/metrics"

import (
	"testing"

	"go.opentelemetry.io/collector/otelcol"

	"github.com/open-telemetry/opentelemetry-collector-contrib/testbed/testbed"
)

type correctnessTestCase struct {
	t         *testing.T
	sender    testbed.DataSender
	receiver  testbed.DataReceiver
	harness   *testHarness
	collector testbed.OtelcolRunner
}

func newCorrectnessTestCase(
	t *testing.T,
	sender testbed.DataSender,
	receiver testbed.DataReceiver,
	harness *testHarness,
) *correctnessTestCase {
	_ = "STUB: not implemented"
	return nil
}

func (tc *correctnessTestCase) startCollector() { _ = "STUB: not implemented"; return }

func (tc *correctnessTestCase) stopCollector() { _ = "STUB: not implemented"; return }

func (tc *correctnessTestCase) startTestbedSender() { _ = "STUB: not implemented"; return }

func (tc *correctnessTestCase) startTestbedReceiver() { _ = "STUB: not implemented"; return }

func (tc *correctnessTestCase) stopTestbedReceiver() { _ = "STUB: not implemented"; return }

func (tc *correctnessTestCase) sendFirstMetric() { _ = "STUB: not implemented"; return }

func (tc *correctnessTestCase) waitForAllMetrics() { _ = "STUB: not implemented"; return }

func componentFactories(t *testing.T) otelcol.Factories {
	_ = "STUB: not implemented"
	return *new(otelcol.Factories)
}

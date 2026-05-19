// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package scraperinttest // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/coreinternal/scraperinttest"

import (
	"sync"
	"testing"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/receiver"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/pdatatest/pmetrictest"
)

func NewIntegrationTest(f receiver.Factory, opts ...TestOption) *IntegrationTest {
	_ = "STUB: not implemented"
	return nil
}

type IntegrationTest struct {
	containerRequests      []testcontainers.ContainerRequest
	allowHardcodedHostPort bool
	createContainerTimeout time.Duration

	factory      receiver.Factory
	customConfig customConfigFunc

	expectedFile   string
	compareOptions []pmetrictest.CompareMetricsOption
	compareTimeout time.Duration

	failOnErrorLogs bool
	writeExpected   bool
}

func (it *IntegrationTest) Run(t *testing.T) { _ = "STUB: not implemented"; return }

// Defined outside of Eventually so it can be printed if the test fails

func (it *IntegrationTest) createContainers(t *testing.T) *ContainerInfo {
	_ = "STUB: not implemented"
	return nil
}

func (it *IntegrationTest) validate(t *testing.T) { _ = "STUB: not implemented"; return }

type TestOption func(*IntegrationTest)

func WithContainerRequest(cr testcontainers.ContainerRequest) TestOption {
	_ = "STUB: not implemented"
	return *new(TestOption)
}

func AllowHardcodedHostPort() TestOption { _ = "STUB: not implemented"; return *new(TestOption) }

func WithCreateContainerTimeout(t time.Duration) TestOption {
	_ = "STUB: not implemented"
	return *new(TestOption)
}

func WithCustomConfig(c customConfigFunc) TestOption {
	_ = "STUB: not implemented"
	return *new(TestOption)
}

func WithExpectedFile(f string) TestOption { _ = "STUB: not implemented"; return *new(TestOption) }

// This option is useful for debugging scrapers but should not be used permanently
// because the logs do not correlate to a single scrape interval. In other words,
// when a retryable failure occurs, this setting will likely force a failure anyways.
func FailOnErrorLogs() TestOption { _ = "STUB: not implemented"; return *new(TestOption) }

func WriteExpected() TestOption { _ = "STUB: not implemented"; return *new(TestOption) }

func WithCompareOptions(opts ...pmetrictest.CompareMetricsOption) TestOption {
	_ = "STUB: not implemented"
	return *new(TestOption)
}

func WithCompareTimeout(t time.Duration) TestOption {
	_ = "STUB: not implemented"
	return *new(TestOption)
}

type customConfigFunc func(*testing.T, component.Config, *ContainerInfo)

type ContainerInfo struct {
	sync.Mutex
	containers map[string]testcontainers.Container
}

func (ci *ContainerInfo) Host(t *testing.T) string { _ = "STUB: not implemented"; return "" }

func (ci *ContainerInfo) HostForNamedContainer(t *testing.T, containerName string) string {
	_ = "STUB: not implemented"
	return ""
}

func (ci *ContainerInfo) MappedPort(t *testing.T, port string) string {
	_ = "STUB: not implemented"
	return ""
}

func (ci *ContainerInfo) MappedPortForNamedContainer(t *testing.T, containerName, port string) string {
	_ = "STUB: not implemented"
	return ""
}

func (ci *ContainerInfo) container(t *testing.T, name string) testcontainers.Container {
	_ = "STUB: not implemented"
	return *new(testcontainers.Container)
}

func (ci *ContainerInfo) add(name string, c testcontainers.Container) {
	_ = "STUB: not implemented"
	return
}

func (ci *ContainerInfo) terminate(t *testing.T) { _ = "STUB: not implemented"; return }

func RunScript(script []string) testcontainers.ContainerHook {
	_ = "STUB: not implemented"
	return *new(testcontainers.ContainerHook)
}

// Try to read the error message for the sake of debugging

// Error message may have non-printable chars, so clean it up

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package testutil // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/testutil"

import (
	context "context"
	"testing"
	"time"

	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/entry"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
)

// NewMockOperator will return a basic operator mock
func NewMockOperator(id string) *Operator { _ = "STUB: not implemented"; return nil }

// FakeOutput is an empty output used primarily for testing
type FakeOutput struct {
	Received         chan *entry.Entry
	logger           *zap.Logger
	processWithError bool
}

// NewFakeOutput creates a new fake output with default settings
func NewFakeOutput(tb testing.TB) *FakeOutput { _ = "STUB: not implemented"; return nil }

// NewFakeOutputWithProcessError creates a new fake output with default settings, which returns error on Process
func NewFakeOutputWithProcessError(tb testing.TB) *FakeOutput {
	_ = "STUB: not implemented"
	return nil
}

// CanOutput always returns false for a fake output
func (*FakeOutput) CanOutput() bool {
	_ = "STUB: not implemented"

	// CanProcess always returns true for a fake output
	return false
}

func (*FakeOutput) CanProcess() bool {
	_ = "STUB: not implemented"

	// ID always returns `fake` as the ID of a fake output operator
	return false
}

func (*FakeOutput) ID() string {
	_ = "STUB: not implemented"

	// Logger returns the logger of a fake output
	return ""
}

func (f *FakeOutput) Logger() *zap.Logger {
	_ = "STUB: not implemented"

	// Outputs always returns nil for a fake output
	return nil
}

func (*FakeOutput) Outputs() []operator.Operator {
	_ = "STUB: not implemented"

	// Outputs always returns nil for a fake output
	return nil
}

func (*FakeOutput) GetOutputIDs() []string {
	_ = "STUB: not implemented"

	// SetOutputs immediately returns nil for a fake output
	return nil
}

func (*FakeOutput) SetOutputs([]operator.Operator) error {
	_ = "STUB: not implemented"

	// SetOutputIDs immediately returns nil for a fake output
	return nil
}

func (*FakeOutput) SetOutputIDs([]string) {
	_ = "STUB: not implemented"

	// Start immediately returns nil for a fake output
	return
}

func (*FakeOutput) Start(operator.Persister) error {
	_ = "STUB: not implemented"

	// Stop immediately returns nil for a fake output
	return nil
}

func (*FakeOutput) Stop() error {
	_ = "STUB: not implemented"

	// Type always return `fake_output` for a fake output
	return nil
}

func (*FakeOutput) Type() string { _ = "STUB: not implemented"; return "" }

func (f *FakeOutput) ProcessBatch(ctx context.Context, entries []*entry.Entry) error {
	_ = "STUB: not implemented"
	return nil
}

// Process will place all incoming entries on the Received channel of a fake output
func (f *FakeOutput) Process(_ context.Context, entry *entry.Entry) error {
	_ = "STUB: not implemented"
	return nil
}

// ExpectBody expects that a body will be received by the fake operator within a second
// and that it is equal to the given body
func (f *FakeOutput) ExpectBody(tb testing.TB, body any) { _ = "STUB: not implemented"; return }

// ExpectEntry expects that an entry will be received by the fake operator within a second
// and that it is equal to the given body
func (f *FakeOutput) ExpectEntry(tb testing.TB, expected *entry.Entry) {
	_ = "STUB: not implemented"
	return
}

// ExpectEntries expects that the given entries will be received in any order
func (f *FakeOutput) ExpectEntries(tb testing.TB, expected []*entry.Entry) {
	_ = "STUB: not implemented"
	return
}

// ExpectNoEntry expects that no entry will be received within the specified time
func (f *FakeOutput) ExpectNoEntry(tb testing.TB, timeout time.Duration) {
	_ = "STUB: not implemented"
	return
}

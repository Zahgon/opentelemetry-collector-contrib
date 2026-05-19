// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package carbonreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/carbonreceiver"

import (
	"context"

	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/receiverhelper"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/carbonreceiver/internal/transport"
)

// reporter struct implements the transport.Reporter interface to give consistent
// observability per Collector metric observability package.
type reporter struct {
	logger        *zap.Logger
	sugaredLogger *zap.SugaredLogger // Used for generic debug logging
	obsrecv       *receiverhelper.ObsReport
}

var _ transport.Reporter = (*reporter)(nil)

func newReporter(set receiver.Settings) (transport.Reporter, error) {
	_ = "STUB: not implemented"
	return *new(transport.Reporter), nil
}

// OnDataReceived is called when a message or request is received from
// a client. The returned context should be used in other calls to the same
// reporter instance. The caller code should include a call to end the
// returned span.
func (r *reporter) OnDataReceived(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// OnTranslationError is used to report a translation error from original
// format to the internal format of the Collector. The context and span
// passed to it should be the ones returned by OnDataReceived.
func (r *reporter) OnTranslationError(ctx context.Context, err error) {
	_ = "STUB: not implemented"
	return
}

// OnMetricsProcessed is called when the received data is passed to next
// consumer on the pipeline. The context and span passed to it should be the
// ones returned by OnDataReceived. The error should be error returned by
// the next consumer - the reporter is expected to handle nil error too.
func (r *reporter) OnMetricsProcessed(
	ctx context.Context,
	numReceivedMetricPoints int,
	err error,
) {
	_ = "STUB: not implemented"
	return
}

func (r *reporter) OnDebugf(template string, args ...any) { _ = "STUB: not implemented"; return }

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package arrow // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/otelarrowreceiver/internal/arrow"

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"

	arrowpb "github.com/open-telemetry/otel-arrow/go/api/experimental/arrow/v1"
	arrowRecord "github.com/open-telemetry/otel-arrow/go/pkg/otel/arrow_record"
	"go.opentelemetry.io/collector/client"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configgrpc"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/consumer/consumererror"
	"go.opentelemetry.io/collector/extension/extensionauth"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/receiverhelper"
	otelcodes "go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/net/http2/hpack"
	"google.golang.org/grpc"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/otelarrow/admission2"
	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/otelarrow/netstats"
)

const (
	streamFormat        = "arrow"
	hpackMaxDynamicSize = 4096
)

var (
	ErrNoMetricsConsumer   = errors.New("no metrics consumer")
	ErrNoLogsConsumer      = errors.New("no logs consumer")
	ErrNoTracesConsumer    = errors.New("no traces consumer")
	ErrUnrecognizedPayload = consumererror.NewPermanent(errors.New("unrecognized OTel-Arrow payload"))
)

type Consumers interface {
	Traces() consumer.Traces
	Metrics() consumer.Metrics
	Logs() consumer.Logs
}

type Receiver struct {
	Consumers

	arrowpb.UnsafeArrowTracesServiceServer
	arrowpb.UnsafeArrowLogsServiceServer
	arrowpb.UnsafeArrowMetricsServiceServer

	telemetry    component.TelemetrySettings
	tracer       trace.Tracer
	obsrecv      *receiverhelper.ObsReport
	gsettings    configgrpc.ServerConfig
	authServer   extensionauth.Server
	newConsumer  func() arrowRecord.ConsumerAPI
	netReporter  netstats.Interface
	boundedQueue admission2.Queue
}

// receiverStream holds the inFlightWG for a single stream.
type receiverStream struct {
	*Receiver
	inFlightWG sync.WaitGroup
}

// New creates a new Receiver reference.
func New(
	cs Consumers,
	set receiver.Settings,
	obsrecv *receiverhelper.ObsReport,
	gsettings configgrpc.ServerConfig,
	authServer extensionauth.Server,
	newConsumer func() arrowRecord.ConsumerAPI,
	bq admission2.Queue,
	netReporter netstats.Interface,
) (*Receiver, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// headerReceiver contains the state necessary to decode per-request metadata
// from an arrow stream.
type headerReceiver struct {
	// decoder maintains state across the stream.
	decoder *hpack.Decoder

	// includeMetadata as configured by gRPC settings.
	includeMetadata bool

	// hasAuthServer indicates that headers must be produced
	// independent of includeMetadata.
	hasAuthServer bool

	// client connection info from the stream context, (optionally
	// if includeMetadata) to be extended with per-request metadata.
	connInfo client.Info

	// streamHdrs was translated from the incoming context, will be
	// merged with per-request metadata.  Note that the contents of
	// this map are equivalent to connInfo.Metadata, however that
	// library does not let us iterate over the map so we recalculate
	// this from the gRPC incoming stream context.
	streamHdrs map[string][]string

	// tmpHdrs is used by the decoder's emit function during Write.
	tmpHdrs map[string][]string
}

func newHeaderReceiver(streamCtx context.Context, as extensionauth.Server, includeMetadata bool) *headerReceiver {
	_ = "STUB: not implemented"
	return nil
}

// Note that we capture the incoming context if there is an
// Auth plugin configured or includeMetadata is set.

// Note the hpack decoder supports additional protections,
// such as SetMaxStringLength(), but as we already have limits
// on stream request size, this seems unnecessary.

// combineHeaders calculates per-request Metadata by combining the stream's
// client.Info with additional key:values associated with the arrow batch.
func (h *headerReceiver) combineHeaders(ctx context.Context, hdrsBytes []byte) (context.Context, map[string][]string, error) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil, nil
}

// Note: call newContext in this case to ensure that
// connInfo is added to the context, for Auth.

// Note that we will parse the headers even if they are not
// used, to check for validity and/or trace context.  Also
// note this code was once optimized to avoid the following
// map allocation in cases where the return value would not be
// used.  This logic was "is metadata present" or "is auth
// server used".  Then we added to this, "is trace propagation
// in use" and simplified this function to always store the
// headers into a temporary map.

// Write calls the emitFunc, appending directly into `tmpHdrs`.

// Get the global propagator, to extract context.  When there
// are no fields, it's a no-op propagator implementation and
// we can skip the allocations inside this block.

// When there are no fields, it's a no-op
// implementation and we can skip the allocations.

// Add streamHdrs that were not carried in the per-request headers.

// Note: This is done after the per-request metadata is defined
// in recognition of a potential for duplicated values stemming
// from the Arrow exporter's independent call to the Auth
// extension's GetRequestMetadata().  This paired with the
// headersetter's return of empty-string values means, we would
// end up with an empty-string element for any headersetter
// `from_context` rules b/c the stream uses background context.
// This allows static headers through.
//
// See https://github.com/open-telemetry/opentelemetry-collector/issues/6965

// Release the temporary copy used in emitFunc().

// Note: newHdrs is passed to the Auth plugin.  Whether
// newHdrs is set in the context depends on h.includeMetadata.

// tmpHdrsAppend appends to tmpHdrs, from decoder's emit function.
func (h *headerReceiver) tmpHdrsAppend(hf hpack.HeaderField) { _ = "STUB: not implemented"; return }

// We force strings.ToLower to ensure consistency.  gRPC itself
// does this and would do the same.

func (h *headerReceiver) newContext(ctx context.Context, hdrs map[string][]string) context.Context {
	_ = "STUB: not implemented"
	// Retain the Addr/Auth of the stream connection, update the
	// per-request metadata from the Arrow batch.
	return *new(context.Context)
}

// logStreamError decides how to log an error.
func (r *Receiver) logStreamError(err error, where string) (occode otelcodes.Code, msg string) {
	_ = "STUB: not implemented"

	// gRPC tends to supply status-wrapped errors, so we always
	// unpack them.  A wrapped Canceled code indicates intentional
	// shutdown, which can be due to normal causes (EOF, e.g.,
	// max-stream-lifetime reached) or unusual causes (Canceled,
	// e.g., because the other stream direction reached an error).
	return *new(otelcodes.Code), ""
}

func gRPCName(desc grpc.ServiceDesc) string { _ = "STUB: not implemented"; return "" }

var (
	arrowTracesMethod  = gRPCName(arrowpb.ArrowTracesService_ServiceDesc)
	arrowMetricsMethod = gRPCName(arrowpb.ArrowMetricsService_ServiceDesc)
	arrowLogsMethod    = gRPCName(arrowpb.ArrowLogsService_ServiceDesc)
)

type signalType int

const (
	signalTraces signalType = iota + 1
	signalMetrics
	signalLogs
)

// method returns the gRPC method name for netstats reporting.
func (s signalType) method() string { _ = "STUB: not implemented"; return "" }

func (r *Receiver) ArrowTraces(serverStream arrowpb.ArrowTracesService_ArrowTracesServer) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Receiver) ArrowLogs(serverStream arrowpb.ArrowLogsService_ArrowLogsServer) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Receiver) ArrowMetrics(serverStream arrowpb.ArrowMetricsService_ArrowMetricsServer) error {
	_ = "STUB: not implemented"
	return nil
}

type anyStreamServer interface {
	Send(*arrowpb.BatchStatus) error
	Recv() (*arrowpb.BatchArrowRecords, error)
	grpc.ServerStream
}

type batchResp struct {
	id  int64
	err error
}

func (r *Receiver) recoverErr(retErr *error) { _ = "STUB: not implemented"; return }

// When this happens, the stacktrace is
// important and lost if we don't capture it
// here.

func (r *Receiver) anyStream(serverStream anyStreamServer, sig signalType) (retErr error) {
	_ = "STUB: not implemented"
	return nil
}

// doneCancel allows an error in the sender/receiver to
// interrupt the corresponding thread.

// wg is used to ensure this thread returns after both
// sender and receiver threads return.

// flushCtx controls the start of flushing.  when this is canceled
// after the receiver finishes, the flush operation begins.

// the sender receives flushCtx, which is canceled after the
// receiver returns (success or no).

// Wait for sender/receiver threads to return before returning.

// Before returning Canceled, drain any pending recvErrCh to
// avoid masking admission/resource errors with a shutdown race.

// the receiver returned EOF, next we
// expect the sender to finish.

// explicit cancel here, in case the sender fails before
// the receiver does. break the receiver loop here:

func (r *receiverStream) newInFlightData(ctx context.Context, sig signalType, batchID int64, pendingCh chan<- batchResp) *inFlightData {
	_ = "STUB: not implemented"
	return nil
}

// inFlightData is responsible for storing the resources held by one request.
type inFlightData struct {
	// Receiver is the owner of the resources held by this object.
	*receiverStream

	sig       signalType
	batchID   int64
	pendingCh chan<- batchResp
	span      trace.Span

	// refs counts the number of goroutines holding this object.
	// initially the recvOne() body, on success the
	// consumeAndRespond() function.
	refs atomic.Int32

	numItems   int   // how many items
	uncompSize int64 // uncompressed data size == how many bytes held in the semaphore
	releaser   admission2.ReleaseFunc
}

func (id *inFlightData) recvDone(ctx context.Context, recvErrPtr *error) {
	_ = "STUB: not implemented"
	return
}

// logStreamError because this response will break the stream.

func (id *inFlightData) consumeDone(ctx context.Context, consumeErrPtr *error) {
	_ = "STUB: not implemented"
	return
}

// debug-level because the error was external from the pipeline.

func (id *inFlightData) replyToCaller(ctx context.Context, callerErr error) {
	_ = "STUB: not implemented"
	return
}

// OK: Responded.

// OK: Never responded due to cancelation.

func (id *inFlightData) anyDone(ctx context.Context) {
	_ = "STUB: not implemented"
	// check if there are still refs, in which case leave the in-flight
	// counts where they are.
	return
}

// The netstats code knows that uncompressed size is
// unreliable for arrow transport, so we instrument it
// directly here.  Only the primary direction of transport
// is instrumented this way.

// recvOne begins processing a single Arrow batch.
//
// If an error is encountered before Arrow data is successfully consumed,
// the stream will break and the error will be returned immediately.
//
// If the error is due to authorization, the stream remains unbroken
// and the request fails.
//
// If not enough resources are available, the stream will block (if
// waiting permitted) or break (insufficient waiters).
//
// Assuming success, a new goroutine is created to handle consuming the
// data.
//
// This handles constructing an inFlightData object, which itself
// tracks everything that needs to be used by instrumentation when the
// batch finishes.
func (r *receiverStream) recvOne(streamCtx context.Context, serverStream anyStreamServer, hrcv *headerReceiver, pendingCh chan<- batchResp, sig signalType, ac arrowRecord.ConsumerAPI) (retErr error) {
	_ = "STUB: not implemented"
	// Receive a batch corresponding with one ptrace.Traces, pmetric.Metrics,
	// or plog.Logs item.
	return nil
}

// the incoming stream context is the parent of the in-flight context, which
// carries a span covering sequential stream-processing work.  the context
// is severed at this point, with flight.span a contextless child that will be
// finished in recvDone().

// inflightCtx is carried through into consumeAndProcess on the success path.
// this inherits the stream context so that its auth headers are present
// when the per-data Auth call is made.

// This is a special case to avoid introducing a span error
// for a canceled operation.

// This is a special case to avoid introducing a span error
// for a canceled operation.

// Note: err is directly from gRPC, should already have status.

// Check for optional headers and set the incoming context.

// Failing to parse the incoming headers breaks the stream.

// start this span after hrcv.combineHeaders returns extracted context. This will allow this span
// to be a part of the data path trace instead of only being included as a child of the stream inflight trace.

// Authorize the request, if configured, prior to acquiring resources.

// timeout parsed successfully

// if we return before the new goroutine is started below
// cancel the context.  callerCancel will be non-nil until
// the new goroutine is created at the end of this function.

// Use the bounded queue to memory limit based on incoming
// uncompressed request size and waiters.  Acquire will fail
// immediately if there are too many waiters, or will
// otherwise block until timeout or enough memory becomes
// available.

// Recognize that the request is still in-flight via consumeAndRespond()

// consumeAndRespond consumes the data and returns control to the sender loop.

// Reset callerCancel so the deferred function above does not call it here.

// consumeAndRespond finishes the span started in recvOne and logs the
// result after invoking the pipeline to consume the data.
func (r *Receiver) consumeAndRespond(ctx, streamCtx context.Context, data any, flight *inFlightData) {
	_ = "STUB: not implemented"
	return
}

// recoverErr is a special function because it recovers panics, so we
// keep it in a separate defer than the processing above, which will
// run after the panic is recovered into an ordinary error.

// srvReceiveLoop repeatedly receives one batch of data.
func (r *receiverStream) srvReceiveLoop(ctx context.Context, serverStream anyStreamServer, pendingCh chan<- batchResp, sig signalType, ac arrowRecord.ConsumerAPI) (retErr error) {
	_ = "STUB: not implemented"
	return nil
}

// srvReceiveLoop repeatedly sends one batch data response.
func (r *receiverStream) sendOne(serverStream anyStreamServer, resp batchResp) error {
	_ = "STUB: not implemented"
	// Note: Statuses can be batched, but we do not take
	// advantage of this feature.
	return nil
}

// Generally, code in the receiver should use
// status.Errorf(codes.XXX, ...)  so that we take the
// first branch.

// Ideally, we don't take this branch because all code uses
// gRPC status constructors and we've taken the branch above.
//
// This is a fallback for several broad categories of error.

// Some kind of pipeline error, somewhere downstream.

// Probably a pipeline error, retryable.

// logStreamError because this response will break the stream.

func (r *receiverStream) flushSender(serverStream anyStreamServer, recvWG *sync.WaitGroup, pendingCh <-chan batchResp) error {
	_ = "STUB: not implemented"
	// wait to ensure no more items are accepted
	return nil
}

// wait for all responses to be sent

// Currently nothing left in pendingCh.

func (r *receiverStream) srvSendLoop(ctx context.Context, serverStream anyStreamServer, recvWG *sync.WaitGroup, pendingCh <-chan batchResp) error {
	_ = "STUB: not implemented"
	return nil
}

// consumeBatch applies the batch to the Arrow Consumer, returns a
// slice of pdata objects of the corresponding data type as `any`.
// along with the number of items and true uncompressed size.
//
// The signal type is determined by the gRPC service method (e.g.,
// ArrowTraces, ArrowLogs, ArrowMetrics)
func (r *Receiver) consumeBatch(arrowConsumer arrowRecord.ConsumerAPI, records *arrowpb.BatchArrowRecords, sig signalType) (retData any, numItems int, uncompSize int64, retErr error) {
	_ = "STUB: not implemented"
	return *new(any), 0, 0, nil
}

// consumeData invokes the next pipeline consumer for a received batch of data.
// it uses the standard OTel collector instrumentation (receiverhelper.ObsReport).
//
// if any errors are permanent, returns a permanent error.
func (r *Receiver) consumeData(ctx context.Context, data any, flight *inFlightData) (retErr error) {
	_ = "STUB: not implemented"
	return nil
}

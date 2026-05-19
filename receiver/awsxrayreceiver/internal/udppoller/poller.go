// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package udppoller // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsxrayreceiver/internal/udppoller"

import (
	"context"
	"sync"

	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/receiverhelper"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsxrayreceiver/internal/socketconn"
)

const (
	// Transport is the network transport protocol used
	// by the poller
	Transport = "udp"

	// size of the buffer used by each poller.
	// https://github.com/aws/aws-xray-daemon/blob/master/pkg/cfg/cfg.go#L182
	// https://github.com/aws/aws-xray-daemon/blob/master/cmd/tracing/daemon.go#L171
	pollerBufferSizeKB = 64 * 1024

	// the size of the channel between the UDP poller
	// and OT consumer
	segChanSize = 30
)

// Poller represents one or more goroutines that are
// polling from a UDP socket
type Poller interface {
	SegmentsChan() <-chan RawSegment
	Start(receiverLongTermCtx context.Context)
	Close() error
}

// RawSegment represents a raw X-Ray segment document.
type RawSegment struct {
	// Payload is the raw bytes that represent one X-Ray segment.
	Payload []byte
	// Ctx is the short-lived context created per raw segment received
	Ctx context.Context
}

// Config represents the configurations needed to
// start the UDP poller
type Config struct {
	Transport          string
	Endpoint           string
	NumOfPollerToStart int
}

type poller struct {
	udpSock              socketconn.SocketConn
	logger               *zap.Logger
	wg                   sync.WaitGroup
	receiverLongLivedCtx context.Context
	maxPollerCount       int
	// closing this channel will shutdown all goroutines
	// within this poller
	shutDown chan struct{}

	// all segments read by the poller will be sent to this channel
	segChan chan RawSegment

	obsrecv *receiverhelper.ObsReport
}

// New creates a new UDP poller
func New(cfg *Config, set receiver.Settings) (Poller, error) {
	_ = "STUB: not implemented"
	return *new(Poller), nil
}

func (p *poller) Start(receiverLongTermCtx context.Context) { _ = "STUB: not implemented"; return }

func (p *poller) Close() error { _ = "STUB: not implemented"; return nil }

// inform the consumers of segChan that the poller is stopped

func (p *poller) SegmentsChan() <-chan RawSegment { _ = "STUB: not implemented"; return nil }

func (p *poller) read(buf *[]byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (p *poller) poll() { _ = "STUB: not implemented"; return }

// TODO: We may want to attempt to shutdown/clean the broken socket and open a new one
// with the same address

// For now tracesegment.SplitHeaderBody does not return irrecoverable error
// so we don't check for it

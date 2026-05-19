// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package tcp // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/input/tcp"

import (
	"bufio"
	"context"
	"crypto/tls"
	"net"
	"sync"

	"github.com/jpillora/backoff"
	"golang.org/x/text/encoding"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"
)

// Input is an operator that listens for log entries over tcp.
type Input struct {
	helper.InputOperator
	address         string
	MaxLogSize      int
	addAttributes   bool
	OneLogPerPacket bool

	listener net.Listener
	cancel   context.CancelFunc
	wg       sync.WaitGroup
	tls      *tls.Config
	backoff  backoff.Backoff

	encoding  encoding.Encoding
	splitFunc bufio.SplitFunc
	resolver  *helper.IPResolver
}

// Start will start listening for log entries over tcp.
func (i *Input) Start(_ operator.Persister) error { _ = "STUB: not implemented"; return nil }

func (i *Input) configureListener() error { _ = "STUB: not implemented"; return nil }

// goListenn will listen for tcp connections.
func (i *Input) goListen(ctx context.Context) { _ = "STUB: not implemented"; return }

// goHandleClose will wait for the context to finish before closing a connection.
func (i *Input) goHandleClose(ctx context.Context, conn net.Conn) {
	_ = "STUB: not implemented"
	return
}

// goHandleMessages will handles messages from a tcp connection.
func (i *Input) goHandleMessages(ctx context.Context, conn net.Conn, cancel context.CancelFunc) {
	_ = "STUB: not implemented"
	return
}

func (i *Input) handleMessage(ctx context.Context, conn net.Conn, dec *encoding.Decoder, log []byte) {
	_ = "STUB: not implemented"
	return
}

func truncateMaxLog(data []byte, maxLogSize int) (token []byte) {
	_ = "STUB: not implemented"
	return nil
}

// Stop will stop listening for log entries over TCP.
func (i *Input) Stop() error { _ = "STUB: not implemented"; return nil }

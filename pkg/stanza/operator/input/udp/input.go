// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package udp // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/input/udp"

import (
	"bufio"
	"context"
	"net"
	"sync"

	"golang.org/x/text/encoding"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"
)

// Input is an operator that listens to a socket for log entries.
type Input struct {
	buffer []byte
	helper.InputOperator
	address         *net.UDPAddr
	addAttributes   bool
	OneLogPerPacket bool
	AsyncConfig     *AsyncConfig

	connection net.PacketConn
	cancel     context.CancelFunc
	wg         sync.WaitGroup
	wgReader   sync.WaitGroup

	encoding  encoding.Encoding
	splitFunc bufio.SplitFunc
	resolver  *helper.IPResolver

	messageQueue   chan messageAndAddress
	readBufferPool sync.Pool
	stopOnce       sync.Once
}

type messageAndAddress struct {
	Message       *[]byte
	RemoteAddr    net.Addr
	MessageLength int
}

// Start will start listening for messages on a socket.
func (i *Input) Start(_ operator.Persister) error { _ = "STUB: not implemented"; return nil }

// goHandleMessages will handle messages from a udp connection.
func (i *Input) goHandleMessages(ctx context.Context) { _ = "STUB: not implemented"; return }

func (i *Input) readAndProcessMessages(ctx context.Context) { _ = "STUB: not implemented"; return }

func (i *Input) processMessage(ctx context.Context, message []byte, remoteAddr net.Addr, dec *encoding.Decoder, scannerBuffer []byte) {
	_ = "STUB: not implemented"
	return
}

func (i *Input) readMessagesAsync(ctx context.Context) { _ = "STUB: not implemented"; return }

// Can't reuse the same buffer since same references would be written multiple times to the messageQueue (and cause data override of previous entries)

// Send the message to the message queue for processing

func (i *Input) processMessagesAsync(ctx context.Context) { _ = "STUB: not implemented"; return }

// Read a message from the message queue.

// Channel closed, exit the goroutine.

func truncateMaxLog(data []byte) (token []byte) { _ = "STUB: not implemented"; return nil }

func (i *Input) handleMessage(ctx context.Context, remoteAddr net.Addr, dec *encoding.Decoder, log []byte) {
	_ = "STUB: not implemented"
	return
}

// readMessage will read log messages from the connection.
func (i *Input) readMessage(buffer []byte) ([]byte, net.Addr, int, error) {
	_ = "STUB: not implemented"
	return nil, *new(net.Addr), 0, nil
}

// This will remove trailing characters and NULs from the buffer
func (*Input) removeTrailingCharactersAndNULsFromBuffer(buffer []byte, n int) []byte {
	_ = "STUB: not implemented"
	// Remove trailing characters and NULs
	return nil
}

//nolint:revive

// Stop will stop listening for udp messages.
func (i *Input) Stop() error { _ = "STUB: not implemented"; return nil }

// only when all async readers are finished, so there's no risk of sending to a closed channel, do we close messageQueue (which allows the async processors to finish)

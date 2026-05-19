// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package emittest // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/fileconsumer/internal/emittest"

import (
	"testing"
	"time"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/fileconsumer/emit"
)

type sinkCfg struct {
	emitChanLen int
	timeout     time.Duration
}

type SinkOpt func(*sinkCfg)

type Sink struct {
	emitChan chan emit.Token
	timeout  time.Duration
	emit.Callback
}

func WithCallBuffer(n int) SinkOpt { _ = "STUB: not implemented"; return *new(SinkOpt) }

func WithTimeout(d time.Duration) SinkOpt { _ = "STUB: not implemented"; return *new(SinkOpt) }

func NewSink(opts ...SinkOpt) *Sink { _ = "STUB: not implemented"; return nil }

func (s *Sink) NextToken(t *testing.T) []byte { _ = "STUB: not implemented"; return nil }

func (s *Sink) NextTokens(t *testing.T, n int) [][]byte { _ = "STUB: not implemented"; return nil }

func (s *Sink) NextCall(t *testing.T) ([]byte, map[string]any) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Sink) ExpectToken(t *testing.T, expected []byte) { _ = "STUB: not implemented"; return }

func (s *Sink) ExpectTokens(t *testing.T, expected ...[]byte) { _ = "STUB: not implemented"; return }

func (s *Sink) ExpectCall(t *testing.T, expected []byte, attrs map[string]any) {
	_ = "STUB: not implemented"
	return
}

func (s *Sink) ExpectCalls(t *testing.T, expected ...emit.Token) { _ = "STUB: not implemented"; return }

func (s *Sink) ExpectNoCalls(t *testing.T) { _ = "STUB: not implemented"; return }

func (s *Sink) ExpectNoCallsUntil(t *testing.T, d time.Duration) { _ = "STUB: not implemented"; return }

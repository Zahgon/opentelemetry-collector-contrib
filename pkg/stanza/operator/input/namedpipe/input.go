// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build linux

package namedpipe // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/input/namedpipe"

import (
	"bufio"
	"context"
	"os"
	"sync"
	"time"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/trim"
)

const ReadTimeout = 2 * time.Second

type Input struct {
	helper.InputOperator

	buffer      []byte
	path        string
	permissions uint32
	splitFunc   bufio.SplitFunc
	trimFunc    trim.Func
	cancel      context.CancelFunc
	pipe        *os.File
	wg          sync.WaitGroup
}

func (i *Input) Start(_ operator.Persister) error { _ = "STUB: not implemented"; return nil }

// chmod the named pipe because mkfifo respects the umask which may result
// in a named pipe with incorrect permissions.

// Open the pipe with O_RDWR so it won't block on opening the pipe when there is no writer.
// The current process is both a writer and reader, which prevents the read from receiving
// EOF because there is always a writer (the process itself) connects to the pipe.

func (i *Input) Stop() error { _ = "STUB: not implemented"; return nil }

func (i *Input) readLoop(ctx context.Context) { _ = "STUB: not implemented"; return }

// The process exits due to whatever reason, wait for ReadTimeout and try again.

func (i *Input) process(ctx context.Context, pipe *os.File) error {
	_ = "STUB: not implemented"
	return nil
}

// sendEntry sends an entry to the next operator in the pipeline.
func (i *Input) sendEntry(ctx context.Context, bytes []byte) error {
	_ = "STUB: not implemented"
	return nil
}

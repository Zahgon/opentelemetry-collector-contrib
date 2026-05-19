// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package zstd // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/otelarrow/compression/zstd"

import (
	"io"
	"sync"

	zstdlib "github.com/klauspost/compress/zstd"
	"google.golang.org/grpc"
	"google.golang.org/grpc/encoding"
)

// NamePrefix is prefix, with N for compression level.
const NamePrefix = "zstdarrow"

// Level is an integer value mapping to compression level.
// [0] implies disablement; not registered in grpc
// [1,2] fastest i.e., "zstdarrow1", "zstdarrow2"
// [3-5] default
// [6-9] better
// [10] best.
type Level uint

const (
	// DefaultLevel is a reasonable balance of compression and cpu usage.
	DefaultLevel Level = 5
	// MinLevel is fast and cheap.
	MinLevel Level = 1
	// MaxLevel is slow and expensive.
	MaxLevel Level = 10
)

type EncoderConfig struct {
	// Level is meaningful in the range [0, 10].  No invalid
	// values, they all map into 4 default configurations.  (default: 5)
	// See `zstdlib.WithEncoderLevel()`.
	Level Level `mapstructure:"level"`
	// WindowSizeMiB is a Zstd-library parameter that controls how
	// much window of text is visible to the compressor at a time.
	// It is the dominant factor that determines memory usage.
	// If zero, the window size is determined by level.  (default: 0)
	// See `zstdlib.WithWindowSize()`.
	WindowSizeMiB uint32 `mapstructure:"window_size_mib"`
	// Concurrency is a Zstd-library parameter that configures the
	// use of background goroutines to improve compression speed.
	// 0 means to let the library decide (it will use up to GOMAXPROCS),
	// and 1 means to avoid background workers.  (default: 1)
	// See `zstdlib.WithEncoderConcurrency()`.
	Concurrency uint `mapstructure:"concurrency"`
}

type DecoderConfig struct {
	// MemoryLimitMiB is a memory limit control for the decoder,
	// as a way to limit overall memory use by Zstd.
	// See `zstdlib.WithDecoderMaxMemory()`.
	MemoryLimitMiB uint32 `mapstructure:"memory_limit_mib"`
	// MaxWindowSizeMiB limits window sizes that can be configured
	// in the corresponding encoder's `EncoderConfig.WindowSizeMiB`
	// setting, as a way to control memory usage.
	// See `zstdlib.WithDecoderMaxWindow()`.
	MaxWindowSizeMiB uint32 `mapstructure:"max_window_size_mib"`
	// Concurrency is a Zstd-library parameter that configures the
	// use of background goroutines to improve decompression speed.
	// 0 means to let the library decide (it will use up to GOMAXPROCS),
	// and 1 means to avoid background workers.  (default: 1)
	// See `zstdlib.WithDecoderConcurrency()`.
	Concurrency uint `mapstructure:"concurrency"`
}

type encoder struct {
	lock sync.Mutex // protects cfg
	cfg  EncoderConfig
	pool mru[*writer]
}

type decoder struct {
	lock sync.Mutex // protects cfg
	cfg  DecoderConfig
	pool mru[*reader]
}

type reader struct {
	*zstdlib.Decoder
	Gen
	pool *mru[*reader]
}

type writer struct {
	*zstdlib.Encoder
	Gen
	pool *mru[*writer]
}

type combined struct {
	enc encoder
	dec decoder
}

type instance struct {
	lock    sync.Mutex
	byLevel map[Level]*combined
}

var _ encoding.Compressor = &combined{}

var staticInstances = &instance{
	byLevel: map[Level]*combined{},
}

func (g *Gen) generation() Gen { _ = "STUB: not implemented"; return *new(Gen) }

func DefaultEncoderConfig() EncoderConfig { _ = "STUB: not implemented"; return *new(EncoderConfig) }

// Determines other defaults
// Avoids extra CPU/memory

func DefaultDecoderConfig() DecoderConfig { _ = "STUB: not implemented"; return *new(DecoderConfig) }

// Avoids extra CPU/memory
// More conservative than library default
// Corresponds w/ "best" level default

func validate(level Level, f func() error) error { _ = "STUB: not implemented"; return nil }

func (cfg EncoderConfig) Validate() error { _ = "STUB: not implemented"; return nil }

func (cfg DecoderConfig) Validate() error { _ = "STUB: not implemented"; return nil }

func init() {
	staticInstances.lock.Lock()
	defer staticInstances.lock.Unlock()
	resetLibrary()
}

func resetLibrary() { _ = "STUB: not implemented"; return }

func SetEncoderConfig(cfg EncoderConfig) error { _ = "STUB: not implemented"; return nil }

func SetDecoderConfig(cfg DecoderConfig) error { _ = "STUB: not implemented"; return nil }

func (cfg EncoderConfig) options() (opts []zstdlib.EOption) { _ = "STUB: not implemented"; return nil }

func (e *encoder) getConfig() EncoderConfig { _ = "STUB: not implemented"; return *new(EncoderConfig) }

func (cfg EncoderConfig) Name() string { _ = "STUB: not implemented"; return "" }

func (cfg EncoderConfig) CallOption() grpc.CallOption {
	_ = "STUB: not implemented"
	return *new(grpc.CallOption)
}

func (cfg DecoderConfig) options() (opts []zstdlib.DOption) { _ = "STUB: not implemented"; return nil }

func (d *decoder) getConfig() DecoderConfig { _ = "STUB: not implemented"; return *new(DecoderConfig) }

func (c *combined) Compress(w io.Writer) (io.WriteCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.WriteCloser), nil
}

func (w *writer) Close() error { _ = "STUB: not implemented"; return nil }

func (c *combined) Decompress(r io.Reader) (io.Reader, error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), nil
}

// zstd decoders need to be closed when they are evicted from
// the freelist. Note that the finalizer is attached to the
// reader object, not to the decoder, because zstd maintains
// background references to the decoder that prevent it from
// being GC'ed.

func (r *reader) Read(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (c *combined) Name() string { _ = "STUB: not implemented"; return "" }

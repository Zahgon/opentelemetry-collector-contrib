// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package internal // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/stefexporter/internal"

import (
	"context"
	"sync/atomic"
	"time"

	"github.com/jonboulle/clockwork"
	"go.uber.org/zap"
)

// ConnManager manages a pool of connections.
//
// It is responsible for creating connections, reconnecting or flushing connections
// in the background when needed, and ensuring each connection is used exclusively
// by one user at a time.
//
// In order to use the connections call Acquire, Release and DiscardAndClose methods.
type ConnManager struct {
	logger *zap.Logger

	// clock is used for mocking the clock for testing purposes.
	// Outside of tests it is a real clock.
	clock clockwork.Clock

	targetConnCount uint

	// Number of current connections, collectively in all pools or acquired.
	curConnCount atomic.Int64

	connCreator ConnCreator

	// Connection pools. curConnCount connections are either in one
	// of these pools, or are acquired temporarily,
	// i.e. curConnCount = len(idleConns) + len(recreateConns) + (acquired count).
	idleConns     chan *ManagedConn // Ready to be acquired.
	recreateConns chan *ManagedConn // Pending to be recreated.

	flushPeriod     time.Duration
	reconnectPeriod time.Duration

	// Flags to indicate if the goroutines are stopped.
	flusherStopped         bool
	durationLimiterStopped bool
	recreatorStopped       bool
	// stoppedCond is used to wait until all goroutines are stopped.
	stoppedCond *CancellableCond

	// stopSignal is used to signal all goroutines to stop.
	stopSignal chan struct{}
}

// ManagedConn wraps a Conn and keeps some tracking information that
// ConnManager needs.
type ManagedConn struct {
	// The underlying connection.
	conn Conn

	// The time when the connection was created.
	startTime time.Time

	// The time when the connection was last flushed. Set after conn.Flush() is called.
	lastFlush time.Time

	// needsFlush indicates if the connection needs to be flushed at the next periodic
	// background check that runs at the flushPeriod intervals.
	needsFlush bool

	// Keep track of the connection usage, whether it is acquired or no.
	isAcquired bool
}

// Conn returns the underlying connection.
func (c *ManagedConn) Conn() Conn {
	_ = "STUB: not implemented"

	// ConnCreator allow creating connections.
	return *new(Conn)
}

type ConnCreator interface {
	// Create a new connection. May be called concurrently.
	// The attempt to create the connection should be cancelled if ctx is done.
	Create(ctx context.Context) (Conn, error)
}

// Conn represents a connection that can be closed or flushed.
type Conn interface {
	// Close the connection. The connection will be discarded
	// after this call returns.
	Close(ctx context.Context) error

	// Flush the connection. This is typically to send any buffered data.
	// Will be called periodically (see ConnManager flushPeriod) and
	// before ConnManager.Stop returns.
	Flush(context.Context) error
}

type ConnManagerSettings struct {
	Logger *zap.Logger

	// Creator helps create new connections.
	Creator ConnCreator

	// TargetConnCount is the number of connections desirable to maintain.
	// Must be >0.
	TargetConnCount uint

	// FlushPeriod is the approximate period to wait before flushing connections after
	// the user releases the connection. Setting this value to >0 avoids
	// unnecessary flushes when the connection is acquired and released frequently.
	// Setting this to 0 ensures flushing is done after every connection release.
	FlushPeriod time.Duration

	// ReconnectPeriod is the interval to reconnect connections. Each connection is
	// periodically reconnected approximately every reconnectPeriod.
	ReconnectPeriod time.Duration
}

func NewConnManager(set ConnManagerSettings) (*ConnManager, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Start starts the connection manager. It will immediately start
// creating targetConnCount new connections by calling connCreator.Create().
// All successfully created connections will become available for acquisition.
func (c *ConnManager) Start() {
	_ = "STUB: not implemented"
	// Put some dummy connections in recreateConns pool and let the recreator()
	// to replace them by proper connections in the background.
	return
}

// Stop stops the connection manager. It will wait until all acquired
// connections are returned. Then it will flush connections that
// are marked as needing to flush, and then will close all connections.
// Stop will be aborted if the context is done before all connections are closed.
func (c *ConnManager) Stop(ctx context.Context) error {
	_ = "STUB: not implemented"
	// Signal goroutines to stop
	return nil
}

// Wait until all goroutines stop

// Close all connections.
func (c *ConnManager) closeAll(ctx context.Context) error {
	_ = "STUB: not implemented"
	// We must close exactly curConnCount connections in total.
	// All goroutines are stopped at this point, so they won't interfere.
	// All connections are either in the idleConns or recreateConns pools
	// or are acquired and will be returned to one of the pools soon.
	return nil
}

// Get a connection from one of the connection pools.

// Connections in recreateConns are discarded and are candidates for recreation.
// We don't need to flush them.

// Flush if needs a flush and is not discarded.

// And close the connection.

// Join all errors (if any)

// Acquire an idle connection for exclusive use.
//
// Must call Release() or DiscardAndClose() when done.
// Returns an error if the connection is not available til ctx is done
// or if the manager is stopped.
func (c *ConnManager) Acquire(ctx context.Context) (*ManagedConn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Release returns a previously acquired connection to the ConnManager
// and makes it available to be acquired again.
//
// If the connection was last flushed more than flushPeriod ago, it will
// be flushed otherwise it will be marked as needing to be flushed at the
// next opportunity.
func (c *ConnManager) Release(ctx context.Context, conn *ManagedConn) {
	_ = "STUB: not implemented"
	return
}

// Time to flush the connection.

// Something went wrong, we need to recreate the connection since it
// may no longer be usable.

// Remember that it needs to be flushed sometime in the future.

// DiscardAndClose discards the acquired connection and closes it.
// This is normally used when the connection goes bad in some way
// and should not be reused.
// This will result in calling the Close method on the connection.
func (c *ConnManager) DiscardAndClose(conn *ManagedConn) { _ = "STUB: not implemented"; return }

func (c *ConnManager) flusher() { _ = "STUB: not implemented"; return }

// flusher is not needed, we will always flush immediately in Release().

// Context that cancels on stopSignal.

// Flush all idle connections that need flushing.

// Get one idle conn (if any).

// Something went wrong, we need to recreate the connection since it
// may no longer be usable.

// Return to idle pool.

// No more available idle connections.

// durationLimiter periodically checks connections and reconnects them if they
// were connected for more than reconnectPeriod. It will stagger the reconnections
// to avoid all connections reconnecting at the same c.clock.
func (c *ConnManager) durationLimiter() { _ = "STUB: not implemented"; return }

// Context that cancels on stopSignal.

// Each connection will be reconnected at approximately reconnectPeriod interval.
// We reconnect one per tick.

// Find an idle connection

// Check if it is time to reconnect.

// Flush it first.

// Send it for reconnection (regardless of whether it was flushed or not).

// Put it back, too soon to reconnect.

// recreator closes connections from recreateConns pool
// and replaces them by new connections.
func (c *ConnManager) recreator() {
	_ = "STUB: not implemented"

	// Indicate we are stopped on exit.
	return
}

// Track the number of connections.

// Close the connection.

// contextFromStopSignal creates a context that is cancelled when the
// provided channel is closed.
func contextFromStopSignal(stopSignal <-chan struct{}) (context.Context, context.CancelFunc) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(context.CancelFunc)
}

// createNewConn creates a new connection and adds it to the idleConns pool.
// It will keep trying with backoff until connection creation succeeds or
// until ConnManager is stopped.
func (c *ConnManager) createNewConn() { _ = "STUB: not implemented"; return }

// Create a context that will be cancelled when the manager is stopped.

// Try to connect, retrying until succeeding, with exponential backoff.

// Wait until the next retry or until the manager is stopped.

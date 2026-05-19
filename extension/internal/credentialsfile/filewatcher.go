// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package credentialsfile // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/internal/credentialsfile"

import (
	"context"
	"sync/atomic"

	"github.com/fsnotify/fsnotify"
	"go.uber.org/zap"
)

// fileWatcher implements ValueResolver by watching a file for changes
// and caching its contents atomically.
type fileWatcher struct {
	path       string
	value      atomic.Pointer[string]
	logger     *zap.Logger
	onChange   func(string)
	shutdownCH chan struct{}
	doneCH     chan struct{}
}

func newFileWatcher(path string, logger *zap.Logger, onChange func(string)) *fileWatcher {
	_ = "STUB: not implemented"
	return nil
}

func (w *fileWatcher) Value() string { _ = "STUB: not implemented"; return "" }

func (w *fileWatcher) Start(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Watch the file itself. For symlinked files (e.g. Kubernetes projected volumes),
// fsnotify follows the symlink and watches the underlying inode. On Remove/Chmod
// events the watcher is re-added to follow the new symlink target.

func (w *fileWatcher) Shutdown() error { _ = "STUB: not implemented"; return nil }

func (w *fileWatcher) watch(ctx context.Context, watcher *fsnotify.Watcher) {
	_ = "STUB: not implemented"
	return
}

// NOTE: Kubernetes projected volumes use symlinks. When the backing
// symlink target is removed, fsnotify may auto-remove the watch.
// We must re-add it so it follows the new symlink target.
// See: https://martensson.io/go-fsnotify-and-kubernetes-configmaps/

func (w *fileWatcher) reloadQuietly() { _ = "STUB: not implemented"; return }

func (w *fileWatcher) reload() error { _ = "STUB: not implemented"; return nil }

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build unix && !aix && !solaris

package reader // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/fileconsumer/internal/reader"

func (r *Reader) tryLockFile() bool { _ = "STUB: not implemented"; return false }

func (r *Reader) unlockFile() { _ = "STUB: not implemented"; return }

// If delete_after_read is set then the file may already have been deleted by this point,
// in which case we'll get EBADF.  This is harmless and not worth logging.

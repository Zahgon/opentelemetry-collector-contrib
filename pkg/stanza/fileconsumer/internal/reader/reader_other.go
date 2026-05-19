// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build !unix || aix || solaris

package reader // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/fileconsumer/internal/reader"

func (*Reader) tryLockFile() bool { _ = "STUB: not implemented"; return false }

func (*Reader) unlockFile() { _ = "STUB: not implemented"; return }

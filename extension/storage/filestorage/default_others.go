// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build !windows

package filestorage // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/storage/filestorage"

func getDefaultDirectory() string { _ = "STUB: not implemented"; return "" }

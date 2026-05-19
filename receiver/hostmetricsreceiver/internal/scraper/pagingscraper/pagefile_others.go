// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build !windows && !linux

package pagingscraper // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/pagingscraper"

func getPageFileStats() ([]*pageFileStats, error) { _ = "STUB: not implemented"; return nil, nil }

// We do not support per-device swap

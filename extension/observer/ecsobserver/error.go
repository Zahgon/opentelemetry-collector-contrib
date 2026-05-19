// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ecsobserver // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/ecsobserver"

import (
	"go.uber.org/zap"
)

// error.go defines common error interfaces and util methods for generating reports
// for log and metrics that can be used for debugging.

const (
	errKeyTask   = "task"
	errKeyTarget = "target"
)

type errWithAttributes interface {
	// message does not include attributes like task arn etc.
	// and expect the caller extract them using getters.
	message() string
	// zapFields will be logged as json attribute and allows searching and filter backend like cloudwatch.
	// For example { $.ErrScope == "Target" } list all the error whose scope is a (scrape) target.
	zapFields() []zap.Field
}

// hasCriticalError returns first critical error.
// Currently only access error and cluster not found are treated as critical.
func hasCriticalError(logger *zap.Logger, err error) error { _ = "STUB: not implemented"; return nil }

// fake a multi error

func printErrors(logger *zap.Logger, err error) { _ = "STUB: not implemented"; return }

// Use the short message, this makes searching the code via error message easier
// as additional info are flushed as fields.

func extractErrorFields(err error) ([]zap.Field, string) { _ = "STUB: not implemented"; return nil, "" }

// Stop early because we are only attaching value for our internal errors.

// Rename ok to tok because linter says it shadows outer ok.
// Though the linter seems to allow the similar block to shadow...

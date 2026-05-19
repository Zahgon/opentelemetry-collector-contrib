// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package dbstorage // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/storage/dbstorage"

import (
	"go.uber.org/zap"
)

const (
	sqliteNativePragmaOpt = "_pragma"
	sqliteNativeTimeOpt   = "_time_format"
	sqliteNativeTXLockOpt = "_txlock"
	sqliteNativeModeOpt   = "mode"
)

var sqlitePragmaOptsMapping = map[string]string{
	"_auto_vacuum":              "auto_vacuum",
	"_vacuum":                   "auto_vacuum",
	"_busy_timeout":             "busy_timeout",
	"_timeout":                  "busy_timeout",
	"_case_sensitive_like":      "case_sensitive_like",
	"_cslike":                   "case_sensitive_like",
	"_defer_foreign_keys":       "defer_foreign_keys",
	"_defer_fk":                 "defer_foreign_keys",
	"_foreign_keys":             "foreign_keys",
	"_fk":                       "foreign_keys",
	"_ignore_check_constraints": "ignore_check_constraints",
	"_journal_mode":             "journal_mode",
	"_journal":                  "journal_mode",
	"_locking_mode":             "locking_mode",
	"_locking":                  "locking_mode",
	"_query_only":               "query_only",
	"_recursive_triggers":       "recursive_triggers",
	"_rt":                       "recursive_triggers",
	"_secure_delete":            "secure_delete",
	"_synchronous":              "synchronous",
	"_sync":                     "synchronous",
	"_cache_size":               "cache_size",
	"_writable_schema":          "writable_schema",
}

func replaceCompatDSNOptions(logger *zap.Logger, dsn string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// If no query params present in DSN - we have nothing to do here

// This is unrecoverable error we should stop processing
// `sqlite` driver is using the same approach for options parsing

// Native driver options, no need to process

// Set of options that could be converter to new _pragma option

// Unknown or non-conversable option - add to errors

// Convert options back to query string and substitute it in DSN

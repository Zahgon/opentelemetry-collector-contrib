// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package operationsmanagement // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/bmchelixexporter/internal/operationsmanagement"

// allowedEntityValueChars contains all special characters allowed in entity values
// (entityTypeId, entityName) besides alphanumerics and whitespace.
// https://docs.bmc.com/xwiki/bin/view/IT-Operations-Management/Operations-Management/BMC-Helix-Operations-Management/bhom244/Policy-event-data-and-metric-data-management-endpoints-in-the-REST-API/Metric-operation-management-endpoints-in-the-REST-API/
const allowedEntityValueChars = "~!@#$^&*()-_=[];'?./\\"

// NormalizeEntityValue normalizes entity values (entityTypeId, entityName) to contain only valid characters.
// Allowed characters: a-zA-Z0-9~!@#$^&*()-_=[];'?./\ and whitespace
// Invalid characters are replaced with underscores.
func NormalizeEntityValue(value string) string { _ = "STUB: not implemented"; return "" }

// Replace all invalid characters with underscores

// sanitizeEntityValueRune returns the rune if it's valid for an entity value,
// otherwise returns '_'.
// Valid characters: letters, digits, whitespace, and special chars in allowedEntityValueChars
func sanitizeEntityValueRune(r rune) rune { _ = "STUB: not implemented"; return 0 }

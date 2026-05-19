// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package azurelogs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/azurelogs"

// az.service_request_id was removed from semconv in v1.36.0 without a replacement
// ("removed due to lack of applicability or use"). Keeping as a custom attribute
// for backward compatibility. TODO: revisit if semconv adds an equivalent.
const azServiceRequestIDKey = "az.service_request_id"

// TODO @constanca-m remove this file once the logic for the remaining categories
// is added to category_logs.go

func handleFrontDoorAccessLog(field string, value any, attrs, attrsProps map[string]any) {
	_ = "STUB: not implemented"
	return
}

// TODO Should be a port

func handleFrontDoorHealthProbeLog(field string, value any, attrs, attrsProps map[string]any) {
	_ = "STUB: not implemented"
	return
}

func handleAppServiceAppLogs(field string, value any, attrs, attrsProps map[string]any) {
	_ = "STUB: not implemented"
	return
}

func handleAppServiceAuditLogs(field string, value any, attrs, attrsProps map[string]any) {
	_ = "STUB: not implemented"
	return
}

func handleAppServiceAuthenticationLogs(field string, value any, attrs, attrsProps map[string]any) {
	_ = "STUB: not implemented"
	return
}

func handleAppServiceConsoleLogs(field string, value any, attrs, attrsProps map[string]any) {
	_ = "STUB: not implemented"
	return
}

func handleAppServiceHTTPLogs(field string, value any, attrs, attrsProps map[string]any) {
	_ = "STUB: not implemented"
	return
}

func handleAppServiceIPSecAuditLogs(field string, value any, attrs, attrsProps map[string]any) {
	_ = "STUB: not implemented"
	return
}

func handleAppServicePlatformLogs(field string, value any, attrs, attrsProps map[string]any) {
	_ = "STUB: not implemented"
	return
}

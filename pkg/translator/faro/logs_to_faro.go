// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package faro // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/faro"

import (
	"context"
	"regexp"
	"time"

	faroTypes "github.com/grafana/faro/pkg/go"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
)

const (
	faroKind          = "kind"
	faroTimestamp     = "timestamp"
	faroContextPrefix = "context_"

	faroSDKName         = "sdk_name"
	faroSDKVersion      = "sdk_version"
	faroSDKIntegrations = "sdk_integrations"

	faroApp               = "app"
	faroAppName           = "app_name"
	faroAppNamespace      = "app_namespace"
	faroAppRelease        = "app_release"
	faroAppVersion        = "app_version"
	faroAppBundleID       = "app_bundle_id"
	faroAppEnvironment    = "app_environment"
	faroAppInstallationID = "app_installation_id"

	faroBrowserName           = "browser_name"
	faroBrowserVersion        = "browser_version"
	faroBrowserOS             = "browser_os"
	faroBrowserMobile         = "browser_mobile"
	faroBrowserLanguage       = "browser_language"
	faroBrowserUserAgent      = "browser_userAgent"
	faroBrowserViewportHeight = "browser_viewportHeight"
	faroBrowserViewportWidth  = "browser_viewportWidth"
	faroBrowserBrands         = "browser_brands"
	faroBrowserBrandPrefix    = "browser_brand_"

	faroBrand        = "brand"
	faroBrandVersion = "version"

	faroDeviceManufacturer    = "device_manufacturer"
	faroDeviceModelIdentifier = "device_model_identifier"
	faroDeviceModelName       = "device_model_name"
	faroDeviceBrand           = "device_brand"
	faroDeviceIsPhysical      = "device_is_physical"
	faroDeviceType            = "device_type"

	faroOSName    = "os_name"
	faroOSVersion = "os_version"
	faroOSBuildID = "os_build_id"
	faroOSDetail  = "os_detail"

	faroGeoContinentIso   = "geo_continent_iso"
	faroGeoCountryIso     = "geo_country_iso"
	faroGeoSubdivisionIso = "geo_subdivision_iso"
	faroGeoCity           = "geo_city"
	faroGeoASNOrg         = "geo_asn_org"
	faroGeoASNID          = "geo_asn_id"

	faroIsK6Browser = "k6_isK6Browser"
	faroK6TestRunID = "k6_testRunId"

	faroPageID         = "page_id"
	faroPageURL        = "page_url"
	faroPageAttrPrefix = "page_attr_"

	faroSessionID         = "session_id"
	faroSessionAttrPrefix = "session_attr_"

	faroUserID         = "user_id"
	faroUserEmail      = "user_email"
	faroUsername       = "user_username"
	faroUserAttrPrefix = "user_attr_"

	faroViewName = "view_name"

	faroLogMessage = "message"
	faroLogLevel   = "level"

	faroTraceID = "traceID"
	faroSpanID  = "spanID"

	faroEventDomain     = "event_domain"
	faroEventName       = "event_name"
	faroEventDataPrefix = "event_data_"

	faroExceptionType       = "type"
	faroExceptionValue      = "value"
	faroExceptionHash       = "hash"
	faroExceptionStacktrace = "stacktrace"
	faroExceptionFatal      = "fatal"

	faroStacktraceFunction = "function"
	faroStackTraceModule   = "module"
	faroStackTraceFilename = "filename"
	faroStackTraceLineno   = "lineno"
	faroStackTraceColno    = "colno"

	faroMeasurementType        = "type"
	faroMeasurementValuePrefix = "value_"

	faroActionID       = "action_id"
	faroActionName     = "action_name"
	faroActionParentID = "action_parent_id"
)

var stacktraceRegexp *regexp.Regexp

func init() {
	stacktraceRegexp = regexp.MustCompile(`(?P<function>.+)?\s\(((?P<module>.+)\|)?(?P<filename>.+)?:(?P<lineno>\d+)?:(?P<colno>\d+)?\)$`)
}

// TranslateFromLogs converts a Logs pipeline data into []*faro.Payload
func TranslateFromLogs(ctx context.Context, ld plog.Logs) ([]faroTypes.Payload, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Reset hash before encoding

// if payload meta already exists in the metaMap merge payload to the existing payload

func mergePayloads(target *faroTypes.Payload, source faroTypes.Payload) {
	_ = "STUB: not implemented"
	// merge logs
	return
}

// merge events

// merge measurements

// merge exceptions

// merge traces

func translateLogToFaroPayload(lr plog.LogRecord, rl pcommon.Resource) (faroTypes.Payload, error) {
	_ = "STUB: not implemented"
	return *new(faroTypes.Payload), nil
}

func parseLogfmtLine(line string) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func convertLogKeyValToPayload(kv map[string]string) (faroTypes.Payload, error) {
	_ = "STUB: not implemented"
	return *new(faroTypes.Payload), nil
}

func convertEventKeyValsToPayload(kv map[string]string) (faroTypes.Payload, error) {
	_ = "STUB: not implemented"
	return *new(faroTypes.Payload), nil
}

func convertExceptionKeyValsToPayload(kv map[string]string) (faroTypes.Payload, error) {
	_ = "STUB: not implemented"
	return *new(faroTypes.Payload), nil
}

func convertMeasurementKeyValsToPayload(kv map[string]string) (faroTypes.Payload, error) {
	_ = "STUB: not implemented"
	return *new(faroTypes.Payload), nil
}

func extractMetaFromKeyVal(kv map[string]string, rl pcommon.Resource) (faroTypes.Meta, error) {
	_ = "STUB: not implemented"
	return *new(faroTypes.Meta), nil
}

func extractTimestampFromKeyVal(kv map[string]string) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

func extractSDKFromKeyVal(kv map[string]string) faroTypes.SDK {
	_ = "STUB: not implemented"
	return *new(faroTypes.SDK)
}

func parseIntegrationsFromString(integrationsString string) []faroTypes.SDKIntegration {
	_ = "STUB: not implemented"
	return nil
}

func extractAppFromKeyVal(kv map[string]string, rl pcommon.Resource) faroTypes.App {
	_ = "STUB: not implemented"
	return *new(faroTypes.App)
}

// force the app name stored in resource attribute service.name
// if service.name resource attribute is missing try to get app name from the custom "app" resource attribute

// force the app name stored in resource attribute service.name or in custom "app" resource attribute
// if service.name resource attribute is missing as well as custom "app" attribute try to get app name from the log line

// force the app namespace stored in resource attribute service.namespace
// if service.namespace resource attribute is missing try to get app namespace from the log line

// force the app version stored in resource attribute service.version
// if service.version resource attribute is missing try to get app version from the log line

// force the app environment stored in resource attribute deployment.environment
// if deployment.environment resource attribute is missing try to get app environment from the log line

func extractBrowserFromKeyVal(kv map[string]string) (*faroTypes.Browser, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func extractBrowserBrandsFromKeyVal(kv map[string]string) (faroTypes.Browser_Brands, error) {
	_ = "STUB: not implemented"
	return *new(faroTypes.Browser_Brands), nil
}

func extractDeviceFromKeyVal(kv map[string]string) (faroTypes.Device, error) {
	_ = "STUB: not implemented"
	return *new(faroTypes.Device), nil
}

func extractOSFromKeyVal(kv map[string]string) faroTypes.OS {
	_ = "STUB: not implemented"
	return *new(faroTypes.OS)
}

func extractGeoFromKeyVal(kv map[string]string) faroTypes.Geo {
	_ = "STUB: not implemented"
	return *new(faroTypes.Geo)
}

func extractK6FromKeyVal(kv map[string]string) (faroTypes.K6, error) {
	_ = "STUB: not implemented"
	return *new(faroTypes.K6), nil
}

func extractPageFromKeyVal(kv map[string]string) faroTypes.Page {
	_ = "STUB: not implemented"
	return *new(faroTypes.Page)
}

func extractSessionFromKeyVal(kv map[string]string) faroTypes.Session {
	_ = "STUB: not implemented"
	return *new(faroTypes.Session)
}

func extractUserFromKeyVal(kv map[string]string) faroTypes.User {
	_ = "STUB: not implemented"
	return *new(faroTypes.User)
}

func extractViewFromKeyVal(kv map[string]string) faroTypes.View {
	_ = "STUB: not implemented"
	return *new(faroTypes.View)
}

func extractLogFromKeyVal(kv map[string]string) (faroTypes.Log, error) {
	_ = "STUB: not implemented"
	return *new(faroTypes.Log), nil
}

func extractLogContextFromKeyVal(kv map[string]string) faroTypes.LogContext {
	_ = "STUB: not implemented"
	return *new(faroTypes.LogContext)
}

func extractTraceFromKeyVal(kv map[string]string) faroTypes.TraceContext {
	_ = "STUB: not implemented"
	return *new(faroTypes.TraceContext)
}

func extractActionFromKeyVal(kv map[string]string) faroTypes.Action {
	_ = "STUB: not implemented"
	return *new(faroTypes.Action)
}

func extractEventFromKeyVal(kv map[string]string) (faroTypes.Event, error) {
	_ = "STUB: not implemented"
	return *new(faroTypes.Event), nil
}

func extractExceptionFromKeyVal(kv map[string]string) (faroTypes.Exception, error) {
	_ = "STUB: not implemented"
	return *new(faroTypes.Exception), nil
}

func extractExceptionContextFromKeyVal(kv map[string]string) faroTypes.ExceptionContext {
	_ = "STUB: not implemented"
	return *new(faroTypes.ExceptionContext)
}

func extractStacktraceFromKeyVal(kv map[string]string, exceptionType, exceptionValue string) (*faroTypes.Stacktrace, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseStacktraceFromString(stacktraceStr, exceptionType, exceptionValue string) (*faroTypes.Stacktrace, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseFrameFromString(frameStr string) (*faroTypes.Frame, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func extractMeasurementFromKeyVal(kv map[string]string) (faroTypes.Measurement, error) {
	_ = "STUB: not implemented"
	return *new(faroTypes.Measurement), nil
}

func extractMeasurementContextFromKeyVal(kv map[string]string) faroTypes.MeasurementContext {
	_ = "STUB: not implemented"
	return *new(faroTypes.MeasurementContext)
}

func extractMeasurementValuesFromKeyVal(kv map[string]string) (map[string]float64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func extractAttributesWithPrefixFromKeyVal(prefix string, kv map[string]string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package logs // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding/azureencodingextension/internal/unmarshaler/logs"

import (
	"encoding/json"

	"go.opentelemetry.io/collector/pdata/pcommon"
)

// Non-SemConv attributes that are used for common Azure Log Record fields
const (
	// OpenTelemetry attribute name for the name of the task being performed
	attributeAzureTaskName = "azure.app_service.task.name"

	// OpenTelemetry attribute name for original log level from Azure Log Record
	// as it sent by App Service
	attributeAzureOriginalLogLevel = "log.record.severity.original"

	// OpenTelemetry attribute name for Web Instance Id the application running
	attributeAzureWebInstanceID = "azure.app_service.instance.id"

	// OpenTelemetry attribute name for the authentication event details
	attributeAzureAuthEventDetails = "azure.auth.event.details"

	// OpenTelemetry attribute name for the version of App Service Authentication running
	attributeAzureModuleRuntimeVersion = "azure.auth.module.runtime.version"

	// OpenTelemetry attribute name for the runtime name of the application
	attributeAzureSiteName = "azure.app_service.site.name"

	// OpenTelemetry attribute name for HTTP sub-status code of the request,
	// in Azure App Service logs it differs from standard HTTP sub-status code semantic convention
	attributeAzureSubStatusCode = "azure.http.response.sub_status_code"

	// OpenTelemetry attribute name for the duration of the indicator of the access via
	// Virtual Network Service Endpoint communication
	attributeAzureIsServiceEndpoint = "azure.app_service.endpoint"

	// OpenTelemetry attribute name for the Deployment ID of the application deployment
	attributeAzureDeploymentID = "azure.deployment.id"

	// OpenTelemetry attribute name for Logger name
	attributeLogLogger = "log.record.logger"

	// OpenTelemetry attribute name for for "x-azure-fdid" HTTP Header value
	attributeHTTPHeaderAzureFDID = "http.request.header.x-azure-fdid"

	// OpenTelemetry attribute name for for "x-fd-healthprobe" HTTP Header value
	attributeHTTPHeaderFDHealthProbe = "http.request.header.x-fd-healthprobe"

	// OpenTelemetry attribute name for for "x-forwarded-for" HTTP Header value
	attributeHTTPHeaderForwardedFor = "http.request.header.x-forwarded-for"

	// OpenTelemetry attribute name for for "x-forwarded-host" HTTP Header value
	attributeHTTPHeaderForwardedHost = "http.request.header.x-forwarded-host"

	// OpenTelemetry attribute name for "Referer" HTTP Header value
	attributeHTTPHeaderReferer = "http.request.header.referer"
)

type appServiceAppLogProperties struct {
	ContainerID       string `json:"containerId"`
	CustomLevel       string `json:"customLevel"`
	ExceptionClass    string `json:"exceptionClass"`
	Host              string `json:"host"`
	Logger            string `json:"logger"`
	Message           string `json:"message"`
	Method            string `json:"method"`
	Source            string `json:"source"`
	StackTrace        string `json:"stackTrace"`
	WebSiteInstanceID string `json:"webSiteInstanceId"`
}

func (p *appServiceAppLogProperties) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// See https://learn.microsoft.com/en-us/azure/azure-monitor/reference/tables/appserviceapplogs
// There is no documentation about the structure of the logs, so we will
// implement this based on the fields from Azure Monitor table excluding generic fields
type azureAppServiceAppLog struct {
	azureLogRecordBase

	Properties appServiceAppLogProperties `json:"properties"`
}

func (r *azureAppServiceAppLog) PutProperties(attrs pcommon.Map, body pcommon.Value) error {
	_ = "STUB: not implemented"
	return nil
}

type appServiceAuditLogProperties struct {
	User            string `json:"user"`
	UserDisplayName string `json:"userDisplayName"`
	UserAddress     string `json:"userAddress"`
	Protocol        string `json:"protocol"`
}

func (p *appServiceAuditLogProperties) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// See https://github.com/MicrosoftDocs/azure-docs/blob/main/articles/app-service/configure-basic-auth-disable.md#monitor-for-basic-authentication-attempts
type azureAppServiceAuditLog struct {
	azureLogRecordBase

	Properties appServiceAuditLogProperties `json:"properties"`
}

func (r *azureAppServiceAuditLog) PutProperties(attrs pcommon.Map, _ pcommon.Value) error {
	_ = "STUB: not implemented"
	return nil
}

type appServiceAuthenticationLogProperties struct {
	Details              string      `json:"details"`
	Host                 string      `json:"hostName"`
	Message              string      `json:"message"`
	ModuleRuntimeVersion string      `json:"moduleRuntimeVersion"`
	SiteName             string      `json:"siteName"`
	StatusCode           json.Number `json:"statusCode"`    // int
	SubStatusCode        json.Number `json:"subStatusCode"` // int
	TaskName             string      `json:"taskName"`
}

func (p *appServiceAuthenticationLogProperties) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// See https://learn.microsoft.com/en-us/azure/azure-monitor/reference/tables/appserviceauthenticationlogs
// There is no documentation about the structure of the logs, so we will
// implement this based on the fields from Azure Monitor table excluding generic fields
type azureAppServiceAuthenticationLog struct {
	azureLogRecordBase

	Properties appServiceAuthenticationLogProperties `json:"properties"`
}

func (r *azureAppServiceAuthenticationLog) PutProperties(attrs pcommon.Map, body pcommon.Value) error {
	_ = "STUB: not implemented"
	return nil
}

type appServiceConsoleLogProperties struct {
	ContainerID string `json:"containerId"`
	Host        string `json:"host"`
}

func (p *appServiceConsoleLogProperties) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// See https://learn.microsoft.com/en-us/azure/azure-monitor/reference/tables/appserviceconsolelogs
// There is no documentation about the structure of the logs, so we will
// implement this based on the fields from Azure Monitor table excluding generic fields
type azureAppServiceConsoleLog struct {
	azureLogRecordBase

	Properties appServiceConsoleLogProperties `json:"properties"`
}

func (r *azureAppServiceConsoleLog) PutProperties(attrs pcommon.Map, _ pcommon.Value) error {
	_ = "STUB: not implemented"
	return nil
}

// appServiceHTTPLogProperties represents the properties field of AppServiceHTTPLogs.
// Azure can emit this field as either a JSON object or a stringified JSON string
// (common when logs are routed via Azure Event Hub).
type appServiceHTTPLogProperties struct {
	ClientIP       string      `json:"CIp"`
	Host           string      `json:"ComputerName"`
	Cookie         string      `json:"Cookie"`
	RequestBytes   json.Number `json:"CsBytes"` // int
	HostHeader     string      `json:"CsHost"`
	RequestMethod  string      `json:"CsMethod"`
	URIQuery       string      `json:"CsUriQuery"`
	RequestPath    string      `json:"CsUriStem"`
	UserName       string      `json:"CsUsername"`
	Referer        string      `json:"Referer"`
	Result         string      `json:"Result"`
	ResponseBytes  json.Number `json:"ScBytes"`  // int
	HTTPStatusCode json.Number `json:"ScStatus"` // int
	HTTPSubStatus  string      `json:"ScSubStatus"`
	ServerPort     json.Number `json:"SPort"`     // int
	TimeTaken      json.Number `json:"TimeTaken"` // int
	UserAgent      string      `json:"UserAgent"`
}

func (p *appServiceHTTPLogProperties) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	// Define an alias type to avoid infinite recursion
	return nil
}

// See https://learn.microsoft.com/en-us/azure/azure-monitor/reference/tables/appservicehttplogs
// There is no documentation about the structure of the logs, so we will
// implement this based on the fields from Azure Monitor table excluding generic fields
type azureAppServiceHTTPLog struct {
	azureLogRecordBase

	Properties appServiceHTTPLogProperties `json:"properties"`
}

func (r *azureAppServiceHTTPLog) PutProperties(attrs pcommon.Map, body pcommon.Value) error {
	_ = "STUB: not implemented"
	// In general it's unsafe to put Cookie values in Log Attributes as it may contain sensitive information,
	// and there is no generic masking available for it as well
	// So we will skip "Cookie" field here
	return nil
}

type appServiceIPSecAuditLogProperties struct {
	ClientIP          string `json:"CIp"`
	HostHeader        string `json:"CsHost"`
	Details           string `json:"details"`
	Result            string `json:"Result"`
	IsServiceEndpoint string `json:"ServiceEndpoint"`
	XAzureFDID        string `json:"XAzureFDID"`     // X-Azure-FDID header
	XFDHealthProbe    string `json:"XFDHealthProbe"` // X-FD-HealthProbe header
	XForwardedFor     string `json:"XForwardedFor"`  // X-Forwarded-For header
	XForwardedHost    string `json:"XForwardedHost"` // X-Forwarded-Host header
}

func (p *appServiceIPSecAuditLogProperties) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// See https://learn.microsoft.com/en-us/azure/azure-monitor/reference/tables/appserviceipsecauditlogs
// There is no documentation about the structure of the logs, so we will
// implement this based on the fields from Azure Monitor table excluding generic fields
type azureAppServiceIPSecAuditLog struct {
	azureLogRecordBase

	Properties appServiceIPSecAuditLogProperties `json:"properties"`
}

func (r *azureAppServiceIPSecAuditLog) PutProperties(attrs pcommon.Map, body pcommon.Value) error {
	_ = "STUB: not implemented"
	return nil
}

type appServicePlatformLogProperties struct {
	ContainerID  string `json:"containerId"`
	DeploymentID string `json:"deploymentId"`
	Exception    string `json:"Exception"`
	Host         string `json:"host"`
	Message      string `json:"message"`
	StackTrace   string `json:"stackTrace"`
}

func (p *appServicePlatformLogProperties) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// See https://learn.microsoft.com/en-us/azure/azure-monitor/reference/tables/appserviceplatformlogs
// There is no documentation about the structure of the logs, so we will
// implement this based on the fields from Azure Monitor table excluding generic fields
type azureAppServicePlatformLog struct {
	azureLogRecordBase

	Properties appServicePlatformLogProperties `json:"properties"`
}

func (r *azureAppServicePlatformLog) PutProperties(attrs pcommon.Map, body pcommon.Value) error {
	_ = "STUB: not implemented"
	return nil
}

type appServiceFileAuditLogProperties struct {
	Path    string `json:"path"`
	Process string `json:"process"`
}

func (p *appServiceFileAuditLogProperties) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// See https://learn.microsoft.com/en-us/azure/azure-monitor/reference/tables/appservicefileauditlogs
// There is no documentation about the structure of the logs, so we will
// implement this based on the fields from Azure Monitor table excluding generic fields
type azureAppServiceFileAuditLog struct {
	azureLogRecordBase

	Properties appServiceFileAuditLogProperties `json:"properties"`
}

func (r *azureAppServiceFileAuditLog) PutProperties(attrs pcommon.Map, _ pcommon.Value) error {
	_ = "STUB: not implemented"
	return nil
}

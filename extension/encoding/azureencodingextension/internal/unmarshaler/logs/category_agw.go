// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package logs // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding/azureencodingextension/internal/unmarshaler/logs"

import (
	"encoding/json"

	"go.opentelemetry.io/collector/pdata/pcommon"
)

const (
	// OpenTelemetry attribute name for the whether communication to the backend pools used TLS ("on" or "off")
	attributeTLSEnabled = "tls.enabled"

	// OpenTelemetry attribute name for the name of the listener associated with given request log
	attributeAzureAGWListenerName = "azure.agw.listener.name"

	// OpenTelemetry attribute name for the name of the rule associated with given request log
	attributeAzureAGWRuleName = "azure.agw.rule.name"

	// OpenTelemetry attribute name for the name of the backend pool associated with the given request log.
	attributeAzureAGWBackendPoolName = "azure.agw.backend.pool.name"

	// OpenTelemetry attribute name for the name of the backend setting associated with the given request log.
	attributeAzureAGWBackendSettingName = "azure.agw.backend.setting.name"

	// OpenTelemetry attribute name for the HTTP status code returned by the backend server
	attributeAzureAGWBackendStatusCode = "azure.agw.backend.status_code"

	// OpenTelemetry attribute name for the count of healthy hosts in the back-end pool
	attributeAzureAGWHostHealthyCount = "azure.agw.backend.healthy.count"

	// OpenTelemetry attribute name for the count of unhealthy hosts in the back-end pool
	attributeAzureAGWHostUnhealthyCount = "azure.agw.backend.unhealthy.count"

	// OpenTelemetry attribute name for the total number of failed requests served by the Application Gateway
	attributeAzureAGWRequestCount = "azure.agw.request.count"

	// OpenTelemetry attribute name for the total number of failed requests
	attributeAzureAGWFailedRequestCount = "azure.agw.request.failed"

	// OpenTelemetry attribute name for the Average latency (in milliseconds) of requests
	// from the AGW instance to the back-end that serves the requests
	attributeAzureAGWBackendLatency = "azure.agw.backend.latency"

	// OpenTelemetry attribute name for the time difference (in seconds) between first byte received from the backend
	// to first byte sent to the client
	attributeAzureAGWLatency = "azure.agw.latency"

	// OpenTelemetry attribute name for the average throughput since the last log,
	// measured in bytes per second
	attributeAzureAGWThroughput = "azure.agw.throughput"

	// OpenTelemetry attribute name for the length of time (in seconds)
	// that it takes for the request to be processed by the WAF
	attributeAzureFirewallLatency = "azure.firewall.evaluation.duration"

	// OpenTelemetry attribute name for the site to which the firewall log was generated,
	// currently, only "Global" is listed because rules are global
	attributeAzureFirewallSite = "azure.firewall.site"

	// OpenTelemetry attribute name for detailed information about the firewall triggering event
	attributeAzureFirewallEventDetails = "azure.firewall.evaluation.details"

	// OpenTelemetry attribute name for the unique ID of the Firewall Policy
	// associated with the Application Gateway, Listener, or Path
	attributeAzureFirewallPolicyID = "azure.firewall.policy.id"

	// OpenTelemetry attribute name for the location of the policy ("Global", "Listener", or "Location")
	attributeAzureFirewallPolicyScope = "azure.firewall.policy.scope.type"

	// OpenTelemetry attribute name for the name of the object where the policy is applied
	attributeAzureFirewallPolicyScopeName = "azure.firewall.policy.object.name"
)

// See https://github.com/MicrosoftDocs/azure-docs/blob/main/articles/web-application-firewall/ag/web-application-firewall-logs.md
type azureApplicationGatewayAccessLog struct {
	azureLogRecordBase

	// Custom top-level attributes
	ListenerName       string `json:"listenerName"`
	RuleName           string `json:"ruleName"`
	BackendPoolName    string `json:"backendPoolName"`
	BackendSettingName string `json:"backendSettingName"`

	Properties struct {
		InstanceID     string      `json:"instanceId"`
		ClientIP       string      `json:"clientIP"`
		ClientPort     json.Number `json:"clientPort"` // int
		HTTPMethod     string      `json:"httpMethod"`
		RequestURI     string      `json:"requestUri"`
		RequestQuery   string      `json:"requestQuery"`
		UserAgent      string      `json:"userAgent"`
		HTTPStatusCode json.Number `json:"httpStatus"` // int
		HTTPVersion    string      `json:"httpVersion"`
		ReceivedBytes  json.Number `json:"receivedBytes"` // int
		SentBytes      json.Number `json:"sentBytes"`     // int
		TimeTaken      json.Number `json:"timeTaken"`     // float, s
		SSLEnabled     string      `json:"sslEnabled"`
		Host           string      `json:"host"`
		OriginalHost   string      `json:"originalHost"`
		// v2 fields
		ClientResponseTime         json.Number `json:"clientResponseTime"` // float
		OriginalRequestURIWithArgs string      `json:"originalRequestUriWithArgs"`
		SSLCipher                  string      `json:"sslCipher"`
		SSLProtocol                string      `json:"sslProtocol"`
		ServerRouted               string      `json:"serverRouted"`
		ServerStatus               string      `json:"serverStatus"`          // int + stringToNumber
		ServerResponseLatency      string      `json:"serverResponseLatency"` // float + stringToNumber
		TransactionID              string      `json:"transactionId"`
		WAFEvaluationTime          string      `json:"WAFEvaluationTime"` // float + stringToNumber
		WAFMode                    string      `json:"WAFMode"`
		UpstreamSourcePort         string      `json:"upstreamSourcePort"` // int + stringToNumber
		ErrorInfo                  string      `json:"error_info"`
	} `json:"properties"`
}

// Override GetResource to add ServiceInstanceID from Properties
func (r *azureApplicationGatewayAccessLog) GetResource() logsResourceAttributes {
	_ = "STUB: not implemented"
	return *new(logsResourceAttributes)
}

func (r *azureApplicationGatewayAccessLog) PutCommonAttributes(attrs pcommon.Map, body pcommon.Value) {
	_ = "STUB: not implemented"
	// Put common attributes first
	return
}

// Then put custom top-level attributes

func (r *azureApplicationGatewayAccessLog) PutProperties(attrs pcommon.Map, _ pcommon.Value) error {
	_ = "STUB: not implemented"
	return nil
}

// requestUri is not an absolute URL, so we cannot use unmarshaler.AttrPutURLParsed here

// See https://github.com/MicrosoftDocs/azure-docs/blob/main/articles/web-application-firewall/ag/web-application-firewall-logs.md
type azureApplicationGatewayPerformanceLog struct {
	azureLogRecordBase

	Properties struct {
		InstanceID         string      `json:"instanceId"`
		HealthyHostCount   json.Number `json:"healthyHostCount"`   // int
		UnHealthyHostCount json.Number `json:"unHealthyHostCount"` // int
		RequestCount       json.Number `json:"requestCount"`       // int
		Latency            json.Number `json:"latency"`            // int, ms
		FailedRequestCount json.Number `json:"failedRequestCount"` // int
		Throughput         json.Number `json:"throughput"`         // int, bytes
	} `json:"properties"`
}

// Override GetResource to add ServiceInstanceID from Properties
func (r *azureApplicationGatewayPerformanceLog) GetResource() logsResourceAttributes {
	_ = "STUB: not implemented"
	return *new(logsResourceAttributes)
}

// addApplicationGatewayAccessLogsProperties parses the Azure Resource Log record and adds
// the relevant attributes to the OpenTelemetry Log Record Attributes and/or Log Body
func (r *azureApplicationGatewayPerformanceLog) PutProperties(attrs pcommon.Map, _ pcommon.Value) error {
	_ = "STUB: not implemented"
	return nil
}

// See https://github.com/MicrosoftDocs/azure-docs/blob/main/articles/web-application-firewall/ag/web-application-firewall-logs.md
type azureApplicationGatewayFirewallLog struct {
	azureLogRecordBase

	Properties struct {
		InstanceID      string         `json:"instanceId"`
		ClientIP        string         `json:"clientIp"`
		ClientPort      json.Number    `json:"clientPort"` // int
		RequestURI      string         `json:"requestUri"`
		RuleSetType     string         `json:"ruleSetType"`
		RuleSetVersion  string         `json:"ruleSetVersion"`
		RuleID          string         `json:"ruleId"`
		RuleGroup       string         `json:"ruleGroup"`
		Message         string         `json:"message"`
		Action          string         `json:"action"`
		Site            string         `json:"site"`
		Details         map[string]any `json:"details"`
		HostName        string         `json:"hostname"`
		TransactionID   string         `json:"transactionId"`
		PolicyID        string         `json:"policyId"`
		PolicyScope     string         `json:"policyScope"`
		PolicyScopeName string         `json:"policyScopeName"`
	} `json:"properties"`
}

// Override GetResource to add ServiceInstanceID from Properties
func (r *azureApplicationGatewayFirewallLog) GetResource() logsResourceAttributes {
	_ = "STUB: not implemented"
	return *new(logsResourceAttributes)
}

func (r *azureApplicationGatewayFirewallLog) PutProperties(attrs pcommon.Map, body pcommon.Value) error {
	_ = "STUB: not implemented"
	return nil
}

// requestUri is not an absolute URL, so we cannot use unmarshaler.AttrPutURLParsed here

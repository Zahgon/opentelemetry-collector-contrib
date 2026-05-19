// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package logs // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding/azureencodingextension/internal/unmarshaler/logs"

import (
	"encoding/json"

	"go.opentelemetry.io/collector/pdata/pcommon"
)

const (
	// OpenTelemetry attribute name for the unique ID to identify the health probe request
	attributeAzureFrontDoorHealthProbeID = "azure.frontdoor.health_probe.id"

	// OpenTelemetry attribute name for the unique ID to identify the health probe request
	attributeAzureFrontDoorHealthOriginName = "azure.frontdoor.health_probe.origin.name"

	// OpenTelemetry attribute name for the time from when the Azure Front Door edge sent
	// the health probe request to the origin to when the origin sent the last response to Azure Front Door.
	attributeAzureFrontDoorHealthTotalLatency = "azure.frontdoor.health_probe.origin.latency.total"

	// OpenTelemetry attribute name for the time spent setting up the TCP connection
	// to send the HTTP probe request to the origin
	attributeAzureFrontDoorHealthConnLatency = "azure.frontdoor.health_probe.origin.latency.connection"

	// OpenTelemetry attribute name for the time spent on DNS resolution
	attributeAzureFrontDoorHealthDNSLatency = "azure.frontdoor.health_probe.origin.latency.dns"
)

// NOTE: For "FrontDoorAccessLog" category - see "azureHTTPAccessLog" struct in category_azurecdn.go

// See https://github.com/MicrosoftDocs/azure-docs/blob/main/articles/frontdoor/monitor-front-door.md#health-probe-log
type frontDoorHealthProbeLog struct {
	azureLogRecordBase

	Properties struct {
		HealthProbeID     string      `json:"healthProbeId"`
		Pop               string      `json:"POP"`
		HTTPVerb          string      `json:"httpVerb"`
		Result            string      `json:"result"`
		HTTPStatusCode    json.Number `json:"httpStatusCode"` // int
		ProbeURL          string      `json:"probeURL"`
		OriginName        string      `json:"originName"`
		OriginIP          string      `json:"originIP"`
		TotalLatency      json.Number `json:"totalLatencyMilliseconds"`      // int, ms
		ConnectionLatency json.Number `json:"connectionLatencyMilliseconds"` // int, ms
		DNSLatency        json.Number `json:"DNSLatencyMicroseconds"`        // int, us
	} `json:"properties"`
}

func (r *frontDoorHealthProbeLog) PutProperties(attrs pcommon.Map, body pcommon.Value) error {
	_ = "STUB: not implemented"
	return nil
}

// See https://learn.microsoft.com/en-us/azure/web-application-firewall/afds/waf-front-door-monitor?pivots=front-door-standard-premium#waf-logs
type frontDoorWAFLog struct {
	azureLogRecordBase

	Properties struct {
		ClientIP          string      `json:"clientIP"`
		ClientPort        json.Number `json:"clientPort"` // int
		SocketIP          string      `json:"socketIP"`
		RequestURI        string      `json:"requestUri"`
		RuleName          string      `json:"ruleName"`
		Policy            string      `json:"policy"`
		Action            string      `json:"action"`
		Host              string      `json:"host"`
		TrackingReference string      `json:"trackingReference"`
		PolicyMode        string      `json:"policyMode"`
	} `json:"properties"`
}

func (r *frontDoorWAFLog) PutProperties(attrs pcommon.Map, _ pcommon.Value) error {
	_ = "STUB: not implemented"
	return nil
}

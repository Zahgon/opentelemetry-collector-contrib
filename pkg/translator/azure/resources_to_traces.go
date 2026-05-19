// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package azure // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/azure"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"
)

const (
	// Constants for OpenTelemetry Specs
	traceAzureResourceID = "azure.resource.id"
)

type azureTracesRecords struct {
	Records []azureTracesRecord `json:"records"`
}

// Azure Trace Records based on Azure AppRequests & AppDependencies table data
// the common record schema reference:
// https://learn.microsoft.com/en-us/azure/azure-monitor/reference/tables/apprequests
// https://learn.microsoft.com/en-us/azure/azure-monitor/reference/tables/appdependencies
type azureTracesRecord struct {
	Time                  string             `json:"time"`
	ResourceID            string             `json:"resourceId"`
	ResourceGUID          string             `json:"ResourceGUID"`
	Type                  string             `json:"Type"`
	AppRoleInstance       string             `json:"AppRoleInstance"`
	AppRoleName           string             `json:"AppRoleName"`
	AppVersion            string             `json:"AppVersion"`
	ClientCity            string             `json:"ClientCity"`
	ClientCountryOrRegion string             `json:"ClientCountryOrRegion"`
	ClientIP              string             `json:"ClientIP"`
	ClientStateOrProvince string             `json:"ClientStateOrProvince"`
	ClientType            string             `json:"ClientType"`
	IKey                  string             `json:"IKey"`
	OperationName         string             `json:"OperationName"`
	OperationID           string             `json:"OperationId"`
	ParentID              string             `json:"ParentId"`
	SDKVersion            string             `json:"SDKVersion"`
	Properties            map[string]string  `json:"Properties"`
	Measurements          map[string]float64 `json:"Measurements"`
	SpanID                string             `json:"Id"`
	Name                  string             `json:"Name"`
	URL                   string             `json:"Url"`
	Source                string             `json:"Source"`
	Success               bool               `json:"Success"`
	ResultCode            string             `json:"ResultCode"`
	DurationMs            float64            `json:"DurationMs"`
	PerformanceBucket     string             `json:"PerformanceBucket"`
	ItemCount             float64            `json:"ItemCount"`
}

var _ ptrace.Unmarshaler = (*TracesUnmarshaler)(nil)

type TracesUnmarshaler struct {
	Version     string
	Logger      *zap.Logger
	TimeFormats []string
}

func (r TracesUnmarshaler) UnmarshalTraces(buf []byte) (ptrace.Traces, error) {
	_ = "STUB: not implemented"
	return *new(ptrace.Traces), nil
}

// HTTP Method is already mapped to http.method

func TraceIDFromHex(hexStr string) (pcommon.TraceID, error) {
	_ = "STUB: not implemented"
	return *new(pcommon.TraceID), nil
}

func SpanIDFromHex(hexStr string) (pcommon.SpanID, error) {
	_ = "STUB: not implemented"
	return *new(pcommon.SpanID), nil
}

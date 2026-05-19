// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Find more information about Cloud dns logs at:
// https://docs.cloud.google.com/dns/docs/monitoring#dns-log-record-format
package dnslog // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding/googlecloudlogentryencodingextension/internal/dnslog"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
)

const (
	CloudDNSQueryLogSuffix = "dns.googleapis.com%2Fdns_queries"

	// Query related attributes. Ref: https://datatracker.ietf.org/doc/html/rfc1035#section-4.1.2
	// gcpDNSQueryType holds the DNS query type
	gcpDNSQueryType = "dns.question.type" // TBD in SemConv

	// Response related attributes
	// gcpDNSResponseCode holds the DNS response code
	gcpDNSResponseCode = "dns.response_code" // TBD in SemConv
	// gcpDNSAliasQueryResponseCode holds the response code for alias queries
	gcpDNSAliasQueryResponseCode = "gcp.dns.alias_query.response.code"
	// gcpDNSAuthAnswer indicates whether the response is authoritative
	gcpDNSAuthAnswer = "gcp.dns.auth_answer" // Ref: https://datatracker.ietf.org/doc/html/rfc1035
	// gcpDNSAnswerData holds DNS answer in presentation format
	gcpDNSAnswerData = "dns.answer.data" // TBD in SemConv

	// Network related attributes
	// gcpDNSClientVPCNetwork holds the client network name where the DNS query originated
	gcpDNSClientVPCNetwork = "gcp.dns.client.vpc.name"
	// gcpDNSClientType holds the client type of the DNS query
	gcpDNSClientType = "gcp.dns.client.type"

	// Server related attributes
	// gcpDNSServerName holds the name name of Google Cloud instance that is responsible of resolving the query
	gcpDNSServerName = "gcp.dns.server.name"
	// gcpDNServerType holds the type of target resolving the DNS query
	gcpDNServerType = "gcp.dns.server.type"

	// Performance and error related attributes
	// gcpDNSServerLatency holds the server-side latency in seconds
	gcpDNSServerLatency = "gcp.dns.server.latency"
	// gcpDNSEgressError holds Egress proxy error, the actual error as received from the on-premises DNS server
	gcpDNSEgressError = "gcp.dns.egress.error"
	// gcpDNSHealthyIPs holds addresses in the ResourceRecordSet that are known to be HEALTHY.
	gcpDNSHealthyIPs = "gcp.dns.healthy.ips"
	// gcpDNSUnhealthyIPs holds addresses in the ResourceRecordSet that are known to be UNHEALTHY.
	gcpDNSUnhealthyIPs = "gcp.dns.unhealthy.ips"

	// DNS feature related attributes
	// gcpDNSDNS64Translated indicates whether DNS64 translation was applied
	gcpDNSDNS64Translated = "gcp.dns.dns64.translated"

	// VM instance related attributes
	// gcpProjectID holds the Google Cloud project ID of the network from which the query was sent
	gcpProjectID = "gcp.project.id"
)

type dnslog struct {
	AliasQueryResponseCode string   `json:"alias_query_response_code"`
	AuthAnswer             *bool    `json:"authAnswer"`
	DestinationIP          string   `json:"destinationIP"`
	DNS64Translated        *bool    `json:"dns64Translated"`
	EgressError            string   `json:"egressError"`
	HealthyIps             string   `json:"healthyIps"`
	Location               string   `json:"location"`
	Protocol               string   `json:"protocol"`
	ProjectID              string   `json:"project_id"`
	QueryName              string   `json:"queryName"`
	QueryType              string   `json:"queryType"`
	Rdata                  string   `json:"rdata"`
	ResponseCode           string   `json:"responseCode"`
	ServerLatency          *float64 `json:"serverLatency"`
	SourceIP               string   `json:"sourceIP"`
	SourceNetwork          string   `json:"sourceNetwork"`
	SourceType             string   `json:"source_type"`
	TargetName             string   `json:"target_name"`
	TargetType             string   `json:"target_type"`
	UnhealthyIps           string   `json:"unhealthyIps"`
	VMInstanceID           *int64   `json:"vmInstanceId"`
	VMInstanceName         string   `json:"vmInstanceName"`
	VMProjectID            string   `json:"vmProjectId"`
	VMZoneName             string   `json:"vmZoneName"`
}

func handleQueryAttributes(log *dnslog, attr pcommon.Map) { _ = "STUB: not implemented"; return }

// TBD in SemConv

func handleResponseAttributes(log *dnslog, attr pcommon.Map) { _ = "STUB: not implemented"; return }

// TBD in SemConv

// TBD in SemConv

func handleNetworkAttributes(log *dnslog, attr pcommon.Map) { _ = "STUB: not implemented"; return }

func handleTargetAttributes(log *dnslog, attr pcommon.Map) { _ = "STUB: not implemented"; return }

func handlePerformanceAndErrorAttributes(log *dnslog, attr pcommon.Map) {
	_ = "STUB: not implemented"
	return
}

func handleDNSFeatureAttributes(log *dnslog, attr pcommon.Map) { _ = "STUB: not implemented"; return }

func handleVMInstanceAttributes(log *dnslog, attr pcommon.Map) { _ = "STUB: not implemented"; return }

func ParsePayloadIntoAttributes(payload []byte, attr pcommon.Map) error {
	_ = "STUB: not implemented"
	return nil
}

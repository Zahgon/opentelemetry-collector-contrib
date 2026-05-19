// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package elbaccesslogs // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding/awslogsencodingextension/internal/unmarshaler/elb-access-log"

type logSyntaxType string

const (
	albAccessLogs  logSyntaxType = "alb_access_logs"
	nlbAccessLogs  logSyntaxType = "nlb_access_logs"
	clbAccessLogs  logSyntaxType = "clb_access_logs"
	controlMessage logSyntaxType = "control_message"

	// any field can be set to - to indicate that the data was unknown
	// or unavailable, or that the field was not applicable to this request.
	unknownField = "-"
	// First field of ELB control message
	EnableControlMessage = "Enable"
)

// CLBAccessLogRecord represents a record of access logs from an AWS Classic Load Balancer.
// Documentation Reference: https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/access-log-collection.html
type CLBAccessLogRecord struct {
	Time                   string  // Timestamp the load balancer received the request from the client, in ISO 8601 format
	ELB                    string  // The name of the load balancer
	ClientIP               string  // Client IP
	ClientPort             int64   // Client  port
	BackendIPPort          string  // Backend IP:Port or -
	BackendIP              string  // Backend IP split from BackendIPPort
	BackendPort            int64   // Backend port split from BackendIPPort
	RequestProcessingTime  float64 // Time taken to process the request in seconds (HTTP/TCP)
	BackendProcessingTime  float64 // Time taken for the registered instance to respond
	ResponseProcessingTime float64 // Time taken to send the response to the client
	ELBStatusCode          int64   // Status code from the load balancer
	BackendStatusCode      int64   // Status code from the backend instance
	ReceivedBytes          int64   // Size of the request in bytes received from the client
	SentBytes              int64   // Size of the response in bytes sent to the client
	RequestMethod          string  // HTTP method used in the request (e.g., GET, POST)
	RequestURI             string  // Full URI requested by the client
	ProtocolName           string  // Network protocol name (e.g., HTTP)
	ProtocolVersion        string  // Network protocol version (e.g., 1.1)
	UserAgent              string  // The User-Agent string identifying the client
	SSLCipher              string  // The SSL cipher used, available for HTTPS listeners
	SSLProtocol            string  // The SSL protocol negotiated, available for HTTPS listeners
}

// convertTextToCLBAccessLogRecord converts a slice of strings into a CLBAccessLogRecord
func convertTextToCLBAccessLogRecord(fields []string) (CLBAccessLogRecord, error) {
	_ = "STUB: not implemented"
	return *new(CLBAccessLogRecord), nil
}

// Map fields to the struct

// Timestamp
// Load balancer name
// Backend IP:Port or -
// Placeholder for ELB status code
// Placeholder for Backend status code
// User-Agent
// SSL cipher
// SSL protocol

// Process the fields for numerical values (convenient to parse from string)

// Parse BackendIPPort into BackendIP and BackendPort

// ELB status code and backend status code can be - in case of TCP and SSL entry

// Network Load Balancer Access Logs record
// Doc: https://docs.aws.amazon.com/elasticloadbalancing/latest//network/load-balancer-access-logs.html#access-log-entry-format
type NLBAccessLogRecord struct {
	Type                      string // Type of request (tls)
	Version                   string // Version of the log entry
	Time                      string // Timestamp the load balancer generated a response to the client in ISO 8601 format
	ELB                       string // Load balancer resource ID
	Listener                  string // Resource ID of the TLS listener for the connection
	ClientIP                  string // Client IP
	ClientPort                int64  // Client  port
	DestinationIP             string // Destination IP
	DestinationPort           int64  // Destination port
	ConnectionTime            int64  // Total time for the connection to complete, in milliseconds
	TLSHandshakeTime          int64  // Time for the TLS handshake to complete, in milliseconds, or -
	ReceivedBytes             int64  // Count of bytes received by the load balancer from the client, after decryption
	SentBytes                 int64  // Count of bytes sent by the load balancer to the client, before encryption
	IncomingTLSAlert          string // TLS alerts received from the client, if present, or -
	ChosenCertARN             string // ARN of the certificate served to the client, or -
	ChosenCertSerial          string // Reserved for future use, always set to -
	TLSCipher                 string // The cipher suite negotiated with the client, or -
	TLSProtocolVersion        string // The TLS protocol negotiated with the client, or -
	TLSNamedGroup             string // Reserved for future use, always set to -
	DomainName                string // Server name extension in the client hello message, or -
	ALPNFeProtocol            string // Application protocol negotiated with the client, or -
	ALPNBeProtocol            string // Application protocol negotiated with the target, or -
	ALPNClientPreferenceList  string // Client hello message ALPN list, URL-encoded, or -
	TLSConnectionCreationTime string // Time recorded at the start of the TLS connection, in ISO 8601 format
}

// convertTextToNLBAccessLogRecord converts a slice of strings into a NLBAccessLogRecord
func convertTextToNLBAccessLogRecord(fields []string) (NLBAccessLogRecord, error) {
	_ = "STUB: not implemented"

	// Check if the fields contain enough data
	return *new(NLBAccessLogRecord), nil
}

// Map fields to the struct

// Type of request
// Log version
// Timestamp
// Load balancer resource ID
// Listener ID
// TLSHandshakeTime placeholder value
// Incoming TLS alert
// Chosen certificate ARN
// Reserved for future use, usually set to '-'
// Negotiated TLS cipher
// Negotiated TLS protocol version
// Reserved for future use, usually set to '-'
// SNI domain provided by the client
// Protocol negotiated with client
// Protocol negotiated with target
// ALPN preference list from the client
// Time of the start of the TLS connection

// Processing additional fields if applicable

// Application Load Balancer Access Logs record
// Doc: https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-access-logs.html
type ALBAccessLogRecord struct {
	Type                   string // Type of request (http, https, etc.)
	Time                   string // Timestamp the load balancer generated a response to the client in ISO 8601 format
	ELB                    string // Load balancer resource ID
	ClientIP               string // Client IP
	ClientPort             int64  // Client  port
	TargetIPPort           string // Target IP:Port or -
	TargetIP               string // Target IP
	TargetPort             int64  // Target port
	RequestProcessingTime  string // Time taken to process the request in milliseconds
	TargetProcessingTime   string // Time taken for the target to process the request in milliseconds
	ResponseProcessingTime string // Time taken to send the response to the client in milliseconds
	ELBStatusCode          int64  // Status code from the load balancer
	TargetStatusCode       string // Status code from the target
	ReceivedBytes          int64  // Size of the request in bytes
	SentBytes              int64  // Size of the response in bytes
	RequestMethod          string // HTTP method used in the request (e.g., GET, POST)
	RequestURI             string // Full URI requested by the client
	ProtocolName           string // Network protocol name (e.g., HTTP)
	ProtocolVersion        string // Network protocol version (e.g., 1.1)
	UserAgent              string // Client's User-Agent string
	SSLCipher              string // SSL cipher
	SSLProtocol            string // SSL protocol
	TargetGroupARN         string // Target group ARN
	TraceID                string // X-Amzn-Trace-Id header contents
	DomainName             string // SNI domain provided by the client
	ChosenCertARN          string // ARN of the chosen certificate
	MatchedRulePriority    string // Priority of the rule that matched
	RequestCreationTime    string // Time when the load balancer received the request from the client in ISO 8601 format
	ActionsExecuted        string // Actions executed, comma-separated
	RedirectURL            string // URL of the redirect target
	ErrorReason            string // Reason for request failure
	TargetPortList         string // List of target IPs/ports (if applicable)
	TargetStatusCodeList   string // List of status codes from targets
	Classification         string // Classification of the request
	ClassificationReason   string // Reason for classification
	ConnectionTraceID      string // The connection traceability ID
	TransformedHost        string // The transformed host header
	TransformedURI         string // The URI after it is modified by a URL rewrite transform
	RequestTransformStatus string // The status of the rewrite transform
}

// convertTextToALBAccessLogRecord converts a slice of strings into a ALBAccessLogRecord
func convertTextToALBAccessLogRecord(fields []string) (ALBAccessLogRecord, error) {
	_ = "STUB: not implemented"
	return *new(ALBAccessLogRecord), nil
}

// Map fields to the struct

// Parse TargetIPPort into TargetIP and TargetPort

func safeConvertStrToInt(stringNum string) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func safeConvertStrToFloat(stringNum string) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// findLogSyntaxByField determines the log syntax type based on the first field of a log entry.
// It first checks for known protocol types (ALB and NLB).
// ALB supports http, https, h2, grpcs, ws, wss and NLB supports tls.
// Only if those are not matched, it checks if the field is a valid timestamp (for CLB logs).
// If none match, it returns an error.
func findLogSyntaxByField(field string) (logSyntaxType, error) {
	_ = "STUB: not implemented"
	return *new(logSyntaxType), nil
}

// isValidTimestamp checks if the given field is a valid timestamp
func isValidTimestamp(field string) bool { _ = "STUB: not implemented"; return false }

// convertToUnixEpoch converts ISO 8601 timestamps to UNIX Epoch time in nanoseconds
func convertToUnixEpoch(isoTimestamp string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Check if the timestamp has sub-second precision

// Parse complete ISO 8601 timestamp with microseconds including Zone

//  Parse complete ISO 8601 timestamp with microseconds without Zone

// Parse timestamp without sub-second precision

// Return the UNIX epoch time in nanoseconds

// scanField gets the next value in the log line by moving through space delimiters.
// If the value starts with a quote, it moves forward until it finds the ending quote.
// Note that quotes are not preserved. For example "a","b" becomes a,b.
// Otherwise, it returns the value as it is.
func scanField(logLine string) (value, remainder string, err error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

// preserve values without quotes, terminate at space after the closing quote or at line end

// terminating space found, return the value

// No quotes means we are at the end of log line

// Invalid log line - must not happen in well-formed logs

func extractFields(logLine string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

// parseRequestField splits a raw HTTP request line into its components:
// method, URI, protocol name, and protocol version.
// Expected format: "<METHOD> <URI> <PROTOCOL>/<VERSION>", e.g. "GET http://example.com HTTP/1.1".
func parseRequestField(raw string) (method, uri, protoName, protoVersion string, err error) {
	_ = "STUB: not implemented"
	return "", "", "", "", nil
}

// netProtocol returns protocol name and version based on proto value
func netProtocol(proto string) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

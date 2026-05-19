// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package azurelogs // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/azurelogs"

import (
	"errors"

	"go.opentelemetry.io/collector/pdata/plog"
)

const (
	noError = "NoError"

	categoryAzureCdnAccessLog                  = "AzureCdnAccessLog"
	categoryFrontDoorAccessLog                 = "FrontDoorAccessLog"
	categoryFrontDoorHealthProbeLog            = "FrontDoorHealthProbeLog"
	categoryFrontdoorWebApplicationFirewallLog = "FrontDoorWebApplicationFirewallLog"
	categoryAppServiceAppLogs                  = "AppServiceAppLogs"
	categoryAppServiceAuditLogs                = "AppServiceAuditLogs"
	categoryAppServiceAuthenticationLogs       = "AppServiceAuthenticationLogs"
	categoryAppServiceConsoleLogs              = "AppServiceConsoleLogs"
	categoryAppServiceHTTPLogs                 = "AppServiceHTTPLogs"
	categoryAppServiceIPSecAuditLogs           = "AppServiceIPSecAuditLogs"
	categoryAppServicePlatformLogs             = "AppServicePlatformLogs"
	categoryAdministrative                     = "Administrative"
	categoryAlert                              = "Alert"
	categoryAutoscale                          = "Autoscale"
	categorySecurity                           = "Security"
	categoryPolicy                             = "Policy"
	categoryRecommendation                     = "Recommendation"
	categoryServiceHealth                      = "ServiceHealth"
	categoryResourceHealth                     = "ResourceHealth"

	// attributeAzureRef holds the request tracking reference, also
	// placed in the request header "X-Azure-Ref".
	attributeAzureRef = "azure.ref"

	// attributeTimeToFirstByte holds the length of time in milliseconds
	// from when Microsoft service (CDN, Front Door, etc) receives the
	// request to the time the first byte gets sent to the client.
	attributeTimeToFirstByte = "azure.time_to_first_byte"

	// attributeDuration holds the length of time from first byte of
	// request to last byte of response out, in seconds.
	attributeDuration = "duration"

	// attributeAzurePop holds the point of presence (POP) that
	// processed the request
	attributeAzurePop = "azure.pop"

	// attributeCacheStatus holds the result of the cache hit/miss
	// at the POP
	attributeCacheStatus = "azure.cache_status"

	// attributeTLSServerName holds the server name indication (SNI)
	// value
	attributeTLSServerName = "tls.server.name"

	missingPort = "missing port in address"
)

const (
	// Identity specific attributes
	attributeIdentityAuthorizationScope  = "azure.identity.authorization.scope"
	attributeIdentityAuthorizationAction = "azure.identity.authorization.action"
	// Identity > authorization > evidence
	attributeIdentityAuthorizationEvidenceRole                = "azure.identity.authorization.evidence.role"
	attributeIdentityAuthorizationEvidenceRoleAssignmentScope = "azure.identity.authorization.evidence.role.assignment.scope"
	attributeIdentityAuthorizationEvidenceRoleAssignmentID    = "azure.identity.authorization.evidence.role.assignment.id"
	attributeIdentityAuthorizationEvidenceRoleDefinitionID    = "azure.identity.authorization.evidence.role.definition.id"
	attributeIdentityAuthorizationEvidencePrincipalID         = "azure.identity.authorization.evidence.principal.id"
	attributeIdentityAuthorizationEvidencePrincipalType       = "azure.identity.authorization.evidence.principal.type"
	// Identity > claims (standard JWT claims)
	attributeIdentityClaimsAudience  = "azure.identity.audience"
	attributeIdentityClaimsIssuer    = "azure.identity.issuer"
	attributeIdentityClaimsSubject   = "azure.identity.subject"
	attributeIdentityClaimsNotAfter  = "azure.identity.not_after"
	attributeIdentityClaimsNotBefore = "azure.identity.not_before"
	attributeIdentityClaimsCreated   = "azure.identity.created"
	// Identity > claims (Azure specific claims)
	attributeIdentityClaimsScope                 = "azure.identity.scope"
	attributeIdentityClaimsType                  = "azure.identity.type"
	attributeIdentityClaimsApplicationID         = "azure.identity.application.id"
	attributeIdentityClaimsAuthMethodsReferences = "azure.identity.auth.methods.references"
	attributeIdentityClaimsIdentifierObject      = "azure.identity.identifier.object"
	attributeIdentityClaimsIdentifierName        = "user.name"
	attributeIdentityClaimsProvider              = "azure.identity.provider"

	// azure front door WAF attributes

	// attributeAzureFrontDoorWAFRuleName holds the name of the WAF rule that
	// the request matched.
	attributeAzureFrontDoorWAFRuleName = "azure.frontdoor.waf.rule.name"

	// attributeAzureFrontDoorWAFPolicyName holds the name of the WAF policy
	// that processed the request.
	attributeAzureFrontDoorWAFPolicyName = "azure.frontdoor.waf.policy.name"

	// attributeAzureFrontDoorWAFPolicyMode holds the operations mode of the
	// WAF policy.
	attributeAzureFrontDoorWAFPolicyMode = "azure.frontdoor.waf.policy.mode"

	// attributeAzureFrontDoorWAFAction holds the action taken on the request.
	attributeAzureFrontDoorWAFAction = "azure.frontdoor.waf.action"

	// Administrative specific attributes
	attributeAzureAdministrativeEntity    = "azure.administrative.entity"
	attributeAzureAdministrativeMessage   = "azure.administrative.message"
	attributeAzureAdministrativeHierarchy = "azure.administrative.hierarchy"

	// Alert specific attributes
	attributeAzureAlertWebhookURI      = "azure.alert.webhook.uri"
	attributeAzureAlertRuleURI         = "azure.alert.rule.uri"
	attributeAzureAlertRuleName        = "azure.alert.rule.name"
	attributeAzureAlertRuleDescription = "azure.alert.rule.description"
	attributeAzureAlertThreshold       = "azure.alert.threshold"
	attributeAzureAlertWindowSize      = "azure.alert.window_size_minutes"
	attributeAzureAlertAggregation     = "azure.alert.aggregation"
	attributeAzureAlertOperator        = "azure.alert.operator"
	attributeAzureAlertMetricName      = "azure.alert.metric.name"
	attributeAzureAlertMetricUnit      = "azure.alert.metric.unit"

	// Autoscale specific attributes
	attributeAzureAutoscaleDescription     = "azure.autoscale.description"
	attributeAzureAutoscaleResourceName    = "azure.autoscale.resource.name"
	attributeAzureAutoscaleOldInstances    = "azure.autoscale.instances.previous_count"
	attributeAzureAutoscaleNewInstances    = "azure.autoscale.instances.count"
	attributeAzureAutoscaleLastScaleAction = "azure.autoscale.resource.last_scale"

	// Policy specific attributes
	attributeAzurePolicyIsComplianceCheck = "azure.policy.compliance_check"
	attributeAzurePolicyAncestors         = "azure.policy.ancestors"
	attributeAzurePolicyHierarchy         = "azure.policy.hierarchy"

	// Recommendation specific attributes
	attributeAzureRecommendationCategory      = "azure.recommendation.category"
	attributeAzureRecommendationImpact        = "azure.recommendation.impact"
	attributeAzureRecommendationName          = "azure.recommendation.name"
	attributeAzureRecommendationType          = "azure.recommendation.type"
	attributeAzureRecommendationSchemaVersion = "azure.recommendation.schema_version"
	attributeAzureRecommendationLink          = "azure.recommendation.link"

	// Security specific attributes
	attributeAzureSecurityAccountLogonID = "azure.security.account_logon_id"
	attributeAzureSecurityDomainName     = "azure.security.domain_name"
	attributeAzureSecurityActionTaken    = "azure.security.action_taken"
	attributeAzureSecuritySeverity       = "azure.security.severity"

	// Service Health specific attributes
	attributeAzureServiceHealthTitle                    = "azure.servicehealth.title"
	attributeAzureServiceHealthService                  = "azure.servicehealth.service"
	attributeAzureServiceHealthRegion                   = "azure.servicehealth.region"
	attributeAzureServiceHealthCommunicationID          = "azure.servicehealth.communication.id"
	attributeAzureServiceHealthCommunicationBody        = "azure.servicehealth.communication.body"
	attributeAzureServiceHealthCommunicationRouteType   = "azure.servicehealth.communication.route_type"
	attributeAzureServiceHealthIncidentType             = "azure.servicehealth.incident.type"
	attributeAzureServiceHealthTrackingID               = "azure.servicehealth.tracking.id"
	attributeAzureServiceHealthImpactStartTime          = "azure.servicehealth.impact.start"
	attributeAzureServiceHealthImpactMitigationTime     = "azure.servicehealth.impact.mitigation"
	attributeAzureServiceHealthImpactedServices         = "azure.servicehealth.impact.services"
	attributeAzureServiceHealthImpactType               = "azure.servicehealth.impact.type"
	attributeAzureServiceHealthImpactCategory           = "azure.servicehealth.impact.category"
	attributeAzureServiceHealthDefaultLanguageTitle     = "azure.servicehealth.default_language.title"
	attributeAzureServiceHealthDefaultLanguageContent   = "azure.servicehealth.default_language.content"
	attributeAzureServiceHealthState                    = "azure.servicehealth.state"
	attributeAzureServiceHealthMaintenanceID            = "azure.servicehealth.maintenance.id"
	attributeAzureServiceHealthMaintenanceType          = "azure.servicehealth.maintenance.type"
	attributeAzureServiceHealthIsHIR                    = "azure.servicehealth.is_hir"
	attributeAzureServiceHealthIsSynthetic              = "azure.servicehealth.is_synthetic"
	attributeAzureServiceHealthEmailTemplateID          = "azure.servicehealth.email.template.id"
	attributeAzureServiceHealthEmailTemplateFullVersion = "azure.servicehealth.email.template.full_version"
	attributeAzureServiceHealthEmailTemplateLocale      = "azure.servicehealth.email.template.locale"
	attributeAzureServiceHealthSMSText                  = "azure.servicehealth.sms.text"
	attributeAzureServiceHealthVersion                  = "azure.servicehealth.version"
	attributeAzureServiceHealthArgQuery                 = "azure.servicehealth.arg_query"
	attributeAzureServiceHealthRateNew                  = "azure.servicehealth.new_rate"
	attributeAzureServiceHealthRateOld                  = "azure.servicehealth.old_rate"

	// Resource Health specific attributes
	attributeAzureResourceHealthTitle                = "azure.resourcehealth.title"
	attributeAzureResourceHealthDetails              = "azure.resourcehealth.details"
	attributeAzureResourceHealthCurrentHealthStatus  = "azure.resourcehealth.state"
	attributeAzureResourceHealthPreviousHealthStatus = "azure.resourcehealth.previous_state"
	attributeAzureResourceHealthType                 = "azure.resourcehealth.type"
	attributeAzureResourceHealthCause                = "azure.resourcehealth.cause"
)

var (
	errStillToImplement    = errors.New("still to implement")
	errUnsupportedCategory = errors.New("category not supported")
)

func addRecordAttributes(category string, data []byte, record plog.LogRecord) error {
	_ = "STUB: not implemented"
	return nil
}

// putInt parses value as an int and puts it in the record
func putInt(field, value string, record plog.LogRecord) error {
	_ = "STUB: not implemented"
	return nil
}

// putStr puts the value in the record if the value holds
// meaningful data. Meaningful data is defined as not being empty
// or "N/A".
func putStr(field, value string, record plog.LogRecord) { _ = "STUB: not implemented"; return }

// ignore

func putBool(field, value string, record plog.LogRecord) { _ = "STUB: not implemented"; return }

// handleTime parses the time value and always multiplies it by
// 1e3. This is so we don't loose so much data if the time is for
// example "0.154". In that case, the output would be "154".
func handleTime(field, value string, record plog.LogRecord) error {
	_ = "STUB: not implemented"
	return nil
}

// See https://github.com/MicrosoftDocs/azure-docs/blob/main/articles/cdn/monitoring-and-access-log.md
type azureCdnAccessLogProperties struct {
	TrackingReference    string `json:"trackingReference"`
	HTTPMethod           string `json:"httpMethod"`
	HTTPVersion          string `json:"httpVersion"`
	RequestURI           string `json:"requestUri"`
	SNI                  string `json:"sni"`
	RequestBytes         string `json:"requestBytes"`
	ResponseBytes        string `json:"responseBytes"`
	UserAgent            string `json:"userAgent"`
	ClientIP             string `json:"clientIp"`
	ClientPort           string `json:"clientPort"`
	SocketIP             string `json:"socketIp"`
	TimeToFirstByte      string `json:"timeToFirstByte"`
	TimeTaken            string `json:"timeTaken"`
	RequestProtocol      string `json:"requestProtocol"`
	SecurityProtocol     string `json:"securityProtocol"`
	HTTPStatusCode       string `json:"httpStatusCode"`
	Pop                  string `json:"pop"`
	CacheStatus          string `json:"cacheStatus"`
	ErrorInfo            string `json:"errorInfo"`
	ErrorInfo1           string `json:"ErrorInfo"`
	Endpoint             string `json:"endpoint"`
	IsReceivedFromClient bool   `json:"isReceivedFromClient"`
	BackendHostname      string `json:"backendHostname"`
}

// addRequestURIProperties parses the request URI and adds the
// relevant attributes to the record
func addRequestURIProperties(uri string, record plog.LogRecord) error {
	_ = "STUB: not implemented"
	return nil
}

// addSecurityProtocolProperties based on the security protocol
func addSecurityProtocolProperties(securityProtocol string, record plog.LogRecord) error {
	_ = "STUB: not implemented"
	return nil
}

// addErrorInfoProperties checks if there is an error and adds it
// to the record attributes as an exception in case it exists.
func addErrorInfoProperties(errorInfo string, record plog.LogRecord) {
	_ = "STUB: not implemented"
	return
}

// handleDestination puts the value for the backend host name and endpoint
// in the expected field names. If backend hostname is empty, then the
// destination address and port depend only on the endpoint. If backend
// hostname is filled, then the destination address and port are based on
// it. If both are filled but different, then the endpoint will cover the
// network address and port.
func handleDestination(backendHostname, endpoint string, record plog.LogRecord) error {
	_ = "STUB: not implemented"
	return nil
}

// there is no port, so let's keep using the full endpoint for the address

// addAzureCdnAccessLogProperties parses the Azure CDN access log, and adds
// the relevant attributes to the record
func addAzureCdnAccessLogProperties(data []byte, record plog.LogRecord) error {
	_ = "STUB: not implemented"
	return nil
}

// See https://learn.microsoft.com/en-us/azure/frontdoor/monitor-front-door?pivots=front-door-standard-premium#access-log.
type frontDoorAccessLog struct {
	TrackingReference string `json:"trackingReference"`
	HTTPMethod        string `json:"httpMethod"`
	HTTPVersion       string `json:"httpVersion"`
	RequestURI        string `json:"requestUri"`
	SNI               string `json:"sni"`
	RequestBytes      string `json:"requestBytes"`
	ResponseBytes     string `json:"responseBytes"`
	UserAgent         string `json:"userAgent"`
	ClientIP          string `json:"clientIp"`
	ClientPort        string `json:"clientPort"`
	SocketIP          string `json:"socketIp"`
	TimeToFirstByte   string `json:"timeToFirstByte"`
	TimeTaken         string `json:"timeTaken"`
	RequestProtocol   string `json:"requestProtocol"`
	SecurityProtocol  string `json:"securityProtocol"`
	HTTPStatusCode    string `json:"httpStatusCode"`
	Pop               string `json:"pop"`
	CacheStatus       string `json:"cacheStatus"`
	ErrorInfo         string `json:"errorInfo"`
	ErrorInfo1        string `json:"ErrorInfo"`
	Result            string `json:"result"`
	Endpoint          string `json:"endpoint"`
	HostName          string `json:"hostName"`
	SecurityCipher    string `json:"securityCipher"`
	SecurityCurves    string `json:"securityCurves"`
	OriginIP          string `json:"originIp"`
}

// addFrontDoorAccessLogProperties parses the Front Door access log, and adds
// the relevant attributes to the record
func addFrontDoorAccessLogProperties(data []byte, record plog.LogRecord) error {
	_ = "STUB: not implemented"
	return nil
}

// addFrontDoorHealthProbeLogProperties parses the Front Door access log, and adds
// the relevant attributes to the record
func addFrontDoorHealthProbeLogProperties(_ []byte, _ plog.LogRecord) error {
	_ = "STUB: not implemented"
	// TODO @constanca-m implement this the same way as addAzureCdnAccessLogProperties
	return nil
}

// See https://learn.microsoft.com/en-us/azure/web-application-firewall/afds/waf-front-door-monitor?pivots=front-door-standard-premium#waf-logs
type frontDoorWAFLogProperties struct {
	ClientIP          string `json:"clientIP"`
	ClientPort        string `json:"clientPort"`
	SocketIP          string `json:"socketIP"`
	RequestURI        string `json:"requestUri"`
	RuleName          string `json:"ruleName"`
	Policy            string `json:"policy"`
	Action            string `json:"action"`
	Host              string `json:"host"`
	TrackingReference string `json:"trackingReference"`
	PolicyMode        string `json:"policyMode"`
}

// addFrontDoorWAFLogProperties parses the Front Door access log, and adds
// the relevant attributes to the record
func addFrontDoorWAFLogProperties(data []byte, record plog.LogRecord) error {
	_ = "STUB: not implemented"
	return nil
}

// addAppServiceAppLogsProperties parses the App Service access log, and adds
// the relevant attributes to the record
func addAppServiceAppLogsProperties(_ []byte, _ plog.LogRecord) error {
	_ = "STUB: not implemented"
	// TODO @constanca-m implement this the same way as addAzureCdnAccessLogProperties
	return nil
}

// addAppServiceAuditLogsProperties parses the App Service access log, and adds
// the relevant attributes to the record
func addAppServiceAuditLogsProperties(_ []byte, _ plog.LogRecord) error {
	_ = "STUB: not implemented"
	// TODO @constanca-m implement this the same way as addAzureCdnAccessLogProperties
	return nil
}

// addAppServiceAuthenticationLogsProperties parses the App Service access log, and adds
// the relevant attributes to the record
func addAppServiceAuthenticationLogsProperties(_ []byte, _ plog.LogRecord) error {
	_ = "STUB: not implemented"
	// TODO @constanca-m implement this the same way as addAzureCdnAccessLogProperties
	return nil
}

// addAppServiceConsoleLogsProperties parses the App Service access log, and adds
// the relevant attributes to the record
func addAppServiceConsoleLogsProperties(_ []byte, _ plog.LogRecord) error {
	_ = "STUB: not implemented"
	// TODO @constanca-m implement this the same way as addAzureCdnAccessLogProperties
	return nil
}

// addAppServiceHTTPLogsProperties parses the App Service access log, and adds
// the relevant attributes to the record
func addAppServiceHTTPLogsProperties(_ []byte, _ plog.LogRecord) error {
	_ = "STUB: not implemented"
	// TODO @constanca-m implement this the same way as addAzureCdnAccessLogProperties
	return nil
}

// addAppServiceIPSecAuditLogsProperties parses the App Service access log, and adds
// the relevant attributes to the record
func addAppServiceIPSecAuditLogsProperties(_ []byte, _ plog.LogRecord) error {
	_ = "STUB: not implemented"
	// TODO @constanca-m implement this the same way as addAzureCdnAccessLogProperties
	return nil
}

// addAppServicePlatformLogsProperties parses the App Service access log, and adds
// the relevant attributes to the record
func addAppServicePlatformLogsProperties(_ []byte, _ plog.LogRecord) error {
	_ = "STUB: not implemented"
	// TODO @constanca-m implement this the same way as addAzureCdnAccessLogProperties
	return nil
}

// ------------------------------------------------------------
// Activity Log - Administrative category
// ------------------------------------------------------------

// administrativeLogProperties represents the properties field of an Administrative activity log.
// See: https://learn.microsoft.com/en-us/azure/azure-monitor/essentials/activity-log-schema#administrative-category
type administrativeLogProperties struct {
	Entity    string `json:"entity"`
	Message   string `json:"message"`
	Hierarchy string `json:"hierarchy"`
}

// addAdministrativeLogProperties parses Administrative activity logs
// and maps them to OpenTelemetry semantic conventions.
// See: https://learn.microsoft.com/en-us/azure/azure-monitor/essentials/activity-log-schema#administrative-category
func addAdministrativeLogProperties(data []byte, record plog.LogRecord) error {
	_ = "STUB: not implemented"
	return nil
}

// ------------------------------------------------------------
// Activity Log - Alert category
// ------------------------------------------------------------

// alertLogProperties represents the properties field of an Alert activity log.
// See: https://learn.microsoft.com/en-us/azure/azure-monitor/essentials/activity-log-schema#alert-category
type alertLogProperties struct {
	WebHookURI          string `json:"webHookUri"`
	RuleURI             string `json:"RuleUri"`
	RuleName            string `json:"RuleName"`
	RuleDescription     string `json:"RuleDescription"`
	Threshold           string `json:"Threshold"`
	WindowSizeInMinutes string `json:"WindowSizeInMinutes"`
	Aggregation         string `json:"Aggregation"`
	Operator            string `json:"Operator"`
	MetricName          string `json:"MetricName"`
	MetricUnit          string `json:"MetricUnit"`
}

// addAlertLogProperties parses Alert activity logs
// and maps them to OpenTelemetry semantic conventions.
// See: https://learn.microsoft.com/en-us/azure/azure-monitor/essentials/activity-log-schema#alert-category
func addAlertLogProperties(data []byte, record plog.LogRecord) error {
	_ = "STUB: not implemented"
	return nil
}

// ------------------------------------------------------------
// Activity Log - Autoscale category
// ------------------------------------------------------------

// autoscaleLogProperties represents the properties field of an Autoscale activity log.
// See: https://learn.microsoft.com/en-us/azure/azure-monitor/essentials/activity-log-schema#autoscale-category
type autoscaleLogProperties struct {
	Description         string `json:"Description"`
	ResourceName        string `json:"ResourceName"`
	OldInstancesCount   string `json:"OldInstancesCount"`
	NewInstancesCount   string `json:"NewInstancesCount"`
	LastScaleActionTime string `json:"LastScaleActionTime"`
}

// addAutoscaleLogProperties parses Autoscale activity logs
// and maps them to OpenTelemetry semantic conventions.
// See: https://learn.microsoft.com/en-us/azure/azure-monitor/essentials/activity-log-schema#autoscale-category
func addAutoscaleLogProperties(data []byte, record plog.LogRecord) error {
	_ = "STUB: not implemented"
	return nil
}

// ------------------------------------------------------------
// Activity Log - Policy category
// ------------------------------------------------------------

type policyElement struct {
	DefinitionID             string   `json:"policyDefinitionId"`
	SetDefinitionID          string   `json:"policySetDefinitionId"`
	ReferenceID              string   `json:"policyDefinitionReferenceId"`
	SetDefinitionName        string   `json:"policySetDefinitionName"`
	SetDefinitionDisplayName string   `json:"policySetDefinitionDisplayName"`
	SetDefinitionVersion     string   `json:"policySetDefinitionVersion"`
	DefinitionName           string   `json:"policyDefinitionName"`
	DefinitionDisplayName    string   `json:"policyDefinitionDisplayName"`
	DefinitionVersion        string   `json:"policyDefinitionVersion"`
	DefinitionEffect         string   `json:"policyDefinitionEffect"`
	AssignmentID             string   `json:"policyAssignmentId"`
	AssignmentName           string   `json:"policyAssignmentName"`
	AssignmentDisplayName    string   `json:"policyAssignmentDisplayName"`
	AssignmentScope          string   `json:"policyAssignmentScope"`
	ExemptionIDs             []string `json:"policyExemptionIds"`
	AssignmentIDs            []string `json:"policyAssignmentIds"`
}

type policyLogProperties struct {
	IsComplianceCheck string `json:"isComplianceCheck"`
	ResourceLocation  string `json:"resourceLocation"`
	Ancestors         string `json:"ancestors"`
	Policies          string `json:"policies"`
	Hierarchy         string `json:"hierarchy"`
}

func addPolicyLogProperties(data []byte, record plog.LogRecord) error {
	_ = "STUB: not implemented"
	return nil
}

// check if Policies is a string and unmarshal the embedded JSON
// object in the `policyElement` struct

// Add policies as a slice of maps

// ------------------------------------------------------------
// Activity Log - Recommendation category
// ------------------------------------------------------------

// recommendationLogProperties represents the properties field of a Recommendation activity log.
// See: https://learn.microsoft.com/en-us/azure/azure-monitor/essentials/activity-log-schema#recommendation-category
type recommendationLogProperties struct {
	RecommendationSchemaVersion string `json:"recommendationSchemaVersion"`
	RecommendationCategory      string `json:"recommendationCategory"`
	RecommendationImpact        string `json:"recommendationImpact"`
	RecommendationName          string `json:"recommendationName"`
	RecommendationResourceLink  string `json:"recommendationResourceLink"`
	RecommendationType          string `json:"recommendationType"`
}

// addRecommendationLogProperties parses Recommendation activity logs from Azure Advisor
// and maps them to OpenTelemetry semantic conventions.
// See: https://learn.microsoft.com/en-us/azure/azure-monitor/essentials/activity-log-schema#recommendation-category
func addRecommendationLogProperties(data []byte, record plog.LogRecord) error {
	_ = "STUB: not implemented"
	return nil
}

// ------------------------------------------------------------
// Activity Log - Security category
// ------------------------------------------------------------

// securityLogProperties represents the properties field of a Security activity log.
// See: https://learn.microsoft.com/en-us/azure/azure-monitor/essentials/activity-log-schema#security-category
type securityLogProperties struct {
	AccountLogonID  string `json:"accountLogonId"`
	CommandLine     string `json:"commandLine"`
	DomainName      string `json:"domainName"`
	ParentProcess   string `json:"parentProcess"`
	ParentProcessID string `json:"parentProcess id"`
	ProcessID       string `json:"processId"`
	ProcessName     string `json:"processName"`
	UserName        string `json:"userName"`
	UserSID         string `json:"UserSID"`
	ActionTaken     string `json:"ActionTaken"`
	Severity        string `json:"Severity"`
}

// addSecurityLogProperties parses Security activity logs
// and maps them to OpenTelemetry semantic conventions.
// See: https://learn.microsoft.com/en-us/azure/azure-monitor/essentials/activity-log-schema#security-category
func addSecurityLogProperties(data []byte, record plog.LogRecord) error {
	_ = "STUB: not implemented"
	return nil
}

// Map to OTel process semantic conventions

// Azure-specific fields that don't have OTel equivalents

// ------------------------------------------------------------
// Activity Log - Service Health category
// ------------------------------------------------------------

// serviceHealthLogProperties represents the properties field of a Service Health activity log.
// See: https://learn.microsoft.com/en-us/azure/azure-monitor/essentials/activity-log-schema#service-health-category
type serviceHealthLogProperties struct {
	Title                  string `json:"title"`
	Service                string `json:"service"`
	Region                 string `json:"region"`
	CommunicationText      string `json:"communication"`
	CommunicationID        string `json:"communicationId"`
	IncidentType           string `json:"incidentType"`
	TrackingID             string `json:"trackingId"`
	ImpactStartTime        string `json:"impactStartTime"`
	ImpactMitigationTime   string `json:"impactMitigationTime"`
	ImpactedServices       string `json:"impactedServices"`
	DefaultLanguageTitle   string `json:"defaultLanguageTitle"`
	DefaultLanguageContent string `json:"defaultLanguageContent"`
	Stage                  string `json:"stage"`
	MaintenanceID          string `json:"maintenanceId"`
	MaintenanceType        string `json:"maintenanceType"`
	IsHIR                  bool   `json:"isHIR"`
	IsSynthetic            string `json:"IsSynthetic"`
	ImpactType             string `json:"impactType"`
	ImpactCategory         string `json:"impactCategory"`
}

type impactedService struct {
	Name    string `json:"ServiceName"`
	ID      string `json:"ServiceId"`
	GUID    string `json:"ServiceGuid"`
	Regions []struct {
		Name string `json:"RegionName"`
		ID   string `json:"RegionId"`
	} `json:"ImpactedRegions"`
}

// addServiceHealthLogProperties parses Service Health activity logs
// and maps them to OpenTelemetry semantic conventions.
// See: https://learn.microsoft.com/en-us/azure/azure-monitor/essentials/activity-log-schema#service-health-category
func addServiceHealthLogProperties(data []byte, record plog.LogRecord) error {
	_ = "STUB: not implemented"
	return nil
}

// check if Policies is a string and unmarshal the embedded JSON
// object in the `policyElement` struct

// Add impacted services as a slice of maps

// ------------------------------------------------------------
// Activity Log - Resource Health category
// ------------------------------------------------------------

// resourceHealthLogProperties represents the properties field of a Resource Health activity log.
// See: https://learn.microsoft.com/en-us/azure/azure-monitor/essentials/activity-log-schema#resource-health-category
type resourceHealthLogProperties struct {
	Title                string `json:"title"`
	Details              string `json:"details"`
	CurrentHealthStatus  string `json:"currentHealthStatus"`
	PreviousHealthStatus string `json:"previousHealthStatus"`
	Type                 string `json:"type"`
	Cause                string `json:"cause"`
}

// addResourceHealthLogProperties parses Resource Health activity logs
// and maps them to OpenTelemetry semantic conventions.
// See: https://learn.microsoft.com/en-us/azure/azure-monitor/essentials/activity-log-schema#resource-health-category
func addResourceHealthLogProperties(data []byte, record plog.LogRecord) error {
	_ = "STUB: not implemented"
	return nil
}

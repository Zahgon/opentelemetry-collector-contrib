// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package auditlog // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding/googlecloudlogentryencodingextension/internal/auditlog"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
)

const (
	ActivityLogNameSuffix    = "cloudaudit.googleapis.com%2Factivity"
	DataAccessLogNameSuffix  = "cloudaudit.googleapis.com%2Fdata_access"
	SystemEventLogNameSuffix = "cloudaudit.googleapis.com%2Fsystem_event"
	PolicyLogNameSuffix      = "cloudaudit.googleapis.com%2Fpolicy"

	auditLogType = "type.googleapis.com/google.cloud.audit.AuditLog"

	// gcpAuditServiceName holds the name of the API service performing the operation
	gcpAuditServiceName = "gcp.audit.service.name"
	// gcpAuditMethodName holds the name of the service method or operation
	gcpAuditMethodName = "gcp.audit.method.name"
	// gcpAuditResourceName holds the value for the operation target
	gcpAuditResourceName = "gcp.audit.resource.name"

	// gcpAuditResourceLocationCurrent holds the locations of a resource after the execution of the operation
	gcpAuditResourceLocationCurrent = "gcp.audit.resource.location.current"
	// gcpAuditResourceLocationOriginal holds the locations of a resource prior to the execution of the operation
	gcpAuditResourceLocationOriginal = "gcp.audit.resource.location.original"

	// gcpAuditNUmResponseItems holds the number of items returned from a List or Query API method, if applicable
	gcpAuditNumResponseItems = "gcp.audit.response.items"

	// gcpAuditAuthenticationAuthoritySelector holds the authority selector specified by the requestor, if any
	gcpAuditAuthenticationAuthoritySelector = "gcp.audit.authentication.authority_selector"
	// gcpAuditAuthenticationServiceAccountKeyName holds the name of the service account key used to create/exchange
	// credentials for authenticating the service account making the request
	gcpAuditAuthenticationServiceAccountKeyName = "gcp.audit.authentication.service_account.key.name"

	gcpAuditAuthorization = "gcp.audit.authorization"
	// gcpAuditAuthorizationResource holds the value of the resource being accessed
	gcpAuditAuthorizationResource = "resource"
	// gcpAuditAuthorizationPermission holds the required IAM permission
	gcpAuditAuthorizationPermission = "permission"
	// gcpAuditAuthorizationGranted holds the value on whether the authorization for resource and permission was granted
	gcpAuditAuthorizationGranted = "granted"

	// gcpAuditRequestCallerNetwork holds the network of the request caller
	gcpAuditRequestCallerNetwork = "gcp.audit.request.caller.network"
	// gcpAuditRequestReason holds the reason for the request.
	gcpAuditRequestReason = "gcp.audit.request.reason"
	// gcpAuditRequestTime holds the timestamp for when the destination receives the last byte of the request
	gcpAuditRequestTime = "gcp.audit.request.time"
	// httpRequestID will hold the request ID from requestMetadata.requestAttributes.id
	httpRequestID = "http.request.id"
	// gcpAuditRequestAuthPrincipal holds the principal at transport level layer
	gcpAuditRequestAuthPrincipal = "gcp.audit.request.auth.principal"
	// gcpAuditRequestAuthAudiences holds the audience at transport level layer
	gcpAuditRequestAuthAudiences = "gcp.audit.request.auth.audiences"
	// gcpAuditRequestAuthPresenter holds the presenter at transport level layer
	gcpAuditRequestAuthPresenter = "gcp.audit.request.auth.presenter"
	// gcpAuditRequestAuthAccessLevels holds the list of access level resource names that allow
	// resources to be accessed, at transport level layer
	gcpAuditRequestAuthAccessLevels = "gcp.audit.request.auth.access_levels"

	// gcpAuditDestinationLabels holds the labels in the destination attributes
	gcpAuditDestinationLabels = "gcp.audit.destination.label"
	// gcpAuditDestinationPrincipal holds the identity of the destination
	gcpAuditDestinationPrincipal = "gcp.audit.destination.principal"
	// gcpAuditDestinationRegionCode holds the region code of the destination
	gcpAuditDestinationRegionCode = "gcp.audit.destination.region_code"

	// gcpAuditPolicyViolationResourceType holds the esource type that the orgpolicy is checked against
	gcpAuditPolicyViolationResourceType = "gcp.audit.policy_violation.resource.type"
	// gcpAuditPolicyViolationResourceTags holds the tags referenced on the resource at the time of evaluation
	gcpAuditPolicyViolationResourceTags = "gcp.audit.policy_violation.resource.tags"
	// gcpAuditPolicyViolationInfo holds the policy violations list
	gcpAuditPolicyViolationInfo = "gcp.audit.policy_violation.info"
	// gcpAuditPolicyViolationInfoErrorMessage holds the error message
	gcpAuditPolicyViolationInfoErrorMessage = "error_message"
	// gcpAuditPolicyViolationInfoConstraint holds the constraint
	gcpAuditPolicyViolationInfoConstraint = "constraint"
	// gcpAuditPolicyViolationInfoCheckedValue holds the value that is being checked for the policy
	gcpAuditPolicyViolationInfoCheckedValue = "checked_value"
	// gcpAuditPolicyViolationInfoPolicyType holds the policy type
	gcpAuditPolicyViolationInfoPolicyType = "policy_type"
)

type auditLog struct {
	Type        string `json:"@type"`
	ServiceName string `json:"serviceName"`
	MethodName  string `json:"methodName"`

	ResourceName     string            `json:"resourceName"`
	ResourceLocation *resourceLocation `json:"resourceLocation"`
	// TODO Add support for resourceOriginalState
	// We can deduct the format of this struct by the @type field. We should use it
	// to unmarshal this struct without affecting performance and following the right
	// semantic conventions.

	NumResponseItems string `json:"numResponseItems"`

	Status *status `json:"status"`

	AuthenticationInfo *authenticationInfo `json:"authenticationInfo"`

	AuthorizationInfo []authorizationInfo `json:"authorizationInfo"`

	PolicyViolationInfo *policyViolationInfo `json:"policyViolationInfo"`

	RequestMetadata *requestMetadata `json:"requestMetadata"`
}

type resourceLocation struct {
	CurrentLocations  []string `json:"currentLocations"`
	OriginalLocations []string `json:"originalLocations"`
}

type status struct {
	Code    *int64 `json:"code"`
	Message string `json:"message"`
	// TODO Add support for details
}

type authenticationInfo struct {
	PrincipalEmail        string `json:"principalEmail"`
	PrincipalSubject      string `json:"principalSubject"`
	AuthoritySelector     string `json:"authoritySelector"`
	ServiceAccountKeyName string `json:"serviceAccountKeyName"`

	// TODO Add support for thirdPartyPrincipal
	// TODO Add support for serviceAccountDelegationInfo
}

type authorizationInfo struct {
	Resource   string `json:"resource"`
	Permission string `json:"permission"`
	Granted    *bool  `json:"granted"`

	// TODO Add support for resourceAttributes
}

type policyViolationInfo struct {
	OrgPolicyViolationInfo *orgPolicyViolationInfo `json:"orgPolicyViolationInfo"`
}

type orgPolicyViolationInfo struct {
	ResourceType  string            `json:"resourceType"`
	ResourceTags  map[string]string `json:"resourceTags"`
	ViolationInfo []violationInfo   `json:"violationInfo"`

	// TODO Add support for payload
}

type violationInfo struct {
	Constraint   string `json:"constraint"`
	ErrorMessage string `json:"errorMessage"`
	CheckedValue string `json:"checkedValue"`
	PolicyType   string `json:"policyType"`
}

type requestMetadata struct {
	CallerIP                string                 `json:"callerIp"`
	CallerSuppliedUserAgent string                 `json:"callerSuppliedUserAgent"`
	CallerNetwork           string                 `json:"callerNetwork"`
	RequestAttributes       *requestAttributes     `json:"requestAttributes"`
	DestinationAttributes   *destinationAttributes `json:"destinationAttributes"`
}

type requestAttributes struct {
	ID       string            `json:"id"`
	Method   string            `json:"method"`
	Headers  map[string]string `json:"headers"`
	Path     string            `json:"path"`
	Host     string            `json:"host"`
	Scheme   string            `json:"scheme"`
	Query    string            `json:"query"`
	Time     string            `json:"time"`
	Size     string            `json:"size"`
	Protocol string            `json:"protocol"`
	Reason   string            `json:"reason"`
	Auth     auth              `json:"auth"`
}

type auth struct {
	Principal    string   `json:"principal"`
	Audiences    []string `json:"audiences"`
	Presenter    string   `json:"presenter"`
	AccessLevels []string `json:"accessLevels"`
	// TODO Add support for claims
}

type destinationAttributes struct {
	IP         string            `json:"ip"`
	Port       string            `json:"port"`
	Labels     map[string]string `json:"labels"`
	Principal  string            `json:"principal"`
	RegionCode string            `json:"regionCode"`
}

// isValid checks that the log meets requirements
// See: https://cloud.google.com/logging/docs/audit/understanding-audit-logs#interpreting_the_sample_audit_log_entry
func isValid(log auditLog) error { _ = "STUB: not implemented"; return nil }

func handleResourceLocation(loc *resourceLocation, attr pcommon.Map) {
	_ = "STUB: not implemented"
	return
}

func handleStatus(s *status, attr pcommon.Map) { _ = "STUB: not implemented"; return }

func handleAuthenticationInfo(info *authenticationInfo, attr pcommon.Map) {
	_ = "STUB: not implemented"
	return
}

func handleAuthorizationInfo(info []authorizationInfo, attr pcommon.Map) {
	_ = "STUB: not implemented"
	return
}

func handlePolicyViolationInfo(info *policyViolationInfo, attr pcommon.Map) {
	_ = "STUB: not implemented"
	return
}

// Some logs have:
// 		policyViolationInfo: {
// 			orgPolicyViolationInfo: {}
// 		}
// We should ignore those cases.

func handleRequestMetadata(metadata *requestMetadata, attr pcommon.Map) error {
	_ = "STUB: not implemented"
	return nil
}

func ParsePayloadIntoAttributes(payload []byte, attr pcommon.Map) error {
	_ = "STUB: not implemented"
	return nil
}

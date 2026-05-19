// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Find more information about Armor Logs:
// https://docs.cloud.google.com/armor/docs/request-logging

package apploadbalancerlog // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding/googlecloudlogentryencodingextension/internal/apploadbalancerlog"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
)

const (
	// Security policy types
	gcpArmorSecurityPolicyEnforced     = "gcp.armor.security_policy.enforced"
	gcpArmorSecurityPolicyPreview      = "gcp.armor.security_policy.preview"
	gcpArmorSecurityPolicyEnforcedEdge = "gcp.armor.security_policy.enforced_edge"
	gcpArmorSecurityPolicyPreviewEdge  = "gcp.armor.security_policy.preview_edge"

	// Security policy base attributes
	// gcpArmorSecurityPolicyName holds the security policy rule that was enforced
	gcpArmorSecurityPolicyName = "name"
	// gcpArmorSecurityPolicyPriority holds a numerical priority of the matching rule in the security policy
	gcpArmorSecurityPolicyPriority = "priority"
	// gcpArmorSecurityPolicyConfiguredAction holds  the name of the configured action in the matching rule
	gcpArmorSecurityPolicyConfiguredAction = "configured_action"
	// gcpArmorSecurityPolicyOutcome holds the outcome of executing the configured action
	gcpArmorSecurityPolicyOutcome = "outcome"

	// Rate limit action attributes
	// gcpArmorRateLimitActionKey holds the rate limit key value (up to 36 bytes)
	gcpArmorRateLimitActionKey = "rate_limit.action.key"
	// gcpArmorRateLimitActionOutcome holds the outcome of the rate limit action
	gcpArmorRateLimitActionOutcome = "rate_limit.action.outcome"

	// Extended attributes
	// gcpArmorWAFRuleExpressionIDs holds the IDs of all preconfigured WAF rule expressions that triggered the rule
	gcpArmorWAFRuleExpressionIDs = "preconfigured.expr_ids"
	// gcpArmorThreatIntelligenceCategories holds information about the matched IP address lists from Google Threat Intelligence
	gcpArmorThreatIntelligenceCategories = "threat_intelligence.categories"
	// gcpArmorAddressGroupNames holds the names of the matched address groups
	gcpArmorAddressGroupNames = "address_group.names"

	// Enforced policy attributes
	// gcpArmorAdaptiveProtectionAutoDeployAlertID holds the alert ID of the events that Adaptive Protection detected
	gcpArmorAdaptiveProtectionAutoDeployAlertID = "adaptive_protection.auto_deploy.alert_id"

	// Security policy request data attributes
	// gcpArmorRecaptchaActionTokenScore holds the legitimacy score embedded in a of the reCAPTCHA action-token
	gcpArmorRecaptchaActionTokenScore = "gcp.armor.request_data.recaptcha_action_token.score" // #nosec G101 -- This is not a credential but an attribute name
	// gcpArmorRecaptchaSessionTokenScore holds the legitimacy score embedded in a of the reCAPTCHA session-token
	gcpArmorRecaptchaSessionTokenScore = "gcp.armor.request_data.recaptcha_session_token.score" // #nosec G101
	// gcpArmorUserIPInfoSource holds a field that is typically the header from which the user IP was resolved
	gcpArmorUserIPInfoSource = "gcp.armor.request_data.user_ip.source"
	// gcpArmorRemoteIPInfoAsn holds the five-digit autonomous system number (ASN) for the IP address
	gcpArmorRemoteIPInfoAsn = "gcp.armor.request_data.remote_ip.asn"
	// gcpArmorTLSJa4Fingerprint holds a JA4 TTL/SSL fingerprint if the client connects using HTTPS, HTTP/2, or HTTP/3
	gcpArmorTLSJa4Fingerprint = "tls.client.ja4"
)

type armorlog struct {
	SecurityPolicyRequestData *securityPolicyRequestData `json:"securityPolicyRequestData"`

	// Security policy fields: exactly one must be non-nil.
	EnforcedSecurityPolicy     *enforcedSecurityPolicy `json:"enforcedSecurityPolicy"`
	PreviewSecurityPolicy      *securityPolicyExtended `json:"previewSecurityPolicy"`
	EnforcedEdgeSecurityPolicy *securityPolicyBase     `json:"enforcedEdgeSecurityPolicy"`
	PreviewEdgeSecurityPolicy  *securityPolicyBase     `json:"previewEdgeSecurityPolicy"`
}

type securityPolicyRequestData struct {
	RecaptchaActionToken  *recaptchaToken `json:"recaptchaActionToken"`
	RecaptchaSessionToken *recaptchaToken `json:"recaptchaSessionToken"`
	UserIPInfo            *userIPInfo     `json:"userIpInfo"`
	RemoteIPInfo          *remoteIPInfo   `json:"remoteIpInfo"`
	TLSJa4Fingerprint     string          `json:"tlsJa4Fingerprint"`
	TLSJa3Fingerprint     string          `json:"tlsJa3Fingerprint"`
}

type recaptchaToken struct {
	Score float64 `json:"score"`
}

type userIPInfo struct {
	Source    string `json:"source"`
	IPAddress string `json:"ipAddress"`
}

type remoteIPInfo struct {
	IPAddress  string `json:"ipAddress"`
	RegionCode string `json:"regionCode"`
	ASN        *int64 `json:"asn"`
}

type securityPolicyBase struct {
	Name             string `json:"name"`
	Priority         *int64 `json:"priority"`
	ConfiguredAction string `json:"configuredAction"`
	Outcome          string `json:"outcome"`
}

type securityPolicyExtended struct {
	securityPolicyBase
	RateLimitAction      *rateLimitAction    `json:"rateLimitAction"`
	PreconfiguredExprIDs []string            `json:"preconfiguredExprIds"`
	ThreatIntelligence   *threatIntelligence `json:"threatIntelligence"`
	AddressGroup         *addressGroup       `json:"addressGroup"`
}

type enforcedSecurityPolicy struct {
	securityPolicyExtended
	AdaptiveProtection *adaptiveProtection `json:"adaptiveProtection"`
}

type rateLimitAction struct {
	Key     string `json:"key"`
	Outcome string `json:"outcome"`
}

type adaptiveProtection struct {
	AutoDeployAlertID string `json:"autoDeployAlertId"`
}

type threatIntelligence struct {
	Categories []string `json:"categories"`
}

type addressGroup struct {
	Names []string `json:"names"`
}

func handleRecaptchaTokens(data *securityPolicyRequestData, attr pcommon.Map) {
	_ = "STUB: not implemented"
	return
}

func handleUserIPInfo(info *userIPInfo, attr pcommon.Map) { _ = "STUB: not implemented"; return }

func handleRemoteIPInfo(info *remoteIPInfo, attr pcommon.Map) error {
	_ = "STUB: not implemented"
	return nil
}

func handleTLSFingerprints(data *securityPolicyRequestData, attr pcommon.Map) {
	_ = "STUB: not implemented"
	return
}

func handleSecurityPolicyRequestData(data *securityPolicyRequestData, attr pcommon.Map) error {
	_ = "STUB: not implemented"
	return nil
}

func handleRateLimitAction(rl *rateLimitAction, attr pcommon.Map) {
	_ = "STUB: not implemented"
	return
}

func handleAddressGroup(ag *addressGroup, attr pcommon.Map) { _ = "STUB: not implemented"; return }

func handleThreatIntelligence(ti *threatIntelligence, attr pcommon.Map) {
	_ = "STUB: not implemented"
	return
}

func handleSecurityPolicyBase(sp *securityPolicyBase, attr pcommon.Map) {
	_ = "STUB: not implemented"
	return
}

func handleSecurityPolicyExtended(sp *securityPolicyExtended, attr pcommon.Map) {
	_ = "STUB: not implemented"
	return
}

func handleEnforcedSecurityPolicy(sp *enforcedSecurityPolicy, attr pcommon.Map) {
	_ = "STUB: not implemented"
	return
}

func handleSecurityPolicies(armorlog *armorlog, attr pcommon.Map) {
	_ = "STUB: not implemented"
	return
}

func handleArmorLogAttributes(armorlog *armorlog, attr pcommon.Map) error {
	_ = "STUB: not implemented"
	return nil
}

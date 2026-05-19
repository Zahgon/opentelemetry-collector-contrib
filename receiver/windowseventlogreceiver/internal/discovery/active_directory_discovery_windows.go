// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build windows

package discovery // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/windowseventlogreceiver/internal/discovery"

import (
	"errors"
	"fmt"

	"github.com/alexbrainman/sspi"
	"github.com/alexbrainman/sspi/ntlm"
	"github.com/go-ldap/ldap/v3"
	"go.uber.org/zap"

	stanza "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/input/windows"
)

// sspINegotiator implements ldap.NTLMNegotiator using current user's SSPI credentials.
// This mirrors C++ ADsOpenObject(..., ADS_SECURE_AUTHENTICATION) with null username/password.
type sspINegotiator struct {
	creds *sspi.Credentials   // current user creds acquired from SSPI
	ctx   *ntlm.ClientContext // NTLM client context for challenge/response
}

var _ ldap.NTLMNegotiator = (*sspINegotiator)(nil)

// Acquire current logged-on user credentials no username/password needed
func (n *sspINegotiator) Negotiate(_, _ string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create NTLM client context and generate Type 1 (Negotiate) message

func (n *sspINegotiator) ChallengeResponse(challenge []byte, _, _ string) ([]byte, error) {
	_ = "STUB: not implemented"
	// Process Type 2 (Challenge) from server, produce Type 3 (Authenticate) message
	return nil, nil
}

func (n *sspINegotiator) Release() { _ = "STUB: not implemented"; return }

// getLDAPDomainPath discovers the root domain path of the Active Directory service.
// It first tries querying the LDAP Root DSE, then falls back to the Windows API.
// Returns a path like "LDAP://DC=example,DC=com".
func getLDAPDomainPath(logger *zap.Logger) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Primary: query Root DSE with the current joined domain as the LDAP server

// Fallback: current Joined Domain

// getRootLDAPDomainPath connects to the current machine joined DC and reads
// the defaultNamingContext attribute.
func getRootLDAPDomainPath(domain string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Base DN: empty string = Root DSE
// Only the root entry itself

// Size limit of 1 since we only expect one Root DSE entry
// 10s time limit to avoid hanging if something is wrong with the server
// attrs only = false

// namingContext is already in DN format, e.g. "DC=example,DC=com"

// getLDAPDomainPathFromWindowsAPI uses GetComputerNameEx to get the DNS domain
// name and converts it to an LDAP path.
func getCurrentMachineJoinedDomain() (string, error) {
	_ = "STUB: not implemented"
	// First call to get required buffer size
	return "", nil
}

// Decode UTF-16 to string, trimming null terminator

func discoverDomainControllersForJoinedDomain(logger *zap.Logger) (string, []string, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

func getDomainControllersForDomain(domain string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Size limit, keeping 1000 for safety, though we expect far fewer DCs
// Time limit of 30s to avoid hanging if something is wrong with the server
// Types only
// LDAP filter for DCs
// Attributes to retrieve

// domainToDN converts a domain name to an LDAP Distinguished Name.
// Example: "example.com" -> "DC=example,DC=com"
func domainToDN(domain string) string { _ = "STUB: not implemented"; return "" }

func dnToHostname(dn string) string { _ = "STUB: not implemented"; return "" }

// GetJoinedDomainControllersRemoteConfig is a variable holding the function that discovers domain controllers
// for the machine's joined domain and returns a RemoteConfig slice for each discovered controller.
// It is a variable to allow mocking in tests.
var GetJoinedDomainControllersRemoteConfig = func(logger *zap.Logger, username, password string) ([]stanza.RemoteConfig, error) {
	var domainControllerConfigs []stanza.RemoteConfig
	domain, domainControllers, err := discoverDomainControllersForJoinedDomain(logger)
	if err != nil {
		return nil, fmt.Errorf("failed to discover domain controllers: %w", err)
	}
	if len(domainControllers) == 0 {
		return nil, errors.New("no domain controllers found during discovery")
	}
	for _, dc := range domainControllers {
		config := stanza.RemoteConfig{Server: dc, Username: username, Password: password, Domain: domain}
		domainControllerConfigs = append(domainControllerConfigs, config)
	}
	logger.Info("Discovered domain controllers", zap.Int("count", len(domainControllers)))
	return domainControllerConfigs, nil
}

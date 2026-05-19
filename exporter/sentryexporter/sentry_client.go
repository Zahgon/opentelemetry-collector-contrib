// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package sentryexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/sentryexporter"

import (
	"context"
	"net/http"

	"github.com/getsentry/sentry-go"
)

// sentryAPIClient defines the interface for interacting with Sentry API.
type sentryAPIClient interface {
	GetAllProjects(ctx context.Context, orgSlug string) ([]projectInfo, error)
	GetProjectKeys(ctx context.Context, orgSlug, projectSlug string) ([]projectKey, error)
	GetOrgProjectKeys(ctx context.Context, orgSlug string) ([]projectKey, error)
	GetOTLPEndpoints(ctx context.Context, orgSlug, projectSlug string) (*otlpEndpoints, error)
	CreateProject(ctx context.Context, orgSlug, teamSlug, projectSlug, projectName, platform string) (*projectInfo, error)
}

// sentryClient handles communication with the Sentry API.
type sentryClient struct {
	baseURL   string
	authToken string
	client    *http.Client
}

// newSentryClient is used to override the sentry client factory. While running tests we need to mock the http transport, so that
// we don't open real sockets.
var newSentryClient = func(baseURL, authToken string, httpClient *http.Client) sentryAPIClient {
	return newSentryClientImpl(baseURL, authToken, httpClient)
}

// projectKey represents a Sentry project key.
type projectKey struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	Public    string   `json:"public"`
	Secret    string   `json:"secret"`
	ProjectID int      `json:"projectId"`
	IsActive  bool     `json:"isActive"`
	DSN       dsnField `json:"dsn"`
}

// dsnField represents the DSN field from API response.
type dsnField struct {
	Public string `json:"public"`
}

// ParsePublicDSN parses the public DSN string into a sentry.Dsn.
func (d *dsnField) ParsePublicDSN() (*sentry.Dsn, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// teamInfo represents a Sentry team.
}

type teamInfo struct {
	ID   string `json:"id"`
	Slug string `json:"slug"`
	Name string `json:"name"`
}

// projectInfo represents a Sentry project.
type projectInfo struct {
	ID    string     `json:"id"`
	Slug  string     `json:"slug"`
	Name  string     `json:"name"`
	Team  teamInfo   `json:"team"`
	Teams []teamInfo `json:"teams"`
}

// otlpEndpoints contains the OTLP endpoint URLs for a project.
type otlpEndpoints struct {
	TracesURL  string
	LogsURL    string
	PublicKey  string
	AuthHeader string
}

func parseProjectID(id string) int { _ = "STUB: not implemented"; return 0 }

// parseNextCursor extracts the cursor for the next page from a Sentry Link header.
func parseNextCursor(header http.Header) (cursor string, hasMore, found bool) {
	_ = "STUB: not implemented"
	return "", false, false
}

func newSentryClientImpl(baseURL, authToken string, httpClient *http.Client) *sentryClient {
	_ = "STUB: not implemented"
	return nil
}

// GetAllProjects fetches all projects for a given organization.
func (c *sentryClient) GetAllProjects(ctx context.Context, orgSlug string) ([]projectInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetProjectKeys fetches the project keys for a given organization and project.
func (c *sentryClient) GetProjectKeys(ctx context.Context, orgSlug, projectSlug string) ([]projectKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetOrgProjectKeys fetches all project keys for an organization.
func (c *sentryClient) GetOrgProjectKeys(ctx context.Context, orgSlug string) ([]projectKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateProject creates a new project in the given organization and team.
func (c *sentryClient) CreateProject(ctx context.Context, orgSlug, teamSlug, projectSlug, projectName, platform string) (*projectInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetOTLPEndpoints returns the OTLP endpoint URLs for a project.
func (c *sentryClient) GetOTLPEndpoints(ctx context.Context, orgSlug, projectSlug string) (*otlpEndpoints, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

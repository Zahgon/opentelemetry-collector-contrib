// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package splunkenterprisereceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/splunkenterprisereceiver"

import (
	"context"
	"errors"
	"net/http"
	"net/url"

	"go.opentelemetry.io/collector/component"
)

// Indexer type "enum". Included in context sent from scraper functions
const (
	typeIdx = "IDX"
	typeSh  = "SH"
	typeCm  = "CM"
)

var (
	errCtxMissingEndpointType = errors.New("context was passed without the endpoint type included")
	errEndpointTypeNotFound   = errors.New("requested client is not configured and could not be found in splunkEntClient")
)

// Wrapper around splunkClientMap to avoid awkward reference/dereference stuff that arises when using maps in golang
type splunkEntClient struct {
	clients splunkClientMap
}

func (c *splunkEntClient) newClientNotFoundError(eptType, apiEndpoint string) error {
	_ = "STUB: not implemented"
	return nil
}

// Type wrapper for accessing context value
type endpointType string

// The splunkEntClient is made up of a number of splunkClients defined for each configured endpoint
type splunkClientMap map[string]splunkClient

// The client does not carry the endpoint that is configured with it and golang does not support mixed
// type arrays so this struct contains the pair: the client configured for the endpoint and the endpoint
// itself
type splunkClient struct {
	client   *http.Client
	endpoint *url.URL
}

func newSplunkEntClient(ctx context.Context, cfg *Config, h component.Host, s component.TelemetrySettings) (*splunkEntClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// if the endpoint is defined, put it in the endpoints map for later use
// we already checked that url.Parse does not fail in cfg.Validate()

// For running ad hoc searches only
func (c *splunkEntClient) createRequest(eptType string, sr *searchResponse) (req *http.Request, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Running searches via Splunk's REST API is a two step process: First you submit the job to run
// this returns a jobid which is then used in the second part to retrieve the search results

// reader for the response data

// return the build request, ready to be run by makeRequest

// forms an *http.Request for use with Splunk built-in API's (like introspection).
func (c *splunkEntClient) createAPIRequest(eptType, apiEndpoint string) (req *http.Request, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Perform a request.
func (c *splunkEntClient) makeRequest(req *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	// get endpoint type from the context
	return nil, nil
}

// Check if the splunkEntClient contains a configured endpoint for the type of scraper
// Returns true if an entry exists, false if not.
func (c *splunkEntClient) isConfigured(v string) bool { _ = "STUB: not implemented"; return false }

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package sigv4authextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/sigv4authextension"

import (
	"errors"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	sigv4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"go.uber.org/zap"
)

var errNilRequest = errors.New("sigv4: unable to sign nil *http.Request")

// signingRoundTripper is a custom RoundTripper that performs AWS Sigv4.
type signingRoundTripper struct {
	transport     http.RoundTripper
	signer        *sigv4.Signer
	region        string
	service       string
	credsProvider *aws.CredentialsProvider
	awsSDKInfo    string
	logger        *zap.Logger
}

// RoundTrip() executes a single HTTP transaction and returns an HTTP response, signing
// the request with Sigv4.
func (si *signingRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Send the request

func (si *signingRoundTripper) signRequest(req *http.Request) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Clone request to ensure thread safety.

// Add the runtime information to the User-Agent header of the request

// Use user provided service/region if specified, use inferred service/region if not, then sign the request

// hashPayload creates a SHA256 hash of the request body
func hashPayload(req *http.Request) (string, error) {
	_ = "STUB: not implemented"
	return "",

		// hash of an empty payload to use if there is no request body
		nil
}

// Hash the request body

// inferServiceAndRegion attempts to infer a service
// and a region from an http.request, and returns either an empty
// string for both or a valid value for both.
func (si *signingRoundTripper) inferServiceAndRegion(r *http.Request) (service, region string) {
	_ = "STUB: not implemented"
	return "", ""
}

func extractServiceAndRegion(service, region, host, defaultService string) (string, string) {
	_ = "STUB: not implemented"
	return "", ""
}

// cloneRequest() is a helper function that makes a shallow copy of the request and a
// deep copy of the header, for thread safety purposes.
func cloneRequest(r *http.Request) *http.Request {
	_ = "STUB: not implemented"
	// shallow copy of the struct
	return nil
}

// deep copy of the Header

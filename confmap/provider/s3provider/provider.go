// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:generate make mdatagen

package s3provider // import "github.com/open-telemetry/opentelemetry-collector-contrib/confmap/provider/s3provider"

import (
	"context"
	"regexp"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"go.opentelemetry.io/collector/confmap"
)

const (
	schemeName = "s3"
	// Pattern for an AWS S3 virtual-hosted-style uri
	s3AWSPattern = `^s3:\/\/([a-z0-9\.\-]{3,63})\.s3(?:-fips)?(?:\.dualstack)?\.([a-z0-9\-]+)\.(api\.amazonwebservices\.com\.cn|api\.amazonwebservices\.eu|api\.cloud-aws\.adc-e\.uk|api\.aws\.hci\.ic\.gov|amazonaws\.com\.cn|cloud\.adc-e\.uk|api\.aws\.scloud|api\.aws\.ic\.gov|csp\.hci\.ic\.gov|sc2s\.sgov\.gov|amazonaws\.com|amazonaws\.eu|c2s\.ic\.gov|api\.aws)\/.`
)

var s3AWSRegexp = regexp.MustCompile(s3AWSPattern)

type s3Client interface {
	GetObject(context.Context, *s3.GetObjectInput, ...func(*s3.Options)) (*s3.GetObjectOutput, error)
}

type provider struct {
	client s3Client
}

// NewFactory returns a new confmap.ProviderFactory that creates a confmap.Provider
// which reads configuration from a file obtained from an s3 bucket.
//
// This Provider supports the "s3" scheme with two URI formats:
//
// AWS virtual-hosted-style (standard Amazon S3):
//
//	s3://[BUCKET].s3.[REGION].amazonaws.com/[KEY]
//	s3://doc-example-bucket.s3.us-west-2.amazonaws.com/config.yaml
//
// S3-compatible path-style (for MinIO, DigitalOcean Spaces, and other S3-compatible services):
//
//	s3://[ENDPOINT_HOST]/[BUCKET]/[KEY]?region=[REGION]
//	s3://minio.example.com/my-bucket/config.yaml
func NewFactory() confmap.ProviderFactory {
	_ = "STUB: not implemented"
	return *new(confmap.ProviderFactory)
}

func newWithSettings(confmap.ProviderSettings) confmap.Provider {
	_ = "STUB: not implemented"
	return *new(confmap.Provider)
}

func (fmp *provider) Retrieve(ctx context.Context, uri string, _ confmap.WatcherFunc) (*confmap.Retrieved, error) {
	_ = "STUB: not implemented"
	// Split the uri and get [BUCKET], [REGION], [KEY], and optional [ENDPOINT]
	return nil, nil
}

// s3 downloading

// read config from response body

func (*provider) Scheme() string { _ = "STUB: not implemented"; return "" }

func (*provider) Shutdown(context.Context) error {
	_ = "STUB: not implemented"

	// s3URISplit splits the s3 uri and returns [BUCKET], [REGION], [KEY], and optional [ENDPOINT].
	//
	// Two URI formats are supported:
	//
	//  1. AWS virtual-hosted-style (host contains "amazonaws.com"):
	//     s3://[BUCKET].s3.[REGION].amazonaws.com/[KEY]
	//
	//  2. S3-compatible path-style (any other host):
	//     s3://[ENDPOINT_HOST]/[BUCKET]/[KEY]?region=[REGION]
	//     The host is used as the endpoint; region is optional.
	return nil
}

func s3URISplit(uri string) (bucket, region, key, endpoint string, err error) {
	_ = "STUB: not implemented"
	// parse the uri as [scheme:][//[userinfo@]host][/]path[?query][#fragment]
	return "", "", "", "", nil
}

// S3-compatible path-style: s3://endpoint-host/bucket/key?region=...
// The host is the endpoint; the path provides bucket and key.

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package elasticsearchexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/elasticsearchexporter"

import (
	"context"
	"net/http"
	"net/url"
	"sync/atomic"
	"time"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"go.opentelemetry.io/collector/component"
	"go.uber.org/zap"
)

type attemptCounterKey struct{}

// attemptCounter tracks the number of round-trip attempts for a single
// _bulk request call in the Elasticsearch client.
type attemptCounter struct {
	value atomic.Int64
}

func (c *attemptCounter) Attempts() int { _ = "STUB: not implemented"; return 0 }
func (c *attemptCounter) Retries() int  { _ = "STUB: not implemented"; return 0 }

// newAttemptContext returns a context carrying a fresh attemptCounter,
// along with the counter itself so the caller can inspect it after.
func newAttemptContext(ctx context.Context) (context.Context, *attemptCounter) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

// countRetriesInterceptor returns an interceptor that increments the
// attemptCounter stored in the request context on every round-trip.
func countRetriesInterceptor() elastictransport.InterceptorFunc {
	_ = "STUB: not implemented"
	return *new(elastictransport.InterceptorFunc)
}

// clientLogger implements the estransport.Logger interface
// that is required by the Elasticsearch client for logging.
type clientLogger struct {
	*zap.Logger
	logRequestBody  bool
	logResponseBody bool
	componentHost   component.Host
}

// LogRoundTrip should not modify the request or response, except for consuming and closing the body.
// Implementations have to check for nil values in request and response.
func (cl *clientLogger) LogRoundTrip(requ *http.Request, resp *http.Response, clientErr error, _ time.Time, dur time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// Success

// RequestBodyEnabled makes the client pass a copy of request body to the logger.
func (cl *clientLogger) RequestBodyEnabled() bool { _ = "STUB: not implemented"; return false }

// ResponseBodyEnabled makes the client pass a copy of response body to the logger.
func (cl *clientLogger) ResponseBodyEnabled() bool { _ = "STUB: not implemented"; return false }

const (
	unknownProduct = "the client noticed that the server is not Elasticsearch and we do not support this unknown product"
	defaultURL     = "http://localhost:9200"
)

// genuineCheckHeader validates the presence of the X-Elastic-Product header
func genuineCheckHeader(header http.Header) error { _ = "STUB: not implemented"; return nil }

type esClient struct {
	transport           elastictransport.Interface
	productCheckSuccess atomic.Bool
}

func (e *esClient) Perform(req *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *esClient) doProductCheck(f func() error) error { _ = "STUB: not implemented"; return nil }

// newElasticsearchClient returns a new elastictransport.Interface.
func newElasticsearchClient(
	ctx context.Context,
	config *Config,
	host component.Host,
	telemetry component.TelemetrySettings,
	userAgent string,
) (elastictransport.Interface, error) {
	_ = "STUB: not implemented"
	return *new(elastictransport.Interface), nil
}

// endpoints converts Config.Endpoints, Config.CloudID,
// and Config.ClientConfig.Endpoint to a list of addresses.

// Convert addresses to URLs

// Create transport configuration matching elasticsearch.newTransport structure

// TODO
// TODO

/* captureSearchBody */

// Handle node discovery on start, matching elasticsearch.NewClient behavior

// addrsToURLs creates a list of url.URL structures from url list.
func addrsToURLs(addrs []string) ([]*url.URL, error) { _ = "STUB: not implemented"; return nil, nil }

// createElasticsearchBackoffFunc creates an exponential backoff with equal jitter.
func createElasticsearchBackoffFunc(config *RetrySettings) func(int) time.Duration {
	_ = "STUB: not implemented"
	return nil
}

// config.InitialInterval * 2 ^ (attempts - 1)
// guard against overflow

func httpRecoverableErrorStatus(statusCode int) bool {
	_ = "STUB: not implemented"
	// Elasticsearch uses 409 conflict to report duplicates, which aren't really
	// an error state, so those return false (but if we were already in an error
	// state, we will still wait until we get an actual 200 OK before changing
	// our state back).
	return false
}

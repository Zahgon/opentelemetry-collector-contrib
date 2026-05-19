// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0
// Originally copied from https://github.com/signalfx/signalfx-agent/blob/fbc24b0fdd3884bd0bbfbd69fe3c83f49d4c0b77/pkg/apm/correlations/client.go

package correlations // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/signalfxexporter/internal/apm/correlations"

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/signalfxexporter/internal/apm/log"
	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/signalfxexporter/internal/apm/requests"
)

var (
	ErrChFull           = errors.New("request channel full")
	errRetryChFull      = errors.New("retry channel full")
	errMaxAttempts      = errors.New("maximum attempts exceeded")
	errRequestCancelled = errors.New("request cancelled")
)

// ErrMaxEntries is an error returned when the correlation endpoint returns a 418 http status
// code indicating that the set of services or environments is too large to add another value
type ErrMaxEntries struct {
	MaxEntries int64 `json:"max,omitempty"`
}

func (m *ErrMaxEntries) Error() string { _ = "STUB: not implemented"; return "" }

var _ error = (*ErrMaxEntries)(nil)

// CorrelationClient is an interface for correlations.Client
type CorrelationClient interface {
	Correlate(*Correlation, CorrelateCB)
	Delete(*Correlation, SuccessfulDeleteCB)
	Get(dimName, dimValue string, cb SuccessfulGetCB)
	Start()
	Shutdown()
}

type request struct {
	*Correlation
	ctx       context.Context
	cancel    context.CancelFunc
	operation string
	callback  func(body []byte, statuscode int, err error)
	sendAt    time.Time
}

// Client is a client for making dimensional correlations
type Client struct {
	sync.RWMutex
	log           log.Logger
	ctx           context.Context
	wg            sync.WaitGroup
	Token         string
	APIURL        *url.URL
	client        *http.Client
	requestSender *requests.ReqSender
	requestChan   chan *request
	retryChan     chan *request
	dedup         *deduplicator

	// For easier unit testing
	now        func() time.Time
	logUpdates bool

	retryDelay  time.Duration
	maxAttempts uint32

	TotalClientError4xxResponses int64
	TotalRetriedUpdates          int64
	TotalInvalidDimensions       int64
	dedupCleanupInterval         time.Duration
}

// Config defines configuration for correlation settings.
type Config struct {
	MaxRequests     uint          `mapstructure:"max_requests"`
	MaxBuffered     uint          `mapstructure:"max_buffered"`
	MaxRetries      uint          `mapstructure:"max_retries"`
	LogUpdates      bool          `mapstructure:"log_updates"`
	RetryDelay      time.Duration `mapstructure:"retry_delay"`
	CleanupInterval time.Duration `mapstructure:"cleanup_interval"`
}

// ClientConfig for correlation client.
type ClientConfig struct {
	Config
	AccessToken string
	URL         *url.URL
}

// NewCorrelationClient returns a new Client
func NewCorrelationClient(ctx context.Context, log log.Logger, client *http.Client, conf ClientConfig) (CorrelationClient, error) {
	_ = "STUB: not implemented"
	return *new(CorrelationClient), nil
}

func (cc *Client) putRequestOnChan(r *request) error {
	_ = "STUB: not implemented"
	// prevent requests against empty dimension names and values
	return nil
}

// logging this as debug because this means there's no actual dimension to correlate with
// and because this isn't being taken off on the request sender and subject to retries, this could
// potentially spam the logs

func (cc *Client) putRequestOnRetryChan(r *request) error {
	_ = "STUB: not implemented"
	// handle request counter
	return nil
}

// set the time to retry

// CorrelateCB is a call back invoked with Correlate requests
// it is not invoked if the request is deduplicated, cancelled, or the client context is cancelled
type CorrelateCB func(cor *Correlation, err error)

// Correlate
func (cc *Client) Correlate(cor *Correlation, cb CorrelateCB) { _ = "STUB: not implemented"; return }

// SuccessfulDeleteCB is a call back that is only invoked on successful Deletion operations
type SuccessfulDeleteCB func(cor *Correlation)

// Delete removes a correlation
func (cc *Client) Delete(cor *Correlation, callback SuccessfulDeleteCB) {
	_ = "STUB: not implemented"
	return
}

// SuccessfulGetCB
type SuccessfulGetCB func(map[string][]string)

// Get
func (cc *Client) Get(dimName, dimValue string, callback SuccessfulGetCB) {
	_ = "STUB: not implemented"
	return
}

// only log this as debug because we do a blanket fetch of correlations on the backend
// and if the backend fails to find anything this isn't really an error for us

func (cc *Client) makeRequest(r *request) { _ = "STUB: not implemented"; return }

// build endpoint url

// TODO: pool the reader

// logging this as debug because this means there's something fundamentally wrong with the request
// and because this isn't being taken off on the request sender and subject to retries, this could
// potentially spam the logs long term.  This would be a really good candidate for a throttled error logger

// retry if the http status code is not 4XX. A 4xx or http client error implies
// an error that is not going to be remedied by retrying.

// The retry (for non 400 errors) is meant to provide some measure of robustness against
// temporary API failures.  If the API is down for significant
// periods of time, correlation updates will probably eventually back
// up beyond conf.MaxBuffered and start dropping.

// invoke the callback

// cancel the request context

// close the request context

// This will block if we don't have enough requests

// routines
// processChan processes incoming requests, drops duplicates, and cancels conflicting requests
func (cc *Client) processChan() { _ = "STUB: not implemented"; return }

// processRetryChan is a routine that drains the retry channel and waits until the appropriate time to retry the request
func (cc *Client) processRetryChan() { _ = "STUB: not implemented"; return }

// client is shutdown

// wait and resend the request

// request is cancelled

// client is shutdown

// Start the client's processing queue
func (cc *Client) Start() { _ = "STUB: not implemented"; return }

// Shutdown the client. This will block until the context's cancel
// function is complete.
func (cc *Client) Shutdown() { _ = "STUB: not implemented"; return }

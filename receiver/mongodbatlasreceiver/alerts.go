// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package mongodbatlasreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mongodbatlasreceiver"

import (
	"context" // #nosec G505 -- SHA1 is the algorithm mongodbatlas uses, it must be used to calculate the HMAC signature
	"net/http"
	"sync"
	"time"

	"go.mongodb.org/atlas/mongodbatlas"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/config/configretry"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/extension/xextension/storage"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	rcvr "go.opentelemetry.io/collector/receiver"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mongodbatlasreceiver/internal"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mongodbatlasreceiver/internal/model"
)

// maxContentLength is the maximum payload size we will accept from incoming requests.
// Requests are generally ~1000 bytes, so we overshoot that by an order of magnitude.
// This is to protect from overly large requests.
const (
	maxContentLength    int64  = 16384
	signatureHeaderName string = "X-MMS-Signature"

	alertModeListen = "listen"
	alertModePoll   = "poll"
	alertCacheKey   = "last_recorded_alert"

	defaultAlertsPollInterval = 5 * time.Minute
	// defaults were based off API docs https://www.mongodb.com/docs/atlas/reference/api/alerts-get-all-alerts/
	defaultAlertsPageSize = 100
	defaultAlertsMaxPages = 10
)

type alertsClient interface {
	GetProject(ctx context.Context, groupID string) (*mongodbatlas.Project, error)
	GetAlerts(ctx context.Context, groupID string, opts *internal.AlertPollOptions) ([]mongodbatlas.Alert, bool, error)
}

type alertsReceiver struct {
	secret       string
	mode         string
	serverConfig *confighttp.ServerConfig
	consumer     consumer.Logs
	wg           *sync.WaitGroup
	listenClose  func(ctx context.Context) error

	// only relevant in `poll` mode
	projects          []*ProjectConfig
	client            alertsClient
	baseURL           string
	privateKey        string
	publicKey         string
	backoffConfig     configretry.BackOffConfig
	pollInterval      time.Duration
	record            *alertRecord
	pageSize          int64
	maxPages          int64
	doneChan          chan bool
	storageClient     storage.Client
	telemetrySettings component.TelemetrySettings
}

func newAlertsReceiver(params rcvr.Settings, baseConfig *Config, consumer consumer.Logs) (*alertsReceiver, error) {
	_ = "STUB: not implemented"
	return nil,

		// Validate TLS
		nil
}

func (a *alertsReceiver) Start(ctx context.Context, host component.Host, storageClient storage.Client) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *alertsReceiver) startPolling(ctx context.Context, storageClient storage.Client) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *alertsReceiver) retrieveAndProcessAlerts(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *alertsReceiver) pollAndProcess(ctx context.Context, pc *ProjectConfig, project *mongodbatlas.Project) {
	_ = "STUB: not implemented"
	return
}

func (a *alertsReceiver) startListening(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *alertsReceiver) handleRequest(rw http.ResponseWriter, req *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (a *alertsReceiver) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (a *alertsReceiver) shutdownListener(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *alertsReceiver) shutdownPoller(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *alertsReceiver) convertAlerts(now pcommon.Timestamp, alerts []*mongodbatlas.Alert, project *mongodbatlas.Project) (plog.Logs, error) {
	_ = "STUB: not implemented"
	return *new(plog.Logs), nil
}

// this could be fairly expensive to do, expecting not too many issues unless there are a ton
// of unrecognized alerts to process.

// These attributes are always present

// These attributes are optional and may not be present, depending on the alert type.

// Only present for HOST, HOST_METRIC, and REPLICA_SET alerts

func verifyHMACSignature(secret string, payload []byte, signatureHeader string) error {
	_ = "STUB: not implemented"
	return nil
}

func payloadToLogs(now time.Time, payload []byte) (plog.Logs, error) {
	_ = "STUB: not implemented"
	return *new(plog.Logs), nil
}

// These attributes are always present

// These attributes are optional and may not be present, depending on the alert type.

// alertRecord wraps a sync Map so it is goroutine safe as well as
// can have custom marshaling
type alertRecord struct {
	sync.Mutex
	LastRecordedTime *time.Time `mapstructure:"last_recorded"`
}

func (a *alertRecord) SetLastRecorded(lastUpdated *time.Time) { _ = "STUB: not implemented"; return }

func (a *alertsReceiver) syncPersistence(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *alertsReceiver) writeCheckpoint(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *alertsReceiver) applyFilters(pConf *ProjectConfig, alerts []mongodbatlas.Alert) []*mongodbatlas.Alert {
	_ = "STUB: not implemented"
	return nil
}

// we need to maintain two timestamps in order to not conflict while iterating

// already processed if the updated time was before or equal to the last recorded

func timestampFromAlert(a model.Alert) pcommon.Timestamp {
	_ = "STUB: not implemented"
	return *new(pcommon.Timestamp)
}

// severityFromAlert maps the alert to a severity number.
// Currently, it just maps "OPEN" alerts to WARN, and everything else to INFO.
func severityFromAlert(a model.Alert) plog.SeverityNumber {
	_ = "STUB: not implemented"
	// Status is defined here: https://www.mongodb.com/docs/atlas/reference/api/alerts-get-alert/#response-elements
	// It may also be "INFORMATIONAL" for single-fire alerts (events)
	return *new(plog.SeverityNumber)
}

// severityFromAPIAlert is a workaround for shared types between the API and the model
func severityFromAPIAlert(a string) plog.SeverityNumber {
	_ = "STUB: not implemented"
	return *new(plog.SeverityNumber)
}

func putStringToMapNotNil(m pcommon.Map, k string, v *string) { _ = "STUB: not implemented"; return }

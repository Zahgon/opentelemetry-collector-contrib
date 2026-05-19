// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package mongodbatlasreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mongodbatlasreceiver"

import (
	"context"
	"sync"
	"time"

	"go.mongodb.org/atlas/mongodbatlas"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/extension/xextension/storage"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	rcvr "go.opentelemetry.io/collector/receiver"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mongodbatlasreceiver/internal"
)

const (
	eventStorageKey       = "last_recorded_event"
	defaultEventsMaxPages = 25
	defaultEventsPageSize = 100
	defaultPollInterval   = time.Minute
)

type eventsClient interface {
	GetProject(ctx context.Context, groupID string) (*mongodbatlas.Project, error)
	GetProjectEvents(ctx context.Context, groupID string, opts *internal.GetEventsOptions) (ret []*mongodbatlas.Event, nextPage bool, err error)
	GetOrganization(ctx context.Context, orgID string) (*mongodbatlas.Organization, error)
	GetOrganizationEvents(ctx context.Context, orgID string, opts *internal.GetEventsOptions) (ret []*mongodbatlas.Event, nextPage bool, err error)
	Shutdown() error
}

type eventsReceiver struct {
	client        eventsClient
	logger        *zap.Logger
	storageClient storage.Client
	cfg           *Config
	consumer      consumer.Logs

	maxPages     int
	pageSize     int
	pollInterval time.Duration
	wg           *sync.WaitGroup
	record       *eventRecord // this record is used for checkpointing last processed events
	cancel       context.CancelFunc
}

type eventRecord struct {
	NextStartTime *time.Time `mapstructure:"next_start_time"`
}

func newEventsReceiver(settings rcvr.Settings, c *Config, consumer consumer.Logs) (*eventsReceiver, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (er *eventsReceiver) Start(ctx context.Context, _ component.Host, storageClient storage.Client) error {
	_ = "STUB: not implemented"
	return nil
}

func (er *eventsReceiver) Shutdown(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (er *eventsReceiver) startPolling(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (er *eventsReceiver) pollEvents(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (er *eventsReceiver) pollProject(ctx context.Context, project *mongodbatlas.Project, p *ProjectConfig, startTime, now time.Time) {
	_ = "STUB: not implemented"
	return
}

func (er *eventsReceiver) pollOrg(ctx context.Context, org *mongodbatlas.Organization, p *OrgConfig, startTime, now time.Time) {
	_ = "STUB: not implemented"
	return
}

func (er *eventsReceiver) transformProjectEvents(now pcommon.Timestamp, events []*mongodbatlas.Event, p *mongodbatlas.Project) plog.Logs {
	_ = "STUB: not implemented"
	return *new(plog.Logs)
}

func (er *eventsReceiver) transformOrgEvents(now pcommon.Timestamp, events []*mongodbatlas.Event, o *mongodbatlas.Organization) plog.Logs {
	_ = "STUB: not implemented"
	return *new(plog.Logs)
}

func (er *eventsReceiver) transformEvents(now pcommon.Timestamp, events []*mongodbatlas.Event, resourceLogs *plog.ResourceLogs) {
	_ = "STUB: not implemented"
	return
}

// ISO-8601 formatted

// always present attributes

func (er *eventsReceiver) checkpoint(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (er *eventsReceiver) loadCheckpoint(ctx context.Context) { _ = "STUB: not implemented"; return }

func parseOptionalAttributes(m *pcommon.Map, event *mongodbatlas.Event) {
	_ = "STUB: not implemented"
	return
}

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package googlecloudpubsubpushreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudpubsubpushreceiver"

import (
	"context"
	"io"
	"net/http"
	"sync"
	"time"

	"cloud.google.com/go/storage"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudpubsubpushreceiver/internal/metadata"
)

const (
	bucketIDKey         = "bucketId"
	objectIDKey         = "objectId"
	eventTypeKey        = "eventType"
	eventObjectFinalize = "OBJECT_FINALIZE"

	bucketMetadataKey          = "bucket"
	objectMetadataKey          = "object"
	subscriptionMetadataKey    = "subscription"
	messageIDMetadataKey       = "message_id"
	deliveryAttemptMetadataKey = "delivery_attempt"

	bucketNameAttr = "gcp.gcs.bucket.name"
)

type pubSubPushReceiver struct {
	cfg              *Config
	settings         receiver.Settings
	storageClient    *storage.Client
	telemetryBuilder *metadata.TelemetryBuilder

	server     *http.Server
	shutdownWG sync.WaitGroup

	nextLogs consumer.Logs
}

var _ receiver.Logs = (*pubSubPushReceiver)(nil)

func newPubSubPushReceiver(
	cfg *Config,
	set receiver.Settings,
	nextLogs consumer.Logs,
) (*pubSubPushReceiver, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func addHandlerFunc[T any](
	tb *metadata.TelemetryBuilder,
	mux *http.ServeMux,
	endpoint string,
	unmarshal func([]byte) (T, error),
	consume func(context.Context, T) error,
	storageClient *storage.Client,
	includeMetadata bool,
	logger *zap.Logger,
) {
	_ = "STUB: not implemented"
	return
}

// Pub/Sub retries everything that is not a valid response. A valid response
// has the following HTTP codes: [102, 200, 201, 202, 204].
// You can verify this in the official documentation. For this, refer to
// https://cloud.google.com/pubsub/docs/push.
//
// This becomes an issue for permanent errors, because it means that if we
// don't return one of those codes, Pub/Sub will keep retrying. This would
// only stop retrying if:
// 1. We add event arc advanced after Pub/Sub. For this, we need to set the
// correct HTTP code for permanent/transient errors. You can refer to
// the official documentation for this. See:
// https://docs.cloud.google.com/eventarc/advanced/docs/retry-events#transient
// 2. Or return 2xx for permanent errors. Since this is not semantically correct,
// we are holding on this option.

func (p *pubSubPushReceiver) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *pubSubPushReceiver) Shutdown(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// See: https://cloud.google.com/pubsub/docs/push
type pubSubPushRequest struct {
	Message         pubSubPushMessage `json:"message"`
	Subscription    string            `json:"subscription"`
	DeliveryAttempt int               `json:"deliveryAttempt"`
}

// See: https://cloud.google.com/pubsub/docs/reference/rest/v1/PubsubMessage
type pubSubPushMessage struct {
	Attributes  map[string]string `json:"attributes"`
	Data        []byte            `json:"data"` // go will automatically decode base64
	MessageID   string            `json:"messageId"`
	OrderingKey string            `json:"orderingKey,omitempty"`
	PublishTime time.Time         `json:"publishTime"`
}

// getFileContent retrieves file content from cloud storage notifications.
// Returns:
//   - (true, nil, nil) if the notification should be ignored (non-create event)
//   - (false, content, nil) if file content was successfully read
//   - (false, nil, error) if required attributes are missing or file read fails
//
// The function filters for OBJECT_FINALIZE events (new file creation) and
// ignores other storage events like deletion or metadata updates.
func getFileContent(
	ctx context.Context,
	attributes map[string]string,
	storageClient *storage.Client,
	extraMetadata map[string][]string,
) (bool, []byte, error) {
	_ = "STUB: not implemented"
	return false, nil, nil
}

// See: https://cloud.google.com/storage/docs/pubsub-notifications#events

// check that this notification is coming from a new file, otherwise ignore it

// TODO https://github.com/open-telemetry/opentelemetry-collector-contrib/issues/38780
// Reading the whole file into memory is not good.

// enrich metadata

func handlePubSubPushRequest[T any](
	ctx context.Context,
	r io.Reader,
	unmarshal func([]byte) (T, error),
	consume func(context.Context, T) error,
	storageClient *storage.Client,
	includeMetadata bool,
	tb *metadata.TelemetryBuilder,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Not coming from a storage notification, so we can
// use the message data instead. This happens for the
// cases in which the log is directly sent to Pub/Sub
// instead of being placed in a GCS file.

// this field is empty if the log is sent directly to Pub/Sub

// err is already marked as permanent, so any error wrapping that
// will be marked as permanent as well

func loadEncodingExtension[T any](host component.Host, encodingID *component.ID, signal string) (T, error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}

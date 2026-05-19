// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package solacereceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/solacereceiver"

import (
	"errors"

	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/solacereceiver/internal/metadata"
)

// tracesUnmarshaller deserializes the message body.
type tracesUnmarshaller interface {
	// unmarshal the amqp-message into traces.
	// Only valid traces are produced or error is returned
	unmarshal(message *inboundMessage) (ptrace.Traces, error)
}

// newTracesUnmarshaller returns a new unmarshaller ready for message unmarshalling
func newTracesUnmarshaller(logger *zap.Logger, telemetryBuilder *metadata.TelemetryBuilder, metricAttrs attribute.Set) tracesUnmarshaller {
	_ = "STUB: not implemented"
	return *new(tracesUnmarshaller)
}

// v1 unmarshaller is implemented by solaceMessageUnmarshallerV1

// solaceTracesUnmarshaller implements tracesUnmarshaller.
type solaceTracesUnmarshaller struct {
	logger                *zap.Logger
	telemetryBuilder      *metadata.TelemetryBuilder
	moveUnmarshallerV1    tracesUnmarshaller
	receiveUnmarshallerV1 tracesUnmarshaller
	egressUnmarshallerV1  tracesUnmarshaller
}

var (
	errUpgradeRequired = errors.New("unsupported trace message, upgrade required")
	errUnknownTopic    = errors.New("unknown topic")
	errEmptyPayload    = errors.New("no binary attachment")
)

// unmarshal will unmarshal an *solaceMessage into ptrace.Traces.
// It will make a decision based on the version of the message which unmarshalling strategy to use.
// For now, only receive v1 messages are used.
func (u *solaceTracesUnmarshaller) unmarshal(message *inboundMessage) (ptrace.Traces, error) {
	_ = "STUB: not implemented"
	return *new(ptrace.Traces), nil
}

// no topic

// Multiplex the topic string. For now we only have a single type handled

// we are a telemetry string

// we are handling a move span, validate the version is v1

// otherwise we are an unknown version

// we are handling a receive span, validate the version is v1

// otherwise we are an unknown version

// make lint happy, wants two boolean expressions to be written as a switch?!

// otherwise we are an unknown version

// if we don't know the type, we must upgrade

// unknown topic, do not require an upgrade

// common helper functions used by all unmarshallers

// Endpoint types
const (
	queueKind         = "queue"
	topicEndpointKind = "topic-endpoint"
)

// Transaction event keys
const (
	transactionInitiatorEventKey    = "messaging.solace.transaction_initiator"
	transactionIDEventKey           = "messaging.solace.transaction_id"
	transactedSessionNameEventKey   = "messaging.solace.transacted_session_name"
	transactedSessionIDEventKey     = "messaging.solace.transacted_session_id"
	transactionErrorMessageEventKey = "messaging.solace.transaction_error_message"
	transactionXIDEventKey          = "messaging.solace.transaction_xid"
)

// span keys
const (
	protocolAttrKey                    = "network.protocol.name"
	protocolVersionAttrKey             = "network.protocol.version"
	messageIDAttrKey                   = "messaging.message.id"
	conversationIDAttrKey              = "messaging.message.conversation_id"
	messageBodySizeBytesAttrKey        = "messaging.message.body.size"
	messageEnvelopeSizeBytesAttrKey    = "messaging.message.envelope.size"
	destinationNameAttrKey             = "messaging.destination.name"
	destinationTypeAttrKey             = "messaging.solace.destination.type"
	clientUsernameAttrKey              = "messaging.solace.client_username"
	clientNameAttrKey                  = "messaging.solace.client_name"
	partitionNumberKey                 = "messaging.solace.partition_number"
	replicationGroupMessageIDAttrKey   = "messaging.solace.replication_group_message_id"
	priorityAttrKey                    = "messaging.solace.priority"
	ttlAttrKey                         = "messaging.solace.ttl"
	dmqEligibleAttrKey                 = "messaging.solace.dmq_eligible"
	droppedEnqueueEventsSuccessAttrKey = "messaging.solace.dropped_enqueue_events_success"
	droppedEnqueueEventsFailedAttrKey  = "messaging.solace.dropped_enqueue_events_failed"
	replyToAttrKey                     = "messaging.solace.reply_to_topic"
	receiveTimeAttrKey                 = "messaging.solace.broker_receive_time_unix_nano"
	droppedUserPropertiesAttrKey       = "messaging.solace.dropped_application_message_properties"
	deliveryModeAttrKey                = "messaging.solace.delivery_mode"
	hostIPAttrKey                      = "server.address"
	hostPortAttrKey                    = "server.port"
	peerIPAttrKey                      = "network.peer.address"
	peerPortAttrKey                    = "network.peer.port"
)

// constant attributes
const (
	systemAttrKey        = "messaging.system"
	systemAttrValue      = "SolacePubSub+"
	operationNameAttrKey = "messaging.operation.name"
	operationTypeAttrKey = "messaging.operation.type"
)

func setResourceSpanAttributes(attrMap pcommon.Map, routerName, version string, messageVpnName *string) {
	_ = "STUB: not implemented"
	return
}

func rgmidToString(rgmid []byte, otelMetricAttrs attribute.Set, telemetryBuilder *metadata.TelemetryBuilder, logger *zap.Logger) string {
	_ = "STUB: not implemented"
	// rgmid[0] is the version of the rgmid
	return ""
}

// may be cases where the rgmid is empty or nil, len(rgmid) will return 0 if nil

// format: rmid1:aaaaa-bbbbbbbbbbb-cccccccc-dddddddd

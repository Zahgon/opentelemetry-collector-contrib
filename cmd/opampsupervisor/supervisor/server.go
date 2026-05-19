// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package supervisor

import (
	"context"
	"net/http"

	"github.com/open-telemetry/opamp-go/protobufs"
	"github.com/open-telemetry/opamp-go/server"
	serverTypes "github.com/open-telemetry/opamp-go/server/types"
)

type flattenedSettings struct {
	onMessage         func(conn serverTypes.Connection, message *protobufs.AgentToServer) *protobufs.ServerToAgent
	onConnecting      func(request *http.Request) (shouldConnect bool, rejectStatusCode int)
	onConnectionClose func(conn serverTypes.Connection)
	endpoint          string
}

func (fs flattenedSettings) toServerSettings() server.StartSettings {
	_ = "STUB: not implemented"
	return *new(server.StartSettings)
}

func (fs flattenedSettings) OnConnecting(request *http.Request) serverTypes.ConnectionResponse {
	_ = "STUB: not implemented"
	return *new(serverTypes.ConnectionResponse)
}

func (flattenedSettings) OnConnected(context.Context, serverTypes.Connection) {
	_ = "STUB: not implemented"
	return
}

func (fs flattenedSettings) OnMessage(_ context.Context, conn serverTypes.Connection, message *protobufs.AgentToServer) *protobufs.ServerToAgent {
	_ = "STUB: not implemented"
	return nil
}

func (fs flattenedSettings) OnConnectionClose(conn serverTypes.Connection) {
	_ = "STUB: not implemented"
	return
}

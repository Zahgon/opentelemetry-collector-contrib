// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package azureeventhubreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/azureeventhubreceiver"

import (
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azeventhubs/v2"
)

type azureEvent struct {
	AzEventData *azeventhubs.ReceivedEventData
}

func (a *azureEvent) EnqueueTime() *time.Time { _ = "STUB: not implemented"; return nil }

func (a *azureEvent) Properties() map[string]any { _ = "STUB: not implemented"; return nil }

func (a *azureEvent) Data() []byte { _ = "STUB: not implemented"; return nil }

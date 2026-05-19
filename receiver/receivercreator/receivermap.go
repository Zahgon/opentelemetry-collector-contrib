// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package receivercreator // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/receivercreator"

import (
	"go.opentelemetry.io/collector/component"

	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer"
)

// receiverMap is a multimap for mapping one id to many receivers. It does
// not deduplicate the same value being associated with the same key.
type receiverMap map[observer.EndpointID][]component.Component

// Put rcvr into key id. If rcvr is a duplicate it will still be added.
func (rm receiverMap) Put(id observer.EndpointID, rcvr component.Component) {
	_ = "STUB: not implemented"
	return
}

// Get receivers by id.
func (rm receiverMap) Get(id observer.EndpointID) []component.Component {
	_ = "STUB: not implemented"

	// Remove all receivers by id.
	return nil
}

func (rm receiverMap) RemoveAll(id observer.EndpointID) {
	_ = "STUB: not implemented"

	// Get all receivers in the map.
	return
}

func (rm receiverMap) Values() (out []component.Component) { _ = "STUB: not implemented"; return nil }

// Size is the number of total receivers in the map.
func (rm receiverMap) Size() (out int) { _ = "STUB: not implemented"; return 0 }

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package splunk // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/splunk"

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
)

// HostIDKey represents a host identifier.
type HostIDKey string

const (
	// HostIDKeyAWS AWS HostIDKey.
	HostIDKeyAWS HostIDKey = "AWSUniqueId"
	// HostIDKeyGCP GCP HostIDKey.
	HostIDKeyGCP HostIDKey = "gcp_id"
	// HostIDKeyAzure Azure HostIDKey.
	HostIDKeyAzure HostIDKey = "azure_resource_id"
	// HostIDKeyHost Host HostIDKey.
	HostIDKeyHost HostIDKey = "host.name"
)

// HostID is a unique key and value (usually used as a dimension) to uniquely identify a host
// using metadata about a cloud instance.
type HostID struct {
	// Key is the key name/type.
	Key HostIDKey
	// Value is the unique ID.
	ID string
}

// ResourceToHostID returns a boolean determining whether or not a HostID was able to be
// computed or not.
func ResourceToHostID(res pcommon.Resource) (HostID, bool) {
	_ = "STUB: not implemented"
	return *new(HostID), false
}

func azureID(attrs pcommon.Map, cloudAccount string) string { _ = "STUB: not implemented"; return "" }

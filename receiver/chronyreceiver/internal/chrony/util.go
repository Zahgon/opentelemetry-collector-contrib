// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package chrony // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/chronyreceiver/internal/chrony"

import (
	"errors"
)

var ErrInvalidNetwork = errors.New("invalid network format")

// splitSchemeEndpoint splits addr on "://" and returns the scheme and path.
func splitSchemeEndpoint(addr string) (scheme, endpoint string, err error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

// SplitNetworkEndpoint takes in a URL like string of the format: [network type]://[network endpoint]
// and then will return the network and the endpoint for the client to use for connection.
func SplitNetworkEndpoint(addr string) (network, endpoint string, err error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

// Chrony uses socket type DGRAM which converts to `unixgram`,
// in order to preserve configuration of existing clients, this will overwrite the network type

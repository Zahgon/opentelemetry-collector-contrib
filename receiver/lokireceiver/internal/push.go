// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package internal // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/lokireceiver/internal"

import (
	"net/http"

	"github.com/grafana/loki/pkg/push"
)

var (
	contentType = http.CanonicalHeaderKey("Content-Type")
	contentEnc  = http.CanonicalHeaderKey("Content-Encoding")
)

const applicationJSON = "application/json"

func ParseRequest(req *http.Request) (*push.PushRequest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

/* params */

// When no content-type header is set or when it is set to
// `application/x-protobuf`: expect snappy compression.

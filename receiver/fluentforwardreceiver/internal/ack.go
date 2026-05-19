// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package internal // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/fluentforwardreceiver/internal"

import "github.com/tinylib/msgp/msgp"

type AckResponse struct {
	Ack string `msg:"ack"`
}

func (z AckResponse) EncodeMsg(en *msgp.Writer) error {
	_ = "STUB: not implemented"
	// map header, size 1
	// write "ack"
	return nil
}

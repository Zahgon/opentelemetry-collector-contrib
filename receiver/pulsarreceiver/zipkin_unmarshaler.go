// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build !aix

package pulsarreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/pulsarreceiver"

// copy from kafka receiver
const (
	zipkinProtobufEncoding = "zipkin_proto"
	zipkinJSONEncoding     = "zipkin_json"
	zipkinThriftEncoding   = "zipkin_thrift"
)

func newZipkinProtobufUnmarshaler() TracesUnmarshaler {
	_ = "STUB: not implemented"
	return *new(TracesUnmarshaler)
}

func newZipkinJSONUnmarshaler() TracesUnmarshaler {
	_ = "STUB: not implemented"
	return *new(TracesUnmarshaler)
}

func newZipkinThriftUnmarshaler() TracesUnmarshaler {
	_ = "STUB: not implemented"
	return *new(TracesUnmarshaler)
}

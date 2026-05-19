// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package avrologencodingextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding/avrologencodingextension"

import (
	"github.com/linkedin/goavro/v2"
)

type avroDeserializer interface {
	Deserialize([]byte) (map[string]any, error)
}

type avroStaticSchemaDeserializer struct {
	codec *goavro.Codec
}

func newAVROStaticSchemaDeserializer(schema string) (avroDeserializer, error) {
	_ = "STUB: not implemented"
	return *new(avroDeserializer), nil
}

func (d *avroStaticSchemaDeserializer) Deserialize(data []byte) (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

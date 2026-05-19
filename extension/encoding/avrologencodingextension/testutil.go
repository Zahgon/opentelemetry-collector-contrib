// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package avrologencodingextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding/avrologencodingextension"

import (
	"testing"

	"github.com/linkedin/goavro/v2"
)

func encodeAVROLogTestData(codec *goavro.Codec, data string) []byte {
	_ = "STUB: not implemented"
	return nil
}

func loadAVROSchemaFromFile(path string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createAVROTestData(t *testing.T) (string, []byte) { _ = "STUB: not implemented"; return "", nil }

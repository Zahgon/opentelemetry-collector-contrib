// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package header // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/fileconsumer/internal/header"

import (
	"bufio"
	"regexp"

	"go.opentelemetry.io/collector/component"
	"golang.org/x/text/encoding"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
)

type Config struct {
	regex             *regexp.Regexp
	SplitFunc         bufio.SplitFunc
	metadataOperators []operator.Config
}

func NewConfig(set component.TelemetrySettings, matchRegex string, metadataOperators []operator.Config, enc encoding.Encoding) (*Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// This is the default output we created, it's always valid

// Filter processor also may fail to propagate some entries

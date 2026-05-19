// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package main // import "github.com/open-telemetry/opentelemetry-collector-contrib/cmd/schemagen"

import (
	"log"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error { _ = "STUB: not implemented"; return nil }

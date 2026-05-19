// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package main // import "github.com/open-telemetry/opentelemetry-collector-contrib/cmd/golden"

import (
	"log"
	"os"

	"go.opentelemetry.io/collector/config/configoptional"
)

// insertDefault is a helper function to insert a default value for a configoptional.Optional type.
func insertDefault[T any](opt *configoptional.Optional[T]) error {
	_ = "STUB: not implemented"
	return nil
}

func main() {
	if err := run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(args []string) error { _ = "STUB: not implemented"; return nil }

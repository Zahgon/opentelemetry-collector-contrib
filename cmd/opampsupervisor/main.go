// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"log"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func runInteractive() error { _ = "STUB: not implemented"; return nil }

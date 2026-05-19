// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package internal

type Metadata struct {
	Type   string `mapstructure:"type"`
	Status struct {
		Class string `mapstructure:"class"`
	} `mapstructure:"status"`
	Parent string `mapstructure:"parent"`
}

func ReadMetadata(dir string) (*Metadata, bool) { _ = "STUB: not implemented"; return nil, false }

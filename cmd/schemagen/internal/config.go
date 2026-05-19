// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package internal

import (
	"flag"
)

type RunMode string

const (
	Component RunMode = "component"
	Package   RunMode = "package"
)

type Config = struct {
	Mode          RunMode
	DirPath       string
	OutputFolder  string
	ConfigPackage string
	ConfigType    string
	FileType      string
	Class         string
	Mappings      Mappings
	AllowedRefs   []string
	Namespace     string
}

var (
	configType   = flag.String("c", "Config", "Config type name for component schema generation")
	outputFolder = flag.String("o", "", "Output schema folder (defaults to input folder)")
	fileType     = flag.String("t", "yaml", "Output file type (yaml or json)")
)

func usage() { _ = "STUB: not implemented"; return }

func ReadConfig() (*Config, error) { _ = "STUB: not implemented"; return nil, nil }

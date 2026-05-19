// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package internal

const (
	SettingsFileName = ".schemagen.yaml"
)

type (
	Settings struct {
		Namespace          string             `yaml:"namespace"`
		Mappings           Mappings           `yaml:"mappings"`
		ComponentOverrides ComponentOverrides `yaml:"componentOverrides"`
		AllowedRefs        []string           `yaml:"allowedRefs"`
	}
	Mappings           map[string]PackagesMapping
	PackagesMapping    map[string]TypeDesc
	ComponentOverrides map[string]ComponentOverride
	ComponentOverride  struct {
		ConfigName string `yaml:"configName"`
	}
	TypeDesc struct {
		SchemaType     SchemaType `yaml:"schemaType"`
		Format         string     `yaml:"format"`
		SkipAnnotation bool       `yaml:"skipAnnotation"`
	}
)

func ReadSettingsFile() (*Settings, bool) { _ = "STUB: not implemented"; return nil, false }

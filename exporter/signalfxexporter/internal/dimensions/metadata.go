// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package dimensions // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/signalfxexporter/internal/dimensions"

import (
	"regexp"

	metadata "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/experimentalmetricmetadata"
)

// MetadataUpdateClient is an interface for pushing metadata updates
type MetadataUpdateClient interface {
	PushMetadata([]*metadata.MetadataUpdate) error
}

// resourceIDsSkipSanitization lists resource ID keys whose properties should
// pass through without label-prefix stripping or k8s.service.* rewriting.
var resourceIDsSkipSanitization = map[string]bool{
	"k8s.service.uid":               true,
	"k8s.persistentvolume.uid":      true,
	"k8s.persistentvolumeclaim.uid": true,
}

func getDimensionUpdateFromMetadata(
	defaults map[string]string,
	metadata metadata.MetadataUpdate,
	nonAlphanumericDimChars string,
	stripK8sLabelPrefix bool,
) *DimensionUpdate {
	_ = "STUB: not implemented"
	return nil
}

const (
	oTelK8sServicePrefix = "k8s.service."
	sfxK8sServicePrefix  = "kubernetes_service_"
)

var oTelK8sLabelRe = regexp.MustCompile(`^k8s\.[^.]+\.label\.(.+)$`)

func sanitizeProperty(property string, stripK8sLabelPrefix bool) string {
	_ = "STUB: not implemented"
	return ""
}

func getPropertiesAndTags(defaults map[string]string, kmu metadata.MetadataUpdate, skipSanitization, stripK8sLabelPrefix bool) (map[string]*string, map[string]bool) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Treat it as a remove if a property update has empty value since
// this cannot be a tag as tags can either be added or removed but
// not updated.

func (dc *DimensionClient) PushMetadata(metadata []*metadata.MetadataUpdate) error {
	_ = "STUB: not implemented"
	return nil
}

// FilterKeyChars filters dimension key characters, replacing non-alphanumeric characters
// (except those in nonAlphanumericDimChars) with underscores.
func FilterKeyChars(str, nonAlphanumericDimChars string) string {
	_ = "STUB: not implemented"
	return ""
}

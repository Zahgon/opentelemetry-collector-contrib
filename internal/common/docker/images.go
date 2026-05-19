// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package docker // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/common/docker"

import (
	"go.uber.org/zap"
)

type ImageRef struct {
	Repository      string
	Tag             string
	Digest          string
	DigestAlgorithm string
}

// ParseImageName extracts image repository, tag, digest and digest algorithm from a combined image reference
// e.g. example.com:5000/alpine/alpine:test --> `example.com:5000/alpine/alpine` and `test`
func ParseImageName(image string) (ImageRef, error) {
	_ = "STUB: not implemented"
	return *new(ImageRef), nil
}

// Basic image reference

// If no tag provided in image reference - default to "latest"

// No image digest provided

// Image digest is incorrect or no digest

// CanonicalImageRef tries to parse provided image reference
// and return canonical image reference instead, i.e. image reference
// that have qualified repository path and image name
// If it's not possible to normalize image reference - empty string and error returned
// Extracted from k8sattributesprocessor for future reuse
func CanonicalImageRef(image string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func LogParseError(err error, image string, logger *zap.Logger) { _ = "STUB: not implemented"; return }

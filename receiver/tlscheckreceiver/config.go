// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package tlscheckreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/tlscheckreceiver"

import (
	"errors"

	"go.opentelemetry.io/collector/config/confignet"
	"go.opentelemetry.io/collector/config/configopaque"
	"go.opentelemetry.io/collector/scraper/scraperhelper"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/tlscheckreceiver/internal/metadata"
)

// Predefined error responses for configuration validation failures
var (
	errInvalidEndpoint   = errors.New(`"endpoint" must be in the form of <hostname>:<port>`)
	errInvalidFileFormat = errors.New(`"file_format" must be one of: auto, pem, jks, pkcs12`)
)

// FileFormat represents the format of a local certificate file.
type FileFormat string

const (
	// FileFormatAuto infers the format from the file extension.
	FileFormatAuto FileFormat = "auto"
	// FileFormatPEM indicates a PEM-encoded certificate file.
	FileFormatPEM FileFormat = "pem"
	// FileFormatJKS indicates a Java KeyStore file.
	FileFormatJKS FileFormat = "jks"
	// FileFormatPKCS12 indicates a PKCS#12 / PFX keystore file.
	FileFormatPKCS12 FileFormat = "pkcs12"
)

// CertificateTarget represents a target for certificate checking, which can be either
// a network endpoint or a local file
type CertificateTarget struct {
	confignet.TCPAddrConfig `mapstructure:",squash"`
	FilePath                string              `mapstructure:"file_path"`
	FileFormat              FileFormat          `mapstructure:"file_format"`
	Password                configopaque.String `mapstructure:"password"`

	// prevent unkeyed literal initialization
	_ struct{}
}

// Config defines the configuration for the various elements of the receiver agent.
type Config struct {
	scraperhelper.ControllerConfig `mapstructure:",squash"`
	metadata.MetricsBuilderConfig  `mapstructure:",squash"`
	Targets                        []*CertificateTarget `mapstructure:"targets"`

	// prevent unkeyed literal initialization
	_ struct{}
}

func validateTarget(ct *CertificateTarget) error { _ = "STUB: not implemented"; return nil }

// valid — "" is treated the same as "auto"

// Endpoint-based target: file-related options must not be set.

func (cfg *Config) Validate() error { _ = "STUB: not implemented"; return nil }

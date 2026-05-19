// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package system // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/metadataproviders/system"

import (
	"context"
	"net"

	"github.com/shirou/gopsutil/v4/cpu"
	"go.opentelemetry.io/otel/sdk/resource"
)

// nameInfoProvider abstracts domain name resolution so it can be swapped for
// testing
type nameInfoProvider struct {
	osHostname  func() (string, error)
	lookupCNAME func(string) (string, error)
	lookupHost  func(string) ([]string, error)
	lookupAddr  func(string) ([]string, error)
}

// newNameInfoProvider creates a name info provider for production use, using
// DNS to resolve domain names
func newNameInfoProvider() nameInfoProvider {
	_ = "STUB: not implemented"
	return *new(nameInfoProvider)
}

type Provider interface {
	// Hostname returns the OS hostname
	Hostname() (string, error)

	// FQDN returns the fully qualified domain name
	FQDN() (string, error)

	// OSDescription returns a human readable description of the OS.
	OSDescription(ctx context.Context) (string, error)

	// OSType returns the host operating system
	OSType() (string, error)

	// OSVersion returns the version of the operating system
	OSVersion() (string, error)

	// LookupCNAME returns the canonical name for the current host
	LookupCNAME() (string, error)

	// ReverseLookupHost does a reverse DNS query on the current host's IP address
	ReverseLookupHost() (string, error)

	// HostID returns Host Unique Identifier
	HostID(ctx context.Context) (string, error)

	// HostArch returns the host architecture
	HostArch() (string, error)

	// HostIPs returns the host's IP interfaces
	HostIPs() ([]net.IP, error)

	// HostMACs returns the host's MAC addresses
	HostMACs() ([]net.HardwareAddr, error)

	// CPUInfo returns the host's CPU info
	CPUInfo(ctx context.Context) ([]cpu.InfoStat, error)
	// OSName returns the OS name according to semantic conventions.
	OSName(ctx context.Context) (string, error)

	// OSBuildID returns the OS build ID according to semantic conventions.
	OSBuildID(ctx context.Context) (string, error)

	// HostInterfaces returns the host's Interfaces info
	HostInterfaces() ([]net.Interface, error)
}

type systemMetadataProvider struct {
	nameInfoProvider
	newResource func(context.Context, ...resource.Option) (*resource.Resource, error)
}

func NewProvider() Provider { _ = "STUB: not implemented"; return *new(Provider) }

func (systemMetadataProvider) OSType() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (systemMetadataProvider) OSVersion() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (systemMetadataProvider) FQDN() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (p systemMetadataProvider) Hostname() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (p systemMetadataProvider) LookupCNAME() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (p systemMetadataProvider) ReverseLookupHost() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (p systemMetadataProvider) hostnameToDomainName(hostname string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (p systemMetadataProvider) reverseLookup(ipAddresses []string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (p systemMetadataProvider) fromOption(ctx context.Context, opt resource.Option, semconv string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (p systemMetadataProvider) HostID(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (p systemMetadataProvider) OSDescription(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// OSName returns the OS name from host metadata.
func (systemMetadataProvider) OSName(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// OSBuildID returns the OS build ID based on the platform.
func (systemMetadataProvider) OSBuildID(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// macOS: read ProductBuildVersion from SystemVersion.plist

// Linux: read BUILD_ID from /etc/os-release

// parsePlistValue parses the XML plist and returns the <string> value after the given <key>.
func parsePlistValue(data []byte, key string) string { _ = "STUB: not implemented"; return "" }

// parseOSReleaseValue parses key=value pairs from os-release data.
func parseOSReleaseValue(data []byte, key string) string { _ = "STUB: not implemented"; return "" }

func (systemMetadataProvider) HostArch() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (systemMetadataProvider) HostIPs() (ips []net.IP, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// skip if the interface is down or is a loopback interface

// skip loopback IPs

func (systemMetadataProvider) HostMACs() (macs []net.HardwareAddr, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// skip if the interface is down or is a loopback interface

func (systemMetadataProvider) CPUInfo(ctx context.Context) ([]cpu.InfoStat, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (systemMetadataProvider) HostInterfaces() ([]net.Interface, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

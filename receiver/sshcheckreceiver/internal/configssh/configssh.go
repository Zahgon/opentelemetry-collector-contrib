// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package configssh // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/sshcheckreceiver/internal/configssh"

import (
	"errors"
	"time"

	"github.com/pkg/sftp"
	"go.opentelemetry.io/collector/component"
	"golang.org/x/crypto/ssh"
)

const (
	defaultClientVersion = "SSH-2.0-OTelClient"
)

var errMissingKnownHosts = errors.New(`known_hosts file is missing`)

type SSHClientSettings struct {
	// Endpoint is always required
	Endpoint string        `mapstructure:"endpoint"`
	Timeout  time.Duration `mapstructure:"timeout"`

	// authentication requires a Username and either a Password or KeyFile
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	KeyFile  string `mapstructure:"key_file"`

	// file path to the known_hosts
	KnownHosts string `mapstructure:"known_hosts"`

	// IgnoreHostKey provides an insecure path to quickstarts and testing
	IgnoreHostKey bool `mapstructure:"ignore_host_key"`
}

type Client struct {
	*ssh.Client
	*ssh.ClientConfig
	DialFunc func(network, address string, config *ssh.ClientConfig) (*ssh.Client, error)
}

// Dial starts an SSH session.
func (c *Client) Dial(endpoint string) (err error) { _ = "STUB: not implemented"; return nil }

func (c *Client) SFTPClient() (*SFTPClient, error) { _ = "STUB: not implemented"; return nil, nil }

type SFTPClient struct {
	*sftp.Client
	*ssh.ClientConfig
}

// ToClient creates an SSHClient.
func (scs *SSHClientSettings) ToClient(component.Host, component.TelemetrySettings) (*Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:gosec // #nosec G106

func defaultKnownHostsPath() (string, error) { _ = "STUB: not implemented"; return "", nil }

func defaultKnownHostsCallback() (hkc ssh.HostKeyCallback, err error) {
	_ = "STUB: not implemented"
	return *new(ssh.HostKeyCallback), nil
}

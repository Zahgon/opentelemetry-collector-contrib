// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build !windows

package podmanreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/podmanreceiver"

import (
	"net/http"
	"net/url"

	"go.uber.org/zap"
	"golang.org/x/crypto/ssh"
)

// most of this file has been adopted from https://github.com/containers/podman/blob/main/pkg/bindings/connection.go
// and then simplified to remove things we do not need.

func newPodmanConnection(logger *zap.Logger, endpoint, sshKey, sshPassphrase string) (*http.Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// autofix unix://path_element vs unix:///path_element

func tcpConnection(_url *url.URL) *http.Client { _ = "STUB: not implemented"; return nil }

func unixConnection(_url *url.URL) *http.Client { _ = "STUB: not implemented"; return nil }

func sshConnection(logger *zap.Logger, _url *url.URL, secure bool, key, passphrase string) (*http.Client, error) {
	_ = "STUB: not implemented"
	return nil,
		// order Signers are appended to this list determines which key is presented to server
		nil
}

// Dedup signers based on fingerprint, ssh-agent keys override CONTAINER_SSHKEY

// #nosec

func publicKey(path string, passphrase []byte) (ssh.Signer, error) {
	_ = "STUB: not implemented"
	return *new(ssh.Signer), nil
}

func hostKey(logger *zap.Logger, host string) ssh.PublicKey {
	_ = "STUB: not implemented"
	// parse OpenSSH known_hosts file
	// ssh or use ssh-keyscan to get initial key
	return *new(ssh.PublicKey)
}

// support -H parameter for ssh-keyscan

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package redfish // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/redfishreceiver/internal/redfish"

import (
	"net/http"
	"net/url"
	"time"

	"go.opentelemetry.io/collector/config/configopaque"
)

type Client struct {
	client           http.Client
	BaseURL          *url.URL
	redfishVersion   string
	host             string
	userName         string
	password         configopaque.String
	computerSystemID string
}

// redfish client options
type ClientOption func(*clientOptions)

type clientOptions struct {
	RedfishVersion string
	ClientTimeout  time.Duration
	Insecure       bool
}

func WithRedfishVersion(version string) ClientOption {
	_ = "STUB: not implemented"
	return *new(ClientOption)
}

func WithClientTimeout(timeout time.Duration) ClientOption {
	_ = "STUB: not implemented"
	return *new(ClientOption)
}

func WithInsecure(insecure bool) ClientOption { _ = "STUB: not implemented"; return *new(ClientOption) }

// NewRedfishClient is a function to create new redfish clients
func NewClient(computerSystemID, addr, user string, pwd configopaque.String, opts ...ClientOption) (*Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) setHeaders(req *http.Request) { _ = "STUB: not implemented"; return }

// GetComputerSystem is a method to get the client's computer system data
func (c *Client) GetComputerSystem() (*ComputerSystem, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetChassis is a method that gets chassis data given a chassis odata ref
func (c *Client) GetChassis(ref string) (*Chassis, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetThermal is a method that gets thermal data given a thermal odata ref
func (c *Client) GetThermal(ref string) (*Thermal, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

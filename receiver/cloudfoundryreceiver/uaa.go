// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package cloudfoundryreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/cloudfoundryreceiver"

import (
	"sync"
	"time"

	"github.com/cloudfoundry-incubator/uaago"
	"go.uber.org/zap"
)

const (
	// time added to expiration time to reduce chance of using expired token due to network latency
	expirationTimeBuffer = -5 * time.Second
)

type uaaTokenProvider struct {
	client         *uaago.Client
	logger         *zap.Logger
	username       string
	password       string
	tlsSkipVerify  bool
	cachedToken    string
	expirationTime *time.Time
	mutex          *sync.Mutex
}

func newUAATokenProvider(logger *zap.Logger, config LimitedClientConfig, username, password string) (*uaaTokenProvider, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (utp *uaaTokenProvider) ProvideToken() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

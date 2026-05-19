// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package huaweicloudcesreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/huaweicloudcesreceiver"

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/config"
)

func createHTTPConfig(cfg huaweiSessionConfig) (*config.HttpConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func configureHTTPProxy(cfg huaweiSessionConfig) (*config.Proxy, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Configure the username and password if the proxy requires authentication

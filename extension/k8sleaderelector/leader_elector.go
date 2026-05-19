// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package k8sleaderelector // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/k8sleaderelector"

import (
	"context"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/leaderelection"
)

func newK8sLeaderElector(
	cfg *Config,
	client kubernetes.Interface,
	onStartedLeading func(context.Context),
	onStoppedLeading func(),
	identity string,
) (*leaderelection.LeaderElector, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

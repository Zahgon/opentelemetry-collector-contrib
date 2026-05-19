// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package k8sclient // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/k8s/k8sclient"

import (
	"sync"
	"time"

	"go.uber.org/zap"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
)

const (
	cronJob = "CronJob"
)

type JobClient interface {
	// get the mapping between job and cronjob
	JobToCronJob() map[string]string
}

type noOpJobClient struct{}

func (*noOpJobClient) JobToCronJob() map[string]string { _ = "STUB: not implemented"; return nil }

func (*noOpJobClient) shutdown() { _ = "STUB: not implemented"; return }

type jobClientOption func(*jobClient)

func jobSyncCheckerOption(checker initialSyncChecker) jobClientOption {
	_ = "STUB: not implemented"
	return *new(jobClientOption)
}

type jobClient struct {
	stopChan chan struct{}
	stopped  bool

	store *ObjStore

	syncChecker initialSyncChecker

	mu              sync.RWMutex
	cachedJobMap    map[string]time.Time
	jobToCronJobMap map[string]string
}

func (c *jobClient) JobToCronJob() map[string]string { _ = "STUB: not implemented"; return nil }

func (c *jobClient) refresh() { _ = "STUB: not implemented"; return }

func newJobClient(clientSet kubernetes.Interface, logger *zap.Logger, options ...jobClientOption) (*jobClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// check the init sync for potential connection issue

func (c *jobClient) shutdown() { _ = "STUB: not implemented"; return }

func transformFuncJob(obj any) (any, error) { _ = "STUB: not implemented"; return *new(any), nil }

func createJobListWatch(client kubernetes.Interface, ns string) cache.ListerWatcher {
	_ = "STUB: not implemented"
	return *new(cache.ListerWatcher)
}

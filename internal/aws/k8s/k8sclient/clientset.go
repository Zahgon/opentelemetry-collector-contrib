// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package k8sclient // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/k8s/k8sclient"

import (
	"os"
	"path/filepath"
	"sync"
	"time"

	"go.uber.org/zap"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

const (
	cacheTTL = 10 * time.Minute
)

// Option is a struct that can be used to change the configuration of passed K8sClient
// It can be used as an option to the Get(...) function to create a customized K8sClient
type Option struct {
	name string
	set  func(*K8sClient)
}

var mu = &sync.Mutex{}

var optionsToK8sClient = map[string]*K8sClient{}

type stopper interface {
	shutdown()
}

func shutdownClient(client stopper, mu *sync.Mutex, afterShutdown func()) {
	_ = "STUB: not implemented"
	return
}

type cacheReflector interface {
	LastSyncResourceVersion() string
	Run(<-chan struct{})
}

type initialSyncChecker interface {
	// check the initial sync of cache reflector and log the warnMessage if timeout
	Check(reflector cacheReflector, warnMessage string)
}

// reflectorSyncChecker implements initialSyncChecker interface
type reflectorSyncChecker struct {
	pollInterval time.Duration
	pollTimeout  time.Duration
	logger       *zap.Logger
}

func (r *reflectorSyncChecker) Check(reflector cacheReflector, warnMessage string) {
	_ = "STUB: not implemented"
	return
}

// KubeConfigPath provides the option to set the kube config which will be used if the
// service account that kubernetes gives to pods can't be used
func KubeConfigPath(kubeConfigPath string) Option { _ = "STUB: not implemented"; return *new(Option) }

// InitSyncPollInterval provides the option to set the init sync poll interval
// for testing connection to kubernetes api server
func InitSyncPollInterval(pollInterval time.Duration) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// InitSyncPollTimeout provides the option to set the init sync poll timeout
// for testing connection to kubernetes api server
func InitSyncPollTimeout(pollTimeout time.Duration) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func getStringifiedOptions(options ...Option) string { _ = "STUB: not implemented"; return "" }

// Get returns a singleton instance of k8s client
// If the intialization fails, it returns nil
func Get(logger *zap.Logger, options ...Option) *K8sClient { _ = "STUB: not implemented"; return nil }

// construct the k8s client

type epClientWithStopper interface {
	EpClient
	stopper
}

type jobClientWithStopper interface {
	JobClient
	stopper
}

type nodeClientWithStopper interface {
	NodeClient
	stopper
}

type podClientWithStopper interface {
	PodClient
	stopper
}

type replicaSetClientWithStopper interface {
	ReplicaSetClient
	stopper
}

type K8sClient struct {
	kubeConfigPath       string
	initSyncPollInterval time.Duration
	initSyncPollTimeout  time.Duration

	clientSet kubernetes.Interface

	syncChecker *reflectorSyncChecker

	epMu sync.Mutex
	ep   epClientWithStopper

	podMu sync.Mutex
	pod   podClientWithStopper

	nodeMu sync.Mutex
	node   nodeClientWithStopper

	jobMu sync.Mutex
	job   jobClientWithStopper

	rsMu       sync.Mutex
	replicaSet replicaSetClientWithStopper

	logger *zap.Logger
}

func (c *K8sClient) init(logger *zap.Logger, options ...Option) error {
	c.logger = logger

	// set up some default configs
	c.kubeConfigPath = filepath.Join(os.Getenv("HOME"), ".kube/config")
	c.initSyncPollInterval = 50 * time.Millisecond
	c.initSyncPollTimeout = 2 * time.Second

	// take additional options passed in
	for _, opt := range options {
		opt.set(c)
	}

	config, err := rest.InClusterConfig()
	if err != nil {
		c.logger.Warn("cannot find in cluster config", zap.Error(err))
		config, err = clientcmd.BuildConfigFromFlags("", c.kubeConfigPath)
		if err != nil {
			c.logger.Error("failed to build config", zap.Error(err))
			return err
		}
	}
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		c.logger.Error("failed to build ClientSet", zap.Error(err))
		return err
	}

	c.syncChecker = &reflectorSyncChecker{
		pollInterval: c.initSyncPollInterval,
		pollTimeout:  c.initSyncPollTimeout,
		logger:       c.logger,
	}

	c.clientSet = client
	c.ep = nil
	c.pod = nil
	c.node = nil
	c.job = nil
	c.replicaSet = nil

	return nil
}

func (c *K8sClient) GetEpClient() EpClient { _ = "STUB: not implemented"; return *new(EpClient) }

func (c *K8sClient) ShutdownEpClient() { _ = "STUB: not implemented"; return }

func (c *K8sClient) GetPodClient() PodClient { _ = "STUB: not implemented"; return *new(PodClient) }

func (c *K8sClient) ShutdownPodClient() { _ = "STUB: not implemented"; return }

func (c *K8sClient) GetNodeClient() NodeClient { _ = "STUB: not implemented"; return *new(NodeClient) }

func (c *K8sClient) ShutdownNodeClient() { _ = "STUB: not implemented"; return }

func (c *K8sClient) GetJobClient() JobClient { _ = "STUB: not implemented"; return *new(JobClient) }

func (c *K8sClient) ShutdownJobClient() { _ = "STUB: not implemented"; return }

func (c *K8sClient) GetReplicaSetClient() ReplicaSetClient {
	_ = "STUB: not implemented"
	return *new(ReplicaSetClient)
}

func (c *K8sClient) ShutdownReplicaSetClient() { _ = "STUB: not implemented"; return }

func (c *K8sClient) GetClientSet() kubernetes.Interface {
	_ = "STUB: not implemented"
	return *

	// Shutdown stops K8sClient
	new(kubernetes.Interface)
}

func (c *K8sClient) Shutdown() { _ = "STUB: not implemented"; return }

// remove the current instance of k8s client from map

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build linux

package cadvisor // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awscontainerinsightreceiver/internal/cadvisor"

import (
	"net/http"
	"time"

	"github.com/google/cadvisor/cache/memory"
	cadvisormetrics "github.com/google/cadvisor/container"

	// Register filesystem plugins via init() functions
	_ "github.com/google/cadvisor/fs/overlay/install"
	_ "github.com/google/cadvisor/fs/tmpfs/install"
	_ "github.com/google/cadvisor/fs/vfs/install"
	cInfo "github.com/google/cadvisor/info/v1"
	"github.com/google/cadvisor/manager"
	"github.com/google/cadvisor/utils/sysfs"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awscontainerinsightreceiver/internal/cadvisor/extractors"
)

// The amount of time for which to keep stats in memory.
const statsCacheDuration = 2 * time.Minute

// Max collection interval, it is not meaningful if allowDynamicHousekeeping = false
var maxHousekeepingInterval = 15 * time.Second

// When allowDynamicHousekeeping is true, the collection interval is floating between 1s(default) to maxHousekeepingInterval
var allowDynamicHousekeeping = true

const defaultHousekeepingInterval = 10 * time.Second

// define the interface for cadvisor manager so that we can mock it in tests.
// For detailed information about these APIs, see https://github.com/google/cadvisor/blob/release-v0.39/manager/manager.go
type cadvisorManager interface {
	// Start the manager. Calling other manager methods before this returns
	// may produce undefined behavior.
	Start() error

	// Get information about all subcontainers of the specified container (includes self).
	SubcontainersInfo(containerName string, query *cInfo.ContainerInfoRequest) ([]*cInfo.ContainerInfo, error)
}

// define a function type for creating a cadvisor manager
type createCadvisorManager func(*memory.InMemoryCache, sysfs.SysFs, manager.HousekeepingConfig, cadvisormetrics.MetricSet, *http.Client,
	[]string, string) (cadvisorManager, error)

// this is the default function that are used in production code to create a cadvisor manager
// We define defaultCreateManager and use it as a field value for Cadvisor mainly for the purpose of
// unit testing. Ideally we should just replace the cadvisor manager with a mock in unit tests instead of calling
// createCadvisorManager() to create a mock. However this means that we will not be able
// to unit test the code in `initManager(...)`. For now, I will leave code as it is. Hopefully we can find
// a better way to mock the cadvisor related part in the future.
var defaultCreateManager = func(memoryCache *memory.InMemoryCache, sysfs sysfs.SysFs, housekeepingConfig manager.HousekeepingConfig,
	includedMetricsSet cadvisormetrics.MetricSet, collectorHTTPClient *http.Client, rawContainerCgroupPathPrefixWhiteList []string,
	perfEventsFile string,
) (cadvisorManager, error) {
	return manager.New(memoryCache, sysfs, housekeepingConfig, includedMetricsSet, collectorHTTPClient, rawContainerCgroupPathPrefixWhiteList, []string{}, perfEventsFile, 0)
}

// Option is a function that can be used to configure Cadvisor struct
type Option func(*Cadvisor)

func cadvisorManagerCreator(f createCadvisorManager) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithDecorator constructs an option for configuring the metric decorator
func WithDecorator(d Decorator) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithECSInfoCreator(f EcsInfo) Option { _ = "STUB: not implemented"; return *new(Option) }

type hostInfo interface {
	GetNumCores() int64
	GetMemoryCapacity() int64
	GetClusterName() string
	GetEBSVolumeID(string) string
	ExtractEbsIDsUsedByKubernetes() map[string]string
	GetInstanceID() string
	GetInstanceType() string
	GetAutoScalingGroupName() string
}

type EcsInfo interface {
	GetCPUReserved() int64
	GetMemReserved() int64
	GetRunningTaskCount() int64
	GetContainerInstanceID() string
	GetClusterName() string
}

type Decorator interface {
	Decorate(*extractors.CAdvisorMetric) *extractors.CAdvisorMetric
	Shutdown() error
}

type Cadvisor struct {
	logger                *zap.Logger
	nodeName              string // get the value from downward API
	createCadvisorManager createCadvisorManager
	manager               cadvisorManager
	version               string
	hostInfo              hostInfo
	k8sDecorator          Decorator
	ecsInfo               EcsInfo
	containerOrchestrator string
	metricsExtractors     []extractors.MetricExtractor
}

func init() {
	// Override the default cAdvisor housekeeping interval.
	// We should use a proper way to configure once the issue is resolved: https://github.com/google/cadvisor/issues/2886
	*manager.HousekeepingInterval = defaultHousekeepingInterval
}

// New creates a Cadvisor struct which can generate metrics from embedded cadvisor lib
func New(containerOrchestrator string, hostInfo hostInfo, logger *zap.Logger, options ...Option) (*Cadvisor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// apply additional options

func (c *Cadvisor) GetMetricsExtractors() []extractors.MetricExtractor {
	_ = "STUB: not implemented"
	return nil
}

func (c *Cadvisor) Shutdown() error { _ = "STUB: not implemented"; return nil }

func (c *Cadvisor) addEbsVolumeInfo(tags, ebsVolumeIDsUsedAsPV map[string]string) {
	_ = "STUB: not implemented"
	return
}

func (c *Cadvisor) addECSMetrics(cadvisormetrics []*extractors.CAdvisorMetric) {
	_ = "STUB: not implemented"
	return
}

// cgroup standard cpulimits should be cadvisor standard * 1.024

func addECSResources(tags map[string]string) { _ = "STUB: not implemented"; return }

func (c *Cadvisor) decorateMetrics(cadvisormetrics []*extractors.CAdvisorMetric) []*extractors.CAdvisorMetric {
	_ = "STUB: not implemented"
	return nil
}

// add version

// add NodeName for node, pod and container

// add instance id and type

// add scaling group name

// add ECS cluster name and container instance id

// add tags for EKS

// GetMetrics generates metrics from cadvisor
func (c *Cadvisor) GetMetrics() []pmetric.Metrics { _ = "STUB: not implemented"; return nil }

// For EKS don't emit metrics if the cluster name is not detected

// initManager accepts a function of type createCadvisorManager which can be used
// to create a cadvisor manager
func (c *Cadvisor) initManager(createManager createCadvisorManager) error {
	_ = "STUB: not implemented"
	return nil
}

// Create and start the cAdvisor container manager.

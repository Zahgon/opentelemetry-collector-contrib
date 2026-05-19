// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ecsmock // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/ecsobserver/internal/ecsmock"

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	ec2types "github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	ecstypes "github.com/aws/aws-sdk-go-v2/service/ecs/types"
)

// PageLimit defines number of items in a single page for different APIs.
// Those numbers can be found on the Input and Output struct comments.
// Call DefaultPageLimit() to config the mock to use numbers same as the actual AWS API.
type PageLimit struct {
	ListTaskOutput                 int // default 100, max 100
	DescribeTaskInput              int // max 100
	ListServiceOutput              int // default 10, max 100
	DescribeServiceInput           int // max 10
	DescribeContainerInstanceInput int // max 100
	DescribeInstanceOutput         int // max 1000
}

func DefaultPageLimit() PageLimit { _ = "STUB: not implemented"; return *new(PageLimit) }

// Cluster implements both ECS and EC2 API for a single cluster.
type Cluster struct {
	name                  string                              // optional
	definitions           map[string]*ecstypes.TaskDefinition // key is task definition arn
	taskMap               map[string]ecstypes.Task            // key is task arn
	taskList              []ecstypes.Task
	containerInstanceMap  map[string]ecstypes.ContainerInstance // key is container instance arn
	containerInstanceList []ecstypes.ContainerInstance
	ec2Map                map[string]ec2types.Instance // key is instance id
	ec2List               []ec2types.Instance
	serviceMap            map[string]ecstypes.Service
	serviceList           []ecstypes.Service
	limit                 PageLimit
	stats                 ClusterStats
}

// NewCluster creates a mock ECS cluster with default limits.
func NewCluster() *Cluster { _ = "STUB: not implemented"; return nil }

// NewClusterWithName creates a cluster that checks for cluster name if request includes a non empty cluster name.
func NewClusterWithName(name string) *Cluster { _ = "STUB: not implemented"; return nil }

// NOTE: we don't set the maps by design, they should be injected and API calls
// without setting up data should just panic.

// APIStat keep track of individual API calls.
type APIStat struct {
	Called int
	Error  int
}

// ClusterStats keep track of API methods for one ECS cluster.
// Not all methods are tracked
type ClusterStats struct {
	DescribeTaskDefinition APIStat
}

// API Start

func (c *Cluster) Stats() ClusterStats { _ = "STUB: not implemented"; return *new(ClusterStats) }

func (c *Cluster) ListTasks(_ context.Context, input *ecs.ListTasksInput, _ ...func(*ecs.Options)) (*ecs.ListTasksOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Cluster) DescribeTasks(_ context.Context, input *ecs.DescribeTasksInput, _ ...func(*ecs.Options)) (*ecs.DescribeTasksOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Cluster) DescribeTaskDefinition(_ context.Context, input *ecs.DescribeTaskDefinitionInput, _ ...func(*ecs.Options)) (*ecs.DescribeTaskDefinitionOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Cluster) DescribeContainerInstances(_ context.Context, input *ecs.DescribeContainerInstancesInput, _ ...func(*ecs.Options)) (*ecs.DescribeContainerInstancesOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DescribeInstances supports get all the instances and get instance by ids.
// It does NOT support filter. Result always has a single reservation, which is not the case in actual EC2 API.
func (c *Cluster) DescribeInstances(_ context.Context, input *ec2.DescribeInstancesInput, _ ...func(*ec2.Options)) (*ec2.DescribeInstancesOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Cluster) ListServices(_ context.Context, input *ecs.ListServicesInput, _ ...func(*ecs.Options)) (*ecs.ListServicesOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Cluster) DescribeServices(_ context.Context, input *ecs.DescribeServicesInput, _ ...func(*ecs.Options)) (*ecs.DescribeServicesOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// API End

// Hook Start

// SetTasks update both list and map.
func (c *Cluster) SetTasks(tasks []ecstypes.Task) { _ = "STUB: not implemented"; return }

// SetTaskDefinitions updates the map.
// NOTE: we could have both list and map, but we are not using list task def in ecsobserver.
func (c *Cluster) SetTaskDefinitions(defs []*ecstypes.TaskDefinition) {
	_ = "STUB: not implemented"
	return
}

// SetContainerInstances updates the list and map.
func (c *Cluster) SetContainerInstances(instances []ecstypes.ContainerInstance) {
	_ = "STUB: not implemented"
	return
}

// SetEc2Instances updates the list and map.
func (c *Cluster) SetEc2Instances(instances []ec2types.Instance) { _ = "STUB: not implemented"; return }

// SetServices updates the list and map.
func (c *Cluster) SetServices(services []ecstypes.Service) { _ = "STUB: not implemented"; return }

// Hook End

// Generator Start

// GenTasks returns tasks with TaskArn set to arnPrefix+offset, where offset is [0, count).
func GenTasks(arnPrefix string, count int, modifier func(i int, task *ecstypes.Task)) []ecstypes.Task {
	_ = "STUB: not implemented"
	return nil
}

// GenTaskDefinitions returns tasks with TaskArn set to arnPrefix+offset+version, where offset is [0, count).
// e.g. foo0:1, foo1:1 the `:` is following the task family version syntax.
func GenTaskDefinitions(arnPrefix string, count, version int, modifier func(i int, def *ecstypes.TaskDefinition)) []*ecstypes.TaskDefinition {
	_ = "STUB: not implemented"
	return nil
}

func GenContainerInstances(arnPrefix string, count int, modifier func(i int, ci *ecstypes.ContainerInstance)) []ecstypes.ContainerInstance {
	_ = "STUB: not implemented"
	return nil
}

func GenEc2Instances(idPrefix string, count int, modifier func(i int, ins *ec2types.Instance)) []ec2types.Instance {
	_ = "STUB: not implemented"
	return nil
}

func GenServices(arnPrefix string, count int, modifier func(i int, s *ecstypes.Service)) []ecstypes.Service {
	_ = "STUB: not implemented"
	return nil
}

func checkCluster(reqClusterName *string, mockClusterName string) error {
	_ = "STUB: not implemented"
	return nil
}

// pagination Start

type pageInput struct {
	nextToken *string
	size      int
	limit     int
}

type pageOutput struct {
	start     int
	end       int
	nextToken *string
}

// getPage returns new page offset based on existing one.
// It is not using the actual AWS token format, it simply uses number string to keep track of offset.
func getPage(p pageInput) (*pageOutput, error) { _ = "STUB: not implemented"; return nil, nil }

// pagination Emd

// 'generic' Start

// getArns is used by both ListTasks and ListServices
func getArns(items any, arnGetter func(i int) string) []string {
	_ = "STUB: not implemented"
	return nil
}

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ecsobserver // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/ecsobserver"

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	ec2types "github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	ecstypes "github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/hashicorp/golang-lru/v2/simplelru"
	"go.uber.org/zap"
)

const (
	// ECS Service Quota: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/service-quotas.html
	taskDefCacheSize = 2000
	// Based on existing number from cloudwatch-agent
	ec2CacheSize                   = 2000
	describeContainerInstanceLimit = 100
	describeServiceLimit           = 10
	// NOTE: these constants are not defined in go sdk, there are three values for deployment status.
	deploymentStatusActive  = "ACTIVE"
	deploymentStatusPrimary = "PRIMARY"
)

// ecsClient includes API required by taskFetcher.
type ecsClient interface {
	ListTasks(ctx context.Context, params *ecs.ListTasksInput, optFns ...func(*ecs.Options)) (*ecs.ListTasksOutput, error)
	DescribeTasks(ctx context.Context, params *ecs.DescribeTasksInput, optFns ...func(*ecs.Options)) (*ecs.DescribeTasksOutput, error)
	DescribeTaskDefinition(ctx context.Context, params *ecs.DescribeTaskDefinitionInput, optFns ...func(*ecs.Options)) (*ecs.DescribeTaskDefinitionOutput, error)
	DescribeContainerInstances(ctx context.Context, params *ecs.DescribeContainerInstancesInput, optFns ...func(*ecs.Options)) (*ecs.DescribeContainerInstancesOutput, error)
	ListServices(ctx context.Context, params *ecs.ListServicesInput, optFns ...func(*ecs.Options)) (*ecs.ListServicesOutput, error)
	DescribeServices(ctx context.Context, params *ecs.DescribeServicesInput, optFns ...func(*ecs.Options)) (*ecs.DescribeServicesOutput, error)
}

// ec2Client includes API required by TaskFetcher.
type ec2Client interface {
	DescribeInstances(ctx context.Context, params *ec2.DescribeInstancesInput, optFns ...func(*ec2.Options)) (*ec2.DescribeInstancesOutput, error)
}

type taskFetcher struct {
	logger            *zap.Logger
	ecs               ecsClient
	ec2               ec2Client
	cluster           string
	taskDefCache      *simplelru.LRU[string, *ecstypes.TaskDefinition]
	ec2Cache          *simplelru.LRU[string, *ec2types.Instance]
	serviceNameFilter serviceNameFilter
}

type taskFetcherOptions struct {
	Logger            *zap.Logger
	Cluster           string
	Region            string
	serviceNameFilter serviceNameFilter

	// test overrides
	ecsOverride ecsClient
	ec2Override ec2Client
}

func newTaskFetcherFromConfig(ctx context.Context, cfg Config, logger *zap.Logger) (*taskFetcher, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newTaskFetcher(ctx context.Context, opts taskFetcherOptions) (*taskFetcher, error) {
	_ = "STUB: not implemented"
	// Init cache
	return nil, nil
}

// Even if user didn't specify any service related config, we still generates a valid filter
// that matches nothing. See service.go serviceConfigsToFilter.

// Return early if any clients are mocked, caller should overrides all the clients when mocking.

func (f *taskFetcher) fetchAndDecorate(ctx context.Context) ([]*taskAnnotated, error) {
	_ = "STUB: not implemented"
	// taskAnnotated
	return nil, nil
}

// EC2

// Services

// getDiscoverableTasks get arns of all running tasks and describe those tasks
// and filter only fargate tasks or EC2 task which container instance is known.
// There is no API to list task detail without arn so we need to call two APIs.
func (f *taskFetcher) getDiscoverableTasks(ctx context.Context) ([]*ecstypes.Task, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NOTE: the limit for list task response and describe task request are both 100.

// Preserve only fargate tasks or EC2 tasks with non-nil ContainerInstanceArn.
// When ECS task of EC2 launch type is in state Provisioning/Pending, it may
// not have EC2 instance. Such tasks have `nil` instance arn and the
// attachContainerInstance call will fail

// attachTaskDefinition converts ecs.Task into a taskAnnotated to include its ecs.TaskDefinition.
func (f *taskFetcher) attachTaskDefinition(ctx context.Context, tasks []*ecstypes.Task) ([]*taskAnnotated, error) {
	_ = "STUB: not implemented"

	// key is task definition arn
	return nil, nil
}

// attachContainerInstance fetches all the container instances' underlying EC2 vms
// and attach EC2 info to tasks.
func (f *taskFetcher) attachContainerInstance(ctx context.Context, tasks []*taskAnnotated) error {
	_ = "STUB: not implemented"
	// Map container instance to EC2, key is container instance id.
	return nil
}

// Only EC2 instance type need to fetch EC2 info

// All fargate, skip

// Describe container instances that do not have cached EC2 info.

// use value from cache

// DescribeContainerInstance size limit is 100, do it in batch.

// Assign the info back to task

// NOTE: we need to skip fargate here because we are looping all tasks again.

// Update the cache

// Run ecs.DescribeContainerInstances and ec2.DescribeInstances for a batch (less than 100 container instances).
func (f *taskFetcher) describeContainerInstances(ctx context.Context, instanceList []string,
	ci2EC2 map[string]*ec2types.Instance,
) error {
	_ = "STUB: not implemented"
	// Get container instances
	return nil
}

// Create the index to map ec2 id back to container instance id.

// Fetch all ec2 instances and update mapping from container instance id to ec2 info.
// NOTE: because the limit on ec2 is 1000, much larger than ecs container instance's 100,
// we don't do paging logic here.

// update mapping

// serviceNameFilter decides if we should get detail info for a service, i.e. make the describe API call.
type serviceNameFilter func(name string) bool

// getAllServices does not have cache like task definition or ec2 instances
// because we need to get the deployment id to map service to task, which changes frequently.
func (f *taskFetcher) getAllServices(ctx context.Context) ([]ecstypes.Service, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// List and filter out services we need to describe.

// DescribeServices size limit is 10 so we need to do paging on client side.

// attachService map service to task using deployment id.
// Each service can have multiple deployment and each task keep track of the deployment in task.StartedBy.
func (*taskFetcher) attachService(tasks []*taskAnnotated, services []ecstypes.Service) {
	_ = "STUB: not implemented"
	// Map deployment ID to service name
	return
}

// Attach service to task

// taskAnnotated is created using RunTask i.e. not managed by a service.

// Service not found happen a lot because we only fetch services defined in ServiceConfig.
// However, we fetch all the tasks, which could be started by other services no mentioned in config
// or started using RunTasks API directly.

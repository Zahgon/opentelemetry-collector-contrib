// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package ecsobserver // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/ecsobserver"

import (
	ec2types "github.com/aws/aws-sdk-go-v2/service/ec2/types"
	ecstypes "github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"go.uber.org/zap"
)

// taskAnnotated contains both raw task info and its definition.
// It is generated from taskFetcher.
type taskAnnotated struct {
	Task       *ecstypes.Task
	Definition *ecstypes.TaskDefinition
	EC2        *ec2types.Instance
	Service    *ecstypes.Service
	Matched    []matchedContainer
}

// AddMatchedContainer tries to add a new matched container.
// If the container already exists will merge targets within one container (i.e. different port/metrics path).
func (t *taskAnnotated) AddMatchedContainer(newContainer matchedContainer) {
	_ = "STUB: not implemented"
	return
}

func (t *taskAnnotated) TaskTags() map[string]string { _ = "STUB: not implemented"; return nil }

// EC2Tags returns ec2 instance tags as it is. Sanitize to prometheus label format is done during export.
// NOTE: the tag to string conversion is duplicated because the Tag struct is defined in each service's own API package.
// i.e. services don't import a common package that includes tag definition.
func (t *taskAnnotated) EC2Tags() map[string]string { _ = "STUB: not implemented"; return nil }

func (t *taskAnnotated) ContainerLabels(containerIndex int) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

// errPrivateIPNotFound indicates the awsvpc private ip or EC2 instance ip is not found.
type errPrivateIPNotFound struct {
	TaskArn     string
	NetworkMode ecstypes.NetworkMode
	Extra       string // extra message
}

func (e *errPrivateIPNotFound) Error() string { _ = "STUB: not implemented"; return "" }

func (*errPrivateIPNotFound) message() string { _ = "STUB: not implemented"; return "" }

func (e *errPrivateIPNotFound) zapFields() []zap.Field { _ = "STUB: not implemented"; return nil }

// PrivateIP returns private ip address based on network mode.
// EC2 launch type can use host/bridge mode and the private ip is the EC2 instance's ip.
// awsvpc has its own ip regardless of launch type.
func (t *taskAnnotated) PrivateIP() (string, error) { _ = "STUB: not implemented"; return "", nil }

// When network mode is empty and launch type is EC2, ECS uses bridge network.
// For fargate it has to be awsvpc and will error on task creation if invalid config is given.
// See https://docs.aws.amazon.com/AmazonECS/latest/userguide/fargate-task-defs.html#fargate-tasks-networkmod
// In another word, when network mode is empty, it must be EC2 bridge.

// errMappedPortNotFound indicates the port specified in config does not exists
// or the location for mapped ports has changed on ECS side.
type errMappedPortNotFound struct {
	TaskArn       string
	NetworkMode   ecstypes.NetworkMode
	ContainerName string
	ContainerPort int32
}

func (e *errMappedPortNotFound) Error() string {
	_ = "STUB: not implemented"
	// Output the error message in this order to make searching easier as only task arn changes frequently.
	// %q for network mode because empty string is valid for ECS EC2.
	return ""
}

func (*errMappedPortNotFound) message() string { _ = "STUB: not implemented"; return "" }

func (e *errMappedPortNotFound) zapFields() []zap.Field { _ = "STUB: not implemented"; return nil }

// MappedPort returns 'external' port based on network mode.
// EC2 bridge uses a random host port while EC2 host/awsvpc uses a port specified by the user.
func (t *taskAnnotated) MappedPort(def ecstypes.ContainerDefinition, containerPort int32) (int32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// the error is same for all network modes (if any)

// taskDefinition->containerDefinitions->portMappings

//  task->containers->networkBindings

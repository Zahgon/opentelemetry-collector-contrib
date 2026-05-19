// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package eks // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/metadataproviders/aws/eks"

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"k8s.io/client-go/kubernetes"
)

const (
	clusterNameAwsEksTag     = "aws:eks:cluster-name"
	clusterNameEksTag        = "eks:cluster-name"
	kubernetesClusterNameTag = "kubernetes.io/cluster/"
)

type Provider interface {
	ClusterVersion() (string, error)
	OIDCIssuer(ctx context.Context) (string, error)
	GetK8sInstanceMetadata(ctx context.Context) (InstanceMetadata, error)
	GetInstanceMetadata(ctx context.Context) (InstanceMetadata, error)
	SetRegionInstanceID(region, instanceID string)
}

type ec2Client interface {
	DescribeInstances(ctx context.Context, params *ec2.DescribeInstancesInput, optFns ...func(*ec2.Options)) (*ec2.DescribeInstancesOutput, error)
}

type metadataClient struct {
	instanceMetadata InstanceMetadata
	clientset        kubernetes.Interface
	ec2Client        ec2Client
	nodeName         string
}

type InstanceMetadata struct {
	AvailabilityZone string
	AccountID        string
	InstanceID       string
	Region           string
	ClusterName      string
	ImageID          string
	InstanceType     string
	Hostname         string
}

var _ Provider = (*metadataClient)(nil)

func NewProvider(cfg aws.Config, nodeName string) (Provider, error) {
	_ = "STUB: not implemented"
	return *new(Provider), nil
}

func (c *metadataClient) ClusterVersion() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (c *metadataClient) OIDCIssuer(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// GetK8sInstanceMetadata retrieves region, instanceID, and availabilityZone attributes from the K8s node's providerID.
// It requires the node name which is passed to the constructor.
// If region or instanceID are not found, it returns an error as they are mandatory to query ec2 API for the full metadata.
func (c *metadataClient) GetK8sInstanceMetadata(ctx context.Context) (InstanceMetadata, error) {
	_ = "STUB: not implemented"
	return *new(InstanceMetadata), nil
}

func parseRegionAndInstanceID(input string) (string, string, string) {
	_ = "STUB: not implemented"
	// Example mockProviderID: "aws:///us-west-2a/i-049ca2df511bec762"
	return "", "", ""
}

// Prevent NPE and return empty strings if the format is invalid

// Remove the last character (zone letter)

// GetInstanceMetadata retrieves detailed instance metadata from AWS EC2.
// It requires region and InstanceID to be set, if not set, it calls GetK8sInstanceMetadata to set them
func (c *metadataClient) GetInstanceMetadata(ctx context.Context) (InstanceMetadata, error) {
	_ = "STUB: not implemented"
	return *new(InstanceMetadata), nil
}

func getClusterNameTagFromReservations(reservations []types.Reservation) string {
	_ = "STUB: not implemented"
	return ""
}

// SetRegionInstanceID sets the region and instance ID.
// This method is useful when caller has the region and instanceID and wants to fetch the ec2 metadata
// without calling the kube api through GetK8sInstanceMetadata first.
func (c *metadataClient) SetRegionInstanceID(region, instanceID string) {
	_ = "STUB: not implemented"
	return
}

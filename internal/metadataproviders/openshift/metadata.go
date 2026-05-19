// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package openshift // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/metadataproviders/openshift"

import (
	"context"
	"crypto/tls"
	"net/http"
)

// Provider gets cluster metadata from Openshift.
type Provider interface {
	K8SClusterVersion(context.Context) (string, error)
	OpenShiftClusterVersion(context.Context) (string, error)
	Infrastructure(context.Context) (*InfrastructureAPIResponse, error)
}

// NewProvider creates a new metadata provider.
func NewProvider(address, token string, tlsCfg *tls.Config) Provider {
	_ = "STUB: not implemented"
	return *new(Provider)
}

type openshiftProvider struct {
	client  *http.Client
	address string
	token   string
}

func (o *openshiftProvider) makeOCPRequest(ctx context.Context, endpoint, target string) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// OpenShiftClusterVersion requests the ClusterVersion from the openshift api.
func (o *openshiftProvider) OpenShiftClusterVersion(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// ClusterVersion requests Infrastructure details from the openshift api.
func (o *openshiftProvider) Infrastructure(ctx context.Context) (*InfrastructureAPIResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// K8SClusterVersion requests the ClusterVersion from the kubernetes api.
func (o *openshiftProvider) K8SClusterVersion(ctx context.Context) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

type ocpClusterVersionAPIResponse struct {
	Status struct {
		Desired struct {
			Version string `json:"version"`
		} `json:"desired"`
	} `json:"status"`
}

// InfrastructureAPIResponse from OpenShift API.
type InfrastructureAPIResponse struct {
	Status InfrastructureStatus `json:"status"`
}

// InfrastructureStatus holds cluster-wide information about Infrastructure.
// https://docs.openshift.com/container-platform/4.11/rest_api/config_apis/infrastructure-config-openshift-io-v1.html#apisconfig-openshift-iov1infrastructuresnamestatus
type InfrastructureStatus struct {
	// ControlPlaneTopology expresses the expectations for operands that normally
	// run on control nodes. The default is 'HighlyAvailable', which represents
	// the behavior operators have in a "normal" cluster. The 'SingleReplica' mode
	// will be used in single-node deployments and the operators should not
	// configure the operand for highly-available operation The 'External' mode
	// indicates that the control plane is hosted externally to the cluster and
	// that its components are not visible within the cluster.
	ControlPlaneTopology string `json:"controlPlaneTopology"`
	// InfrastructureName uniquely identifies a cluster with a human friendly
	// name. Once set it should not be changed. Must be of max length 27 and must
	// have only alphanumeric or hyphen characters.
	InfrastructureName string `json:"infrastructureName"`
	// InfrastructureTopology expresses the expectations for infrastructure
	// services that do not run on control plane nodes, usually indicated by a
	// node selector for a role value other than master. The default is
	// 'HighlyAvailable', which represents the behavior operators have in a
	// "normal" cluster. The 'SingleReplica' mode will be used in single-node
	// deployments and the operators should not configure the operand for
	// highly-available operation.
	InfrastructureTopology string `json:"infrastructureTopology"`
	// PlatformStatus holds status information specific to the underlying
	// infrastructure provider.
	PlatformStatus InfrastructurePlatformStatus `json:"platformStatus"`
}

// InfrastructurePlatformStatus reported by the OpenShift API.
type InfrastructurePlatformStatus struct {
	Aws       InfrastructureStatusAWS       `json:"aws"`
	Azure     InfrastructureStatusAzure     `json:"azure"`
	Baremetal struct{}                      `json:"baremetal"`
	GCP       InfrastructureStatusGCP       `json:"gcp"`
	IBMCloud  InfrastructureStatusIBMCloud  `json:"ibmcloud"`
	OpenStack InfrastructureStatusOpenStack `json:"openstack"`
	OVirt     struct{}                      `json:"ovirt"`
	VSphere   struct{}                      `json:"vsphere"`
	Type      string                        `json:"type"`
}

// InfrastructureStatusAWS reported by the OpenShift API.
type InfrastructureStatusAWS struct {
	// Region holds the default AWS region for new AWS resources created by the
	// cluster.
	Region string `json:"region"`
}

// InfrastructureStatusAzure reported by the OpenShift API.
type InfrastructureStatusAzure struct {
	// CloudName is the name of the Azure cloud environment which can be used to
	// configure the Azure SDK with the appropriate Azure API endpoints. If empty,
	// the value is equal to AzurePublicCloud.
	CloudName string `json:"cloudName"`
}

// InfrastructureStatusGCP reported by the OpenShift API.
type InfrastructureStatusGCP struct {
	// Region holds the region for new GCP resources created for the cluster.
	Region string `json:"region"`
}

// InfrastructureStatusIBMCloud reported by the OpenShift API.
type InfrastructureStatusIBMCloud struct {
	// Location is where the cluster has been deployed.
	Location string `json:"location"`
}

// InfrastructureStatusOpenStack reported by the OpenShift API.
type InfrastructureStatusOpenStack struct {
	// CloudName is the name of the desired OpenStack cloud in the client
	// configuration file (clouds.yaml).
	CloudName string `json:"cloudName"`
}

// k8sClusterVersionAPIResponse of OpenShift.
// https://docs.openshift.com/container-platform/4.11/rest_api/config_apis/clusterversion-config-openshift-io-v1.html#apisconfig-openshift-iov1clusterversionsnamestatus
type k8sClusterVersionAPIResponse struct {
	GitVersion string `json:"gitVersion"`
}

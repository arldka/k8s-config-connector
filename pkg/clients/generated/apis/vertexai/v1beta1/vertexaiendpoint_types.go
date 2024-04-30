// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Config Connector and manual
//     changes will be clobbered when the file is regenerated.
//
// ----------------------------------------------------------------------------

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

package v1beta1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type EndpointEncryptionSpec struct {
	/* Required. The Cloud KMS resource identifier of the customer managed encryption key used to protect a resource.
	Has the form: projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key.
	The key needs to be in the same region as where the compute resource is created. */
	KmsKeyNameRef v1alpha1.ResourceRef `json:"kmsKeyNameRef"`
}

type VertexAIEndpointSpec struct {
	/* The description of the Endpoint. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* Required. The display name of the Endpoint. The name can be up to 128 characters long and can consist of any UTF-8 characters. */
	DisplayName string `json:"displayName"`

	/* Immutable. Customer-managed encryption key spec for an Endpoint. If set, this Endpoint and all sub-resources of this Endpoint will be secured by this key. */
	// +optional
	EncryptionSpec *EndpointEncryptionSpec `json:"encryptionSpec,omitempty"`

	/* Optional. The full name of the Google Compute Engine network to which the Endpoint should be peered.
	Private services access must already be configured for the network. If left unspecified, the Endpoint is not peered with any network.
	Only one of the fields, network or enablePrivateServiceConnect, can be set.
	Format: projects/{project_id}/global/networks/{network_name}. */
	// +optional
	NetworkRef *v1alpha1.ResourceRef `json:"networkRef,omitempty"`

	/* The project that this resource belongs to. */
	ProjectRef v1alpha1.ResourceRef `json:"projectRef"`

	/* Immutable. The region for the resource. */
	Region string `json:"region"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`
}

type EndpointObservedStateStatus struct {
	/* Output only. Timestamp when this Endpoint was created. */
	// +optional
	CreateTime *string `json:"createTime,omitempty"`

	/* Output only. Resource name of the Model Monitoring job associated with this Endpoint if monitoring is enabled by CreateModelDeploymentMonitoringJob. Format: 'projects/{project}/locations/{location}/modelDeploymentMonitoringJobs/{model_deployment_monitoring_job}'. */
	// +optional
	ModelDeploymentMonitoringJob *string `json:"modelDeploymentMonitoringJob,omitempty"`
}

type VertexAIEndpointStatus struct {
	/* Conditions represent the latest available observations of the
	   VertexAIEndpoint's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`

	/* The observed state of the underlying GCP resource. */
	// +optional
	ObservedState *EndpointObservedStateStatus `json:"observedState,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpvertexaiendpoint;gcpvertexaiendpoints
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/stability-level=stable";"cnrm.cloud.google.com/system=true";"cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VertexAIEndpoint is the Schema for the vertexai API
// +k8s:openapi-gen=true
type VertexAIEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VertexAIEndpointSpec   `json:"spec,omitempty"`
	Status VertexAIEndpointStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VertexAIEndpointList contains a list of VertexAIEndpoint
type VertexAIEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VertexAIEndpoint `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VertexAIEndpoint{}, &VertexAIEndpointList{})
}

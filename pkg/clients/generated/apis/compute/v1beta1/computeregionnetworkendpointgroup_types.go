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

type RegionnetworkendpointgroupCloudFunction struct {
	/* Immutable. A user-defined name of the Cloud Function.
	The function name is case-sensitive and must be 1-63 characters long.
	Example value: "func1". */
	// +optional
	FunctionRef *v1alpha1.ResourceRef `json:"functionRef,omitempty"`

	/* Immutable. A template to parse function field from a request URL. URL mask allows
	for routing to multiple Cloud Functions without having to create
	multiple Network Endpoint Groups and backend services.

	For example, request URLs "mydomain.com/function1" and "mydomain.com/function2"
	can be backed by the same Serverless NEG with URL mask "/". The URL mask
	will parse them to { function = "function1" } and { function = "function2" } respectively. */
	// +optional
	UrlMask *string `json:"urlMask,omitempty"`
}

type RegionnetworkendpointgroupCloudRun struct {
	/* Immutable. Cloud Run service is the main resource of Cloud Run.
	The service must be 1-63 characters long, and comply with RFC1035.
	Example value: "run-service". */
	// +optional
	ServiceRef *v1alpha1.ResourceRef `json:"serviceRef,omitempty"`

	/* Immutable. Cloud Run tag represents the "named-revision" to provide
	additional fine-grained traffic routing information.
	The tag must be 1-63 characters long, and comply with RFC1035.
	Example value: "revision-0010". */
	// +optional
	Tag *string `json:"tag,omitempty"`

	/* Immutable. A template to parse service and tag fields from a request URL.
	URL mask allows for routing to multiple Run services without having
	to create multiple network endpoint groups and backend services.

	For example, request URLs "foo1.domain.com/bar1" and "foo1.domain.com/bar2"
	an be backed by the same Serverless Network Endpoint Group (NEG) with
	URL mask ".domain.com/". The URL mask will parse them to { service="bar1", tag="foo1" }
	and { service="bar2", tag="foo2" } respectively. */
	// +optional
	UrlMask *string `json:"urlMask,omitempty"`
}

type ComputeRegionNetworkEndpointGroupSpec struct {
	/* Immutable. Only valid when networkEndpointType is "SERVERLESS".
	Only one of cloud_run, app_engine, cloud_function or serverless_deployment may be set. */
	// +optional
	CloudFunction *RegionnetworkendpointgroupCloudFunction `json:"cloudFunction,omitempty"`

	/* Immutable. Only valid when networkEndpointType is "SERVERLESS".
	Only one of cloud_run, app_engine, cloud_function or serverless_deployment may be set. */
	// +optional
	CloudRun *RegionnetworkendpointgroupCloudRun `json:"cloudRun,omitempty"`

	/* Immutable. An optional description of this resource. Provide this property when
	you create the resource. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* Immutable. Type of network endpoints in this network endpoint group. Defaults to SERVERLESS Default value: "SERVERLESS" Possible values: ["SERVERLESS", "PRIVATE_SERVICE_CONNECT"]. */
	// +optional
	NetworkEndpointType *string `json:"networkEndpointType,omitempty"`

	/* Immutable. This field is only used for PSC.
	The URL of the network to which all network endpoints in the NEG belong. Uses
	"default" project network if unspecified. */
	// +optional
	NetworkRef *v1alpha1.ResourceRef `json:"networkRef,omitempty"`

	/* Immutable. The target service url used to set up private service connection to
	a Google API or a PSC Producer Service Attachment. */
	// +optional
	PscTargetService *string `json:"pscTargetService,omitempty"`

	/* Immutable. A reference to the region where the Serverless NEGs Reside. */
	Region string `json:"region"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* Immutable. This field is only used for PSC.
	Optional URL of the subnetwork to which all network endpoints in the NEG belong. */
	// +optional
	SubnetworkRef *v1alpha1.ResourceRef `json:"subnetworkRef,omitempty"`
}

type ComputeRegionNetworkEndpointGroupStatus struct {
	/* Conditions represent the latest available observations of the
	   ComputeRegionNetworkEndpointGroup's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`

	// +optional
	SelfLink *string `json:"selfLink,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ComputeRegionNetworkEndpointGroup is the Schema for the compute API
// +k8s:openapi-gen=true
type ComputeRegionNetworkEndpointGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ComputeRegionNetworkEndpointGroupSpec   `json:"spec,omitempty"`
	Status ComputeRegionNetworkEndpointGroupStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ComputeRegionNetworkEndpointGroupList contains a list of ComputeRegionNetworkEndpointGroup
type ComputeRegionNetworkEndpointGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeRegionNetworkEndpointGroup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeRegionNetworkEndpointGroup{}, &ComputeRegionNetworkEndpointGroupList{})
}

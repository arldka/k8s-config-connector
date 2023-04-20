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

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ApiconfigBackendConfig struct {
	/* Immutable. Google Cloud IAM service account used to sign OIDC tokens for backends that have authentication configured
	(https://cloud.google.com/service-infrastructure/docs/service-management/reference/rest/v1/services.configs#backend). */
	GoogleServiceAccount string `json:"googleServiceAccount"`
}

type ApiconfigDocument struct {
	/* Immutable. Base64 encoded content of the file. */
	Contents string `json:"contents"`

	/* Immutable. The file path (full or relative path). This is typically the path of the file when it is uploaded. */
	Path string `json:"path"`
}

type ApiconfigFileDescriptorSet struct {
	/* Immutable. Base64 encoded content of the file. */
	Contents string `json:"contents"`

	/* Immutable. The file path (full or relative path). This is typically the path of the file when it is uploaded. */
	Path string `json:"path"`
}

type ApiconfigGatewayConfig struct {
	/* Backend settings that are applied to all backends of the Gateway. */
	BackendConfig ApiconfigBackendConfig `json:"backendConfig"`
}

type ApiconfigGrpcServices struct {
	/* Immutable. Input only. File descriptor set, generated by protoc.
	To generate, use protoc with imports and source info included. For an example test.proto file, the following command would put the value in a new file named out.pb.

	$ protoc --include_imports --include_source_info test.proto -o out.pb. */
	FileDescriptorSet ApiconfigFileDescriptorSet `json:"fileDescriptorSet"`

	/* Uncompiled proto files associated with the descriptor set, used for display purposes (server-side compilation is not supported). These should match the inputs to 'protoc' command used to generate fileDescriptorSet. */
	// +optional
	Source []ApiconfigSource `json:"source,omitempty"`
}

type ApiconfigManagedServiceConfigs struct {
	/* Immutable. Base64 encoded content of the file. */
	Contents string `json:"contents"`

	/* Immutable. The file path (full or relative path). This is typically the path of the file when it is uploaded. */
	Path string `json:"path"`
}

type ApiconfigOpenapiDocuments struct {
	/* The OpenAPI Specification document file. */
	Document ApiconfigDocument `json:"document"`
}

type ApiconfigSource struct {
	/* Immutable. Base64 encoded content of the file. */
	Contents string `json:"contents"`

	/* Immutable. The file path (full or relative path). This is typically the path of the file when it is uploaded. */
	Path string `json:"path"`
}

type APIGatewayAPIConfigSpec struct {
	/* Immutable. The API to attach the config to. */
	Api string `json:"api"`

	/* Immutable. Creates a unique name beginning with the specified prefix. If this and api_config_id are unspecified, a random value is chosen for the name. */
	// +optional
	ApiConfigIdPrefix *string `json:"apiConfigIdPrefix,omitempty"`

	/* A user-visible name for the API. */
	// +optional
	DisplayName *string `json:"displayName,omitempty"`

	/* Immutable. Immutable. Gateway specific configuration.
	If not specified, backend authentication will be set to use OIDC authentication using the default compute service account. */
	// +optional
	GatewayConfig *ApiconfigGatewayConfig `json:"gatewayConfig,omitempty"`

	/* gRPC service definition files. If specified, openapiDocuments must not be included. */
	// +optional
	GrpcServices []ApiconfigGrpcServices `json:"grpcServices,omitempty"`

	/* Optional. Service Configuration files. At least one must be included when using gRPC service definitions. See https://cloud.google.com/endpoints/docs/grpc/grpc-service-config#service_configuration_overview for the expected file contents.
	If multiple files are specified, the files are merged with the following rules: * All singular scalar fields are merged using "last one wins" semantics in the order of the files uploaded. * Repeated fields are concatenated. * Singular embedded messages are merged using these rules for nested fields. */
	// +optional
	ManagedServiceConfigs []ApiconfigManagedServiceConfigs `json:"managedServiceConfigs,omitempty"`

	/* OpenAPI specification documents. If specified, grpcServices and managedServiceConfigs must not be included. */
	// +optional
	OpenapiDocuments []ApiconfigOpenapiDocuments `json:"openapiDocuments,omitempty"`

	/* The project that this resource belongs to. */
	ProjectRef v1alpha1.ResourceRef `json:"projectRef"`

	/* Immutable. Optional. The apiConfigId of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`
}

type APIGatewayAPIConfigStatus struct {
	/* Conditions represent the latest available observations of the
	   APIGatewayAPIConfig's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* The resource name of the API Config. */
	// +optional
	Name *string `json:"name,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`

	/* The ID of the associated Service Config (https://cloud.google.com/service-infrastructure/docs/glossary#config). */
	// +optional
	ServiceConfigId *string `json:"serviceConfigId,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// APIGatewayAPIConfig is the Schema for the apigateway API
// +k8s:openapi-gen=true
type APIGatewayAPIConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   APIGatewayAPIConfigSpec   `json:"spec,omitempty"`
	Status APIGatewayAPIConfigStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// APIGatewayAPIConfigList contains a list of APIGatewayAPIConfig
type APIGatewayAPIConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []APIGatewayAPIConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&APIGatewayAPIConfig{}, &APIGatewayAPIConfigList{})
}
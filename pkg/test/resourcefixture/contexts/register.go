// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package contexts

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/core/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller"
	dclcontroller "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/dcl"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/logging"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl"
	dclconversion "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/conversion"
	dclextension "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/extension"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/kcclite"
	dcllivestate "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/livestate"
	dclmetadata "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/metadata"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/schema/dclschemaloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/krmtotf"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/resourceoverrides"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/resourcefixture"
	"github.com/nasa9084/go-openapi"

	mmdcl "github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	dclunstruct "github.com/GoogleCloudPlatform/declarative-resource-client-library/unstructured"
	tfschema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type ResourceContext struct {
	ResourceGVK        schema.GroupVersionKind
	ResourceKind       string
	SkipNoChange       bool
	SkipUpdate         bool
	SkipDriftDetection bool
	SkipDelete         bool

	// fields related to DCL-based resources
	DCLSchema *openapi.Schema
	DCLBased  bool
}

var (
	resourceContextMap = map[string]ResourceContext{}
	emptyGVK           = schema.GroupVersionKind{}
)

func GetResourceContext(fixture resourcefixture.ResourceFixture, serviceMetadataLoader dclmetadata.ServiceMetadataLoader, dclSchemaLoader dclschemaloader.DCLSchemaLoader) ResourceContext {
	rc, ok := resourceContextMap[fixture.Name]
	if !ok {
		rc = ResourceContext{
			ResourceGVK:  fixture.GVK,
			ResourceKind: fixture.GVK.Kind,
		}
	}
	if rc.ResourceGVK == emptyGVK {
		rc.ResourceGVK = fixture.GVK
	}
	if dclmetadata.IsDCLBasedResourceKind(rc.ResourceGVK, serviceMetadataLoader) {
		s, err := dclschemaloader.GetDCLSchemaForGVK(rc.ResourceGVK, serviceMetadataLoader, dclSchemaLoader)
		if err != nil {
			panic(fmt.Sprintf("error getting the DCL schema for GVK %v: %v", rc.ResourceGVK, err))
		}
		rc.DCLSchema = s
		rc.DCLBased = true
	}
	return rc
}

func (rc ResourceContext) SupportsLabels(smLoader *servicemappingloader.ServiceMappingLoader) bool {
	if rc.DCLBased {
		_, _, found, err := dclextension.GetLabelsFieldSchema(rc.DCLSchema)
		if err != nil {
			panic(fmt.Errorf("error getting the DCL schema for labels field: %w", err))
		}
		return found
	}
	// For tf based resources, resolve the label info from ResourceConfig
	resourceConfig := rc.getResourceConfig(smLoader)
	return resourceConfig.MetadataMapping.Labels != ""
}

func (rc ResourceContext) getResourceConfig(smLoader *servicemappingloader.ServiceMappingLoader) v1alpha1.ResourceConfig {
	for _, sm := range smLoader.GetServiceMappings() {
		for _, resourceConfig := range sm.Spec.Resources {
			if resourceConfig.Kind == rc.ResourceKind {
				return resourceConfig
			}
		}
	}
	panic(fmt.Errorf("no resource config found for kind: %v", rc.ResourceKind))
}

func (rc ResourceContext) IsDirectResource() bool {
	switch rc.ResourceGVK.GroupKind() {
	case schema.GroupKind{Group: "logging.cnrm.cloud.google.com", Kind: "LoggingLogMetric"}:
		return true
	}
	return false
}

func (rc ResourceContext) IsAutoGenerated(smLoader *servicemappingloader.ServiceMappingLoader) bool {
	if rc.DCLBased {
		return false
	}
	resourceConfig := rc.getResourceConfig(smLoader)
	return resourceConfig.AutoGenerated
}

func (rc ResourceContext) Create(ctx context.Context, _ *testing.T, u *unstructured.Unstructured, provider *tfschema.Provider, c client.Client, smLoader *servicemappingloader.ServiceMappingLoader, config *mmdcl.Config, dclConverter *dclconversion.Converter) (*unstructured.Unstructured, error) {
	if rc.DCLBased {
		return dclCreate(ctx, u, config, c, dclConverter, smLoader)
	}
	return terraformCreate(ctx, u, provider, c, smLoader)
}

func (rc ResourceContext) Get(ctx context.Context, _ *testing.T, u *unstructured.Unstructured, provider *tfschema.Provider, c client.Client, smLoader *servicemappingloader.ServiceMappingLoader, config *mmdcl.Config, dclConverter *dclconversion.Converter, httpClient *http.Client) (*unstructured.Unstructured, error) {
	// direct controllers
	switch u.GroupVersionKind().GroupKind() {
	case schema.GroupKind{Group: "logging.cnrm.cloud.google.com", Kind: "LoggingLogMetric"}:
		m, err := logging.GetModel(ctx, &controller.Config{
			HTTPClient: httpClient,
		})
		if err != nil {
			return nil, err
		}
		a, err := m.AdapterForObject(ctx, c, u)
		if err != nil {
			return nil, err
		}
		found, err := a.Find(ctx)
		if err != nil {
			return nil, err
		}
		if !found {
			return nil, fmt.Errorf("logging.cnrm.cloud.google.com/LoggingLogMetric '%v' is not found", u.GetName())
		}

		unst, err := a.Export(ctx)
		if err != nil {
			return nil, err
		}

		return unst, nil
	}

	if rc.DCLBased {
		return dclGet(ctx, u, config, c, dclConverter, smLoader)
	}
	return terraformGet(ctx, u, provider, c, smLoader)
}

func (rc ResourceContext) Delete(ctx context.Context, _ *testing.T, u *unstructured.Unstructured, provider *tfschema.Provider, c client.Client, smLoader *servicemappingloader.ServiceMappingLoader, config *mmdcl.Config, dclConverter *dclconversion.Converter, httpClient *http.Client) error {
	// direct controllers
	switch u.GroupVersionKind().GroupKind() {
	case schema.GroupKind{Group: "logging.cnrm.cloud.google.com", Kind: "LoggingLogMetric"}:
		m, err := logging.GetModel(ctx, &controller.Config{
			HTTPClient: httpClient,
		})
		if err != nil {
			return err
		}
		a, err := m.AdapterForObject(ctx, c, u)
		if err != nil {
			return err
		}
		found, err := a.Find(ctx)
		if err != nil {
			return err
		}
		if !found {
			return fmt.Errorf("logging.cnrm.cloud.google.com/LoggingLogMetric '%v' not found", u.GetName())
		}

		deleted, err := a.Delete(ctx)
		if err != nil {
			return err
		}
		if !deleted {
			return fmt.Errorf("did not delete logging.cnrm.cloud.google.com/LoggingLogMetric '%v'", u.GetName())
		}
	}
	if rc.DCLBased {
		return dclDelete(ctx, u, config, c, dclConverter, smLoader)
	}
	return terraformDelete(ctx, u, provider, c, smLoader)
}

func terraformDelete(ctx context.Context, u *unstructured.Unstructured, provider *tfschema.Provider, c client.Client, smLoader *servicemappingloader.ServiceMappingLoader) error {
	resource, liveState, err := getTerraformResourceAndLiveState(ctx, u, provider, c, smLoader)
	if err != nil {
		return err
	}
	if liveState.Empty() {
		return fmt.Errorf("resource '%v' of type '%v' cannot be deleted as it does not exist", u.GetName(), u.GroupVersionKind())
	}
	_, diagnostics := resource.TFResource.Apply(ctx, liveState, &terraform.InstanceDiff{Destroy: true}, provider.Meta())
	if err := krmtotf.NewErrorFromDiagnostics(diagnostics); err != nil {
		return fmt.Errorf("error deleting resource: %w", err)
	}
	return err
}

func terraformCreate(ctx context.Context, u *unstructured.Unstructured, provider *tfschema.Provider, c client.Client, smLoader *servicemappingloader.ServiceMappingLoader) (*unstructured.Unstructured, error) {
	resource, liveState, err := getTerraformResourceAndLiveState(ctx, u, provider, c, smLoader)
	if err != nil {
		return nil, err
	}
	if !liveState.Empty() {
		return nil, fmt.Errorf("resource '%v' of type '%v' cannot be created as it already exists", u.GetName(), u.GroupVersionKind())
	}
	config, _, err := krmtotf.KRMResourceToTFResourceConfig(resource, c, smLoader)
	if err != nil {
		return nil, fmt.Errorf("error expanding resource configuration: %w", err)
	}
	diff, err := resource.TFResource.Diff(ctx, liveState, config, provider.Meta())
	if err != nil {
		return nil, fmt.Errorf("error calculating diff: %w", err)
	}
	newState, diagnostics := resource.TFResource.Apply(ctx, liveState, diff, provider.Meta())
	if err := krmtotf.NewErrorFromDiagnostics(diagnostics); err != nil {
		return nil, fmt.Errorf("error applying resource change: %w", err)
	}
	return resourceToKRM(resource, newState)
}

func terraformGet(ctx context.Context, u *unstructured.Unstructured, provider *tfschema.Provider, c client.Client, smLoader *servicemappingloader.ServiceMappingLoader) (*unstructured.Unstructured, error) {
	resource, liveState, err := getTerraformResourceAndLiveState(ctx, u, provider, c, smLoader)
	if err != nil {
		return nil, err
	}
	if liveState.Empty() {
		return nil, fmt.Errorf("resource '%v' of type '%v' is not found", u.GetName(), u.GroupVersionKind())
	}
	return resourceToKRM(resource, liveState)
}

func dclCreate(ctx context.Context, u *unstructured.Unstructured, config *mmdcl.Config, kubeClient client.Client, converter *dclconversion.Converter, serviceMappingLoader *servicemappingloader.ServiceMappingLoader) (*unstructured.Unstructured, error) {
	resource, err := newDCLResource(u, converter)
	if err != nil {
		return nil, err
	}
	liveLite, err := dcllivestate.FetchLiveState(ctx, resource, config, converter, serviceMappingLoader, kubeClient)
	if err != nil {
		return nil, err
	}
	if liveLite != nil {
		return nil, fmt.Errorf("resource '%v' of type '%v' cannot be created as it already exists", u.GetName(), u.GroupVersionKind())
	}
	lite, err := kcclite.ToKCCLite(resource, converter.MetadataLoader, converter.SchemaLoader, serviceMappingLoader, kubeClient)
	if err != nil {
		return nil, fmt.Errorf("error converting KCC full to KCC lite: %w", err)
	}
	dclObj, err := converter.KRMObjectToDCLObject(lite)
	if err != nil {
		return nil, fmt.Errorf("error converting KCC lite to DCL resource: %w", err)
	}
	createdDCLObj, err := dclunstruct.Apply(ctx, config, dclObj, dclcontroller.LifecycleParams...)
	if err != nil {
		return nil, fmt.Errorf("error applying the desired resource: %w", err)
	}
	// get the new state in KCC lite format
	newStateLite, err := converter.DCLObjectToKRMObject(createdDCLObj)
	if err != nil {
		return nil, fmt.Errorf("error converting DCL resource to KCC lite: %w", err)
	}
	return dclStateToKRM(resource, newStateLite, converter.MetadataLoader)
}

func dclGet(ctx context.Context, u *unstructured.Unstructured, config *mmdcl.Config, kubeClient client.Client, converter *dclconversion.Converter, serviceMappingLoader *servicemappingloader.ServiceMappingLoader) (*unstructured.Unstructured, error) {
	resource, err := newDCLResource(u, converter)
	if err != nil {
		return nil, err
	}
	liveLite, err := dcllivestate.FetchLiveState(ctx, resource, config, converter, serviceMappingLoader, kubeClient)
	if err != nil {
		return nil, err
	}
	if liveLite == nil {
		return nil, fmt.Errorf("resource '%v' of type '%v' is not found", u.GetName(), u.GroupVersionKind())
	}
	return dclStateToKRM(resource, liveLite, converter.MetadataLoader)
}

func dclDelete(ctx context.Context, u *unstructured.Unstructured, config *mmdcl.Config, kubeClient client.Client, converter *dclconversion.Converter, serviceMappingLoader *servicemappingloader.ServiceMappingLoader) error {
	resource, err := newDCLResource(u, converter)
	if err != nil {
		return err
	}
	lite, err := kcclite.ToKCCLiteBestEffort(resource, converter.MetadataLoader, converter.SchemaLoader, serviceMappingLoader, kubeClient)
	if err != nil {
		return fmt.Errorf("error converting KCC full to KCC lite: %w", err)
	}
	dclObj, err := converter.KRMObjectToDCLObject(lite)
	if err != nil {
		return fmt.Errorf("error converting KCC lite to DCL resource: %w", err)
	}
	if err := dclunstruct.Delete(ctx, config, dclObj); err != nil {
		return fmt.Errorf("error deleting the resource %v: %w", resource.GetNamespacedName(), err)
	}
	return nil
}

func newDCLResource(u *unstructured.Unstructured, converter *dclconversion.Converter) (*dcl.Resource, error) {
	s, err := dclschemaloader.GetDCLSchemaForGVK(u.GroupVersionKind(), converter.MetadataLoader, converter.SchemaLoader)
	if err != nil {
		return nil, err
	}
	resource, err := dcl.NewResource(u, s)
	if err != nil {
		return nil, err
	}
	return resource, nil
}

func dclStateToKRM(resource *dcl.Resource, liveState *unstructured.Unstructured, smLoader dclmetadata.ServiceMetadataLoader) (*unstructured.Unstructured, error) {
	spec, status, err := kcclite.ResolveSpecAndStatus(liveState, resource, smLoader)
	if err != nil {
		return nil, err
	}
	resource.Spec = spec
	resource.Status = status
	resource.Labels = liveState.GetLabels()
	return resource.MarshalAsUnstructured()
}

func resourceToKRM(resource *krmtotf.Resource, state *terraform.InstanceState) (*unstructured.Unstructured, error) {
	resource.Spec, resource.Status = krmtotf.ResolveSpecAndStatusWithResourceID(resource, state)
	resource.Labels = krmtotf.GetLabelsFromState(resource, state)
	// Apply post-actuation transformation.
	if err := resourceoverrides.Handler.PostActuationTransform(resource.Original, &resource.Resource, state, nil); err != nil {
		return nil, fmt.Errorf("error applying post-actuation transformation to resource '%v': %w", resource.GetNamespacedName(), err)
	}
	return resource.MarshalAsUnstructured()
}

func getTerraformResourceAndLiveState(ctx context.Context, u *unstructured.Unstructured, provider *tfschema.Provider, c client.Client, smLoader *servicemappingloader.ServiceMappingLoader) (*krmtotf.Resource,
	*terraform.InstanceState, error) {
	resource, err := newTerraformResource(u, provider, smLoader)
	if err != nil {
		return nil, nil, err
	}
	// Apply pre-actuation transformation.
	if err := resourceoverrides.Handler.PreActuationTransform(&resource.Resource); err != nil {
		return nil, nil, fmt.Errorf("error applying pre-actuation transformation to resource '%s': %w", u.GetName(), err)
	}
	liveState, err := krmtotf.FetchLiveState(ctx, resource, provider, c, smLoader)
	if err != nil {
		return nil, nil, fmt.Errorf("error fetching live state: %w", err)
	}
	return resource, liveState, nil
}

func newTerraformResource(u *unstructured.Unstructured, provider *tfschema.Provider, smLoader *servicemappingloader.ServiceMappingLoader) (*krmtotf.Resource, error) {
	sm, err := smLoader.GetServiceMapping(u.GroupVersionKind().Group)
	if err != nil {
		return nil, err
	}
	resource, err := krmtotf.NewResource(u, sm, provider)
	if err != nil {
		return nil, fmt.Errorf("could not parse resource %s: %w", u.GetName(), err)
	}
	return resource, nil
}

func IsNotFoundError(err error) bool {
	if err == nil {
		return false
	}
	return strings.Contains(err.Error(), "is not found")
}

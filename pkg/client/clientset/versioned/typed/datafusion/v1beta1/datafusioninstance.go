// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

// Code generated by main. DO NOT EDIT.

package v1beta1

import (
	"context"
	"time"

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/datafusion/v1beta1"
	scheme "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DataFusionInstancesGetter has a method to return a DataFusionInstanceInterface.
// A group's client should implement this interface.
type DataFusionInstancesGetter interface {
	DataFusionInstances(namespace string) DataFusionInstanceInterface
}

// DataFusionInstanceInterface has methods to work with DataFusionInstance resources.
type DataFusionInstanceInterface interface {
	Create(ctx context.Context, dataFusionInstance *v1beta1.DataFusionInstance, opts v1.CreateOptions) (*v1beta1.DataFusionInstance, error)
	Update(ctx context.Context, dataFusionInstance *v1beta1.DataFusionInstance, opts v1.UpdateOptions) (*v1beta1.DataFusionInstance, error)
	UpdateStatus(ctx context.Context, dataFusionInstance *v1beta1.DataFusionInstance, opts v1.UpdateOptions) (*v1beta1.DataFusionInstance, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.DataFusionInstance, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.DataFusionInstanceList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.DataFusionInstance, err error)
	DataFusionInstanceExpansion
}

// dataFusionInstances implements DataFusionInstanceInterface
type dataFusionInstances struct {
	client rest.Interface
	ns     string
}

// newDataFusionInstances returns a DataFusionInstances
func newDataFusionInstances(c *DatafusionV1beta1Client, namespace string) *dataFusionInstances {
	return &dataFusionInstances{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the dataFusionInstance, and returns the corresponding dataFusionInstance object, and an error if there is any.
func (c *dataFusionInstances) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.DataFusionInstance, err error) {
	result = &v1beta1.DataFusionInstance{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("datafusioninstances").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DataFusionInstances that match those selectors.
func (c *dataFusionInstances) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.DataFusionInstanceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.DataFusionInstanceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("datafusioninstances").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested dataFusionInstances.
func (c *dataFusionInstances) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("datafusioninstances").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a dataFusionInstance and creates it.  Returns the server's representation of the dataFusionInstance, and an error, if there is any.
func (c *dataFusionInstances) Create(ctx context.Context, dataFusionInstance *v1beta1.DataFusionInstance, opts v1.CreateOptions) (result *v1beta1.DataFusionInstance, err error) {
	result = &v1beta1.DataFusionInstance{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("datafusioninstances").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(dataFusionInstance).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a dataFusionInstance and updates it. Returns the server's representation of the dataFusionInstance, and an error, if there is any.
func (c *dataFusionInstances) Update(ctx context.Context, dataFusionInstance *v1beta1.DataFusionInstance, opts v1.UpdateOptions) (result *v1beta1.DataFusionInstance, err error) {
	result = &v1beta1.DataFusionInstance{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("datafusioninstances").
		Name(dataFusionInstance.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(dataFusionInstance).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *dataFusionInstances) UpdateStatus(ctx context.Context, dataFusionInstance *v1beta1.DataFusionInstance, opts v1.UpdateOptions) (result *v1beta1.DataFusionInstance, err error) {
	result = &v1beta1.DataFusionInstance{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("datafusioninstances").
		Name(dataFusionInstance.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(dataFusionInstance).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the dataFusionInstance and deletes it. Returns an error if one occurs.
func (c *dataFusionInstances) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("datafusioninstances").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *dataFusionInstances) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("datafusioninstances").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched dataFusionInstance.
func (c *dataFusionInstances) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.DataFusionInstance, err error) {
	result = &v1beta1.DataFusionInstance{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("datafusioninstances").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

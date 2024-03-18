/*
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	json "encoding/json"
	"fmt"
	"time"

	v1alpha1 "github.com/flomesh-io/fsm/pkg/apis/connector/v1alpha1"
	connectorv1alpha1 "github.com/flomesh-io/fsm/pkg/gen/client/connector/applyconfiguration/connector/v1alpha1"
	scheme "github.com/flomesh-io/fsm/pkg/gen/client/connector/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// GatewayConnectorsGetter has a method to return a GatewayConnectorInterface.
// A group's client should implement this interface.
type GatewayConnectorsGetter interface {
	GatewayConnectors() GatewayConnectorInterface
}

// GatewayConnectorInterface has methods to work with GatewayConnector resources.
type GatewayConnectorInterface interface {
	Create(ctx context.Context, gatewayConnector *v1alpha1.GatewayConnector, opts v1.CreateOptions) (*v1alpha1.GatewayConnector, error)
	Update(ctx context.Context, gatewayConnector *v1alpha1.GatewayConnector, opts v1.UpdateOptions) (*v1alpha1.GatewayConnector, error)
	UpdateStatus(ctx context.Context, gatewayConnector *v1alpha1.GatewayConnector, opts v1.UpdateOptions) (*v1alpha1.GatewayConnector, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.GatewayConnector, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.GatewayConnectorList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.GatewayConnector, err error)
	Apply(ctx context.Context, gatewayConnector *connectorv1alpha1.GatewayConnectorApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.GatewayConnector, err error)
	ApplyStatus(ctx context.Context, gatewayConnector *connectorv1alpha1.GatewayConnectorApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.GatewayConnector, err error)
	GatewayConnectorExpansion
}

// gatewayConnectors implements GatewayConnectorInterface
type gatewayConnectors struct {
	client rest.Interface
}

// newGatewayConnectors returns a GatewayConnectors
func newGatewayConnectors(c *ConnectorV1alpha1Client) *gatewayConnectors {
	return &gatewayConnectors{
		client: c.RESTClient(),
	}
}

// Get takes name of the gatewayConnector, and returns the corresponding gatewayConnector object, and an error if there is any.
func (c *gatewayConnectors) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.GatewayConnector, err error) {
	result = &v1alpha1.GatewayConnector{}
	err = c.client.Get().
		Resource("gatewayconnectors").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GatewayConnectors that match those selectors.
func (c *gatewayConnectors) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.GatewayConnectorList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.GatewayConnectorList{}
	err = c.client.Get().
		Resource("gatewayconnectors").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested gatewayConnectors.
func (c *gatewayConnectors) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("gatewayconnectors").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a gatewayConnector and creates it.  Returns the server's representation of the gatewayConnector, and an error, if there is any.
func (c *gatewayConnectors) Create(ctx context.Context, gatewayConnector *v1alpha1.GatewayConnector, opts v1.CreateOptions) (result *v1alpha1.GatewayConnector, err error) {
	result = &v1alpha1.GatewayConnector{}
	err = c.client.Post().
		Resource("gatewayconnectors").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(gatewayConnector).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a gatewayConnector and updates it. Returns the server's representation of the gatewayConnector, and an error, if there is any.
func (c *gatewayConnectors) Update(ctx context.Context, gatewayConnector *v1alpha1.GatewayConnector, opts v1.UpdateOptions) (result *v1alpha1.GatewayConnector, err error) {
	result = &v1alpha1.GatewayConnector{}
	err = c.client.Put().
		Resource("gatewayconnectors").
		Name(gatewayConnector.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(gatewayConnector).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *gatewayConnectors) UpdateStatus(ctx context.Context, gatewayConnector *v1alpha1.GatewayConnector, opts v1.UpdateOptions) (result *v1alpha1.GatewayConnector, err error) {
	result = &v1alpha1.GatewayConnector{}
	err = c.client.Put().
		Resource("gatewayconnectors").
		Name(gatewayConnector.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(gatewayConnector).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the gatewayConnector and deletes it. Returns an error if one occurs.
func (c *gatewayConnectors) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("gatewayconnectors").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *gatewayConnectors) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("gatewayconnectors").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched gatewayConnector.
func (c *gatewayConnectors) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.GatewayConnector, err error) {
	result = &v1alpha1.GatewayConnector{}
	err = c.client.Patch(pt).
		Resource("gatewayconnectors").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied gatewayConnector.
func (c *gatewayConnectors) Apply(ctx context.Context, gatewayConnector *connectorv1alpha1.GatewayConnectorApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.GatewayConnector, err error) {
	if gatewayConnector == nil {
		return nil, fmt.Errorf("gatewayConnector provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(gatewayConnector)
	if err != nil {
		return nil, err
	}
	name := gatewayConnector.Name
	if name == nil {
		return nil, fmt.Errorf("gatewayConnector.Name must be provided to Apply")
	}
	result = &v1alpha1.GatewayConnector{}
	err = c.client.Patch(types.ApplyPatchType).
		Resource("gatewayconnectors").
		Name(*name).
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *gatewayConnectors) ApplyStatus(ctx context.Context, gatewayConnector *connectorv1alpha1.GatewayConnectorApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.GatewayConnector, err error) {
	if gatewayConnector == nil {
		return nil, fmt.Errorf("gatewayConnector provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(gatewayConnector)
	if err != nil {
		return nil, err
	}

	name := gatewayConnector.Name
	if name == nil {
		return nil, fmt.Errorf("gatewayConnector.Name must be provided to Apply")
	}

	result = &v1alpha1.GatewayConnector{}
	err = c.client.Patch(types.ApplyPatchType).
		Resource("gatewayconnectors").
		Name(*name).
		SubResource("status").
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

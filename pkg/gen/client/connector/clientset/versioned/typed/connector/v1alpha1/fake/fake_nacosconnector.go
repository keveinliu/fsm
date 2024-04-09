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

package fake

import (
	"context"

	v1alpha1 "github.com/flomesh-io/fsm/pkg/apis/connector/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeNacosConnectors implements NacosConnectorInterface
type FakeNacosConnectors struct {
	Fake *FakeConnectorV1alpha1
}

var nacosconnectorsResource = schema.GroupVersionResource{Group: "connector.flomesh.io", Version: "v1alpha1", Resource: "nacosconnectors"}

var nacosconnectorsKind = schema.GroupVersionKind{Group: "connector.flomesh.io", Version: "v1alpha1", Kind: "NacosConnector"}

// Get takes name of the nacosConnector, and returns the corresponding nacosConnector object, and an error if there is any.
func (c *FakeNacosConnectors) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.NacosConnector, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(nacosconnectorsResource, name), &v1alpha1.NacosConnector{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NacosConnector), err
}

// List takes label and field selectors, and returns the list of NacosConnectors that match those selectors.
func (c *FakeNacosConnectors) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.NacosConnectorList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(nacosconnectorsResource, nacosconnectorsKind, opts), &v1alpha1.NacosConnectorList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.NacosConnectorList{ListMeta: obj.(*v1alpha1.NacosConnectorList).ListMeta}
	for _, item := range obj.(*v1alpha1.NacosConnectorList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested nacosConnectors.
func (c *FakeNacosConnectors) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(nacosconnectorsResource, opts))
}

// Create takes the representation of a nacosConnector and creates it.  Returns the server's representation of the nacosConnector, and an error, if there is any.
func (c *FakeNacosConnectors) Create(ctx context.Context, nacosConnector *v1alpha1.NacosConnector, opts v1.CreateOptions) (result *v1alpha1.NacosConnector, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(nacosconnectorsResource, nacosConnector), &v1alpha1.NacosConnector{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NacosConnector), err
}

// Update takes the representation of a nacosConnector and updates it. Returns the server's representation of the nacosConnector, and an error, if there is any.
func (c *FakeNacosConnectors) Update(ctx context.Context, nacosConnector *v1alpha1.NacosConnector, opts v1.UpdateOptions) (result *v1alpha1.NacosConnector, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(nacosconnectorsResource, nacosConnector), &v1alpha1.NacosConnector{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NacosConnector), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNacosConnectors) UpdateStatus(ctx context.Context, nacosConnector *v1alpha1.NacosConnector, opts v1.UpdateOptions) (*v1alpha1.NacosConnector, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(nacosconnectorsResource, "status", nacosConnector), &v1alpha1.NacosConnector{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NacosConnector), err
}

// Delete takes name of the nacosConnector and deletes it. Returns an error if one occurs.
func (c *FakeNacosConnectors) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(nacosconnectorsResource, name, opts), &v1alpha1.NacosConnector{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNacosConnectors) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(nacosconnectorsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.NacosConnectorList{})
	return err
}

// Patch applies the patch and returns the patched nacosConnector.
func (c *FakeNacosConnectors) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.NacosConnector, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(nacosconnectorsResource, name, pt, data, subresources...), &v1alpha1.NacosConnector{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NacosConnector), err
}
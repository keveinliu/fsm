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
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeEurekaConnectors implements EurekaConnectorInterface
type FakeEurekaConnectors struct {
	Fake *FakeConnectorV1alpha1
}

var eurekaconnectorsResource = v1alpha1.SchemeGroupVersion.WithResource("eurekaconnectors")

var eurekaconnectorsKind = v1alpha1.SchemeGroupVersion.WithKind("EurekaConnector")

// Get takes name of the eurekaConnector, and returns the corresponding eurekaConnector object, and an error if there is any.
func (c *FakeEurekaConnectors) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.EurekaConnector, err error) {
	emptyResult := &v1alpha1.EurekaConnector{}
	obj, err := c.Fake.
		Invokes(testing.NewRootGetActionWithOptions(eurekaconnectorsResource, name, options), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.EurekaConnector), err
}

// List takes label and field selectors, and returns the list of EurekaConnectors that match those selectors.
func (c *FakeEurekaConnectors) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.EurekaConnectorList, err error) {
	emptyResult := &v1alpha1.EurekaConnectorList{}
	obj, err := c.Fake.
		Invokes(testing.NewRootListActionWithOptions(eurekaconnectorsResource, eurekaconnectorsKind, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.EurekaConnectorList{ListMeta: obj.(*v1alpha1.EurekaConnectorList).ListMeta}
	for _, item := range obj.(*v1alpha1.EurekaConnectorList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested eurekaConnectors.
func (c *FakeEurekaConnectors) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchActionWithOptions(eurekaconnectorsResource, opts))
}

// Create takes the representation of a eurekaConnector and creates it.  Returns the server's representation of the eurekaConnector, and an error, if there is any.
func (c *FakeEurekaConnectors) Create(ctx context.Context, eurekaConnector *v1alpha1.EurekaConnector, opts v1.CreateOptions) (result *v1alpha1.EurekaConnector, err error) {
	emptyResult := &v1alpha1.EurekaConnector{}
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateActionWithOptions(eurekaconnectorsResource, eurekaConnector, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.EurekaConnector), err
}

// Update takes the representation of a eurekaConnector and updates it. Returns the server's representation of the eurekaConnector, and an error, if there is any.
func (c *FakeEurekaConnectors) Update(ctx context.Context, eurekaConnector *v1alpha1.EurekaConnector, opts v1.UpdateOptions) (result *v1alpha1.EurekaConnector, err error) {
	emptyResult := &v1alpha1.EurekaConnector{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateActionWithOptions(eurekaconnectorsResource, eurekaConnector, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.EurekaConnector), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeEurekaConnectors) UpdateStatus(ctx context.Context, eurekaConnector *v1alpha1.EurekaConnector, opts v1.UpdateOptions) (result *v1alpha1.EurekaConnector, err error) {
	emptyResult := &v1alpha1.EurekaConnector{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceActionWithOptions(eurekaconnectorsResource, "status", eurekaConnector, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.EurekaConnector), err
}

// Delete takes name of the eurekaConnector and deletes it. Returns an error if one occurs.
func (c *FakeEurekaConnectors) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(eurekaconnectorsResource, name, opts), &v1alpha1.EurekaConnector{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEurekaConnectors) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionActionWithOptions(eurekaconnectorsResource, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.EurekaConnectorList{})
	return err
}

// Patch applies the patch and returns the patched eurekaConnector.
func (c *FakeEurekaConnectors) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.EurekaConnector, err error) {
	emptyResult := &v1alpha1.EurekaConnector{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(eurekaconnectorsResource, name, pt, data, opts, subresources...), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.EurekaConnector), err
}

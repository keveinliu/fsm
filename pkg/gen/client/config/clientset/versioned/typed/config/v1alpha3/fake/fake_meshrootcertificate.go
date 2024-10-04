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

	v1alpha3 "github.com/flomesh-io/fsm/pkg/apis/config/v1alpha3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeMeshRootCertificates implements MeshRootCertificateInterface
type FakeMeshRootCertificates struct {
	Fake *FakeConfigV1alpha3
	ns   string
}

var meshrootcertificatesResource = v1alpha3.SchemeGroupVersion.WithResource("meshrootcertificates")

var meshrootcertificatesKind = v1alpha3.SchemeGroupVersion.WithKind("MeshRootCertificate")

// Get takes name of the meshRootCertificate, and returns the corresponding meshRootCertificate object, and an error if there is any.
func (c *FakeMeshRootCertificates) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha3.MeshRootCertificate, err error) {
	emptyResult := &v1alpha3.MeshRootCertificate{}
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithOptions(meshrootcertificatesResource, c.ns, name, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha3.MeshRootCertificate), err
}

// List takes label and field selectors, and returns the list of MeshRootCertificates that match those selectors.
func (c *FakeMeshRootCertificates) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha3.MeshRootCertificateList, err error) {
	emptyResult := &v1alpha3.MeshRootCertificateList{}
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithOptions(meshrootcertificatesResource, meshrootcertificatesKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha3.MeshRootCertificateList{ListMeta: obj.(*v1alpha3.MeshRootCertificateList).ListMeta}
	for _, item := range obj.(*v1alpha3.MeshRootCertificateList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested meshRootCertificates.
func (c *FakeMeshRootCertificates) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithOptions(meshrootcertificatesResource, c.ns, opts))

}

// Create takes the representation of a meshRootCertificate and creates it.  Returns the server's representation of the meshRootCertificate, and an error, if there is any.
func (c *FakeMeshRootCertificates) Create(ctx context.Context, meshRootCertificate *v1alpha3.MeshRootCertificate, opts v1.CreateOptions) (result *v1alpha3.MeshRootCertificate, err error) {
	emptyResult := &v1alpha3.MeshRootCertificate{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithOptions(meshrootcertificatesResource, c.ns, meshRootCertificate, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha3.MeshRootCertificate), err
}

// Update takes the representation of a meshRootCertificate and updates it. Returns the server's representation of the meshRootCertificate, and an error, if there is any.
func (c *FakeMeshRootCertificates) Update(ctx context.Context, meshRootCertificate *v1alpha3.MeshRootCertificate, opts v1.UpdateOptions) (result *v1alpha3.MeshRootCertificate, err error) {
	emptyResult := &v1alpha3.MeshRootCertificate{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithOptions(meshrootcertificatesResource, c.ns, meshRootCertificate, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha3.MeshRootCertificate), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMeshRootCertificates) UpdateStatus(ctx context.Context, meshRootCertificate *v1alpha3.MeshRootCertificate, opts v1.UpdateOptions) (result *v1alpha3.MeshRootCertificate, err error) {
	emptyResult := &v1alpha3.MeshRootCertificate{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceActionWithOptions(meshrootcertificatesResource, "status", c.ns, meshRootCertificate, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha3.MeshRootCertificate), err
}

// Delete takes name of the meshRootCertificate and deletes it. Returns an error if one occurs.
func (c *FakeMeshRootCertificates) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(meshrootcertificatesResource, c.ns, name, opts), &v1alpha3.MeshRootCertificate{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMeshRootCertificates) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithOptions(meshrootcertificatesResource, c.ns, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha3.MeshRootCertificateList{})
	return err
}

// Patch applies the patch and returns the patched meshRootCertificate.
func (c *FakeMeshRootCertificates) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha3.MeshRootCertificate, err error) {
	emptyResult := &v1alpha3.MeshRootCertificate{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(meshrootcertificatesResource, c.ns, name, pt, data, opts, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha3.MeshRootCertificate), err
}

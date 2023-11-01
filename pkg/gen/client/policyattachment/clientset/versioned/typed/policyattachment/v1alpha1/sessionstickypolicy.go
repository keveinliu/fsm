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
	"time"

	v1alpha1 "github.com/flomesh-io/fsm/pkg/apis/policyattachment/v1alpha1"
	scheme "github.com/flomesh-io/fsm/pkg/gen/client/policyattachment/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// SessionStickyPoliciesGetter has a method to return a SessionStickyPolicyInterface.
// A group's client should implement this interface.
type SessionStickyPoliciesGetter interface {
	SessionStickyPolicies(namespace string) SessionStickyPolicyInterface
}

// SessionStickyPolicyInterface has methods to work with SessionStickyPolicy resources.
type SessionStickyPolicyInterface interface {
	Create(ctx context.Context, sessionStickyPolicy *v1alpha1.SessionStickyPolicy, opts v1.CreateOptions) (*v1alpha1.SessionStickyPolicy, error)
	Update(ctx context.Context, sessionStickyPolicy *v1alpha1.SessionStickyPolicy, opts v1.UpdateOptions) (*v1alpha1.SessionStickyPolicy, error)
	UpdateStatus(ctx context.Context, sessionStickyPolicy *v1alpha1.SessionStickyPolicy, opts v1.UpdateOptions) (*v1alpha1.SessionStickyPolicy, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.SessionStickyPolicy, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.SessionStickyPolicyList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.SessionStickyPolicy, err error)
	SessionStickyPolicyExpansion
}

// sessionStickyPolicies implements SessionStickyPolicyInterface
type sessionStickyPolicies struct {
	client rest.Interface
	ns     string
}

// newSessionStickyPolicies returns a SessionStickyPolicies
func newSessionStickyPolicies(c *GatewayV1alpha1Client, namespace string) *sessionStickyPolicies {
	return &sessionStickyPolicies{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the sessionStickyPolicy, and returns the corresponding sessionStickyPolicy object, and an error if there is any.
func (c *sessionStickyPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.SessionStickyPolicy, err error) {
	result = &v1alpha1.SessionStickyPolicy{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sessionstickypolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SessionStickyPolicies that match those selectors.
func (c *sessionStickyPolicies) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SessionStickyPolicyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.SessionStickyPolicyList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sessionstickypolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested sessionStickyPolicies.
func (c *sessionStickyPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("sessionstickypolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a sessionStickyPolicy and creates it.  Returns the server's representation of the sessionStickyPolicy, and an error, if there is any.
func (c *sessionStickyPolicies) Create(ctx context.Context, sessionStickyPolicy *v1alpha1.SessionStickyPolicy, opts v1.CreateOptions) (result *v1alpha1.SessionStickyPolicy, err error) {
	result = &v1alpha1.SessionStickyPolicy{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("sessionstickypolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(sessionStickyPolicy).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a sessionStickyPolicy and updates it. Returns the server's representation of the sessionStickyPolicy, and an error, if there is any.
func (c *sessionStickyPolicies) Update(ctx context.Context, sessionStickyPolicy *v1alpha1.SessionStickyPolicy, opts v1.UpdateOptions) (result *v1alpha1.SessionStickyPolicy, err error) {
	result = &v1alpha1.SessionStickyPolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("sessionstickypolicies").
		Name(sessionStickyPolicy.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(sessionStickyPolicy).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *sessionStickyPolicies) UpdateStatus(ctx context.Context, sessionStickyPolicy *v1alpha1.SessionStickyPolicy, opts v1.UpdateOptions) (result *v1alpha1.SessionStickyPolicy, err error) {
	result = &v1alpha1.SessionStickyPolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("sessionstickypolicies").
		Name(sessionStickyPolicy.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(sessionStickyPolicy).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the sessionStickyPolicy and deletes it. Returns an error if one occurs.
func (c *sessionStickyPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sessionstickypolicies").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *sessionStickyPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sessionstickypolicies").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched sessionStickyPolicy.
func (c *sessionStickyPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.SessionStickyPolicy, err error) {
	result = &v1alpha1.SessionStickyPolicy{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("sessionstickypolicies").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
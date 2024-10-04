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

	v1alpha1 "github.com/flomesh-io/fsm/pkg/apis/policy/v1alpha1"
	scheme "github.com/flomesh-io/fsm/pkg/gen/client/policy/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// AccessCertsGetter has a method to return a AccessCertInterface.
// A group's client should implement this interface.
type AccessCertsGetter interface {
	AccessCerts(namespace string) AccessCertInterface
}

// AccessCertInterface has methods to work with AccessCert resources.
type AccessCertInterface interface {
	Create(ctx context.Context, accessCert *v1alpha1.AccessCert, opts v1.CreateOptions) (*v1alpha1.AccessCert, error)
	Update(ctx context.Context, accessCert *v1alpha1.AccessCert, opts v1.UpdateOptions) (*v1alpha1.AccessCert, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, accessCert *v1alpha1.AccessCert, opts v1.UpdateOptions) (*v1alpha1.AccessCert, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.AccessCert, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.AccessCertList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AccessCert, err error)
	AccessCertExpansion
}

// accessCerts implements AccessCertInterface
type accessCerts struct {
	*gentype.ClientWithList[*v1alpha1.AccessCert, *v1alpha1.AccessCertList]
}

// newAccessCerts returns a AccessCerts
func newAccessCerts(c *PolicyV1alpha1Client, namespace string) *accessCerts {
	return &accessCerts{
		gentype.NewClientWithList[*v1alpha1.AccessCert, *v1alpha1.AccessCertList](
			"accesscerts",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1alpha1.AccessCert { return &v1alpha1.AccessCert{} },
			func() *v1alpha1.AccessCertList { return &v1alpha1.AccessCertList{} }),
	}
}

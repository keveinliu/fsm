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
// Code generated by informer-gen. DO NOT EDIT.

package externalversions

import (
	"fmt"

	v1alpha1 "github.com/flomesh-io/fsm/pkg/apis/extension/v1alpha1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
)

// GenericInformer is type of SharedIndexInformer which will locate and delegate to other
// sharedInformers based on type
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=extension.gateway.flomesh.io, Version=v1alpha1
	case v1alpha1.SchemeGroupVersion.WithResource("circuitbreakers"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Extension().V1alpha1().CircuitBreakers().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("concurrencylimits"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Extension().V1alpha1().ConcurrencyLimits().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("dnsmodifiers"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Extension().V1alpha1().DNSModifiers().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("externalratelimits"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Extension().V1alpha1().ExternalRateLimits().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("faultinjections"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Extension().V1alpha1().FaultInjections().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("filters"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Extension().V1alpha1().Filters().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("filterconfigs"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Extension().V1alpha1().FilterConfigs().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("filterdefinitions"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Extension().V1alpha1().FilterDefinitions().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("httplogs"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Extension().V1alpha1().HTTPLogs().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("iprestrictions"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Extension().V1alpha1().IPRestrictions().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("listenerfilters"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Extension().V1alpha1().ListenerFilters().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("metricses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Extension().V1alpha1().Metricses().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("proxytags"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Extension().V1alpha1().ProxyTags().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("ratelimits"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Extension().V1alpha1().RateLimits().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("requestterminations"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Extension().V1alpha1().RequestTerminations().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("zipkins"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Extension().V1alpha1().Zipkins().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}

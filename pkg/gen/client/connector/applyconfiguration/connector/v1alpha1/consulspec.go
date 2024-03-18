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
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

// ConsulSpecApplyConfiguration represents an declarative configuration of the ConsulSpec type for use
// with apply.
type ConsulSpecApplyConfiguration struct {
	HTTPAddr           *string                                  `json:"httpAddr,omitempty"`
	DeriveNamespace    *string                                  `json:"deriveNamespace,omitempty"`
	AsInternalServices *bool                                    `json:"asInternalServices,omitempty"`
	Auth               *NacosAuthSpecApplyConfiguration         `json:"auth,omitempty"`
	SyncToK8S          *ConsulSyncToK8SSpecApplyConfiguration   `json:"syncToK8S,omitempty"`
	SyncFromK8S        *ConsulSyncFromK8SSpecApplyConfiguration `json:"syncFromK8S,omitempty"`
	Limiter            *LimiterApplyConfiguration               `json:"limiter,omitempty"`
}

// ConsulSpecApplyConfiguration constructs an declarative configuration of the ConsulSpec type for use with
// apply.
func ConsulSpec() *ConsulSpecApplyConfiguration {
	return &ConsulSpecApplyConfiguration{}
}

// WithHTTPAddr sets the HTTPAddr field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the HTTPAddr field is set to the value of the last call.
func (b *ConsulSpecApplyConfiguration) WithHTTPAddr(value string) *ConsulSpecApplyConfiguration {
	b.HTTPAddr = &value
	return b
}

// WithDeriveNamespace sets the DeriveNamespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeriveNamespace field is set to the value of the last call.
func (b *ConsulSpecApplyConfiguration) WithDeriveNamespace(value string) *ConsulSpecApplyConfiguration {
	b.DeriveNamespace = &value
	return b
}

// WithAsInternalServices sets the AsInternalServices field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AsInternalServices field is set to the value of the last call.
func (b *ConsulSpecApplyConfiguration) WithAsInternalServices(value bool) *ConsulSpecApplyConfiguration {
	b.AsInternalServices = &value
	return b
}

// WithAuth sets the Auth field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Auth field is set to the value of the last call.
func (b *ConsulSpecApplyConfiguration) WithAuth(value *NacosAuthSpecApplyConfiguration) *ConsulSpecApplyConfiguration {
	b.Auth = value
	return b
}

// WithSyncToK8S sets the SyncToK8S field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SyncToK8S field is set to the value of the last call.
func (b *ConsulSpecApplyConfiguration) WithSyncToK8S(value *ConsulSyncToK8SSpecApplyConfiguration) *ConsulSpecApplyConfiguration {
	b.SyncToK8S = value
	return b
}

// WithSyncFromK8S sets the SyncFromK8S field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SyncFromK8S field is set to the value of the last call.
func (b *ConsulSpecApplyConfiguration) WithSyncFromK8S(value *ConsulSyncFromK8SSpecApplyConfiguration) *ConsulSpecApplyConfiguration {
	b.SyncFromK8S = value
	return b
}

// WithLimiter sets the Limiter field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Limiter field is set to the value of the last call.
func (b *ConsulSpecApplyConfiguration) WithLimiter(value *LimiterApplyConfiguration) *ConsulSpecApplyConfiguration {
	b.Limiter = value
	return b
}

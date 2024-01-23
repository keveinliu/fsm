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

package v1alpha2

// LocalDNSProxyApplyConfiguration represents an declarative configuration of the LocalDNSProxy type for use
// with apply.
type LocalDNSProxyApplyConfiguration struct {
	Enable                           *bool   `json:"enable,omitempty"`
	PrimaryUpstreamDNSServerIPAddr   *string `json:"primaryUpstreamDNSServerIPAddr,omitempty"`
	SecondaryUpstreamDNSServerIPAddr *string `json:"secondaryUpstreamDNSServerIPAddr,omitempty"`
}

// LocalDNSProxyApplyConfiguration constructs an declarative configuration of the LocalDNSProxy type for use with
// apply.
func LocalDNSProxy() *LocalDNSProxyApplyConfiguration {
	return &LocalDNSProxyApplyConfiguration{}
}

// WithEnable sets the Enable field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Enable field is set to the value of the last call.
func (b *LocalDNSProxyApplyConfiguration) WithEnable(value bool) *LocalDNSProxyApplyConfiguration {
	b.Enable = &value
	return b
}

// WithPrimaryUpstreamDNSServerIPAddr sets the PrimaryUpstreamDNSServerIPAddr field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PrimaryUpstreamDNSServerIPAddr field is set to the value of the last call.
func (b *LocalDNSProxyApplyConfiguration) WithPrimaryUpstreamDNSServerIPAddr(value string) *LocalDNSProxyApplyConfiguration {
	b.PrimaryUpstreamDNSServerIPAddr = &value
	return b
}

// WithSecondaryUpstreamDNSServerIPAddr sets the SecondaryUpstreamDNSServerIPAddr field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SecondaryUpstreamDNSServerIPAddr field is set to the value of the last call.
func (b *LocalDNSProxyApplyConfiguration) WithSecondaryUpstreamDNSServerIPAddr(value string) *LocalDNSProxyApplyConfiguration {
	b.SecondaryUpstreamDNSServerIPAddr = &value
	return b
}

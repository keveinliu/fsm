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

package v1alpha3

// ServiceLBSpecApplyConfiguration represents an declarative configuration of the ServiceLBSpec type for use
// with apply.
type ServiceLBSpecApplyConfiguration struct {
	Enabled *bool   `json:"enabled,omitempty"`
	Image   *string `json:"image,omitempty"`
}

// ServiceLBSpecApplyConfiguration constructs an declarative configuration of the ServiceLBSpec type for use with
// apply.
func ServiceLBSpec() *ServiceLBSpecApplyConfiguration {
	return &ServiceLBSpecApplyConfiguration{}
}

// WithEnabled sets the Enabled field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Enabled field is set to the value of the last call.
func (b *ServiceLBSpecApplyConfiguration) WithEnabled(value bool) *ServiceLBSpecApplyConfiguration {
	b.Enabled = &value
	return b
}

// WithImage sets the Image field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Image field is set to the value of the last call.
func (b *ServiceLBSpecApplyConfiguration) WithImage(value string) *ServiceLBSpecApplyConfiguration {
	b.Image = &value
	return b
}

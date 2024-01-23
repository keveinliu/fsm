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

// GatewayBindingSubjectApplyConfiguration represents an declarative configuration of the GatewayBindingSubject type for use
// with apply.
type GatewayBindingSubjectApplyConfiguration struct {
	Service   *string `json:"service,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
	Weight    *int    `json:"weight,omitempty"`
}

// GatewayBindingSubjectApplyConfiguration constructs an declarative configuration of the GatewayBindingSubject type for use with
// apply.
func GatewayBindingSubject() *GatewayBindingSubjectApplyConfiguration {
	return &GatewayBindingSubjectApplyConfiguration{}
}

// WithService sets the Service field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Service field is set to the value of the last call.
func (b *GatewayBindingSubjectApplyConfiguration) WithService(value string) *GatewayBindingSubjectApplyConfiguration {
	b.Service = &value
	return b
}

// WithNamespace sets the Namespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Namespace field is set to the value of the last call.
func (b *GatewayBindingSubjectApplyConfiguration) WithNamespace(value string) *GatewayBindingSubjectApplyConfiguration {
	b.Namespace = &value
	return b
}

// WithWeight sets the Weight field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Weight field is set to the value of the last call.
func (b *GatewayBindingSubjectApplyConfiguration) WithWeight(value int) *GatewayBindingSubjectApplyConfiguration {
	b.Weight = &value
	return b
}

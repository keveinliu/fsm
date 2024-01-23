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

import (
	v1 "k8s.io/api/core/v1"
)

// TresorCASpecApplyConfiguration represents an declarative configuration of the TresorCASpec type for use
// with apply.
type TresorCASpecApplyConfiguration struct {
	SecretRef *v1.SecretReference `json:"secretRef,omitempty"`
}

// TresorCASpecApplyConfiguration constructs an declarative configuration of the TresorCASpec type for use with
// apply.
func TresorCASpec() *TresorCASpecApplyConfiguration {
	return &TresorCASpecApplyConfiguration{}
}

// WithSecretRef sets the SecretRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SecretRef field is set to the value of the last call.
func (b *TresorCASpecApplyConfiguration) WithSecretRef(value v1.SecretReference) *TresorCASpecApplyConfiguration {
	b.SecretRef = &value
	return b
}

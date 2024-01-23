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

// ExternalAuthzSpecApplyConfiguration represents an declarative configuration of the ExternalAuthzSpec type for use
// with apply.
type ExternalAuthzSpecApplyConfiguration struct {
	Enable           *bool   `json:"enable,omitempty"`
	Address          *string `json:"address,omitempty"`
	Port             *uint16 `json:"port,omitempty"`
	StatPrefix       *string `json:"statPrefix,omitempty"`
	Timeout          *string `json:"timeout,omitempty"`
	FailureModeAllow *bool   `json:"failureModeAllow,omitempty"`
}

// ExternalAuthzSpecApplyConfiguration constructs an declarative configuration of the ExternalAuthzSpec type for use with
// apply.
func ExternalAuthzSpec() *ExternalAuthzSpecApplyConfiguration {
	return &ExternalAuthzSpecApplyConfiguration{}
}

// WithEnable sets the Enable field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Enable field is set to the value of the last call.
func (b *ExternalAuthzSpecApplyConfiguration) WithEnable(value bool) *ExternalAuthzSpecApplyConfiguration {
	b.Enable = &value
	return b
}

// WithAddress sets the Address field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Address field is set to the value of the last call.
func (b *ExternalAuthzSpecApplyConfiguration) WithAddress(value string) *ExternalAuthzSpecApplyConfiguration {
	b.Address = &value
	return b
}

// WithPort sets the Port field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Port field is set to the value of the last call.
func (b *ExternalAuthzSpecApplyConfiguration) WithPort(value uint16) *ExternalAuthzSpecApplyConfiguration {
	b.Port = &value
	return b
}

// WithStatPrefix sets the StatPrefix field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the StatPrefix field is set to the value of the last call.
func (b *ExternalAuthzSpecApplyConfiguration) WithStatPrefix(value string) *ExternalAuthzSpecApplyConfiguration {
	b.StatPrefix = &value
	return b
}

// WithTimeout sets the Timeout field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Timeout field is set to the value of the last call.
func (b *ExternalAuthzSpecApplyConfiguration) WithTimeout(value string) *ExternalAuthzSpecApplyConfiguration {
	b.Timeout = &value
	return b
}

// WithFailureModeAllow sets the FailureModeAllow field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the FailureModeAllow field is set to the value of the last call.
func (b *ExternalAuthzSpecApplyConfiguration) WithFailureModeAllow(value bool) *ExternalAuthzSpecApplyConfiguration {
	b.FailureModeAllow = &value
	return b
}

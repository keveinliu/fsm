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

// RemoteLoggingSpecApplyConfiguration represents an declarative configuration of the RemoteLoggingSpec type for use
// with apply.
type RemoteLoggingSpecApplyConfiguration struct {
	Enable          *bool    `json:"enable,omitempty"`
	Level           *uint16  `json:"level,omitempty"`
	Port            *uint16  `json:"port,omitempty"`
	Address         *string  `json:"address,omitempty"`
	Endpoint        *string  `json:"endpoint,omitempty"`
	Authorization   *string  `json:"authorization,omitempty"`
	SampledFraction *float32 `json:"sampledFraction,omitempty"`
}

// RemoteLoggingSpecApplyConfiguration constructs an declarative configuration of the RemoteLoggingSpec type for use with
// apply.
func RemoteLoggingSpec() *RemoteLoggingSpecApplyConfiguration {
	return &RemoteLoggingSpecApplyConfiguration{}
}

// WithEnable sets the Enable field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Enable field is set to the value of the last call.
func (b *RemoteLoggingSpecApplyConfiguration) WithEnable(value bool) *RemoteLoggingSpecApplyConfiguration {
	b.Enable = &value
	return b
}

// WithLevel sets the Level field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Level field is set to the value of the last call.
func (b *RemoteLoggingSpecApplyConfiguration) WithLevel(value uint16) *RemoteLoggingSpecApplyConfiguration {
	b.Level = &value
	return b
}

// WithPort sets the Port field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Port field is set to the value of the last call.
func (b *RemoteLoggingSpecApplyConfiguration) WithPort(value uint16) *RemoteLoggingSpecApplyConfiguration {
	b.Port = &value
	return b
}

// WithAddress sets the Address field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Address field is set to the value of the last call.
func (b *RemoteLoggingSpecApplyConfiguration) WithAddress(value string) *RemoteLoggingSpecApplyConfiguration {
	b.Address = &value
	return b
}

// WithEndpoint sets the Endpoint field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Endpoint field is set to the value of the last call.
func (b *RemoteLoggingSpecApplyConfiguration) WithEndpoint(value string) *RemoteLoggingSpecApplyConfiguration {
	b.Endpoint = &value
	return b
}

// WithAuthorization sets the Authorization field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Authorization field is set to the value of the last call.
func (b *RemoteLoggingSpecApplyConfiguration) WithAuthorization(value string) *RemoteLoggingSpecApplyConfiguration {
	b.Authorization = &value
	return b
}

// WithSampledFraction sets the SampledFraction field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SampledFraction field is set to the value of the last call.
func (b *RemoteLoggingSpecApplyConfiguration) WithSampledFraction(value float32) *RemoteLoggingSpecApplyConfiguration {
	b.SampledFraction = &value
	return b
}

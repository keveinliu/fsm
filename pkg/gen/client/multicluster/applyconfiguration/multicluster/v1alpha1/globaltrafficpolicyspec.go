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

import (
	v1alpha1 "github.com/flomesh-io/fsm/pkg/apis/multicluster/v1alpha1"
)

// GlobalTrafficPolicySpecApplyConfiguration represents an declarative configuration of the GlobalTrafficPolicySpec type for use
// with apply.
type GlobalTrafficPolicySpecApplyConfiguration struct {
	LbType  *v1alpha1.LoadBalancerType        `json:"lbType,omitempty"`
	Targets []TrafficTargetApplyConfiguration `json:"targets,omitempty"`
}

// GlobalTrafficPolicySpecApplyConfiguration constructs an declarative configuration of the GlobalTrafficPolicySpec type for use with
// apply.
func GlobalTrafficPolicySpec() *GlobalTrafficPolicySpecApplyConfiguration {
	return &GlobalTrafficPolicySpecApplyConfiguration{}
}

// WithLbType sets the LbType field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LbType field is set to the value of the last call.
func (b *GlobalTrafficPolicySpecApplyConfiguration) WithLbType(value v1alpha1.LoadBalancerType) *GlobalTrafficPolicySpecApplyConfiguration {
	b.LbType = &value
	return b
}

// WithTargets adds the given value to the Targets field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Targets field.
func (b *GlobalTrafficPolicySpecApplyConfiguration) WithTargets(values ...*TrafficTargetApplyConfiguration) *GlobalTrafficPolicySpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithTargets")
		}
		b.Targets = append(b.Targets, *values[i])
	}
	return b
}

/*
Copyright 2024 Red Hat Inc.

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

package v1beta1

import (
	"github.com/konflux-ci/integration-service/api/v1beta2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

func (r *IntegrationTestScenario) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// Hub marks this type as a conversion hub.
// ConvertTo converts this ITS to the Hub version (v1beta2).
func (src *IntegrationTestScenario) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1beta2.IntegrationTestScenario)
	dst.ObjectMeta = src.ObjectMeta
	dst.Spec.Application = src.Spec.Application
	dst.Status = v1beta2.IntegrationTestScenarioStatus{Conditions: make([]metav1.Condition, 0)}

	if src.Spec.Params != nil {
		for _, par := range src.Spec.Params {
			dst.Spec.Params = append(dst.Spec.Params, v1beta2.PipelineParameter(par))
		}
	}
	if src.Spec.Contexts != nil {
		for _, par := range src.Spec.Contexts {
			dst.Spec.Contexts = append(dst.Spec.Contexts, v1beta2.TestContext(par))
		}
	}

	if src.Status.Conditions != nil {
		dst.Status.Conditions = append(dst.Status.Conditions, src.Status.Conditions...)
	}

	dst.Spec.ResolverRef = v1beta2.ResolverRef{
		Resolver: src.Spec.ResolverRef.Resolver,
	}

	if src.Spec.ResolverRef.Params != nil {
		for _, par := range src.Spec.ResolverRef.Params {
			dst.Spec.ResolverRef.Params = append(dst.Spec.ResolverRef.Params, v1beta2.ResolverParameter{
				Name:  par.Name,
				Value: par.Value,
			})
		}
	}

	return nil
}

func (dst *IntegrationTestScenario) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1beta2.IntegrationTestScenario)
	dst.ObjectMeta = src.ObjectMeta
	dst.Spec.Application = src.Spec.Application

	if src.Spec.Params != nil {
		for _, par := range src.Spec.Params {
			dst.Spec.Params = append(dst.Spec.Params, PipelineParameter(par))
		}
	}
	if src.Spec.Contexts != nil {
		for _, par := range src.Spec.Contexts {
			dst.Spec.Contexts = append(dst.Spec.Contexts, TestContext(par))
		}
	}

	if src.Status.Conditions != nil {
		dst.Status.Conditions = append(dst.Status.Conditions, src.Status.Conditions...)
	}
	dst.Spec.ResolverRef = ResolverRef{
		Resolver: src.Spec.ResolverRef.Resolver,
	}
	if src.Spec.ResolverRef.Params != nil {
		for _, par := range src.Spec.ResolverRef.Params {
			dst.Spec.ResolverRef.Params = append(dst.Spec.ResolverRef.Params, ResolverParameter{
				Name:  par.Name,
				Value: par.Value,
			})
		}
	}

	return nil
}

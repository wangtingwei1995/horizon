// Copyright © 2023 Horizoncd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pod

import (
	"github.com/argoproj/argo-cd/pkg/apis/application/v1alpha1"
	herrors "github.com/horizoncd/horizon/core/errors"
	perror "github.com/horizoncd/horizon/pkg/errors"
	"github.com/horizoncd/horizon/pkg/workload"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic/dynamicinformer"
)

var (
	GVRPod = schema.GroupVersionResource{
		Group:    "",
		Version:  "v1",
		Resource: "pods",
	}
)

func init() {
	workload.Register(ability, GVRPod)
}

var ability = &pod{}

type pod struct{}

func (*pod) MatchGK(gk schema.GroupKind) bool {
	return gk.Kind == "Pod"
}

func (*pod) getPod(node *v1alpha1.ResourceNode,
	factory dynamicinformer.DynamicSharedInformerFactory) (*corev1.Pod, error) {
	instance, err := factory.ForResource(GVRPod).Lister().ByNamespace(node.Namespace).Get(node.Name)
	if err != nil {
		return nil, perror.Wrapf(
			herrors.NewErrGetFailed(herrors.ResourceInK8S,
				"failed to get deployment in k8s"),
			"failed to get deployment in k8s: deployment = %s, err = %v", node.Name, err)
	}
	return instance.(*corev1.Pod), nil
}

func (p *pod) ListPods(node *v1alpha1.ResourceNode,
	factory dynamicinformer.DynamicSharedInformerFactory) ([]corev1.Pod, error) {
	instance, err := p.getPod(node, factory)
	if err != nil {
		return nil, err
	}

	return []corev1.Pod{*instance}, nil
}

func (*pod) Action(actionName string, un *unstructured.Unstructured) (*unstructured.Unstructured, error) {
	return un, nil
}

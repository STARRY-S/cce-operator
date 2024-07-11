/*
Copyright 2023 [Rancher Labs, Inc](https://rancher.com).

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

// Code generated by main. DO NOT EDIT.

package v1

import (
	"github.com/rancher/lasso/pkg/controller"
	"github.com/rancher/wrangler/v3/pkg/generic"
	"github.com/rancher/wrangler/v3/pkg/schemes"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func init() {
	schemes.Register(v1.AddToScheme)
}

type Interface interface {
	Node() NodeController
	Pod() PodController
	Secret() SecretController
}

func New(controllerFactory controller.SharedControllerFactory) Interface {
	return &version{
		controllerFactory: controllerFactory,
	}
}

type version struct {
	controllerFactory controller.SharedControllerFactory
}

func (v *version) Node() NodeController {
	return generic.NewNonNamespacedController[*v1.Node, *v1.NodeList](schema.GroupVersionKind{Group: "", Version: "v1", Kind: "Node"}, "nodes", v.controllerFactory)
}

func (v *version) Pod() PodController {
	return generic.NewController[*v1.Pod, *v1.PodList](schema.GroupVersionKind{Group: "", Version: "v1", Kind: "Pod"}, "pods", true, v.controllerFactory)
}

func (v *version) Secret() SecretController {
	return generic.NewController[*v1.Secret, *v1.SecretList](schema.GroupVersionKind{Group: "", Version: "v1", Kind: "Secret"}, "secrets", true, v.controllerFactory)
}

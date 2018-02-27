/*
Copyright 2018 The Kubernetes Authors.

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

package fake

import (
	v1alpha1 "github.com/rancher/longhorn-manager/k8s/pkg/apis/longhorn/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeEngines implements EngineInterface
type FakeEngines struct {
	Fake *FakeLonghornV1alpha1
	ns   string
}

var enginesResource = schema.GroupVersionResource{Group: "longhorn.rancher.io", Version: "v1alpha1", Resource: "engines"}

var enginesKind = schema.GroupVersionKind{Group: "longhorn.rancher.io", Version: "v1alpha1", Kind: "Engine"}

// Get takes name of the engine, and returns the corresponding engine object, and an error if there is any.
func (c *FakeEngines) Get(name string, options v1.GetOptions) (result *v1alpha1.Engine, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(enginesResource, c.ns, name), &v1alpha1.Engine{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Engine), err
}

// List takes label and field selectors, and returns the list of Engines that match those selectors.
func (c *FakeEngines) List(opts v1.ListOptions) (result *v1alpha1.EngineList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(enginesResource, enginesKind, c.ns, opts), &v1alpha1.EngineList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.EngineList{}
	for _, item := range obj.(*v1alpha1.EngineList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested engines.
func (c *FakeEngines) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(enginesResource, c.ns, opts))

}

// Create takes the representation of a engine and creates it.  Returns the server's representation of the engine, and an error, if there is any.
func (c *FakeEngines) Create(engine *v1alpha1.Engine) (result *v1alpha1.Engine, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(enginesResource, c.ns, engine), &v1alpha1.Engine{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Engine), err
}

// Update takes the representation of a engine and updates it. Returns the server's representation of the engine, and an error, if there is any.
func (c *FakeEngines) Update(engine *v1alpha1.Engine) (result *v1alpha1.Engine, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(enginesResource, c.ns, engine), &v1alpha1.Engine{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Engine), err
}

// Delete takes name of the engine and deletes it. Returns an error if one occurs.
func (c *FakeEngines) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(enginesResource, c.ns, name), &v1alpha1.Engine{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEngines) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(enginesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.EngineList{})
	return err
}

// Patch applies the patch and returns the patched engine.
func (c *FakeEngines) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Engine, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(enginesResource, c.ns, name, data, subresources...), &v1alpha1.Engine{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Engine), err
}

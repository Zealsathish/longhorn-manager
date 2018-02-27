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

package v1alpha1

import (
	v1alpha1 "github.com/rancher/longhorn-manager/k8s/pkg/apis/longhorn/v1alpha1"
	scheme "github.com/rancher/longhorn-manager/k8s/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// EnginesGetter has a method to return a EngineInterface.
// A group's client should implement this interface.
type EnginesGetter interface {
	Engines(namespace string) EngineInterface
}

// EngineInterface has methods to work with Engine resources.
type EngineInterface interface {
	Create(*v1alpha1.Engine) (*v1alpha1.Engine, error)
	Update(*v1alpha1.Engine) (*v1alpha1.Engine, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Engine, error)
	List(opts v1.ListOptions) (*v1alpha1.EngineList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Engine, err error)
	EngineExpansion
}

// engines implements EngineInterface
type engines struct {
	client rest.Interface
	ns     string
}

// newEngines returns a Engines
func newEngines(c *LonghornV1alpha1Client, namespace string) *engines {
	return &engines{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the engine, and returns the corresponding engine object, and an error if there is any.
func (c *engines) Get(name string, options v1.GetOptions) (result *v1alpha1.Engine, err error) {
	result = &v1alpha1.Engine{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("engines").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Engines that match those selectors.
func (c *engines) List(opts v1.ListOptions) (result *v1alpha1.EngineList, err error) {
	result = &v1alpha1.EngineList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("engines").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested engines.
func (c *engines) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("engines").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a engine and creates it.  Returns the server's representation of the engine, and an error, if there is any.
func (c *engines) Create(engine *v1alpha1.Engine) (result *v1alpha1.Engine, err error) {
	result = &v1alpha1.Engine{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("engines").
		Body(engine).
		Do().
		Into(result)
	return
}

// Update takes the representation of a engine and updates it. Returns the server's representation of the engine, and an error, if there is any.
func (c *engines) Update(engine *v1alpha1.Engine) (result *v1alpha1.Engine, err error) {
	result = &v1alpha1.Engine{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("engines").
		Name(engine.Name).
		Body(engine).
		Do().
		Into(result)
	return
}

// Delete takes name of the engine and deletes it. Returns an error if one occurs.
func (c *engines) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("engines").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *engines) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("engines").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched engine.
func (c *engines) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Engine, err error) {
	result = &v1alpha1.Engine{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("engines").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}

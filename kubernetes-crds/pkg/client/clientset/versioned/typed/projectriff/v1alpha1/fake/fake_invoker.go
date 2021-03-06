/*
 * Copyright 2018 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package fake

import (
	v1alpha1 "github.com/projectriff/riff/kubernetes-crds/pkg/apis/projectriff.io/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeInvokers implements InvokerInterface
type FakeInvokers struct {
	Fake *FakeProjectriffV1alpha1
	ns   string
}

var invokersResource = schema.GroupVersionResource{Group: "projectriff.io", Version: "v1alpha1", Resource: "invokers"}

var invokersKind = schema.GroupVersionKind{Group: "projectriff.io", Version: "v1alpha1", Kind: "Invoker"}

// Get takes name of the invoker, and returns the corresponding invoker object, and an error if there is any.
func (c *FakeInvokers) Get(name string, options v1.GetOptions) (result *v1alpha1.Invoker, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(invokersResource, c.ns, name), &v1alpha1.Invoker{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Invoker), err
}

// List takes label and field selectors, and returns the list of Invokers that match those selectors.
func (c *FakeInvokers) List(opts v1.ListOptions) (result *v1alpha1.InvokerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(invokersResource, invokersKind, c.ns, opts), &v1alpha1.InvokerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.InvokerList{}
	for _, item := range obj.(*v1alpha1.InvokerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested invokers.
func (c *FakeInvokers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(invokersResource, c.ns, opts))

}

// Create takes the representation of a invoker and creates it.  Returns the server's representation of the invoker, and an error, if there is any.
func (c *FakeInvokers) Create(invoker *v1alpha1.Invoker) (result *v1alpha1.Invoker, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(invokersResource, c.ns, invoker), &v1alpha1.Invoker{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Invoker), err
}

// Update takes the representation of a invoker and updates it. Returns the server's representation of the invoker, and an error, if there is any.
func (c *FakeInvokers) Update(invoker *v1alpha1.Invoker) (result *v1alpha1.Invoker, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(invokersResource, c.ns, invoker), &v1alpha1.Invoker{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Invoker), err
}

// Delete takes name of the invoker and deletes it. Returns an error if one occurs.
func (c *FakeInvokers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(invokersResource, c.ns, name), &v1alpha1.Invoker{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeInvokers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(invokersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.InvokerList{})
	return err
}

// Patch applies the patch and returns the patched invoker.
func (c *FakeInvokers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Invoker, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(invokersResource, c.ns, name, data, subresources...), &v1alpha1.Invoker{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Invoker), err
}

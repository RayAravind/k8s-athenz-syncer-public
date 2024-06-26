/*
Copyright 2019, Oath Inc.

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	athenzv1 "github.com/AthenZ/k8s-athenz-syncer/pkg/apis/athenz/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAthenzDomains implements AthenzDomainInterface
type FakeAthenzDomains struct {
	Fake *FakeAthenzV1
}

var athenzdomainsResource = schema.GroupVersionResource{Group: "athenz", Version: "v1", Resource: "athenzdomains"}

var athenzdomainsKind = schema.GroupVersionKind{Group: "athenz", Version: "v1", Kind: "AthenzDomain"}

// Get takes name of the athenzDomain, and returns the corresponding athenzDomain object, and an error if there is any.
func (c *FakeAthenzDomains) Get(ctx context.Context, name string, options v1.GetOptions) (result *athenzv1.AthenzDomain, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(athenzdomainsResource, name), &athenzv1.AthenzDomain{})
	if obj == nil {
		return nil, err
	}
	return obj.(*athenzv1.AthenzDomain), err
}

// List takes label and field selectors, and returns the list of AthenzDomains that match those selectors.
func (c *FakeAthenzDomains) List(ctx context.Context, opts v1.ListOptions) (result *athenzv1.AthenzDomainList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(athenzdomainsResource, athenzdomainsKind, opts), &athenzv1.AthenzDomainList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &athenzv1.AthenzDomainList{ListMeta: obj.(*athenzv1.AthenzDomainList).ListMeta}
	for _, item := range obj.(*athenzv1.AthenzDomainList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested athenzDomains.
func (c *FakeAthenzDomains) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(athenzdomainsResource, opts))
}

// Create takes the representation of a athenzDomain and creates it.  Returns the server's representation of the athenzDomain, and an error, if there is any.
func (c *FakeAthenzDomains) Create(ctx context.Context, athenzDomain *athenzv1.AthenzDomain, opts v1.CreateOptions) (result *athenzv1.AthenzDomain, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(athenzdomainsResource, athenzDomain), &athenzv1.AthenzDomain{})
	if obj == nil {
		return nil, err
	}
	return obj.(*athenzv1.AthenzDomain), err
}

// Update takes the representation of a athenzDomain and updates it. Returns the server's representation of the athenzDomain, and an error, if there is any.
func (c *FakeAthenzDomains) Update(ctx context.Context, athenzDomain *athenzv1.AthenzDomain, opts v1.UpdateOptions) (result *athenzv1.AthenzDomain, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(athenzdomainsResource, athenzDomain), &athenzv1.AthenzDomain{})
	if obj == nil {
		return nil, err
	}
	return obj.(*athenzv1.AthenzDomain), err
}

// Delete takes name of the athenzDomain and deletes it. Returns an error if one occurs.
func (c *FakeAthenzDomains) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(athenzdomainsResource, name), &athenzv1.AthenzDomain{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAthenzDomains) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(athenzdomainsResource, listOpts)

	_, err := c.Fake.Invokes(action, &athenzv1.AthenzDomainList{})
	return err
}

// Patch applies the patch and returns the patched athenzDomain.
func (c *FakeAthenzDomains) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *athenzv1.AthenzDomain, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(athenzdomainsResource, name, pt, data, subresources...), &athenzv1.AthenzDomain{})
	if obj == nil {
		return nil, err
	}
	return obj.(*athenzv1.AthenzDomain), err
}

/*
Copyright 2021 The KubeVela Authors.

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

	v1beta1 "github.com/napptive/kubevela-core-api/apis/core.oam.dev/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeResourceTrackers implements ResourceTrackerInterface
type FakeResourceTrackers struct {
	Fake *FakeCoreV1beta1
	ns   string
}

var resourcetrackersResource = schema.GroupVersionResource{Group: "core.oam.dev", Version: "v1beta1", Resource: "resourcetrackers"}

var resourcetrackersKind = schema.GroupVersionKind{Group: "core.oam.dev", Version: "v1beta1", Kind: "ResourceTracker"}

// Get takes name of the resourceTracker, and returns the corresponding resourceTracker object, and an error if there is any.
func (c *FakeResourceTrackers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.ResourceTracker, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(resourcetrackersResource, c.ns, name), &v1beta1.ResourceTracker{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ResourceTracker), err
}

// List takes label and field selectors, and returns the list of ResourceTrackers that match those selectors.
func (c *FakeResourceTrackers) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.ResourceTrackerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(resourcetrackersResource, resourcetrackersKind, c.ns, opts), &v1beta1.ResourceTrackerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.ResourceTrackerList{ListMeta: obj.(*v1beta1.ResourceTrackerList).ListMeta}
	for _, item := range obj.(*v1beta1.ResourceTrackerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested resourceTrackers.
func (c *FakeResourceTrackers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(resourcetrackersResource, c.ns, opts))

}

// Create takes the representation of a resourceTracker and creates it.  Returns the server's representation of the resourceTracker, and an error, if there is any.
func (c *FakeResourceTrackers) Create(ctx context.Context, resourceTracker *v1beta1.ResourceTracker, opts v1.CreateOptions) (result *v1beta1.ResourceTracker, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(resourcetrackersResource, c.ns, resourceTracker), &v1beta1.ResourceTracker{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ResourceTracker), err
}

// Update takes the representation of a resourceTracker and updates it. Returns the server's representation of the resourceTracker, and an error, if there is any.
func (c *FakeResourceTrackers) Update(ctx context.Context, resourceTracker *v1beta1.ResourceTracker, opts v1.UpdateOptions) (result *v1beta1.ResourceTracker, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(resourcetrackersResource, c.ns, resourceTracker), &v1beta1.ResourceTracker{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ResourceTracker), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeResourceTrackers) UpdateStatus(ctx context.Context, resourceTracker *v1beta1.ResourceTracker, opts v1.UpdateOptions) (*v1beta1.ResourceTracker, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(resourcetrackersResource, "status", c.ns, resourceTracker), &v1beta1.ResourceTracker{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ResourceTracker), err
}

// Delete takes name of the resourceTracker and deletes it. Returns an error if one occurs.
func (c *FakeResourceTrackers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(resourcetrackersResource, c.ns, name), &v1beta1.ResourceTracker{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeResourceTrackers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(resourcetrackersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.ResourceTrackerList{})
	return err
}

// Patch applies the patch and returns the patched resourceTracker.
func (c *FakeResourceTrackers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.ResourceTracker, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(resourcetrackersResource, c.ns, name, pt, data, subresources...), &v1beta1.ResourceTracker{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ResourceTracker), err
}

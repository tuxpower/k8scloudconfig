/*
Copyright 2017 The Kubernetes Authors.

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
	v1alpha1 "github.com/giantswarm/apiextensions/pkg/apis/core/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCertConfigs implements CertConfigInterface
type FakeCertConfigs struct {
	Fake *FakeCoreV1alpha1
	ns   string
}

var certconfigsResource = schema.GroupVersionResource{Group: "core.giantswarm.io", Version: "v1alpha1", Resource: "certconfigs"}

var certconfigsKind = schema.GroupVersionKind{Group: "core.giantswarm.io", Version: "v1alpha1", Kind: "CertConfig"}

// Get takes name of the certConfig, and returns the corresponding certConfig object, and an error if there is any.
func (c *FakeCertConfigs) Get(name string, options v1.GetOptions) (result *v1alpha1.CertConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(certconfigsResource, c.ns, name), &v1alpha1.CertConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CertConfig), err
}

// List takes label and field selectors, and returns the list of CertConfigs that match those selectors.
func (c *FakeCertConfigs) List(opts v1.ListOptions) (result *v1alpha1.CertConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(certconfigsResource, certconfigsKind, c.ns, opts), &v1alpha1.CertConfigList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CertConfigList{}
	for _, item := range obj.(*v1alpha1.CertConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested certConfigs.
func (c *FakeCertConfigs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(certconfigsResource, c.ns, opts))

}

// Create takes the representation of a certConfig and creates it.  Returns the server's representation of the certConfig, and an error, if there is any.
func (c *FakeCertConfigs) Create(certConfig *v1alpha1.CertConfig) (result *v1alpha1.CertConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(certconfigsResource, c.ns, certConfig), &v1alpha1.CertConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CertConfig), err
}

// Update takes the representation of a certConfig and updates it. Returns the server's representation of the certConfig, and an error, if there is any.
func (c *FakeCertConfigs) Update(certConfig *v1alpha1.CertConfig) (result *v1alpha1.CertConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(certconfigsResource, c.ns, certConfig), &v1alpha1.CertConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CertConfig), err
}

// Delete takes name of the certConfig and deletes it. Returns an error if one occurs.
func (c *FakeCertConfigs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(certconfigsResource, c.ns, name), &v1alpha1.CertConfig{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCertConfigs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(certconfigsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.CertConfigList{})
	return err
}

// Patch applies the patch and returns the patched certConfig.
func (c *FakeCertConfigs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CertConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(certconfigsResource, c.ns, name, data, subresources...), &v1alpha1.CertConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CertConfig), err
}

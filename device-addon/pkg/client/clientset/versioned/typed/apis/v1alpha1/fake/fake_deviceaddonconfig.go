// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	v1alpha1 "open-cluster-management-io/addon-contrib/device-addon/pkg/apis/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDeviceAddOnConfigs implements DeviceAddOnConfigInterface
type FakeDeviceAddOnConfigs struct {
	Fake *FakeEdgeV1alpha1
	ns   string
}

var deviceaddonconfigsResource = v1alpha1.SchemeGroupVersion.WithResource("deviceaddonconfigs")

var deviceaddonconfigsKind = v1alpha1.SchemeGroupVersion.WithKind("DeviceAddOnConfig")

// Get takes name of the deviceAddOnConfig, and returns the corresponding deviceAddOnConfig object, and an error if there is any.
func (c *FakeDeviceAddOnConfigs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.DeviceAddOnConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(deviceaddonconfigsResource, c.ns, name), &v1alpha1.DeviceAddOnConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DeviceAddOnConfig), err
}

// List takes label and field selectors, and returns the list of DeviceAddOnConfigs that match those selectors.
func (c *FakeDeviceAddOnConfigs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DeviceAddOnConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(deviceaddonconfigsResource, deviceaddonconfigsKind, c.ns, opts), &v1alpha1.DeviceAddOnConfigList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DeviceAddOnConfigList{ListMeta: obj.(*v1alpha1.DeviceAddOnConfigList).ListMeta}
	for _, item := range obj.(*v1alpha1.DeviceAddOnConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested deviceAddOnConfigs.
func (c *FakeDeviceAddOnConfigs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(deviceaddonconfigsResource, c.ns, opts))

}

// Create takes the representation of a deviceAddOnConfig and creates it.  Returns the server's representation of the deviceAddOnConfig, and an error, if there is any.
func (c *FakeDeviceAddOnConfigs) Create(ctx context.Context, deviceAddOnConfig *v1alpha1.DeviceAddOnConfig, opts v1.CreateOptions) (result *v1alpha1.DeviceAddOnConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(deviceaddonconfigsResource, c.ns, deviceAddOnConfig), &v1alpha1.DeviceAddOnConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DeviceAddOnConfig), err
}

// Update takes the representation of a deviceAddOnConfig and updates it. Returns the server's representation of the deviceAddOnConfig, and an error, if there is any.
func (c *FakeDeviceAddOnConfigs) Update(ctx context.Context, deviceAddOnConfig *v1alpha1.DeviceAddOnConfig, opts v1.UpdateOptions) (result *v1alpha1.DeviceAddOnConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(deviceaddonconfigsResource, c.ns, deviceAddOnConfig), &v1alpha1.DeviceAddOnConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DeviceAddOnConfig), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDeviceAddOnConfigs) UpdateStatus(ctx context.Context, deviceAddOnConfig *v1alpha1.DeviceAddOnConfig, opts v1.UpdateOptions) (*v1alpha1.DeviceAddOnConfig, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(deviceaddonconfigsResource, "status", c.ns, deviceAddOnConfig), &v1alpha1.DeviceAddOnConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DeviceAddOnConfig), err
}

// Delete takes name of the deviceAddOnConfig and deletes it. Returns an error if one occurs.
func (c *FakeDeviceAddOnConfigs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(deviceaddonconfigsResource, c.ns, name, opts), &v1alpha1.DeviceAddOnConfig{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDeviceAddOnConfigs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(deviceaddonconfigsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.DeviceAddOnConfigList{})
	return err
}

// Patch applies the patch and returns the patched deviceAddOnConfig.
func (c *FakeDeviceAddOnConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DeviceAddOnConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(deviceaddonconfigsResource, c.ns, name, pt, data, subresources...), &v1alpha1.DeviceAddOnConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DeviceAddOnConfig), err
}

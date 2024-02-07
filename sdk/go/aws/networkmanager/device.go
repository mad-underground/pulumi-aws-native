// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package networkmanager

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The AWS::NetworkManager::Device type describes a device.
type Device struct {
	pulumi.CustomResourceState

	// The Amazon Web Services location of the device, if applicable.
	AwsLocation DeviceAwsLocationPtrOutput `pulumi:"awsLocation"`
	// The date and time that the device was created.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The description of the device.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The Amazon Resource Name (ARN) of the device.
	DeviceArn pulumi.StringOutput `pulumi:"deviceArn"`
	// The ID of the device.
	DeviceId pulumi.StringOutput `pulumi:"deviceId"`
	// The ID of the global network.
	GlobalNetworkId pulumi.StringOutput `pulumi:"globalNetworkId"`
	// The site location.
	Location DeviceLocationPtrOutput `pulumi:"location"`
	// The device model
	Model pulumi.StringPtrOutput `pulumi:"model"`
	// The device serial number.
	SerialNumber pulumi.StringPtrOutput `pulumi:"serialNumber"`
	// The site ID.
	SiteId pulumi.StringPtrOutput `pulumi:"siteId"`
	// The state of the device.
	State pulumi.StringOutput `pulumi:"state"`
	// The tags for the device.
	Tags DeviceTagArrayOutput `pulumi:"tags"`
	// The device type.
	Type pulumi.StringPtrOutput `pulumi:"type"`
	// The device vendor.
	Vendor pulumi.StringPtrOutput `pulumi:"vendor"`
}

// NewDevice registers a new resource with the given unique name, arguments, and options.
func NewDevice(ctx *pulumi.Context,
	name string, args *DeviceArgs, opts ...pulumi.ResourceOption) (*Device, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GlobalNetworkId == nil {
		return nil, errors.New("invalid value for required argument 'GlobalNetworkId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"globalNetworkId",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Device
	err := ctx.RegisterResource("aws-native:networkmanager:Device", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDevice gets an existing Device resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDevice(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeviceState, opts ...pulumi.ResourceOption) (*Device, error) {
	var resource Device
	err := ctx.ReadResource("aws-native:networkmanager:Device", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Device resources.
type deviceState struct {
}

type DeviceState struct {
}

func (DeviceState) ElementType() reflect.Type {
	return reflect.TypeOf((*deviceState)(nil)).Elem()
}

type deviceArgs struct {
	// The Amazon Web Services location of the device, if applicable.
	AwsLocation *DeviceAwsLocation `pulumi:"awsLocation"`
	// The description of the device.
	Description *string `pulumi:"description"`
	// The ID of the global network.
	GlobalNetworkId string `pulumi:"globalNetworkId"`
	// The site location.
	Location *DeviceLocation `pulumi:"location"`
	// The device model
	Model *string `pulumi:"model"`
	// The device serial number.
	SerialNumber *string `pulumi:"serialNumber"`
	// The site ID.
	SiteId *string `pulumi:"siteId"`
	// The tags for the device.
	Tags []DeviceTag `pulumi:"tags"`
	// The device type.
	Type *string `pulumi:"type"`
	// The device vendor.
	Vendor *string `pulumi:"vendor"`
}

// The set of arguments for constructing a Device resource.
type DeviceArgs struct {
	// The Amazon Web Services location of the device, if applicable.
	AwsLocation DeviceAwsLocationPtrInput
	// The description of the device.
	Description pulumi.StringPtrInput
	// The ID of the global network.
	GlobalNetworkId pulumi.StringInput
	// The site location.
	Location DeviceLocationPtrInput
	// The device model
	Model pulumi.StringPtrInput
	// The device serial number.
	SerialNumber pulumi.StringPtrInput
	// The site ID.
	SiteId pulumi.StringPtrInput
	// The tags for the device.
	Tags DeviceTagArrayInput
	// The device type.
	Type pulumi.StringPtrInput
	// The device vendor.
	Vendor pulumi.StringPtrInput
}

func (DeviceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*deviceArgs)(nil)).Elem()
}

type DeviceInput interface {
	pulumi.Input

	ToDeviceOutput() DeviceOutput
	ToDeviceOutputWithContext(ctx context.Context) DeviceOutput
}

func (*Device) ElementType() reflect.Type {
	return reflect.TypeOf((**Device)(nil)).Elem()
}

func (i *Device) ToDeviceOutput() DeviceOutput {
	return i.ToDeviceOutputWithContext(context.Background())
}

func (i *Device) ToDeviceOutputWithContext(ctx context.Context) DeviceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeviceOutput)
}

type DeviceOutput struct{ *pulumi.OutputState }

func (DeviceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Device)(nil)).Elem()
}

func (o DeviceOutput) ToDeviceOutput() DeviceOutput {
	return o
}

func (o DeviceOutput) ToDeviceOutputWithContext(ctx context.Context) DeviceOutput {
	return o
}

// The Amazon Web Services location of the device, if applicable.
func (o DeviceOutput) AwsLocation() DeviceAwsLocationPtrOutput {
	return o.ApplyT(func(v *Device) DeviceAwsLocationPtrOutput { return v.AwsLocation }).(DeviceAwsLocationPtrOutput)
}

// The date and time that the device was created.
func (o DeviceOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Device) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The description of the device.
func (o DeviceOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Device) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The Amazon Resource Name (ARN) of the device.
func (o DeviceOutput) DeviceArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Device) pulumi.StringOutput { return v.DeviceArn }).(pulumi.StringOutput)
}

// The ID of the device.
func (o DeviceOutput) DeviceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Device) pulumi.StringOutput { return v.DeviceId }).(pulumi.StringOutput)
}

// The ID of the global network.
func (o DeviceOutput) GlobalNetworkId() pulumi.StringOutput {
	return o.ApplyT(func(v *Device) pulumi.StringOutput { return v.GlobalNetworkId }).(pulumi.StringOutput)
}

// The site location.
func (o DeviceOutput) Location() DeviceLocationPtrOutput {
	return o.ApplyT(func(v *Device) DeviceLocationPtrOutput { return v.Location }).(DeviceLocationPtrOutput)
}

// The device model
func (o DeviceOutput) Model() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Device) pulumi.StringPtrOutput { return v.Model }).(pulumi.StringPtrOutput)
}

// The device serial number.
func (o DeviceOutput) SerialNumber() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Device) pulumi.StringPtrOutput { return v.SerialNumber }).(pulumi.StringPtrOutput)
}

// The site ID.
func (o DeviceOutput) SiteId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Device) pulumi.StringPtrOutput { return v.SiteId }).(pulumi.StringPtrOutput)
}

// The state of the device.
func (o DeviceOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *Device) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// The tags for the device.
func (o DeviceOutput) Tags() DeviceTagArrayOutput {
	return o.ApplyT(func(v *Device) DeviceTagArrayOutput { return v.Tags }).(DeviceTagArrayOutput)
}

// The device type.
func (o DeviceOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Device) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

// The device vendor.
func (o DeviceOutput) Vendor() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Device) pulumi.StringPtrOutput { return v.Vendor }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DeviceInput)(nil)).Elem(), &Device{})
	pulumi.RegisterOutputType(DeviceOutput{})
}

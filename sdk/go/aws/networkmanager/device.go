// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package networkmanager

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-device.html
type Device struct {
	pulumi.CustomResourceState

	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-device.html#cfn-networkmanager-device-description
	Description pulumi.StringPtrOutput `pulumi:"description"`
	DeviceArn   pulumi.StringOutput    `pulumi:"deviceArn"`
	DeviceId    pulumi.StringOutput    `pulumi:"deviceId"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-device.html#cfn-networkmanager-device-globalnetworkid
	GlobalNetworkId pulumi.StringOutput `pulumi:"globalNetworkId"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-device.html#cfn-networkmanager-device-location
	Location DeviceLocationPtrOutput `pulumi:"location"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-device.html#cfn-networkmanager-device-model
	Model pulumi.StringPtrOutput `pulumi:"model"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-device.html#cfn-networkmanager-device-serialnumber
	SerialNumber pulumi.StringPtrOutput `pulumi:"serialNumber"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-device.html#cfn-networkmanager-device-siteid
	SiteId pulumi.StringPtrOutput `pulumi:"siteId"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-device.html#cfn-networkmanager-device-tags
	Tags aws.TagArrayOutput `pulumi:"tags"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-device.html#cfn-networkmanager-device-type
	Type pulumi.StringPtrOutput `pulumi:"type"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-device.html#cfn-networkmanager-device-vendor
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
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-device.html#cfn-networkmanager-device-description
	Description *string `pulumi:"description"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-device.html#cfn-networkmanager-device-globalnetworkid
	GlobalNetworkId string `pulumi:"globalNetworkId"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-device.html#cfn-networkmanager-device-location
	Location *DeviceLocation `pulumi:"location"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-device.html#cfn-networkmanager-device-model
	Model *string `pulumi:"model"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-device.html#cfn-networkmanager-device-serialnumber
	SerialNumber *string `pulumi:"serialNumber"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-device.html#cfn-networkmanager-device-siteid
	SiteId *string `pulumi:"siteId"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-device.html#cfn-networkmanager-device-tags
	Tags []aws.Tag `pulumi:"tags"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-device.html#cfn-networkmanager-device-type
	Type *string `pulumi:"type"`
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-device.html#cfn-networkmanager-device-vendor
	Vendor *string `pulumi:"vendor"`
}

// The set of arguments for constructing a Device resource.
type DeviceArgs struct {
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-device.html#cfn-networkmanager-device-description
	Description pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-device.html#cfn-networkmanager-device-globalnetworkid
	GlobalNetworkId pulumi.StringInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-device.html#cfn-networkmanager-device-location
	Location DeviceLocationPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-device.html#cfn-networkmanager-device-model
	Model pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-device.html#cfn-networkmanager-device-serialnumber
	SerialNumber pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-device.html#cfn-networkmanager-device-siteid
	SiteId pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-device.html#cfn-networkmanager-device-tags
	Tags aws.TagArrayInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-device.html#cfn-networkmanager-device-type
	Type pulumi.StringPtrInput
	// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-device.html#cfn-networkmanager-device-vendor
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
	return reflect.TypeOf((*Device)(nil))
}

func (i *Device) ToDeviceOutput() DeviceOutput {
	return i.ToDeviceOutputWithContext(context.Background())
}

func (i *Device) ToDeviceOutputWithContext(ctx context.Context) DeviceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeviceOutput)
}

type DeviceOutput struct{ *pulumi.OutputState }

func (DeviceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Device)(nil))
}

func (o DeviceOutput) ToDeviceOutput() DeviceOutput {
	return o
}

func (o DeviceOutput) ToDeviceOutputWithContext(ctx context.Context) DeviceOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DeviceOutput{})
}
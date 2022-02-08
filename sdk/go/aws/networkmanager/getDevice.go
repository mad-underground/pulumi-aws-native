// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package networkmanager

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The AWS::NetworkManager::Device type describes a device.
func LookupDevice(ctx *pulumi.Context, args *LookupDeviceArgs, opts ...pulumi.InvokeOption) (*LookupDeviceResult, error) {
	var rv LookupDeviceResult
	err := ctx.Invoke("aws-native:networkmanager:getDevice", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDeviceArgs struct {
	// The ID of the device.
	DeviceId string `pulumi:"deviceId"`
	// The ID of the global network.
	GlobalNetworkId string `pulumi:"globalNetworkId"`
}

type LookupDeviceResult struct {
	// The description of the device.
	Description *string `pulumi:"description"`
	// The Amazon Resource Name (ARN) of the device.
	DeviceArn *string `pulumi:"deviceArn"`
	// The ID of the device.
	DeviceId *string `pulumi:"deviceId"`
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

func LookupDeviceOutput(ctx *pulumi.Context, args LookupDeviceOutputArgs, opts ...pulumi.InvokeOption) LookupDeviceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDeviceResult, error) {
			args := v.(LookupDeviceArgs)
			r, err := LookupDevice(ctx, &args, opts...)
			return *r, err
		}).(LookupDeviceResultOutput)
}

type LookupDeviceOutputArgs struct {
	// The ID of the device.
	DeviceId pulumi.StringInput `pulumi:"deviceId"`
	// The ID of the global network.
	GlobalNetworkId pulumi.StringInput `pulumi:"globalNetworkId"`
}

func (LookupDeviceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDeviceArgs)(nil)).Elem()
}

type LookupDeviceResultOutput struct{ *pulumi.OutputState }

func (LookupDeviceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDeviceResult)(nil)).Elem()
}

func (o LookupDeviceResultOutput) ToLookupDeviceResultOutput() LookupDeviceResultOutput {
	return o
}

func (o LookupDeviceResultOutput) ToLookupDeviceResultOutputWithContext(ctx context.Context) LookupDeviceResultOutput {
	return o
}

// The description of the device.
func (o LookupDeviceResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDeviceResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The Amazon Resource Name (ARN) of the device.
func (o LookupDeviceResultOutput) DeviceArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDeviceResult) *string { return v.DeviceArn }).(pulumi.StringPtrOutput)
}

// The ID of the device.
func (o LookupDeviceResultOutput) DeviceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDeviceResult) *string { return v.DeviceId }).(pulumi.StringPtrOutput)
}

// The site location.
func (o LookupDeviceResultOutput) Location() DeviceLocationPtrOutput {
	return o.ApplyT(func(v LookupDeviceResult) *DeviceLocation { return v.Location }).(DeviceLocationPtrOutput)
}

// The device model
func (o LookupDeviceResultOutput) Model() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDeviceResult) *string { return v.Model }).(pulumi.StringPtrOutput)
}

// The device serial number.
func (o LookupDeviceResultOutput) SerialNumber() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDeviceResult) *string { return v.SerialNumber }).(pulumi.StringPtrOutput)
}

// The site ID.
func (o LookupDeviceResultOutput) SiteId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDeviceResult) *string { return v.SiteId }).(pulumi.StringPtrOutput)
}

// The tags for the device.
func (o LookupDeviceResultOutput) Tags() DeviceTagArrayOutput {
	return o.ApplyT(func(v LookupDeviceResult) []DeviceTag { return v.Tags }).(DeviceTagArrayOutput)
}

// The device type.
func (o LookupDeviceResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDeviceResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

// The device vendor.
func (o LookupDeviceResultOutput) Vendor() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDeviceResult) *string { return v.Vendor }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDeviceResultOutput{})
}

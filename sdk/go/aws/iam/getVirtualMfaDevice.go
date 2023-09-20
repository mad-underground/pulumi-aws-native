// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Resource Type definition for AWS::IAM::VirtualMFADevice
func LookupVirtualMfaDevice(ctx *pulumi.Context, args *LookupVirtualMfaDeviceArgs, opts ...pulumi.InvokeOption) (*LookupVirtualMfaDeviceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupVirtualMfaDeviceResult
	err := ctx.Invoke("aws-native:iam:getVirtualMfaDevice", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVirtualMfaDeviceArgs struct {
	SerialNumber string `pulumi:"serialNumber"`
}

type LookupVirtualMfaDeviceResult struct {
	SerialNumber *string               `pulumi:"serialNumber"`
	Tags         []VirtualMfaDeviceTag `pulumi:"tags"`
	Users        []string              `pulumi:"users"`
}

func LookupVirtualMfaDeviceOutput(ctx *pulumi.Context, args LookupVirtualMfaDeviceOutputArgs, opts ...pulumi.InvokeOption) LookupVirtualMfaDeviceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVirtualMfaDeviceResult, error) {
			args := v.(LookupVirtualMfaDeviceArgs)
			r, err := LookupVirtualMfaDevice(ctx, &args, opts...)
			var s LookupVirtualMfaDeviceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVirtualMfaDeviceResultOutput)
}

type LookupVirtualMfaDeviceOutputArgs struct {
	SerialNumber pulumi.StringInput `pulumi:"serialNumber"`
}

func (LookupVirtualMfaDeviceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualMfaDeviceArgs)(nil)).Elem()
}

type LookupVirtualMfaDeviceResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualMfaDeviceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualMfaDeviceResult)(nil)).Elem()
}

func (o LookupVirtualMfaDeviceResultOutput) ToLookupVirtualMfaDeviceResultOutput() LookupVirtualMfaDeviceResultOutput {
	return o
}

func (o LookupVirtualMfaDeviceResultOutput) ToLookupVirtualMfaDeviceResultOutputWithContext(ctx context.Context) LookupVirtualMfaDeviceResultOutput {
	return o
}

func (o LookupVirtualMfaDeviceResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupVirtualMfaDeviceResult] {
	return pulumix.Output[LookupVirtualMfaDeviceResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupVirtualMfaDeviceResultOutput) SerialNumber() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMfaDeviceResult) *string { return v.SerialNumber }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMfaDeviceResultOutput) Tags() VirtualMfaDeviceTagArrayOutput {
	return o.ApplyT(func(v LookupVirtualMfaDeviceResult) []VirtualMfaDeviceTag { return v.Tags }).(VirtualMfaDeviceTagArrayOutput)
}

func (o LookupVirtualMfaDeviceResultOutput) Users() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVirtualMfaDeviceResult) []string { return v.Users }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualMfaDeviceResultOutput{})
}
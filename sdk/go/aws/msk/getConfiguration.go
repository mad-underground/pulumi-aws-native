// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package msk

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::MSK::Configuration
func LookupConfiguration(ctx *pulumi.Context, args *LookupConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupConfigurationResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupConfigurationResult
	err := ctx.Invoke("aws-native:msk:getConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupConfigurationArgs struct {
	Arn string `pulumi:"arn"`
}

type LookupConfigurationResult struct {
	Arn         *string `pulumi:"arn"`
	Description *string `pulumi:"description"`
}

func LookupConfigurationOutput(ctx *pulumi.Context, args LookupConfigurationOutputArgs, opts ...pulumi.InvokeOption) LookupConfigurationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupConfigurationResult, error) {
			args := v.(LookupConfigurationArgs)
			r, err := LookupConfiguration(ctx, &args, opts...)
			var s LookupConfigurationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupConfigurationResultOutput)
}

type LookupConfigurationOutputArgs struct {
	Arn pulumi.StringInput `pulumi:"arn"`
}

func (LookupConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConfigurationArgs)(nil)).Elem()
}

type LookupConfigurationResultOutput struct{ *pulumi.OutputState }

func (LookupConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConfigurationResult)(nil)).Elem()
}

func (o LookupConfigurationResultOutput) ToLookupConfigurationResultOutput() LookupConfigurationResultOutput {
	return o
}

func (o LookupConfigurationResultOutput) ToLookupConfigurationResultOutputWithContext(ctx context.Context) LookupConfigurationResultOutput {
	return o
}

func (o LookupConfigurationResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConfigurationResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupConfigurationResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupConfigurationResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConfigurationResultOutput{})
}

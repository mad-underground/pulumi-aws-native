// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iotanalytics

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::IoTAnalytics::Channel
func LookupChannel(ctx *pulumi.Context, args *LookupChannelArgs, opts ...pulumi.InvokeOption) (*LookupChannelResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupChannelResult
	err := ctx.Invoke("aws-native:iotanalytics:getChannel", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupChannelArgs struct {
	ChannelName string `pulumi:"channelName"`
}

type LookupChannelResult struct {
	ChannelStorage  *ChannelStorage         `pulumi:"channelStorage"`
	Id              *string                 `pulumi:"id"`
	RetentionPeriod *ChannelRetentionPeriod `pulumi:"retentionPeriod"`
	Tags            []ChannelTag            `pulumi:"tags"`
}

func LookupChannelOutput(ctx *pulumi.Context, args LookupChannelOutputArgs, opts ...pulumi.InvokeOption) LookupChannelResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupChannelResult, error) {
			args := v.(LookupChannelArgs)
			r, err := LookupChannel(ctx, &args, opts...)
			var s LookupChannelResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupChannelResultOutput)
}

type LookupChannelOutputArgs struct {
	ChannelName pulumi.StringInput `pulumi:"channelName"`
}

func (LookupChannelOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupChannelArgs)(nil)).Elem()
}

type LookupChannelResultOutput struct{ *pulumi.OutputState }

func (LookupChannelResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupChannelResult)(nil)).Elem()
}

func (o LookupChannelResultOutput) ToLookupChannelResultOutput() LookupChannelResultOutput {
	return o
}

func (o LookupChannelResultOutput) ToLookupChannelResultOutputWithContext(ctx context.Context) LookupChannelResultOutput {
	return o
}

func (o LookupChannelResultOutput) ChannelStorage() ChannelStoragePtrOutput {
	return o.ApplyT(func(v LookupChannelResult) *ChannelStorage { return v.ChannelStorage }).(ChannelStoragePtrOutput)
}

func (o LookupChannelResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupChannelResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupChannelResultOutput) RetentionPeriod() ChannelRetentionPeriodPtrOutput {
	return o.ApplyT(func(v LookupChannelResult) *ChannelRetentionPeriod { return v.RetentionPeriod }).(ChannelRetentionPeriodPtrOutput)
}

func (o LookupChannelResultOutput) Tags() ChannelTagArrayOutput {
	return o.ApplyT(func(v LookupChannelResult) []ChannelTag { return v.Tags }).(ChannelTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupChannelResultOutput{})
}

// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package pinpoint

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Pinpoint::EmailChannel
func LookupEmailChannel(ctx *pulumi.Context, args *LookupEmailChannelArgs, opts ...pulumi.InvokeOption) (*LookupEmailChannelResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupEmailChannelResult
	err := ctx.Invoke("aws-native:pinpoint:getEmailChannel", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEmailChannelArgs struct {
	Id string `pulumi:"id"`
}

type LookupEmailChannelResult struct {
	ConfigurationSet *string `pulumi:"configurationSet"`
	Enabled          *bool   `pulumi:"enabled"`
	FromAddress      *string `pulumi:"fromAddress"`
	Id               *string `pulumi:"id"`
	Identity         *string `pulumi:"identity"`
	RoleArn          *string `pulumi:"roleArn"`
}

func LookupEmailChannelOutput(ctx *pulumi.Context, args LookupEmailChannelOutputArgs, opts ...pulumi.InvokeOption) LookupEmailChannelResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEmailChannelResult, error) {
			args := v.(LookupEmailChannelArgs)
			r, err := LookupEmailChannel(ctx, &args, opts...)
			var s LookupEmailChannelResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEmailChannelResultOutput)
}

type LookupEmailChannelOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupEmailChannelOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEmailChannelArgs)(nil)).Elem()
}

type LookupEmailChannelResultOutput struct{ *pulumi.OutputState }

func (LookupEmailChannelResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEmailChannelResult)(nil)).Elem()
}

func (o LookupEmailChannelResultOutput) ToLookupEmailChannelResultOutput() LookupEmailChannelResultOutput {
	return o
}

func (o LookupEmailChannelResultOutput) ToLookupEmailChannelResultOutputWithContext(ctx context.Context) LookupEmailChannelResultOutput {
	return o
}

func (o LookupEmailChannelResultOutput) ConfigurationSet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEmailChannelResult) *string { return v.ConfigurationSet }).(pulumi.StringPtrOutput)
}

func (o LookupEmailChannelResultOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupEmailChannelResult) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o LookupEmailChannelResultOutput) FromAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEmailChannelResult) *string { return v.FromAddress }).(pulumi.StringPtrOutput)
}

func (o LookupEmailChannelResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEmailChannelResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupEmailChannelResultOutput) Identity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEmailChannelResult) *string { return v.Identity }).(pulumi.StringPtrOutput)
}

func (o LookupEmailChannelResultOutput) RoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEmailChannelResult) *string { return v.RoleArn }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEmailChannelResultOutput{})
}

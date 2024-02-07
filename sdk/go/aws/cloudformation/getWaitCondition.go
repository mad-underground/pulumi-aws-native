// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudformation

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::CloudFormation::WaitCondition
func LookupWaitCondition(ctx *pulumi.Context, args *LookupWaitConditionArgs, opts ...pulumi.InvokeOption) (*LookupWaitConditionResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupWaitConditionResult
	err := ctx.Invoke("aws-native:cloudformation:getWaitCondition", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWaitConditionArgs struct {
	Id string `pulumi:"id"`
}

type LookupWaitConditionResult struct {
	Count   *int        `pulumi:"count"`
	Data    interface{} `pulumi:"data"`
	Handle  *string     `pulumi:"handle"`
	Id      *string     `pulumi:"id"`
	Timeout *string     `pulumi:"timeout"`
}

func LookupWaitConditionOutput(ctx *pulumi.Context, args LookupWaitConditionOutputArgs, opts ...pulumi.InvokeOption) LookupWaitConditionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWaitConditionResult, error) {
			args := v.(LookupWaitConditionArgs)
			r, err := LookupWaitCondition(ctx, &args, opts...)
			var s LookupWaitConditionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWaitConditionResultOutput)
}

type LookupWaitConditionOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupWaitConditionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWaitConditionArgs)(nil)).Elem()
}

type LookupWaitConditionResultOutput struct{ *pulumi.OutputState }

func (LookupWaitConditionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWaitConditionResult)(nil)).Elem()
}

func (o LookupWaitConditionResultOutput) ToLookupWaitConditionResultOutput() LookupWaitConditionResultOutput {
	return o
}

func (o LookupWaitConditionResultOutput) ToLookupWaitConditionResultOutputWithContext(ctx context.Context) LookupWaitConditionResultOutput {
	return o
}

func (o LookupWaitConditionResultOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupWaitConditionResult) *int { return v.Count }).(pulumi.IntPtrOutput)
}

func (o LookupWaitConditionResultOutput) Data() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupWaitConditionResult) interface{} { return v.Data }).(pulumi.AnyOutput)
}

func (o LookupWaitConditionResultOutput) Handle() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWaitConditionResult) *string { return v.Handle }).(pulumi.StringPtrOutput)
}

func (o LookupWaitConditionResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWaitConditionResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupWaitConditionResultOutput) Timeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWaitConditionResult) *string { return v.Timeout }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWaitConditionResultOutput{})
}

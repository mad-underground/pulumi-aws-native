// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::IAM::InstanceProfile
func LookupInstanceProfile(ctx *pulumi.Context, args *LookupInstanceProfileArgs, opts ...pulumi.InvokeOption) (*LookupInstanceProfileResult, error) {
	var rv LookupInstanceProfileResult
	err := ctx.Invoke("aws-native:iam:getInstanceProfile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupInstanceProfileArgs struct {
	Id string `pulumi:"id"`
}

type LookupInstanceProfileResult struct {
	Arn   *string  `pulumi:"arn"`
	Id    *string  `pulumi:"id"`
	Roles []string `pulumi:"roles"`
}

func LookupInstanceProfileOutput(ctx *pulumi.Context, args LookupInstanceProfileOutputArgs, opts ...pulumi.InvokeOption) LookupInstanceProfileResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupInstanceProfileResult, error) {
			args := v.(LookupInstanceProfileArgs)
			r, err := LookupInstanceProfile(ctx, &args, opts...)
			return *r, err
		}).(LookupInstanceProfileResultOutput)
}

type LookupInstanceProfileOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupInstanceProfileOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInstanceProfileArgs)(nil)).Elem()
}

type LookupInstanceProfileResultOutput struct{ *pulumi.OutputState }

func (LookupInstanceProfileResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInstanceProfileResult)(nil)).Elem()
}

func (o LookupInstanceProfileResultOutput) ToLookupInstanceProfileResultOutput() LookupInstanceProfileResultOutput {
	return o
}

func (o LookupInstanceProfileResultOutput) ToLookupInstanceProfileResultOutputWithContext(ctx context.Context) LookupInstanceProfileResultOutput {
	return o
}

func (o LookupInstanceProfileResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInstanceProfileResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupInstanceProfileResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInstanceProfileResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupInstanceProfileResultOutput) Roles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupInstanceProfileResult) []string { return v.Roles }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupInstanceProfileResultOutput{})
}
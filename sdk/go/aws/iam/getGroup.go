// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::IAM::Group
func LookupGroup(ctx *pulumi.Context, args *LookupGroupArgs, opts ...pulumi.InvokeOption) (*LookupGroupResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupGroupResult
	err := ctx.Invoke("aws-native:iam:getGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGroupArgs struct {
	Id string `pulumi:"id"`
}

type LookupGroupResult struct {
	Arn               *string           `pulumi:"arn"`
	Id                *string           `pulumi:"id"`
	ManagedPolicyArns []string          `pulumi:"managedPolicyArns"`
	Path              *string           `pulumi:"path"`
	Policies          []GroupPolicyType `pulumi:"policies"`
}

func LookupGroupOutput(ctx *pulumi.Context, args LookupGroupOutputArgs, opts ...pulumi.InvokeOption) LookupGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGroupResult, error) {
			args := v.(LookupGroupArgs)
			r, err := LookupGroup(ctx, &args, opts...)
			var s LookupGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGroupResultOutput)
}

type LookupGroupOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGroupArgs)(nil)).Elem()
}

type LookupGroupResultOutput struct{ *pulumi.OutputState }

func (LookupGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGroupResult)(nil)).Elem()
}

func (o LookupGroupResultOutput) ToLookupGroupResultOutput() LookupGroupResultOutput {
	return o
}

func (o LookupGroupResultOutput) ToLookupGroupResultOutputWithContext(ctx context.Context) LookupGroupResultOutput {
	return o
}

func (o LookupGroupResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGroupResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupGroupResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGroupResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupGroupResultOutput) ManagedPolicyArns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupGroupResult) []string { return v.ManagedPolicyArns }).(pulumi.StringArrayOutput)
}

func (o LookupGroupResultOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGroupResult) *string { return v.Path }).(pulumi.StringPtrOutput)
}

func (o LookupGroupResultOutput) Policies() GroupPolicyTypeArrayOutput {
	return o.ApplyT(func(v LookupGroupResult) []GroupPolicyType { return v.Policies }).(GroupPolicyTypeArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGroupResultOutput{})
}

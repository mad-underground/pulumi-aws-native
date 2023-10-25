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
	// The name of the group to create
	GroupName string `pulumi:"groupName"`
}

type LookupGroupResult struct {
	// The Arn of the group to create
	Arn *string `pulumi:"arn"`
	// A list of Amazon Resource Names (ARNs) of the IAM managed policies that you want to attach to the role.
	ManagedPolicyArns []string `pulumi:"managedPolicyArns"`
	// The path to the group
	Path *string `pulumi:"path"`
	// Adds or updates an inline policy document that is embedded in the specified IAM group
	Policies []GroupPolicyType `pulumi:"policies"`
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
	// The name of the group to create
	GroupName pulumi.StringInput `pulumi:"groupName"`
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

func (o LookupGroupResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupGroupResult] {
	return pulumix.Output[LookupGroupResult]{
		OutputState: o.OutputState,
	}
}

// The Arn of the group to create
func (o LookupGroupResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGroupResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// A list of Amazon Resource Names (ARNs) of the IAM managed policies that you want to attach to the role.
func (o LookupGroupResultOutput) ManagedPolicyArns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupGroupResult) []string { return v.ManagedPolicyArns }).(pulumi.StringArrayOutput)
}

// The path to the group
func (o LookupGroupResultOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGroupResult) *string { return v.Path }).(pulumi.StringPtrOutput)
}

// Adds or updates an inline policy document that is embedded in the specified IAM group
func (o LookupGroupResultOutput) Policies() GroupPolicyTypeArrayOutput {
	return o.ApplyT(func(v LookupGroupResult) []GroupPolicyType { return v.Policies }).(GroupPolicyTypeArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGroupResultOutput{})
}

// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dlm

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::DLM::LifecyclePolicy
func LookupLifecyclePolicy(ctx *pulumi.Context, args *LookupLifecyclePolicyArgs, opts ...pulumi.InvokeOption) (*LookupLifecyclePolicyResult, error) {
	var rv LookupLifecyclePolicyResult
	err := ctx.Invoke("aws-native:dlm:getLifecyclePolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLifecyclePolicyArgs struct {
	Id string `pulumi:"id"`
}

type LookupLifecyclePolicyResult struct {
	Arn              *string                       `pulumi:"arn"`
	Description      *string                       `pulumi:"description"`
	ExecutionRoleArn *string                       `pulumi:"executionRoleArn"`
	Id               *string                       `pulumi:"id"`
	PolicyDetails    *LifecyclePolicyPolicyDetails `pulumi:"policyDetails"`
	State            *string                       `pulumi:"state"`
	Tags             []LifecyclePolicyTag          `pulumi:"tags"`
}

func LookupLifecyclePolicyOutput(ctx *pulumi.Context, args LookupLifecyclePolicyOutputArgs, opts ...pulumi.InvokeOption) LookupLifecyclePolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLifecyclePolicyResult, error) {
			args := v.(LookupLifecyclePolicyArgs)
			r, err := LookupLifecyclePolicy(ctx, &args, opts...)
			return *r, err
		}).(LookupLifecyclePolicyResultOutput)
}

type LookupLifecyclePolicyOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupLifecyclePolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLifecyclePolicyArgs)(nil)).Elem()
}

type LookupLifecyclePolicyResultOutput struct{ *pulumi.OutputState }

func (LookupLifecyclePolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLifecyclePolicyResult)(nil)).Elem()
}

func (o LookupLifecyclePolicyResultOutput) ToLookupLifecyclePolicyResultOutput() LookupLifecyclePolicyResultOutput {
	return o
}

func (o LookupLifecyclePolicyResultOutput) ToLookupLifecyclePolicyResultOutputWithContext(ctx context.Context) LookupLifecyclePolicyResultOutput {
	return o
}

func (o LookupLifecyclePolicyResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLifecyclePolicyResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupLifecyclePolicyResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLifecyclePolicyResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupLifecyclePolicyResultOutput) ExecutionRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLifecyclePolicyResult) *string { return v.ExecutionRoleArn }).(pulumi.StringPtrOutput)
}

func (o LookupLifecyclePolicyResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLifecyclePolicyResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupLifecyclePolicyResultOutput) PolicyDetails() LifecyclePolicyPolicyDetailsPtrOutput {
	return o.ApplyT(func(v LookupLifecyclePolicyResult) *LifecyclePolicyPolicyDetails { return v.PolicyDetails }).(LifecyclePolicyPolicyDetailsPtrOutput)
}

func (o LookupLifecyclePolicyResultOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLifecyclePolicyResult) *string { return v.State }).(pulumi.StringPtrOutput)
}

func (o LookupLifecyclePolicyResultOutput) Tags() LifecyclePolicyTagArrayOutput {
	return o.ApplyT(func(v LookupLifecyclePolicyResult) []LifecyclePolicyTag { return v.Tags }).(LifecyclePolicyTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLifecyclePolicyResultOutput{})
}

// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package batch

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type schema for AWS::Batch::SchedulingPolicy
func LookupSchedulingPolicy(ctx *pulumi.Context, args *LookupSchedulingPolicyArgs, opts ...pulumi.InvokeOption) (*LookupSchedulingPolicyResult, error) {
	var rv LookupSchedulingPolicyResult
	err := ctx.Invoke("aws-native:batch:getSchedulingPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSchedulingPolicyArgs struct {
	Arn string `pulumi:"arn"`
}

type LookupSchedulingPolicyResult struct {
	Arn             *string                          `pulumi:"arn"`
	FairsharePolicy *SchedulingPolicyFairsharePolicy `pulumi:"fairsharePolicy"`
}

func LookupSchedulingPolicyOutput(ctx *pulumi.Context, args LookupSchedulingPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupSchedulingPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSchedulingPolicyResult, error) {
			args := v.(LookupSchedulingPolicyArgs)
			r, err := LookupSchedulingPolicy(ctx, &args, opts...)
			return *r, err
		}).(LookupSchedulingPolicyResultOutput)
}

type LookupSchedulingPolicyOutputArgs struct {
	Arn pulumi.StringInput `pulumi:"arn"`
}

func (LookupSchedulingPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSchedulingPolicyArgs)(nil)).Elem()
}

type LookupSchedulingPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupSchedulingPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSchedulingPolicyResult)(nil)).Elem()
}

func (o LookupSchedulingPolicyResultOutput) ToLookupSchedulingPolicyResultOutput() LookupSchedulingPolicyResultOutput {
	return o
}

func (o LookupSchedulingPolicyResultOutput) ToLookupSchedulingPolicyResultOutputWithContext(ctx context.Context) LookupSchedulingPolicyResultOutput {
	return o
}

func (o LookupSchedulingPolicyResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSchedulingPolicyResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupSchedulingPolicyResultOutput) FairsharePolicy() SchedulingPolicyFairsharePolicyPtrOutput {
	return o.ApplyT(func(v LookupSchedulingPolicyResult) *SchedulingPolicyFairsharePolicy { return v.FairsharePolicy }).(SchedulingPolicyFairsharePolicyPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSchedulingPolicyResultOutput{})
}

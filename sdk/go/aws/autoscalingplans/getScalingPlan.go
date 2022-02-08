// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package autoscalingplans

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::AutoScalingPlans::ScalingPlan
func LookupScalingPlan(ctx *pulumi.Context, args *LookupScalingPlanArgs, opts ...pulumi.InvokeOption) (*LookupScalingPlanResult, error) {
	var rv LookupScalingPlanResult
	err := ctx.Invoke("aws-native:autoscalingplans:getScalingPlan", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupScalingPlanArgs struct {
	Id string `pulumi:"id"`
}

type LookupScalingPlanResult struct {
	ApplicationSource   *ScalingPlanApplicationSource   `pulumi:"applicationSource"`
	Id                  *string                         `pulumi:"id"`
	ScalingInstructions []ScalingPlanScalingInstruction `pulumi:"scalingInstructions"`
	ScalingPlanName     *string                         `pulumi:"scalingPlanName"`
	ScalingPlanVersion  *string                         `pulumi:"scalingPlanVersion"`
}

func LookupScalingPlanOutput(ctx *pulumi.Context, args LookupScalingPlanOutputArgs, opts ...pulumi.InvokeOption) LookupScalingPlanResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupScalingPlanResult, error) {
			args := v.(LookupScalingPlanArgs)
			r, err := LookupScalingPlan(ctx, &args, opts...)
			return *r, err
		}).(LookupScalingPlanResultOutput)
}

type LookupScalingPlanOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupScalingPlanOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupScalingPlanArgs)(nil)).Elem()
}

type LookupScalingPlanResultOutput struct{ *pulumi.OutputState }

func (LookupScalingPlanResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupScalingPlanResult)(nil)).Elem()
}

func (o LookupScalingPlanResultOutput) ToLookupScalingPlanResultOutput() LookupScalingPlanResultOutput {
	return o
}

func (o LookupScalingPlanResultOutput) ToLookupScalingPlanResultOutputWithContext(ctx context.Context) LookupScalingPlanResultOutput {
	return o
}

func (o LookupScalingPlanResultOutput) ApplicationSource() ScalingPlanApplicationSourcePtrOutput {
	return o.ApplyT(func(v LookupScalingPlanResult) *ScalingPlanApplicationSource { return v.ApplicationSource }).(ScalingPlanApplicationSourcePtrOutput)
}

func (o LookupScalingPlanResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScalingPlanResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupScalingPlanResultOutput) ScalingInstructions() ScalingPlanScalingInstructionArrayOutput {
	return o.ApplyT(func(v LookupScalingPlanResult) []ScalingPlanScalingInstruction { return v.ScalingInstructions }).(ScalingPlanScalingInstructionArrayOutput)
}

func (o LookupScalingPlanResultOutput) ScalingPlanName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScalingPlanResult) *string { return v.ScalingPlanName }).(pulumi.StringPtrOutput)
}

func (o LookupScalingPlanResultOutput) ScalingPlanVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScalingPlanResult) *string { return v.ScalingPlanVersion }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupScalingPlanResultOutput{})
}

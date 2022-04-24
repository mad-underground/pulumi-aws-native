// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package billingconductor

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Pricing Plan enables you to customize your billing details consistent with the usage that accrues in each of your billing groups.
func LookupPricingPlan(ctx *pulumi.Context, args *LookupPricingPlanArgs, opts ...pulumi.InvokeOption) (*LookupPricingPlanResult, error) {
	var rv LookupPricingPlanResult
	err := ctx.Invoke("aws-native:billingconductor:getPricingPlan", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPricingPlanArgs struct {
	// Pricing Plan ARN
	Arn string `pulumi:"arn"`
}

type LookupPricingPlanResult struct {
	// Pricing Plan ARN
	Arn *string `pulumi:"arn"`
	// Creation timestamp in UNIX epoch time format
	CreationTime *int    `pulumi:"creationTime"`
	Description  *string `pulumi:"description"`
	// Latest modified timestamp in UNIX epoch time format
	LastModifiedTime *int     `pulumi:"lastModifiedTime"`
	Name             *string  `pulumi:"name"`
	PricingRuleArns  []string `pulumi:"pricingRuleArns"`
	// Number of associated pricing rules
	Size *int             `pulumi:"size"`
	Tags []PricingPlanTag `pulumi:"tags"`
}

func LookupPricingPlanOutput(ctx *pulumi.Context, args LookupPricingPlanOutputArgs, opts ...pulumi.InvokeOption) LookupPricingPlanResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPricingPlanResult, error) {
			args := v.(LookupPricingPlanArgs)
			r, err := LookupPricingPlan(ctx, &args, opts...)
			var s LookupPricingPlanResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPricingPlanResultOutput)
}

type LookupPricingPlanOutputArgs struct {
	// Pricing Plan ARN
	Arn pulumi.StringInput `pulumi:"arn"`
}

func (LookupPricingPlanOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPricingPlanArgs)(nil)).Elem()
}

type LookupPricingPlanResultOutput struct{ *pulumi.OutputState }

func (LookupPricingPlanResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPricingPlanResult)(nil)).Elem()
}

func (o LookupPricingPlanResultOutput) ToLookupPricingPlanResultOutput() LookupPricingPlanResultOutput {
	return o
}

func (o LookupPricingPlanResultOutput) ToLookupPricingPlanResultOutputWithContext(ctx context.Context) LookupPricingPlanResultOutput {
	return o
}

// Pricing Plan ARN
func (o LookupPricingPlanResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPricingPlanResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// Creation timestamp in UNIX epoch time format
func (o LookupPricingPlanResultOutput) CreationTime() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupPricingPlanResult) *int { return v.CreationTime }).(pulumi.IntPtrOutput)
}

func (o LookupPricingPlanResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPricingPlanResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Latest modified timestamp in UNIX epoch time format
func (o LookupPricingPlanResultOutput) LastModifiedTime() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupPricingPlanResult) *int { return v.LastModifiedTime }).(pulumi.IntPtrOutput)
}

func (o LookupPricingPlanResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPricingPlanResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupPricingPlanResultOutput) PricingRuleArns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPricingPlanResult) []string { return v.PricingRuleArns }).(pulumi.StringArrayOutput)
}

// Number of associated pricing rules
func (o LookupPricingPlanResultOutput) Size() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupPricingPlanResult) *int { return v.Size }).(pulumi.IntPtrOutput)
}

func (o LookupPricingPlanResultOutput) Tags() PricingPlanTagArrayOutput {
	return o.ApplyT(func(v LookupPricingPlanResult) []PricingPlanTag { return v.Tags }).(PricingPlanTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPricingPlanResultOutput{})
}

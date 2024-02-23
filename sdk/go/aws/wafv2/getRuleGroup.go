// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package wafv2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Contains the Rules that identify the requests that you want to allow, block, or count. In a RuleGroup, you also specify a default action (ALLOW or BLOCK), and the action for each Rule that you add to a RuleGroup, for example, block requests from specified IP addresses or block requests from specified referrers. You also associate the RuleGroup with a CloudFront distribution to identify the requests that you want AWS WAF to filter. If you add more than one Rule to a RuleGroup, a request needs to match only one of the specifications to be allowed, blocked, or counted.
func LookupRuleGroup(ctx *pulumi.Context, args *LookupRuleGroupArgs, opts ...pulumi.InvokeOption) (*LookupRuleGroupResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupRuleGroupResult
	err := ctx.Invoke("aws-native:wafv2:getRuleGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRuleGroupArgs struct {
	Id    string         `pulumi:"id"`
	Name  string         `pulumi:"name"`
	Scope RuleGroupScope `pulumi:"scope"`
}

type LookupRuleGroupResult struct {
	Arn *string `pulumi:"arn"`
	// Collection of Available Labels.
	AvailableLabels []RuleGroupLabelSummary `pulumi:"availableLabels"`
	Capacity        *int                    `pulumi:"capacity"`
	// Collection of Consumed Labels.
	ConsumedLabels       []RuleGroupLabelSummary                `pulumi:"consumedLabels"`
	CustomResponseBodies map[string]RuleGroupCustomResponseBody `pulumi:"customResponseBodies"`
	Description          *string                                `pulumi:"description"`
	Id                   *string                                `pulumi:"id"`
	LabelNamespace       *string                                `pulumi:"labelNamespace"`
	// Collection of Rules.
	Rules            []RuleGroupRule            `pulumi:"rules"`
	Tags             []aws.Tag                  `pulumi:"tags"`
	VisibilityConfig *RuleGroupVisibilityConfig `pulumi:"visibilityConfig"`
}

func LookupRuleGroupOutput(ctx *pulumi.Context, args LookupRuleGroupOutputArgs, opts ...pulumi.InvokeOption) LookupRuleGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRuleGroupResult, error) {
			args := v.(LookupRuleGroupArgs)
			r, err := LookupRuleGroup(ctx, &args, opts...)
			var s LookupRuleGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRuleGroupResultOutput)
}

type LookupRuleGroupOutputArgs struct {
	Id    pulumi.StringInput  `pulumi:"id"`
	Name  pulumi.StringInput  `pulumi:"name"`
	Scope RuleGroupScopeInput `pulumi:"scope"`
}

func (LookupRuleGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRuleGroupArgs)(nil)).Elem()
}

type LookupRuleGroupResultOutput struct{ *pulumi.OutputState }

func (LookupRuleGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRuleGroupResult)(nil)).Elem()
}

func (o LookupRuleGroupResultOutput) ToLookupRuleGroupResultOutput() LookupRuleGroupResultOutput {
	return o
}

func (o LookupRuleGroupResultOutput) ToLookupRuleGroupResultOutputWithContext(ctx context.Context) LookupRuleGroupResultOutput {
	return o
}

func (o LookupRuleGroupResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRuleGroupResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// Collection of Available Labels.
func (o LookupRuleGroupResultOutput) AvailableLabels() RuleGroupLabelSummaryArrayOutput {
	return o.ApplyT(func(v LookupRuleGroupResult) []RuleGroupLabelSummary { return v.AvailableLabels }).(RuleGroupLabelSummaryArrayOutput)
}

func (o LookupRuleGroupResultOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRuleGroupResult) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

// Collection of Consumed Labels.
func (o LookupRuleGroupResultOutput) ConsumedLabels() RuleGroupLabelSummaryArrayOutput {
	return o.ApplyT(func(v LookupRuleGroupResult) []RuleGroupLabelSummary { return v.ConsumedLabels }).(RuleGroupLabelSummaryArrayOutput)
}

func (o LookupRuleGroupResultOutput) CustomResponseBodies() RuleGroupCustomResponseBodyMapOutput {
	return o.ApplyT(func(v LookupRuleGroupResult) map[string]RuleGroupCustomResponseBody { return v.CustomResponseBodies }).(RuleGroupCustomResponseBodyMapOutput)
}

func (o LookupRuleGroupResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRuleGroupResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupRuleGroupResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRuleGroupResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupRuleGroupResultOutput) LabelNamespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRuleGroupResult) *string { return v.LabelNamespace }).(pulumi.StringPtrOutput)
}

// Collection of Rules.
func (o LookupRuleGroupResultOutput) Rules() RuleGroupRuleArrayOutput {
	return o.ApplyT(func(v LookupRuleGroupResult) []RuleGroupRule { return v.Rules }).(RuleGroupRuleArrayOutput)
}

func (o LookupRuleGroupResultOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupRuleGroupResult) []aws.Tag { return v.Tags }).(aws.TagArrayOutput)
}

func (o LookupRuleGroupResultOutput) VisibilityConfig() RuleGroupVisibilityConfigPtrOutput {
	return o.ApplyT(func(v LookupRuleGroupResult) *RuleGroupVisibilityConfig { return v.VisibilityConfig }).(RuleGroupVisibilityConfigPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRuleGroupResultOutput{})
}

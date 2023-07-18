// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package xray

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This schema provides construct and validation rules for AWS-XRay SamplingRule resource parameters.
func LookupSamplingRule(ctx *pulumi.Context, args *LookupSamplingRuleArgs, opts ...pulumi.InvokeOption) (*LookupSamplingRuleResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSamplingRuleResult
	err := ctx.Invoke("aws-native:xray:getSamplingRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSamplingRuleArgs struct {
	RuleARN string `pulumi:"ruleARN"`
}

type LookupSamplingRuleResult struct {
	RuleARN            *string             `pulumi:"ruleARN"`
	RuleName           *string             `pulumi:"ruleName"`
	SamplingRule       *SamplingRuleType   `pulumi:"samplingRule"`
	SamplingRuleRecord *SamplingRuleRecord `pulumi:"samplingRuleRecord"`
	SamplingRuleUpdate *SamplingRuleUpdate `pulumi:"samplingRuleUpdate"`
	Tags               []SamplingRuleTag   `pulumi:"tags"`
}

func LookupSamplingRuleOutput(ctx *pulumi.Context, args LookupSamplingRuleOutputArgs, opts ...pulumi.InvokeOption) LookupSamplingRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSamplingRuleResult, error) {
			args := v.(LookupSamplingRuleArgs)
			r, err := LookupSamplingRule(ctx, &args, opts...)
			var s LookupSamplingRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSamplingRuleResultOutput)
}

type LookupSamplingRuleOutputArgs struct {
	RuleARN pulumi.StringInput `pulumi:"ruleARN"`
}

func (LookupSamplingRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSamplingRuleArgs)(nil)).Elem()
}

type LookupSamplingRuleResultOutput struct{ *pulumi.OutputState }

func (LookupSamplingRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSamplingRuleResult)(nil)).Elem()
}

func (o LookupSamplingRuleResultOutput) ToLookupSamplingRuleResultOutput() LookupSamplingRuleResultOutput {
	return o
}

func (o LookupSamplingRuleResultOutput) ToLookupSamplingRuleResultOutputWithContext(ctx context.Context) LookupSamplingRuleResultOutput {
	return o
}

func (o LookupSamplingRuleResultOutput) RuleARN() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSamplingRuleResult) *string { return v.RuleARN }).(pulumi.StringPtrOutput)
}

func (o LookupSamplingRuleResultOutput) RuleName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSamplingRuleResult) *string { return v.RuleName }).(pulumi.StringPtrOutput)
}

func (o LookupSamplingRuleResultOutput) SamplingRule() SamplingRuleTypePtrOutput {
	return o.ApplyT(func(v LookupSamplingRuleResult) *SamplingRuleType { return v.SamplingRule }).(SamplingRuleTypePtrOutput)
}

func (o LookupSamplingRuleResultOutput) SamplingRuleRecord() SamplingRuleRecordPtrOutput {
	return o.ApplyT(func(v LookupSamplingRuleResult) *SamplingRuleRecord { return v.SamplingRuleRecord }).(SamplingRuleRecordPtrOutput)
}

func (o LookupSamplingRuleResultOutput) SamplingRuleUpdate() SamplingRuleUpdatePtrOutput {
	return o.ApplyT(func(v LookupSamplingRuleResult) *SamplingRuleUpdate { return v.SamplingRuleUpdate }).(SamplingRuleUpdatePtrOutput)
}

func (o LookupSamplingRuleResultOutput) Tags() SamplingRuleTagArrayOutput {
	return o.ApplyT(func(v LookupSamplingRuleResult) []SamplingRuleTag { return v.Tags }).(SamplingRuleTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSamplingRuleResultOutput{})
}

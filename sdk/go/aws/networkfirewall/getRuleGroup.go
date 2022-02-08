// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package networkfirewall

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource type definition for AWS::NetworkFirewall::RuleGroup
func LookupRuleGroup(ctx *pulumi.Context, args *LookupRuleGroupArgs, opts ...pulumi.InvokeOption) (*LookupRuleGroupResult, error) {
	var rv LookupRuleGroupResult
	err := ctx.Invoke("aws-native:networkfirewall:getRuleGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRuleGroupArgs struct {
	RuleGroupArn string `pulumi:"ruleGroupArn"`
}

type LookupRuleGroupResult struct {
	Description  *string        `pulumi:"description"`
	RuleGroup    *RuleGroupType `pulumi:"ruleGroup"`
	RuleGroupArn *string        `pulumi:"ruleGroupArn"`
	RuleGroupId  *string        `pulumi:"ruleGroupId"`
	Tags         []RuleGroupTag `pulumi:"tags"`
}

func LookupRuleGroupOutput(ctx *pulumi.Context, args LookupRuleGroupOutputArgs, opts ...pulumi.InvokeOption) LookupRuleGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRuleGroupResult, error) {
			args := v.(LookupRuleGroupArgs)
			r, err := LookupRuleGroup(ctx, &args, opts...)
			return *r, err
		}).(LookupRuleGroupResultOutput)
}

type LookupRuleGroupOutputArgs struct {
	RuleGroupArn pulumi.StringInput `pulumi:"ruleGroupArn"`
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

func (o LookupRuleGroupResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRuleGroupResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupRuleGroupResultOutput) RuleGroup() RuleGroupTypePtrOutput {
	return o.ApplyT(func(v LookupRuleGroupResult) *RuleGroupType { return v.RuleGroup }).(RuleGroupTypePtrOutput)
}

func (o LookupRuleGroupResultOutput) RuleGroupArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRuleGroupResult) *string { return v.RuleGroupArn }).(pulumi.StringPtrOutput)
}

func (o LookupRuleGroupResultOutput) RuleGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRuleGroupResult) *string { return v.RuleGroupId }).(pulumi.StringPtrOutput)
}

func (o LookupRuleGroupResultOutput) Tags() RuleGroupTagArrayOutput {
	return o.ApplyT(func(v LookupRuleGroupResult) []RuleGroupTag { return v.Tags }).(RuleGroupTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRuleGroupResultOutput{})
}

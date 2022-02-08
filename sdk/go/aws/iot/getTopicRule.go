// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iot

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::IoT::TopicRule
func LookupTopicRule(ctx *pulumi.Context, args *LookupTopicRuleArgs, opts ...pulumi.InvokeOption) (*LookupTopicRuleResult, error) {
	var rv LookupTopicRuleResult
	err := ctx.Invoke("aws-native:iot:getTopicRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTopicRuleArgs struct {
	RuleName string `pulumi:"ruleName"`
}

type LookupTopicRuleResult struct {
	Arn              *string           `pulumi:"arn"`
	Tags             []TopicRuleTag    `pulumi:"tags"`
	TopicRulePayload *TopicRulePayload `pulumi:"topicRulePayload"`
}

func LookupTopicRuleOutput(ctx *pulumi.Context, args LookupTopicRuleOutputArgs, opts ...pulumi.InvokeOption) LookupTopicRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTopicRuleResult, error) {
			args := v.(LookupTopicRuleArgs)
			r, err := LookupTopicRule(ctx, &args, opts...)
			return *r, err
		}).(LookupTopicRuleResultOutput)
}

type LookupTopicRuleOutputArgs struct {
	RuleName pulumi.StringInput `pulumi:"ruleName"`
}

func (LookupTopicRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTopicRuleArgs)(nil)).Elem()
}

type LookupTopicRuleResultOutput struct{ *pulumi.OutputState }

func (LookupTopicRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTopicRuleResult)(nil)).Elem()
}

func (o LookupTopicRuleResultOutput) ToLookupTopicRuleResultOutput() LookupTopicRuleResultOutput {
	return o
}

func (o LookupTopicRuleResultOutput) ToLookupTopicRuleResultOutputWithContext(ctx context.Context) LookupTopicRuleResultOutput {
	return o
}

func (o LookupTopicRuleResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTopicRuleResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupTopicRuleResultOutput) Tags() TopicRuleTagArrayOutput {
	return o.ApplyT(func(v LookupTopicRuleResult) []TopicRuleTag { return v.Tags }).(TopicRuleTagArrayOutput)
}

func (o LookupTopicRuleResultOutput) TopicRulePayload() TopicRulePayloadPtrOutput {
	return o.ApplyT(func(v LookupTopicRuleResult) *TopicRulePayload { return v.TopicRulePayload }).(TopicRulePayloadPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTopicRuleResultOutput{})
}

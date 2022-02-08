// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ssmincidents

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource type definition for AWS::SSMIncidents::ResponsePlan
func LookupResponsePlan(ctx *pulumi.Context, args *LookupResponsePlanArgs, opts ...pulumi.InvokeOption) (*LookupResponsePlanResult, error) {
	var rv LookupResponsePlanResult
	err := ctx.Invoke("aws-native:ssmincidents:getResponsePlan", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupResponsePlanArgs struct {
	// The ARN of the response plan.
	Arn string `pulumi:"arn"`
}

type LookupResponsePlanResult struct {
	// The list of actions.
	Actions []ResponsePlanAction `pulumi:"actions"`
	// The ARN of the response plan.
	Arn         *string                  `pulumi:"arn"`
	ChatChannel *ResponsePlanChatChannel `pulumi:"chatChannel"`
	// The display name of the response plan.
	DisplayName *string `pulumi:"displayName"`
	// The list of engagements to use.
	Engagements      []string                      `pulumi:"engagements"`
	IncidentTemplate *ResponsePlanIncidentTemplate `pulumi:"incidentTemplate"`
	// The tags to apply to the response plan.
	Tags []ResponsePlanTag `pulumi:"tags"`
}

func LookupResponsePlanOutput(ctx *pulumi.Context, args LookupResponsePlanOutputArgs, opts ...pulumi.InvokeOption) LookupResponsePlanResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupResponsePlanResult, error) {
			args := v.(LookupResponsePlanArgs)
			r, err := LookupResponsePlan(ctx, &args, opts...)
			return *r, err
		}).(LookupResponsePlanResultOutput)
}

type LookupResponsePlanOutputArgs struct {
	// The ARN of the response plan.
	Arn pulumi.StringInput `pulumi:"arn"`
}

func (LookupResponsePlanOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResponsePlanArgs)(nil)).Elem()
}

type LookupResponsePlanResultOutput struct{ *pulumi.OutputState }

func (LookupResponsePlanResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResponsePlanResult)(nil)).Elem()
}

func (o LookupResponsePlanResultOutput) ToLookupResponsePlanResultOutput() LookupResponsePlanResultOutput {
	return o
}

func (o LookupResponsePlanResultOutput) ToLookupResponsePlanResultOutputWithContext(ctx context.Context) LookupResponsePlanResultOutput {
	return o
}

// The list of actions.
func (o LookupResponsePlanResultOutput) Actions() ResponsePlanActionArrayOutput {
	return o.ApplyT(func(v LookupResponsePlanResult) []ResponsePlanAction { return v.Actions }).(ResponsePlanActionArrayOutput)
}

// The ARN of the response plan.
func (o LookupResponsePlanResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupResponsePlanResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupResponsePlanResultOutput) ChatChannel() ResponsePlanChatChannelPtrOutput {
	return o.ApplyT(func(v LookupResponsePlanResult) *ResponsePlanChatChannel { return v.ChatChannel }).(ResponsePlanChatChannelPtrOutput)
}

// The display name of the response plan.
func (o LookupResponsePlanResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupResponsePlanResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// The list of engagements to use.
func (o LookupResponsePlanResultOutput) Engagements() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupResponsePlanResult) []string { return v.Engagements }).(pulumi.StringArrayOutput)
}

func (o LookupResponsePlanResultOutput) IncidentTemplate() ResponsePlanIncidentTemplatePtrOutput {
	return o.ApplyT(func(v LookupResponsePlanResult) *ResponsePlanIncidentTemplate { return v.IncidentTemplate }).(ResponsePlanIncidentTemplatePtrOutput)
}

// The tags to apply to the response plan.
func (o LookupResponsePlanResultOutput) Tags() ResponsePlanTagArrayOutput {
	return o.ApplyT(func(v LookupResponsePlanResult) []ResponsePlanTag { return v.Tags }).(ResponsePlanTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupResponsePlanResultOutput{})
}

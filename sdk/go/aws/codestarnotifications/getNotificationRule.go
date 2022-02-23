// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package codestarnotifications

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::CodeStarNotifications::NotificationRule
func LookupNotificationRule(ctx *pulumi.Context, args *LookupNotificationRuleArgs, opts ...pulumi.InvokeOption) (*LookupNotificationRuleResult, error) {
	var rv LookupNotificationRuleResult
	err := ctx.Invoke("aws-native:codestarnotifications:getNotificationRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNotificationRuleArgs struct {
	Arn string `pulumi:"arn"`
}

type LookupNotificationRuleResult struct {
	Arn           *string                     `pulumi:"arn"`
	CreatedBy     *string                     `pulumi:"createdBy"`
	DetailType    *NotificationRuleDetailType `pulumi:"detailType"`
	EventTypeId   *string                     `pulumi:"eventTypeId"`
	EventTypeIds  []string                    `pulumi:"eventTypeIds"`
	Name          *string                     `pulumi:"name"`
	Status        *NotificationRuleStatus     `pulumi:"status"`
	TargetAddress *string                     `pulumi:"targetAddress"`
	Targets       []NotificationRuleTarget    `pulumi:"targets"`
}

func LookupNotificationRuleOutput(ctx *pulumi.Context, args LookupNotificationRuleOutputArgs, opts ...pulumi.InvokeOption) LookupNotificationRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNotificationRuleResult, error) {
			args := v.(LookupNotificationRuleArgs)
			r, err := LookupNotificationRule(ctx, &args, opts...)
			return *r, err
		}).(LookupNotificationRuleResultOutput)
}

type LookupNotificationRuleOutputArgs struct {
	Arn pulumi.StringInput `pulumi:"arn"`
}

func (LookupNotificationRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNotificationRuleArgs)(nil)).Elem()
}

type LookupNotificationRuleResultOutput struct{ *pulumi.OutputState }

func (LookupNotificationRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNotificationRuleResult)(nil)).Elem()
}

func (o LookupNotificationRuleResultOutput) ToLookupNotificationRuleResultOutput() LookupNotificationRuleResultOutput {
	return o
}

func (o LookupNotificationRuleResultOutput) ToLookupNotificationRuleResultOutputWithContext(ctx context.Context) LookupNotificationRuleResultOutput {
	return o
}

func (o LookupNotificationRuleResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNotificationRuleResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupNotificationRuleResultOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNotificationRuleResult) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

func (o LookupNotificationRuleResultOutput) DetailType() NotificationRuleDetailTypePtrOutput {
	return o.ApplyT(func(v LookupNotificationRuleResult) *NotificationRuleDetailType { return v.DetailType }).(NotificationRuleDetailTypePtrOutput)
}

func (o LookupNotificationRuleResultOutput) EventTypeId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNotificationRuleResult) *string { return v.EventTypeId }).(pulumi.StringPtrOutput)
}

func (o LookupNotificationRuleResultOutput) EventTypeIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNotificationRuleResult) []string { return v.EventTypeIds }).(pulumi.StringArrayOutput)
}

func (o LookupNotificationRuleResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNotificationRuleResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupNotificationRuleResultOutput) Status() NotificationRuleStatusPtrOutput {
	return o.ApplyT(func(v LookupNotificationRuleResult) *NotificationRuleStatus { return v.Status }).(NotificationRuleStatusPtrOutput)
}

func (o LookupNotificationRuleResultOutput) TargetAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNotificationRuleResult) *string { return v.TargetAddress }).(pulumi.StringPtrOutput)
}

func (o LookupNotificationRuleResultOutput) Targets() NotificationRuleTargetArrayOutput {
	return o.ApplyT(func(v LookupNotificationRuleResult) []NotificationRuleTarget { return v.Targets }).(NotificationRuleTargetArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNotificationRuleResultOutput{})
}
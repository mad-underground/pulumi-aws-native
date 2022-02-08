// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::ApiGateway::UsagePlan
func LookupUsagePlan(ctx *pulumi.Context, args *LookupUsagePlanArgs, opts ...pulumi.InvokeOption) (*LookupUsagePlanResult, error) {
	var rv LookupUsagePlanResult
	err := ctx.Invoke("aws-native:apigateway:getUsagePlan", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupUsagePlanArgs struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}

type LookupUsagePlanResult struct {
	// The API stages to associate with this usage plan.
	ApiStages []UsagePlanApiStage `pulumi:"apiStages"`
	// A description of the usage plan.
	Description *string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id *string `pulumi:"id"`
	// Configures the number of requests that users can make within a given interval.
	Quota *UsagePlanQuotaSettings `pulumi:"quota"`
	// An array of arbitrary tags (key-value pairs) to associate with the usage plan.
	Tags []UsagePlanTag `pulumi:"tags"`
	// Configures the overall request rate (average requests per second) and burst capacity.
	Throttle *UsagePlanThrottleSettings `pulumi:"throttle"`
	// A name for the usage plan.
	UsagePlanName *string `pulumi:"usagePlanName"`
}

func LookupUsagePlanOutput(ctx *pulumi.Context, args LookupUsagePlanOutputArgs, opts ...pulumi.InvokeOption) LookupUsagePlanResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupUsagePlanResult, error) {
			args := v.(LookupUsagePlanArgs)
			r, err := LookupUsagePlan(ctx, &args, opts...)
			return *r, err
		}).(LookupUsagePlanResultOutput)
}

type LookupUsagePlanOutputArgs struct {
	// The provider-assigned unique ID for this managed resource.
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupUsagePlanOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUsagePlanArgs)(nil)).Elem()
}

type LookupUsagePlanResultOutput struct{ *pulumi.OutputState }

func (LookupUsagePlanResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUsagePlanResult)(nil)).Elem()
}

func (o LookupUsagePlanResultOutput) ToLookupUsagePlanResultOutput() LookupUsagePlanResultOutput {
	return o
}

func (o LookupUsagePlanResultOutput) ToLookupUsagePlanResultOutputWithContext(ctx context.Context) LookupUsagePlanResultOutput {
	return o
}

// The API stages to associate with this usage plan.
func (o LookupUsagePlanResultOutput) ApiStages() UsagePlanApiStageArrayOutput {
	return o.ApplyT(func(v LookupUsagePlanResult) []UsagePlanApiStage { return v.ApiStages }).(UsagePlanApiStageArrayOutput)
}

// A description of the usage plan.
func (o LookupUsagePlanResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUsagePlanResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupUsagePlanResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUsagePlanResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// Configures the number of requests that users can make within a given interval.
func (o LookupUsagePlanResultOutput) Quota() UsagePlanQuotaSettingsPtrOutput {
	return o.ApplyT(func(v LookupUsagePlanResult) *UsagePlanQuotaSettings { return v.Quota }).(UsagePlanQuotaSettingsPtrOutput)
}

// An array of arbitrary tags (key-value pairs) to associate with the usage plan.
func (o LookupUsagePlanResultOutput) Tags() UsagePlanTagArrayOutput {
	return o.ApplyT(func(v LookupUsagePlanResult) []UsagePlanTag { return v.Tags }).(UsagePlanTagArrayOutput)
}

// Configures the overall request rate (average requests per second) and burst capacity.
func (o LookupUsagePlanResultOutput) Throttle() UsagePlanThrottleSettingsPtrOutput {
	return o.ApplyT(func(v LookupUsagePlanResult) *UsagePlanThrottleSettings { return v.Throttle }).(UsagePlanThrottleSettingsPtrOutput)
}

// A name for the usage plan.
func (o LookupUsagePlanResultOutput) UsagePlanName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUsagePlanResult) *string { return v.UsagePlanName }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupUsagePlanResultOutput{})
}

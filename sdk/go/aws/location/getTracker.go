// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package location

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of AWS::Location::Tracker Resource Type
func LookupTracker(ctx *pulumi.Context, args *LookupTrackerArgs, opts ...pulumi.InvokeOption) (*LookupTrackerResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupTrackerResult
	err := ctx.Invoke("aws-native:location:getTracker", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTrackerArgs struct {
	TrackerName string `pulumi:"trackerName"`
}

type LookupTrackerResult struct {
	Arn                           *string                   `pulumi:"arn"`
	CreateTime                    *string                   `pulumi:"createTime"`
	Description                   *string                   `pulumi:"description"`
	EventBridgeEnabled            *bool                     `pulumi:"eventBridgeEnabled"`
	KmsKeyEnableGeospatialQueries *bool                     `pulumi:"kmsKeyEnableGeospatialQueries"`
	PositionFiltering             *TrackerPositionFiltering `pulumi:"positionFiltering"`
	PricingPlan                   *TrackerPricingPlan       `pulumi:"pricingPlan"`
	PricingPlanDataSource         *string                   `pulumi:"pricingPlanDataSource"`
	// An array of key-value pairs to apply to this resource.
	Tags       []TrackerTag `pulumi:"tags"`
	TrackerArn *string      `pulumi:"trackerArn"`
	UpdateTime *string      `pulumi:"updateTime"`
}

func LookupTrackerOutput(ctx *pulumi.Context, args LookupTrackerOutputArgs, opts ...pulumi.InvokeOption) LookupTrackerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTrackerResult, error) {
			args := v.(LookupTrackerArgs)
			r, err := LookupTracker(ctx, &args, opts...)
			var s LookupTrackerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTrackerResultOutput)
}

type LookupTrackerOutputArgs struct {
	TrackerName pulumi.StringInput `pulumi:"trackerName"`
}

func (LookupTrackerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTrackerArgs)(nil)).Elem()
}

type LookupTrackerResultOutput struct{ *pulumi.OutputState }

func (LookupTrackerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTrackerResult)(nil)).Elem()
}

func (o LookupTrackerResultOutput) ToLookupTrackerResultOutput() LookupTrackerResultOutput {
	return o
}

func (o LookupTrackerResultOutput) ToLookupTrackerResultOutputWithContext(ctx context.Context) LookupTrackerResultOutput {
	return o
}

func (o LookupTrackerResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTrackerResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupTrackerResultOutput) CreateTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTrackerResult) *string { return v.CreateTime }).(pulumi.StringPtrOutput)
}

func (o LookupTrackerResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTrackerResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupTrackerResultOutput) EventBridgeEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupTrackerResult) *bool { return v.EventBridgeEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupTrackerResultOutput) KmsKeyEnableGeospatialQueries() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupTrackerResult) *bool { return v.KmsKeyEnableGeospatialQueries }).(pulumi.BoolPtrOutput)
}

func (o LookupTrackerResultOutput) PositionFiltering() TrackerPositionFilteringPtrOutput {
	return o.ApplyT(func(v LookupTrackerResult) *TrackerPositionFiltering { return v.PositionFiltering }).(TrackerPositionFilteringPtrOutput)
}

func (o LookupTrackerResultOutput) PricingPlan() TrackerPricingPlanPtrOutput {
	return o.ApplyT(func(v LookupTrackerResult) *TrackerPricingPlan { return v.PricingPlan }).(TrackerPricingPlanPtrOutput)
}

func (o LookupTrackerResultOutput) PricingPlanDataSource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTrackerResult) *string { return v.PricingPlanDataSource }).(pulumi.StringPtrOutput)
}

// An array of key-value pairs to apply to this resource.
func (o LookupTrackerResultOutput) Tags() TrackerTagArrayOutput {
	return o.ApplyT(func(v LookupTrackerResult) []TrackerTag { return v.Tags }).(TrackerTagArrayOutput)
}

func (o LookupTrackerResultOutput) TrackerArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTrackerResult) *string { return v.TrackerArn }).(pulumi.StringPtrOutput)
}

func (o LookupTrackerResultOutput) UpdateTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTrackerResult) *string { return v.UpdateTime }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTrackerResultOutput{})
}

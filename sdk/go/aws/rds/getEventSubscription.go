// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::RDS::EventSubscription
func LookupEventSubscription(ctx *pulumi.Context, args *LookupEventSubscriptionArgs, opts ...pulumi.InvokeOption) (*LookupEventSubscriptionResult, error) {
	var rv LookupEventSubscriptionResult
	err := ctx.Invoke("aws-native:rds:getEventSubscription", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEventSubscriptionArgs struct {
	Id string `pulumi:"id"`
}

type LookupEventSubscriptionResult struct {
	Enabled         *bool    `pulumi:"enabled"`
	EventCategories []string `pulumi:"eventCategories"`
	Id              *string  `pulumi:"id"`
	SourceIds       []string `pulumi:"sourceIds"`
	SourceType      *string  `pulumi:"sourceType"`
}

func LookupEventSubscriptionOutput(ctx *pulumi.Context, args LookupEventSubscriptionOutputArgs, opts ...pulumi.InvokeOption) LookupEventSubscriptionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEventSubscriptionResult, error) {
			args := v.(LookupEventSubscriptionArgs)
			r, err := LookupEventSubscription(ctx, &args, opts...)
			return *r, err
		}).(LookupEventSubscriptionResultOutput)
}

type LookupEventSubscriptionOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupEventSubscriptionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEventSubscriptionArgs)(nil)).Elem()
}

type LookupEventSubscriptionResultOutput struct{ *pulumi.OutputState }

func (LookupEventSubscriptionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEventSubscriptionResult)(nil)).Elem()
}

func (o LookupEventSubscriptionResultOutput) ToLookupEventSubscriptionResultOutput() LookupEventSubscriptionResultOutput {
	return o
}

func (o LookupEventSubscriptionResultOutput) ToLookupEventSubscriptionResultOutputWithContext(ctx context.Context) LookupEventSubscriptionResultOutput {
	return o
}

func (o LookupEventSubscriptionResultOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupEventSubscriptionResult) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o LookupEventSubscriptionResultOutput) EventCategories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupEventSubscriptionResult) []string { return v.EventCategories }).(pulumi.StringArrayOutput)
}

func (o LookupEventSubscriptionResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEventSubscriptionResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupEventSubscriptionResultOutput) SourceIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupEventSubscriptionResult) []string { return v.SourceIds }).(pulumi.StringArrayOutput)
}

func (o LookupEventSubscriptionResultOutput) SourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEventSubscriptionResult) *string { return v.SourceType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEventSubscriptionResultOutput{})
}

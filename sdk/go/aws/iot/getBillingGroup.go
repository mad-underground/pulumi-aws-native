// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iot

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::IoT::BillingGroup
func LookupBillingGroup(ctx *pulumi.Context, args *LookupBillingGroupArgs, opts ...pulumi.InvokeOption) (*LookupBillingGroupResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupBillingGroupResult
	err := ctx.Invoke("aws-native:iot:getBillingGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBillingGroupArgs struct {
	BillingGroupName string `pulumi:"billingGroupName"`
}

type LookupBillingGroupResult struct {
	Arn                    *string                           `pulumi:"arn"`
	BillingGroupProperties *BillingGroupPropertiesProperties `pulumi:"billingGroupProperties"`
	Id                     *string                           `pulumi:"id"`
	// An array of key-value pairs to apply to this resource.
	Tags []BillingGroupTag `pulumi:"tags"`
}

func LookupBillingGroupOutput(ctx *pulumi.Context, args LookupBillingGroupOutputArgs, opts ...pulumi.InvokeOption) LookupBillingGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBillingGroupResult, error) {
			args := v.(LookupBillingGroupArgs)
			r, err := LookupBillingGroup(ctx, &args, opts...)
			var s LookupBillingGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBillingGroupResultOutput)
}

type LookupBillingGroupOutputArgs struct {
	BillingGroupName pulumi.StringInput `pulumi:"billingGroupName"`
}

func (LookupBillingGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBillingGroupArgs)(nil)).Elem()
}

type LookupBillingGroupResultOutput struct{ *pulumi.OutputState }

func (LookupBillingGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBillingGroupResult)(nil)).Elem()
}

func (o LookupBillingGroupResultOutput) ToLookupBillingGroupResultOutput() LookupBillingGroupResultOutput {
	return o
}

func (o LookupBillingGroupResultOutput) ToLookupBillingGroupResultOutputWithContext(ctx context.Context) LookupBillingGroupResultOutput {
	return o
}

func (o LookupBillingGroupResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBillingGroupResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupBillingGroupResultOutput) BillingGroupProperties() BillingGroupPropertiesPropertiesPtrOutput {
	return o.ApplyT(func(v LookupBillingGroupResult) *BillingGroupPropertiesProperties { return v.BillingGroupProperties }).(BillingGroupPropertiesPropertiesPtrOutput)
}

func (o LookupBillingGroupResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBillingGroupResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// An array of key-value pairs to apply to this resource.
func (o LookupBillingGroupResultOutput) Tags() BillingGroupTagArrayOutput {
	return o.ApplyT(func(v LookupBillingGroupResult) []BillingGroupTag { return v.Tags }).(BillingGroupTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBillingGroupResultOutput{})
}

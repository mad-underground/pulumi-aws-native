// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package oam

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Oam::Sink
func LookupSink(ctx *pulumi.Context, args *LookupSinkArgs, opts ...pulumi.InvokeOption) (*LookupSinkResult, error) {
	var rv LookupSinkResult
	err := ctx.Invoke("aws-native:oam:getSink", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSinkArgs struct {
	// The Amazon resource name (ARN) of the ObservabilityAccessManager Sink
	Arn string `pulumi:"arn"`
}

type LookupSinkResult struct {
	// The Amazon resource name (ARN) of the ObservabilityAccessManager Sink
	Arn *string `pulumi:"arn"`
	// The policy of this ObservabilityAccessManager Sink.
	Policy interface{} `pulumi:"policy"`
	// Tags to apply to the sink
	Tags interface{} `pulumi:"tags"`
}

func LookupSinkOutput(ctx *pulumi.Context, args LookupSinkOutputArgs, opts ...pulumi.InvokeOption) LookupSinkResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSinkResult, error) {
			args := v.(LookupSinkArgs)
			r, err := LookupSink(ctx, &args, opts...)
			var s LookupSinkResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSinkResultOutput)
}

type LookupSinkOutputArgs struct {
	// The Amazon resource name (ARN) of the ObservabilityAccessManager Sink
	Arn pulumi.StringInput `pulumi:"arn"`
}

func (LookupSinkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSinkArgs)(nil)).Elem()
}

type LookupSinkResultOutput struct{ *pulumi.OutputState }

func (LookupSinkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSinkResult)(nil)).Elem()
}

func (o LookupSinkResultOutput) ToLookupSinkResultOutput() LookupSinkResultOutput {
	return o
}

func (o LookupSinkResultOutput) ToLookupSinkResultOutputWithContext(ctx context.Context) LookupSinkResultOutput {
	return o
}

// The Amazon resource name (ARN) of the ObservabilityAccessManager Sink
func (o LookupSinkResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSinkResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// The policy of this ObservabilityAccessManager Sink.
func (o LookupSinkResultOutput) Policy() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupSinkResult) interface{} { return v.Policy }).(pulumi.AnyOutput)
}

// Tags to apply to the sink
func (o LookupSinkResultOutput) Tags() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupSinkResult) interface{} { return v.Tags }).(pulumi.AnyOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSinkResultOutput{})
}

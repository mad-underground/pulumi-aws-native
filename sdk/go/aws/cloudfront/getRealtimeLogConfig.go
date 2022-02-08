// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudfront

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::CloudFront::RealtimeLogConfig
func LookupRealtimeLogConfig(ctx *pulumi.Context, args *LookupRealtimeLogConfigArgs, opts ...pulumi.InvokeOption) (*LookupRealtimeLogConfigResult, error) {
	var rv LookupRealtimeLogConfigResult
	err := ctx.Invoke("aws-native:cloudfront:getRealtimeLogConfig", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRealtimeLogConfigArgs struct {
	Arn string `pulumi:"arn"`
}

type LookupRealtimeLogConfigResult struct {
	Arn          *string                     `pulumi:"arn"`
	EndPoints    []RealtimeLogConfigEndPoint `pulumi:"endPoints"`
	Fields       []string                    `pulumi:"fields"`
	SamplingRate *float64                    `pulumi:"samplingRate"`
}

func LookupRealtimeLogConfigOutput(ctx *pulumi.Context, args LookupRealtimeLogConfigOutputArgs, opts ...pulumi.InvokeOption) LookupRealtimeLogConfigResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRealtimeLogConfigResult, error) {
			args := v.(LookupRealtimeLogConfigArgs)
			r, err := LookupRealtimeLogConfig(ctx, &args, opts...)
			return *r, err
		}).(LookupRealtimeLogConfigResultOutput)
}

type LookupRealtimeLogConfigOutputArgs struct {
	Arn pulumi.StringInput `pulumi:"arn"`
}

func (LookupRealtimeLogConfigOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRealtimeLogConfigArgs)(nil)).Elem()
}

type LookupRealtimeLogConfigResultOutput struct{ *pulumi.OutputState }

func (LookupRealtimeLogConfigResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRealtimeLogConfigResult)(nil)).Elem()
}

func (o LookupRealtimeLogConfigResultOutput) ToLookupRealtimeLogConfigResultOutput() LookupRealtimeLogConfigResultOutput {
	return o
}

func (o LookupRealtimeLogConfigResultOutput) ToLookupRealtimeLogConfigResultOutputWithContext(ctx context.Context) LookupRealtimeLogConfigResultOutput {
	return o
}

func (o LookupRealtimeLogConfigResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRealtimeLogConfigResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupRealtimeLogConfigResultOutput) EndPoints() RealtimeLogConfigEndPointArrayOutput {
	return o.ApplyT(func(v LookupRealtimeLogConfigResult) []RealtimeLogConfigEndPoint { return v.EndPoints }).(RealtimeLogConfigEndPointArrayOutput)
}

func (o LookupRealtimeLogConfigResultOutput) Fields() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRealtimeLogConfigResult) []string { return v.Fields }).(pulumi.StringArrayOutput)
}

func (o LookupRealtimeLogConfigResultOutput) SamplingRate() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupRealtimeLogConfigResult) *float64 { return v.SamplingRate }).(pulumi.Float64PtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRealtimeLogConfigResultOutput{})
}

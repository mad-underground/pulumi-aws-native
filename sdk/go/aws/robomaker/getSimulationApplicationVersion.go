// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package robomaker

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An example resource schema demonstrating some basic constructs and validation rules.
func LookupSimulationApplicationVersion(ctx *pulumi.Context, args *LookupSimulationApplicationVersionArgs, opts ...pulumi.InvokeOption) (*LookupSimulationApplicationVersionResult, error) {
	var rv LookupSimulationApplicationVersionResult
	err := ctx.Invoke("aws-native:robomaker:getSimulationApplicationVersion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSimulationApplicationVersionArgs struct {
	Arn string `pulumi:"arn"`
}

type LookupSimulationApplicationVersionResult struct {
	ApplicationVersion *string `pulumi:"applicationVersion"`
	Arn                *string `pulumi:"arn"`
}

func LookupSimulationApplicationVersionOutput(ctx *pulumi.Context, args LookupSimulationApplicationVersionOutputArgs, opts ...pulumi.InvokeOption) LookupSimulationApplicationVersionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSimulationApplicationVersionResult, error) {
			args := v.(LookupSimulationApplicationVersionArgs)
			r, err := LookupSimulationApplicationVersion(ctx, &args, opts...)
			return *r, err
		}).(LookupSimulationApplicationVersionResultOutput)
}

type LookupSimulationApplicationVersionOutputArgs struct {
	Arn pulumi.StringInput `pulumi:"arn"`
}

func (LookupSimulationApplicationVersionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSimulationApplicationVersionArgs)(nil)).Elem()
}

type LookupSimulationApplicationVersionResultOutput struct{ *pulumi.OutputState }

func (LookupSimulationApplicationVersionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSimulationApplicationVersionResult)(nil)).Elem()
}

func (o LookupSimulationApplicationVersionResultOutput) ToLookupSimulationApplicationVersionResultOutput() LookupSimulationApplicationVersionResultOutput {
	return o
}

func (o LookupSimulationApplicationVersionResultOutput) ToLookupSimulationApplicationVersionResultOutputWithContext(ctx context.Context) LookupSimulationApplicationVersionResultOutput {
	return o
}

func (o LookupSimulationApplicationVersionResultOutput) ApplicationVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSimulationApplicationVersionResult) *string { return v.ApplicationVersion }).(pulumi.StringPtrOutput)
}

func (o LookupSimulationApplicationVersionResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSimulationApplicationVersionResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSimulationApplicationVersionResultOutput{})
}

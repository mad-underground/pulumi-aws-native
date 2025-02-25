// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package autoscaling

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::AutoScaling::WarmPool.
func LookupWarmPool(ctx *pulumi.Context, args *LookupWarmPoolArgs, opts ...pulumi.InvokeOption) (*LookupWarmPoolResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupWarmPoolResult
	err := ctx.Invoke("aws-native:autoscaling:getWarmPool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWarmPoolArgs struct {
	AutoScalingGroupName string `pulumi:"autoScalingGroupName"`
}

type LookupWarmPoolResult struct {
	InstanceReusePolicy      *WarmPoolInstanceReusePolicy `pulumi:"instanceReusePolicy"`
	MaxGroupPreparedCapacity *int                         `pulumi:"maxGroupPreparedCapacity"`
	MinSize                  *int                         `pulumi:"minSize"`
	PoolState                *string                      `pulumi:"poolState"`
}

func LookupWarmPoolOutput(ctx *pulumi.Context, args LookupWarmPoolOutputArgs, opts ...pulumi.InvokeOption) LookupWarmPoolResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWarmPoolResult, error) {
			args := v.(LookupWarmPoolArgs)
			r, err := LookupWarmPool(ctx, &args, opts...)
			var s LookupWarmPoolResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWarmPoolResultOutput)
}

type LookupWarmPoolOutputArgs struct {
	AutoScalingGroupName pulumi.StringInput `pulumi:"autoScalingGroupName"`
}

func (LookupWarmPoolOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWarmPoolArgs)(nil)).Elem()
}

type LookupWarmPoolResultOutput struct{ *pulumi.OutputState }

func (LookupWarmPoolResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWarmPoolResult)(nil)).Elem()
}

func (o LookupWarmPoolResultOutput) ToLookupWarmPoolResultOutput() LookupWarmPoolResultOutput {
	return o
}

func (o LookupWarmPoolResultOutput) ToLookupWarmPoolResultOutputWithContext(ctx context.Context) LookupWarmPoolResultOutput {
	return o
}

func (o LookupWarmPoolResultOutput) InstanceReusePolicy() WarmPoolInstanceReusePolicyPtrOutput {
	return o.ApplyT(func(v LookupWarmPoolResult) *WarmPoolInstanceReusePolicy { return v.InstanceReusePolicy }).(WarmPoolInstanceReusePolicyPtrOutput)
}

func (o LookupWarmPoolResultOutput) MaxGroupPreparedCapacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupWarmPoolResult) *int { return v.MaxGroupPreparedCapacity }).(pulumi.IntPtrOutput)
}

func (o LookupWarmPoolResultOutput) MinSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupWarmPoolResult) *int { return v.MinSize }).(pulumi.IntPtrOutput)
}

func (o LookupWarmPoolResultOutput) PoolState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWarmPoolResult) *string { return v.PoolState }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWarmPoolResultOutput{})
}

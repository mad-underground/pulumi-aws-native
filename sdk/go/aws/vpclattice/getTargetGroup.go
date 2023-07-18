// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpclattice

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A target group is a collection of targets, or compute resources, that run your application or service. A target group can only be used by a single service.
func LookupTargetGroup(ctx *pulumi.Context, args *LookupTargetGroupArgs, opts ...pulumi.InvokeOption) (*LookupTargetGroupResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupTargetGroupResult
	err := ctx.Invoke("aws-native:vpclattice:getTargetGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTargetGroupArgs struct {
	Arn string `pulumi:"arn"`
}

type LookupTargetGroupResult struct {
	Arn           *string             `pulumi:"arn"`
	Config        *TargetGroupConfig  `pulumi:"config"`
	CreatedAt     *string             `pulumi:"createdAt"`
	Id            *string             `pulumi:"id"`
	LastUpdatedAt *string             `pulumi:"lastUpdatedAt"`
	Status        *TargetGroupStatus  `pulumi:"status"`
	Tags          []TargetGroupTag    `pulumi:"tags"`
	Targets       []TargetGroupTarget `pulumi:"targets"`
}

func LookupTargetGroupOutput(ctx *pulumi.Context, args LookupTargetGroupOutputArgs, opts ...pulumi.InvokeOption) LookupTargetGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTargetGroupResult, error) {
			args := v.(LookupTargetGroupArgs)
			r, err := LookupTargetGroup(ctx, &args, opts...)
			var s LookupTargetGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTargetGroupResultOutput)
}

type LookupTargetGroupOutputArgs struct {
	Arn pulumi.StringInput `pulumi:"arn"`
}

func (LookupTargetGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTargetGroupArgs)(nil)).Elem()
}

type LookupTargetGroupResultOutput struct{ *pulumi.OutputState }

func (LookupTargetGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTargetGroupResult)(nil)).Elem()
}

func (o LookupTargetGroupResultOutput) ToLookupTargetGroupResultOutput() LookupTargetGroupResultOutput {
	return o
}

func (o LookupTargetGroupResultOutput) ToLookupTargetGroupResultOutputWithContext(ctx context.Context) LookupTargetGroupResultOutput {
	return o
}

func (o LookupTargetGroupResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTargetGroupResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupTargetGroupResultOutput) Config() TargetGroupConfigPtrOutput {
	return o.ApplyT(func(v LookupTargetGroupResult) *TargetGroupConfig { return v.Config }).(TargetGroupConfigPtrOutput)
}

func (o LookupTargetGroupResultOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTargetGroupResult) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o LookupTargetGroupResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTargetGroupResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupTargetGroupResultOutput) LastUpdatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTargetGroupResult) *string { return v.LastUpdatedAt }).(pulumi.StringPtrOutput)
}

func (o LookupTargetGroupResultOutput) Status() TargetGroupStatusPtrOutput {
	return o.ApplyT(func(v LookupTargetGroupResult) *TargetGroupStatus { return v.Status }).(TargetGroupStatusPtrOutput)
}

func (o LookupTargetGroupResultOutput) Tags() TargetGroupTagArrayOutput {
	return o.ApplyT(func(v LookupTargetGroupResult) []TargetGroupTag { return v.Tags }).(TargetGroupTagArrayOutput)
}

func (o LookupTargetGroupResultOutput) Targets() TargetGroupTargetArrayOutput {
	return o.ApplyT(func(v LookupTargetGroupResult) []TargetGroupTarget { return v.Targets }).(TargetGroupTargetArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTargetGroupResultOutput{})
}

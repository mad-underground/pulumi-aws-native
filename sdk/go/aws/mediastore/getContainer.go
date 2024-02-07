// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mediastore

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::MediaStore::Container
func LookupContainer(ctx *pulumi.Context, args *LookupContainerArgs, opts ...pulumi.InvokeOption) (*LookupContainerResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupContainerResult
	err := ctx.Invoke("aws-native:mediastore:getContainer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupContainerArgs struct {
	Id string `pulumi:"id"`
}

type LookupContainerResult struct {
	AccessLoggingEnabled *bool                  `pulumi:"accessLoggingEnabled"`
	CorsPolicy           []ContainerCorsRule    `pulumi:"corsPolicy"`
	Endpoint             *string                `pulumi:"endpoint"`
	Id                   *string                `pulumi:"id"`
	LifecyclePolicy      *string                `pulumi:"lifecyclePolicy"`
	MetricPolicy         *ContainerMetricPolicy `pulumi:"metricPolicy"`
	Policy               *string                `pulumi:"policy"`
	Tags                 []ContainerTag         `pulumi:"tags"`
}

func LookupContainerOutput(ctx *pulumi.Context, args LookupContainerOutputArgs, opts ...pulumi.InvokeOption) LookupContainerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupContainerResult, error) {
			args := v.(LookupContainerArgs)
			r, err := LookupContainer(ctx, &args, opts...)
			var s LookupContainerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupContainerResultOutput)
}

type LookupContainerOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupContainerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContainerArgs)(nil)).Elem()
}

type LookupContainerResultOutput struct{ *pulumi.OutputState }

func (LookupContainerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContainerResult)(nil)).Elem()
}

func (o LookupContainerResultOutput) ToLookupContainerResultOutput() LookupContainerResultOutput {
	return o
}

func (o LookupContainerResultOutput) ToLookupContainerResultOutputWithContext(ctx context.Context) LookupContainerResultOutput {
	return o
}

func (o LookupContainerResultOutput) AccessLoggingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupContainerResult) *bool { return v.AccessLoggingEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupContainerResultOutput) CorsPolicy() ContainerCorsRuleArrayOutput {
	return o.ApplyT(func(v LookupContainerResult) []ContainerCorsRule { return v.CorsPolicy }).(ContainerCorsRuleArrayOutput)
}

func (o LookupContainerResultOutput) Endpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContainerResult) *string { return v.Endpoint }).(pulumi.StringPtrOutput)
}

func (o LookupContainerResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContainerResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupContainerResultOutput) LifecyclePolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContainerResult) *string { return v.LifecyclePolicy }).(pulumi.StringPtrOutput)
}

func (o LookupContainerResultOutput) MetricPolicy() ContainerMetricPolicyPtrOutput {
	return o.ApplyT(func(v LookupContainerResult) *ContainerMetricPolicy { return v.MetricPolicy }).(ContainerMetricPolicyPtrOutput)
}

func (o LookupContainerResultOutput) Policy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContainerResult) *string { return v.Policy }).(pulumi.StringPtrOutput)
}

func (o LookupContainerResultOutput) Tags() ContainerTagArrayOutput {
	return o.ApplyT(func(v LookupContainerResult) []ContainerTag { return v.Tags }).(ContainerTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupContainerResultOutput{})
}

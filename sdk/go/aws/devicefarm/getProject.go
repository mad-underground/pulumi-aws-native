// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package devicefarm

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// AWS::DeviceFarm::Project creates a new Device Farm Project
func LookupProject(ctx *pulumi.Context, args *LookupProjectArgs, opts ...pulumi.InvokeOption) (*LookupProjectResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupProjectResult
	err := ctx.Invoke("aws-native:devicefarm:getProject", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupProjectArgs struct {
	Arn string `pulumi:"arn"`
}

type LookupProjectResult struct {
	Arn                      *string           `pulumi:"arn"`
	DefaultJobTimeoutMinutes *int              `pulumi:"defaultJobTimeoutMinutes"`
	Name                     *string           `pulumi:"name"`
	Tags                     []ProjectTag      `pulumi:"tags"`
	VpcConfig                *ProjectVpcConfig `pulumi:"vpcConfig"`
}

func LookupProjectOutput(ctx *pulumi.Context, args LookupProjectOutputArgs, opts ...pulumi.InvokeOption) LookupProjectResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupProjectResult, error) {
			args := v.(LookupProjectArgs)
			r, err := LookupProject(ctx, &args, opts...)
			var s LookupProjectResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupProjectResultOutput)
}

type LookupProjectOutputArgs struct {
	Arn pulumi.StringInput `pulumi:"arn"`
}

func (LookupProjectOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectArgs)(nil)).Elem()
}

type LookupProjectResultOutput struct{ *pulumi.OutputState }

func (LookupProjectResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectResult)(nil)).Elem()
}

func (o LookupProjectResultOutput) ToLookupProjectResultOutput() LookupProjectResultOutput {
	return o
}

func (o LookupProjectResultOutput) ToLookupProjectResultOutputWithContext(ctx context.Context) LookupProjectResultOutput {
	return o
}

func (o LookupProjectResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProjectResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupProjectResultOutput) DefaultJobTimeoutMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupProjectResult) *int { return v.DefaultJobTimeoutMinutes }).(pulumi.IntPtrOutput)
}

func (o LookupProjectResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupProjectResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupProjectResultOutput) Tags() ProjectTagArrayOutput {
	return o.ApplyT(func(v LookupProjectResult) []ProjectTag { return v.Tags }).(ProjectTagArrayOutput)
}

func (o LookupProjectResultOutput) VpcConfig() ProjectVpcConfigPtrOutput {
	return o.ApplyT(func(v LookupProjectResult) *ProjectVpcConfig { return v.VpcConfig }).(ProjectVpcConfigPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProjectResultOutput{})
}

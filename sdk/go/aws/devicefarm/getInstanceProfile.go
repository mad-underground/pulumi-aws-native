// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package devicefarm

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// AWS::DeviceFarm::InstanceProfile creates a new Device Farm Instance Profile
func LookupInstanceProfile(ctx *pulumi.Context, args *LookupInstanceProfileArgs, opts ...pulumi.InvokeOption) (*LookupInstanceProfileResult, error) {
	var rv LookupInstanceProfileResult
	err := ctx.Invoke("aws-native:devicefarm:getInstanceProfile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupInstanceProfileArgs struct {
	Arn string `pulumi:"arn"`
}

type LookupInstanceProfileResult struct {
	Arn                           *string              `pulumi:"arn"`
	Description                   *string              `pulumi:"description"`
	ExcludeAppPackagesFromCleanup []string             `pulumi:"excludeAppPackagesFromCleanup"`
	Name                          *string              `pulumi:"name"`
	PackageCleanup                *bool                `pulumi:"packageCleanup"`
	RebootAfterUse                *bool                `pulumi:"rebootAfterUse"`
	Tags                          []InstanceProfileTag `pulumi:"tags"`
}

func LookupInstanceProfileOutput(ctx *pulumi.Context, args LookupInstanceProfileOutputArgs, opts ...pulumi.InvokeOption) LookupInstanceProfileResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupInstanceProfileResult, error) {
			args := v.(LookupInstanceProfileArgs)
			r, err := LookupInstanceProfile(ctx, &args, opts...)
			return *r, err
		}).(LookupInstanceProfileResultOutput)
}

type LookupInstanceProfileOutputArgs struct {
	Arn pulumi.StringInput `pulumi:"arn"`
}

func (LookupInstanceProfileOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInstanceProfileArgs)(nil)).Elem()
}

type LookupInstanceProfileResultOutput struct{ *pulumi.OutputState }

func (LookupInstanceProfileResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInstanceProfileResult)(nil)).Elem()
}

func (o LookupInstanceProfileResultOutput) ToLookupInstanceProfileResultOutput() LookupInstanceProfileResultOutput {
	return o
}

func (o LookupInstanceProfileResultOutput) ToLookupInstanceProfileResultOutputWithContext(ctx context.Context) LookupInstanceProfileResultOutput {
	return o
}

func (o LookupInstanceProfileResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInstanceProfileResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupInstanceProfileResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInstanceProfileResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupInstanceProfileResultOutput) ExcludeAppPackagesFromCleanup() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupInstanceProfileResult) []string { return v.ExcludeAppPackagesFromCleanup }).(pulumi.StringArrayOutput)
}

func (o LookupInstanceProfileResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupInstanceProfileResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupInstanceProfileResultOutput) PackageCleanup() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupInstanceProfileResult) *bool { return v.PackageCleanup }).(pulumi.BoolPtrOutput)
}

func (o LookupInstanceProfileResultOutput) RebootAfterUse() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupInstanceProfileResult) *bool { return v.RebootAfterUse }).(pulumi.BoolPtrOutput)
}

func (o LookupInstanceProfileResultOutput) Tags() InstanceProfileTagArrayOutput {
	return o.ApplyT(func(v LookupInstanceProfileResult) []InstanceProfileTag { return v.Tags }).(InstanceProfileTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupInstanceProfileResultOutput{})
}

// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package lakeformation

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::LakeFormation::Resource
func LookupResource(ctx *pulumi.Context, args *LookupResourceArgs, opts ...pulumi.InvokeOption) (*LookupResourceResult, error) {
	var rv LookupResourceResult
	err := ctx.Invoke("aws-native:lakeformation:getResource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupResourceArgs struct {
	Id string `pulumi:"id"`
}

type LookupResourceResult struct {
	Id                   *string `pulumi:"id"`
	ResourceArn          *string `pulumi:"resourceArn"`
	RoleArn              *string `pulumi:"roleArn"`
	UseServiceLinkedRole *bool   `pulumi:"useServiceLinkedRole"`
}

func LookupResourceOutput(ctx *pulumi.Context, args LookupResourceOutputArgs, opts ...pulumi.InvokeOption) LookupResourceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupResourceResult, error) {
			args := v.(LookupResourceArgs)
			r, err := LookupResource(ctx, &args, opts...)
			return *r, err
		}).(LookupResourceResultOutput)
}

type LookupResourceOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupResourceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResourceArgs)(nil)).Elem()
}

type LookupResourceResultOutput struct{ *pulumi.OutputState }

func (LookupResourceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResourceResult)(nil)).Elem()
}

func (o LookupResourceResultOutput) ToLookupResourceResultOutput() LookupResourceResultOutput {
	return o
}

func (o LookupResourceResultOutput) ToLookupResourceResultOutputWithContext(ctx context.Context) LookupResourceResultOutput {
	return o
}

func (o LookupResourceResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupResourceResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupResourceResultOutput) ResourceArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupResourceResult) *string { return v.ResourceArn }).(pulumi.StringPtrOutput)
}

func (o LookupResourceResultOutput) RoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupResourceResult) *string { return v.RoleArn }).(pulumi.StringPtrOutput)
}

func (o LookupResourceResultOutput) UseServiceLinkedRole() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupResourceResult) *bool { return v.UseServiceLinkedRole }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupResourceResultOutput{})
}

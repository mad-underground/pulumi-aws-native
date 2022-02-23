// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package inspector

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Inspector::ResourceGroup
func LookupResourceGroup(ctx *pulumi.Context, args *LookupResourceGroupArgs, opts ...pulumi.InvokeOption) (*LookupResourceGroupResult, error) {
	var rv LookupResourceGroupResult
	err := ctx.Invoke("aws-native:inspector:getResourceGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupResourceGroupArgs struct {
	Id string `pulumi:"id"`
}

type LookupResourceGroupResult struct {
	Arn *string `pulumi:"arn"`
	Id  *string `pulumi:"id"`
}

func LookupResourceGroupOutput(ctx *pulumi.Context, args LookupResourceGroupOutputArgs, opts ...pulumi.InvokeOption) LookupResourceGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupResourceGroupResult, error) {
			args := v.(LookupResourceGroupArgs)
			r, err := LookupResourceGroup(ctx, &args, opts...)
			return *r, err
		}).(LookupResourceGroupResultOutput)
}

type LookupResourceGroupOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupResourceGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResourceGroupArgs)(nil)).Elem()
}

type LookupResourceGroupResultOutput struct{ *pulumi.OutputState }

func (LookupResourceGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResourceGroupResult)(nil)).Elem()
}

func (o LookupResourceGroupResultOutput) ToLookupResourceGroupResultOutput() LookupResourceGroupResultOutput {
	return o
}

func (o LookupResourceGroupResultOutput) ToLookupResourceGroupResultOutputWithContext(ctx context.Context) LookupResourceGroupResultOutput {
	return o
}

func (o LookupResourceGroupResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupResourceGroupResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupResourceGroupResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupResourceGroupResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupResourceGroupResultOutput{})
}
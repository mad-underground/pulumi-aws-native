// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dax

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::DAX::ParameterGroup
func LookupParameterGroup(ctx *pulumi.Context, args *LookupParameterGroupArgs, opts ...pulumi.InvokeOption) (*LookupParameterGroupResult, error) {
	var rv LookupParameterGroupResult
	err := ctx.Invoke("aws-native:dax:getParameterGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupParameterGroupArgs struct {
	Id string `pulumi:"id"`
}

type LookupParameterGroupResult struct {
	Description         *string     `pulumi:"description"`
	Id                  *string     `pulumi:"id"`
	ParameterNameValues interface{} `pulumi:"parameterNameValues"`
}

func LookupParameterGroupOutput(ctx *pulumi.Context, args LookupParameterGroupOutputArgs, opts ...pulumi.InvokeOption) LookupParameterGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupParameterGroupResult, error) {
			args := v.(LookupParameterGroupArgs)
			r, err := LookupParameterGroup(ctx, &args, opts...)
			return *r, err
		}).(LookupParameterGroupResultOutput)
}

type LookupParameterGroupOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupParameterGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupParameterGroupArgs)(nil)).Elem()
}

type LookupParameterGroupResultOutput struct{ *pulumi.OutputState }

func (LookupParameterGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupParameterGroupResult)(nil)).Elem()
}

func (o LookupParameterGroupResultOutput) ToLookupParameterGroupResultOutput() LookupParameterGroupResultOutput {
	return o
}

func (o LookupParameterGroupResultOutput) ToLookupParameterGroupResultOutputWithContext(ctx context.Context) LookupParameterGroupResultOutput {
	return o
}

func (o LookupParameterGroupResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupParameterGroupResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupParameterGroupResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupParameterGroupResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupParameterGroupResultOutput) ParameterNameValues() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupParameterGroupResult) interface{} { return v.ParameterNameValues }).(pulumi.AnyOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupParameterGroupResultOutput{})
}
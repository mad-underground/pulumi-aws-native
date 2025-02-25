// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package neptune

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Neptune::DBParameterGroup
func LookupDbParameterGroup(ctx *pulumi.Context, args *LookupDbParameterGroupArgs, opts ...pulumi.InvokeOption) (*LookupDbParameterGroupResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDbParameterGroupResult
	err := ctx.Invoke("aws-native:neptune:getDbParameterGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDbParameterGroupArgs struct {
	Id string `pulumi:"id"`
}

type LookupDbParameterGroupResult struct {
	Id *string `pulumi:"id"`
	// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Neptune::DBParameterGroup` for more information about the expected schema for this property.
	Parameters interface{} `pulumi:"parameters"`
	Tags       []aws.Tag   `pulumi:"tags"`
}

func LookupDbParameterGroupOutput(ctx *pulumi.Context, args LookupDbParameterGroupOutputArgs, opts ...pulumi.InvokeOption) LookupDbParameterGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDbParameterGroupResult, error) {
			args := v.(LookupDbParameterGroupArgs)
			r, err := LookupDbParameterGroup(ctx, &args, opts...)
			var s LookupDbParameterGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDbParameterGroupResultOutput)
}

type LookupDbParameterGroupOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupDbParameterGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDbParameterGroupArgs)(nil)).Elem()
}

type LookupDbParameterGroupResultOutput struct{ *pulumi.OutputState }

func (LookupDbParameterGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDbParameterGroupResult)(nil)).Elem()
}

func (o LookupDbParameterGroupResultOutput) ToLookupDbParameterGroupResultOutput() LookupDbParameterGroupResultOutput {
	return o
}

func (o LookupDbParameterGroupResultOutput) ToLookupDbParameterGroupResultOutputWithContext(ctx context.Context) LookupDbParameterGroupResultOutput {
	return o
}

func (o LookupDbParameterGroupResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDbParameterGroupResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Neptune::DBParameterGroup` for more information about the expected schema for this property.
func (o LookupDbParameterGroupResultOutput) Parameters() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupDbParameterGroupResult) interface{} { return v.Parameters }).(pulumi.AnyOutput)
}

func (o LookupDbParameterGroupResultOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupDbParameterGroupResult) []aws.Tag { return v.Tags }).(aws.TagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDbParameterGroupResultOutput{})
}

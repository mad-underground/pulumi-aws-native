// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ssm

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::SSM::Parameter
func LookupParameter(ctx *pulumi.Context, args *LookupParameterArgs, opts ...pulumi.InvokeOption) (*LookupParameterResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupParameterResult
	err := ctx.Invoke("aws-native:ssm:getParameter", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupParameterArgs struct {
	// The name of the parameter.
	Name string `pulumi:"name"`
}

type LookupParameterResult struct {
	// The corresponding DataType of the parameter.
	DataType *ParameterDataType `pulumi:"dataType"`
	// The type of the parameter.
	Type *ParameterType `pulumi:"type"`
	// The value associated with the parameter.
	Value *string `pulumi:"value"`
}

func LookupParameterOutput(ctx *pulumi.Context, args LookupParameterOutputArgs, opts ...pulumi.InvokeOption) LookupParameterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupParameterResult, error) {
			args := v.(LookupParameterArgs)
			r, err := LookupParameter(ctx, &args, opts...)
			var s LookupParameterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupParameterResultOutput)
}

type LookupParameterOutputArgs struct {
	// The name of the parameter.
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupParameterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupParameterArgs)(nil)).Elem()
}

type LookupParameterResultOutput struct{ *pulumi.OutputState }

func (LookupParameterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupParameterResult)(nil)).Elem()
}

func (o LookupParameterResultOutput) ToLookupParameterResultOutput() LookupParameterResultOutput {
	return o
}

func (o LookupParameterResultOutput) ToLookupParameterResultOutputWithContext(ctx context.Context) LookupParameterResultOutput {
	return o
}

// The corresponding DataType of the parameter.
func (o LookupParameterResultOutput) DataType() ParameterDataTypePtrOutput {
	return o.ApplyT(func(v LookupParameterResult) *ParameterDataType { return v.DataType }).(ParameterDataTypePtrOutput)
}

// The type of the parameter.
func (o LookupParameterResultOutput) Type() ParameterTypePtrOutput {
	return o.ApplyT(func(v LookupParameterResult) *ParameterType { return v.Type }).(ParameterTypePtrOutput)
}

// The value associated with the parameter.
func (o LookupParameterResultOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupParameterResult) *string { return v.Value }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupParameterResultOutput{})
}

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
	Id string `pulumi:"id"`
}

type LookupParameterResult struct {
	AllowedPattern *string     `pulumi:"allowedPattern"`
	DataType       *string     `pulumi:"dataType"`
	Description    *string     `pulumi:"description"`
	Id             *string     `pulumi:"id"`
	Policies       *string     `pulumi:"policies"`
	Tags           interface{} `pulumi:"tags"`
	Tier           *string     `pulumi:"tier"`
	Type           *string     `pulumi:"type"`
	Value          *string     `pulumi:"value"`
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
	Id pulumi.StringInput `pulumi:"id"`
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

func (o LookupParameterResultOutput) AllowedPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupParameterResult) *string { return v.AllowedPattern }).(pulumi.StringPtrOutput)
}

func (o LookupParameterResultOutput) DataType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupParameterResult) *string { return v.DataType }).(pulumi.StringPtrOutput)
}

func (o LookupParameterResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupParameterResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupParameterResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupParameterResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupParameterResultOutput) Policies() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupParameterResult) *string { return v.Policies }).(pulumi.StringPtrOutput)
}

func (o LookupParameterResultOutput) Tags() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupParameterResult) interface{} { return v.Tags }).(pulumi.AnyOutput)
}

func (o LookupParameterResultOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupParameterResult) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

func (o LookupParameterResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupParameterResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o LookupParameterResultOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupParameterResult) *string { return v.Value }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupParameterResultOutput{})
}

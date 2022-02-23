// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cognito

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Cognito::UserPoolGroup
func LookupUserPoolGroup(ctx *pulumi.Context, args *LookupUserPoolGroupArgs, opts ...pulumi.InvokeOption) (*LookupUserPoolGroupResult, error) {
	var rv LookupUserPoolGroupResult
	err := ctx.Invoke("aws-native:cognito:getUserPoolGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupUserPoolGroupArgs struct {
	Id string `pulumi:"id"`
}

type LookupUserPoolGroupResult struct {
	Description *string  `pulumi:"description"`
	Id          *string  `pulumi:"id"`
	Precedence  *float64 `pulumi:"precedence"`
	RoleArn     *string  `pulumi:"roleArn"`
}

func LookupUserPoolGroupOutput(ctx *pulumi.Context, args LookupUserPoolGroupOutputArgs, opts ...pulumi.InvokeOption) LookupUserPoolGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupUserPoolGroupResult, error) {
			args := v.(LookupUserPoolGroupArgs)
			r, err := LookupUserPoolGroup(ctx, &args, opts...)
			return *r, err
		}).(LookupUserPoolGroupResultOutput)
}

type LookupUserPoolGroupOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupUserPoolGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserPoolGroupArgs)(nil)).Elem()
}

type LookupUserPoolGroupResultOutput struct{ *pulumi.OutputState }

func (LookupUserPoolGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserPoolGroupResult)(nil)).Elem()
}

func (o LookupUserPoolGroupResultOutput) ToLookupUserPoolGroupResultOutput() LookupUserPoolGroupResultOutput {
	return o
}

func (o LookupUserPoolGroupResultOutput) ToLookupUserPoolGroupResultOutputWithContext(ctx context.Context) LookupUserPoolGroupResultOutput {
	return o
}

func (o LookupUserPoolGroupResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUserPoolGroupResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupUserPoolGroupResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUserPoolGroupResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupUserPoolGroupResultOutput) Precedence() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupUserPoolGroupResult) *float64 { return v.Precedence }).(pulumi.Float64PtrOutput)
}

func (o LookupUserPoolGroupResultOutput) RoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUserPoolGroupResult) *string { return v.RoleArn }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupUserPoolGroupResultOutput{})
}
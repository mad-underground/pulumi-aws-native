// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cognito

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Resource Type definition for AWS::Cognito::UserPoolGroup
func LookupUserPoolGroup(ctx *pulumi.Context, args *LookupUserPoolGroupArgs, opts ...pulumi.InvokeOption) (*LookupUserPoolGroupResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupUserPoolGroupResult
	err := ctx.Invoke("aws-native:cognito:getUserPoolGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupUserPoolGroupArgs struct {
	GroupName  string `pulumi:"groupName"`
	UserPoolId string `pulumi:"userPoolId"`
}

type LookupUserPoolGroupResult struct {
	Description *string `pulumi:"description"`
	Precedence  *int    `pulumi:"precedence"`
	RoleArn     *string `pulumi:"roleArn"`
}

func LookupUserPoolGroupOutput(ctx *pulumi.Context, args LookupUserPoolGroupOutputArgs, opts ...pulumi.InvokeOption) LookupUserPoolGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupUserPoolGroupResult, error) {
			args := v.(LookupUserPoolGroupArgs)
			r, err := LookupUserPoolGroup(ctx, &args, opts...)
			var s LookupUserPoolGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupUserPoolGroupResultOutput)
}

type LookupUserPoolGroupOutputArgs struct {
	GroupName  pulumi.StringInput `pulumi:"groupName"`
	UserPoolId pulumi.StringInput `pulumi:"userPoolId"`
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

func (o LookupUserPoolGroupResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupUserPoolGroupResult] {
	return pulumix.Output[LookupUserPoolGroupResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupUserPoolGroupResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUserPoolGroupResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupUserPoolGroupResultOutput) Precedence() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupUserPoolGroupResult) *int { return v.Precedence }).(pulumi.IntPtrOutput)
}

func (o LookupUserPoolGroupResultOutput) RoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupUserPoolGroupResult) *string { return v.RoleArn }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupUserPoolGroupResultOutput{})
}

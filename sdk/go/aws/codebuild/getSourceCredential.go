// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package codebuild

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::CodeBuild::SourceCredential
func LookupSourceCredential(ctx *pulumi.Context, args *LookupSourceCredentialArgs, opts ...pulumi.InvokeOption) (*LookupSourceCredentialResult, error) {
	var rv LookupSourceCredentialResult
	err := ctx.Invoke("aws-native:codebuild:getSourceCredential", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSourceCredentialArgs struct {
	Id string `pulumi:"id"`
}

type LookupSourceCredentialResult struct {
	AuthType *string `pulumi:"authType"`
	Id       *string `pulumi:"id"`
	Token    *string `pulumi:"token"`
	Username *string `pulumi:"username"`
}

func LookupSourceCredentialOutput(ctx *pulumi.Context, args LookupSourceCredentialOutputArgs, opts ...pulumi.InvokeOption) LookupSourceCredentialResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceCredentialResult, error) {
			args := v.(LookupSourceCredentialArgs)
			r, err := LookupSourceCredential(ctx, &args, opts...)
			return *r, err
		}).(LookupSourceCredentialResultOutput)
}

type LookupSourceCredentialOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupSourceCredentialOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceCredentialArgs)(nil)).Elem()
}

type LookupSourceCredentialResultOutput struct{ *pulumi.OutputState }

func (LookupSourceCredentialResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceCredentialResult)(nil)).Elem()
}

func (o LookupSourceCredentialResultOutput) ToLookupSourceCredentialResultOutput() LookupSourceCredentialResultOutput {
	return o
}

func (o LookupSourceCredentialResultOutput) ToLookupSourceCredentialResultOutputWithContext(ctx context.Context) LookupSourceCredentialResultOutput {
	return o
}

func (o LookupSourceCredentialResultOutput) AuthType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceCredentialResult) *string { return v.AuthType }).(pulumi.StringPtrOutput)
}

func (o LookupSourceCredentialResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceCredentialResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupSourceCredentialResultOutput) Token() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceCredentialResult) *string { return v.Token }).(pulumi.StringPtrOutput)
}

func (o LookupSourceCredentialResultOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceCredentialResult) *string { return v.Username }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceCredentialResultOutput{})
}

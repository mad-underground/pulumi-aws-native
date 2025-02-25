// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package appsync

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::AppSync::ApiKey
func LookupApiKey(ctx *pulumi.Context, args *LookupApiKeyArgs, opts ...pulumi.InvokeOption) (*LookupApiKeyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupApiKeyResult
	err := ctx.Invoke("aws-native:appsync:getApiKey", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApiKeyArgs struct {
	ApiKeyId string `pulumi:"apiKeyId"`
}

type LookupApiKeyResult struct {
	ApiKey      *string  `pulumi:"apiKey"`
	ApiKeyId    *string  `pulumi:"apiKeyId"`
	Arn         *string  `pulumi:"arn"`
	Description *string  `pulumi:"description"`
	Expires     *float64 `pulumi:"expires"`
}

func LookupApiKeyOutput(ctx *pulumi.Context, args LookupApiKeyOutputArgs, opts ...pulumi.InvokeOption) LookupApiKeyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApiKeyResult, error) {
			args := v.(LookupApiKeyArgs)
			r, err := LookupApiKey(ctx, &args, opts...)
			var s LookupApiKeyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupApiKeyResultOutput)
}

type LookupApiKeyOutputArgs struct {
	ApiKeyId pulumi.StringInput `pulumi:"apiKeyId"`
}

func (LookupApiKeyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiKeyArgs)(nil)).Elem()
}

type LookupApiKeyResultOutput struct{ *pulumi.OutputState }

func (LookupApiKeyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiKeyResult)(nil)).Elem()
}

func (o LookupApiKeyResultOutput) ToLookupApiKeyResultOutput() LookupApiKeyResultOutput {
	return o
}

func (o LookupApiKeyResultOutput) ToLookupApiKeyResultOutputWithContext(ctx context.Context) LookupApiKeyResultOutput {
	return o
}

func (o LookupApiKeyResultOutput) ApiKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiKeyResult) *string { return v.ApiKey }).(pulumi.StringPtrOutput)
}

func (o LookupApiKeyResultOutput) ApiKeyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiKeyResult) *string { return v.ApiKeyId }).(pulumi.StringPtrOutput)
}

func (o LookupApiKeyResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiKeyResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupApiKeyResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiKeyResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupApiKeyResultOutput) Expires() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupApiKeyResult) *float64 { return v.Expires }).(pulumi.Float64PtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApiKeyResultOutput{})
}

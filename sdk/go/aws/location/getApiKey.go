// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package location

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of AWS::Location::APIKey Resource Type
func LookupApiKey(ctx *pulumi.Context, args *LookupApiKeyArgs, opts ...pulumi.InvokeOption) (*LookupApiKeyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupApiKeyResult
	err := ctx.Invoke("aws-native:location:getApiKey", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApiKeyArgs struct {
	KeyName string `pulumi:"keyName"`
}

type LookupApiKeyResult struct {
	Arn          *string             `pulumi:"arn"`
	CreateTime   *string             `pulumi:"createTime"`
	Description  *string             `pulumi:"description"`
	ExpireTime   *string             `pulumi:"expireTime"`
	KeyArn       *string             `pulumi:"keyArn"`
	Restrictions *ApiKeyRestrictions `pulumi:"restrictions"`
	// An array of key-value pairs to apply to this resource.
	Tags       []aws.Tag `pulumi:"tags"`
	UpdateTime *string   `pulumi:"updateTime"`
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
	KeyName pulumi.StringInput `pulumi:"keyName"`
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

func (o LookupApiKeyResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiKeyResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupApiKeyResultOutput) CreateTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiKeyResult) *string { return v.CreateTime }).(pulumi.StringPtrOutput)
}

func (o LookupApiKeyResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiKeyResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupApiKeyResultOutput) ExpireTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiKeyResult) *string { return v.ExpireTime }).(pulumi.StringPtrOutput)
}

func (o LookupApiKeyResultOutput) KeyArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiKeyResult) *string { return v.KeyArn }).(pulumi.StringPtrOutput)
}

func (o LookupApiKeyResultOutput) Restrictions() ApiKeyRestrictionsPtrOutput {
	return o.ApplyT(func(v LookupApiKeyResult) *ApiKeyRestrictions { return v.Restrictions }).(ApiKeyRestrictionsPtrOutput)
}

// An array of key-value pairs to apply to this resource.
func (o LookupApiKeyResultOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v LookupApiKeyResult) []aws.Tag { return v.Tags }).(aws.TagArrayOutput)
}

func (o LookupApiKeyResultOutput) UpdateTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiKeyResult) *string { return v.UpdateTime }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApiKeyResultOutput{})
}

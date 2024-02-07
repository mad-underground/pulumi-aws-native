// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package appsync

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::AppSync::ApiCache
func LookupApiCache(ctx *pulumi.Context, args *LookupApiCacheArgs, opts ...pulumi.InvokeOption) (*LookupApiCacheResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupApiCacheResult
	err := ctx.Invoke("aws-native:appsync:getApiCache", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApiCacheArgs struct {
	Id string `pulumi:"id"`
}

type LookupApiCacheResult struct {
	ApiCachingBehavior       *string  `pulumi:"apiCachingBehavior"`
	AtRestEncryptionEnabled  *bool    `pulumi:"atRestEncryptionEnabled"`
	Id                       *string  `pulumi:"id"`
	TransitEncryptionEnabled *bool    `pulumi:"transitEncryptionEnabled"`
	Ttl                      *float64 `pulumi:"ttl"`
	Type                     *string  `pulumi:"type"`
}

func LookupApiCacheOutput(ctx *pulumi.Context, args LookupApiCacheOutputArgs, opts ...pulumi.InvokeOption) LookupApiCacheResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApiCacheResult, error) {
			args := v.(LookupApiCacheArgs)
			r, err := LookupApiCache(ctx, &args, opts...)
			var s LookupApiCacheResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupApiCacheResultOutput)
}

type LookupApiCacheOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupApiCacheOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiCacheArgs)(nil)).Elem()
}

type LookupApiCacheResultOutput struct{ *pulumi.OutputState }

func (LookupApiCacheResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiCacheResult)(nil)).Elem()
}

func (o LookupApiCacheResultOutput) ToLookupApiCacheResultOutput() LookupApiCacheResultOutput {
	return o
}

func (o LookupApiCacheResultOutput) ToLookupApiCacheResultOutputWithContext(ctx context.Context) LookupApiCacheResultOutput {
	return o
}

func (o LookupApiCacheResultOutput) ApiCachingBehavior() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiCacheResult) *string { return v.ApiCachingBehavior }).(pulumi.StringPtrOutput)
}

func (o LookupApiCacheResultOutput) AtRestEncryptionEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupApiCacheResult) *bool { return v.AtRestEncryptionEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupApiCacheResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiCacheResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupApiCacheResultOutput) TransitEncryptionEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupApiCacheResult) *bool { return v.TransitEncryptionEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupApiCacheResultOutput) Ttl() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupApiCacheResult) *float64 { return v.Ttl }).(pulumi.Float64PtrOutput)
}

func (o LookupApiCacheResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiCacheResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApiCacheResultOutput{})
}

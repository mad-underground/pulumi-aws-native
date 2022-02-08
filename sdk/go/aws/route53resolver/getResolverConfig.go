// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package route53resolver

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::Route53Resolver::ResolverConfig.
func LookupResolverConfig(ctx *pulumi.Context, args *LookupResolverConfigArgs, opts ...pulumi.InvokeOption) (*LookupResolverConfigResult, error) {
	var rv LookupResolverConfigResult
	err := ctx.Invoke("aws-native:route53resolver:getResolverConfig", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupResolverConfigArgs struct {
	// Id
	Id string `pulumi:"id"`
}

type LookupResolverConfigResult struct {
	// ResolverAutodefinedReverseStatus, possible values are ENABLING, ENABLED, DISABLING AND DISABLED.
	AutodefinedReverse *ResolverConfigAutodefinedReverse `pulumi:"autodefinedReverse"`
	// Id
	Id *string `pulumi:"id"`
	// AccountId
	OwnerId *string `pulumi:"ownerId"`
}

func LookupResolverConfigOutput(ctx *pulumi.Context, args LookupResolverConfigOutputArgs, opts ...pulumi.InvokeOption) LookupResolverConfigResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupResolverConfigResult, error) {
			args := v.(LookupResolverConfigArgs)
			r, err := LookupResolverConfig(ctx, &args, opts...)
			return *r, err
		}).(LookupResolverConfigResultOutput)
}

type LookupResolverConfigOutputArgs struct {
	// Id
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupResolverConfigOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResolverConfigArgs)(nil)).Elem()
}

type LookupResolverConfigResultOutput struct{ *pulumi.OutputState }

func (LookupResolverConfigResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResolverConfigResult)(nil)).Elem()
}

func (o LookupResolverConfigResultOutput) ToLookupResolverConfigResultOutput() LookupResolverConfigResultOutput {
	return o
}

func (o LookupResolverConfigResultOutput) ToLookupResolverConfigResultOutputWithContext(ctx context.Context) LookupResolverConfigResultOutput {
	return o
}

// ResolverAutodefinedReverseStatus, possible values are ENABLING, ENABLED, DISABLING AND DISABLED.
func (o LookupResolverConfigResultOutput) AutodefinedReverse() ResolverConfigAutodefinedReversePtrOutput {
	return o.ApplyT(func(v LookupResolverConfigResult) *ResolverConfigAutodefinedReverse { return v.AutodefinedReverse }).(ResolverConfigAutodefinedReversePtrOutput)
}

// Id
func (o LookupResolverConfigResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupResolverConfigResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// AccountId
func (o LookupResolverConfigResultOutput) OwnerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupResolverConfigResult) *string { return v.OwnerId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupResolverConfigResultOutput{})
}

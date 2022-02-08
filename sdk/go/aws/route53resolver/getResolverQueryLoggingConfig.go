// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package route53resolver

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::Route53Resolver::ResolverQueryLoggingConfig.
func LookupResolverQueryLoggingConfig(ctx *pulumi.Context, args *LookupResolverQueryLoggingConfigArgs, opts ...pulumi.InvokeOption) (*LookupResolverQueryLoggingConfigResult, error) {
	var rv LookupResolverQueryLoggingConfigResult
	err := ctx.Invoke("aws-native:route53resolver:getResolverQueryLoggingConfig", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupResolverQueryLoggingConfigArgs struct {
	// ResourceId
	Id string `pulumi:"id"`
}

type LookupResolverQueryLoggingConfigResult struct {
	// Arn
	Arn *string `pulumi:"arn"`
	// Count
	AssociationCount *int `pulumi:"associationCount"`
	// Rfc3339TimeString
	CreationTime *string `pulumi:"creationTime"`
	// The id of the creator request.
	CreatorRequestId *string `pulumi:"creatorRequestId"`
	// ResourceId
	Id *string `pulumi:"id"`
	// AccountId
	OwnerId *string `pulumi:"ownerId"`
	// ShareStatus, possible values are NOT_SHARED, SHARED_WITH_ME, SHARED_BY_ME.
	ShareStatus *ResolverQueryLoggingConfigShareStatus `pulumi:"shareStatus"`
	// ResolverQueryLogConfigStatus, possible values are CREATING, CREATED, DELETED AND FAILED.
	Status *ResolverQueryLoggingConfigStatus `pulumi:"status"`
}

func LookupResolverQueryLoggingConfigOutput(ctx *pulumi.Context, args LookupResolverQueryLoggingConfigOutputArgs, opts ...pulumi.InvokeOption) LookupResolverQueryLoggingConfigResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupResolverQueryLoggingConfigResult, error) {
			args := v.(LookupResolverQueryLoggingConfigArgs)
			r, err := LookupResolverQueryLoggingConfig(ctx, &args, opts...)
			return *r, err
		}).(LookupResolverQueryLoggingConfigResultOutput)
}

type LookupResolverQueryLoggingConfigOutputArgs struct {
	// ResourceId
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupResolverQueryLoggingConfigOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResolverQueryLoggingConfigArgs)(nil)).Elem()
}

type LookupResolverQueryLoggingConfigResultOutput struct{ *pulumi.OutputState }

func (LookupResolverQueryLoggingConfigResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResolverQueryLoggingConfigResult)(nil)).Elem()
}

func (o LookupResolverQueryLoggingConfigResultOutput) ToLookupResolverQueryLoggingConfigResultOutput() LookupResolverQueryLoggingConfigResultOutput {
	return o
}

func (o LookupResolverQueryLoggingConfigResultOutput) ToLookupResolverQueryLoggingConfigResultOutputWithContext(ctx context.Context) LookupResolverQueryLoggingConfigResultOutput {
	return o
}

// Arn
func (o LookupResolverQueryLoggingConfigResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupResolverQueryLoggingConfigResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// Count
func (o LookupResolverQueryLoggingConfigResultOutput) AssociationCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupResolverQueryLoggingConfigResult) *int { return v.AssociationCount }).(pulumi.IntPtrOutput)
}

// Rfc3339TimeString
func (o LookupResolverQueryLoggingConfigResultOutput) CreationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupResolverQueryLoggingConfigResult) *string { return v.CreationTime }).(pulumi.StringPtrOutput)
}

// The id of the creator request.
func (o LookupResolverQueryLoggingConfigResultOutput) CreatorRequestId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupResolverQueryLoggingConfigResult) *string { return v.CreatorRequestId }).(pulumi.StringPtrOutput)
}

// ResourceId
func (o LookupResolverQueryLoggingConfigResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupResolverQueryLoggingConfigResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// AccountId
func (o LookupResolverQueryLoggingConfigResultOutput) OwnerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupResolverQueryLoggingConfigResult) *string { return v.OwnerId }).(pulumi.StringPtrOutput)
}

// ShareStatus, possible values are NOT_SHARED, SHARED_WITH_ME, SHARED_BY_ME.
func (o LookupResolverQueryLoggingConfigResultOutput) ShareStatus() ResolverQueryLoggingConfigShareStatusPtrOutput {
	return o.ApplyT(func(v LookupResolverQueryLoggingConfigResult) *ResolverQueryLoggingConfigShareStatus {
		return v.ShareStatus
	}).(ResolverQueryLoggingConfigShareStatusPtrOutput)
}

// ResolverQueryLogConfigStatus, possible values are CREATING, CREATED, DELETED AND FAILED.
func (o LookupResolverQueryLoggingConfigResultOutput) Status() ResolverQueryLoggingConfigStatusPtrOutput {
	return o.ApplyT(func(v LookupResolverQueryLoggingConfigResult) *ResolverQueryLoggingConfigStatus { return v.Status }).(ResolverQueryLoggingConfigStatusPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupResolverQueryLoggingConfigResultOutput{})
}

// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package omics

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of AWS::Omics::ReferenceStore Resource Type
func LookupReferenceStore(ctx *pulumi.Context, args *LookupReferenceStoreArgs, opts ...pulumi.InvokeOption) (*LookupReferenceStoreResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupReferenceStoreResult
	err := ctx.Invoke("aws-native:omics:getReferenceStore", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupReferenceStoreArgs struct {
	ReferenceStoreId string `pulumi:"referenceStoreId"`
}

type LookupReferenceStoreResult struct {
	// The store's ARN.
	Arn *string `pulumi:"arn"`
	// When the store was created.
	CreationTime     *string `pulumi:"creationTime"`
	ReferenceStoreId *string `pulumi:"referenceStoreId"`
}

func LookupReferenceStoreOutput(ctx *pulumi.Context, args LookupReferenceStoreOutputArgs, opts ...pulumi.InvokeOption) LookupReferenceStoreResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupReferenceStoreResult, error) {
			args := v.(LookupReferenceStoreArgs)
			r, err := LookupReferenceStore(ctx, &args, opts...)
			var s LookupReferenceStoreResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupReferenceStoreResultOutput)
}

type LookupReferenceStoreOutputArgs struct {
	ReferenceStoreId pulumi.StringInput `pulumi:"referenceStoreId"`
}

func (LookupReferenceStoreOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReferenceStoreArgs)(nil)).Elem()
}

type LookupReferenceStoreResultOutput struct{ *pulumi.OutputState }

func (LookupReferenceStoreResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupReferenceStoreResult)(nil)).Elem()
}

func (o LookupReferenceStoreResultOutput) ToLookupReferenceStoreResultOutput() LookupReferenceStoreResultOutput {
	return o
}

func (o LookupReferenceStoreResultOutput) ToLookupReferenceStoreResultOutputWithContext(ctx context.Context) LookupReferenceStoreResultOutput {
	return o
}

// The store's ARN.
func (o LookupReferenceStoreResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupReferenceStoreResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

// When the store was created.
func (o LookupReferenceStoreResultOutput) CreationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupReferenceStoreResult) *string { return v.CreationTime }).(pulumi.StringPtrOutput)
}

func (o LookupReferenceStoreResultOutput) ReferenceStoreId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupReferenceStoreResult) *string { return v.ReferenceStoreId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupReferenceStoreResultOutput{})
}

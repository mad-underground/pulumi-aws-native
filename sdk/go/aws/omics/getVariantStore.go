// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package omics

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of AWS::Omics::VariantStore Resource Type
func LookupVariantStore(ctx *pulumi.Context, args *LookupVariantStoreArgs, opts ...pulumi.InvokeOption) (*LookupVariantStoreResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupVariantStoreResult
	err := ctx.Invoke("aws-native:omics:getVariantStore", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVariantStoreArgs struct {
	Name string `pulumi:"name"`
}

type LookupVariantStoreResult struct {
	CreationTime   *string                  `pulumi:"creationTime"`
	Description    *string                  `pulumi:"description"`
	Id             *string                  `pulumi:"id"`
	Status         *VariantStoreStoreStatus `pulumi:"status"`
	StatusMessage  *string                  `pulumi:"statusMessage"`
	StoreArn       *string                  `pulumi:"storeArn"`
	StoreSizeBytes *float64                 `pulumi:"storeSizeBytes"`
	UpdateTime     *string                  `pulumi:"updateTime"`
}

func LookupVariantStoreOutput(ctx *pulumi.Context, args LookupVariantStoreOutputArgs, opts ...pulumi.InvokeOption) LookupVariantStoreResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVariantStoreResult, error) {
			args := v.(LookupVariantStoreArgs)
			r, err := LookupVariantStore(ctx, &args, opts...)
			var s LookupVariantStoreResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVariantStoreResultOutput)
}

type LookupVariantStoreOutputArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupVariantStoreOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVariantStoreArgs)(nil)).Elem()
}

type LookupVariantStoreResultOutput struct{ *pulumi.OutputState }

func (LookupVariantStoreResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVariantStoreResult)(nil)).Elem()
}

func (o LookupVariantStoreResultOutput) ToLookupVariantStoreResultOutput() LookupVariantStoreResultOutput {
	return o
}

func (o LookupVariantStoreResultOutput) ToLookupVariantStoreResultOutputWithContext(ctx context.Context) LookupVariantStoreResultOutput {
	return o
}

func (o LookupVariantStoreResultOutput) CreationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVariantStoreResult) *string { return v.CreationTime }).(pulumi.StringPtrOutput)
}

func (o LookupVariantStoreResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVariantStoreResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupVariantStoreResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVariantStoreResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupVariantStoreResultOutput) Status() VariantStoreStoreStatusPtrOutput {
	return o.ApplyT(func(v LookupVariantStoreResult) *VariantStoreStoreStatus { return v.Status }).(VariantStoreStoreStatusPtrOutput)
}

func (o LookupVariantStoreResultOutput) StatusMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVariantStoreResult) *string { return v.StatusMessage }).(pulumi.StringPtrOutput)
}

func (o LookupVariantStoreResultOutput) StoreArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVariantStoreResult) *string { return v.StoreArn }).(pulumi.StringPtrOutput)
}

func (o LookupVariantStoreResultOutput) StoreSizeBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupVariantStoreResult) *float64 { return v.StoreSizeBytes }).(pulumi.Float64PtrOutput)
}

func (o LookupVariantStoreResultOutput) UpdateTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVariantStoreResult) *string { return v.UpdateTime }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVariantStoreResultOutput{})
}

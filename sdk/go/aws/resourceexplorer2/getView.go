// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package resourceexplorer2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of AWS::ResourceExplorer2::View Resource Type
func LookupView(ctx *pulumi.Context, args *LookupViewArgs, opts ...pulumi.InvokeOption) (*LookupViewResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupViewResult
	err := ctx.Invoke("aws-native:resourceexplorer2:getView", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupViewArgs struct {
	ViewArn string `pulumi:"viewArn"`
}

type LookupViewResult struct {
	Filters            *ViewFilters           `pulumi:"filters"`
	IncludedProperties []ViewIncludedProperty `pulumi:"includedProperties"`
	Tags               *ViewTagMap            `pulumi:"tags"`
	ViewArn            *string                `pulumi:"viewArn"`
}

func LookupViewOutput(ctx *pulumi.Context, args LookupViewOutputArgs, opts ...pulumi.InvokeOption) LookupViewResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupViewResult, error) {
			args := v.(LookupViewArgs)
			r, err := LookupView(ctx, &args, opts...)
			var s LookupViewResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupViewResultOutput)
}

type LookupViewOutputArgs struct {
	ViewArn pulumi.StringInput `pulumi:"viewArn"`
}

func (LookupViewOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupViewArgs)(nil)).Elem()
}

type LookupViewResultOutput struct{ *pulumi.OutputState }

func (LookupViewResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupViewResult)(nil)).Elem()
}

func (o LookupViewResultOutput) ToLookupViewResultOutput() LookupViewResultOutput {
	return o
}

func (o LookupViewResultOutput) ToLookupViewResultOutputWithContext(ctx context.Context) LookupViewResultOutput {
	return o
}

func (o LookupViewResultOutput) Filters() ViewFiltersPtrOutput {
	return o.ApplyT(func(v LookupViewResult) *ViewFilters { return v.Filters }).(ViewFiltersPtrOutput)
}

func (o LookupViewResultOutput) IncludedProperties() ViewIncludedPropertyArrayOutput {
	return o.ApplyT(func(v LookupViewResult) []ViewIncludedProperty { return v.IncludedProperties }).(ViewIncludedPropertyArrayOutput)
}

func (o LookupViewResultOutput) Tags() ViewTagMapPtrOutput {
	return o.ApplyT(func(v LookupViewResult) *ViewTagMap { return v.Tags }).(ViewTagMapPtrOutput)
}

func (o LookupViewResultOutput) ViewArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupViewResult) *string { return v.ViewArn }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupViewResultOutput{})
}

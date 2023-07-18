// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package servicecatalogappregistry

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Schema for AWS::ServiceCatalogAppRegistry::AttributeGroupAssociation.
func LookupAttributeGroupAssociation(ctx *pulumi.Context, args *LookupAttributeGroupAssociationArgs, opts ...pulumi.InvokeOption) (*LookupAttributeGroupAssociationResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupAttributeGroupAssociationResult
	err := ctx.Invoke("aws-native:servicecatalogappregistry:getAttributeGroupAssociation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAttributeGroupAssociationArgs struct {
	Id string `pulumi:"id"`
}

type LookupAttributeGroupAssociationResult struct {
	ApplicationArn    *string `pulumi:"applicationArn"`
	AttributeGroupArn *string `pulumi:"attributeGroupArn"`
	Id                *string `pulumi:"id"`
}

func LookupAttributeGroupAssociationOutput(ctx *pulumi.Context, args LookupAttributeGroupAssociationOutputArgs, opts ...pulumi.InvokeOption) LookupAttributeGroupAssociationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAttributeGroupAssociationResult, error) {
			args := v.(LookupAttributeGroupAssociationArgs)
			r, err := LookupAttributeGroupAssociation(ctx, &args, opts...)
			var s LookupAttributeGroupAssociationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAttributeGroupAssociationResultOutput)
}

type LookupAttributeGroupAssociationOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupAttributeGroupAssociationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAttributeGroupAssociationArgs)(nil)).Elem()
}

type LookupAttributeGroupAssociationResultOutput struct{ *pulumi.OutputState }

func (LookupAttributeGroupAssociationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAttributeGroupAssociationResult)(nil)).Elem()
}

func (o LookupAttributeGroupAssociationResultOutput) ToLookupAttributeGroupAssociationResultOutput() LookupAttributeGroupAssociationResultOutput {
	return o
}

func (o LookupAttributeGroupAssociationResultOutput) ToLookupAttributeGroupAssociationResultOutputWithContext(ctx context.Context) LookupAttributeGroupAssociationResultOutput {
	return o
}

func (o LookupAttributeGroupAssociationResultOutput) ApplicationArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAttributeGroupAssociationResult) *string { return v.ApplicationArn }).(pulumi.StringPtrOutput)
}

func (o LookupAttributeGroupAssociationResultOutput) AttributeGroupArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAttributeGroupAssociationResult) *string { return v.AttributeGroupArn }).(pulumi.StringPtrOutput)
}

func (o LookupAttributeGroupAssociationResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAttributeGroupAssociationResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAttributeGroupAssociationResultOutput{})
}

// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package servicecatalogappregistry

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Schema for AWS::ServiceCatalogAppRegistry::AttributeGroupAssociation.
func LookupAttributeGroupAssociation(ctx *pulumi.Context, args *LookupAttributeGroupAssociationArgs, opts ...pulumi.InvokeOption) (*LookupAttributeGroupAssociationResult, error) {
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
	// The name or the Id of the Application.
	Application    *string `pulumi:"application"`
	ApplicationArn *string `pulumi:"applicationArn"`
	// The name or the Id of the AttributeGroup.
	AttributeGroup    *string `pulumi:"attributeGroup"`
	AttributeGroupArn *string `pulumi:"attributeGroupArn"`
	Id                *string `pulumi:"id"`
}

func LookupAttributeGroupAssociationOutput(ctx *pulumi.Context, args LookupAttributeGroupAssociationOutputArgs, opts ...pulumi.InvokeOption) LookupAttributeGroupAssociationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAttributeGroupAssociationResult, error) {
			args := v.(LookupAttributeGroupAssociationArgs)
			r, err := LookupAttributeGroupAssociation(ctx, &args, opts...)
			return *r, err
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

// The name or the Id of the Application.
func (o LookupAttributeGroupAssociationResultOutput) Application() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAttributeGroupAssociationResult) *string { return v.Application }).(pulumi.StringPtrOutput)
}

func (o LookupAttributeGroupAssociationResultOutput) ApplicationArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAttributeGroupAssociationResult) *string { return v.ApplicationArn }).(pulumi.StringPtrOutput)
}

// The name or the Id of the AttributeGroup.
func (o LookupAttributeGroupAssociationResultOutput) AttributeGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAttributeGroupAssociationResult) *string { return v.AttributeGroup }).(pulumi.StringPtrOutput)
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

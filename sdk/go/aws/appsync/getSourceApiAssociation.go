// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package appsync

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::AppSync::SourceApiAssociation
func LookupSourceApiAssociation(ctx *pulumi.Context, args *LookupSourceApiAssociationArgs, opts ...pulumi.InvokeOption) (*LookupSourceApiAssociationResult, error) {
	var rv LookupSourceApiAssociationResult
	err := ctx.Invoke("aws-native:appsync:getSourceApiAssociation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSourceApiAssociationArgs struct {
	// ARN of the SourceApiAssociation.
	AssociationArn string `pulumi:"associationArn"`
}

type LookupSourceApiAssociationResult struct {
	// ARN of the SourceApiAssociation.
	AssociationArn *string `pulumi:"associationArn"`
	// Id of the SourceApiAssociation.
	AssociationId *string `pulumi:"associationId"`
	// Description of the SourceApiAssociation.
	Description *string `pulumi:"description"`
	// Date of last schema successful merge.
	LastSuccessfulMergeDate *string `pulumi:"lastSuccessfulMergeDate"`
	// ARN of the Merged API in the association.
	MergedApiArn *string `pulumi:"mergedApiArn"`
	// GraphQLApiId of the Merged API in the association.
	MergedApiId *string `pulumi:"mergedApiId"`
	// ARN of the source API in the association.
	SourceApiArn *string `pulumi:"sourceApiArn"`
	// Customized configuration for SourceApiAssociation.
	SourceApiAssociationConfig *SourceApiAssociationConfig `pulumi:"sourceApiAssociationConfig"`
	// Current status of SourceApiAssociation.
	SourceApiAssociationStatus *SourceApiAssociationStatus `pulumi:"sourceApiAssociationStatus"`
	// Current SourceApiAssociation status details.
	SourceApiAssociationStatusDetail *string `pulumi:"sourceApiAssociationStatusDetail"`
	// GraphQLApiId of the source API in the association.
	SourceApiId *string `pulumi:"sourceApiId"`
}

func LookupSourceApiAssociationOutput(ctx *pulumi.Context, args LookupSourceApiAssociationOutputArgs, opts ...pulumi.InvokeOption) LookupSourceApiAssociationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSourceApiAssociationResult, error) {
			args := v.(LookupSourceApiAssociationArgs)
			r, err := LookupSourceApiAssociation(ctx, &args, opts...)
			var s LookupSourceApiAssociationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSourceApiAssociationResultOutput)
}

type LookupSourceApiAssociationOutputArgs struct {
	// ARN of the SourceApiAssociation.
	AssociationArn pulumi.StringInput `pulumi:"associationArn"`
}

func (LookupSourceApiAssociationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceApiAssociationArgs)(nil)).Elem()
}

type LookupSourceApiAssociationResultOutput struct{ *pulumi.OutputState }

func (LookupSourceApiAssociationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSourceApiAssociationResult)(nil)).Elem()
}

func (o LookupSourceApiAssociationResultOutput) ToLookupSourceApiAssociationResultOutput() LookupSourceApiAssociationResultOutput {
	return o
}

func (o LookupSourceApiAssociationResultOutput) ToLookupSourceApiAssociationResultOutputWithContext(ctx context.Context) LookupSourceApiAssociationResultOutput {
	return o
}

// ARN of the SourceApiAssociation.
func (o LookupSourceApiAssociationResultOutput) AssociationArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceApiAssociationResult) *string { return v.AssociationArn }).(pulumi.StringPtrOutput)
}

// Id of the SourceApiAssociation.
func (o LookupSourceApiAssociationResultOutput) AssociationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceApiAssociationResult) *string { return v.AssociationId }).(pulumi.StringPtrOutput)
}

// Description of the SourceApiAssociation.
func (o LookupSourceApiAssociationResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceApiAssociationResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Date of last schema successful merge.
func (o LookupSourceApiAssociationResultOutput) LastSuccessfulMergeDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceApiAssociationResult) *string { return v.LastSuccessfulMergeDate }).(pulumi.StringPtrOutput)
}

// ARN of the Merged API in the association.
func (o LookupSourceApiAssociationResultOutput) MergedApiArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceApiAssociationResult) *string { return v.MergedApiArn }).(pulumi.StringPtrOutput)
}

// GraphQLApiId of the Merged API in the association.
func (o LookupSourceApiAssociationResultOutput) MergedApiId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceApiAssociationResult) *string { return v.MergedApiId }).(pulumi.StringPtrOutput)
}

// ARN of the source API in the association.
func (o LookupSourceApiAssociationResultOutput) SourceApiArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceApiAssociationResult) *string { return v.SourceApiArn }).(pulumi.StringPtrOutput)
}

// Customized configuration for SourceApiAssociation.
func (o LookupSourceApiAssociationResultOutput) SourceApiAssociationConfig() SourceApiAssociationConfigPtrOutput {
	return o.ApplyT(func(v LookupSourceApiAssociationResult) *SourceApiAssociationConfig {
		return v.SourceApiAssociationConfig
	}).(SourceApiAssociationConfigPtrOutput)
}

// Current status of SourceApiAssociation.
func (o LookupSourceApiAssociationResultOutput) SourceApiAssociationStatus() SourceApiAssociationStatusPtrOutput {
	return o.ApplyT(func(v LookupSourceApiAssociationResult) *SourceApiAssociationStatus {
		return v.SourceApiAssociationStatus
	}).(SourceApiAssociationStatusPtrOutput)
}

// Current SourceApiAssociation status details.
func (o LookupSourceApiAssociationResultOutput) SourceApiAssociationStatusDetail() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceApiAssociationResult) *string { return v.SourceApiAssociationStatusDetail }).(pulumi.StringPtrOutput)
}

// GraphQLApiId of the source API in the association.
func (o LookupSourceApiAssociationResultOutput) SourceApiId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSourceApiAssociationResult) *string { return v.SourceApiId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSourceApiAssociationResultOutput{})
}

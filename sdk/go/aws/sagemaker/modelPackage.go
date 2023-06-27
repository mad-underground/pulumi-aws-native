// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sagemaker

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::SageMaker::ModelPackage
type ModelPackage struct {
	pulumi.CustomResourceState

	AdditionalInferenceSpecifications      ModelPackageAdditionalInferenceSpecificationDefinitionArrayOutput `pulumi:"additionalInferenceSpecifications"`
	AdditionalInferenceSpecificationsToAdd ModelPackageAdditionalInferenceSpecificationDefinitionArrayOutput `pulumi:"additionalInferenceSpecificationsToAdd"`
	ApprovalDescription                    pulumi.StringPtrOutput                                            `pulumi:"approvalDescription"`
	CertifyForMarketplace                  pulumi.BoolPtrOutput                                              `pulumi:"certifyForMarketplace"`
	ClientToken                            pulumi.StringPtrOutput                                            `pulumi:"clientToken"`
	CreationTime                           pulumi.StringOutput                                               `pulumi:"creationTime"`
	CustomerMetadataProperties             ModelPackageCustomerMetadataPropertiesPtrOutput                   `pulumi:"customerMetadataProperties"`
	Domain                                 pulumi.StringPtrOutput                                            `pulumi:"domain"`
	DriftCheckBaselines                    ModelPackageDriftCheckBaselinesPtrOutput                          `pulumi:"driftCheckBaselines"`
	InferenceSpecification                 ModelPackageInferenceSpecificationPtrOutput                       `pulumi:"inferenceSpecification"`
	LastModifiedTime                       pulumi.StringPtrOutput                                            `pulumi:"lastModifiedTime"`
	MetadataProperties                     ModelPackageMetadataPropertiesPtrOutput                           `pulumi:"metadataProperties"`
	ModelApprovalStatus                    ModelPackageModelApprovalStatusPtrOutput                          `pulumi:"modelApprovalStatus"`
	ModelMetrics                           ModelPackageModelMetricsPtrOutput                                 `pulumi:"modelMetrics"`
	ModelPackageArn                        pulumi.StringOutput                                               `pulumi:"modelPackageArn"`
	ModelPackageDescription                pulumi.StringPtrOutput                                            `pulumi:"modelPackageDescription"`
	ModelPackageGroupName                  pulumi.StringPtrOutput                                            `pulumi:"modelPackageGroupName"`
	ModelPackageName                       pulumi.StringPtrOutput                                            `pulumi:"modelPackageName"`
	ModelPackageStatus                     ModelPackageStatusOutput                                          `pulumi:"modelPackageStatus"`
	ModelPackageStatusDetails              ModelPackageStatusDetailsPtrOutput                                `pulumi:"modelPackageStatusDetails"`
	ModelPackageVersion                    pulumi.IntPtrOutput                                               `pulumi:"modelPackageVersion"`
	SamplePayloadUrl                       pulumi.StringPtrOutput                                            `pulumi:"samplePayloadUrl"`
	SourceAlgorithmSpecification           ModelPackageSourceAlgorithmSpecificationPtrOutput                 `pulumi:"sourceAlgorithmSpecification"`
	// An array of key-value pairs to apply to this resource.
	Tags                    ModelPackageTagArrayOutput                   `pulumi:"tags"`
	Task                    pulumi.StringPtrOutput                       `pulumi:"task"`
	ValidationSpecification ModelPackageValidationSpecificationPtrOutput `pulumi:"validationSpecification"`
}

// NewModelPackage registers a new resource with the given unique name, arguments, and options.
func NewModelPackage(ctx *pulumi.Context,
	name string, args *ModelPackageArgs, opts ...pulumi.ResourceOption) (*ModelPackage, error) {
	if args == nil {
		args = &ModelPackageArgs{}
	}

	var resource ModelPackage
	err := ctx.RegisterResource("aws-native:sagemaker:ModelPackage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetModelPackage gets an existing ModelPackage resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetModelPackage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ModelPackageState, opts ...pulumi.ResourceOption) (*ModelPackage, error) {
	var resource ModelPackage
	err := ctx.ReadResource("aws-native:sagemaker:ModelPackage", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ModelPackage resources.
type modelPackageState struct {
}

type ModelPackageState struct {
}

func (ModelPackageState) ElementType() reflect.Type {
	return reflect.TypeOf((*modelPackageState)(nil)).Elem()
}

type modelPackageArgs struct {
	AdditionalInferenceSpecifications      []ModelPackageAdditionalInferenceSpecificationDefinition `pulumi:"additionalInferenceSpecifications"`
	AdditionalInferenceSpecificationsToAdd []ModelPackageAdditionalInferenceSpecificationDefinition `pulumi:"additionalInferenceSpecificationsToAdd"`
	ApprovalDescription                    *string                                                  `pulumi:"approvalDescription"`
	CertifyForMarketplace                  *bool                                                    `pulumi:"certifyForMarketplace"`
	ClientToken                            *string                                                  `pulumi:"clientToken"`
	CustomerMetadataProperties             *ModelPackageCustomerMetadataProperties                  `pulumi:"customerMetadataProperties"`
	Domain                                 *string                                                  `pulumi:"domain"`
	DriftCheckBaselines                    *ModelPackageDriftCheckBaselines                         `pulumi:"driftCheckBaselines"`
	InferenceSpecification                 *ModelPackageInferenceSpecification                      `pulumi:"inferenceSpecification"`
	LastModifiedTime                       *string                                                  `pulumi:"lastModifiedTime"`
	MetadataProperties                     *ModelPackageMetadataProperties                          `pulumi:"metadataProperties"`
	ModelApprovalStatus                    *ModelPackageModelApprovalStatus                         `pulumi:"modelApprovalStatus"`
	ModelMetrics                           *ModelPackageModelMetrics                                `pulumi:"modelMetrics"`
	ModelPackageDescription                *string                                                  `pulumi:"modelPackageDescription"`
	ModelPackageGroupName                  *string                                                  `pulumi:"modelPackageGroupName"`
	ModelPackageName                       *string                                                  `pulumi:"modelPackageName"`
	ModelPackageStatusDetails              *ModelPackageStatusDetails                               `pulumi:"modelPackageStatusDetails"`
	ModelPackageVersion                    *int                                                     `pulumi:"modelPackageVersion"`
	SamplePayloadUrl                       *string                                                  `pulumi:"samplePayloadUrl"`
	SourceAlgorithmSpecification           *ModelPackageSourceAlgorithmSpecification                `pulumi:"sourceAlgorithmSpecification"`
	// An array of key-value pairs to apply to this resource.
	Tags                    []ModelPackageTag                    `pulumi:"tags"`
	Task                    *string                              `pulumi:"task"`
	ValidationSpecification *ModelPackageValidationSpecification `pulumi:"validationSpecification"`
}

// The set of arguments for constructing a ModelPackage resource.
type ModelPackageArgs struct {
	AdditionalInferenceSpecifications      ModelPackageAdditionalInferenceSpecificationDefinitionArrayInput
	AdditionalInferenceSpecificationsToAdd ModelPackageAdditionalInferenceSpecificationDefinitionArrayInput
	ApprovalDescription                    pulumi.StringPtrInput
	CertifyForMarketplace                  pulumi.BoolPtrInput
	ClientToken                            pulumi.StringPtrInput
	CustomerMetadataProperties             ModelPackageCustomerMetadataPropertiesPtrInput
	Domain                                 pulumi.StringPtrInput
	DriftCheckBaselines                    ModelPackageDriftCheckBaselinesPtrInput
	InferenceSpecification                 ModelPackageInferenceSpecificationPtrInput
	LastModifiedTime                       pulumi.StringPtrInput
	MetadataProperties                     ModelPackageMetadataPropertiesPtrInput
	ModelApprovalStatus                    ModelPackageModelApprovalStatusPtrInput
	ModelMetrics                           ModelPackageModelMetricsPtrInput
	ModelPackageDescription                pulumi.StringPtrInput
	ModelPackageGroupName                  pulumi.StringPtrInput
	ModelPackageName                       pulumi.StringPtrInput
	ModelPackageStatusDetails              ModelPackageStatusDetailsPtrInput
	ModelPackageVersion                    pulumi.IntPtrInput
	SamplePayloadUrl                       pulumi.StringPtrInput
	SourceAlgorithmSpecification           ModelPackageSourceAlgorithmSpecificationPtrInput
	// An array of key-value pairs to apply to this resource.
	Tags                    ModelPackageTagArrayInput
	Task                    pulumi.StringPtrInput
	ValidationSpecification ModelPackageValidationSpecificationPtrInput
}

func (ModelPackageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*modelPackageArgs)(nil)).Elem()
}

type ModelPackageInput interface {
	pulumi.Input

	ToModelPackageOutput() ModelPackageOutput
	ToModelPackageOutputWithContext(ctx context.Context) ModelPackageOutput
}

func (*ModelPackage) ElementType() reflect.Type {
	return reflect.TypeOf((**ModelPackage)(nil)).Elem()
}

func (i *ModelPackage) ToModelPackageOutput() ModelPackageOutput {
	return i.ToModelPackageOutputWithContext(context.Background())
}

func (i *ModelPackage) ToModelPackageOutputWithContext(ctx context.Context) ModelPackageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModelPackageOutput)
}

type ModelPackageOutput struct{ *pulumi.OutputState }

func (ModelPackageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ModelPackage)(nil)).Elem()
}

func (o ModelPackageOutput) ToModelPackageOutput() ModelPackageOutput {
	return o
}

func (o ModelPackageOutput) ToModelPackageOutputWithContext(ctx context.Context) ModelPackageOutput {
	return o
}

func (o ModelPackageOutput) AdditionalInferenceSpecifications() ModelPackageAdditionalInferenceSpecificationDefinitionArrayOutput {
	return o.ApplyT(func(v *ModelPackage) ModelPackageAdditionalInferenceSpecificationDefinitionArrayOutput {
		return v.AdditionalInferenceSpecifications
	}).(ModelPackageAdditionalInferenceSpecificationDefinitionArrayOutput)
}

func (o ModelPackageOutput) AdditionalInferenceSpecificationsToAdd() ModelPackageAdditionalInferenceSpecificationDefinitionArrayOutput {
	return o.ApplyT(func(v *ModelPackage) ModelPackageAdditionalInferenceSpecificationDefinitionArrayOutput {
		return v.AdditionalInferenceSpecificationsToAdd
	}).(ModelPackageAdditionalInferenceSpecificationDefinitionArrayOutput)
}

func (o ModelPackageOutput) ApprovalDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ModelPackage) pulumi.StringPtrOutput { return v.ApprovalDescription }).(pulumi.StringPtrOutput)
}

func (o ModelPackageOutput) CertifyForMarketplace() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ModelPackage) pulumi.BoolPtrOutput { return v.CertifyForMarketplace }).(pulumi.BoolPtrOutput)
}

func (o ModelPackageOutput) ClientToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ModelPackage) pulumi.StringPtrOutput { return v.ClientToken }).(pulumi.StringPtrOutput)
}

func (o ModelPackageOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ModelPackage) pulumi.StringOutput { return v.CreationTime }).(pulumi.StringOutput)
}

func (o ModelPackageOutput) CustomerMetadataProperties() ModelPackageCustomerMetadataPropertiesPtrOutput {
	return o.ApplyT(func(v *ModelPackage) ModelPackageCustomerMetadataPropertiesPtrOutput {
		return v.CustomerMetadataProperties
	}).(ModelPackageCustomerMetadataPropertiesPtrOutput)
}

func (o ModelPackageOutput) Domain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ModelPackage) pulumi.StringPtrOutput { return v.Domain }).(pulumi.StringPtrOutput)
}

func (o ModelPackageOutput) DriftCheckBaselines() ModelPackageDriftCheckBaselinesPtrOutput {
	return o.ApplyT(func(v *ModelPackage) ModelPackageDriftCheckBaselinesPtrOutput { return v.DriftCheckBaselines }).(ModelPackageDriftCheckBaselinesPtrOutput)
}

func (o ModelPackageOutput) InferenceSpecification() ModelPackageInferenceSpecificationPtrOutput {
	return o.ApplyT(func(v *ModelPackage) ModelPackageInferenceSpecificationPtrOutput { return v.InferenceSpecification }).(ModelPackageInferenceSpecificationPtrOutput)
}

func (o ModelPackageOutput) LastModifiedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ModelPackage) pulumi.StringPtrOutput { return v.LastModifiedTime }).(pulumi.StringPtrOutput)
}

func (o ModelPackageOutput) MetadataProperties() ModelPackageMetadataPropertiesPtrOutput {
	return o.ApplyT(func(v *ModelPackage) ModelPackageMetadataPropertiesPtrOutput { return v.MetadataProperties }).(ModelPackageMetadataPropertiesPtrOutput)
}

func (o ModelPackageOutput) ModelApprovalStatus() ModelPackageModelApprovalStatusPtrOutput {
	return o.ApplyT(func(v *ModelPackage) ModelPackageModelApprovalStatusPtrOutput { return v.ModelApprovalStatus }).(ModelPackageModelApprovalStatusPtrOutput)
}

func (o ModelPackageOutput) ModelMetrics() ModelPackageModelMetricsPtrOutput {
	return o.ApplyT(func(v *ModelPackage) ModelPackageModelMetricsPtrOutput { return v.ModelMetrics }).(ModelPackageModelMetricsPtrOutput)
}

func (o ModelPackageOutput) ModelPackageArn() pulumi.StringOutput {
	return o.ApplyT(func(v *ModelPackage) pulumi.StringOutput { return v.ModelPackageArn }).(pulumi.StringOutput)
}

func (o ModelPackageOutput) ModelPackageDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ModelPackage) pulumi.StringPtrOutput { return v.ModelPackageDescription }).(pulumi.StringPtrOutput)
}

func (o ModelPackageOutput) ModelPackageGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ModelPackage) pulumi.StringPtrOutput { return v.ModelPackageGroupName }).(pulumi.StringPtrOutput)
}

func (o ModelPackageOutput) ModelPackageName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ModelPackage) pulumi.StringPtrOutput { return v.ModelPackageName }).(pulumi.StringPtrOutput)
}

func (o ModelPackageOutput) ModelPackageStatus() ModelPackageStatusOutput {
	return o.ApplyT(func(v *ModelPackage) ModelPackageStatusOutput { return v.ModelPackageStatus }).(ModelPackageStatusOutput)
}

func (o ModelPackageOutput) ModelPackageStatusDetails() ModelPackageStatusDetailsPtrOutput {
	return o.ApplyT(func(v *ModelPackage) ModelPackageStatusDetailsPtrOutput { return v.ModelPackageStatusDetails }).(ModelPackageStatusDetailsPtrOutput)
}

func (o ModelPackageOutput) ModelPackageVersion() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ModelPackage) pulumi.IntPtrOutput { return v.ModelPackageVersion }).(pulumi.IntPtrOutput)
}

func (o ModelPackageOutput) SamplePayloadUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ModelPackage) pulumi.StringPtrOutput { return v.SamplePayloadUrl }).(pulumi.StringPtrOutput)
}

func (o ModelPackageOutput) SourceAlgorithmSpecification() ModelPackageSourceAlgorithmSpecificationPtrOutput {
	return o.ApplyT(func(v *ModelPackage) ModelPackageSourceAlgorithmSpecificationPtrOutput {
		return v.SourceAlgorithmSpecification
	}).(ModelPackageSourceAlgorithmSpecificationPtrOutput)
}

// An array of key-value pairs to apply to this resource.
func (o ModelPackageOutput) Tags() ModelPackageTagArrayOutput {
	return o.ApplyT(func(v *ModelPackage) ModelPackageTagArrayOutput { return v.Tags }).(ModelPackageTagArrayOutput)
}

func (o ModelPackageOutput) Task() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ModelPackage) pulumi.StringPtrOutput { return v.Task }).(pulumi.StringPtrOutput)
}

func (o ModelPackageOutput) ValidationSpecification() ModelPackageValidationSpecificationPtrOutput {
	return o.ApplyT(func(v *ModelPackage) ModelPackageValidationSpecificationPtrOutput { return v.ValidationSpecification }).(ModelPackageValidationSpecificationPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ModelPackageInput)(nil)).Elem(), &ModelPackage{})
	pulumi.RegisterOutputType(ModelPackageOutput{})
}

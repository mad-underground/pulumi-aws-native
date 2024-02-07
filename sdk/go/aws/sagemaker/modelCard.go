// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sagemaker

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::SageMaker::ModelCard.
type ModelCard struct {
	pulumi.CustomResourceState

	Content ModelCardContentOutput `pulumi:"content"`
	// Information about the user who created or modified an experiment, trial, trial component, lineage group, project, or model card.
	CreatedBy ModelCardUserContextPtrOutput `pulumi:"createdBy"`
	// The date and time the model card was created.
	CreationTime pulumi.StringOutput `pulumi:"creationTime"`
	// Information about the user who created or modified an experiment, trial, trial component, lineage group, project, or model card.
	LastModifiedBy ModelCardUserContextPtrOutput `pulumi:"lastModifiedBy"`
	// The date and time the model card was last modified.
	LastModifiedTime pulumi.StringOutput `pulumi:"lastModifiedTime"`
	// The Amazon Resource Name (ARN) of the successfully created model card.
	ModelCardArn pulumi.StringOutput `pulumi:"modelCardArn"`
	// The unique name of the model card.
	ModelCardName pulumi.StringOutput `pulumi:"modelCardName"`
	// The processing status of model card deletion. The ModelCardProcessingStatus updates throughout the different deletion steps.
	ModelCardProcessingStatus ModelCardProcessingStatusOutput `pulumi:"modelCardProcessingStatus"`
	// The approval status of the model card within your organization. Different organizations might have different criteria for model card review and approval.
	ModelCardStatus ModelCardStatusOutput `pulumi:"modelCardStatus"`
	// A version of the model card.
	ModelCardVersion pulumi.IntOutput                 `pulumi:"modelCardVersion"`
	SecurityConfig   ModelCardSecurityConfigPtrOutput `pulumi:"securityConfig"`
	// Key-value pairs used to manage metadata for model cards.
	Tags ModelCardTagArrayOutput `pulumi:"tags"`
}

// NewModelCard registers a new resource with the given unique name, arguments, and options.
func NewModelCard(ctx *pulumi.Context,
	name string, args *ModelCardArgs, opts ...pulumi.ResourceOption) (*ModelCard, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Content == nil {
		return nil, errors.New("invalid value for required argument 'Content'")
	}
	if args.ModelCardStatus == nil {
		return nil, errors.New("invalid value for required argument 'ModelCardStatus'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"modelCardName",
		"securityConfig",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ModelCard
	err := ctx.RegisterResource("aws-native:sagemaker:ModelCard", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetModelCard gets an existing ModelCard resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetModelCard(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ModelCardState, opts ...pulumi.ResourceOption) (*ModelCard, error) {
	var resource ModelCard
	err := ctx.ReadResource("aws-native:sagemaker:ModelCard", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ModelCard resources.
type modelCardState struct {
}

type ModelCardState struct {
}

func (ModelCardState) ElementType() reflect.Type {
	return reflect.TypeOf((*modelCardState)(nil)).Elem()
}

type modelCardArgs struct {
	Content ModelCardContent `pulumi:"content"`
	// Information about the user who created or modified an experiment, trial, trial component, lineage group, project, or model card.
	CreatedBy *ModelCardUserContext `pulumi:"createdBy"`
	// Information about the user who created or modified an experiment, trial, trial component, lineage group, project, or model card.
	LastModifiedBy *ModelCardUserContext `pulumi:"lastModifiedBy"`
	// The unique name of the model card.
	ModelCardName *string `pulumi:"modelCardName"`
	// The approval status of the model card within your organization. Different organizations might have different criteria for model card review and approval.
	ModelCardStatus ModelCardStatus          `pulumi:"modelCardStatus"`
	SecurityConfig  *ModelCardSecurityConfig `pulumi:"securityConfig"`
	// Key-value pairs used to manage metadata for model cards.
	Tags []ModelCardTag `pulumi:"tags"`
}

// The set of arguments for constructing a ModelCard resource.
type ModelCardArgs struct {
	Content ModelCardContentInput
	// Information about the user who created or modified an experiment, trial, trial component, lineage group, project, or model card.
	CreatedBy ModelCardUserContextPtrInput
	// Information about the user who created or modified an experiment, trial, trial component, lineage group, project, or model card.
	LastModifiedBy ModelCardUserContextPtrInput
	// The unique name of the model card.
	ModelCardName pulumi.StringPtrInput
	// The approval status of the model card within your organization. Different organizations might have different criteria for model card review and approval.
	ModelCardStatus ModelCardStatusInput
	SecurityConfig  ModelCardSecurityConfigPtrInput
	// Key-value pairs used to manage metadata for model cards.
	Tags ModelCardTagArrayInput
}

func (ModelCardArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*modelCardArgs)(nil)).Elem()
}

type ModelCardInput interface {
	pulumi.Input

	ToModelCardOutput() ModelCardOutput
	ToModelCardOutputWithContext(ctx context.Context) ModelCardOutput
}

func (*ModelCard) ElementType() reflect.Type {
	return reflect.TypeOf((**ModelCard)(nil)).Elem()
}

func (i *ModelCard) ToModelCardOutput() ModelCardOutput {
	return i.ToModelCardOutputWithContext(context.Background())
}

func (i *ModelCard) ToModelCardOutputWithContext(ctx context.Context) ModelCardOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModelCardOutput)
}

type ModelCardOutput struct{ *pulumi.OutputState }

func (ModelCardOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ModelCard)(nil)).Elem()
}

func (o ModelCardOutput) ToModelCardOutput() ModelCardOutput {
	return o
}

func (o ModelCardOutput) ToModelCardOutputWithContext(ctx context.Context) ModelCardOutput {
	return o
}

func (o ModelCardOutput) Content() ModelCardContentOutput {
	return o.ApplyT(func(v *ModelCard) ModelCardContentOutput { return v.Content }).(ModelCardContentOutput)
}

// Information about the user who created or modified an experiment, trial, trial component, lineage group, project, or model card.
func (o ModelCardOutput) CreatedBy() ModelCardUserContextPtrOutput {
	return o.ApplyT(func(v *ModelCard) ModelCardUserContextPtrOutput { return v.CreatedBy }).(ModelCardUserContextPtrOutput)
}

// The date and time the model card was created.
func (o ModelCardOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ModelCard) pulumi.StringOutput { return v.CreationTime }).(pulumi.StringOutput)
}

// Information about the user who created or modified an experiment, trial, trial component, lineage group, project, or model card.
func (o ModelCardOutput) LastModifiedBy() ModelCardUserContextPtrOutput {
	return o.ApplyT(func(v *ModelCard) ModelCardUserContextPtrOutput { return v.LastModifiedBy }).(ModelCardUserContextPtrOutput)
}

// The date and time the model card was last modified.
func (o ModelCardOutput) LastModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ModelCard) pulumi.StringOutput { return v.LastModifiedTime }).(pulumi.StringOutput)
}

// The Amazon Resource Name (ARN) of the successfully created model card.
func (o ModelCardOutput) ModelCardArn() pulumi.StringOutput {
	return o.ApplyT(func(v *ModelCard) pulumi.StringOutput { return v.ModelCardArn }).(pulumi.StringOutput)
}

// The unique name of the model card.
func (o ModelCardOutput) ModelCardName() pulumi.StringOutput {
	return o.ApplyT(func(v *ModelCard) pulumi.StringOutput { return v.ModelCardName }).(pulumi.StringOutput)
}

// The processing status of model card deletion. The ModelCardProcessingStatus updates throughout the different deletion steps.
func (o ModelCardOutput) ModelCardProcessingStatus() ModelCardProcessingStatusOutput {
	return o.ApplyT(func(v *ModelCard) ModelCardProcessingStatusOutput { return v.ModelCardProcessingStatus }).(ModelCardProcessingStatusOutput)
}

// The approval status of the model card within your organization. Different organizations might have different criteria for model card review and approval.
func (o ModelCardOutput) ModelCardStatus() ModelCardStatusOutput {
	return o.ApplyT(func(v *ModelCard) ModelCardStatusOutput { return v.ModelCardStatus }).(ModelCardStatusOutput)
}

// A version of the model card.
func (o ModelCardOutput) ModelCardVersion() pulumi.IntOutput {
	return o.ApplyT(func(v *ModelCard) pulumi.IntOutput { return v.ModelCardVersion }).(pulumi.IntOutput)
}

func (o ModelCardOutput) SecurityConfig() ModelCardSecurityConfigPtrOutput {
	return o.ApplyT(func(v *ModelCard) ModelCardSecurityConfigPtrOutput { return v.SecurityConfig }).(ModelCardSecurityConfigPtrOutput)
}

// Key-value pairs used to manage metadata for model cards.
func (o ModelCardOutput) Tags() ModelCardTagArrayOutput {
	return o.ApplyT(func(v *ModelCard) ModelCardTagArrayOutput { return v.Tags }).(ModelCardTagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ModelCardInput)(nil)).Elem(), &ModelCard{})
	pulumi.RegisterOutputType(ModelCardOutput{})
}

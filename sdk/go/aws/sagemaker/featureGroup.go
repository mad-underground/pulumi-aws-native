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

// Resource Type definition for AWS::SageMaker::FeatureGroup
type FeatureGroup struct {
	pulumi.CustomResourceState

	// A timestamp of FeatureGroup creation time.
	CreationTime pulumi.StringOutput `pulumi:"creationTime"`
	// Description about the FeatureGroup.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The Event Time Feature Name.
	EventTimeFeatureName pulumi.StringOutput `pulumi:"eventTimeFeatureName"`
	// An Array of Feature Definition
	FeatureDefinitions FeatureGroupFeatureDefinitionArrayOutput `pulumi:"featureDefinitions"`
	// The Name of the FeatureGroup.
	FeatureGroupName pulumi.StringOutput `pulumi:"featureGroupName"`
	// The status of the feature group.
	FeatureGroupStatus pulumi.StringOutput                   `pulumi:"featureGroupStatus"`
	OfflineStoreConfig OfflineStoreConfigPropertiesPtrOutput `pulumi:"offlineStoreConfig"`
	OnlineStoreConfig  OnlineStoreConfigPropertiesPtrOutput  `pulumi:"onlineStoreConfig"`
	// The Record Identifier Feature Name.
	RecordIdentifierFeatureName pulumi.StringOutput `pulumi:"recordIdentifierFeatureName"`
	// Role Arn
	RoleArn pulumi.StringPtrOutput `pulumi:"roleArn"`
	// An array of key-value pair to apply to this resource.
	Tags FeatureGroupTagArrayOutput `pulumi:"tags"`
}

// NewFeatureGroup registers a new resource with the given unique name, arguments, and options.
func NewFeatureGroup(ctx *pulumi.Context,
	name string, args *FeatureGroupArgs, opts ...pulumi.ResourceOption) (*FeatureGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EventTimeFeatureName == nil {
		return nil, errors.New("invalid value for required argument 'EventTimeFeatureName'")
	}
	if args.FeatureDefinitions == nil {
		return nil, errors.New("invalid value for required argument 'FeatureDefinitions'")
	}
	if args.RecordIdentifierFeatureName == nil {
		return nil, errors.New("invalid value for required argument 'RecordIdentifierFeatureName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FeatureGroup
	err := ctx.RegisterResource("aws-native:sagemaker:FeatureGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFeatureGroup gets an existing FeatureGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFeatureGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FeatureGroupState, opts ...pulumi.ResourceOption) (*FeatureGroup, error) {
	var resource FeatureGroup
	err := ctx.ReadResource("aws-native:sagemaker:FeatureGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FeatureGroup resources.
type featureGroupState struct {
}

type FeatureGroupState struct {
}

func (FeatureGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*featureGroupState)(nil)).Elem()
}

type featureGroupArgs struct {
	// Description about the FeatureGroup.
	Description *string `pulumi:"description"`
	// The Event Time Feature Name.
	EventTimeFeatureName string `pulumi:"eventTimeFeatureName"`
	// An Array of Feature Definition
	FeatureDefinitions []FeatureGroupFeatureDefinition `pulumi:"featureDefinitions"`
	// The Name of the FeatureGroup.
	FeatureGroupName   *string                       `pulumi:"featureGroupName"`
	OfflineStoreConfig *OfflineStoreConfigProperties `pulumi:"offlineStoreConfig"`
	OnlineStoreConfig  *OnlineStoreConfigProperties  `pulumi:"onlineStoreConfig"`
	// The Record Identifier Feature Name.
	RecordIdentifierFeatureName string `pulumi:"recordIdentifierFeatureName"`
	// Role Arn
	RoleArn *string `pulumi:"roleArn"`
	// An array of key-value pair to apply to this resource.
	Tags []FeatureGroupTag `pulumi:"tags"`
}

// The set of arguments for constructing a FeatureGroup resource.
type FeatureGroupArgs struct {
	// Description about the FeatureGroup.
	Description pulumi.StringPtrInput
	// The Event Time Feature Name.
	EventTimeFeatureName pulumi.StringInput
	// An Array of Feature Definition
	FeatureDefinitions FeatureGroupFeatureDefinitionArrayInput
	// The Name of the FeatureGroup.
	FeatureGroupName   pulumi.StringPtrInput
	OfflineStoreConfig OfflineStoreConfigPropertiesPtrInput
	OnlineStoreConfig  OnlineStoreConfigPropertiesPtrInput
	// The Record Identifier Feature Name.
	RecordIdentifierFeatureName pulumi.StringInput
	// Role Arn
	RoleArn pulumi.StringPtrInput
	// An array of key-value pair to apply to this resource.
	Tags FeatureGroupTagArrayInput
}

func (FeatureGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*featureGroupArgs)(nil)).Elem()
}

type FeatureGroupInput interface {
	pulumi.Input

	ToFeatureGroupOutput() FeatureGroupOutput
	ToFeatureGroupOutputWithContext(ctx context.Context) FeatureGroupOutput
}

func (*FeatureGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**FeatureGroup)(nil)).Elem()
}

func (i *FeatureGroup) ToFeatureGroupOutput() FeatureGroupOutput {
	return i.ToFeatureGroupOutputWithContext(context.Background())
}

func (i *FeatureGroup) ToFeatureGroupOutputWithContext(ctx context.Context) FeatureGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FeatureGroupOutput)
}

type FeatureGroupOutput struct{ *pulumi.OutputState }

func (FeatureGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FeatureGroup)(nil)).Elem()
}

func (o FeatureGroupOutput) ToFeatureGroupOutput() FeatureGroupOutput {
	return o
}

func (o FeatureGroupOutput) ToFeatureGroupOutputWithContext(ctx context.Context) FeatureGroupOutput {
	return o
}

// A timestamp of FeatureGroup creation time.
func (o FeatureGroupOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *FeatureGroup) pulumi.StringOutput { return v.CreationTime }).(pulumi.StringOutput)
}

// Description about the FeatureGroup.
func (o FeatureGroupOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FeatureGroup) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The Event Time Feature Name.
func (o FeatureGroupOutput) EventTimeFeatureName() pulumi.StringOutput {
	return o.ApplyT(func(v *FeatureGroup) pulumi.StringOutput { return v.EventTimeFeatureName }).(pulumi.StringOutput)
}

// An Array of Feature Definition
func (o FeatureGroupOutput) FeatureDefinitions() FeatureGroupFeatureDefinitionArrayOutput {
	return o.ApplyT(func(v *FeatureGroup) FeatureGroupFeatureDefinitionArrayOutput { return v.FeatureDefinitions }).(FeatureGroupFeatureDefinitionArrayOutput)
}

// The Name of the FeatureGroup.
func (o FeatureGroupOutput) FeatureGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v *FeatureGroup) pulumi.StringOutput { return v.FeatureGroupName }).(pulumi.StringOutput)
}

// The status of the feature group.
func (o FeatureGroupOutput) FeatureGroupStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *FeatureGroup) pulumi.StringOutput { return v.FeatureGroupStatus }).(pulumi.StringOutput)
}

func (o FeatureGroupOutput) OfflineStoreConfig() OfflineStoreConfigPropertiesPtrOutput {
	return o.ApplyT(func(v *FeatureGroup) OfflineStoreConfigPropertiesPtrOutput { return v.OfflineStoreConfig }).(OfflineStoreConfigPropertiesPtrOutput)
}

func (o FeatureGroupOutput) OnlineStoreConfig() OnlineStoreConfigPropertiesPtrOutput {
	return o.ApplyT(func(v *FeatureGroup) OnlineStoreConfigPropertiesPtrOutput { return v.OnlineStoreConfig }).(OnlineStoreConfigPropertiesPtrOutput)
}

// The Record Identifier Feature Name.
func (o FeatureGroupOutput) RecordIdentifierFeatureName() pulumi.StringOutput {
	return o.ApplyT(func(v *FeatureGroup) pulumi.StringOutput { return v.RecordIdentifierFeatureName }).(pulumi.StringOutput)
}

// Role Arn
func (o FeatureGroupOutput) RoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FeatureGroup) pulumi.StringPtrOutput { return v.RoleArn }).(pulumi.StringPtrOutput)
}

// An array of key-value pair to apply to this resource.
func (o FeatureGroupOutput) Tags() FeatureGroupTagArrayOutput {
	return o.ApplyT(func(v *FeatureGroup) FeatureGroupTagArrayOutput { return v.Tags }).(FeatureGroupTagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FeatureGroupInput)(nil)).Elem(), &FeatureGroup{})
	pulumi.RegisterOutputType(FeatureGroupOutput{})
}

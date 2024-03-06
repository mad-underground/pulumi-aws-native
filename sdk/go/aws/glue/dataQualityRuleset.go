// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package glue

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Glue::DataQualityRuleset
//
// Deprecated: DataQualityRuleset is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type DataQualityRuleset struct {
	pulumi.CustomResourceState

	AwsId       pulumi.StringOutput    `pulumi:"awsId"`
	ClientToken pulumi.StringPtrOutput `pulumi:"clientToken"`
	Description pulumi.StringPtrOutput `pulumi:"description"`
	Name        pulumi.StringPtrOutput `pulumi:"name"`
	Ruleset     pulumi.StringPtrOutput `pulumi:"ruleset"`
	// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Glue::DataQualityRuleset` for more information about the expected schema for this property.
	Tags        pulumi.AnyOutput                                  `pulumi:"tags"`
	TargetTable DataQualityRulesetDataQualityTargetTablePtrOutput `pulumi:"targetTable"`
}

// NewDataQualityRuleset registers a new resource with the given unique name, arguments, and options.
func NewDataQualityRuleset(ctx *pulumi.Context,
	name string, args *DataQualityRulesetArgs, opts ...pulumi.ResourceOption) (*DataQualityRuleset, error) {
	if args == nil {
		args = &DataQualityRulesetArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DataQualityRuleset
	err := ctx.RegisterResource("aws-native:glue:DataQualityRuleset", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataQualityRuleset gets an existing DataQualityRuleset resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataQualityRuleset(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataQualityRulesetState, opts ...pulumi.ResourceOption) (*DataQualityRuleset, error) {
	var resource DataQualityRuleset
	err := ctx.ReadResource("aws-native:glue:DataQualityRuleset", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataQualityRuleset resources.
type dataQualityRulesetState struct {
}

type DataQualityRulesetState struct {
}

func (DataQualityRulesetState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataQualityRulesetState)(nil)).Elem()
}

type dataQualityRulesetArgs struct {
	ClientToken *string `pulumi:"clientToken"`
	Description *string `pulumi:"description"`
	Name        *string `pulumi:"name"`
	Ruleset     *string `pulumi:"ruleset"`
	// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Glue::DataQualityRuleset` for more information about the expected schema for this property.
	Tags        interface{}                               `pulumi:"tags"`
	TargetTable *DataQualityRulesetDataQualityTargetTable `pulumi:"targetTable"`
}

// The set of arguments for constructing a DataQualityRuleset resource.
type DataQualityRulesetArgs struct {
	ClientToken pulumi.StringPtrInput
	Description pulumi.StringPtrInput
	Name        pulumi.StringPtrInput
	Ruleset     pulumi.StringPtrInput
	// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Glue::DataQualityRuleset` for more information about the expected schema for this property.
	Tags        pulumi.Input
	TargetTable DataQualityRulesetDataQualityTargetTablePtrInput
}

func (DataQualityRulesetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataQualityRulesetArgs)(nil)).Elem()
}

type DataQualityRulesetInput interface {
	pulumi.Input

	ToDataQualityRulesetOutput() DataQualityRulesetOutput
	ToDataQualityRulesetOutputWithContext(ctx context.Context) DataQualityRulesetOutput
}

func (*DataQualityRuleset) ElementType() reflect.Type {
	return reflect.TypeOf((**DataQualityRuleset)(nil)).Elem()
}

func (i *DataQualityRuleset) ToDataQualityRulesetOutput() DataQualityRulesetOutput {
	return i.ToDataQualityRulesetOutputWithContext(context.Background())
}

func (i *DataQualityRuleset) ToDataQualityRulesetOutputWithContext(ctx context.Context) DataQualityRulesetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataQualityRulesetOutput)
}

type DataQualityRulesetOutput struct{ *pulumi.OutputState }

func (DataQualityRulesetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataQualityRuleset)(nil)).Elem()
}

func (o DataQualityRulesetOutput) ToDataQualityRulesetOutput() DataQualityRulesetOutput {
	return o
}

func (o DataQualityRulesetOutput) ToDataQualityRulesetOutputWithContext(ctx context.Context) DataQualityRulesetOutput {
	return o
}

func (o DataQualityRulesetOutput) AwsId() pulumi.StringOutput {
	return o.ApplyT(func(v *DataQualityRuleset) pulumi.StringOutput { return v.AwsId }).(pulumi.StringOutput)
}

func (o DataQualityRulesetOutput) ClientToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataQualityRuleset) pulumi.StringPtrOutput { return v.ClientToken }).(pulumi.StringPtrOutput)
}

func (o DataQualityRulesetOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataQualityRuleset) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o DataQualityRulesetOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataQualityRuleset) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o DataQualityRulesetOutput) Ruleset() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataQualityRuleset) pulumi.StringPtrOutput { return v.Ruleset }).(pulumi.StringPtrOutput)
}

// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Glue::DataQualityRuleset` for more information about the expected schema for this property.
func (o DataQualityRulesetOutput) Tags() pulumi.AnyOutput {
	return o.ApplyT(func(v *DataQualityRuleset) pulumi.AnyOutput { return v.Tags }).(pulumi.AnyOutput)
}

func (o DataQualityRulesetOutput) TargetTable() DataQualityRulesetDataQualityTargetTablePtrOutput {
	return o.ApplyT(func(v *DataQualityRuleset) DataQualityRulesetDataQualityTargetTablePtrOutput { return v.TargetTable }).(DataQualityRulesetDataQualityTargetTablePtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DataQualityRulesetInput)(nil)).Elem(), &DataQualityRuleset{})
	pulumi.RegisterOutputType(DataQualityRulesetOutput{})
}

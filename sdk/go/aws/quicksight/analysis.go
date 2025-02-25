// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package quicksight

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of the AWS::QuickSight::Analysis Resource Type.
type Analysis struct {
	pulumi.CustomResourceState

	AnalysisId         pulumi.StringOutput                   `pulumi:"analysisId"`
	Arn                pulumi.StringOutput                   `pulumi:"arn"`
	AwsAccountId       pulumi.StringOutput                   `pulumi:"awsAccountId"`
	CreatedTime        pulumi.StringOutput                   `pulumi:"createdTime"`
	DataSetArns        pulumi.StringArrayOutput              `pulumi:"dataSetArns"`
	Definition         AnalysisDefinitionPtrOutput           `pulumi:"definition"`
	Errors             AnalysisErrorArrayOutput              `pulumi:"errors"`
	LastUpdatedTime    pulumi.StringOutput                   `pulumi:"lastUpdatedTime"`
	Name               pulumi.StringOutput                   `pulumi:"name"`
	Parameters         AnalysisParametersPtrOutput           `pulumi:"parameters"`
	Permissions        AnalysisResourcePermissionArrayOutput `pulumi:"permissions"`
	Sheets             AnalysisSheetArrayOutput              `pulumi:"sheets"`
	SourceEntity       AnalysisSourceEntityPtrOutput         `pulumi:"sourceEntity"`
	Status             AnalysisResourceStatusPtrOutput       `pulumi:"status"`
	Tags               aws.TagArrayOutput                    `pulumi:"tags"`
	ThemeArn           pulumi.StringPtrOutput                `pulumi:"themeArn"`
	ValidationStrategy AnalysisValidationStrategyPtrOutput   `pulumi:"validationStrategy"`
}

// NewAnalysis registers a new resource with the given unique name, arguments, and options.
func NewAnalysis(ctx *pulumi.Context,
	name string, args *AnalysisArgs, opts ...pulumi.ResourceOption) (*Analysis, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AnalysisId == nil {
		return nil, errors.New("invalid value for required argument 'AnalysisId'")
	}
	if args.AwsAccountId == nil {
		return nil, errors.New("invalid value for required argument 'AwsAccountId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"analysisId",
		"awsAccountId",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Analysis
	err := ctx.RegisterResource("aws-native:quicksight:Analysis", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAnalysis gets an existing Analysis resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAnalysis(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AnalysisState, opts ...pulumi.ResourceOption) (*Analysis, error) {
	var resource Analysis
	err := ctx.ReadResource("aws-native:quicksight:Analysis", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Analysis resources.
type analysisState struct {
}

type AnalysisState struct {
}

func (AnalysisState) ElementType() reflect.Type {
	return reflect.TypeOf((*analysisState)(nil)).Elem()
}

type analysisArgs struct {
	AnalysisId         string                       `pulumi:"analysisId"`
	AwsAccountId       string                       `pulumi:"awsAccountId"`
	Definition         *AnalysisDefinition          `pulumi:"definition"`
	Errors             []AnalysisError              `pulumi:"errors"`
	Name               *string                      `pulumi:"name"`
	Parameters         *AnalysisParameters          `pulumi:"parameters"`
	Permissions        []AnalysisResourcePermission `pulumi:"permissions"`
	Sheets             []AnalysisSheet              `pulumi:"sheets"`
	SourceEntity       *AnalysisSourceEntity        `pulumi:"sourceEntity"`
	Status             *AnalysisResourceStatus      `pulumi:"status"`
	Tags               []aws.Tag                    `pulumi:"tags"`
	ThemeArn           *string                      `pulumi:"themeArn"`
	ValidationStrategy *AnalysisValidationStrategy  `pulumi:"validationStrategy"`
}

// The set of arguments for constructing a Analysis resource.
type AnalysisArgs struct {
	AnalysisId         pulumi.StringInput
	AwsAccountId       pulumi.StringInput
	Definition         AnalysisDefinitionPtrInput
	Errors             AnalysisErrorArrayInput
	Name               pulumi.StringPtrInput
	Parameters         AnalysisParametersPtrInput
	Permissions        AnalysisResourcePermissionArrayInput
	Sheets             AnalysisSheetArrayInput
	SourceEntity       AnalysisSourceEntityPtrInput
	Status             AnalysisResourceStatusPtrInput
	Tags               aws.TagArrayInput
	ThemeArn           pulumi.StringPtrInput
	ValidationStrategy AnalysisValidationStrategyPtrInput
}

func (AnalysisArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*analysisArgs)(nil)).Elem()
}

type AnalysisInput interface {
	pulumi.Input

	ToAnalysisOutput() AnalysisOutput
	ToAnalysisOutputWithContext(ctx context.Context) AnalysisOutput
}

func (*Analysis) ElementType() reflect.Type {
	return reflect.TypeOf((**Analysis)(nil)).Elem()
}

func (i *Analysis) ToAnalysisOutput() AnalysisOutput {
	return i.ToAnalysisOutputWithContext(context.Background())
}

func (i *Analysis) ToAnalysisOutputWithContext(ctx context.Context) AnalysisOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AnalysisOutput)
}

type AnalysisOutput struct{ *pulumi.OutputState }

func (AnalysisOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Analysis)(nil)).Elem()
}

func (o AnalysisOutput) ToAnalysisOutput() AnalysisOutput {
	return o
}

func (o AnalysisOutput) ToAnalysisOutputWithContext(ctx context.Context) AnalysisOutput {
	return o
}

func (o AnalysisOutput) AnalysisId() pulumi.StringOutput {
	return o.ApplyT(func(v *Analysis) pulumi.StringOutput { return v.AnalysisId }).(pulumi.StringOutput)
}

func (o AnalysisOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Analysis) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

func (o AnalysisOutput) AwsAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *Analysis) pulumi.StringOutput { return v.AwsAccountId }).(pulumi.StringOutput)
}

func (o AnalysisOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Analysis) pulumi.StringOutput { return v.CreatedTime }).(pulumi.StringOutput)
}

func (o AnalysisOutput) DataSetArns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Analysis) pulumi.StringArrayOutput { return v.DataSetArns }).(pulumi.StringArrayOutput)
}

func (o AnalysisOutput) Definition() AnalysisDefinitionPtrOutput {
	return o.ApplyT(func(v *Analysis) AnalysisDefinitionPtrOutput { return v.Definition }).(AnalysisDefinitionPtrOutput)
}

func (o AnalysisOutput) Errors() AnalysisErrorArrayOutput {
	return o.ApplyT(func(v *Analysis) AnalysisErrorArrayOutput { return v.Errors }).(AnalysisErrorArrayOutput)
}

func (o AnalysisOutput) LastUpdatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Analysis) pulumi.StringOutput { return v.LastUpdatedTime }).(pulumi.StringOutput)
}

func (o AnalysisOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Analysis) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AnalysisOutput) Parameters() AnalysisParametersPtrOutput {
	return o.ApplyT(func(v *Analysis) AnalysisParametersPtrOutput { return v.Parameters }).(AnalysisParametersPtrOutput)
}

func (o AnalysisOutput) Permissions() AnalysisResourcePermissionArrayOutput {
	return o.ApplyT(func(v *Analysis) AnalysisResourcePermissionArrayOutput { return v.Permissions }).(AnalysisResourcePermissionArrayOutput)
}

func (o AnalysisOutput) Sheets() AnalysisSheetArrayOutput {
	return o.ApplyT(func(v *Analysis) AnalysisSheetArrayOutput { return v.Sheets }).(AnalysisSheetArrayOutput)
}

func (o AnalysisOutput) SourceEntity() AnalysisSourceEntityPtrOutput {
	return o.ApplyT(func(v *Analysis) AnalysisSourceEntityPtrOutput { return v.SourceEntity }).(AnalysisSourceEntityPtrOutput)
}

func (o AnalysisOutput) Status() AnalysisResourceStatusPtrOutput {
	return o.ApplyT(func(v *Analysis) AnalysisResourceStatusPtrOutput { return v.Status }).(AnalysisResourceStatusPtrOutput)
}

func (o AnalysisOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v *Analysis) aws.TagArrayOutput { return v.Tags }).(aws.TagArrayOutput)
}

func (o AnalysisOutput) ThemeArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Analysis) pulumi.StringPtrOutput { return v.ThemeArn }).(pulumi.StringPtrOutput)
}

func (o AnalysisOutput) ValidationStrategy() AnalysisValidationStrategyPtrOutput {
	return o.ApplyT(func(v *Analysis) AnalysisValidationStrategyPtrOutput { return v.ValidationStrategy }).(AnalysisValidationStrategyPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AnalysisInput)(nil)).Elem(), &Analysis{})
	pulumi.RegisterOutputType(AnalysisOutput{})
}

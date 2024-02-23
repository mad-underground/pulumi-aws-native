// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package fis

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for AWS::FIS::ExperimentTemplate
type ExperimentTemplate struct {
	pulumi.CustomResourceState

	Actions           ExperimentTemplateActionMapOutput            `pulumi:"actions"`
	Description       pulumi.StringOutput                          `pulumi:"description"`
	ExperimentOptions ExperimentTemplateExperimentOptionsPtrOutput `pulumi:"experimentOptions"`
	LogConfiguration  ExperimentTemplateLogConfigurationPtrOutput  `pulumi:"logConfiguration"`
	RoleArn           pulumi.StringOutput                          `pulumi:"roleArn"`
	StopConditions    ExperimentTemplateStopConditionArrayOutput   `pulumi:"stopConditions"`
	Tags              pulumi.StringMapOutput                       `pulumi:"tags"`
	Targets           ExperimentTemplateTargetMapOutput            `pulumi:"targets"`
}

// NewExperimentTemplate registers a new resource with the given unique name, arguments, and options.
func NewExperimentTemplate(ctx *pulumi.Context,
	name string, args *ExperimentTemplateArgs, opts ...pulumi.ResourceOption) (*ExperimentTemplate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	if args.RoleArn == nil {
		return nil, errors.New("invalid value for required argument 'RoleArn'")
	}
	if args.StopConditions == nil {
		return nil, errors.New("invalid value for required argument 'StopConditions'")
	}
	if args.Tags == nil {
		return nil, errors.New("invalid value for required argument 'Tags'")
	}
	if args.Targets == nil {
		return nil, errors.New("invalid value for required argument 'Targets'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"tags.*",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ExperimentTemplate
	err := ctx.RegisterResource("aws-native:fis:ExperimentTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExperimentTemplate gets an existing ExperimentTemplate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExperimentTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExperimentTemplateState, opts ...pulumi.ResourceOption) (*ExperimentTemplate, error) {
	var resource ExperimentTemplate
	err := ctx.ReadResource("aws-native:fis:ExperimentTemplate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ExperimentTemplate resources.
type experimentTemplateState struct {
}

type ExperimentTemplateState struct {
}

func (ExperimentTemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*experimentTemplateState)(nil)).Elem()
}

type experimentTemplateArgs struct {
	Actions           map[string]ExperimentTemplateAction  `pulumi:"actions"`
	Description       string                               `pulumi:"description"`
	ExperimentOptions *ExperimentTemplateExperimentOptions `pulumi:"experimentOptions"`
	LogConfiguration  *ExperimentTemplateLogConfiguration  `pulumi:"logConfiguration"`
	RoleArn           string                               `pulumi:"roleArn"`
	StopConditions    []ExperimentTemplateStopCondition    `pulumi:"stopConditions"`
	Tags              map[string]string                    `pulumi:"tags"`
	Targets           map[string]ExperimentTemplateTarget  `pulumi:"targets"`
}

// The set of arguments for constructing a ExperimentTemplate resource.
type ExperimentTemplateArgs struct {
	Actions           ExperimentTemplateActionMapInput
	Description       pulumi.StringInput
	ExperimentOptions ExperimentTemplateExperimentOptionsPtrInput
	LogConfiguration  ExperimentTemplateLogConfigurationPtrInput
	RoleArn           pulumi.StringInput
	StopConditions    ExperimentTemplateStopConditionArrayInput
	Tags              pulumi.StringMapInput
	Targets           ExperimentTemplateTargetMapInput
}

func (ExperimentTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*experimentTemplateArgs)(nil)).Elem()
}

type ExperimentTemplateInput interface {
	pulumi.Input

	ToExperimentTemplateOutput() ExperimentTemplateOutput
	ToExperimentTemplateOutputWithContext(ctx context.Context) ExperimentTemplateOutput
}

func (*ExperimentTemplate) ElementType() reflect.Type {
	return reflect.TypeOf((**ExperimentTemplate)(nil)).Elem()
}

func (i *ExperimentTemplate) ToExperimentTemplateOutput() ExperimentTemplateOutput {
	return i.ToExperimentTemplateOutputWithContext(context.Background())
}

func (i *ExperimentTemplate) ToExperimentTemplateOutputWithContext(ctx context.Context) ExperimentTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExperimentTemplateOutput)
}

type ExperimentTemplateOutput struct{ *pulumi.OutputState }

func (ExperimentTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExperimentTemplate)(nil)).Elem()
}

func (o ExperimentTemplateOutput) ToExperimentTemplateOutput() ExperimentTemplateOutput {
	return o
}

func (o ExperimentTemplateOutput) ToExperimentTemplateOutputWithContext(ctx context.Context) ExperimentTemplateOutput {
	return o
}

func (o ExperimentTemplateOutput) Actions() ExperimentTemplateActionMapOutput {
	return o.ApplyT(func(v *ExperimentTemplate) ExperimentTemplateActionMapOutput { return v.Actions }).(ExperimentTemplateActionMapOutput)
}

func (o ExperimentTemplateOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *ExperimentTemplate) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

func (o ExperimentTemplateOutput) ExperimentOptions() ExperimentTemplateExperimentOptionsPtrOutput {
	return o.ApplyT(func(v *ExperimentTemplate) ExperimentTemplateExperimentOptionsPtrOutput { return v.ExperimentOptions }).(ExperimentTemplateExperimentOptionsPtrOutput)
}

func (o ExperimentTemplateOutput) LogConfiguration() ExperimentTemplateLogConfigurationPtrOutput {
	return o.ApplyT(func(v *ExperimentTemplate) ExperimentTemplateLogConfigurationPtrOutput { return v.LogConfiguration }).(ExperimentTemplateLogConfigurationPtrOutput)
}

func (o ExperimentTemplateOutput) RoleArn() pulumi.StringOutput {
	return o.ApplyT(func(v *ExperimentTemplate) pulumi.StringOutput { return v.RoleArn }).(pulumi.StringOutput)
}

func (o ExperimentTemplateOutput) StopConditions() ExperimentTemplateStopConditionArrayOutput {
	return o.ApplyT(func(v *ExperimentTemplate) ExperimentTemplateStopConditionArrayOutput { return v.StopConditions }).(ExperimentTemplateStopConditionArrayOutput)
}

func (o ExperimentTemplateOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ExperimentTemplate) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ExperimentTemplateOutput) Targets() ExperimentTemplateTargetMapOutput {
	return o.ApplyT(func(v *ExperimentTemplate) ExperimentTemplateTargetMapOutput { return v.Targets }).(ExperimentTemplateTargetMapOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ExperimentTemplateInput)(nil)).Elem(), &ExperimentTemplate{})
	pulumi.RegisterOutputType(ExperimentTemplateOutput{})
}

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

// Resource Type definition for AWS::SageMaker::DataQualityJobDefinition
type DataQualityJobDefinition struct {
	pulumi.CustomResourceState

	// The time at which the job definition was created.
	CreationTime                pulumi.StringOutput                                        `pulumi:"creationTime"`
	DataQualityAppSpecification DataQualityJobDefinitionDataQualityAppSpecificationOutput  `pulumi:"dataQualityAppSpecification"`
	DataQualityBaselineConfig   DataQualityJobDefinitionDataQualityBaselineConfigPtrOutput `pulumi:"dataQualityBaselineConfig"`
	DataQualityJobInput         DataQualityJobDefinitionDataQualityJobInputOutput          `pulumi:"dataQualityJobInput"`
	DataQualityJobOutputConfig  DataQualityJobDefinitionMonitoringOutputConfigOutput       `pulumi:"dataQualityJobOutputConfig"`
	EndpointName                pulumi.StringPtrOutput                                     `pulumi:"endpointName"`
	// The Amazon Resource Name (ARN) of job definition.
	JobDefinitionArn  pulumi.StringOutput                               `pulumi:"jobDefinitionArn"`
	JobDefinitionName pulumi.StringPtrOutput                            `pulumi:"jobDefinitionName"`
	JobResources      DataQualityJobDefinitionMonitoringResourcesOutput `pulumi:"jobResources"`
	NetworkConfig     DataQualityJobDefinitionNetworkConfigPtrOutput    `pulumi:"networkConfig"`
	// The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker can assume to perform tasks on your behalf.
	RoleArn           pulumi.StringOutput                                `pulumi:"roleArn"`
	StoppingCondition DataQualityJobDefinitionStoppingConditionPtrOutput `pulumi:"stoppingCondition"`
	// An array of key-value pairs to apply to this resource.
	Tags DataQualityJobDefinitionTagArrayOutput `pulumi:"tags"`
}

// NewDataQualityJobDefinition registers a new resource with the given unique name, arguments, and options.
func NewDataQualityJobDefinition(ctx *pulumi.Context,
	name string, args *DataQualityJobDefinitionArgs, opts ...pulumi.ResourceOption) (*DataQualityJobDefinition, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataQualityAppSpecification == nil {
		return nil, errors.New("invalid value for required argument 'DataQualityAppSpecification'")
	}
	if args.DataQualityJobInput == nil {
		return nil, errors.New("invalid value for required argument 'DataQualityJobInput'")
	}
	if args.DataQualityJobOutputConfig == nil {
		return nil, errors.New("invalid value for required argument 'DataQualityJobOutputConfig'")
	}
	if args.JobResources == nil {
		return nil, errors.New("invalid value for required argument 'JobResources'")
	}
	if args.RoleArn == nil {
		return nil, errors.New("invalid value for required argument 'RoleArn'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"dataQualityAppSpecification",
		"dataQualityBaselineConfig",
		"dataQualityJobInput",
		"dataQualityJobOutputConfig",
		"endpointName",
		"jobDefinitionName",
		"jobResources",
		"networkConfig",
		"roleArn",
		"stoppingCondition",
		"tags[*]",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DataQualityJobDefinition
	err := ctx.RegisterResource("aws-native:sagemaker:DataQualityJobDefinition", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataQualityJobDefinition gets an existing DataQualityJobDefinition resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataQualityJobDefinition(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataQualityJobDefinitionState, opts ...pulumi.ResourceOption) (*DataQualityJobDefinition, error) {
	var resource DataQualityJobDefinition
	err := ctx.ReadResource("aws-native:sagemaker:DataQualityJobDefinition", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataQualityJobDefinition resources.
type dataQualityJobDefinitionState struct {
}

type DataQualityJobDefinitionState struct {
}

func (DataQualityJobDefinitionState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataQualityJobDefinitionState)(nil)).Elem()
}

type dataQualityJobDefinitionArgs struct {
	DataQualityAppSpecification DataQualityJobDefinitionDataQualityAppSpecification `pulumi:"dataQualityAppSpecification"`
	DataQualityBaselineConfig   *DataQualityJobDefinitionDataQualityBaselineConfig  `pulumi:"dataQualityBaselineConfig"`
	DataQualityJobInput         DataQualityJobDefinitionDataQualityJobInput         `pulumi:"dataQualityJobInput"`
	DataQualityJobOutputConfig  DataQualityJobDefinitionMonitoringOutputConfig      `pulumi:"dataQualityJobOutputConfig"`
	EndpointName                *string                                             `pulumi:"endpointName"`
	JobDefinitionName           *string                                             `pulumi:"jobDefinitionName"`
	JobResources                DataQualityJobDefinitionMonitoringResources         `pulumi:"jobResources"`
	NetworkConfig               *DataQualityJobDefinitionNetworkConfig              `pulumi:"networkConfig"`
	// The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker can assume to perform tasks on your behalf.
	RoleArn           string                                     `pulumi:"roleArn"`
	StoppingCondition *DataQualityJobDefinitionStoppingCondition `pulumi:"stoppingCondition"`
	// An array of key-value pairs to apply to this resource.
	Tags []DataQualityJobDefinitionTag `pulumi:"tags"`
}

// The set of arguments for constructing a DataQualityJobDefinition resource.
type DataQualityJobDefinitionArgs struct {
	DataQualityAppSpecification DataQualityJobDefinitionDataQualityAppSpecificationInput
	DataQualityBaselineConfig   DataQualityJobDefinitionDataQualityBaselineConfigPtrInput
	DataQualityJobInput         DataQualityJobDefinitionDataQualityJobInputInput
	DataQualityJobOutputConfig  DataQualityJobDefinitionMonitoringOutputConfigInput
	EndpointName                pulumi.StringPtrInput
	JobDefinitionName           pulumi.StringPtrInput
	JobResources                DataQualityJobDefinitionMonitoringResourcesInput
	NetworkConfig               DataQualityJobDefinitionNetworkConfigPtrInput
	// The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker can assume to perform tasks on your behalf.
	RoleArn           pulumi.StringInput
	StoppingCondition DataQualityJobDefinitionStoppingConditionPtrInput
	// An array of key-value pairs to apply to this resource.
	Tags DataQualityJobDefinitionTagArrayInput
}

func (DataQualityJobDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataQualityJobDefinitionArgs)(nil)).Elem()
}

type DataQualityJobDefinitionInput interface {
	pulumi.Input

	ToDataQualityJobDefinitionOutput() DataQualityJobDefinitionOutput
	ToDataQualityJobDefinitionOutputWithContext(ctx context.Context) DataQualityJobDefinitionOutput
}

func (*DataQualityJobDefinition) ElementType() reflect.Type {
	return reflect.TypeOf((**DataQualityJobDefinition)(nil)).Elem()
}

func (i *DataQualityJobDefinition) ToDataQualityJobDefinitionOutput() DataQualityJobDefinitionOutput {
	return i.ToDataQualityJobDefinitionOutputWithContext(context.Background())
}

func (i *DataQualityJobDefinition) ToDataQualityJobDefinitionOutputWithContext(ctx context.Context) DataQualityJobDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataQualityJobDefinitionOutput)
}

type DataQualityJobDefinitionOutput struct{ *pulumi.OutputState }

func (DataQualityJobDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataQualityJobDefinition)(nil)).Elem()
}

func (o DataQualityJobDefinitionOutput) ToDataQualityJobDefinitionOutput() DataQualityJobDefinitionOutput {
	return o
}

func (o DataQualityJobDefinitionOutput) ToDataQualityJobDefinitionOutputWithContext(ctx context.Context) DataQualityJobDefinitionOutput {
	return o
}

// The time at which the job definition was created.
func (o DataQualityJobDefinitionOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *DataQualityJobDefinition) pulumi.StringOutput { return v.CreationTime }).(pulumi.StringOutput)
}

func (o DataQualityJobDefinitionOutput) DataQualityAppSpecification() DataQualityJobDefinitionDataQualityAppSpecificationOutput {
	return o.ApplyT(func(v *DataQualityJobDefinition) DataQualityJobDefinitionDataQualityAppSpecificationOutput {
		return v.DataQualityAppSpecification
	}).(DataQualityJobDefinitionDataQualityAppSpecificationOutput)
}

func (o DataQualityJobDefinitionOutput) DataQualityBaselineConfig() DataQualityJobDefinitionDataQualityBaselineConfigPtrOutput {
	return o.ApplyT(func(v *DataQualityJobDefinition) DataQualityJobDefinitionDataQualityBaselineConfigPtrOutput {
		return v.DataQualityBaselineConfig
	}).(DataQualityJobDefinitionDataQualityBaselineConfigPtrOutput)
}

func (o DataQualityJobDefinitionOutput) DataQualityJobInput() DataQualityJobDefinitionDataQualityJobInputOutput {
	return o.ApplyT(func(v *DataQualityJobDefinition) DataQualityJobDefinitionDataQualityJobInputOutput {
		return v.DataQualityJobInput
	}).(DataQualityJobDefinitionDataQualityJobInputOutput)
}

func (o DataQualityJobDefinitionOutput) DataQualityJobOutputConfig() DataQualityJobDefinitionMonitoringOutputConfigOutput {
	return o.ApplyT(func(v *DataQualityJobDefinition) DataQualityJobDefinitionMonitoringOutputConfigOutput {
		return v.DataQualityJobOutputConfig
	}).(DataQualityJobDefinitionMonitoringOutputConfigOutput)
}

func (o DataQualityJobDefinitionOutput) EndpointName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataQualityJobDefinition) pulumi.StringPtrOutput { return v.EndpointName }).(pulumi.StringPtrOutput)
}

// The Amazon Resource Name (ARN) of job definition.
func (o DataQualityJobDefinitionOutput) JobDefinitionArn() pulumi.StringOutput {
	return o.ApplyT(func(v *DataQualityJobDefinition) pulumi.StringOutput { return v.JobDefinitionArn }).(pulumi.StringOutput)
}

func (o DataQualityJobDefinitionOutput) JobDefinitionName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataQualityJobDefinition) pulumi.StringPtrOutput { return v.JobDefinitionName }).(pulumi.StringPtrOutput)
}

func (o DataQualityJobDefinitionOutput) JobResources() DataQualityJobDefinitionMonitoringResourcesOutput {
	return o.ApplyT(func(v *DataQualityJobDefinition) DataQualityJobDefinitionMonitoringResourcesOutput {
		return v.JobResources
	}).(DataQualityJobDefinitionMonitoringResourcesOutput)
}

func (o DataQualityJobDefinitionOutput) NetworkConfig() DataQualityJobDefinitionNetworkConfigPtrOutput {
	return o.ApplyT(func(v *DataQualityJobDefinition) DataQualityJobDefinitionNetworkConfigPtrOutput {
		return v.NetworkConfig
	}).(DataQualityJobDefinitionNetworkConfigPtrOutput)
}

// The Amazon Resource Name (ARN) of an IAM role that Amazon SageMaker can assume to perform tasks on your behalf.
func (o DataQualityJobDefinitionOutput) RoleArn() pulumi.StringOutput {
	return o.ApplyT(func(v *DataQualityJobDefinition) pulumi.StringOutput { return v.RoleArn }).(pulumi.StringOutput)
}

func (o DataQualityJobDefinitionOutput) StoppingCondition() DataQualityJobDefinitionStoppingConditionPtrOutput {
	return o.ApplyT(func(v *DataQualityJobDefinition) DataQualityJobDefinitionStoppingConditionPtrOutput {
		return v.StoppingCondition
	}).(DataQualityJobDefinitionStoppingConditionPtrOutput)
}

// An array of key-value pairs to apply to this resource.
func (o DataQualityJobDefinitionOutput) Tags() DataQualityJobDefinitionTagArrayOutput {
	return o.ApplyT(func(v *DataQualityJobDefinition) DataQualityJobDefinitionTagArrayOutput { return v.Tags }).(DataQualityJobDefinitionTagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DataQualityJobDefinitionInput)(nil)).Elem(), &DataQualityJobDefinition{})
	pulumi.RegisterOutputType(DataQualityJobDefinitionOutput{})
}

// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package lookoutequipment

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource schema for LookoutEquipment InferenceScheduler.
//
// Deprecated: InferenceScheduler is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type InferenceScheduler struct {
	pulumi.CustomResourceState

	// A period of time (in minutes) by which inference on the data is delayed after the data starts.
	DataDelayOffsetInMinutes pulumi.IntPtrOutput `pulumi:"dataDelayOffsetInMinutes"`
	// Specifies configuration information for the input data for the inference scheduler, including delimiter, format, and dataset location.
	DataInputConfiguration DataInputConfigurationPropertiesOutput `pulumi:"dataInputConfiguration"`
	// Specifies configuration information for the output results for the inference scheduler, including the S3 location for the output.
	DataOutputConfiguration DataOutputConfigurationPropertiesOutput `pulumi:"dataOutputConfiguration"`
	// How often data is uploaded to the source S3 bucket for the input data.
	DataUploadFrequency InferenceSchedulerDataUploadFrequencyOutput `pulumi:"dataUploadFrequency"`
	// The Amazon Resource Name (ARN) of the inference scheduler being created.
	InferenceSchedulerArn pulumi.StringOutput `pulumi:"inferenceSchedulerArn"`
	// The name of the inference scheduler being created.
	InferenceSchedulerName pulumi.StringPtrOutput `pulumi:"inferenceSchedulerName"`
	// The name of the previously trained ML model being used to create the inference scheduler.
	ModelName pulumi.StringOutput `pulumi:"modelName"`
	// The Amazon Resource Name (ARN) of a role with permission to access the data source being used for the inference.
	RoleArn pulumi.StringOutput `pulumi:"roleArn"`
	// Provides the identifier of the AWS KMS customer master key (CMK) used to encrypt inference scheduler data by Amazon Lookout for Equipment.
	ServerSideKmsKeyId pulumi.StringPtrOutput `pulumi:"serverSideKmsKeyId"`
	// Any tags associated with the inference scheduler.
	Tags InferenceSchedulerTagArrayOutput `pulumi:"tags"`
}

// NewInferenceScheduler registers a new resource with the given unique name, arguments, and options.
func NewInferenceScheduler(ctx *pulumi.Context,
	name string, args *InferenceSchedulerArgs, opts ...pulumi.ResourceOption) (*InferenceScheduler, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataInputConfiguration == nil {
		return nil, errors.New("invalid value for required argument 'DataInputConfiguration'")
	}
	if args.DataOutputConfiguration == nil {
		return nil, errors.New("invalid value for required argument 'DataOutputConfiguration'")
	}
	if args.DataUploadFrequency == nil {
		return nil, errors.New("invalid value for required argument 'DataUploadFrequency'")
	}
	if args.ModelName == nil {
		return nil, errors.New("invalid value for required argument 'ModelName'")
	}
	if args.RoleArn == nil {
		return nil, errors.New("invalid value for required argument 'RoleArn'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"inferenceSchedulerName",
		"modelName",
		"serverSideKmsKeyId",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource InferenceScheduler
	err := ctx.RegisterResource("aws-native:lookoutequipment:InferenceScheduler", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInferenceScheduler gets an existing InferenceScheduler resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInferenceScheduler(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InferenceSchedulerState, opts ...pulumi.ResourceOption) (*InferenceScheduler, error) {
	var resource InferenceScheduler
	err := ctx.ReadResource("aws-native:lookoutequipment:InferenceScheduler", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InferenceScheduler resources.
type inferenceSchedulerState struct {
}

type InferenceSchedulerState struct {
}

func (InferenceSchedulerState) ElementType() reflect.Type {
	return reflect.TypeOf((*inferenceSchedulerState)(nil)).Elem()
}

type inferenceSchedulerArgs struct {
	// A period of time (in minutes) by which inference on the data is delayed after the data starts.
	DataDelayOffsetInMinutes *int `pulumi:"dataDelayOffsetInMinutes"`
	// Specifies configuration information for the input data for the inference scheduler, including delimiter, format, and dataset location.
	DataInputConfiguration DataInputConfigurationProperties `pulumi:"dataInputConfiguration"`
	// Specifies configuration information for the output results for the inference scheduler, including the S3 location for the output.
	DataOutputConfiguration DataOutputConfigurationProperties `pulumi:"dataOutputConfiguration"`
	// How often data is uploaded to the source S3 bucket for the input data.
	DataUploadFrequency InferenceSchedulerDataUploadFrequency `pulumi:"dataUploadFrequency"`
	// The name of the inference scheduler being created.
	InferenceSchedulerName *string `pulumi:"inferenceSchedulerName"`
	// The name of the previously trained ML model being used to create the inference scheduler.
	ModelName string `pulumi:"modelName"`
	// The Amazon Resource Name (ARN) of a role with permission to access the data source being used for the inference.
	RoleArn string `pulumi:"roleArn"`
	// Provides the identifier of the AWS KMS customer master key (CMK) used to encrypt inference scheduler data by Amazon Lookout for Equipment.
	ServerSideKmsKeyId *string `pulumi:"serverSideKmsKeyId"`
	// Any tags associated with the inference scheduler.
	Tags []InferenceSchedulerTag `pulumi:"tags"`
}

// The set of arguments for constructing a InferenceScheduler resource.
type InferenceSchedulerArgs struct {
	// A period of time (in minutes) by which inference on the data is delayed after the data starts.
	DataDelayOffsetInMinutes pulumi.IntPtrInput
	// Specifies configuration information for the input data for the inference scheduler, including delimiter, format, and dataset location.
	DataInputConfiguration DataInputConfigurationPropertiesInput
	// Specifies configuration information for the output results for the inference scheduler, including the S3 location for the output.
	DataOutputConfiguration DataOutputConfigurationPropertiesInput
	// How often data is uploaded to the source S3 bucket for the input data.
	DataUploadFrequency InferenceSchedulerDataUploadFrequencyInput
	// The name of the inference scheduler being created.
	InferenceSchedulerName pulumi.StringPtrInput
	// The name of the previously trained ML model being used to create the inference scheduler.
	ModelName pulumi.StringInput
	// The Amazon Resource Name (ARN) of a role with permission to access the data source being used for the inference.
	RoleArn pulumi.StringInput
	// Provides the identifier of the AWS KMS customer master key (CMK) used to encrypt inference scheduler data by Amazon Lookout for Equipment.
	ServerSideKmsKeyId pulumi.StringPtrInput
	// Any tags associated with the inference scheduler.
	Tags InferenceSchedulerTagArrayInput
}

func (InferenceSchedulerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*inferenceSchedulerArgs)(nil)).Elem()
}

type InferenceSchedulerInput interface {
	pulumi.Input

	ToInferenceSchedulerOutput() InferenceSchedulerOutput
	ToInferenceSchedulerOutputWithContext(ctx context.Context) InferenceSchedulerOutput
}

func (*InferenceScheduler) ElementType() reflect.Type {
	return reflect.TypeOf((**InferenceScheduler)(nil)).Elem()
}

func (i *InferenceScheduler) ToInferenceSchedulerOutput() InferenceSchedulerOutput {
	return i.ToInferenceSchedulerOutputWithContext(context.Background())
}

func (i *InferenceScheduler) ToInferenceSchedulerOutputWithContext(ctx context.Context) InferenceSchedulerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InferenceSchedulerOutput)
}

type InferenceSchedulerOutput struct{ *pulumi.OutputState }

func (InferenceSchedulerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InferenceScheduler)(nil)).Elem()
}

func (o InferenceSchedulerOutput) ToInferenceSchedulerOutput() InferenceSchedulerOutput {
	return o
}

func (o InferenceSchedulerOutput) ToInferenceSchedulerOutputWithContext(ctx context.Context) InferenceSchedulerOutput {
	return o
}

// A period of time (in minutes) by which inference on the data is delayed after the data starts.
func (o InferenceSchedulerOutput) DataDelayOffsetInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *InferenceScheduler) pulumi.IntPtrOutput { return v.DataDelayOffsetInMinutes }).(pulumi.IntPtrOutput)
}

// Specifies configuration information for the input data for the inference scheduler, including delimiter, format, and dataset location.
func (o InferenceSchedulerOutput) DataInputConfiguration() DataInputConfigurationPropertiesOutput {
	return o.ApplyT(func(v *InferenceScheduler) DataInputConfigurationPropertiesOutput { return v.DataInputConfiguration }).(DataInputConfigurationPropertiesOutput)
}

// Specifies configuration information for the output results for the inference scheduler, including the S3 location for the output.
func (o InferenceSchedulerOutput) DataOutputConfiguration() DataOutputConfigurationPropertiesOutput {
	return o.ApplyT(func(v *InferenceScheduler) DataOutputConfigurationPropertiesOutput { return v.DataOutputConfiguration }).(DataOutputConfigurationPropertiesOutput)
}

// How often data is uploaded to the source S3 bucket for the input data.
func (o InferenceSchedulerOutput) DataUploadFrequency() InferenceSchedulerDataUploadFrequencyOutput {
	return o.ApplyT(func(v *InferenceScheduler) InferenceSchedulerDataUploadFrequencyOutput { return v.DataUploadFrequency }).(InferenceSchedulerDataUploadFrequencyOutput)
}

// The Amazon Resource Name (ARN) of the inference scheduler being created.
func (o InferenceSchedulerOutput) InferenceSchedulerArn() pulumi.StringOutput {
	return o.ApplyT(func(v *InferenceScheduler) pulumi.StringOutput { return v.InferenceSchedulerArn }).(pulumi.StringOutput)
}

// The name of the inference scheduler being created.
func (o InferenceSchedulerOutput) InferenceSchedulerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InferenceScheduler) pulumi.StringPtrOutput { return v.InferenceSchedulerName }).(pulumi.StringPtrOutput)
}

// The name of the previously trained ML model being used to create the inference scheduler.
func (o InferenceSchedulerOutput) ModelName() pulumi.StringOutput {
	return o.ApplyT(func(v *InferenceScheduler) pulumi.StringOutput { return v.ModelName }).(pulumi.StringOutput)
}

// The Amazon Resource Name (ARN) of a role with permission to access the data source being used for the inference.
func (o InferenceSchedulerOutput) RoleArn() pulumi.StringOutput {
	return o.ApplyT(func(v *InferenceScheduler) pulumi.StringOutput { return v.RoleArn }).(pulumi.StringOutput)
}

// Provides the identifier of the AWS KMS customer master key (CMK) used to encrypt inference scheduler data by Amazon Lookout for Equipment.
func (o InferenceSchedulerOutput) ServerSideKmsKeyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InferenceScheduler) pulumi.StringPtrOutput { return v.ServerSideKmsKeyId }).(pulumi.StringPtrOutput)
}

// Any tags associated with the inference scheduler.
func (o InferenceSchedulerOutput) Tags() InferenceSchedulerTagArrayOutput {
	return o.ApplyT(func(v *InferenceScheduler) InferenceSchedulerTagArrayOutput { return v.Tags }).(InferenceSchedulerTagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InferenceSchedulerInput)(nil)).Elem(), &InferenceScheduler{})
	pulumi.RegisterOutputType(InferenceSchedulerOutput{})
}

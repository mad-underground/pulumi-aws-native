// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package batch

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Batch::JobDefinition
type JobDefinition struct {
	pulumi.CustomResourceState

	AwsId               pulumi.StringOutput                       `pulumi:"awsId"`
	ContainerProperties JobDefinitionContainerPropertiesPtrOutput `pulumi:"containerProperties"`
	EcsProperties       JobDefinitionEcsPropertiesPtrOutput       `pulumi:"ecsProperties"`
	EksProperties       JobDefinitionEksPropertiesPtrOutput       `pulumi:"eksProperties"`
	JobDefinitionName   pulumi.StringPtrOutput                    `pulumi:"jobDefinitionName"`
	NodeProperties      JobDefinitionNodePropertiesPtrOutput      `pulumi:"nodeProperties"`
	// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Batch::JobDefinition` for more information about the expected schema for this property.
	Parameters           pulumi.AnyOutput                    `pulumi:"parameters"`
	PlatformCapabilities pulumi.StringArrayOutput            `pulumi:"platformCapabilities"`
	PropagateTags        pulumi.BoolPtrOutput                `pulumi:"propagateTags"`
	RetryStrategy        JobDefinitionRetryStrategyPtrOutput `pulumi:"retryStrategy"`
	SchedulingPriority   pulumi.IntPtrOutput                 `pulumi:"schedulingPriority"`
	// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Batch::JobDefinition` for more information about the expected schema for this property.
	Tags    pulumi.AnyOutput              `pulumi:"tags"`
	Timeout JobDefinitionTimeoutPtrOutput `pulumi:"timeout"`
	Type    pulumi.StringOutput           `pulumi:"type"`
}

// NewJobDefinition registers a new resource with the given unique name, arguments, and options.
func NewJobDefinition(ctx *pulumi.Context,
	name string, args *JobDefinitionArgs, opts ...pulumi.ResourceOption) (*JobDefinition, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"jobDefinitionName",
		"tags",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource JobDefinition
	err := ctx.RegisterResource("aws-native:batch:JobDefinition", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetJobDefinition gets an existing JobDefinition resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetJobDefinition(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *JobDefinitionState, opts ...pulumi.ResourceOption) (*JobDefinition, error) {
	var resource JobDefinition
	err := ctx.ReadResource("aws-native:batch:JobDefinition", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering JobDefinition resources.
type jobDefinitionState struct {
}

type JobDefinitionState struct {
}

func (JobDefinitionState) ElementType() reflect.Type {
	return reflect.TypeOf((*jobDefinitionState)(nil)).Elem()
}

type jobDefinitionArgs struct {
	ContainerProperties *JobDefinitionContainerProperties `pulumi:"containerProperties"`
	EcsProperties       *JobDefinitionEcsProperties       `pulumi:"ecsProperties"`
	EksProperties       *JobDefinitionEksProperties       `pulumi:"eksProperties"`
	JobDefinitionName   *string                           `pulumi:"jobDefinitionName"`
	NodeProperties      *JobDefinitionNodeProperties      `pulumi:"nodeProperties"`
	// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Batch::JobDefinition` for more information about the expected schema for this property.
	Parameters           interface{}                 `pulumi:"parameters"`
	PlatformCapabilities []string                    `pulumi:"platformCapabilities"`
	PropagateTags        *bool                       `pulumi:"propagateTags"`
	RetryStrategy        *JobDefinitionRetryStrategy `pulumi:"retryStrategy"`
	SchedulingPriority   *int                        `pulumi:"schedulingPriority"`
	// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Batch::JobDefinition` for more information about the expected schema for this property.
	Tags    interface{}           `pulumi:"tags"`
	Timeout *JobDefinitionTimeout `pulumi:"timeout"`
	Type    string                `pulumi:"type"`
}

// The set of arguments for constructing a JobDefinition resource.
type JobDefinitionArgs struct {
	ContainerProperties JobDefinitionContainerPropertiesPtrInput
	EcsProperties       JobDefinitionEcsPropertiesPtrInput
	EksProperties       JobDefinitionEksPropertiesPtrInput
	JobDefinitionName   pulumi.StringPtrInput
	NodeProperties      JobDefinitionNodePropertiesPtrInput
	// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Batch::JobDefinition` for more information about the expected schema for this property.
	Parameters           pulumi.Input
	PlatformCapabilities pulumi.StringArrayInput
	PropagateTags        pulumi.BoolPtrInput
	RetryStrategy        JobDefinitionRetryStrategyPtrInput
	SchedulingPriority   pulumi.IntPtrInput
	// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Batch::JobDefinition` for more information about the expected schema for this property.
	Tags    pulumi.Input
	Timeout JobDefinitionTimeoutPtrInput
	Type    pulumi.StringInput
}

func (JobDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*jobDefinitionArgs)(nil)).Elem()
}

type JobDefinitionInput interface {
	pulumi.Input

	ToJobDefinitionOutput() JobDefinitionOutput
	ToJobDefinitionOutputWithContext(ctx context.Context) JobDefinitionOutput
}

func (*JobDefinition) ElementType() reflect.Type {
	return reflect.TypeOf((**JobDefinition)(nil)).Elem()
}

func (i *JobDefinition) ToJobDefinitionOutput() JobDefinitionOutput {
	return i.ToJobDefinitionOutputWithContext(context.Background())
}

func (i *JobDefinition) ToJobDefinitionOutputWithContext(ctx context.Context) JobDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobDefinitionOutput)
}

type JobDefinitionOutput struct{ *pulumi.OutputState }

func (JobDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobDefinition)(nil)).Elem()
}

func (o JobDefinitionOutput) ToJobDefinitionOutput() JobDefinitionOutput {
	return o
}

func (o JobDefinitionOutput) ToJobDefinitionOutputWithContext(ctx context.Context) JobDefinitionOutput {
	return o
}

func (o JobDefinitionOutput) AwsId() pulumi.StringOutput {
	return o.ApplyT(func(v *JobDefinition) pulumi.StringOutput { return v.AwsId }).(pulumi.StringOutput)
}

func (o JobDefinitionOutput) ContainerProperties() JobDefinitionContainerPropertiesPtrOutput {
	return o.ApplyT(func(v *JobDefinition) JobDefinitionContainerPropertiesPtrOutput { return v.ContainerProperties }).(JobDefinitionContainerPropertiesPtrOutput)
}

func (o JobDefinitionOutput) EcsProperties() JobDefinitionEcsPropertiesPtrOutput {
	return o.ApplyT(func(v *JobDefinition) JobDefinitionEcsPropertiesPtrOutput { return v.EcsProperties }).(JobDefinitionEcsPropertiesPtrOutput)
}

func (o JobDefinitionOutput) EksProperties() JobDefinitionEksPropertiesPtrOutput {
	return o.ApplyT(func(v *JobDefinition) JobDefinitionEksPropertiesPtrOutput { return v.EksProperties }).(JobDefinitionEksPropertiesPtrOutput)
}

func (o JobDefinitionOutput) JobDefinitionName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobDefinition) pulumi.StringPtrOutput { return v.JobDefinitionName }).(pulumi.StringPtrOutput)
}

func (o JobDefinitionOutput) NodeProperties() JobDefinitionNodePropertiesPtrOutput {
	return o.ApplyT(func(v *JobDefinition) JobDefinitionNodePropertiesPtrOutput { return v.NodeProperties }).(JobDefinitionNodePropertiesPtrOutput)
}

// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Batch::JobDefinition` for more information about the expected schema for this property.
func (o JobDefinitionOutput) Parameters() pulumi.AnyOutput {
	return o.ApplyT(func(v *JobDefinition) pulumi.AnyOutput { return v.Parameters }).(pulumi.AnyOutput)
}

func (o JobDefinitionOutput) PlatformCapabilities() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *JobDefinition) pulumi.StringArrayOutput { return v.PlatformCapabilities }).(pulumi.StringArrayOutput)
}

func (o JobDefinitionOutput) PropagateTags() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *JobDefinition) pulumi.BoolPtrOutput { return v.PropagateTags }).(pulumi.BoolPtrOutput)
}

func (o JobDefinitionOutput) RetryStrategy() JobDefinitionRetryStrategyPtrOutput {
	return o.ApplyT(func(v *JobDefinition) JobDefinitionRetryStrategyPtrOutput { return v.RetryStrategy }).(JobDefinitionRetryStrategyPtrOutput)
}

func (o JobDefinitionOutput) SchedulingPriority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *JobDefinition) pulumi.IntPtrOutput { return v.SchedulingPriority }).(pulumi.IntPtrOutput)
}

// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Batch::JobDefinition` for more information about the expected schema for this property.
func (o JobDefinitionOutput) Tags() pulumi.AnyOutput {
	return o.ApplyT(func(v *JobDefinition) pulumi.AnyOutput { return v.Tags }).(pulumi.AnyOutput)
}

func (o JobDefinitionOutput) Timeout() JobDefinitionTimeoutPtrOutput {
	return o.ApplyT(func(v *JobDefinition) JobDefinitionTimeoutPtrOutput { return v.Timeout }).(JobDefinitionTimeoutPtrOutput)
}

func (o JobDefinitionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *JobDefinition) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*JobDefinitionInput)(nil)).Elem(), &JobDefinition{})
	pulumi.RegisterOutputType(JobDefinitionOutput{})
}

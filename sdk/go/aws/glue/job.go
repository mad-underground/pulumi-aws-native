// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package glue

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Glue::Job
//
// Deprecated: Job is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type Job struct {
	pulumi.CustomResourceState

	AllocatedCapacity pulumi.Float64PtrOutput     `pulumi:"allocatedCapacity"`
	AwsId             pulumi.StringOutput         `pulumi:"awsId"`
	Command           JobCommandOutput            `pulumi:"command"`
	Connections       JobConnectionsListPtrOutput `pulumi:"connections"`
	// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Glue::Job` for more information about the expected schema for this property.
	DefaultArguments  pulumi.AnyOutput              `pulumi:"defaultArguments"`
	Description       pulumi.StringPtrOutput        `pulumi:"description"`
	ExecutionClass    pulumi.StringPtrOutput        `pulumi:"executionClass"`
	ExecutionProperty JobExecutionPropertyPtrOutput `pulumi:"executionProperty"`
	GlueVersion       pulumi.StringPtrOutput        `pulumi:"glueVersion"`
	LogUri            pulumi.StringPtrOutput        `pulumi:"logUri"`
	MaxCapacity       pulumi.Float64PtrOutput       `pulumi:"maxCapacity"`
	MaxRetries        pulumi.Float64PtrOutput       `pulumi:"maxRetries"`
	Name              pulumi.StringPtrOutput        `pulumi:"name"`
	// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Glue::Job` for more information about the expected schema for this property.
	NonOverridableArguments pulumi.AnyOutput                 `pulumi:"nonOverridableArguments"`
	NotificationProperty    JobNotificationPropertyPtrOutput `pulumi:"notificationProperty"`
	NumberOfWorkers         pulumi.IntPtrOutput              `pulumi:"numberOfWorkers"`
	Role                    pulumi.StringOutput              `pulumi:"role"`
	SecurityConfiguration   pulumi.StringPtrOutput           `pulumi:"securityConfiguration"`
	// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Glue::Job` for more information about the expected schema for this property.
	Tags       pulumi.AnyOutput       `pulumi:"tags"`
	Timeout    pulumi.IntPtrOutput    `pulumi:"timeout"`
	WorkerType pulumi.StringPtrOutput `pulumi:"workerType"`
}

// NewJob registers a new resource with the given unique name, arguments, and options.
func NewJob(ctx *pulumi.Context,
	name string, args *JobArgs, opts ...pulumi.ResourceOption) (*Job, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Command == nil {
		return nil, errors.New("invalid value for required argument 'Command'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"name",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Job
	err := ctx.RegisterResource("aws-native:glue:Job", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetJob gets an existing Job resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetJob(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *JobState, opts ...pulumi.ResourceOption) (*Job, error) {
	var resource Job
	err := ctx.ReadResource("aws-native:glue:Job", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Job resources.
type jobState struct {
}

type JobState struct {
}

func (JobState) ElementType() reflect.Type {
	return reflect.TypeOf((*jobState)(nil)).Elem()
}

type jobArgs struct {
	AllocatedCapacity *float64            `pulumi:"allocatedCapacity"`
	Command           JobCommand          `pulumi:"command"`
	Connections       *JobConnectionsList `pulumi:"connections"`
	// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Glue::Job` for more information about the expected schema for this property.
	DefaultArguments  interface{}           `pulumi:"defaultArguments"`
	Description       *string               `pulumi:"description"`
	ExecutionClass    *string               `pulumi:"executionClass"`
	ExecutionProperty *JobExecutionProperty `pulumi:"executionProperty"`
	GlueVersion       *string               `pulumi:"glueVersion"`
	LogUri            *string               `pulumi:"logUri"`
	MaxCapacity       *float64              `pulumi:"maxCapacity"`
	MaxRetries        *float64              `pulumi:"maxRetries"`
	Name              *string               `pulumi:"name"`
	// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Glue::Job` for more information about the expected schema for this property.
	NonOverridableArguments interface{}              `pulumi:"nonOverridableArguments"`
	NotificationProperty    *JobNotificationProperty `pulumi:"notificationProperty"`
	NumberOfWorkers         *int                     `pulumi:"numberOfWorkers"`
	Role                    string                   `pulumi:"role"`
	SecurityConfiguration   *string                  `pulumi:"securityConfiguration"`
	// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Glue::Job` for more information about the expected schema for this property.
	Tags       interface{} `pulumi:"tags"`
	Timeout    *int        `pulumi:"timeout"`
	WorkerType *string     `pulumi:"workerType"`
}

// The set of arguments for constructing a Job resource.
type JobArgs struct {
	AllocatedCapacity pulumi.Float64PtrInput
	Command           JobCommandInput
	Connections       JobConnectionsListPtrInput
	// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Glue::Job` for more information about the expected schema for this property.
	DefaultArguments  pulumi.Input
	Description       pulumi.StringPtrInput
	ExecutionClass    pulumi.StringPtrInput
	ExecutionProperty JobExecutionPropertyPtrInput
	GlueVersion       pulumi.StringPtrInput
	LogUri            pulumi.StringPtrInput
	MaxCapacity       pulumi.Float64PtrInput
	MaxRetries        pulumi.Float64PtrInput
	Name              pulumi.StringPtrInput
	// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Glue::Job` for more information about the expected schema for this property.
	NonOverridableArguments pulumi.Input
	NotificationProperty    JobNotificationPropertyPtrInput
	NumberOfWorkers         pulumi.IntPtrInput
	Role                    pulumi.StringInput
	SecurityConfiguration   pulumi.StringPtrInput
	// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Glue::Job` for more information about the expected schema for this property.
	Tags       pulumi.Input
	Timeout    pulumi.IntPtrInput
	WorkerType pulumi.StringPtrInput
}

func (JobArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*jobArgs)(nil)).Elem()
}

type JobInput interface {
	pulumi.Input

	ToJobOutput() JobOutput
	ToJobOutputWithContext(ctx context.Context) JobOutput
}

func (*Job) ElementType() reflect.Type {
	return reflect.TypeOf((**Job)(nil)).Elem()
}

func (i *Job) ToJobOutput() JobOutput {
	return i.ToJobOutputWithContext(context.Background())
}

func (i *Job) ToJobOutputWithContext(ctx context.Context) JobOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobOutput)
}

type JobOutput struct{ *pulumi.OutputState }

func (JobOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Job)(nil)).Elem()
}

func (o JobOutput) ToJobOutput() JobOutput {
	return o
}

func (o JobOutput) ToJobOutputWithContext(ctx context.Context) JobOutput {
	return o
}

func (o JobOutput) AllocatedCapacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *Job) pulumi.Float64PtrOutput { return v.AllocatedCapacity }).(pulumi.Float64PtrOutput)
}

func (o JobOutput) AwsId() pulumi.StringOutput {
	return o.ApplyT(func(v *Job) pulumi.StringOutput { return v.AwsId }).(pulumi.StringOutput)
}

func (o JobOutput) Command() JobCommandOutput {
	return o.ApplyT(func(v *Job) JobCommandOutput { return v.Command }).(JobCommandOutput)
}

func (o JobOutput) Connections() JobConnectionsListPtrOutput {
	return o.ApplyT(func(v *Job) JobConnectionsListPtrOutput { return v.Connections }).(JobConnectionsListPtrOutput)
}

// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Glue::Job` for more information about the expected schema for this property.
func (o JobOutput) DefaultArguments() pulumi.AnyOutput {
	return o.ApplyT(func(v *Job) pulumi.AnyOutput { return v.DefaultArguments }).(pulumi.AnyOutput)
}

func (o JobOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Job) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o JobOutput) ExecutionClass() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Job) pulumi.StringPtrOutput { return v.ExecutionClass }).(pulumi.StringPtrOutput)
}

func (o JobOutput) ExecutionProperty() JobExecutionPropertyPtrOutput {
	return o.ApplyT(func(v *Job) JobExecutionPropertyPtrOutput { return v.ExecutionProperty }).(JobExecutionPropertyPtrOutput)
}

func (o JobOutput) GlueVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Job) pulumi.StringPtrOutput { return v.GlueVersion }).(pulumi.StringPtrOutput)
}

func (o JobOutput) LogUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Job) pulumi.StringPtrOutput { return v.LogUri }).(pulumi.StringPtrOutput)
}

func (o JobOutput) MaxCapacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *Job) pulumi.Float64PtrOutput { return v.MaxCapacity }).(pulumi.Float64PtrOutput)
}

func (o JobOutput) MaxRetries() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *Job) pulumi.Float64PtrOutput { return v.MaxRetries }).(pulumi.Float64PtrOutput)
}

func (o JobOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Job) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Glue::Job` for more information about the expected schema for this property.
func (o JobOutput) NonOverridableArguments() pulumi.AnyOutput {
	return o.ApplyT(func(v *Job) pulumi.AnyOutput { return v.NonOverridableArguments }).(pulumi.AnyOutput)
}

func (o JobOutput) NotificationProperty() JobNotificationPropertyPtrOutput {
	return o.ApplyT(func(v *Job) JobNotificationPropertyPtrOutput { return v.NotificationProperty }).(JobNotificationPropertyPtrOutput)
}

func (o JobOutput) NumberOfWorkers() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Job) pulumi.IntPtrOutput { return v.NumberOfWorkers }).(pulumi.IntPtrOutput)
}

func (o JobOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *Job) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func (o JobOutput) SecurityConfiguration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Job) pulumi.StringPtrOutput { return v.SecurityConfiguration }).(pulumi.StringPtrOutput)
}

// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Glue::Job` for more information about the expected schema for this property.
func (o JobOutput) Tags() pulumi.AnyOutput {
	return o.ApplyT(func(v *Job) pulumi.AnyOutput { return v.Tags }).(pulumi.AnyOutput)
}

func (o JobOutput) Timeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Job) pulumi.IntPtrOutput { return v.Timeout }).(pulumi.IntPtrOutput)
}

func (o JobOutput) WorkerType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Job) pulumi.StringPtrOutput { return v.WorkerType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*JobInput)(nil)).Elem(), &Job{})
	pulumi.RegisterOutputType(JobOutput{})
}

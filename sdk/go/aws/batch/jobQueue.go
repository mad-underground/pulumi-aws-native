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

// Resource Type definition for AWS::Batch::JobQueue
type JobQueue struct {
	pulumi.CustomResourceState

	ComputeEnvironmentOrder JobQueueComputeEnvironmentOrderArrayOutput `pulumi:"computeEnvironmentOrder"`
	JobQueueArn             pulumi.StringOutput                        `pulumi:"jobQueueArn"`
	JobQueueName            pulumi.StringPtrOutput                     `pulumi:"jobQueueName"`
	Priority                pulumi.IntOutput                           `pulumi:"priority"`
	SchedulingPolicyArn     pulumi.StringPtrOutput                     `pulumi:"schedulingPolicyArn"`
	State                   JobQueueStateEnumPtrOutput                 `pulumi:"state"`
	// A key-value pair to associate with a resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewJobQueue registers a new resource with the given unique name, arguments, and options.
func NewJobQueue(ctx *pulumi.Context,
	name string, args *JobQueueArgs, opts ...pulumi.ResourceOption) (*JobQueue, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ComputeEnvironmentOrder == nil {
		return nil, errors.New("invalid value for required argument 'ComputeEnvironmentOrder'")
	}
	if args.Priority == nil {
		return nil, errors.New("invalid value for required argument 'Priority'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"jobQueueName",
		"tags.*",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource JobQueue
	err := ctx.RegisterResource("aws-native:batch:JobQueue", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetJobQueue gets an existing JobQueue resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetJobQueue(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *JobQueueState, opts ...pulumi.ResourceOption) (*JobQueue, error) {
	var resource JobQueue
	err := ctx.ReadResource("aws-native:batch:JobQueue", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering JobQueue resources.
type jobQueueState struct {
}

type JobQueueState struct {
}

func (JobQueueState) ElementType() reflect.Type {
	return reflect.TypeOf((*jobQueueState)(nil)).Elem()
}

type jobQueueArgs struct {
	ComputeEnvironmentOrder []JobQueueComputeEnvironmentOrder `pulumi:"computeEnvironmentOrder"`
	JobQueueName            *string                           `pulumi:"jobQueueName"`
	Priority                int                               `pulumi:"priority"`
	SchedulingPolicyArn     *string                           `pulumi:"schedulingPolicyArn"`
	State                   *JobQueueStateEnum                `pulumi:"state"`
	// A key-value pair to associate with a resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a JobQueue resource.
type JobQueueArgs struct {
	ComputeEnvironmentOrder JobQueueComputeEnvironmentOrderArrayInput
	JobQueueName            pulumi.StringPtrInput
	Priority                pulumi.IntInput
	SchedulingPolicyArn     pulumi.StringPtrInput
	State                   JobQueueStateEnumPtrInput
	// A key-value pair to associate with a resource.
	Tags pulumi.StringMapInput
}

func (JobQueueArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*jobQueueArgs)(nil)).Elem()
}

type JobQueueInput interface {
	pulumi.Input

	ToJobQueueOutput() JobQueueOutput
	ToJobQueueOutputWithContext(ctx context.Context) JobQueueOutput
}

func (*JobQueue) ElementType() reflect.Type {
	return reflect.TypeOf((**JobQueue)(nil)).Elem()
}

func (i *JobQueue) ToJobQueueOutput() JobQueueOutput {
	return i.ToJobQueueOutputWithContext(context.Background())
}

func (i *JobQueue) ToJobQueueOutputWithContext(ctx context.Context) JobQueueOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobQueueOutput)
}

type JobQueueOutput struct{ *pulumi.OutputState }

func (JobQueueOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobQueue)(nil)).Elem()
}

func (o JobQueueOutput) ToJobQueueOutput() JobQueueOutput {
	return o
}

func (o JobQueueOutput) ToJobQueueOutputWithContext(ctx context.Context) JobQueueOutput {
	return o
}

func (o JobQueueOutput) ComputeEnvironmentOrder() JobQueueComputeEnvironmentOrderArrayOutput {
	return o.ApplyT(func(v *JobQueue) JobQueueComputeEnvironmentOrderArrayOutput { return v.ComputeEnvironmentOrder }).(JobQueueComputeEnvironmentOrderArrayOutput)
}

func (o JobQueueOutput) JobQueueArn() pulumi.StringOutput {
	return o.ApplyT(func(v *JobQueue) pulumi.StringOutput { return v.JobQueueArn }).(pulumi.StringOutput)
}

func (o JobQueueOutput) JobQueueName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobQueue) pulumi.StringPtrOutput { return v.JobQueueName }).(pulumi.StringPtrOutput)
}

func (o JobQueueOutput) Priority() pulumi.IntOutput {
	return o.ApplyT(func(v *JobQueue) pulumi.IntOutput { return v.Priority }).(pulumi.IntOutput)
}

func (o JobQueueOutput) SchedulingPolicyArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobQueue) pulumi.StringPtrOutput { return v.SchedulingPolicyArn }).(pulumi.StringPtrOutput)
}

func (o JobQueueOutput) State() JobQueueStateEnumPtrOutput {
	return o.ApplyT(func(v *JobQueue) JobQueueStateEnumPtrOutput { return v.State }).(JobQueueStateEnumPtrOutput)
}

// A key-value pair to associate with a resource.
func (o JobQueueOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *JobQueue) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*JobQueueInput)(nil)).Elem(), &JobQueue{})
	pulumi.RegisterOutputType(JobQueueOutput{})
}

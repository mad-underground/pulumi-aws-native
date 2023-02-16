// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sqs

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::SQS::QueuePolicy
//
// Deprecated: QueuePolicy is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type QueuePolicy struct {
	pulumi.CustomResourceState

	PolicyDocument pulumi.AnyOutput         `pulumi:"policyDocument"`
	Queues         pulumi.StringArrayOutput `pulumi:"queues"`
}

// NewQueuePolicy registers a new resource with the given unique name, arguments, and options.
func NewQueuePolicy(ctx *pulumi.Context,
	name string, args *QueuePolicyArgs, opts ...pulumi.ResourceOption) (*QueuePolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PolicyDocument == nil {
		return nil, errors.New("invalid value for required argument 'PolicyDocument'")
	}
	if args.Queues == nil {
		return nil, errors.New("invalid value for required argument 'Queues'")
	}
	var resource QueuePolicy
	err := ctx.RegisterResource("aws-native:sqs:QueuePolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetQueuePolicy gets an existing QueuePolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetQueuePolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *QueuePolicyState, opts ...pulumi.ResourceOption) (*QueuePolicy, error) {
	var resource QueuePolicy
	err := ctx.ReadResource("aws-native:sqs:QueuePolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering QueuePolicy resources.
type queuePolicyState struct {
}

type QueuePolicyState struct {
}

func (QueuePolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*queuePolicyState)(nil)).Elem()
}

type queuePolicyArgs struct {
	PolicyDocument interface{} `pulumi:"policyDocument"`
	Queues         []string    `pulumi:"queues"`
}

// The set of arguments for constructing a QueuePolicy resource.
type QueuePolicyArgs struct {
	PolicyDocument pulumi.Input
	Queues         pulumi.StringArrayInput
}

func (QueuePolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*queuePolicyArgs)(nil)).Elem()
}

type QueuePolicyInput interface {
	pulumi.Input

	ToQueuePolicyOutput() QueuePolicyOutput
	ToQueuePolicyOutputWithContext(ctx context.Context) QueuePolicyOutput
}

func (*QueuePolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**QueuePolicy)(nil)).Elem()
}

func (i *QueuePolicy) ToQueuePolicyOutput() QueuePolicyOutput {
	return i.ToQueuePolicyOutputWithContext(context.Background())
}

func (i *QueuePolicy) ToQueuePolicyOutputWithContext(ctx context.Context) QueuePolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueuePolicyOutput)
}

type QueuePolicyOutput struct{ *pulumi.OutputState }

func (QueuePolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**QueuePolicy)(nil)).Elem()
}

func (o QueuePolicyOutput) ToQueuePolicyOutput() QueuePolicyOutput {
	return o
}

func (o QueuePolicyOutput) ToQueuePolicyOutputWithContext(ctx context.Context) QueuePolicyOutput {
	return o
}

func (o QueuePolicyOutput) PolicyDocument() pulumi.AnyOutput {
	return o.ApplyT(func(v *QueuePolicy) pulumi.AnyOutput { return v.PolicyDocument }).(pulumi.AnyOutput)
}

func (o QueuePolicyOutput) Queues() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *QueuePolicy) pulumi.StringArrayOutput { return v.Queues }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*QueuePolicyInput)(nil)).Elem(), &QueuePolicy{})
	pulumi.RegisterOutputType(QueuePolicyOutput{})
}

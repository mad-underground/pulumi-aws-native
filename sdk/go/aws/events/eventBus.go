// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package events

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource type definition for AWS::Events::EventBus
type EventBus struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) for the event bus.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// If you are creating a partner event bus, this specifies the partner event source that the new event bus will be matched with.
	EventSourceName pulumi.StringPtrOutput `pulumi:"eventSourceName"`
	// The name of the event bus.
	Name pulumi.StringOutput `pulumi:"name"`
	// A JSON string that describes the permission policy statement for the event bus.
	//
	// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Events::EventBus` for more information about the expected schema for this property.
	Policy pulumi.AnyOutput `pulumi:"policy"`
	// Any tags assigned to the event bus.
	Tags aws.TagArrayOutput `pulumi:"tags"`
}

// NewEventBus registers a new resource with the given unique name, arguments, and options.
func NewEventBus(ctx *pulumi.Context,
	name string, args *EventBusArgs, opts ...pulumi.ResourceOption) (*EventBus, error) {
	if args == nil {
		args = &EventBusArgs{}
	}

	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"name",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource EventBus
	err := ctx.RegisterResource("aws-native:events:EventBus", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEventBus gets an existing EventBus resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEventBus(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EventBusState, opts ...pulumi.ResourceOption) (*EventBus, error) {
	var resource EventBus
	err := ctx.ReadResource("aws-native:events:EventBus", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EventBus resources.
type eventBusState struct {
}

type EventBusState struct {
}

func (EventBusState) ElementType() reflect.Type {
	return reflect.TypeOf((*eventBusState)(nil)).Elem()
}

type eventBusArgs struct {
	// If you are creating a partner event bus, this specifies the partner event source that the new event bus will be matched with.
	EventSourceName *string `pulumi:"eventSourceName"`
	// The name of the event bus.
	Name *string `pulumi:"name"`
	// A JSON string that describes the permission policy statement for the event bus.
	//
	// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Events::EventBus` for more information about the expected schema for this property.
	Policy interface{} `pulumi:"policy"`
	// Any tags assigned to the event bus.
	Tags []aws.Tag `pulumi:"tags"`
}

// The set of arguments for constructing a EventBus resource.
type EventBusArgs struct {
	// If you are creating a partner event bus, this specifies the partner event source that the new event bus will be matched with.
	EventSourceName pulumi.StringPtrInput
	// The name of the event bus.
	Name pulumi.StringPtrInput
	// A JSON string that describes the permission policy statement for the event bus.
	//
	// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Events::EventBus` for more information about the expected schema for this property.
	Policy pulumi.Input
	// Any tags assigned to the event bus.
	Tags aws.TagArrayInput
}

func (EventBusArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eventBusArgs)(nil)).Elem()
}

type EventBusInput interface {
	pulumi.Input

	ToEventBusOutput() EventBusOutput
	ToEventBusOutputWithContext(ctx context.Context) EventBusOutput
}

func (*EventBus) ElementType() reflect.Type {
	return reflect.TypeOf((**EventBus)(nil)).Elem()
}

func (i *EventBus) ToEventBusOutput() EventBusOutput {
	return i.ToEventBusOutputWithContext(context.Background())
}

func (i *EventBus) ToEventBusOutputWithContext(ctx context.Context) EventBusOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventBusOutput)
}

type EventBusOutput struct{ *pulumi.OutputState }

func (EventBusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EventBus)(nil)).Elem()
}

func (o EventBusOutput) ToEventBusOutput() EventBusOutput {
	return o
}

func (o EventBusOutput) ToEventBusOutputWithContext(ctx context.Context) EventBusOutput {
	return o
}

// The Amazon Resource Name (ARN) for the event bus.
func (o EventBusOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *EventBus) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// If you are creating a partner event bus, this specifies the partner event source that the new event bus will be matched with.
func (o EventBusOutput) EventSourceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventBus) pulumi.StringPtrOutput { return v.EventSourceName }).(pulumi.StringPtrOutput)
}

// The name of the event bus.
func (o EventBusOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EventBus) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// A JSON string that describes the permission policy statement for the event bus.
//
// Search the [CloudFormation User Guide](https://docs.aws.amazon.com/cloudformation/) for `AWS::Events::EventBus` for more information about the expected schema for this property.
func (o EventBusOutput) Policy() pulumi.AnyOutput {
	return o.ApplyT(func(v *EventBus) pulumi.AnyOutput { return v.Policy }).(pulumi.AnyOutput)
}

// Any tags assigned to the event bus.
func (o EventBusOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v *EventBus) aws.TagArrayOutput { return v.Tags }).(aws.TagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EventBusInput)(nil)).Elem(), &EventBus{})
	pulumi.RegisterOutputType(EventBusOutput{})
}

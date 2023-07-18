// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package events

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Events::EventBus
//
// Deprecated: EventBus is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type EventBus struct {
	pulumi.CustomResourceState

	Arn             pulumi.StringOutput         `pulumi:"arn"`
	EventSourceName pulumi.StringPtrOutput      `pulumi:"eventSourceName"`
	Name            pulumi.StringOutput         `pulumi:"name"`
	Policy          pulumi.StringOutput         `pulumi:"policy"`
	Tags            EventBusTagEntryArrayOutput `pulumi:"tags"`
}

// NewEventBus registers a new resource with the given unique name, arguments, and options.
func NewEventBus(ctx *pulumi.Context,
	name string, args *EventBusArgs, opts ...pulumi.ResourceOption) (*EventBus, error) {
	if args == nil {
		args = &EventBusArgs{}
	}

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
	EventSourceName *string            `pulumi:"eventSourceName"`
	Name            *string            `pulumi:"name"`
	Tags            []EventBusTagEntry `pulumi:"tags"`
}

// The set of arguments for constructing a EventBus resource.
type EventBusArgs struct {
	EventSourceName pulumi.StringPtrInput
	Name            pulumi.StringPtrInput
	Tags            EventBusTagEntryArrayInput
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

func (o EventBusOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *EventBus) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

func (o EventBusOutput) EventSourceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventBus) pulumi.StringPtrOutput { return v.EventSourceName }).(pulumi.StringPtrOutput)
}

func (o EventBusOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EventBus) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o EventBusOutput) Policy() pulumi.StringOutput {
	return o.ApplyT(func(v *EventBus) pulumi.StringOutput { return v.Policy }).(pulumi.StringOutput)
}

func (o EventBusOutput) Tags() EventBusTagEntryArrayOutput {
	return o.ApplyT(func(v *EventBus) EventBusTagEntryArrayOutput { return v.Tags }).(EventBusTagEntryArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EventBusInput)(nil)).Elem(), &EventBus{})
	pulumi.RegisterOutputType(EventBusOutput{})
}

// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ivschat

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource type definition for AWS::IVSChat::Room.
type Room struct {
	pulumi.CustomResourceState

	// Room ARN is automatically generated on creation and assigned as the unique identifier.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Array of logging configuration identifiers attached to the room.
	LoggingConfigurationIdentifiers pulumi.StringArrayOutput `pulumi:"loggingConfigurationIdentifiers"`
	// The maximum number of characters in a single message.
	MaximumMessageLength pulumi.IntPtrOutput `pulumi:"maximumMessageLength"`
	// The maximum number of messages per second that can be sent to the room.
	MaximumMessageRatePerSecond pulumi.IntPtrOutput               `pulumi:"maximumMessageRatePerSecond"`
	MessageReviewHandler        RoomMessageReviewHandlerPtrOutput `pulumi:"messageReviewHandler"`
	// The name of the room. The value does not need to be unique.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// An array of key-value pairs to apply to this resource.
	Tags RoomTagArrayOutput `pulumi:"tags"`
}

// NewRoom registers a new resource with the given unique name, arguments, and options.
func NewRoom(ctx *pulumi.Context,
	name string, args *RoomArgs, opts ...pulumi.ResourceOption) (*Room, error) {
	if args == nil {
		args = &RoomArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Room
	err := ctx.RegisterResource("aws-native:ivschat:Room", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRoom gets an existing Room resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRoom(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RoomState, opts ...pulumi.ResourceOption) (*Room, error) {
	var resource Room
	err := ctx.ReadResource("aws-native:ivschat:Room", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Room resources.
type roomState struct {
}

type RoomState struct {
}

func (RoomState) ElementType() reflect.Type {
	return reflect.TypeOf((*roomState)(nil)).Elem()
}

type roomArgs struct {
	// Array of logging configuration identifiers attached to the room.
	LoggingConfigurationIdentifiers []string `pulumi:"loggingConfigurationIdentifiers"`
	// The maximum number of characters in a single message.
	MaximumMessageLength *int `pulumi:"maximumMessageLength"`
	// The maximum number of messages per second that can be sent to the room.
	MaximumMessageRatePerSecond *int                      `pulumi:"maximumMessageRatePerSecond"`
	MessageReviewHandler        *RoomMessageReviewHandler `pulumi:"messageReviewHandler"`
	// The name of the room. The value does not need to be unique.
	Name *string `pulumi:"name"`
	// An array of key-value pairs to apply to this resource.
	Tags []RoomTag `pulumi:"tags"`
}

// The set of arguments for constructing a Room resource.
type RoomArgs struct {
	// Array of logging configuration identifiers attached to the room.
	LoggingConfigurationIdentifiers pulumi.StringArrayInput
	// The maximum number of characters in a single message.
	MaximumMessageLength pulumi.IntPtrInput
	// The maximum number of messages per second that can be sent to the room.
	MaximumMessageRatePerSecond pulumi.IntPtrInput
	MessageReviewHandler        RoomMessageReviewHandlerPtrInput
	// The name of the room. The value does not need to be unique.
	Name pulumi.StringPtrInput
	// An array of key-value pairs to apply to this resource.
	Tags RoomTagArrayInput
}

func (RoomArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*roomArgs)(nil)).Elem()
}

type RoomInput interface {
	pulumi.Input

	ToRoomOutput() RoomOutput
	ToRoomOutputWithContext(ctx context.Context) RoomOutput
}

func (*Room) ElementType() reflect.Type {
	return reflect.TypeOf((**Room)(nil)).Elem()
}

func (i *Room) ToRoomOutput() RoomOutput {
	return i.ToRoomOutputWithContext(context.Background())
}

func (i *Room) ToRoomOutputWithContext(ctx context.Context) RoomOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoomOutput)
}

type RoomOutput struct{ *pulumi.OutputState }

func (RoomOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Room)(nil)).Elem()
}

func (o RoomOutput) ToRoomOutput() RoomOutput {
	return o
}

func (o RoomOutput) ToRoomOutputWithContext(ctx context.Context) RoomOutput {
	return o
}

// Room ARN is automatically generated on creation and assigned as the unique identifier.
func (o RoomOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Room) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Array of logging configuration identifiers attached to the room.
func (o RoomOutput) LoggingConfigurationIdentifiers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Room) pulumi.StringArrayOutput { return v.LoggingConfigurationIdentifiers }).(pulumi.StringArrayOutput)
}

// The maximum number of characters in a single message.
func (o RoomOutput) MaximumMessageLength() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Room) pulumi.IntPtrOutput { return v.MaximumMessageLength }).(pulumi.IntPtrOutput)
}

// The maximum number of messages per second that can be sent to the room.
func (o RoomOutput) MaximumMessageRatePerSecond() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Room) pulumi.IntPtrOutput { return v.MaximumMessageRatePerSecond }).(pulumi.IntPtrOutput)
}

func (o RoomOutput) MessageReviewHandler() RoomMessageReviewHandlerPtrOutput {
	return o.ApplyT(func(v *Room) RoomMessageReviewHandlerPtrOutput { return v.MessageReviewHandler }).(RoomMessageReviewHandlerPtrOutput)
}

// The name of the room. The value does not need to be unique.
func (o RoomOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Room) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// An array of key-value pairs to apply to this resource.
func (o RoomOutput) Tags() RoomTagArrayOutput {
	return o.ApplyT(func(v *Room) RoomTagArrayOutput { return v.Tags }).(RoomTagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RoomInput)(nil)).Elem(), &Room{})
	pulumi.RegisterOutputType(RoomOutput{})
}

// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gamelift

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::GameLift::GameSessionQueue
//
// Deprecated: GameSessionQueue is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type GameSessionQueue struct {
	pulumi.CustomResourceState

	Arn                   pulumi.StringOutput                            `pulumi:"arn"`
	CustomEventData       pulumi.StringPtrOutput                         `pulumi:"customEventData"`
	Destinations          GameSessionQueueDestinationArrayOutput         `pulumi:"destinations"`
	FilterConfiguration   GameSessionQueueFilterConfigurationPtrOutput   `pulumi:"filterConfiguration"`
	Name                  pulumi.StringOutput                            `pulumi:"name"`
	NotificationTarget    pulumi.StringPtrOutput                         `pulumi:"notificationTarget"`
	PlayerLatencyPolicies GameSessionQueuePlayerLatencyPolicyArrayOutput `pulumi:"playerLatencyPolicies"`
	PriorityConfiguration GameSessionQueuePriorityConfigurationPtrOutput `pulumi:"priorityConfiguration"`
	Tags                  GameSessionQueueTagArrayOutput                 `pulumi:"tags"`
	TimeoutInSeconds      pulumi.IntPtrOutput                            `pulumi:"timeoutInSeconds"`
}

// NewGameSessionQueue registers a new resource with the given unique name, arguments, and options.
func NewGameSessionQueue(ctx *pulumi.Context,
	name string, args *GameSessionQueueArgs, opts ...pulumi.ResourceOption) (*GameSessionQueue, error) {
	if args == nil {
		args = &GameSessionQueueArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource GameSessionQueue
	err := ctx.RegisterResource("aws-native:gamelift:GameSessionQueue", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGameSessionQueue gets an existing GameSessionQueue resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGameSessionQueue(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GameSessionQueueState, opts ...pulumi.ResourceOption) (*GameSessionQueue, error) {
	var resource GameSessionQueue
	err := ctx.ReadResource("aws-native:gamelift:GameSessionQueue", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GameSessionQueue resources.
type gameSessionQueueState struct {
}

type GameSessionQueueState struct {
}

func (GameSessionQueueState) ElementType() reflect.Type {
	return reflect.TypeOf((*gameSessionQueueState)(nil)).Elem()
}

type gameSessionQueueArgs struct {
	CustomEventData       *string                                `pulumi:"customEventData"`
	Destinations          []GameSessionQueueDestination          `pulumi:"destinations"`
	FilterConfiguration   *GameSessionQueueFilterConfiguration   `pulumi:"filterConfiguration"`
	Name                  *string                                `pulumi:"name"`
	NotificationTarget    *string                                `pulumi:"notificationTarget"`
	PlayerLatencyPolicies []GameSessionQueuePlayerLatencyPolicy  `pulumi:"playerLatencyPolicies"`
	PriorityConfiguration *GameSessionQueuePriorityConfiguration `pulumi:"priorityConfiguration"`
	Tags                  []GameSessionQueueTag                  `pulumi:"tags"`
	TimeoutInSeconds      *int                                   `pulumi:"timeoutInSeconds"`
}

// The set of arguments for constructing a GameSessionQueue resource.
type GameSessionQueueArgs struct {
	CustomEventData       pulumi.StringPtrInput
	Destinations          GameSessionQueueDestinationArrayInput
	FilterConfiguration   GameSessionQueueFilterConfigurationPtrInput
	Name                  pulumi.StringPtrInput
	NotificationTarget    pulumi.StringPtrInput
	PlayerLatencyPolicies GameSessionQueuePlayerLatencyPolicyArrayInput
	PriorityConfiguration GameSessionQueuePriorityConfigurationPtrInput
	Tags                  GameSessionQueueTagArrayInput
	TimeoutInSeconds      pulumi.IntPtrInput
}

func (GameSessionQueueArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gameSessionQueueArgs)(nil)).Elem()
}

type GameSessionQueueInput interface {
	pulumi.Input

	ToGameSessionQueueOutput() GameSessionQueueOutput
	ToGameSessionQueueOutputWithContext(ctx context.Context) GameSessionQueueOutput
}

func (*GameSessionQueue) ElementType() reflect.Type {
	return reflect.TypeOf((**GameSessionQueue)(nil)).Elem()
}

func (i *GameSessionQueue) ToGameSessionQueueOutput() GameSessionQueueOutput {
	return i.ToGameSessionQueueOutputWithContext(context.Background())
}

func (i *GameSessionQueue) ToGameSessionQueueOutputWithContext(ctx context.Context) GameSessionQueueOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GameSessionQueueOutput)
}

type GameSessionQueueOutput struct{ *pulumi.OutputState }

func (GameSessionQueueOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GameSessionQueue)(nil)).Elem()
}

func (o GameSessionQueueOutput) ToGameSessionQueueOutput() GameSessionQueueOutput {
	return o
}

func (o GameSessionQueueOutput) ToGameSessionQueueOutputWithContext(ctx context.Context) GameSessionQueueOutput {
	return o
}

func (o GameSessionQueueOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *GameSessionQueue) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

func (o GameSessionQueueOutput) CustomEventData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GameSessionQueue) pulumi.StringPtrOutput { return v.CustomEventData }).(pulumi.StringPtrOutput)
}

func (o GameSessionQueueOutput) Destinations() GameSessionQueueDestinationArrayOutput {
	return o.ApplyT(func(v *GameSessionQueue) GameSessionQueueDestinationArrayOutput { return v.Destinations }).(GameSessionQueueDestinationArrayOutput)
}

func (o GameSessionQueueOutput) FilterConfiguration() GameSessionQueueFilterConfigurationPtrOutput {
	return o.ApplyT(func(v *GameSessionQueue) GameSessionQueueFilterConfigurationPtrOutput { return v.FilterConfiguration }).(GameSessionQueueFilterConfigurationPtrOutput)
}

func (o GameSessionQueueOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *GameSessionQueue) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o GameSessionQueueOutput) NotificationTarget() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GameSessionQueue) pulumi.StringPtrOutput { return v.NotificationTarget }).(pulumi.StringPtrOutput)
}

func (o GameSessionQueueOutput) PlayerLatencyPolicies() GameSessionQueuePlayerLatencyPolicyArrayOutput {
	return o.ApplyT(func(v *GameSessionQueue) GameSessionQueuePlayerLatencyPolicyArrayOutput {
		return v.PlayerLatencyPolicies
	}).(GameSessionQueuePlayerLatencyPolicyArrayOutput)
}

func (o GameSessionQueueOutput) PriorityConfiguration() GameSessionQueuePriorityConfigurationPtrOutput {
	return o.ApplyT(func(v *GameSessionQueue) GameSessionQueuePriorityConfigurationPtrOutput {
		return v.PriorityConfiguration
	}).(GameSessionQueuePriorityConfigurationPtrOutput)
}

func (o GameSessionQueueOutput) Tags() GameSessionQueueTagArrayOutput {
	return o.ApplyT(func(v *GameSessionQueue) GameSessionQueueTagArrayOutput { return v.Tags }).(GameSessionQueueTagArrayOutput)
}

func (o GameSessionQueueOutput) TimeoutInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *GameSessionQueue) pulumi.IntPtrOutput { return v.TimeoutInSeconds }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GameSessionQueueInput)(nil)).Elem(), &GameSessionQueue{})
	pulumi.RegisterOutputType(GameSessionQueueOutput{})
}

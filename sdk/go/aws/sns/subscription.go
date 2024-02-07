// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sns

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::SNS::Subscription
//
// Deprecated: Subscription is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type Subscription struct {
	pulumi.CustomResourceState

	DeliveryPolicy      pulumi.AnyOutput       `pulumi:"deliveryPolicy"`
	Endpoint            pulumi.StringPtrOutput `pulumi:"endpoint"`
	FilterPolicy        pulumi.AnyOutput       `pulumi:"filterPolicy"`
	FilterPolicyScope   pulumi.StringPtrOutput `pulumi:"filterPolicyScope"`
	Protocol            pulumi.StringOutput    `pulumi:"protocol"`
	RawMessageDelivery  pulumi.BoolPtrOutput   `pulumi:"rawMessageDelivery"`
	RedrivePolicy       pulumi.AnyOutput       `pulumi:"redrivePolicy"`
	Region              pulumi.StringPtrOutput `pulumi:"region"`
	ReplayPolicy        pulumi.AnyOutput       `pulumi:"replayPolicy"`
	SubscriptionRoleArn pulumi.StringPtrOutput `pulumi:"subscriptionRoleArn"`
	TopicArn            pulumi.StringOutput    `pulumi:"topicArn"`
}

// NewSubscription registers a new resource with the given unique name, arguments, and options.
func NewSubscription(ctx *pulumi.Context,
	name string, args *SubscriptionArgs, opts ...pulumi.ResourceOption) (*Subscription, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Protocol == nil {
		return nil, errors.New("invalid value for required argument 'Protocol'")
	}
	if args.TopicArn == nil {
		return nil, errors.New("invalid value for required argument 'TopicArn'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"endpoint",
		"protocol",
		"topicArn",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Subscription
	err := ctx.RegisterResource("aws-native:sns:Subscription", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSubscription gets an existing Subscription resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSubscription(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SubscriptionState, opts ...pulumi.ResourceOption) (*Subscription, error) {
	var resource Subscription
	err := ctx.ReadResource("aws-native:sns:Subscription", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Subscription resources.
type subscriptionState struct {
}

type SubscriptionState struct {
}

func (SubscriptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*subscriptionState)(nil)).Elem()
}

type subscriptionArgs struct {
	DeliveryPolicy      interface{} `pulumi:"deliveryPolicy"`
	Endpoint            *string     `pulumi:"endpoint"`
	FilterPolicy        interface{} `pulumi:"filterPolicy"`
	FilterPolicyScope   *string     `pulumi:"filterPolicyScope"`
	Protocol            string      `pulumi:"protocol"`
	RawMessageDelivery  *bool       `pulumi:"rawMessageDelivery"`
	RedrivePolicy       interface{} `pulumi:"redrivePolicy"`
	Region              *string     `pulumi:"region"`
	ReplayPolicy        interface{} `pulumi:"replayPolicy"`
	SubscriptionRoleArn *string     `pulumi:"subscriptionRoleArn"`
	TopicArn            string      `pulumi:"topicArn"`
}

// The set of arguments for constructing a Subscription resource.
type SubscriptionArgs struct {
	DeliveryPolicy      pulumi.Input
	Endpoint            pulumi.StringPtrInput
	FilterPolicy        pulumi.Input
	FilterPolicyScope   pulumi.StringPtrInput
	Protocol            pulumi.StringInput
	RawMessageDelivery  pulumi.BoolPtrInput
	RedrivePolicy       pulumi.Input
	Region              pulumi.StringPtrInput
	ReplayPolicy        pulumi.Input
	SubscriptionRoleArn pulumi.StringPtrInput
	TopicArn            pulumi.StringInput
}

func (SubscriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*subscriptionArgs)(nil)).Elem()
}

type SubscriptionInput interface {
	pulumi.Input

	ToSubscriptionOutput() SubscriptionOutput
	ToSubscriptionOutputWithContext(ctx context.Context) SubscriptionOutput
}

func (*Subscription) ElementType() reflect.Type {
	return reflect.TypeOf((**Subscription)(nil)).Elem()
}

func (i *Subscription) ToSubscriptionOutput() SubscriptionOutput {
	return i.ToSubscriptionOutputWithContext(context.Background())
}

func (i *Subscription) ToSubscriptionOutputWithContext(ctx context.Context) SubscriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionOutput)
}

type SubscriptionOutput struct{ *pulumi.OutputState }

func (SubscriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Subscription)(nil)).Elem()
}

func (o SubscriptionOutput) ToSubscriptionOutput() SubscriptionOutput {
	return o
}

func (o SubscriptionOutput) ToSubscriptionOutputWithContext(ctx context.Context) SubscriptionOutput {
	return o
}

func (o SubscriptionOutput) DeliveryPolicy() pulumi.AnyOutput {
	return o.ApplyT(func(v *Subscription) pulumi.AnyOutput { return v.DeliveryPolicy }).(pulumi.AnyOutput)
}

func (o SubscriptionOutput) Endpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Subscription) pulumi.StringPtrOutput { return v.Endpoint }).(pulumi.StringPtrOutput)
}

func (o SubscriptionOutput) FilterPolicy() pulumi.AnyOutput {
	return o.ApplyT(func(v *Subscription) pulumi.AnyOutput { return v.FilterPolicy }).(pulumi.AnyOutput)
}

func (o SubscriptionOutput) FilterPolicyScope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Subscription) pulumi.StringPtrOutput { return v.FilterPolicyScope }).(pulumi.StringPtrOutput)
}

func (o SubscriptionOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v *Subscription) pulumi.StringOutput { return v.Protocol }).(pulumi.StringOutput)
}

func (o SubscriptionOutput) RawMessageDelivery() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Subscription) pulumi.BoolPtrOutput { return v.RawMessageDelivery }).(pulumi.BoolPtrOutput)
}

func (o SubscriptionOutput) RedrivePolicy() pulumi.AnyOutput {
	return o.ApplyT(func(v *Subscription) pulumi.AnyOutput { return v.RedrivePolicy }).(pulumi.AnyOutput)
}

func (o SubscriptionOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Subscription) pulumi.StringPtrOutput { return v.Region }).(pulumi.StringPtrOutput)
}

func (o SubscriptionOutput) ReplayPolicy() pulumi.AnyOutput {
	return o.ApplyT(func(v *Subscription) pulumi.AnyOutput { return v.ReplayPolicy }).(pulumi.AnyOutput)
}

func (o SubscriptionOutput) SubscriptionRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Subscription) pulumi.StringPtrOutput { return v.SubscriptionRoleArn }).(pulumi.StringPtrOutput)
}

func (o SubscriptionOutput) TopicArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Subscription) pulumi.StringOutput { return v.TopicArn }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SubscriptionInput)(nil)).Elem(), &Subscription{})
	pulumi.RegisterOutputType(SubscriptionOutput{})
}

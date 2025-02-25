// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dms

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::DMS::EventSubscription
//
// Deprecated: EventSubscription is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type EventSubscription struct {
	pulumi.CustomResourceState

	AwsId            pulumi.StringOutput      `pulumi:"awsId"`
	Enabled          pulumi.BoolPtrOutput     `pulumi:"enabled"`
	EventCategories  pulumi.StringArrayOutput `pulumi:"eventCategories"`
	SnsTopicArn      pulumi.StringOutput      `pulumi:"snsTopicArn"`
	SourceIds        pulumi.StringArrayOutput `pulumi:"sourceIds"`
	SourceType       pulumi.StringPtrOutput   `pulumi:"sourceType"`
	SubscriptionName pulumi.StringPtrOutput   `pulumi:"subscriptionName"`
	Tags             aws.TagArrayOutput       `pulumi:"tags"`
}

// NewEventSubscription registers a new resource with the given unique name, arguments, and options.
func NewEventSubscription(ctx *pulumi.Context,
	name string, args *EventSubscriptionArgs, opts ...pulumi.ResourceOption) (*EventSubscription, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SnsTopicArn == nil {
		return nil, errors.New("invalid value for required argument 'SnsTopicArn'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"sourceIds[*]",
		"subscriptionName",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource EventSubscription
	err := ctx.RegisterResource("aws-native:dms:EventSubscription", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEventSubscription gets an existing EventSubscription resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEventSubscription(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EventSubscriptionState, opts ...pulumi.ResourceOption) (*EventSubscription, error) {
	var resource EventSubscription
	err := ctx.ReadResource("aws-native:dms:EventSubscription", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EventSubscription resources.
type eventSubscriptionState struct {
}

type EventSubscriptionState struct {
}

func (EventSubscriptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*eventSubscriptionState)(nil)).Elem()
}

type eventSubscriptionArgs struct {
	Enabled          *bool     `pulumi:"enabled"`
	EventCategories  []string  `pulumi:"eventCategories"`
	SnsTopicArn      string    `pulumi:"snsTopicArn"`
	SourceIds        []string  `pulumi:"sourceIds"`
	SourceType       *string   `pulumi:"sourceType"`
	SubscriptionName *string   `pulumi:"subscriptionName"`
	Tags             []aws.Tag `pulumi:"tags"`
}

// The set of arguments for constructing a EventSubscription resource.
type EventSubscriptionArgs struct {
	Enabled          pulumi.BoolPtrInput
	EventCategories  pulumi.StringArrayInput
	SnsTopicArn      pulumi.StringInput
	SourceIds        pulumi.StringArrayInput
	SourceType       pulumi.StringPtrInput
	SubscriptionName pulumi.StringPtrInput
	Tags             aws.TagArrayInput
}

func (EventSubscriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eventSubscriptionArgs)(nil)).Elem()
}

type EventSubscriptionInput interface {
	pulumi.Input

	ToEventSubscriptionOutput() EventSubscriptionOutput
	ToEventSubscriptionOutputWithContext(ctx context.Context) EventSubscriptionOutput
}

func (*EventSubscription) ElementType() reflect.Type {
	return reflect.TypeOf((**EventSubscription)(nil)).Elem()
}

func (i *EventSubscription) ToEventSubscriptionOutput() EventSubscriptionOutput {
	return i.ToEventSubscriptionOutputWithContext(context.Background())
}

func (i *EventSubscription) ToEventSubscriptionOutputWithContext(ctx context.Context) EventSubscriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventSubscriptionOutput)
}

type EventSubscriptionOutput struct{ *pulumi.OutputState }

func (EventSubscriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EventSubscription)(nil)).Elem()
}

func (o EventSubscriptionOutput) ToEventSubscriptionOutput() EventSubscriptionOutput {
	return o
}

func (o EventSubscriptionOutput) ToEventSubscriptionOutputWithContext(ctx context.Context) EventSubscriptionOutput {
	return o
}

func (o EventSubscriptionOutput) AwsId() pulumi.StringOutput {
	return o.ApplyT(func(v *EventSubscription) pulumi.StringOutput { return v.AwsId }).(pulumi.StringOutput)
}

func (o EventSubscriptionOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EventSubscription) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o EventSubscriptionOutput) EventCategories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *EventSubscription) pulumi.StringArrayOutput { return v.EventCategories }).(pulumi.StringArrayOutput)
}

func (o EventSubscriptionOutput) SnsTopicArn() pulumi.StringOutput {
	return o.ApplyT(func(v *EventSubscription) pulumi.StringOutput { return v.SnsTopicArn }).(pulumi.StringOutput)
}

func (o EventSubscriptionOutput) SourceIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *EventSubscription) pulumi.StringArrayOutput { return v.SourceIds }).(pulumi.StringArrayOutput)
}

func (o EventSubscriptionOutput) SourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventSubscription) pulumi.StringPtrOutput { return v.SourceType }).(pulumi.StringPtrOutput)
}

func (o EventSubscriptionOutput) SubscriptionName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventSubscription) pulumi.StringPtrOutput { return v.SubscriptionName }).(pulumi.StringPtrOutput)
}

func (o EventSubscriptionOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v *EventSubscription) aws.TagArrayOutput { return v.Tags }).(aws.TagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EventSubscriptionInput)(nil)).Elem(), &EventSubscription{})
	pulumi.RegisterOutputType(EventSubscriptionOutput{})
}

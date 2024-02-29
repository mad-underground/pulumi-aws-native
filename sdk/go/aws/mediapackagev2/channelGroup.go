// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mediapackagev2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// <p>Represents a channel group that facilitates the grouping of multiple channels.</p>
type ChannelGroup struct {
	pulumi.CustomResourceState

	// <p>The Amazon Resource Name (ARN) associated with the resource.</p>
	Arn              pulumi.StringOutput `pulumi:"arn"`
	ChannelGroupName pulumi.StringOutput `pulumi:"channelGroupName"`
	// <p>The date and time the channel group was created.</p>
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// <p>Enter any descriptive text that helps you to identify the channel group.</p>
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// <p>The output domain where the source stream should be sent. Integrate the domain with a downstream CDN (such as Amazon CloudFront) or playback device.</p>
	EgressDomain pulumi.StringOutput `pulumi:"egressDomain"`
	// <p>The date and time the channel group was modified.</p>
	ModifiedAt pulumi.StringOutput `pulumi:"modifiedAt"`
	Tags       aws.TagArrayOutput  `pulumi:"tags"`
}

// NewChannelGroup registers a new resource with the given unique name, arguments, and options.
func NewChannelGroup(ctx *pulumi.Context,
	name string, args *ChannelGroupArgs, opts ...pulumi.ResourceOption) (*ChannelGroup, error) {
	if args == nil {
		args = &ChannelGroupArgs{}
	}

	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"channelGroupName",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ChannelGroup
	err := ctx.RegisterResource("aws-native:mediapackagev2:ChannelGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetChannelGroup gets an existing ChannelGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetChannelGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ChannelGroupState, opts ...pulumi.ResourceOption) (*ChannelGroup, error) {
	var resource ChannelGroup
	err := ctx.ReadResource("aws-native:mediapackagev2:ChannelGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ChannelGroup resources.
type channelGroupState struct {
}

type ChannelGroupState struct {
}

func (ChannelGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*channelGroupState)(nil)).Elem()
}

type channelGroupArgs struct {
	ChannelGroupName *string `pulumi:"channelGroupName"`
	// <p>Enter any descriptive text that helps you to identify the channel group.</p>
	Description *string   `pulumi:"description"`
	Tags        []aws.Tag `pulumi:"tags"`
}

// The set of arguments for constructing a ChannelGroup resource.
type ChannelGroupArgs struct {
	ChannelGroupName pulumi.StringPtrInput
	// <p>Enter any descriptive text that helps you to identify the channel group.</p>
	Description pulumi.StringPtrInput
	Tags        aws.TagArrayInput
}

func (ChannelGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*channelGroupArgs)(nil)).Elem()
}

type ChannelGroupInput interface {
	pulumi.Input

	ToChannelGroupOutput() ChannelGroupOutput
	ToChannelGroupOutputWithContext(ctx context.Context) ChannelGroupOutput
}

func (*ChannelGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**ChannelGroup)(nil)).Elem()
}

func (i *ChannelGroup) ToChannelGroupOutput() ChannelGroupOutput {
	return i.ToChannelGroupOutputWithContext(context.Background())
}

func (i *ChannelGroup) ToChannelGroupOutputWithContext(ctx context.Context) ChannelGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ChannelGroupOutput)
}

type ChannelGroupOutput struct{ *pulumi.OutputState }

func (ChannelGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ChannelGroup)(nil)).Elem()
}

func (o ChannelGroupOutput) ToChannelGroupOutput() ChannelGroupOutput {
	return o
}

func (o ChannelGroupOutput) ToChannelGroupOutputWithContext(ctx context.Context) ChannelGroupOutput {
	return o
}

// <p>The Amazon Resource Name (ARN) associated with the resource.</p>
func (o ChannelGroupOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *ChannelGroup) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

func (o ChannelGroupOutput) ChannelGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v *ChannelGroup) pulumi.StringOutput { return v.ChannelGroupName }).(pulumi.StringOutput)
}

// <p>The date and time the channel group was created.</p>
func (o ChannelGroupOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ChannelGroup) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// <p>Enter any descriptive text that helps you to identify the channel group.</p>
func (o ChannelGroupOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ChannelGroup) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// <p>The output domain where the source stream should be sent. Integrate the domain with a downstream CDN (such as Amazon CloudFront) or playback device.</p>
func (o ChannelGroupOutput) EgressDomain() pulumi.StringOutput {
	return o.ApplyT(func(v *ChannelGroup) pulumi.StringOutput { return v.EgressDomain }).(pulumi.StringOutput)
}

// <p>The date and time the channel group was modified.</p>
func (o ChannelGroupOutput) ModifiedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ChannelGroup) pulumi.StringOutput { return v.ModifiedAt }).(pulumi.StringOutput)
}

func (o ChannelGroupOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v *ChannelGroup) aws.TagArrayOutput { return v.Tags }).(aws.TagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ChannelGroupInput)(nil)).Elem(), &ChannelGroup{})
	pulumi.RegisterOutputType(ChannelGroupOutput{})
}

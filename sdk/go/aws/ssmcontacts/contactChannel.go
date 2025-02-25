// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ssmcontacts

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::SSMContacts::ContactChannel
type ContactChannel struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name (ARN) of the engagement to a contact channel.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The details that SSM Incident Manager uses when trying to engage the contact channel.
	ChannelAddress pulumi.StringPtrOutput `pulumi:"channelAddress"`
	// The device name. String of 6 to 50 alphabetical, numeric, dash, and underscore characters.
	ChannelName pulumi.StringPtrOutput `pulumi:"channelName"`
	// Device type, which specify notification channel. Currently supported values: "SMS", "VOICE", "EMAIL", "CHATBOT.
	ChannelType ContactChannelChannelTypePtrOutput `pulumi:"channelType"`
	// ARN of the contact resource
	ContactId pulumi.StringPtrOutput `pulumi:"contactId"`
	// If you want to activate the channel at a later time, you can choose to defer activation. SSM Incident Manager can't engage your contact channel until it has been activated.
	DeferActivation pulumi.BoolPtrOutput `pulumi:"deferActivation"`
}

// NewContactChannel registers a new resource with the given unique name, arguments, and options.
func NewContactChannel(ctx *pulumi.Context,
	name string, args *ContactChannelArgs, opts ...pulumi.ResourceOption) (*ContactChannel, error) {
	if args == nil {
		args = &ContactChannelArgs{}
	}

	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"channelType",
		"contactId",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ContactChannel
	err := ctx.RegisterResource("aws-native:ssmcontacts:ContactChannel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetContactChannel gets an existing ContactChannel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetContactChannel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ContactChannelState, opts ...pulumi.ResourceOption) (*ContactChannel, error) {
	var resource ContactChannel
	err := ctx.ReadResource("aws-native:ssmcontacts:ContactChannel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ContactChannel resources.
type contactChannelState struct {
}

type ContactChannelState struct {
}

func (ContactChannelState) ElementType() reflect.Type {
	return reflect.TypeOf((*contactChannelState)(nil)).Elem()
}

type contactChannelArgs struct {
	// The details that SSM Incident Manager uses when trying to engage the contact channel.
	ChannelAddress *string `pulumi:"channelAddress"`
	// The device name. String of 6 to 50 alphabetical, numeric, dash, and underscore characters.
	ChannelName *string `pulumi:"channelName"`
	// Device type, which specify notification channel. Currently supported values: "SMS", "VOICE", "EMAIL", "CHATBOT.
	ChannelType *ContactChannelChannelType `pulumi:"channelType"`
	// ARN of the contact resource
	ContactId *string `pulumi:"contactId"`
	// If you want to activate the channel at a later time, you can choose to defer activation. SSM Incident Manager can't engage your contact channel until it has been activated.
	DeferActivation *bool `pulumi:"deferActivation"`
}

// The set of arguments for constructing a ContactChannel resource.
type ContactChannelArgs struct {
	// The details that SSM Incident Manager uses when trying to engage the contact channel.
	ChannelAddress pulumi.StringPtrInput
	// The device name. String of 6 to 50 alphabetical, numeric, dash, and underscore characters.
	ChannelName pulumi.StringPtrInput
	// Device type, which specify notification channel. Currently supported values: "SMS", "VOICE", "EMAIL", "CHATBOT.
	ChannelType ContactChannelChannelTypePtrInput
	// ARN of the contact resource
	ContactId pulumi.StringPtrInput
	// If you want to activate the channel at a later time, you can choose to defer activation. SSM Incident Manager can't engage your contact channel until it has been activated.
	DeferActivation pulumi.BoolPtrInput
}

func (ContactChannelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*contactChannelArgs)(nil)).Elem()
}

type ContactChannelInput interface {
	pulumi.Input

	ToContactChannelOutput() ContactChannelOutput
	ToContactChannelOutputWithContext(ctx context.Context) ContactChannelOutput
}

func (*ContactChannel) ElementType() reflect.Type {
	return reflect.TypeOf((**ContactChannel)(nil)).Elem()
}

func (i *ContactChannel) ToContactChannelOutput() ContactChannelOutput {
	return i.ToContactChannelOutputWithContext(context.Background())
}

func (i *ContactChannel) ToContactChannelOutputWithContext(ctx context.Context) ContactChannelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContactChannelOutput)
}

type ContactChannelOutput struct{ *pulumi.OutputState }

func (ContactChannelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContactChannel)(nil)).Elem()
}

func (o ContactChannelOutput) ToContactChannelOutput() ContactChannelOutput {
	return o
}

func (o ContactChannelOutput) ToContactChannelOutputWithContext(ctx context.Context) ContactChannelOutput {
	return o
}

// The Amazon Resource Name (ARN) of the engagement to a contact channel.
func (o ContactChannelOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *ContactChannel) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// The details that SSM Incident Manager uses when trying to engage the contact channel.
func (o ContactChannelOutput) ChannelAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContactChannel) pulumi.StringPtrOutput { return v.ChannelAddress }).(pulumi.StringPtrOutput)
}

// The device name. String of 6 to 50 alphabetical, numeric, dash, and underscore characters.
func (o ContactChannelOutput) ChannelName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContactChannel) pulumi.StringPtrOutput { return v.ChannelName }).(pulumi.StringPtrOutput)
}

// Device type, which specify notification channel. Currently supported values: "SMS", "VOICE", "EMAIL", "CHATBOT.
func (o ContactChannelOutput) ChannelType() ContactChannelChannelTypePtrOutput {
	return o.ApplyT(func(v *ContactChannel) ContactChannelChannelTypePtrOutput { return v.ChannelType }).(ContactChannelChannelTypePtrOutput)
}

// ARN of the contact resource
func (o ContactChannelOutput) ContactId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContactChannel) pulumi.StringPtrOutput { return v.ContactId }).(pulumi.StringPtrOutput)
}

// If you want to activate the channel at a later time, you can choose to defer activation. SSM Incident Manager can't engage your contact channel until it has been activated.
func (o ContactChannelOutput) DeferActivation() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ContactChannel) pulumi.BoolPtrOutput { return v.DeferActivation }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ContactChannelInput)(nil)).Elem(), &ContactChannel{})
	pulumi.RegisterOutputType(ContactChannelOutput{})
}

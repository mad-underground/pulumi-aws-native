// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package pinpoint

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Pinpoint::SMSChannel
//
// Deprecated: SmsChannel is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type SmsChannel struct {
	pulumi.CustomResourceState

	ApplicationId pulumi.StringOutput    `pulumi:"applicationId"`
	AwsId         pulumi.StringOutput    `pulumi:"awsId"`
	Enabled       pulumi.BoolPtrOutput   `pulumi:"enabled"`
	SenderId      pulumi.StringPtrOutput `pulumi:"senderId"`
	ShortCode     pulumi.StringPtrOutput `pulumi:"shortCode"`
}

// NewSmsChannel registers a new resource with the given unique name, arguments, and options.
func NewSmsChannel(ctx *pulumi.Context,
	name string, args *SmsChannelArgs, opts ...pulumi.ResourceOption) (*SmsChannel, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApplicationId == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"applicationId",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SmsChannel
	err := ctx.RegisterResource("aws-native:pinpoint:SmsChannel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSmsChannel gets an existing SmsChannel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSmsChannel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SmsChannelState, opts ...pulumi.ResourceOption) (*SmsChannel, error) {
	var resource SmsChannel
	err := ctx.ReadResource("aws-native:pinpoint:SmsChannel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SmsChannel resources.
type smsChannelState struct {
}

type SmsChannelState struct {
}

func (SmsChannelState) ElementType() reflect.Type {
	return reflect.TypeOf((*smsChannelState)(nil)).Elem()
}

type smsChannelArgs struct {
	ApplicationId string  `pulumi:"applicationId"`
	Enabled       *bool   `pulumi:"enabled"`
	SenderId      *string `pulumi:"senderId"`
	ShortCode     *string `pulumi:"shortCode"`
}

// The set of arguments for constructing a SmsChannel resource.
type SmsChannelArgs struct {
	ApplicationId pulumi.StringInput
	Enabled       pulumi.BoolPtrInput
	SenderId      pulumi.StringPtrInput
	ShortCode     pulumi.StringPtrInput
}

func (SmsChannelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*smsChannelArgs)(nil)).Elem()
}

type SmsChannelInput interface {
	pulumi.Input

	ToSmsChannelOutput() SmsChannelOutput
	ToSmsChannelOutputWithContext(ctx context.Context) SmsChannelOutput
}

func (*SmsChannel) ElementType() reflect.Type {
	return reflect.TypeOf((**SmsChannel)(nil)).Elem()
}

func (i *SmsChannel) ToSmsChannelOutput() SmsChannelOutput {
	return i.ToSmsChannelOutputWithContext(context.Background())
}

func (i *SmsChannel) ToSmsChannelOutputWithContext(ctx context.Context) SmsChannelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SmsChannelOutput)
}

type SmsChannelOutput struct{ *pulumi.OutputState }

func (SmsChannelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SmsChannel)(nil)).Elem()
}

func (o SmsChannelOutput) ToSmsChannelOutput() SmsChannelOutput {
	return o
}

func (o SmsChannelOutput) ToSmsChannelOutputWithContext(ctx context.Context) SmsChannelOutput {
	return o
}

func (o SmsChannelOutput) ApplicationId() pulumi.StringOutput {
	return o.ApplyT(func(v *SmsChannel) pulumi.StringOutput { return v.ApplicationId }).(pulumi.StringOutput)
}

func (o SmsChannelOutput) AwsId() pulumi.StringOutput {
	return o.ApplyT(func(v *SmsChannel) pulumi.StringOutput { return v.AwsId }).(pulumi.StringOutput)
}

func (o SmsChannelOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SmsChannel) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o SmsChannelOutput) SenderId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SmsChannel) pulumi.StringPtrOutput { return v.SenderId }).(pulumi.StringPtrOutput)
}

func (o SmsChannelOutput) ShortCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SmsChannel) pulumi.StringPtrOutput { return v.ShortCode }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SmsChannelInput)(nil)).Elem(), &SmsChannel{})
	pulumi.RegisterOutputType(SmsChannelOutput{})
}

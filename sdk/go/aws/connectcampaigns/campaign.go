// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package connectcampaigns

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of AWS::ConnectCampaigns::Campaign Resource Type
type Campaign struct {
	pulumi.CustomResourceState

	// Amazon Connect Campaign Arn
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Amazon Connect Instance Arn
	ConnectInstanceArn pulumi.StringOutput        `pulumi:"connectInstanceArn"`
	DialerConfig       CampaignDialerConfigOutput `pulumi:"dialerConfig"`
	// Amazon Connect Campaign Name
	Name               pulumi.StringOutput              `pulumi:"name"`
	OutboundCallConfig CampaignOutboundCallConfigOutput `pulumi:"outboundCallConfig"`
	// One or more tags.
	Tags CampaignTagArrayOutput `pulumi:"tags"`
}

// NewCampaign registers a new resource with the given unique name, arguments, and options.
func NewCampaign(ctx *pulumi.Context,
	name string, args *CampaignArgs, opts ...pulumi.ResourceOption) (*Campaign, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConnectInstanceArn == nil {
		return nil, errors.New("invalid value for required argument 'ConnectInstanceArn'")
	}
	if args.DialerConfig == nil {
		return nil, errors.New("invalid value for required argument 'DialerConfig'")
	}
	if args.OutboundCallConfig == nil {
		return nil, errors.New("invalid value for required argument 'OutboundCallConfig'")
	}
	var resource Campaign
	err := ctx.RegisterResource("aws-native:connectcampaigns:Campaign", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCampaign gets an existing Campaign resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCampaign(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CampaignState, opts ...pulumi.ResourceOption) (*Campaign, error) {
	var resource Campaign
	err := ctx.ReadResource("aws-native:connectcampaigns:Campaign", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Campaign resources.
type campaignState struct {
}

type CampaignState struct {
}

func (CampaignState) ElementType() reflect.Type {
	return reflect.TypeOf((*campaignState)(nil)).Elem()
}

type campaignArgs struct {
	// Amazon Connect Instance Arn
	ConnectInstanceArn string               `pulumi:"connectInstanceArn"`
	DialerConfig       CampaignDialerConfig `pulumi:"dialerConfig"`
	// Amazon Connect Campaign Name
	Name               *string                    `pulumi:"name"`
	OutboundCallConfig CampaignOutboundCallConfig `pulumi:"outboundCallConfig"`
	// One or more tags.
	Tags []CampaignTag `pulumi:"tags"`
}

// The set of arguments for constructing a Campaign resource.
type CampaignArgs struct {
	// Amazon Connect Instance Arn
	ConnectInstanceArn pulumi.StringInput
	DialerConfig       CampaignDialerConfigInput
	// Amazon Connect Campaign Name
	Name               pulumi.StringPtrInput
	OutboundCallConfig CampaignOutboundCallConfigInput
	// One or more tags.
	Tags CampaignTagArrayInput
}

func (CampaignArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*campaignArgs)(nil)).Elem()
}

type CampaignInput interface {
	pulumi.Input

	ToCampaignOutput() CampaignOutput
	ToCampaignOutputWithContext(ctx context.Context) CampaignOutput
}

func (*Campaign) ElementType() reflect.Type {
	return reflect.TypeOf((**Campaign)(nil)).Elem()
}

func (i *Campaign) ToCampaignOutput() CampaignOutput {
	return i.ToCampaignOutputWithContext(context.Background())
}

func (i *Campaign) ToCampaignOutputWithContext(ctx context.Context) CampaignOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CampaignOutput)
}

type CampaignOutput struct{ *pulumi.OutputState }

func (CampaignOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Campaign)(nil)).Elem()
}

func (o CampaignOutput) ToCampaignOutput() CampaignOutput {
	return o
}

func (o CampaignOutput) ToCampaignOutputWithContext(ctx context.Context) CampaignOutput {
	return o
}

// Amazon Connect Campaign Arn
func (o CampaignOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Campaign) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Amazon Connect Instance Arn
func (o CampaignOutput) ConnectInstanceArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Campaign) pulumi.StringOutput { return v.ConnectInstanceArn }).(pulumi.StringOutput)
}

func (o CampaignOutput) DialerConfig() CampaignDialerConfigOutput {
	return o.ApplyT(func(v *Campaign) CampaignDialerConfigOutput { return v.DialerConfig }).(CampaignDialerConfigOutput)
}

// Amazon Connect Campaign Name
func (o CampaignOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Campaign) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o CampaignOutput) OutboundCallConfig() CampaignOutboundCallConfigOutput {
	return o.ApplyT(func(v *Campaign) CampaignOutboundCallConfigOutput { return v.OutboundCallConfig }).(CampaignOutboundCallConfigOutput)
}

// One or more tags.
func (o CampaignOutput) Tags() CampaignTagArrayOutput {
	return o.ApplyT(func(v *Campaign) CampaignTagArrayOutput { return v.Tags }).(CampaignTagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CampaignInput)(nil)).Elem(), &Campaign{})
	pulumi.RegisterOutputType(CampaignOutput{})
}
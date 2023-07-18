// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package pinpointemail

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::PinpointEmail::ConfigurationSet
//
// Deprecated: ConfigurationSet is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type ConfigurationSet struct {
	pulumi.CustomResourceState

	DeliveryOptions   ConfigurationSetDeliveryOptionsPtrOutput   `pulumi:"deliveryOptions"`
	Name              pulumi.StringOutput                        `pulumi:"name"`
	ReputationOptions ConfigurationSetReputationOptionsPtrOutput `pulumi:"reputationOptions"`
	SendingOptions    ConfigurationSetSendingOptionsPtrOutput    `pulumi:"sendingOptions"`
	Tags              ConfigurationSetTagsArrayOutput            `pulumi:"tags"`
	TrackingOptions   ConfigurationSetTrackingOptionsPtrOutput   `pulumi:"trackingOptions"`
}

// NewConfigurationSet registers a new resource with the given unique name, arguments, and options.
func NewConfigurationSet(ctx *pulumi.Context,
	name string, args *ConfigurationSetArgs, opts ...pulumi.ResourceOption) (*ConfigurationSet, error) {
	if args == nil {
		args = &ConfigurationSetArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ConfigurationSet
	err := ctx.RegisterResource("aws-native:pinpointemail:ConfigurationSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConfigurationSet gets an existing ConfigurationSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConfigurationSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigurationSetState, opts ...pulumi.ResourceOption) (*ConfigurationSet, error) {
	var resource ConfigurationSet
	err := ctx.ReadResource("aws-native:pinpointemail:ConfigurationSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConfigurationSet resources.
type configurationSetState struct {
}

type ConfigurationSetState struct {
}

func (ConfigurationSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationSetState)(nil)).Elem()
}

type configurationSetArgs struct {
	DeliveryOptions   *ConfigurationSetDeliveryOptions   `pulumi:"deliveryOptions"`
	Name              *string                            `pulumi:"name"`
	ReputationOptions *ConfigurationSetReputationOptions `pulumi:"reputationOptions"`
	SendingOptions    *ConfigurationSetSendingOptions    `pulumi:"sendingOptions"`
	Tags              []ConfigurationSetTags             `pulumi:"tags"`
	TrackingOptions   *ConfigurationSetTrackingOptions   `pulumi:"trackingOptions"`
}

// The set of arguments for constructing a ConfigurationSet resource.
type ConfigurationSetArgs struct {
	DeliveryOptions   ConfigurationSetDeliveryOptionsPtrInput
	Name              pulumi.StringPtrInput
	ReputationOptions ConfigurationSetReputationOptionsPtrInput
	SendingOptions    ConfigurationSetSendingOptionsPtrInput
	Tags              ConfigurationSetTagsArrayInput
	TrackingOptions   ConfigurationSetTrackingOptionsPtrInput
}

func (ConfigurationSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationSetArgs)(nil)).Elem()
}

type ConfigurationSetInput interface {
	pulumi.Input

	ToConfigurationSetOutput() ConfigurationSetOutput
	ToConfigurationSetOutputWithContext(ctx context.Context) ConfigurationSetOutput
}

func (*ConfigurationSet) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationSet)(nil)).Elem()
}

func (i *ConfigurationSet) ToConfigurationSetOutput() ConfigurationSetOutput {
	return i.ToConfigurationSetOutputWithContext(context.Background())
}

func (i *ConfigurationSet) ToConfigurationSetOutputWithContext(ctx context.Context) ConfigurationSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationSetOutput)
}

type ConfigurationSetOutput struct{ *pulumi.OutputState }

func (ConfigurationSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationSet)(nil)).Elem()
}

func (o ConfigurationSetOutput) ToConfigurationSetOutput() ConfigurationSetOutput {
	return o
}

func (o ConfigurationSetOutput) ToConfigurationSetOutputWithContext(ctx context.Context) ConfigurationSetOutput {
	return o
}

func (o ConfigurationSetOutput) DeliveryOptions() ConfigurationSetDeliveryOptionsPtrOutput {
	return o.ApplyT(func(v *ConfigurationSet) ConfigurationSetDeliveryOptionsPtrOutput { return v.DeliveryOptions }).(ConfigurationSetDeliveryOptionsPtrOutput)
}

func (o ConfigurationSetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ConfigurationSet) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ConfigurationSetOutput) ReputationOptions() ConfigurationSetReputationOptionsPtrOutput {
	return o.ApplyT(func(v *ConfigurationSet) ConfigurationSetReputationOptionsPtrOutput { return v.ReputationOptions }).(ConfigurationSetReputationOptionsPtrOutput)
}

func (o ConfigurationSetOutput) SendingOptions() ConfigurationSetSendingOptionsPtrOutput {
	return o.ApplyT(func(v *ConfigurationSet) ConfigurationSetSendingOptionsPtrOutput { return v.SendingOptions }).(ConfigurationSetSendingOptionsPtrOutput)
}

func (o ConfigurationSetOutput) Tags() ConfigurationSetTagsArrayOutput {
	return o.ApplyT(func(v *ConfigurationSet) ConfigurationSetTagsArrayOutput { return v.Tags }).(ConfigurationSetTagsArrayOutput)
}

func (o ConfigurationSetOutput) TrackingOptions() ConfigurationSetTrackingOptionsPtrOutput {
	return o.ApplyT(func(v *ConfigurationSet) ConfigurationSetTrackingOptionsPtrOutput { return v.TrackingOptions }).(ConfigurationSetTrackingOptionsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigurationSetInput)(nil)).Elem(), &ConfigurationSet{})
	pulumi.RegisterOutputType(ConfigurationSetOutput{})
}

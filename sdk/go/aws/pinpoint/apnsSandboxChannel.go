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

// Resource Type definition for AWS::Pinpoint::APNSSandboxChannel
//
// Deprecated: ApnsSandboxChannel is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type ApnsSandboxChannel struct {
	pulumi.CustomResourceState

	ApplicationId               pulumi.StringOutput    `pulumi:"applicationId"`
	AwsId                       pulumi.StringOutput    `pulumi:"awsId"`
	BundleId                    pulumi.StringPtrOutput `pulumi:"bundleId"`
	Certificate                 pulumi.StringPtrOutput `pulumi:"certificate"`
	DefaultAuthenticationMethod pulumi.StringPtrOutput `pulumi:"defaultAuthenticationMethod"`
	Enabled                     pulumi.BoolPtrOutput   `pulumi:"enabled"`
	PrivateKey                  pulumi.StringPtrOutput `pulumi:"privateKey"`
	TeamId                      pulumi.StringPtrOutput `pulumi:"teamId"`
	TokenKey                    pulumi.StringPtrOutput `pulumi:"tokenKey"`
	TokenKeyId                  pulumi.StringPtrOutput `pulumi:"tokenKeyId"`
}

// NewApnsSandboxChannel registers a new resource with the given unique name, arguments, and options.
func NewApnsSandboxChannel(ctx *pulumi.Context,
	name string, args *ApnsSandboxChannelArgs, opts ...pulumi.ResourceOption) (*ApnsSandboxChannel, error) {
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
	var resource ApnsSandboxChannel
	err := ctx.RegisterResource("aws-native:pinpoint:ApnsSandboxChannel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApnsSandboxChannel gets an existing ApnsSandboxChannel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApnsSandboxChannel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApnsSandboxChannelState, opts ...pulumi.ResourceOption) (*ApnsSandboxChannel, error) {
	var resource ApnsSandboxChannel
	err := ctx.ReadResource("aws-native:pinpoint:ApnsSandboxChannel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApnsSandboxChannel resources.
type apnsSandboxChannelState struct {
}

type ApnsSandboxChannelState struct {
}

func (ApnsSandboxChannelState) ElementType() reflect.Type {
	return reflect.TypeOf((*apnsSandboxChannelState)(nil)).Elem()
}

type apnsSandboxChannelArgs struct {
	ApplicationId               string  `pulumi:"applicationId"`
	BundleId                    *string `pulumi:"bundleId"`
	Certificate                 *string `pulumi:"certificate"`
	DefaultAuthenticationMethod *string `pulumi:"defaultAuthenticationMethod"`
	Enabled                     *bool   `pulumi:"enabled"`
	PrivateKey                  *string `pulumi:"privateKey"`
	TeamId                      *string `pulumi:"teamId"`
	TokenKey                    *string `pulumi:"tokenKey"`
	TokenKeyId                  *string `pulumi:"tokenKeyId"`
}

// The set of arguments for constructing a ApnsSandboxChannel resource.
type ApnsSandboxChannelArgs struct {
	ApplicationId               pulumi.StringInput
	BundleId                    pulumi.StringPtrInput
	Certificate                 pulumi.StringPtrInput
	DefaultAuthenticationMethod pulumi.StringPtrInput
	Enabled                     pulumi.BoolPtrInput
	PrivateKey                  pulumi.StringPtrInput
	TeamId                      pulumi.StringPtrInput
	TokenKey                    pulumi.StringPtrInput
	TokenKeyId                  pulumi.StringPtrInput
}

func (ApnsSandboxChannelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apnsSandboxChannelArgs)(nil)).Elem()
}

type ApnsSandboxChannelInput interface {
	pulumi.Input

	ToApnsSandboxChannelOutput() ApnsSandboxChannelOutput
	ToApnsSandboxChannelOutputWithContext(ctx context.Context) ApnsSandboxChannelOutput
}

func (*ApnsSandboxChannel) ElementType() reflect.Type {
	return reflect.TypeOf((**ApnsSandboxChannel)(nil)).Elem()
}

func (i *ApnsSandboxChannel) ToApnsSandboxChannelOutput() ApnsSandboxChannelOutput {
	return i.ToApnsSandboxChannelOutputWithContext(context.Background())
}

func (i *ApnsSandboxChannel) ToApnsSandboxChannelOutputWithContext(ctx context.Context) ApnsSandboxChannelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApnsSandboxChannelOutput)
}

type ApnsSandboxChannelOutput struct{ *pulumi.OutputState }

func (ApnsSandboxChannelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApnsSandboxChannel)(nil)).Elem()
}

func (o ApnsSandboxChannelOutput) ToApnsSandboxChannelOutput() ApnsSandboxChannelOutput {
	return o
}

func (o ApnsSandboxChannelOutput) ToApnsSandboxChannelOutputWithContext(ctx context.Context) ApnsSandboxChannelOutput {
	return o
}

func (o ApnsSandboxChannelOutput) ApplicationId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApnsSandboxChannel) pulumi.StringOutput { return v.ApplicationId }).(pulumi.StringOutput)
}

func (o ApnsSandboxChannelOutput) AwsId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApnsSandboxChannel) pulumi.StringOutput { return v.AwsId }).(pulumi.StringOutput)
}

func (o ApnsSandboxChannelOutput) BundleId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApnsSandboxChannel) pulumi.StringPtrOutput { return v.BundleId }).(pulumi.StringPtrOutput)
}

func (o ApnsSandboxChannelOutput) Certificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApnsSandboxChannel) pulumi.StringPtrOutput { return v.Certificate }).(pulumi.StringPtrOutput)
}

func (o ApnsSandboxChannelOutput) DefaultAuthenticationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApnsSandboxChannel) pulumi.StringPtrOutput { return v.DefaultAuthenticationMethod }).(pulumi.StringPtrOutput)
}

func (o ApnsSandboxChannelOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApnsSandboxChannel) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o ApnsSandboxChannelOutput) PrivateKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApnsSandboxChannel) pulumi.StringPtrOutput { return v.PrivateKey }).(pulumi.StringPtrOutput)
}

func (o ApnsSandboxChannelOutput) TeamId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApnsSandboxChannel) pulumi.StringPtrOutput { return v.TeamId }).(pulumi.StringPtrOutput)
}

func (o ApnsSandboxChannelOutput) TokenKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApnsSandboxChannel) pulumi.StringPtrOutput { return v.TokenKey }).(pulumi.StringPtrOutput)
}

func (o ApnsSandboxChannelOutput) TokenKeyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApnsSandboxChannel) pulumi.StringPtrOutput { return v.TokenKeyId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApnsSandboxChannelInput)(nil)).Elem(), &ApnsSandboxChannel{})
	pulumi.RegisterOutputType(ApnsSandboxChannelOutput{})
}

// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ivs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::IVS::PlaybackKeyPair
type PlaybackKeyPair struct {
	pulumi.CustomResourceState

	// Key-pair identifier.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Key-pair identifier.
	Fingerprint pulumi.StringOutput `pulumi:"fingerprint"`
	// An arbitrary string (a nickname) assigned to a playback key pair that helps the customer identify that resource. The value does not need to be unique.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// The public portion of a customer-generated key pair.
	PublicKeyMaterial pulumi.StringPtrOutput `pulumi:"publicKeyMaterial"`
	// A list of key-value pairs that contain metadata for the asset model.
	Tags PlaybackKeyPairTagArrayOutput `pulumi:"tags"`
}

// NewPlaybackKeyPair registers a new resource with the given unique name, arguments, and options.
func NewPlaybackKeyPair(ctx *pulumi.Context,
	name string, args *PlaybackKeyPairArgs, opts ...pulumi.ResourceOption) (*PlaybackKeyPair, error) {
	if args == nil {
		args = &PlaybackKeyPairArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PlaybackKeyPair
	err := ctx.RegisterResource("aws-native:ivs:PlaybackKeyPair", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPlaybackKeyPair gets an existing PlaybackKeyPair resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPlaybackKeyPair(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PlaybackKeyPairState, opts ...pulumi.ResourceOption) (*PlaybackKeyPair, error) {
	var resource PlaybackKeyPair
	err := ctx.ReadResource("aws-native:ivs:PlaybackKeyPair", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PlaybackKeyPair resources.
type playbackKeyPairState struct {
}

type PlaybackKeyPairState struct {
}

func (PlaybackKeyPairState) ElementType() reflect.Type {
	return reflect.TypeOf((*playbackKeyPairState)(nil)).Elem()
}

type playbackKeyPairArgs struct {
	// An arbitrary string (a nickname) assigned to a playback key pair that helps the customer identify that resource. The value does not need to be unique.
	Name *string `pulumi:"name"`
	// The public portion of a customer-generated key pair.
	PublicKeyMaterial *string `pulumi:"publicKeyMaterial"`
	// A list of key-value pairs that contain metadata for the asset model.
	Tags []PlaybackKeyPairTag `pulumi:"tags"`
}

// The set of arguments for constructing a PlaybackKeyPair resource.
type PlaybackKeyPairArgs struct {
	// An arbitrary string (a nickname) assigned to a playback key pair that helps the customer identify that resource. The value does not need to be unique.
	Name pulumi.StringPtrInput
	// The public portion of a customer-generated key pair.
	PublicKeyMaterial pulumi.StringPtrInput
	// A list of key-value pairs that contain metadata for the asset model.
	Tags PlaybackKeyPairTagArrayInput
}

func (PlaybackKeyPairArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*playbackKeyPairArgs)(nil)).Elem()
}

type PlaybackKeyPairInput interface {
	pulumi.Input

	ToPlaybackKeyPairOutput() PlaybackKeyPairOutput
	ToPlaybackKeyPairOutputWithContext(ctx context.Context) PlaybackKeyPairOutput
}

func (*PlaybackKeyPair) ElementType() reflect.Type {
	return reflect.TypeOf((**PlaybackKeyPair)(nil)).Elem()
}

func (i *PlaybackKeyPair) ToPlaybackKeyPairOutput() PlaybackKeyPairOutput {
	return i.ToPlaybackKeyPairOutputWithContext(context.Background())
}

func (i *PlaybackKeyPair) ToPlaybackKeyPairOutputWithContext(ctx context.Context) PlaybackKeyPairOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlaybackKeyPairOutput)
}

type PlaybackKeyPairOutput struct{ *pulumi.OutputState }

func (PlaybackKeyPairOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PlaybackKeyPair)(nil)).Elem()
}

func (o PlaybackKeyPairOutput) ToPlaybackKeyPairOutput() PlaybackKeyPairOutput {
	return o
}

func (o PlaybackKeyPairOutput) ToPlaybackKeyPairOutputWithContext(ctx context.Context) PlaybackKeyPairOutput {
	return o
}

// Key-pair identifier.
func (o PlaybackKeyPairOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *PlaybackKeyPair) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Key-pair identifier.
func (o PlaybackKeyPairOutput) Fingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v *PlaybackKeyPair) pulumi.StringOutput { return v.Fingerprint }).(pulumi.StringOutput)
}

// An arbitrary string (a nickname) assigned to a playback key pair that helps the customer identify that resource. The value does not need to be unique.
func (o PlaybackKeyPairOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PlaybackKeyPair) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// The public portion of a customer-generated key pair.
func (o PlaybackKeyPairOutput) PublicKeyMaterial() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PlaybackKeyPair) pulumi.StringPtrOutput { return v.PublicKeyMaterial }).(pulumi.StringPtrOutput)
}

// A list of key-value pairs that contain metadata for the asset model.
func (o PlaybackKeyPairOutput) Tags() PlaybackKeyPairTagArrayOutput {
	return o.ApplyT(func(v *PlaybackKeyPair) PlaybackKeyPairTagArrayOutput { return v.Tags }).(PlaybackKeyPairTagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PlaybackKeyPairInput)(nil)).Elem(), &PlaybackKeyPair{})
	pulumi.RegisterOutputType(PlaybackKeyPairOutput{})
}

// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iotfleetwise

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of AWS::IoTFleetWise::DecoderManifest Resource Type
//
// Deprecated: DecoderManifest is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type DecoderManifest struct {
	pulumi.CustomResourceState

	Arn                  pulumi.StringOutput                    `pulumi:"arn"`
	CreationTime         pulumi.StringOutput                    `pulumi:"creationTime"`
	Description          pulumi.StringPtrOutput                 `pulumi:"description"`
	LastModificationTime pulumi.StringOutput                    `pulumi:"lastModificationTime"`
	ModelManifestArn     pulumi.StringOutput                    `pulumi:"modelManifestArn"`
	Name                 pulumi.StringOutput                    `pulumi:"name"`
	NetworkInterfaces    pulumi.ArrayOutput                     `pulumi:"networkInterfaces"`
	SignalDecoders       pulumi.ArrayOutput                     `pulumi:"signalDecoders"`
	Status               DecoderManifestManifestStatusPtrOutput `pulumi:"status"`
	Tags                 DecoderManifestTagArrayOutput          `pulumi:"tags"`
}

// NewDecoderManifest registers a new resource with the given unique name, arguments, and options.
func NewDecoderManifest(ctx *pulumi.Context,
	name string, args *DecoderManifestArgs, opts ...pulumi.ResourceOption) (*DecoderManifest, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ModelManifestArn == nil {
		return nil, errors.New("invalid value for required argument 'ModelManifestArn'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DecoderManifest
	err := ctx.RegisterResource("aws-native:iotfleetwise:DecoderManifest", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDecoderManifest gets an existing DecoderManifest resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDecoderManifest(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DecoderManifestState, opts ...pulumi.ResourceOption) (*DecoderManifest, error) {
	var resource DecoderManifest
	err := ctx.ReadResource("aws-native:iotfleetwise:DecoderManifest", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DecoderManifest resources.
type decoderManifestState struct {
}

type DecoderManifestState struct {
}

func (DecoderManifestState) ElementType() reflect.Type {
	return reflect.TypeOf((*decoderManifestState)(nil)).Elem()
}

type decoderManifestArgs struct {
	Description       *string                        `pulumi:"description"`
	ModelManifestArn  string                         `pulumi:"modelManifestArn"`
	Name              *string                        `pulumi:"name"`
	NetworkInterfaces []interface{}                  `pulumi:"networkInterfaces"`
	SignalDecoders    []interface{}                  `pulumi:"signalDecoders"`
	Status            *DecoderManifestManifestStatus `pulumi:"status"`
	Tags              []DecoderManifestTag           `pulumi:"tags"`
}

// The set of arguments for constructing a DecoderManifest resource.
type DecoderManifestArgs struct {
	Description       pulumi.StringPtrInput
	ModelManifestArn  pulumi.StringInput
	Name              pulumi.StringPtrInput
	NetworkInterfaces pulumi.ArrayInput
	SignalDecoders    pulumi.ArrayInput
	Status            DecoderManifestManifestStatusPtrInput
	Tags              DecoderManifestTagArrayInput
}

func (DecoderManifestArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*decoderManifestArgs)(nil)).Elem()
}

type DecoderManifestInput interface {
	pulumi.Input

	ToDecoderManifestOutput() DecoderManifestOutput
	ToDecoderManifestOutputWithContext(ctx context.Context) DecoderManifestOutput
}

func (*DecoderManifest) ElementType() reflect.Type {
	return reflect.TypeOf((**DecoderManifest)(nil)).Elem()
}

func (i *DecoderManifest) ToDecoderManifestOutput() DecoderManifestOutput {
	return i.ToDecoderManifestOutputWithContext(context.Background())
}

func (i *DecoderManifest) ToDecoderManifestOutputWithContext(ctx context.Context) DecoderManifestOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DecoderManifestOutput)
}

type DecoderManifestOutput struct{ *pulumi.OutputState }

func (DecoderManifestOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DecoderManifest)(nil)).Elem()
}

func (o DecoderManifestOutput) ToDecoderManifestOutput() DecoderManifestOutput {
	return o
}

func (o DecoderManifestOutput) ToDecoderManifestOutputWithContext(ctx context.Context) DecoderManifestOutput {
	return o
}

func (o DecoderManifestOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *DecoderManifest) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

func (o DecoderManifestOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *DecoderManifest) pulumi.StringOutput { return v.CreationTime }).(pulumi.StringOutput)
}

func (o DecoderManifestOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DecoderManifest) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o DecoderManifestOutput) LastModificationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *DecoderManifest) pulumi.StringOutput { return v.LastModificationTime }).(pulumi.StringOutput)
}

func (o DecoderManifestOutput) ModelManifestArn() pulumi.StringOutput {
	return o.ApplyT(func(v *DecoderManifest) pulumi.StringOutput { return v.ModelManifestArn }).(pulumi.StringOutput)
}

func (o DecoderManifestOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DecoderManifest) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DecoderManifestOutput) NetworkInterfaces() pulumi.ArrayOutput {
	return o.ApplyT(func(v *DecoderManifest) pulumi.ArrayOutput { return v.NetworkInterfaces }).(pulumi.ArrayOutput)
}

func (o DecoderManifestOutput) SignalDecoders() pulumi.ArrayOutput {
	return o.ApplyT(func(v *DecoderManifest) pulumi.ArrayOutput { return v.SignalDecoders }).(pulumi.ArrayOutput)
}

func (o DecoderManifestOutput) Status() DecoderManifestManifestStatusPtrOutput {
	return o.ApplyT(func(v *DecoderManifest) DecoderManifestManifestStatusPtrOutput { return v.Status }).(DecoderManifestManifestStatusPtrOutput)
}

func (o DecoderManifestOutput) Tags() DecoderManifestTagArrayOutput {
	return o.ApplyT(func(v *DecoderManifest) DecoderManifestTagArrayOutput { return v.Tags }).(DecoderManifestTagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DecoderManifestInput)(nil)).Elem(), &DecoderManifest{})
	pulumi.RegisterOutputType(DecoderManifestOutput{})
}

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

// Definition of AWS::IoTFleetWise::ModelManifest Resource Type
//
// Deprecated: ModelManifest is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type ModelManifest struct {
	pulumi.CustomResourceState

	Arn                  pulumi.StringOutput                  `pulumi:"arn"`
	CreationTime         pulumi.StringOutput                  `pulumi:"creationTime"`
	Description          pulumi.StringPtrOutput               `pulumi:"description"`
	LastModificationTime pulumi.StringOutput                  `pulumi:"lastModificationTime"`
	Name                 pulumi.StringOutput                  `pulumi:"name"`
	Nodes                pulumi.StringArrayOutput             `pulumi:"nodes"`
	SignalCatalogArn     pulumi.StringOutput                  `pulumi:"signalCatalogArn"`
	Status               ModelManifestManifestStatusPtrOutput `pulumi:"status"`
	Tags                 ModelManifestTagArrayOutput          `pulumi:"tags"`
}

// NewModelManifest registers a new resource with the given unique name, arguments, and options.
func NewModelManifest(ctx *pulumi.Context,
	name string, args *ModelManifestArgs, opts ...pulumi.ResourceOption) (*ModelManifest, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SignalCatalogArn == nil {
		return nil, errors.New("invalid value for required argument 'SignalCatalogArn'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"name",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ModelManifest
	err := ctx.RegisterResource("aws-native:iotfleetwise:ModelManifest", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetModelManifest gets an existing ModelManifest resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetModelManifest(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ModelManifestState, opts ...pulumi.ResourceOption) (*ModelManifest, error) {
	var resource ModelManifest
	err := ctx.ReadResource("aws-native:iotfleetwise:ModelManifest", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ModelManifest resources.
type modelManifestState struct {
}

type ModelManifestState struct {
}

func (ModelManifestState) ElementType() reflect.Type {
	return reflect.TypeOf((*modelManifestState)(nil)).Elem()
}

type modelManifestArgs struct {
	Description      *string                      `pulumi:"description"`
	Name             *string                      `pulumi:"name"`
	Nodes            []string                     `pulumi:"nodes"`
	SignalCatalogArn string                       `pulumi:"signalCatalogArn"`
	Status           *ModelManifestManifestStatus `pulumi:"status"`
	Tags             []ModelManifestTag           `pulumi:"tags"`
}

// The set of arguments for constructing a ModelManifest resource.
type ModelManifestArgs struct {
	Description      pulumi.StringPtrInput
	Name             pulumi.StringPtrInput
	Nodes            pulumi.StringArrayInput
	SignalCatalogArn pulumi.StringInput
	Status           ModelManifestManifestStatusPtrInput
	Tags             ModelManifestTagArrayInput
}

func (ModelManifestArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*modelManifestArgs)(nil)).Elem()
}

type ModelManifestInput interface {
	pulumi.Input

	ToModelManifestOutput() ModelManifestOutput
	ToModelManifestOutputWithContext(ctx context.Context) ModelManifestOutput
}

func (*ModelManifest) ElementType() reflect.Type {
	return reflect.TypeOf((**ModelManifest)(nil)).Elem()
}

func (i *ModelManifest) ToModelManifestOutput() ModelManifestOutput {
	return i.ToModelManifestOutputWithContext(context.Background())
}

func (i *ModelManifest) ToModelManifestOutputWithContext(ctx context.Context) ModelManifestOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModelManifestOutput)
}

type ModelManifestOutput struct{ *pulumi.OutputState }

func (ModelManifestOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ModelManifest)(nil)).Elem()
}

func (o ModelManifestOutput) ToModelManifestOutput() ModelManifestOutput {
	return o
}

func (o ModelManifestOutput) ToModelManifestOutputWithContext(ctx context.Context) ModelManifestOutput {
	return o
}

func (o ModelManifestOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *ModelManifest) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

func (o ModelManifestOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ModelManifest) pulumi.StringOutput { return v.CreationTime }).(pulumi.StringOutput)
}

func (o ModelManifestOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ModelManifest) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ModelManifestOutput) LastModificationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ModelManifest) pulumi.StringOutput { return v.LastModificationTime }).(pulumi.StringOutput)
}

func (o ModelManifestOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ModelManifest) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ModelManifestOutput) Nodes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ModelManifest) pulumi.StringArrayOutput { return v.Nodes }).(pulumi.StringArrayOutput)
}

func (o ModelManifestOutput) SignalCatalogArn() pulumi.StringOutput {
	return o.ApplyT(func(v *ModelManifest) pulumi.StringOutput { return v.SignalCatalogArn }).(pulumi.StringOutput)
}

func (o ModelManifestOutput) Status() ModelManifestManifestStatusPtrOutput {
	return o.ApplyT(func(v *ModelManifest) ModelManifestManifestStatusPtrOutput { return v.Status }).(ModelManifestManifestStatusPtrOutput)
}

func (o ModelManifestOutput) Tags() ModelManifestTagArrayOutput {
	return o.ApplyT(func(v *ModelManifest) ModelManifestTagArrayOutput { return v.Tags }).(ModelManifestTagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ModelManifestInput)(nil)).Elem(), &ModelManifest{})
	pulumi.RegisterOutputType(ModelManifestOutput{})
}

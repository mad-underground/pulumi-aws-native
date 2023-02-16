// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package oam

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of AWS::Oam::Link Resource Type
type Link struct {
	pulumi.CustomResourceState

	Arn            pulumi.StringOutput         `pulumi:"arn"`
	Label          pulumi.StringOutput         `pulumi:"label"`
	LabelTemplate  pulumi.StringPtrOutput      `pulumi:"labelTemplate"`
	ResourceTypes  LinkResourceTypeArrayOutput `pulumi:"resourceTypes"`
	SinkIdentifier pulumi.StringOutput         `pulumi:"sinkIdentifier"`
	// Tags to apply to the link
	Tags pulumi.AnyOutput `pulumi:"tags"`
}

// NewLink registers a new resource with the given unique name, arguments, and options.
func NewLink(ctx *pulumi.Context,
	name string, args *LinkArgs, opts ...pulumi.ResourceOption) (*Link, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceTypes == nil {
		return nil, errors.New("invalid value for required argument 'ResourceTypes'")
	}
	if args.SinkIdentifier == nil {
		return nil, errors.New("invalid value for required argument 'SinkIdentifier'")
	}
	var resource Link
	err := ctx.RegisterResource("aws-native:oam:Link", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLink gets an existing Link resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLink(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LinkState, opts ...pulumi.ResourceOption) (*Link, error) {
	var resource Link
	err := ctx.ReadResource("aws-native:oam:Link", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Link resources.
type linkState struct {
}

type LinkState struct {
}

func (LinkState) ElementType() reflect.Type {
	return reflect.TypeOf((*linkState)(nil)).Elem()
}

type linkArgs struct {
	LabelTemplate  *string            `pulumi:"labelTemplate"`
	ResourceTypes  []LinkResourceType `pulumi:"resourceTypes"`
	SinkIdentifier string             `pulumi:"sinkIdentifier"`
	// Tags to apply to the link
	Tags interface{} `pulumi:"tags"`
}

// The set of arguments for constructing a Link resource.
type LinkArgs struct {
	LabelTemplate  pulumi.StringPtrInput
	ResourceTypes  LinkResourceTypeArrayInput
	SinkIdentifier pulumi.StringInput
	// Tags to apply to the link
	Tags pulumi.Input
}

func (LinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*linkArgs)(nil)).Elem()
}

type LinkInput interface {
	pulumi.Input

	ToLinkOutput() LinkOutput
	ToLinkOutputWithContext(ctx context.Context) LinkOutput
}

func (*Link) ElementType() reflect.Type {
	return reflect.TypeOf((**Link)(nil)).Elem()
}

func (i *Link) ToLinkOutput() LinkOutput {
	return i.ToLinkOutputWithContext(context.Background())
}

func (i *Link) ToLinkOutputWithContext(ctx context.Context) LinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkOutput)
}

type LinkOutput struct{ *pulumi.OutputState }

func (LinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Link)(nil)).Elem()
}

func (o LinkOutput) ToLinkOutput() LinkOutput {
	return o
}

func (o LinkOutput) ToLinkOutputWithContext(ctx context.Context) LinkOutput {
	return o
}

func (o LinkOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Link) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

func (o LinkOutput) Label() pulumi.StringOutput {
	return o.ApplyT(func(v *Link) pulumi.StringOutput { return v.Label }).(pulumi.StringOutput)
}

func (o LinkOutput) LabelTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Link) pulumi.StringPtrOutput { return v.LabelTemplate }).(pulumi.StringPtrOutput)
}

func (o LinkOutput) ResourceTypes() LinkResourceTypeArrayOutput {
	return o.ApplyT(func(v *Link) LinkResourceTypeArrayOutput { return v.ResourceTypes }).(LinkResourceTypeArrayOutput)
}

func (o LinkOutput) SinkIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v *Link) pulumi.StringOutput { return v.SinkIdentifier }).(pulumi.StringOutput)
}

// Tags to apply to the link
func (o LinkOutput) Tags() pulumi.AnyOutput {
	return o.ApplyT(func(v *Link) pulumi.AnyOutput { return v.Tags }).(pulumi.AnyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LinkInput)(nil)).Elem(), &Link{})
	pulumi.RegisterOutputType(LinkOutput{})
}

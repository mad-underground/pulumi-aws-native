// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package evidently

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource Type definition for AWS::Evidently::Segment
type Segment struct {
	pulumi.CustomResourceState

	Arn         pulumi.StringOutput    `pulumi:"arn"`
	Description pulumi.StringPtrOutput `pulumi:"description"`
	Name        pulumi.StringOutput    `pulumi:"name"`
	Pattern     pulumi.StringPtrOutput `pulumi:"pattern"`
	// An array of key-value pairs to apply to this resource.
	Tags SegmentTagArrayOutput `pulumi:"tags"`
}

// NewSegment registers a new resource with the given unique name, arguments, and options.
func NewSegment(ctx *pulumi.Context,
	name string, args *SegmentArgs, opts ...pulumi.ResourceOption) (*Segment, error) {
	if args == nil {
		args = &SegmentArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Segment
	err := ctx.RegisterResource("aws-native:evidently:Segment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSegment gets an existing Segment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSegment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SegmentState, opts ...pulumi.ResourceOption) (*Segment, error) {
	var resource Segment
	err := ctx.ReadResource("aws-native:evidently:Segment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Segment resources.
type segmentState struct {
}

type SegmentState struct {
}

func (SegmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*segmentState)(nil)).Elem()
}

type segmentArgs struct {
	Description *string `pulumi:"description"`
	Name        *string `pulumi:"name"`
	Pattern     *string `pulumi:"pattern"`
	// An array of key-value pairs to apply to this resource.
	Tags []SegmentTag `pulumi:"tags"`
}

// The set of arguments for constructing a Segment resource.
type SegmentArgs struct {
	Description pulumi.StringPtrInput
	Name        pulumi.StringPtrInput
	Pattern     pulumi.StringPtrInput
	// An array of key-value pairs to apply to this resource.
	Tags SegmentTagArrayInput
}

func (SegmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*segmentArgs)(nil)).Elem()
}

type SegmentInput interface {
	pulumi.Input

	ToSegmentOutput() SegmentOutput
	ToSegmentOutputWithContext(ctx context.Context) SegmentOutput
}

func (*Segment) ElementType() reflect.Type {
	return reflect.TypeOf((**Segment)(nil)).Elem()
}

func (i *Segment) ToSegmentOutput() SegmentOutput {
	return i.ToSegmentOutputWithContext(context.Background())
}

func (i *Segment) ToSegmentOutputWithContext(ctx context.Context) SegmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SegmentOutput)
}

type SegmentOutput struct{ *pulumi.OutputState }

func (SegmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Segment)(nil)).Elem()
}

func (o SegmentOutput) ToSegmentOutput() SegmentOutput {
	return o
}

func (o SegmentOutput) ToSegmentOutputWithContext(ctx context.Context) SegmentOutput {
	return o
}

func (o SegmentOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Segment) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

func (o SegmentOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Segment) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o SegmentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Segment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SegmentOutput) Pattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Segment) pulumi.StringPtrOutput { return v.Pattern }).(pulumi.StringPtrOutput)
}

// An array of key-value pairs to apply to this resource.
func (o SegmentOutput) Tags() SegmentTagArrayOutput {
	return o.ApplyT(func(v *Segment) SegmentTagArrayOutput { return v.Tags }).(SegmentTagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SegmentInput)(nil)).Elem(), &Segment{})
	pulumi.RegisterOutputType(SegmentOutput{})
}

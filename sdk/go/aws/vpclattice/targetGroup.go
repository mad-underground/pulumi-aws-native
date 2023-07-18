// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpclattice

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A target group is a collection of targets, or compute resources, that run your application or service. A target group can only be used by a single service.
type TargetGroup struct {
	pulumi.CustomResourceState

	Arn           pulumi.StringOutput          `pulumi:"arn"`
	Config        TargetGroupConfigPtrOutput   `pulumi:"config"`
	CreatedAt     pulumi.StringOutput          `pulumi:"createdAt"`
	LastUpdatedAt pulumi.StringOutput          `pulumi:"lastUpdatedAt"`
	Name          pulumi.StringPtrOutput       `pulumi:"name"`
	Status        TargetGroupStatusOutput      `pulumi:"status"`
	Tags          TargetGroupTagArrayOutput    `pulumi:"tags"`
	Targets       TargetGroupTargetArrayOutput `pulumi:"targets"`
	Type          TargetGroupTypeOutput        `pulumi:"type"`
}

// NewTargetGroup registers a new resource with the given unique name, arguments, and options.
func NewTargetGroup(ctx *pulumi.Context,
	name string, args *TargetGroupArgs, opts ...pulumi.ResourceOption) (*TargetGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource TargetGroup
	err := ctx.RegisterResource("aws-native:vpclattice:TargetGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTargetGroup gets an existing TargetGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTargetGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TargetGroupState, opts ...pulumi.ResourceOption) (*TargetGroup, error) {
	var resource TargetGroup
	err := ctx.ReadResource("aws-native:vpclattice:TargetGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TargetGroup resources.
type targetGroupState struct {
}

type TargetGroupState struct {
}

func (TargetGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*targetGroupState)(nil)).Elem()
}

type targetGroupArgs struct {
	Config  *TargetGroupConfig  `pulumi:"config"`
	Name    *string             `pulumi:"name"`
	Tags    []TargetGroupTag    `pulumi:"tags"`
	Targets []TargetGroupTarget `pulumi:"targets"`
	Type    TargetGroupType     `pulumi:"type"`
}

// The set of arguments for constructing a TargetGroup resource.
type TargetGroupArgs struct {
	Config  TargetGroupConfigPtrInput
	Name    pulumi.StringPtrInput
	Tags    TargetGroupTagArrayInput
	Targets TargetGroupTargetArrayInput
	Type    TargetGroupTypeInput
}

func (TargetGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*targetGroupArgs)(nil)).Elem()
}

type TargetGroupInput interface {
	pulumi.Input

	ToTargetGroupOutput() TargetGroupOutput
	ToTargetGroupOutputWithContext(ctx context.Context) TargetGroupOutput
}

func (*TargetGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**TargetGroup)(nil)).Elem()
}

func (i *TargetGroup) ToTargetGroupOutput() TargetGroupOutput {
	return i.ToTargetGroupOutputWithContext(context.Background())
}

func (i *TargetGroup) ToTargetGroupOutputWithContext(ctx context.Context) TargetGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetGroupOutput)
}

type TargetGroupOutput struct{ *pulumi.OutputState }

func (TargetGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TargetGroup)(nil)).Elem()
}

func (o TargetGroupOutput) ToTargetGroupOutput() TargetGroupOutput {
	return o
}

func (o TargetGroupOutput) ToTargetGroupOutputWithContext(ctx context.Context) TargetGroupOutput {
	return o
}

func (o TargetGroupOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *TargetGroup) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

func (o TargetGroupOutput) Config() TargetGroupConfigPtrOutput {
	return o.ApplyT(func(v *TargetGroup) TargetGroupConfigPtrOutput { return v.Config }).(TargetGroupConfigPtrOutput)
}

func (o TargetGroupOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *TargetGroup) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o TargetGroupOutput) LastUpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *TargetGroup) pulumi.StringOutput { return v.LastUpdatedAt }).(pulumi.StringOutput)
}

func (o TargetGroupOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TargetGroup) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o TargetGroupOutput) Status() TargetGroupStatusOutput {
	return o.ApplyT(func(v *TargetGroup) TargetGroupStatusOutput { return v.Status }).(TargetGroupStatusOutput)
}

func (o TargetGroupOutput) Tags() TargetGroupTagArrayOutput {
	return o.ApplyT(func(v *TargetGroup) TargetGroupTagArrayOutput { return v.Tags }).(TargetGroupTagArrayOutput)
}

func (o TargetGroupOutput) Targets() TargetGroupTargetArrayOutput {
	return o.ApplyT(func(v *TargetGroup) TargetGroupTargetArrayOutput { return v.Targets }).(TargetGroupTargetArrayOutput)
}

func (o TargetGroupOutput) Type() TargetGroupTypeOutput {
	return o.ApplyT(func(v *TargetGroup) TargetGroupTypeOutput { return v.Type }).(TargetGroupTypeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TargetGroupInput)(nil)).Elem(), &TargetGroup{})
	pulumi.RegisterOutputType(TargetGroupOutput{})
}

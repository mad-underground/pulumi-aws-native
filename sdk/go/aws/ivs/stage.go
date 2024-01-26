// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ivs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Resource Definition for type AWS::IVS::Stage.
type Stage struct {
	pulumi.CustomResourceState

	// ID of the active session within the stage.
	ActiveSessionId pulumi.StringOutput `pulumi:"activeSessionId"`
	// Stage ARN is automatically generated on creation and assigned as the unique identifier.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Stage name
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// An array of key-value pairs to apply to this resource.
	Tags StageTagArrayOutput `pulumi:"tags"`
}

// NewStage registers a new resource with the given unique name, arguments, and options.
func NewStage(ctx *pulumi.Context,
	name string, args *StageArgs, opts ...pulumi.ResourceOption) (*Stage, error) {
	if args == nil {
		args = &StageArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Stage
	err := ctx.RegisterResource("aws-native:ivs:Stage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStage gets an existing Stage resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StageState, opts ...pulumi.ResourceOption) (*Stage, error) {
	var resource Stage
	err := ctx.ReadResource("aws-native:ivs:Stage", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Stage resources.
type stageState struct {
}

type StageState struct {
}

func (StageState) ElementType() reflect.Type {
	return reflect.TypeOf((*stageState)(nil)).Elem()
}

type stageArgs struct {
	// Stage name
	Name *string `pulumi:"name"`
	// An array of key-value pairs to apply to this resource.
	Tags []StageTag `pulumi:"tags"`
}

// The set of arguments for constructing a Stage resource.
type StageArgs struct {
	// Stage name
	Name pulumi.StringPtrInput
	// An array of key-value pairs to apply to this resource.
	Tags StageTagArrayInput
}

func (StageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*stageArgs)(nil)).Elem()
}

type StageInput interface {
	pulumi.Input

	ToStageOutput() StageOutput
	ToStageOutputWithContext(ctx context.Context) StageOutput
}

func (*Stage) ElementType() reflect.Type {
	return reflect.TypeOf((**Stage)(nil)).Elem()
}

func (i *Stage) ToStageOutput() StageOutput {
	return i.ToStageOutputWithContext(context.Background())
}

func (i *Stage) ToStageOutputWithContext(ctx context.Context) StageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StageOutput)
}

func (i *Stage) ToOutput(ctx context.Context) pulumix.Output[*Stage] {
	return pulumix.Output[*Stage]{
		OutputState: i.ToStageOutputWithContext(ctx).OutputState,
	}
}

type StageOutput struct{ *pulumi.OutputState }

func (StageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Stage)(nil)).Elem()
}

func (o StageOutput) ToStageOutput() StageOutput {
	return o
}

func (o StageOutput) ToStageOutputWithContext(ctx context.Context) StageOutput {
	return o
}

func (o StageOutput) ToOutput(ctx context.Context) pulumix.Output[*Stage] {
	return pulumix.Output[*Stage]{
		OutputState: o.OutputState,
	}
}

// ID of the active session within the stage.
func (o StageOutput) ActiveSessionId() pulumi.StringOutput {
	return o.ApplyT(func(v *Stage) pulumi.StringOutput { return v.ActiveSessionId }).(pulumi.StringOutput)
}

// Stage ARN is automatically generated on creation and assigned as the unique identifier.
func (o StageOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Stage) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

// Stage name
func (o StageOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Stage) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

// An array of key-value pairs to apply to this resource.
func (o StageOutput) Tags() StageTagArrayOutput {
	return o.ApplyT(func(v *Stage) StageTagArrayOutput { return v.Tags }).(StageTagArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StageInput)(nil)).Elem(), &Stage{})
	pulumi.RegisterOutputType(StageOutput{})
}

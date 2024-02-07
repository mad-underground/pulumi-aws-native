// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package controltower

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Enables a control on a specified target.
type EnabledControl struct {
	pulumi.CustomResourceState

	// Arn of the control.
	ControlIdentifier pulumi.StringOutput `pulumi:"controlIdentifier"`
	// Parameters to configure the enabled control behavior.
	Parameters EnabledControlParameterArrayOutput `pulumi:"parameters"`
	// Arn for Organizational unit to which the control needs to be applied
	TargetIdentifier pulumi.StringOutput `pulumi:"targetIdentifier"`
}

// NewEnabledControl registers a new resource with the given unique name, arguments, and options.
func NewEnabledControl(ctx *pulumi.Context,
	name string, args *EnabledControlArgs, opts ...pulumi.ResourceOption) (*EnabledControl, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ControlIdentifier == nil {
		return nil, errors.New("invalid value for required argument 'ControlIdentifier'")
	}
	if args.TargetIdentifier == nil {
		return nil, errors.New("invalid value for required argument 'TargetIdentifier'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"controlIdentifier",
		"targetIdentifier",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource EnabledControl
	err := ctx.RegisterResource("aws-native:controltower:EnabledControl", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEnabledControl gets an existing EnabledControl resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEnabledControl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EnabledControlState, opts ...pulumi.ResourceOption) (*EnabledControl, error) {
	var resource EnabledControl
	err := ctx.ReadResource("aws-native:controltower:EnabledControl", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EnabledControl resources.
type enabledControlState struct {
}

type EnabledControlState struct {
}

func (EnabledControlState) ElementType() reflect.Type {
	return reflect.TypeOf((*enabledControlState)(nil)).Elem()
}

type enabledControlArgs struct {
	// Arn of the control.
	ControlIdentifier string `pulumi:"controlIdentifier"`
	// Parameters to configure the enabled control behavior.
	Parameters []EnabledControlParameter `pulumi:"parameters"`
	// Arn for Organizational unit to which the control needs to be applied
	TargetIdentifier string `pulumi:"targetIdentifier"`
}

// The set of arguments for constructing a EnabledControl resource.
type EnabledControlArgs struct {
	// Arn of the control.
	ControlIdentifier pulumi.StringInput
	// Parameters to configure the enabled control behavior.
	Parameters EnabledControlParameterArrayInput
	// Arn for Organizational unit to which the control needs to be applied
	TargetIdentifier pulumi.StringInput
}

func (EnabledControlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*enabledControlArgs)(nil)).Elem()
}

type EnabledControlInput interface {
	pulumi.Input

	ToEnabledControlOutput() EnabledControlOutput
	ToEnabledControlOutputWithContext(ctx context.Context) EnabledControlOutput
}

func (*EnabledControl) ElementType() reflect.Type {
	return reflect.TypeOf((**EnabledControl)(nil)).Elem()
}

func (i *EnabledControl) ToEnabledControlOutput() EnabledControlOutput {
	return i.ToEnabledControlOutputWithContext(context.Background())
}

func (i *EnabledControl) ToEnabledControlOutputWithContext(ctx context.Context) EnabledControlOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnabledControlOutput)
}

type EnabledControlOutput struct{ *pulumi.OutputState }

func (EnabledControlOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnabledControl)(nil)).Elem()
}

func (o EnabledControlOutput) ToEnabledControlOutput() EnabledControlOutput {
	return o
}

func (o EnabledControlOutput) ToEnabledControlOutputWithContext(ctx context.Context) EnabledControlOutput {
	return o
}

// Arn of the control.
func (o EnabledControlOutput) ControlIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v *EnabledControl) pulumi.StringOutput { return v.ControlIdentifier }).(pulumi.StringOutput)
}

// Parameters to configure the enabled control behavior.
func (o EnabledControlOutput) Parameters() EnabledControlParameterArrayOutput {
	return o.ApplyT(func(v *EnabledControl) EnabledControlParameterArrayOutput { return v.Parameters }).(EnabledControlParameterArrayOutput)
}

// Arn for Organizational unit to which the control needs to be applied
func (o EnabledControlOutput) TargetIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v *EnabledControl) pulumi.StringOutput { return v.TargetIdentifier }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EnabledControlInput)(nil)).Elem(), &EnabledControl{})
	pulumi.RegisterOutputType(EnabledControlOutput{})
}

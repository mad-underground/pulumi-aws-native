// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package controltower

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Definition of AWS::ControlTower::EnabledBaseline Resource Type
type EnabledBaseline struct {
	pulumi.CustomResourceState

	BaselineIdentifier        pulumi.StringOutput                 `pulumi:"baselineIdentifier"`
	BaselineVersion           pulumi.StringOutput                 `pulumi:"baselineVersion"`
	EnabledBaselineIdentifier pulumi.StringOutput                 `pulumi:"enabledBaselineIdentifier"`
	Parameters                EnabledBaselineParameterArrayOutput `pulumi:"parameters"`
	Tags                      aws.TagArrayOutput                  `pulumi:"tags"`
	TargetIdentifier          pulumi.StringOutput                 `pulumi:"targetIdentifier"`
}

// NewEnabledBaseline registers a new resource with the given unique name, arguments, and options.
func NewEnabledBaseline(ctx *pulumi.Context,
	name string, args *EnabledBaselineArgs, opts ...pulumi.ResourceOption) (*EnabledBaseline, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BaselineIdentifier == nil {
		return nil, errors.New("invalid value for required argument 'BaselineIdentifier'")
	}
	if args.BaselineVersion == nil {
		return nil, errors.New("invalid value for required argument 'BaselineVersion'")
	}
	if args.TargetIdentifier == nil {
		return nil, errors.New("invalid value for required argument 'TargetIdentifier'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"baselineIdentifier",
		"targetIdentifier",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource EnabledBaseline
	err := ctx.RegisterResource("aws-native:controltower:EnabledBaseline", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEnabledBaseline gets an existing EnabledBaseline resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEnabledBaseline(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EnabledBaselineState, opts ...pulumi.ResourceOption) (*EnabledBaseline, error) {
	var resource EnabledBaseline
	err := ctx.ReadResource("aws-native:controltower:EnabledBaseline", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EnabledBaseline resources.
type enabledBaselineState struct {
}

type EnabledBaselineState struct {
}

func (EnabledBaselineState) ElementType() reflect.Type {
	return reflect.TypeOf((*enabledBaselineState)(nil)).Elem()
}

type enabledBaselineArgs struct {
	BaselineIdentifier string                     `pulumi:"baselineIdentifier"`
	BaselineVersion    string                     `pulumi:"baselineVersion"`
	Parameters         []EnabledBaselineParameter `pulumi:"parameters"`
	Tags               []aws.Tag                  `pulumi:"tags"`
	TargetIdentifier   string                     `pulumi:"targetIdentifier"`
}

// The set of arguments for constructing a EnabledBaseline resource.
type EnabledBaselineArgs struct {
	BaselineIdentifier pulumi.StringInput
	BaselineVersion    pulumi.StringInput
	Parameters         EnabledBaselineParameterArrayInput
	Tags               aws.TagArrayInput
	TargetIdentifier   pulumi.StringInput
}

func (EnabledBaselineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*enabledBaselineArgs)(nil)).Elem()
}

type EnabledBaselineInput interface {
	pulumi.Input

	ToEnabledBaselineOutput() EnabledBaselineOutput
	ToEnabledBaselineOutputWithContext(ctx context.Context) EnabledBaselineOutput
}

func (*EnabledBaseline) ElementType() reflect.Type {
	return reflect.TypeOf((**EnabledBaseline)(nil)).Elem()
}

func (i *EnabledBaseline) ToEnabledBaselineOutput() EnabledBaselineOutput {
	return i.ToEnabledBaselineOutputWithContext(context.Background())
}

func (i *EnabledBaseline) ToEnabledBaselineOutputWithContext(ctx context.Context) EnabledBaselineOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnabledBaselineOutput)
}

type EnabledBaselineOutput struct{ *pulumi.OutputState }

func (EnabledBaselineOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnabledBaseline)(nil)).Elem()
}

func (o EnabledBaselineOutput) ToEnabledBaselineOutput() EnabledBaselineOutput {
	return o
}

func (o EnabledBaselineOutput) ToEnabledBaselineOutputWithContext(ctx context.Context) EnabledBaselineOutput {
	return o
}

func (o EnabledBaselineOutput) BaselineIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v *EnabledBaseline) pulumi.StringOutput { return v.BaselineIdentifier }).(pulumi.StringOutput)
}

func (o EnabledBaselineOutput) BaselineVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *EnabledBaseline) pulumi.StringOutput { return v.BaselineVersion }).(pulumi.StringOutput)
}

func (o EnabledBaselineOutput) EnabledBaselineIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v *EnabledBaseline) pulumi.StringOutput { return v.EnabledBaselineIdentifier }).(pulumi.StringOutput)
}

func (o EnabledBaselineOutput) Parameters() EnabledBaselineParameterArrayOutput {
	return o.ApplyT(func(v *EnabledBaseline) EnabledBaselineParameterArrayOutput { return v.Parameters }).(EnabledBaselineParameterArrayOutput)
}

func (o EnabledBaselineOutput) Tags() aws.TagArrayOutput {
	return o.ApplyT(func(v *EnabledBaseline) aws.TagArrayOutput { return v.Tags }).(aws.TagArrayOutput)
}

func (o EnabledBaselineOutput) TargetIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v *EnabledBaseline) pulumi.StringOutput { return v.TargetIdentifier }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EnabledBaselineInput)(nil)).Elem(), &EnabledBaseline{})
	pulumi.RegisterOutputType(EnabledBaselineOutput{})
}

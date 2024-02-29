// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package controltower

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

type EnabledBaselineParameter struct {
	Key   *string     `pulumi:"key"`
	Value interface{} `pulumi:"value"`
}

// EnabledBaselineParameterInput is an input type that accepts EnabledBaselineParameterArgs and EnabledBaselineParameterOutput values.
// You can construct a concrete instance of `EnabledBaselineParameterInput` via:
//
//	EnabledBaselineParameterArgs{...}
type EnabledBaselineParameterInput interface {
	pulumi.Input

	ToEnabledBaselineParameterOutput() EnabledBaselineParameterOutput
	ToEnabledBaselineParameterOutputWithContext(context.Context) EnabledBaselineParameterOutput
}

type EnabledBaselineParameterArgs struct {
	Key   pulumi.StringPtrInput `pulumi:"key"`
	Value pulumi.Input          `pulumi:"value"`
}

func (EnabledBaselineParameterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EnabledBaselineParameter)(nil)).Elem()
}

func (i EnabledBaselineParameterArgs) ToEnabledBaselineParameterOutput() EnabledBaselineParameterOutput {
	return i.ToEnabledBaselineParameterOutputWithContext(context.Background())
}

func (i EnabledBaselineParameterArgs) ToEnabledBaselineParameterOutputWithContext(ctx context.Context) EnabledBaselineParameterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnabledBaselineParameterOutput)
}

// EnabledBaselineParameterArrayInput is an input type that accepts EnabledBaselineParameterArray and EnabledBaselineParameterArrayOutput values.
// You can construct a concrete instance of `EnabledBaselineParameterArrayInput` via:
//
//	EnabledBaselineParameterArray{ EnabledBaselineParameterArgs{...} }
type EnabledBaselineParameterArrayInput interface {
	pulumi.Input

	ToEnabledBaselineParameterArrayOutput() EnabledBaselineParameterArrayOutput
	ToEnabledBaselineParameterArrayOutputWithContext(context.Context) EnabledBaselineParameterArrayOutput
}

type EnabledBaselineParameterArray []EnabledBaselineParameterInput

func (EnabledBaselineParameterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EnabledBaselineParameter)(nil)).Elem()
}

func (i EnabledBaselineParameterArray) ToEnabledBaselineParameterArrayOutput() EnabledBaselineParameterArrayOutput {
	return i.ToEnabledBaselineParameterArrayOutputWithContext(context.Background())
}

func (i EnabledBaselineParameterArray) ToEnabledBaselineParameterArrayOutputWithContext(ctx context.Context) EnabledBaselineParameterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnabledBaselineParameterArrayOutput)
}

type EnabledBaselineParameterOutput struct{ *pulumi.OutputState }

func (EnabledBaselineParameterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnabledBaselineParameter)(nil)).Elem()
}

func (o EnabledBaselineParameterOutput) ToEnabledBaselineParameterOutput() EnabledBaselineParameterOutput {
	return o
}

func (o EnabledBaselineParameterOutput) ToEnabledBaselineParameterOutputWithContext(ctx context.Context) EnabledBaselineParameterOutput {
	return o
}

func (o EnabledBaselineParameterOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnabledBaselineParameter) *string { return v.Key }).(pulumi.StringPtrOutput)
}

func (o EnabledBaselineParameterOutput) Value() pulumi.AnyOutput {
	return o.ApplyT(func(v EnabledBaselineParameter) interface{} { return v.Value }).(pulumi.AnyOutput)
}

type EnabledBaselineParameterArrayOutput struct{ *pulumi.OutputState }

func (EnabledBaselineParameterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EnabledBaselineParameter)(nil)).Elem()
}

func (o EnabledBaselineParameterArrayOutput) ToEnabledBaselineParameterArrayOutput() EnabledBaselineParameterArrayOutput {
	return o
}

func (o EnabledBaselineParameterArrayOutput) ToEnabledBaselineParameterArrayOutputWithContext(ctx context.Context) EnabledBaselineParameterArrayOutput {
	return o
}

func (o EnabledBaselineParameterArrayOutput) Index(i pulumi.IntInput) EnabledBaselineParameterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EnabledBaselineParameter {
		return vs[0].([]EnabledBaselineParameter)[vs[1].(int)]
	}).(EnabledBaselineParameterOutput)
}

type EnabledBaselineTag struct {
	Key   *string `pulumi:"key"`
	Value *string `pulumi:"value"`
}

type EnabledControlParameter struct {
	Key   string      `pulumi:"key"`
	Value interface{} `pulumi:"value"`
}

// EnabledControlParameterInput is an input type that accepts EnabledControlParameterArgs and EnabledControlParameterOutput values.
// You can construct a concrete instance of `EnabledControlParameterInput` via:
//
//	EnabledControlParameterArgs{...}
type EnabledControlParameterInput interface {
	pulumi.Input

	ToEnabledControlParameterOutput() EnabledControlParameterOutput
	ToEnabledControlParameterOutputWithContext(context.Context) EnabledControlParameterOutput
}

type EnabledControlParameterArgs struct {
	Key   pulumi.StringInput `pulumi:"key"`
	Value pulumi.Input       `pulumi:"value"`
}

func (EnabledControlParameterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EnabledControlParameter)(nil)).Elem()
}

func (i EnabledControlParameterArgs) ToEnabledControlParameterOutput() EnabledControlParameterOutput {
	return i.ToEnabledControlParameterOutputWithContext(context.Background())
}

func (i EnabledControlParameterArgs) ToEnabledControlParameterOutputWithContext(ctx context.Context) EnabledControlParameterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnabledControlParameterOutput)
}

// EnabledControlParameterArrayInput is an input type that accepts EnabledControlParameterArray and EnabledControlParameterArrayOutput values.
// You can construct a concrete instance of `EnabledControlParameterArrayInput` via:
//
//	EnabledControlParameterArray{ EnabledControlParameterArgs{...} }
type EnabledControlParameterArrayInput interface {
	pulumi.Input

	ToEnabledControlParameterArrayOutput() EnabledControlParameterArrayOutput
	ToEnabledControlParameterArrayOutputWithContext(context.Context) EnabledControlParameterArrayOutput
}

type EnabledControlParameterArray []EnabledControlParameterInput

func (EnabledControlParameterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EnabledControlParameter)(nil)).Elem()
}

func (i EnabledControlParameterArray) ToEnabledControlParameterArrayOutput() EnabledControlParameterArrayOutput {
	return i.ToEnabledControlParameterArrayOutputWithContext(context.Background())
}

func (i EnabledControlParameterArray) ToEnabledControlParameterArrayOutputWithContext(ctx context.Context) EnabledControlParameterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnabledControlParameterArrayOutput)
}

type EnabledControlParameterOutput struct{ *pulumi.OutputState }

func (EnabledControlParameterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnabledControlParameter)(nil)).Elem()
}

func (o EnabledControlParameterOutput) ToEnabledControlParameterOutput() EnabledControlParameterOutput {
	return o
}

func (o EnabledControlParameterOutput) ToEnabledControlParameterOutputWithContext(ctx context.Context) EnabledControlParameterOutput {
	return o
}

func (o EnabledControlParameterOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v EnabledControlParameter) string { return v.Key }).(pulumi.StringOutput)
}

func (o EnabledControlParameterOutput) Value() pulumi.AnyOutput {
	return o.ApplyT(func(v EnabledControlParameter) interface{} { return v.Value }).(pulumi.AnyOutput)
}

type EnabledControlParameterArrayOutput struct{ *pulumi.OutputState }

func (EnabledControlParameterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EnabledControlParameter)(nil)).Elem()
}

func (o EnabledControlParameterArrayOutput) ToEnabledControlParameterArrayOutput() EnabledControlParameterArrayOutput {
	return o
}

func (o EnabledControlParameterArrayOutput) ToEnabledControlParameterArrayOutputWithContext(ctx context.Context) EnabledControlParameterArrayOutput {
	return o
}

func (o EnabledControlParameterArrayOutput) Index(i pulumi.IntInput) EnabledControlParameterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EnabledControlParameter {
		return vs[0].([]EnabledControlParameter)[vs[1].(int)]
	}).(EnabledControlParameterOutput)
}

// A key-value pair to associate with a resource.
type EnabledControlTag struct {
	// The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:.
	Key string `pulumi:"key"`
	// The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:.
	Value string `pulumi:"value"`
}

type LandingZoneTag struct {
	Key   *string `pulumi:"key"`
	Value *string `pulumi:"value"`
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EnabledBaselineParameterInput)(nil)).Elem(), EnabledBaselineParameterArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnabledBaselineParameterArrayInput)(nil)).Elem(), EnabledBaselineParameterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnabledControlParameterInput)(nil)).Elem(), EnabledControlParameterArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnabledControlParameterArrayInput)(nil)).Elem(), EnabledControlParameterArray{})
	pulumi.RegisterOutputType(EnabledBaselineParameterOutput{})
	pulumi.RegisterOutputType(EnabledBaselineParameterArrayOutput{})
	pulumi.RegisterOutputType(EnabledControlParameterOutput{})
	pulumi.RegisterOutputType(EnabledControlParameterArrayOutput{})
}

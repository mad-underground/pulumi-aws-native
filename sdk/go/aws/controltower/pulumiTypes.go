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

type LandingZoneTag struct {
	Key   *string `pulumi:"key"`
	Value *string `pulumi:"value"`
}

// LandingZoneTagInput is an input type that accepts LandingZoneTagArgs and LandingZoneTagOutput values.
// You can construct a concrete instance of `LandingZoneTagInput` via:
//
//	LandingZoneTagArgs{...}
type LandingZoneTagInput interface {
	pulumi.Input

	ToLandingZoneTagOutput() LandingZoneTagOutput
	ToLandingZoneTagOutputWithContext(context.Context) LandingZoneTagOutput
}

type LandingZoneTagArgs struct {
	Key   pulumi.StringPtrInput `pulumi:"key"`
	Value pulumi.StringPtrInput `pulumi:"value"`
}

func (LandingZoneTagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LandingZoneTag)(nil)).Elem()
}

func (i LandingZoneTagArgs) ToLandingZoneTagOutput() LandingZoneTagOutput {
	return i.ToLandingZoneTagOutputWithContext(context.Background())
}

func (i LandingZoneTagArgs) ToLandingZoneTagOutputWithContext(ctx context.Context) LandingZoneTagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LandingZoneTagOutput)
}

// LandingZoneTagArrayInput is an input type that accepts LandingZoneTagArray and LandingZoneTagArrayOutput values.
// You can construct a concrete instance of `LandingZoneTagArrayInput` via:
//
//	LandingZoneTagArray{ LandingZoneTagArgs{...} }
type LandingZoneTagArrayInput interface {
	pulumi.Input

	ToLandingZoneTagArrayOutput() LandingZoneTagArrayOutput
	ToLandingZoneTagArrayOutputWithContext(context.Context) LandingZoneTagArrayOutput
}

type LandingZoneTagArray []LandingZoneTagInput

func (LandingZoneTagArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LandingZoneTag)(nil)).Elem()
}

func (i LandingZoneTagArray) ToLandingZoneTagArrayOutput() LandingZoneTagArrayOutput {
	return i.ToLandingZoneTagArrayOutputWithContext(context.Background())
}

func (i LandingZoneTagArray) ToLandingZoneTagArrayOutputWithContext(ctx context.Context) LandingZoneTagArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LandingZoneTagArrayOutput)
}

type LandingZoneTagOutput struct{ *pulumi.OutputState }

func (LandingZoneTagOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LandingZoneTag)(nil)).Elem()
}

func (o LandingZoneTagOutput) ToLandingZoneTagOutput() LandingZoneTagOutput {
	return o
}

func (o LandingZoneTagOutput) ToLandingZoneTagOutputWithContext(ctx context.Context) LandingZoneTagOutput {
	return o
}

func (o LandingZoneTagOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LandingZoneTag) *string { return v.Key }).(pulumi.StringPtrOutput)
}

func (o LandingZoneTagOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LandingZoneTag) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type LandingZoneTagArrayOutput struct{ *pulumi.OutputState }

func (LandingZoneTagArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LandingZoneTag)(nil)).Elem()
}

func (o LandingZoneTagArrayOutput) ToLandingZoneTagArrayOutput() LandingZoneTagArrayOutput {
	return o
}

func (o LandingZoneTagArrayOutput) ToLandingZoneTagArrayOutputWithContext(ctx context.Context) LandingZoneTagArrayOutput {
	return o
}

func (o LandingZoneTagArrayOutput) Index(i pulumi.IntInput) LandingZoneTagOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LandingZoneTag {
		return vs[0].([]LandingZoneTag)[vs[1].(int)]
	}).(LandingZoneTagOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EnabledControlParameterInput)(nil)).Elem(), EnabledControlParameterArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnabledControlParameterArrayInput)(nil)).Elem(), EnabledControlParameterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LandingZoneTagInput)(nil)).Elem(), LandingZoneTagArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LandingZoneTagArrayInput)(nil)).Elem(), LandingZoneTagArray{})
	pulumi.RegisterOutputType(EnabledControlParameterOutput{})
	pulumi.RegisterOutputType(EnabledControlParameterArrayOutput{})
	pulumi.RegisterOutputType(LandingZoneTagOutput{})
	pulumi.RegisterOutputType(LandingZoneTagArrayOutput{})
}

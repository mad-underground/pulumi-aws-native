// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rolesanywhere

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

type CrlTag struct {
	Key   string `pulumi:"key"`
	Value string `pulumi:"value"`
}

type ProfileTag struct {
	Key   string `pulumi:"key"`
	Value string `pulumi:"value"`
}

type TrustAnchorNotificationSetting struct {
	Channel   *TrustAnchorNotificationChannel `pulumi:"channel"`
	Enabled   bool                            `pulumi:"enabled"`
	Event     TrustAnchorNotificationEvent    `pulumi:"event"`
	Threshold *float64                        `pulumi:"threshold"`
}

// TrustAnchorNotificationSettingInput is an input type that accepts TrustAnchorNotificationSettingArgs and TrustAnchorNotificationSettingOutput values.
// You can construct a concrete instance of `TrustAnchorNotificationSettingInput` via:
//
//	TrustAnchorNotificationSettingArgs{...}
type TrustAnchorNotificationSettingInput interface {
	pulumi.Input

	ToTrustAnchorNotificationSettingOutput() TrustAnchorNotificationSettingOutput
	ToTrustAnchorNotificationSettingOutputWithContext(context.Context) TrustAnchorNotificationSettingOutput
}

type TrustAnchorNotificationSettingArgs struct {
	Channel   TrustAnchorNotificationChannelPtrInput `pulumi:"channel"`
	Enabled   pulumi.BoolInput                       `pulumi:"enabled"`
	Event     TrustAnchorNotificationEventInput      `pulumi:"event"`
	Threshold pulumi.Float64PtrInput                 `pulumi:"threshold"`
}

func (TrustAnchorNotificationSettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TrustAnchorNotificationSetting)(nil)).Elem()
}

func (i TrustAnchorNotificationSettingArgs) ToTrustAnchorNotificationSettingOutput() TrustAnchorNotificationSettingOutput {
	return i.ToTrustAnchorNotificationSettingOutputWithContext(context.Background())
}

func (i TrustAnchorNotificationSettingArgs) ToTrustAnchorNotificationSettingOutputWithContext(ctx context.Context) TrustAnchorNotificationSettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrustAnchorNotificationSettingOutput)
}

// TrustAnchorNotificationSettingArrayInput is an input type that accepts TrustAnchorNotificationSettingArray and TrustAnchorNotificationSettingArrayOutput values.
// You can construct a concrete instance of `TrustAnchorNotificationSettingArrayInput` via:
//
//	TrustAnchorNotificationSettingArray{ TrustAnchorNotificationSettingArgs{...} }
type TrustAnchorNotificationSettingArrayInput interface {
	pulumi.Input

	ToTrustAnchorNotificationSettingArrayOutput() TrustAnchorNotificationSettingArrayOutput
	ToTrustAnchorNotificationSettingArrayOutputWithContext(context.Context) TrustAnchorNotificationSettingArrayOutput
}

type TrustAnchorNotificationSettingArray []TrustAnchorNotificationSettingInput

func (TrustAnchorNotificationSettingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TrustAnchorNotificationSetting)(nil)).Elem()
}

func (i TrustAnchorNotificationSettingArray) ToTrustAnchorNotificationSettingArrayOutput() TrustAnchorNotificationSettingArrayOutput {
	return i.ToTrustAnchorNotificationSettingArrayOutputWithContext(context.Background())
}

func (i TrustAnchorNotificationSettingArray) ToTrustAnchorNotificationSettingArrayOutputWithContext(ctx context.Context) TrustAnchorNotificationSettingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrustAnchorNotificationSettingArrayOutput)
}

type TrustAnchorNotificationSettingOutput struct{ *pulumi.OutputState }

func (TrustAnchorNotificationSettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TrustAnchorNotificationSetting)(nil)).Elem()
}

func (o TrustAnchorNotificationSettingOutput) ToTrustAnchorNotificationSettingOutput() TrustAnchorNotificationSettingOutput {
	return o
}

func (o TrustAnchorNotificationSettingOutput) ToTrustAnchorNotificationSettingOutputWithContext(ctx context.Context) TrustAnchorNotificationSettingOutput {
	return o
}

func (o TrustAnchorNotificationSettingOutput) Channel() TrustAnchorNotificationChannelPtrOutput {
	return o.ApplyT(func(v TrustAnchorNotificationSetting) *TrustAnchorNotificationChannel { return v.Channel }).(TrustAnchorNotificationChannelPtrOutput)
}

func (o TrustAnchorNotificationSettingOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v TrustAnchorNotificationSetting) bool { return v.Enabled }).(pulumi.BoolOutput)
}

func (o TrustAnchorNotificationSettingOutput) Event() TrustAnchorNotificationEventOutput {
	return o.ApplyT(func(v TrustAnchorNotificationSetting) TrustAnchorNotificationEvent { return v.Event }).(TrustAnchorNotificationEventOutput)
}

func (o TrustAnchorNotificationSettingOutput) Threshold() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v TrustAnchorNotificationSetting) *float64 { return v.Threshold }).(pulumi.Float64PtrOutput)
}

type TrustAnchorNotificationSettingArrayOutput struct{ *pulumi.OutputState }

func (TrustAnchorNotificationSettingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TrustAnchorNotificationSetting)(nil)).Elem()
}

func (o TrustAnchorNotificationSettingArrayOutput) ToTrustAnchorNotificationSettingArrayOutput() TrustAnchorNotificationSettingArrayOutput {
	return o
}

func (o TrustAnchorNotificationSettingArrayOutput) ToTrustAnchorNotificationSettingArrayOutputWithContext(ctx context.Context) TrustAnchorNotificationSettingArrayOutput {
	return o
}

func (o TrustAnchorNotificationSettingArrayOutput) Index(i pulumi.IntInput) TrustAnchorNotificationSettingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TrustAnchorNotificationSetting {
		return vs[0].([]TrustAnchorNotificationSetting)[vs[1].(int)]
	}).(TrustAnchorNotificationSettingOutput)
}

type TrustAnchorSource struct {
	SourceData interface{}      `pulumi:"sourceData"`
	SourceType *TrustAnchorType `pulumi:"sourceType"`
}

// TrustAnchorSourceInput is an input type that accepts TrustAnchorSourceArgs and TrustAnchorSourceOutput values.
// You can construct a concrete instance of `TrustAnchorSourceInput` via:
//
//	TrustAnchorSourceArgs{...}
type TrustAnchorSourceInput interface {
	pulumi.Input

	ToTrustAnchorSourceOutput() TrustAnchorSourceOutput
	ToTrustAnchorSourceOutputWithContext(context.Context) TrustAnchorSourceOutput
}

type TrustAnchorSourceArgs struct {
	SourceData pulumi.Input            `pulumi:"sourceData"`
	SourceType TrustAnchorTypePtrInput `pulumi:"sourceType"`
}

func (TrustAnchorSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TrustAnchorSource)(nil)).Elem()
}

func (i TrustAnchorSourceArgs) ToTrustAnchorSourceOutput() TrustAnchorSourceOutput {
	return i.ToTrustAnchorSourceOutputWithContext(context.Background())
}

func (i TrustAnchorSourceArgs) ToTrustAnchorSourceOutputWithContext(ctx context.Context) TrustAnchorSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrustAnchorSourceOutput)
}

type TrustAnchorSourceOutput struct{ *pulumi.OutputState }

func (TrustAnchorSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TrustAnchorSource)(nil)).Elem()
}

func (o TrustAnchorSourceOutput) ToTrustAnchorSourceOutput() TrustAnchorSourceOutput {
	return o
}

func (o TrustAnchorSourceOutput) ToTrustAnchorSourceOutputWithContext(ctx context.Context) TrustAnchorSourceOutput {
	return o
}

func (o TrustAnchorSourceOutput) SourceData() pulumi.AnyOutput {
	return o.ApplyT(func(v TrustAnchorSource) interface{} { return v.SourceData }).(pulumi.AnyOutput)
}

func (o TrustAnchorSourceOutput) SourceType() TrustAnchorTypePtrOutput {
	return o.ApplyT(func(v TrustAnchorSource) *TrustAnchorType { return v.SourceType }).(TrustAnchorTypePtrOutput)
}

type TrustAnchorSourcePtrOutput struct{ *pulumi.OutputState }

func (TrustAnchorSourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TrustAnchorSource)(nil)).Elem()
}

func (o TrustAnchorSourcePtrOutput) ToTrustAnchorSourcePtrOutput() TrustAnchorSourcePtrOutput {
	return o
}

func (o TrustAnchorSourcePtrOutput) ToTrustAnchorSourcePtrOutputWithContext(ctx context.Context) TrustAnchorSourcePtrOutput {
	return o
}

func (o TrustAnchorSourcePtrOutput) Elem() TrustAnchorSourceOutput {
	return o.ApplyT(func(v *TrustAnchorSource) TrustAnchorSource {
		if v != nil {
			return *v
		}
		var ret TrustAnchorSource
		return ret
	}).(TrustAnchorSourceOutput)
}

func (o TrustAnchorSourcePtrOutput) SourceData() pulumi.AnyOutput {
	return o.ApplyT(func(v *TrustAnchorSource) interface{} {
		if v == nil {
			return nil
		}
		return v.SourceData
	}).(pulumi.AnyOutput)
}

func (o TrustAnchorSourcePtrOutput) SourceType() TrustAnchorTypePtrOutput {
	return o.ApplyT(func(v *TrustAnchorSource) *TrustAnchorType {
		if v == nil {
			return nil
		}
		return v.SourceType
	}).(TrustAnchorTypePtrOutput)
}

type TrustAnchorSourceData0Properties struct {
	X509CertificateData string `pulumi:"x509CertificateData"`
}

// TrustAnchorSourceData0PropertiesInput is an input type that accepts TrustAnchorSourceData0PropertiesArgs and TrustAnchorSourceData0PropertiesOutput values.
// You can construct a concrete instance of `TrustAnchorSourceData0PropertiesInput` via:
//
//	TrustAnchorSourceData0PropertiesArgs{...}
type TrustAnchorSourceData0PropertiesInput interface {
	pulumi.Input

	ToTrustAnchorSourceData0PropertiesOutput() TrustAnchorSourceData0PropertiesOutput
	ToTrustAnchorSourceData0PropertiesOutputWithContext(context.Context) TrustAnchorSourceData0PropertiesOutput
}

type TrustAnchorSourceData0PropertiesArgs struct {
	X509CertificateData pulumi.StringInput `pulumi:"x509CertificateData"`
}

func (TrustAnchorSourceData0PropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TrustAnchorSourceData0Properties)(nil)).Elem()
}

func (i TrustAnchorSourceData0PropertiesArgs) ToTrustAnchorSourceData0PropertiesOutput() TrustAnchorSourceData0PropertiesOutput {
	return i.ToTrustAnchorSourceData0PropertiesOutputWithContext(context.Background())
}

func (i TrustAnchorSourceData0PropertiesArgs) ToTrustAnchorSourceData0PropertiesOutputWithContext(ctx context.Context) TrustAnchorSourceData0PropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrustAnchorSourceData0PropertiesOutput)
}

func (i TrustAnchorSourceData0PropertiesArgs) ToTrustAnchorSourceData0PropertiesPtrOutput() TrustAnchorSourceData0PropertiesPtrOutput {
	return i.ToTrustAnchorSourceData0PropertiesPtrOutputWithContext(context.Background())
}

func (i TrustAnchorSourceData0PropertiesArgs) ToTrustAnchorSourceData0PropertiesPtrOutputWithContext(ctx context.Context) TrustAnchorSourceData0PropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrustAnchorSourceData0PropertiesOutput).ToTrustAnchorSourceData0PropertiesPtrOutputWithContext(ctx)
}

// TrustAnchorSourceData0PropertiesPtrInput is an input type that accepts TrustAnchorSourceData0PropertiesArgs, TrustAnchorSourceData0PropertiesPtr and TrustAnchorSourceData0PropertiesPtrOutput values.
// You can construct a concrete instance of `TrustAnchorSourceData0PropertiesPtrInput` via:
//
//	        TrustAnchorSourceData0PropertiesArgs{...}
//
//	or:
//
//	        nil
type TrustAnchorSourceData0PropertiesPtrInput interface {
	pulumi.Input

	ToTrustAnchorSourceData0PropertiesPtrOutput() TrustAnchorSourceData0PropertiesPtrOutput
	ToTrustAnchorSourceData0PropertiesPtrOutputWithContext(context.Context) TrustAnchorSourceData0PropertiesPtrOutput
}

type trustAnchorSourceData0PropertiesPtrType TrustAnchorSourceData0PropertiesArgs

func TrustAnchorSourceData0PropertiesPtr(v *TrustAnchorSourceData0PropertiesArgs) TrustAnchorSourceData0PropertiesPtrInput {
	return (*trustAnchorSourceData0PropertiesPtrType)(v)
}

func (*trustAnchorSourceData0PropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TrustAnchorSourceData0Properties)(nil)).Elem()
}

func (i *trustAnchorSourceData0PropertiesPtrType) ToTrustAnchorSourceData0PropertiesPtrOutput() TrustAnchorSourceData0PropertiesPtrOutput {
	return i.ToTrustAnchorSourceData0PropertiesPtrOutputWithContext(context.Background())
}

func (i *trustAnchorSourceData0PropertiesPtrType) ToTrustAnchorSourceData0PropertiesPtrOutputWithContext(ctx context.Context) TrustAnchorSourceData0PropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrustAnchorSourceData0PropertiesPtrOutput)
}

type TrustAnchorSourceData0PropertiesOutput struct{ *pulumi.OutputState }

func (TrustAnchorSourceData0PropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TrustAnchorSourceData0Properties)(nil)).Elem()
}

func (o TrustAnchorSourceData0PropertiesOutput) ToTrustAnchorSourceData0PropertiesOutput() TrustAnchorSourceData0PropertiesOutput {
	return o
}

func (o TrustAnchorSourceData0PropertiesOutput) ToTrustAnchorSourceData0PropertiesOutputWithContext(ctx context.Context) TrustAnchorSourceData0PropertiesOutput {
	return o
}

func (o TrustAnchorSourceData0PropertiesOutput) ToTrustAnchorSourceData0PropertiesPtrOutput() TrustAnchorSourceData0PropertiesPtrOutput {
	return o.ToTrustAnchorSourceData0PropertiesPtrOutputWithContext(context.Background())
}

func (o TrustAnchorSourceData0PropertiesOutput) ToTrustAnchorSourceData0PropertiesPtrOutputWithContext(ctx context.Context) TrustAnchorSourceData0PropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TrustAnchorSourceData0Properties) *TrustAnchorSourceData0Properties {
		return &v
	}).(TrustAnchorSourceData0PropertiesPtrOutput)
}

func (o TrustAnchorSourceData0PropertiesOutput) X509CertificateData() pulumi.StringOutput {
	return o.ApplyT(func(v TrustAnchorSourceData0Properties) string { return v.X509CertificateData }).(pulumi.StringOutput)
}

type TrustAnchorSourceData0PropertiesPtrOutput struct{ *pulumi.OutputState }

func (TrustAnchorSourceData0PropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TrustAnchorSourceData0Properties)(nil)).Elem()
}

func (o TrustAnchorSourceData0PropertiesPtrOutput) ToTrustAnchorSourceData0PropertiesPtrOutput() TrustAnchorSourceData0PropertiesPtrOutput {
	return o
}

func (o TrustAnchorSourceData0PropertiesPtrOutput) ToTrustAnchorSourceData0PropertiesPtrOutputWithContext(ctx context.Context) TrustAnchorSourceData0PropertiesPtrOutput {
	return o
}

func (o TrustAnchorSourceData0PropertiesPtrOutput) Elem() TrustAnchorSourceData0PropertiesOutput {
	return o.ApplyT(func(v *TrustAnchorSourceData0Properties) TrustAnchorSourceData0Properties {
		if v != nil {
			return *v
		}
		var ret TrustAnchorSourceData0Properties
		return ret
	}).(TrustAnchorSourceData0PropertiesOutput)
}

func (o TrustAnchorSourceData0PropertiesPtrOutput) X509CertificateData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TrustAnchorSourceData0Properties) *string {
		if v == nil {
			return nil
		}
		return &v.X509CertificateData
	}).(pulumi.StringPtrOutput)
}

type TrustAnchorSourceData1Properties struct {
	AcmPcaArn string `pulumi:"acmPcaArn"`
}

// TrustAnchorSourceData1PropertiesInput is an input type that accepts TrustAnchorSourceData1PropertiesArgs and TrustAnchorSourceData1PropertiesOutput values.
// You can construct a concrete instance of `TrustAnchorSourceData1PropertiesInput` via:
//
//	TrustAnchorSourceData1PropertiesArgs{...}
type TrustAnchorSourceData1PropertiesInput interface {
	pulumi.Input

	ToTrustAnchorSourceData1PropertiesOutput() TrustAnchorSourceData1PropertiesOutput
	ToTrustAnchorSourceData1PropertiesOutputWithContext(context.Context) TrustAnchorSourceData1PropertiesOutput
}

type TrustAnchorSourceData1PropertiesArgs struct {
	AcmPcaArn pulumi.StringInput `pulumi:"acmPcaArn"`
}

func (TrustAnchorSourceData1PropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TrustAnchorSourceData1Properties)(nil)).Elem()
}

func (i TrustAnchorSourceData1PropertiesArgs) ToTrustAnchorSourceData1PropertiesOutput() TrustAnchorSourceData1PropertiesOutput {
	return i.ToTrustAnchorSourceData1PropertiesOutputWithContext(context.Background())
}

func (i TrustAnchorSourceData1PropertiesArgs) ToTrustAnchorSourceData1PropertiesOutputWithContext(ctx context.Context) TrustAnchorSourceData1PropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrustAnchorSourceData1PropertiesOutput)
}

func (i TrustAnchorSourceData1PropertiesArgs) ToTrustAnchorSourceData1PropertiesPtrOutput() TrustAnchorSourceData1PropertiesPtrOutput {
	return i.ToTrustAnchorSourceData1PropertiesPtrOutputWithContext(context.Background())
}

func (i TrustAnchorSourceData1PropertiesArgs) ToTrustAnchorSourceData1PropertiesPtrOutputWithContext(ctx context.Context) TrustAnchorSourceData1PropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrustAnchorSourceData1PropertiesOutput).ToTrustAnchorSourceData1PropertiesPtrOutputWithContext(ctx)
}

// TrustAnchorSourceData1PropertiesPtrInput is an input type that accepts TrustAnchorSourceData1PropertiesArgs, TrustAnchorSourceData1PropertiesPtr and TrustAnchorSourceData1PropertiesPtrOutput values.
// You can construct a concrete instance of `TrustAnchorSourceData1PropertiesPtrInput` via:
//
//	        TrustAnchorSourceData1PropertiesArgs{...}
//
//	or:
//
//	        nil
type TrustAnchorSourceData1PropertiesPtrInput interface {
	pulumi.Input

	ToTrustAnchorSourceData1PropertiesPtrOutput() TrustAnchorSourceData1PropertiesPtrOutput
	ToTrustAnchorSourceData1PropertiesPtrOutputWithContext(context.Context) TrustAnchorSourceData1PropertiesPtrOutput
}

type trustAnchorSourceData1PropertiesPtrType TrustAnchorSourceData1PropertiesArgs

func TrustAnchorSourceData1PropertiesPtr(v *TrustAnchorSourceData1PropertiesArgs) TrustAnchorSourceData1PropertiesPtrInput {
	return (*trustAnchorSourceData1PropertiesPtrType)(v)
}

func (*trustAnchorSourceData1PropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TrustAnchorSourceData1Properties)(nil)).Elem()
}

func (i *trustAnchorSourceData1PropertiesPtrType) ToTrustAnchorSourceData1PropertiesPtrOutput() TrustAnchorSourceData1PropertiesPtrOutput {
	return i.ToTrustAnchorSourceData1PropertiesPtrOutputWithContext(context.Background())
}

func (i *trustAnchorSourceData1PropertiesPtrType) ToTrustAnchorSourceData1PropertiesPtrOutputWithContext(ctx context.Context) TrustAnchorSourceData1PropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrustAnchorSourceData1PropertiesPtrOutput)
}

type TrustAnchorSourceData1PropertiesOutput struct{ *pulumi.OutputState }

func (TrustAnchorSourceData1PropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TrustAnchorSourceData1Properties)(nil)).Elem()
}

func (o TrustAnchorSourceData1PropertiesOutput) ToTrustAnchorSourceData1PropertiesOutput() TrustAnchorSourceData1PropertiesOutput {
	return o
}

func (o TrustAnchorSourceData1PropertiesOutput) ToTrustAnchorSourceData1PropertiesOutputWithContext(ctx context.Context) TrustAnchorSourceData1PropertiesOutput {
	return o
}

func (o TrustAnchorSourceData1PropertiesOutput) ToTrustAnchorSourceData1PropertiesPtrOutput() TrustAnchorSourceData1PropertiesPtrOutput {
	return o.ToTrustAnchorSourceData1PropertiesPtrOutputWithContext(context.Background())
}

func (o TrustAnchorSourceData1PropertiesOutput) ToTrustAnchorSourceData1PropertiesPtrOutputWithContext(ctx context.Context) TrustAnchorSourceData1PropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TrustAnchorSourceData1Properties) *TrustAnchorSourceData1Properties {
		return &v
	}).(TrustAnchorSourceData1PropertiesPtrOutput)
}

func (o TrustAnchorSourceData1PropertiesOutput) AcmPcaArn() pulumi.StringOutput {
	return o.ApplyT(func(v TrustAnchorSourceData1Properties) string { return v.AcmPcaArn }).(pulumi.StringOutput)
}

type TrustAnchorSourceData1PropertiesPtrOutput struct{ *pulumi.OutputState }

func (TrustAnchorSourceData1PropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TrustAnchorSourceData1Properties)(nil)).Elem()
}

func (o TrustAnchorSourceData1PropertiesPtrOutput) ToTrustAnchorSourceData1PropertiesPtrOutput() TrustAnchorSourceData1PropertiesPtrOutput {
	return o
}

func (o TrustAnchorSourceData1PropertiesPtrOutput) ToTrustAnchorSourceData1PropertiesPtrOutputWithContext(ctx context.Context) TrustAnchorSourceData1PropertiesPtrOutput {
	return o
}

func (o TrustAnchorSourceData1PropertiesPtrOutput) Elem() TrustAnchorSourceData1PropertiesOutput {
	return o.ApplyT(func(v *TrustAnchorSourceData1Properties) TrustAnchorSourceData1Properties {
		if v != nil {
			return *v
		}
		var ret TrustAnchorSourceData1Properties
		return ret
	}).(TrustAnchorSourceData1PropertiesOutput)
}

func (o TrustAnchorSourceData1PropertiesPtrOutput) AcmPcaArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TrustAnchorSourceData1Properties) *string {
		if v == nil {
			return nil
		}
		return &v.AcmPcaArn
	}).(pulumi.StringPtrOutput)
}

type TrustAnchorTag struct {
	Key   string `pulumi:"key"`
	Value string `pulumi:"value"`
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TrustAnchorNotificationSettingInput)(nil)).Elem(), TrustAnchorNotificationSettingArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TrustAnchorNotificationSettingArrayInput)(nil)).Elem(), TrustAnchorNotificationSettingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TrustAnchorSourceInput)(nil)).Elem(), TrustAnchorSourceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TrustAnchorSourceData0PropertiesInput)(nil)).Elem(), TrustAnchorSourceData0PropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TrustAnchorSourceData0PropertiesPtrInput)(nil)).Elem(), TrustAnchorSourceData0PropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TrustAnchorSourceData1PropertiesInput)(nil)).Elem(), TrustAnchorSourceData1PropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TrustAnchorSourceData1PropertiesPtrInput)(nil)).Elem(), TrustAnchorSourceData1PropertiesArgs{})
	pulumi.RegisterOutputType(TrustAnchorNotificationSettingOutput{})
	pulumi.RegisterOutputType(TrustAnchorNotificationSettingArrayOutput{})
	pulumi.RegisterOutputType(TrustAnchorSourceOutput{})
	pulumi.RegisterOutputType(TrustAnchorSourcePtrOutput{})
	pulumi.RegisterOutputType(TrustAnchorSourceData0PropertiesOutput{})
	pulumi.RegisterOutputType(TrustAnchorSourceData0PropertiesPtrOutput{})
	pulumi.RegisterOutputType(TrustAnchorSourceData1PropertiesOutput{})
	pulumi.RegisterOutputType(TrustAnchorSourceData1PropertiesPtrOutput{})
}

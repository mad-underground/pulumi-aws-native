// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package groundstation

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ConfigBandwidthUnits string

const (
	ConfigBandwidthUnitsGHz = ConfigBandwidthUnits("GHz")
	ConfigBandwidthUnitsMHz = ConfigBandwidthUnits("MHz")
	ConfigBandwidthUnitsKHz = ConfigBandwidthUnits("kHz")
)

func (ConfigBandwidthUnits) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigBandwidthUnits)(nil)).Elem()
}

func (e ConfigBandwidthUnits) ToConfigBandwidthUnitsOutput() ConfigBandwidthUnitsOutput {
	return pulumi.ToOutput(e).(ConfigBandwidthUnitsOutput)
}

func (e ConfigBandwidthUnits) ToConfigBandwidthUnitsOutputWithContext(ctx context.Context) ConfigBandwidthUnitsOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ConfigBandwidthUnitsOutput)
}

func (e ConfigBandwidthUnits) ToConfigBandwidthUnitsPtrOutput() ConfigBandwidthUnitsPtrOutput {
	return e.ToConfigBandwidthUnitsPtrOutputWithContext(context.Background())
}

func (e ConfigBandwidthUnits) ToConfigBandwidthUnitsPtrOutputWithContext(ctx context.Context) ConfigBandwidthUnitsPtrOutput {
	return ConfigBandwidthUnits(e).ToConfigBandwidthUnitsOutputWithContext(ctx).ToConfigBandwidthUnitsPtrOutputWithContext(ctx)
}

func (e ConfigBandwidthUnits) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ConfigBandwidthUnits) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ConfigBandwidthUnits) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ConfigBandwidthUnits) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ConfigBandwidthUnitsOutput struct{ *pulumi.OutputState }

func (ConfigBandwidthUnitsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigBandwidthUnits)(nil)).Elem()
}

func (o ConfigBandwidthUnitsOutput) ToConfigBandwidthUnitsOutput() ConfigBandwidthUnitsOutput {
	return o
}

func (o ConfigBandwidthUnitsOutput) ToConfigBandwidthUnitsOutputWithContext(ctx context.Context) ConfigBandwidthUnitsOutput {
	return o
}

func (o ConfigBandwidthUnitsOutput) ToConfigBandwidthUnitsPtrOutput() ConfigBandwidthUnitsPtrOutput {
	return o.ToConfigBandwidthUnitsPtrOutputWithContext(context.Background())
}

func (o ConfigBandwidthUnitsOutput) ToConfigBandwidthUnitsPtrOutputWithContext(ctx context.Context) ConfigBandwidthUnitsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConfigBandwidthUnits) *ConfigBandwidthUnits {
		return &v
	}).(ConfigBandwidthUnitsPtrOutput)
}

func (o ConfigBandwidthUnitsOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ConfigBandwidthUnitsOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ConfigBandwidthUnits) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ConfigBandwidthUnitsOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ConfigBandwidthUnitsOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ConfigBandwidthUnits) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ConfigBandwidthUnitsPtrOutput struct{ *pulumi.OutputState }

func (ConfigBandwidthUnitsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigBandwidthUnits)(nil)).Elem()
}

func (o ConfigBandwidthUnitsPtrOutput) ToConfigBandwidthUnitsPtrOutput() ConfigBandwidthUnitsPtrOutput {
	return o
}

func (o ConfigBandwidthUnitsPtrOutput) ToConfigBandwidthUnitsPtrOutputWithContext(ctx context.Context) ConfigBandwidthUnitsPtrOutput {
	return o
}

func (o ConfigBandwidthUnitsPtrOutput) Elem() ConfigBandwidthUnitsOutput {
	return o.ApplyT(func(v *ConfigBandwidthUnits) ConfigBandwidthUnits {
		if v != nil {
			return *v
		}
		var ret ConfigBandwidthUnits
		return ret
	}).(ConfigBandwidthUnitsOutput)
}

func (o ConfigBandwidthUnitsPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ConfigBandwidthUnitsPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ConfigBandwidthUnits) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ConfigBandwidthUnitsInput is an input type that accepts ConfigBandwidthUnitsArgs and ConfigBandwidthUnitsOutput values.
// You can construct a concrete instance of `ConfigBandwidthUnitsInput` via:
//
//	ConfigBandwidthUnitsArgs{...}
type ConfigBandwidthUnitsInput interface {
	pulumi.Input

	ToConfigBandwidthUnitsOutput() ConfigBandwidthUnitsOutput
	ToConfigBandwidthUnitsOutputWithContext(context.Context) ConfigBandwidthUnitsOutput
}

var configBandwidthUnitsPtrType = reflect.TypeOf((**ConfigBandwidthUnits)(nil)).Elem()

type ConfigBandwidthUnitsPtrInput interface {
	pulumi.Input

	ToConfigBandwidthUnitsPtrOutput() ConfigBandwidthUnitsPtrOutput
	ToConfigBandwidthUnitsPtrOutputWithContext(context.Context) ConfigBandwidthUnitsPtrOutput
}

type configBandwidthUnitsPtr string

func ConfigBandwidthUnitsPtr(v string) ConfigBandwidthUnitsPtrInput {
	return (*configBandwidthUnitsPtr)(&v)
}

func (*configBandwidthUnitsPtr) ElementType() reflect.Type {
	return configBandwidthUnitsPtrType
}

func (in *configBandwidthUnitsPtr) ToConfigBandwidthUnitsPtrOutput() ConfigBandwidthUnitsPtrOutput {
	return pulumi.ToOutput(in).(ConfigBandwidthUnitsPtrOutput)
}

func (in *configBandwidthUnitsPtr) ToConfigBandwidthUnitsPtrOutputWithContext(ctx context.Context) ConfigBandwidthUnitsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ConfigBandwidthUnitsPtrOutput)
}

type ConfigEirpUnits string

const (
	ConfigEirpUnitsDBW = ConfigEirpUnits("dBW")
)

func (ConfigEirpUnits) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigEirpUnits)(nil)).Elem()
}

func (e ConfigEirpUnits) ToConfigEirpUnitsOutput() ConfigEirpUnitsOutput {
	return pulumi.ToOutput(e).(ConfigEirpUnitsOutput)
}

func (e ConfigEirpUnits) ToConfigEirpUnitsOutputWithContext(ctx context.Context) ConfigEirpUnitsOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ConfigEirpUnitsOutput)
}

func (e ConfigEirpUnits) ToConfigEirpUnitsPtrOutput() ConfigEirpUnitsPtrOutput {
	return e.ToConfigEirpUnitsPtrOutputWithContext(context.Background())
}

func (e ConfigEirpUnits) ToConfigEirpUnitsPtrOutputWithContext(ctx context.Context) ConfigEirpUnitsPtrOutput {
	return ConfigEirpUnits(e).ToConfigEirpUnitsOutputWithContext(ctx).ToConfigEirpUnitsPtrOutputWithContext(ctx)
}

func (e ConfigEirpUnits) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ConfigEirpUnits) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ConfigEirpUnits) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ConfigEirpUnits) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ConfigEirpUnitsOutput struct{ *pulumi.OutputState }

func (ConfigEirpUnitsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigEirpUnits)(nil)).Elem()
}

func (o ConfigEirpUnitsOutput) ToConfigEirpUnitsOutput() ConfigEirpUnitsOutput {
	return o
}

func (o ConfigEirpUnitsOutput) ToConfigEirpUnitsOutputWithContext(ctx context.Context) ConfigEirpUnitsOutput {
	return o
}

func (o ConfigEirpUnitsOutput) ToConfigEirpUnitsPtrOutput() ConfigEirpUnitsPtrOutput {
	return o.ToConfigEirpUnitsPtrOutputWithContext(context.Background())
}

func (o ConfigEirpUnitsOutput) ToConfigEirpUnitsPtrOutputWithContext(ctx context.Context) ConfigEirpUnitsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConfigEirpUnits) *ConfigEirpUnits {
		return &v
	}).(ConfigEirpUnitsPtrOutput)
}

func (o ConfigEirpUnitsOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ConfigEirpUnitsOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ConfigEirpUnits) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ConfigEirpUnitsOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ConfigEirpUnitsOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ConfigEirpUnits) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ConfigEirpUnitsPtrOutput struct{ *pulumi.OutputState }

func (ConfigEirpUnitsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigEirpUnits)(nil)).Elem()
}

func (o ConfigEirpUnitsPtrOutput) ToConfigEirpUnitsPtrOutput() ConfigEirpUnitsPtrOutput {
	return o
}

func (o ConfigEirpUnitsPtrOutput) ToConfigEirpUnitsPtrOutputWithContext(ctx context.Context) ConfigEirpUnitsPtrOutput {
	return o
}

func (o ConfigEirpUnitsPtrOutput) Elem() ConfigEirpUnitsOutput {
	return o.ApplyT(func(v *ConfigEirpUnits) ConfigEirpUnits {
		if v != nil {
			return *v
		}
		var ret ConfigEirpUnits
		return ret
	}).(ConfigEirpUnitsOutput)
}

func (o ConfigEirpUnitsPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ConfigEirpUnitsPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ConfigEirpUnits) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ConfigEirpUnitsInput is an input type that accepts ConfigEirpUnitsArgs and ConfigEirpUnitsOutput values.
// You can construct a concrete instance of `ConfigEirpUnitsInput` via:
//
//	ConfigEirpUnitsArgs{...}
type ConfigEirpUnitsInput interface {
	pulumi.Input

	ToConfigEirpUnitsOutput() ConfigEirpUnitsOutput
	ToConfigEirpUnitsOutputWithContext(context.Context) ConfigEirpUnitsOutput
}

var configEirpUnitsPtrType = reflect.TypeOf((**ConfigEirpUnits)(nil)).Elem()

type ConfigEirpUnitsPtrInput interface {
	pulumi.Input

	ToConfigEirpUnitsPtrOutput() ConfigEirpUnitsPtrOutput
	ToConfigEirpUnitsPtrOutputWithContext(context.Context) ConfigEirpUnitsPtrOutput
}

type configEirpUnitsPtr string

func ConfigEirpUnitsPtr(v string) ConfigEirpUnitsPtrInput {
	return (*configEirpUnitsPtr)(&v)
}

func (*configEirpUnitsPtr) ElementType() reflect.Type {
	return configEirpUnitsPtrType
}

func (in *configEirpUnitsPtr) ToConfigEirpUnitsPtrOutput() ConfigEirpUnitsPtrOutput {
	return pulumi.ToOutput(in).(ConfigEirpUnitsPtrOutput)
}

func (in *configEirpUnitsPtr) ToConfigEirpUnitsPtrOutputWithContext(ctx context.Context) ConfigEirpUnitsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ConfigEirpUnitsPtrOutput)
}

type ConfigFrequencyUnits string

const (
	ConfigFrequencyUnitsGHz = ConfigFrequencyUnits("GHz")
	ConfigFrequencyUnitsMHz = ConfigFrequencyUnits("MHz")
	ConfigFrequencyUnitsKHz = ConfigFrequencyUnits("kHz")
)

func (ConfigFrequencyUnits) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigFrequencyUnits)(nil)).Elem()
}

func (e ConfigFrequencyUnits) ToConfigFrequencyUnitsOutput() ConfigFrequencyUnitsOutput {
	return pulumi.ToOutput(e).(ConfigFrequencyUnitsOutput)
}

func (e ConfigFrequencyUnits) ToConfigFrequencyUnitsOutputWithContext(ctx context.Context) ConfigFrequencyUnitsOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ConfigFrequencyUnitsOutput)
}

func (e ConfigFrequencyUnits) ToConfigFrequencyUnitsPtrOutput() ConfigFrequencyUnitsPtrOutput {
	return e.ToConfigFrequencyUnitsPtrOutputWithContext(context.Background())
}

func (e ConfigFrequencyUnits) ToConfigFrequencyUnitsPtrOutputWithContext(ctx context.Context) ConfigFrequencyUnitsPtrOutput {
	return ConfigFrequencyUnits(e).ToConfigFrequencyUnitsOutputWithContext(ctx).ToConfigFrequencyUnitsPtrOutputWithContext(ctx)
}

func (e ConfigFrequencyUnits) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ConfigFrequencyUnits) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ConfigFrequencyUnits) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ConfigFrequencyUnits) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ConfigFrequencyUnitsOutput struct{ *pulumi.OutputState }

func (ConfigFrequencyUnitsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigFrequencyUnits)(nil)).Elem()
}

func (o ConfigFrequencyUnitsOutput) ToConfigFrequencyUnitsOutput() ConfigFrequencyUnitsOutput {
	return o
}

func (o ConfigFrequencyUnitsOutput) ToConfigFrequencyUnitsOutputWithContext(ctx context.Context) ConfigFrequencyUnitsOutput {
	return o
}

func (o ConfigFrequencyUnitsOutput) ToConfigFrequencyUnitsPtrOutput() ConfigFrequencyUnitsPtrOutput {
	return o.ToConfigFrequencyUnitsPtrOutputWithContext(context.Background())
}

func (o ConfigFrequencyUnitsOutput) ToConfigFrequencyUnitsPtrOutputWithContext(ctx context.Context) ConfigFrequencyUnitsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConfigFrequencyUnits) *ConfigFrequencyUnits {
		return &v
	}).(ConfigFrequencyUnitsPtrOutput)
}

func (o ConfigFrequencyUnitsOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ConfigFrequencyUnitsOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ConfigFrequencyUnits) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ConfigFrequencyUnitsOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ConfigFrequencyUnitsOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ConfigFrequencyUnits) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ConfigFrequencyUnitsPtrOutput struct{ *pulumi.OutputState }

func (ConfigFrequencyUnitsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigFrequencyUnits)(nil)).Elem()
}

func (o ConfigFrequencyUnitsPtrOutput) ToConfigFrequencyUnitsPtrOutput() ConfigFrequencyUnitsPtrOutput {
	return o
}

func (o ConfigFrequencyUnitsPtrOutput) ToConfigFrequencyUnitsPtrOutputWithContext(ctx context.Context) ConfigFrequencyUnitsPtrOutput {
	return o
}

func (o ConfigFrequencyUnitsPtrOutput) Elem() ConfigFrequencyUnitsOutput {
	return o.ApplyT(func(v *ConfigFrequencyUnits) ConfigFrequencyUnits {
		if v != nil {
			return *v
		}
		var ret ConfigFrequencyUnits
		return ret
	}).(ConfigFrequencyUnitsOutput)
}

func (o ConfigFrequencyUnitsPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ConfigFrequencyUnitsPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ConfigFrequencyUnits) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ConfigFrequencyUnitsInput is an input type that accepts ConfigFrequencyUnitsArgs and ConfigFrequencyUnitsOutput values.
// You can construct a concrete instance of `ConfigFrequencyUnitsInput` via:
//
//	ConfigFrequencyUnitsArgs{...}
type ConfigFrequencyUnitsInput interface {
	pulumi.Input

	ToConfigFrequencyUnitsOutput() ConfigFrequencyUnitsOutput
	ToConfigFrequencyUnitsOutputWithContext(context.Context) ConfigFrequencyUnitsOutput
}

var configFrequencyUnitsPtrType = reflect.TypeOf((**ConfigFrequencyUnits)(nil)).Elem()

type ConfigFrequencyUnitsPtrInput interface {
	pulumi.Input

	ToConfigFrequencyUnitsPtrOutput() ConfigFrequencyUnitsPtrOutput
	ToConfigFrequencyUnitsPtrOutputWithContext(context.Context) ConfigFrequencyUnitsPtrOutput
}

type configFrequencyUnitsPtr string

func ConfigFrequencyUnitsPtr(v string) ConfigFrequencyUnitsPtrInput {
	return (*configFrequencyUnitsPtr)(&v)
}

func (*configFrequencyUnitsPtr) ElementType() reflect.Type {
	return configFrequencyUnitsPtrType
}

func (in *configFrequencyUnitsPtr) ToConfigFrequencyUnitsPtrOutput() ConfigFrequencyUnitsPtrOutput {
	return pulumi.ToOutput(in).(ConfigFrequencyUnitsPtrOutput)
}

func (in *configFrequencyUnitsPtr) ToConfigFrequencyUnitsPtrOutputWithContext(ctx context.Context) ConfigFrequencyUnitsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ConfigFrequencyUnitsPtrOutput)
}

type ConfigPolarization string

const (
	ConfigPolarizationLeftHand  = ConfigPolarization("LEFT_HAND")
	ConfigPolarizationRightHand = ConfigPolarization("RIGHT_HAND")
	ConfigPolarizationNone      = ConfigPolarization("NONE")
)

func (ConfigPolarization) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigPolarization)(nil)).Elem()
}

func (e ConfigPolarization) ToConfigPolarizationOutput() ConfigPolarizationOutput {
	return pulumi.ToOutput(e).(ConfigPolarizationOutput)
}

func (e ConfigPolarization) ToConfigPolarizationOutputWithContext(ctx context.Context) ConfigPolarizationOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ConfigPolarizationOutput)
}

func (e ConfigPolarization) ToConfigPolarizationPtrOutput() ConfigPolarizationPtrOutput {
	return e.ToConfigPolarizationPtrOutputWithContext(context.Background())
}

func (e ConfigPolarization) ToConfigPolarizationPtrOutputWithContext(ctx context.Context) ConfigPolarizationPtrOutput {
	return ConfigPolarization(e).ToConfigPolarizationOutputWithContext(ctx).ToConfigPolarizationPtrOutputWithContext(ctx)
}

func (e ConfigPolarization) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ConfigPolarization) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ConfigPolarization) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ConfigPolarization) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ConfigPolarizationOutput struct{ *pulumi.OutputState }

func (ConfigPolarizationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigPolarization)(nil)).Elem()
}

func (o ConfigPolarizationOutput) ToConfigPolarizationOutput() ConfigPolarizationOutput {
	return o
}

func (o ConfigPolarizationOutput) ToConfigPolarizationOutputWithContext(ctx context.Context) ConfigPolarizationOutput {
	return o
}

func (o ConfigPolarizationOutput) ToConfigPolarizationPtrOutput() ConfigPolarizationPtrOutput {
	return o.ToConfigPolarizationPtrOutputWithContext(context.Background())
}

func (o ConfigPolarizationOutput) ToConfigPolarizationPtrOutputWithContext(ctx context.Context) ConfigPolarizationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConfigPolarization) *ConfigPolarization {
		return &v
	}).(ConfigPolarizationPtrOutput)
}

func (o ConfigPolarizationOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ConfigPolarizationOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ConfigPolarization) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ConfigPolarizationOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ConfigPolarizationOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ConfigPolarization) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ConfigPolarizationPtrOutput struct{ *pulumi.OutputState }

func (ConfigPolarizationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigPolarization)(nil)).Elem()
}

func (o ConfigPolarizationPtrOutput) ToConfigPolarizationPtrOutput() ConfigPolarizationPtrOutput {
	return o
}

func (o ConfigPolarizationPtrOutput) ToConfigPolarizationPtrOutputWithContext(ctx context.Context) ConfigPolarizationPtrOutput {
	return o
}

func (o ConfigPolarizationPtrOutput) Elem() ConfigPolarizationOutput {
	return o.ApplyT(func(v *ConfigPolarization) ConfigPolarization {
		if v != nil {
			return *v
		}
		var ret ConfigPolarization
		return ret
	}).(ConfigPolarizationOutput)
}

func (o ConfigPolarizationPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ConfigPolarizationPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ConfigPolarization) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ConfigPolarizationInput is an input type that accepts ConfigPolarizationArgs and ConfigPolarizationOutput values.
// You can construct a concrete instance of `ConfigPolarizationInput` via:
//
//	ConfigPolarizationArgs{...}
type ConfigPolarizationInput interface {
	pulumi.Input

	ToConfigPolarizationOutput() ConfigPolarizationOutput
	ToConfigPolarizationOutputWithContext(context.Context) ConfigPolarizationOutput
}

var configPolarizationPtrType = reflect.TypeOf((**ConfigPolarization)(nil)).Elem()

type ConfigPolarizationPtrInput interface {
	pulumi.Input

	ToConfigPolarizationPtrOutput() ConfigPolarizationPtrOutput
	ToConfigPolarizationPtrOutputWithContext(context.Context) ConfigPolarizationPtrOutput
}

type configPolarizationPtr string

func ConfigPolarizationPtr(v string) ConfigPolarizationPtrInput {
	return (*configPolarizationPtr)(&v)
}

func (*configPolarizationPtr) ElementType() reflect.Type {
	return configPolarizationPtrType
}

func (in *configPolarizationPtr) ToConfigPolarizationPtrOutput() ConfigPolarizationPtrOutput {
	return pulumi.ToOutput(in).(ConfigPolarizationPtrOutput)
}

func (in *configPolarizationPtr) ToConfigPolarizationPtrOutputWithContext(ctx context.Context) ConfigPolarizationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ConfigPolarizationPtrOutput)
}

type ConfigTrackingConfigAutotrack string

const (
	ConfigTrackingConfigAutotrackRequired  = ConfigTrackingConfigAutotrack("REQUIRED")
	ConfigTrackingConfigAutotrackPreferred = ConfigTrackingConfigAutotrack("PREFERRED")
	ConfigTrackingConfigAutotrackRemoved   = ConfigTrackingConfigAutotrack("REMOVED")
)

func (ConfigTrackingConfigAutotrack) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigTrackingConfigAutotrack)(nil)).Elem()
}

func (e ConfigTrackingConfigAutotrack) ToConfigTrackingConfigAutotrackOutput() ConfigTrackingConfigAutotrackOutput {
	return pulumi.ToOutput(e).(ConfigTrackingConfigAutotrackOutput)
}

func (e ConfigTrackingConfigAutotrack) ToConfigTrackingConfigAutotrackOutputWithContext(ctx context.Context) ConfigTrackingConfigAutotrackOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ConfigTrackingConfigAutotrackOutput)
}

func (e ConfigTrackingConfigAutotrack) ToConfigTrackingConfigAutotrackPtrOutput() ConfigTrackingConfigAutotrackPtrOutput {
	return e.ToConfigTrackingConfigAutotrackPtrOutputWithContext(context.Background())
}

func (e ConfigTrackingConfigAutotrack) ToConfigTrackingConfigAutotrackPtrOutputWithContext(ctx context.Context) ConfigTrackingConfigAutotrackPtrOutput {
	return ConfigTrackingConfigAutotrack(e).ToConfigTrackingConfigAutotrackOutputWithContext(ctx).ToConfigTrackingConfigAutotrackPtrOutputWithContext(ctx)
}

func (e ConfigTrackingConfigAutotrack) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ConfigTrackingConfigAutotrack) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ConfigTrackingConfigAutotrack) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ConfigTrackingConfigAutotrack) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ConfigTrackingConfigAutotrackOutput struct{ *pulumi.OutputState }

func (ConfigTrackingConfigAutotrackOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigTrackingConfigAutotrack)(nil)).Elem()
}

func (o ConfigTrackingConfigAutotrackOutput) ToConfigTrackingConfigAutotrackOutput() ConfigTrackingConfigAutotrackOutput {
	return o
}

func (o ConfigTrackingConfigAutotrackOutput) ToConfigTrackingConfigAutotrackOutputWithContext(ctx context.Context) ConfigTrackingConfigAutotrackOutput {
	return o
}

func (o ConfigTrackingConfigAutotrackOutput) ToConfigTrackingConfigAutotrackPtrOutput() ConfigTrackingConfigAutotrackPtrOutput {
	return o.ToConfigTrackingConfigAutotrackPtrOutputWithContext(context.Background())
}

func (o ConfigTrackingConfigAutotrackOutput) ToConfigTrackingConfigAutotrackPtrOutputWithContext(ctx context.Context) ConfigTrackingConfigAutotrackPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConfigTrackingConfigAutotrack) *ConfigTrackingConfigAutotrack {
		return &v
	}).(ConfigTrackingConfigAutotrackPtrOutput)
}

func (o ConfigTrackingConfigAutotrackOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ConfigTrackingConfigAutotrackOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ConfigTrackingConfigAutotrack) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ConfigTrackingConfigAutotrackOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ConfigTrackingConfigAutotrackOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ConfigTrackingConfigAutotrack) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ConfigTrackingConfigAutotrackPtrOutput struct{ *pulumi.OutputState }

func (ConfigTrackingConfigAutotrackPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigTrackingConfigAutotrack)(nil)).Elem()
}

func (o ConfigTrackingConfigAutotrackPtrOutput) ToConfigTrackingConfigAutotrackPtrOutput() ConfigTrackingConfigAutotrackPtrOutput {
	return o
}

func (o ConfigTrackingConfigAutotrackPtrOutput) ToConfigTrackingConfigAutotrackPtrOutputWithContext(ctx context.Context) ConfigTrackingConfigAutotrackPtrOutput {
	return o
}

func (o ConfigTrackingConfigAutotrackPtrOutput) Elem() ConfigTrackingConfigAutotrackOutput {
	return o.ApplyT(func(v *ConfigTrackingConfigAutotrack) ConfigTrackingConfigAutotrack {
		if v != nil {
			return *v
		}
		var ret ConfigTrackingConfigAutotrack
		return ret
	}).(ConfigTrackingConfigAutotrackOutput)
}

func (o ConfigTrackingConfigAutotrackPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ConfigTrackingConfigAutotrackPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ConfigTrackingConfigAutotrack) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ConfigTrackingConfigAutotrackInput is an input type that accepts ConfigTrackingConfigAutotrackArgs and ConfigTrackingConfigAutotrackOutput values.
// You can construct a concrete instance of `ConfigTrackingConfigAutotrackInput` via:
//
//	ConfigTrackingConfigAutotrackArgs{...}
type ConfigTrackingConfigAutotrackInput interface {
	pulumi.Input

	ToConfigTrackingConfigAutotrackOutput() ConfigTrackingConfigAutotrackOutput
	ToConfigTrackingConfigAutotrackOutputWithContext(context.Context) ConfigTrackingConfigAutotrackOutput
}

var configTrackingConfigAutotrackPtrType = reflect.TypeOf((**ConfigTrackingConfigAutotrack)(nil)).Elem()

type ConfigTrackingConfigAutotrackPtrInput interface {
	pulumi.Input

	ToConfigTrackingConfigAutotrackPtrOutput() ConfigTrackingConfigAutotrackPtrOutput
	ToConfigTrackingConfigAutotrackPtrOutputWithContext(context.Context) ConfigTrackingConfigAutotrackPtrOutput
}

type configTrackingConfigAutotrackPtr string

func ConfigTrackingConfigAutotrackPtr(v string) ConfigTrackingConfigAutotrackPtrInput {
	return (*configTrackingConfigAutotrackPtr)(&v)
}

func (*configTrackingConfigAutotrackPtr) ElementType() reflect.Type {
	return configTrackingConfigAutotrackPtrType
}

func (in *configTrackingConfigAutotrackPtr) ToConfigTrackingConfigAutotrackPtrOutput() ConfigTrackingConfigAutotrackPtrOutput {
	return pulumi.ToOutput(in).(ConfigTrackingConfigAutotrackPtrOutput)
}

func (in *configTrackingConfigAutotrackPtr) ToConfigTrackingConfigAutotrackPtrOutputWithContext(ctx context.Context) ConfigTrackingConfigAutotrackPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ConfigTrackingConfigAutotrackPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigBandwidthUnitsInput)(nil)).Elem(), ConfigBandwidthUnits("GHz"))
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigBandwidthUnitsPtrInput)(nil)).Elem(), ConfigBandwidthUnits("GHz"))
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigEirpUnitsInput)(nil)).Elem(), ConfigEirpUnits("dBW"))
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigEirpUnitsPtrInput)(nil)).Elem(), ConfigEirpUnits("dBW"))
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigFrequencyUnitsInput)(nil)).Elem(), ConfigFrequencyUnits("GHz"))
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigFrequencyUnitsPtrInput)(nil)).Elem(), ConfigFrequencyUnits("GHz"))
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigPolarizationInput)(nil)).Elem(), ConfigPolarization("LEFT_HAND"))
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigPolarizationPtrInput)(nil)).Elem(), ConfigPolarization("LEFT_HAND"))
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigTrackingConfigAutotrackInput)(nil)).Elem(), ConfigTrackingConfigAutotrack("REQUIRED"))
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigTrackingConfigAutotrackPtrInput)(nil)).Elem(), ConfigTrackingConfigAutotrack("REQUIRED"))
	pulumi.RegisterOutputType(ConfigBandwidthUnitsOutput{})
	pulumi.RegisterOutputType(ConfigBandwidthUnitsPtrOutput{})
	pulumi.RegisterOutputType(ConfigEirpUnitsOutput{})
	pulumi.RegisterOutputType(ConfigEirpUnitsPtrOutput{})
	pulumi.RegisterOutputType(ConfigFrequencyUnitsOutput{})
	pulumi.RegisterOutputType(ConfigFrequencyUnitsPtrOutput{})
	pulumi.RegisterOutputType(ConfigPolarizationOutput{})
	pulumi.RegisterOutputType(ConfigPolarizationPtrOutput{})
	pulumi.RegisterOutputType(ConfigTrackingConfigAutotrackOutput{})
	pulumi.RegisterOutputType(ConfigTrackingConfigAutotrackPtrOutput{})
}

// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package inspectorv2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CisScanConfigurationCisSecurityLevel string

const (
	CisScanConfigurationCisSecurityLevelLevel1 = CisScanConfigurationCisSecurityLevel("LEVEL_1")
	CisScanConfigurationCisSecurityLevelLevel2 = CisScanConfigurationCisSecurityLevel("LEVEL_2")
)

func (CisScanConfigurationCisSecurityLevel) ElementType() reflect.Type {
	return reflect.TypeOf((*CisScanConfigurationCisSecurityLevel)(nil)).Elem()
}

func (e CisScanConfigurationCisSecurityLevel) ToCisScanConfigurationCisSecurityLevelOutput() CisScanConfigurationCisSecurityLevelOutput {
	return pulumi.ToOutput(e).(CisScanConfigurationCisSecurityLevelOutput)
}

func (e CisScanConfigurationCisSecurityLevel) ToCisScanConfigurationCisSecurityLevelOutputWithContext(ctx context.Context) CisScanConfigurationCisSecurityLevelOutput {
	return pulumi.ToOutputWithContext(ctx, e).(CisScanConfigurationCisSecurityLevelOutput)
}

func (e CisScanConfigurationCisSecurityLevel) ToCisScanConfigurationCisSecurityLevelPtrOutput() CisScanConfigurationCisSecurityLevelPtrOutput {
	return e.ToCisScanConfigurationCisSecurityLevelPtrOutputWithContext(context.Background())
}

func (e CisScanConfigurationCisSecurityLevel) ToCisScanConfigurationCisSecurityLevelPtrOutputWithContext(ctx context.Context) CisScanConfigurationCisSecurityLevelPtrOutput {
	return CisScanConfigurationCisSecurityLevel(e).ToCisScanConfigurationCisSecurityLevelOutputWithContext(ctx).ToCisScanConfigurationCisSecurityLevelPtrOutputWithContext(ctx)
}

func (e CisScanConfigurationCisSecurityLevel) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e CisScanConfigurationCisSecurityLevel) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e CisScanConfigurationCisSecurityLevel) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e CisScanConfigurationCisSecurityLevel) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type CisScanConfigurationCisSecurityLevelOutput struct{ *pulumi.OutputState }

func (CisScanConfigurationCisSecurityLevelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CisScanConfigurationCisSecurityLevel)(nil)).Elem()
}

func (o CisScanConfigurationCisSecurityLevelOutput) ToCisScanConfigurationCisSecurityLevelOutput() CisScanConfigurationCisSecurityLevelOutput {
	return o
}

func (o CisScanConfigurationCisSecurityLevelOutput) ToCisScanConfigurationCisSecurityLevelOutputWithContext(ctx context.Context) CisScanConfigurationCisSecurityLevelOutput {
	return o
}

func (o CisScanConfigurationCisSecurityLevelOutput) ToCisScanConfigurationCisSecurityLevelPtrOutput() CisScanConfigurationCisSecurityLevelPtrOutput {
	return o.ToCisScanConfigurationCisSecurityLevelPtrOutputWithContext(context.Background())
}

func (o CisScanConfigurationCisSecurityLevelOutput) ToCisScanConfigurationCisSecurityLevelPtrOutputWithContext(ctx context.Context) CisScanConfigurationCisSecurityLevelPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CisScanConfigurationCisSecurityLevel) *CisScanConfigurationCisSecurityLevel {
		return &v
	}).(CisScanConfigurationCisSecurityLevelPtrOutput)
}

func (o CisScanConfigurationCisSecurityLevelOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o CisScanConfigurationCisSecurityLevelOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CisScanConfigurationCisSecurityLevel) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o CisScanConfigurationCisSecurityLevelOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CisScanConfigurationCisSecurityLevelOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CisScanConfigurationCisSecurityLevel) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type CisScanConfigurationCisSecurityLevelPtrOutput struct{ *pulumi.OutputState }

func (CisScanConfigurationCisSecurityLevelPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CisScanConfigurationCisSecurityLevel)(nil)).Elem()
}

func (o CisScanConfigurationCisSecurityLevelPtrOutput) ToCisScanConfigurationCisSecurityLevelPtrOutput() CisScanConfigurationCisSecurityLevelPtrOutput {
	return o
}

func (o CisScanConfigurationCisSecurityLevelPtrOutput) ToCisScanConfigurationCisSecurityLevelPtrOutputWithContext(ctx context.Context) CisScanConfigurationCisSecurityLevelPtrOutput {
	return o
}

func (o CisScanConfigurationCisSecurityLevelPtrOutput) Elem() CisScanConfigurationCisSecurityLevelOutput {
	return o.ApplyT(func(v *CisScanConfigurationCisSecurityLevel) CisScanConfigurationCisSecurityLevel {
		if v != nil {
			return *v
		}
		var ret CisScanConfigurationCisSecurityLevel
		return ret
	}).(CisScanConfigurationCisSecurityLevelOutput)
}

func (o CisScanConfigurationCisSecurityLevelPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CisScanConfigurationCisSecurityLevelPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *CisScanConfigurationCisSecurityLevel) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// CisScanConfigurationCisSecurityLevelInput is an input type that accepts values of the CisScanConfigurationCisSecurityLevel enum
// A concrete instance of `CisScanConfigurationCisSecurityLevelInput` can be one of the following:
//
//	CisScanConfigurationCisSecurityLevelLevel1
//	CisScanConfigurationCisSecurityLevelLevel2
type CisScanConfigurationCisSecurityLevelInput interface {
	pulumi.Input

	ToCisScanConfigurationCisSecurityLevelOutput() CisScanConfigurationCisSecurityLevelOutput
	ToCisScanConfigurationCisSecurityLevelOutputWithContext(context.Context) CisScanConfigurationCisSecurityLevelOutput
}

var cisScanConfigurationCisSecurityLevelPtrType = reflect.TypeOf((**CisScanConfigurationCisSecurityLevel)(nil)).Elem()

type CisScanConfigurationCisSecurityLevelPtrInput interface {
	pulumi.Input

	ToCisScanConfigurationCisSecurityLevelPtrOutput() CisScanConfigurationCisSecurityLevelPtrOutput
	ToCisScanConfigurationCisSecurityLevelPtrOutputWithContext(context.Context) CisScanConfigurationCisSecurityLevelPtrOutput
}

type cisScanConfigurationCisSecurityLevelPtr string

func CisScanConfigurationCisSecurityLevelPtr(v string) CisScanConfigurationCisSecurityLevelPtrInput {
	return (*cisScanConfigurationCisSecurityLevelPtr)(&v)
}

func (*cisScanConfigurationCisSecurityLevelPtr) ElementType() reflect.Type {
	return cisScanConfigurationCisSecurityLevelPtrType
}

func (in *cisScanConfigurationCisSecurityLevelPtr) ToCisScanConfigurationCisSecurityLevelPtrOutput() CisScanConfigurationCisSecurityLevelPtrOutput {
	return pulumi.ToOutput(in).(CisScanConfigurationCisSecurityLevelPtrOutput)
}

func (in *cisScanConfigurationCisSecurityLevelPtr) ToCisScanConfigurationCisSecurityLevelPtrOutputWithContext(ctx context.Context) CisScanConfigurationCisSecurityLevelPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(CisScanConfigurationCisSecurityLevelPtrOutput)
}

type CisScanConfigurationDay string

const (
	CisScanConfigurationDayMon = CisScanConfigurationDay("MON")
	CisScanConfigurationDayTue = CisScanConfigurationDay("TUE")
	CisScanConfigurationDayWed = CisScanConfigurationDay("WED")
	CisScanConfigurationDayThu = CisScanConfigurationDay("THU")
	CisScanConfigurationDayFri = CisScanConfigurationDay("FRI")
	CisScanConfigurationDaySat = CisScanConfigurationDay("SAT")
	CisScanConfigurationDaySun = CisScanConfigurationDay("SUN")
)

func (CisScanConfigurationDay) ElementType() reflect.Type {
	return reflect.TypeOf((*CisScanConfigurationDay)(nil)).Elem()
}

func (e CisScanConfigurationDay) ToCisScanConfigurationDayOutput() CisScanConfigurationDayOutput {
	return pulumi.ToOutput(e).(CisScanConfigurationDayOutput)
}

func (e CisScanConfigurationDay) ToCisScanConfigurationDayOutputWithContext(ctx context.Context) CisScanConfigurationDayOutput {
	return pulumi.ToOutputWithContext(ctx, e).(CisScanConfigurationDayOutput)
}

func (e CisScanConfigurationDay) ToCisScanConfigurationDayPtrOutput() CisScanConfigurationDayPtrOutput {
	return e.ToCisScanConfigurationDayPtrOutputWithContext(context.Background())
}

func (e CisScanConfigurationDay) ToCisScanConfigurationDayPtrOutputWithContext(ctx context.Context) CisScanConfigurationDayPtrOutput {
	return CisScanConfigurationDay(e).ToCisScanConfigurationDayOutputWithContext(ctx).ToCisScanConfigurationDayPtrOutputWithContext(ctx)
}

func (e CisScanConfigurationDay) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e CisScanConfigurationDay) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e CisScanConfigurationDay) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e CisScanConfigurationDay) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type CisScanConfigurationDayOutput struct{ *pulumi.OutputState }

func (CisScanConfigurationDayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CisScanConfigurationDay)(nil)).Elem()
}

func (o CisScanConfigurationDayOutput) ToCisScanConfigurationDayOutput() CisScanConfigurationDayOutput {
	return o
}

func (o CisScanConfigurationDayOutput) ToCisScanConfigurationDayOutputWithContext(ctx context.Context) CisScanConfigurationDayOutput {
	return o
}

func (o CisScanConfigurationDayOutput) ToCisScanConfigurationDayPtrOutput() CisScanConfigurationDayPtrOutput {
	return o.ToCisScanConfigurationDayPtrOutputWithContext(context.Background())
}

func (o CisScanConfigurationDayOutput) ToCisScanConfigurationDayPtrOutputWithContext(ctx context.Context) CisScanConfigurationDayPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CisScanConfigurationDay) *CisScanConfigurationDay {
		return &v
	}).(CisScanConfigurationDayPtrOutput)
}

func (o CisScanConfigurationDayOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o CisScanConfigurationDayOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CisScanConfigurationDay) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o CisScanConfigurationDayOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CisScanConfigurationDayOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CisScanConfigurationDay) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type CisScanConfigurationDayPtrOutput struct{ *pulumi.OutputState }

func (CisScanConfigurationDayPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CisScanConfigurationDay)(nil)).Elem()
}

func (o CisScanConfigurationDayPtrOutput) ToCisScanConfigurationDayPtrOutput() CisScanConfigurationDayPtrOutput {
	return o
}

func (o CisScanConfigurationDayPtrOutput) ToCisScanConfigurationDayPtrOutputWithContext(ctx context.Context) CisScanConfigurationDayPtrOutput {
	return o
}

func (o CisScanConfigurationDayPtrOutput) Elem() CisScanConfigurationDayOutput {
	return o.ApplyT(func(v *CisScanConfigurationDay) CisScanConfigurationDay {
		if v != nil {
			return *v
		}
		var ret CisScanConfigurationDay
		return ret
	}).(CisScanConfigurationDayOutput)
}

func (o CisScanConfigurationDayPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CisScanConfigurationDayPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *CisScanConfigurationDay) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// CisScanConfigurationDayInput is an input type that accepts values of the CisScanConfigurationDay enum
// A concrete instance of `CisScanConfigurationDayInput` can be one of the following:
//
//	CisScanConfigurationDayMon
//	CisScanConfigurationDayTue
//	CisScanConfigurationDayWed
//	CisScanConfigurationDayThu
//	CisScanConfigurationDayFri
//	CisScanConfigurationDaySat
//	CisScanConfigurationDaySun
type CisScanConfigurationDayInput interface {
	pulumi.Input

	ToCisScanConfigurationDayOutput() CisScanConfigurationDayOutput
	ToCisScanConfigurationDayOutputWithContext(context.Context) CisScanConfigurationDayOutput
}

var cisScanConfigurationDayPtrType = reflect.TypeOf((**CisScanConfigurationDay)(nil)).Elem()

type CisScanConfigurationDayPtrInput interface {
	pulumi.Input

	ToCisScanConfigurationDayPtrOutput() CisScanConfigurationDayPtrOutput
	ToCisScanConfigurationDayPtrOutputWithContext(context.Context) CisScanConfigurationDayPtrOutput
}

type cisScanConfigurationDayPtr string

func CisScanConfigurationDayPtr(v string) CisScanConfigurationDayPtrInput {
	return (*cisScanConfigurationDayPtr)(&v)
}

func (*cisScanConfigurationDayPtr) ElementType() reflect.Type {
	return cisScanConfigurationDayPtrType
}

func (in *cisScanConfigurationDayPtr) ToCisScanConfigurationDayPtrOutput() CisScanConfigurationDayPtrOutput {
	return pulumi.ToOutput(in).(CisScanConfigurationDayPtrOutput)
}

func (in *cisScanConfigurationDayPtr) ToCisScanConfigurationDayPtrOutputWithContext(ctx context.Context) CisScanConfigurationDayPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(CisScanConfigurationDayPtrOutput)
}

// CisScanConfigurationDayArrayInput is an input type that accepts CisScanConfigurationDayArray and CisScanConfigurationDayArrayOutput values.
// You can construct a concrete instance of `CisScanConfigurationDayArrayInput` via:
//
//	CisScanConfigurationDayArray{ CisScanConfigurationDayArgs{...} }
type CisScanConfigurationDayArrayInput interface {
	pulumi.Input

	ToCisScanConfigurationDayArrayOutput() CisScanConfigurationDayArrayOutput
	ToCisScanConfigurationDayArrayOutputWithContext(context.Context) CisScanConfigurationDayArrayOutput
}

type CisScanConfigurationDayArray []CisScanConfigurationDay

func (CisScanConfigurationDayArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CisScanConfigurationDay)(nil)).Elem()
}

func (i CisScanConfigurationDayArray) ToCisScanConfigurationDayArrayOutput() CisScanConfigurationDayArrayOutput {
	return i.ToCisScanConfigurationDayArrayOutputWithContext(context.Background())
}

func (i CisScanConfigurationDayArray) ToCisScanConfigurationDayArrayOutputWithContext(ctx context.Context) CisScanConfigurationDayArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CisScanConfigurationDayArrayOutput)
}

type CisScanConfigurationDayArrayOutput struct{ *pulumi.OutputState }

func (CisScanConfigurationDayArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CisScanConfigurationDay)(nil)).Elem()
}

func (o CisScanConfigurationDayArrayOutput) ToCisScanConfigurationDayArrayOutput() CisScanConfigurationDayArrayOutput {
	return o
}

func (o CisScanConfigurationDayArrayOutput) ToCisScanConfigurationDayArrayOutputWithContext(ctx context.Context) CisScanConfigurationDayArrayOutput {
	return o
}

func (o CisScanConfigurationDayArrayOutput) Index(i pulumi.IntInput) CisScanConfigurationDayOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CisScanConfigurationDay {
		return vs[0].([]CisScanConfigurationDay)[vs[1].(int)]
	}).(CisScanConfigurationDayOutput)
}

type FilterAction string

const (
	FilterActionNone     = FilterAction("NONE")
	FilterActionSuppress = FilterAction("SUPPRESS")
)

func (FilterAction) ElementType() reflect.Type {
	return reflect.TypeOf((*FilterAction)(nil)).Elem()
}

func (e FilterAction) ToFilterActionOutput() FilterActionOutput {
	return pulumi.ToOutput(e).(FilterActionOutput)
}

func (e FilterAction) ToFilterActionOutputWithContext(ctx context.Context) FilterActionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(FilterActionOutput)
}

func (e FilterAction) ToFilterActionPtrOutput() FilterActionPtrOutput {
	return e.ToFilterActionPtrOutputWithContext(context.Background())
}

func (e FilterAction) ToFilterActionPtrOutputWithContext(ctx context.Context) FilterActionPtrOutput {
	return FilterAction(e).ToFilterActionOutputWithContext(ctx).ToFilterActionPtrOutputWithContext(ctx)
}

func (e FilterAction) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e FilterAction) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e FilterAction) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e FilterAction) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type FilterActionOutput struct{ *pulumi.OutputState }

func (FilterActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FilterAction)(nil)).Elem()
}

func (o FilterActionOutput) ToFilterActionOutput() FilterActionOutput {
	return o
}

func (o FilterActionOutput) ToFilterActionOutputWithContext(ctx context.Context) FilterActionOutput {
	return o
}

func (o FilterActionOutput) ToFilterActionPtrOutput() FilterActionPtrOutput {
	return o.ToFilterActionPtrOutputWithContext(context.Background())
}

func (o FilterActionOutput) ToFilterActionPtrOutputWithContext(ctx context.Context) FilterActionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FilterAction) *FilterAction {
		return &v
	}).(FilterActionPtrOutput)
}

func (o FilterActionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o FilterActionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e FilterAction) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o FilterActionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o FilterActionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e FilterAction) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type FilterActionPtrOutput struct{ *pulumi.OutputState }

func (FilterActionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FilterAction)(nil)).Elem()
}

func (o FilterActionPtrOutput) ToFilterActionPtrOutput() FilterActionPtrOutput {
	return o
}

func (o FilterActionPtrOutput) ToFilterActionPtrOutputWithContext(ctx context.Context) FilterActionPtrOutput {
	return o
}

func (o FilterActionPtrOutput) Elem() FilterActionOutput {
	return o.ApplyT(func(v *FilterAction) FilterAction {
		if v != nil {
			return *v
		}
		var ret FilterAction
		return ret
	}).(FilterActionOutput)
}

func (o FilterActionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o FilterActionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *FilterAction) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// FilterActionInput is an input type that accepts values of the FilterAction enum
// A concrete instance of `FilterActionInput` can be one of the following:
//
//	FilterActionNone
//	FilterActionSuppress
type FilterActionInput interface {
	pulumi.Input

	ToFilterActionOutput() FilterActionOutput
	ToFilterActionOutputWithContext(context.Context) FilterActionOutput
}

var filterActionPtrType = reflect.TypeOf((**FilterAction)(nil)).Elem()

type FilterActionPtrInput interface {
	pulumi.Input

	ToFilterActionPtrOutput() FilterActionPtrOutput
	ToFilterActionPtrOutputWithContext(context.Context) FilterActionPtrOutput
}

type filterActionPtr string

func FilterActionPtr(v string) FilterActionPtrInput {
	return (*filterActionPtr)(&v)
}

func (*filterActionPtr) ElementType() reflect.Type {
	return filterActionPtrType
}

func (in *filterActionPtr) ToFilterActionPtrOutput() FilterActionPtrOutput {
	return pulumi.ToOutput(in).(FilterActionPtrOutput)
}

func (in *filterActionPtr) ToFilterActionPtrOutputWithContext(ctx context.Context) FilterActionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(FilterActionPtrOutput)
}

type FilterMapComparison string

const (
	FilterMapComparisonEquals = FilterMapComparison("EQUALS")
)

func (FilterMapComparison) ElementType() reflect.Type {
	return reflect.TypeOf((*FilterMapComparison)(nil)).Elem()
}

func (e FilterMapComparison) ToFilterMapComparisonOutput() FilterMapComparisonOutput {
	return pulumi.ToOutput(e).(FilterMapComparisonOutput)
}

func (e FilterMapComparison) ToFilterMapComparisonOutputWithContext(ctx context.Context) FilterMapComparisonOutput {
	return pulumi.ToOutputWithContext(ctx, e).(FilterMapComparisonOutput)
}

func (e FilterMapComparison) ToFilterMapComparisonPtrOutput() FilterMapComparisonPtrOutput {
	return e.ToFilterMapComparisonPtrOutputWithContext(context.Background())
}

func (e FilterMapComparison) ToFilterMapComparisonPtrOutputWithContext(ctx context.Context) FilterMapComparisonPtrOutput {
	return FilterMapComparison(e).ToFilterMapComparisonOutputWithContext(ctx).ToFilterMapComparisonPtrOutputWithContext(ctx)
}

func (e FilterMapComparison) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e FilterMapComparison) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e FilterMapComparison) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e FilterMapComparison) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type FilterMapComparisonOutput struct{ *pulumi.OutputState }

func (FilterMapComparisonOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FilterMapComparison)(nil)).Elem()
}

func (o FilterMapComparisonOutput) ToFilterMapComparisonOutput() FilterMapComparisonOutput {
	return o
}

func (o FilterMapComparisonOutput) ToFilterMapComparisonOutputWithContext(ctx context.Context) FilterMapComparisonOutput {
	return o
}

func (o FilterMapComparisonOutput) ToFilterMapComparisonPtrOutput() FilterMapComparisonPtrOutput {
	return o.ToFilterMapComparisonPtrOutputWithContext(context.Background())
}

func (o FilterMapComparisonOutput) ToFilterMapComparisonPtrOutputWithContext(ctx context.Context) FilterMapComparisonPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FilterMapComparison) *FilterMapComparison {
		return &v
	}).(FilterMapComparisonPtrOutput)
}

func (o FilterMapComparisonOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o FilterMapComparisonOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e FilterMapComparison) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o FilterMapComparisonOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o FilterMapComparisonOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e FilterMapComparison) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type FilterMapComparisonPtrOutput struct{ *pulumi.OutputState }

func (FilterMapComparisonPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FilterMapComparison)(nil)).Elem()
}

func (o FilterMapComparisonPtrOutput) ToFilterMapComparisonPtrOutput() FilterMapComparisonPtrOutput {
	return o
}

func (o FilterMapComparisonPtrOutput) ToFilterMapComparisonPtrOutputWithContext(ctx context.Context) FilterMapComparisonPtrOutput {
	return o
}

func (o FilterMapComparisonPtrOutput) Elem() FilterMapComparisonOutput {
	return o.ApplyT(func(v *FilterMapComparison) FilterMapComparison {
		if v != nil {
			return *v
		}
		var ret FilterMapComparison
		return ret
	}).(FilterMapComparisonOutput)
}

func (o FilterMapComparisonPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o FilterMapComparisonPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *FilterMapComparison) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// FilterMapComparisonInput is an input type that accepts values of the FilterMapComparison enum
// A concrete instance of `FilterMapComparisonInput` can be one of the following:
//
//	FilterMapComparisonEquals
type FilterMapComparisonInput interface {
	pulumi.Input

	ToFilterMapComparisonOutput() FilterMapComparisonOutput
	ToFilterMapComparisonOutputWithContext(context.Context) FilterMapComparisonOutput
}

var filterMapComparisonPtrType = reflect.TypeOf((**FilterMapComparison)(nil)).Elem()

type FilterMapComparisonPtrInput interface {
	pulumi.Input

	ToFilterMapComparisonPtrOutput() FilterMapComparisonPtrOutput
	ToFilterMapComparisonPtrOutputWithContext(context.Context) FilterMapComparisonPtrOutput
}

type filterMapComparisonPtr string

func FilterMapComparisonPtr(v string) FilterMapComparisonPtrInput {
	return (*filterMapComparisonPtr)(&v)
}

func (*filterMapComparisonPtr) ElementType() reflect.Type {
	return filterMapComparisonPtrType
}

func (in *filterMapComparisonPtr) ToFilterMapComparisonPtrOutput() FilterMapComparisonPtrOutput {
	return pulumi.ToOutput(in).(FilterMapComparisonPtrOutput)
}

func (in *filterMapComparisonPtr) ToFilterMapComparisonPtrOutputWithContext(ctx context.Context) FilterMapComparisonPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(FilterMapComparisonPtrOutput)
}

type FilterStringComparison string

const (
	FilterStringComparisonEquals    = FilterStringComparison("EQUALS")
	FilterStringComparisonPrefix    = FilterStringComparison("PREFIX")
	FilterStringComparisonNotEquals = FilterStringComparison("NOT_EQUALS")
)

func (FilterStringComparison) ElementType() reflect.Type {
	return reflect.TypeOf((*FilterStringComparison)(nil)).Elem()
}

func (e FilterStringComparison) ToFilterStringComparisonOutput() FilterStringComparisonOutput {
	return pulumi.ToOutput(e).(FilterStringComparisonOutput)
}

func (e FilterStringComparison) ToFilterStringComparisonOutputWithContext(ctx context.Context) FilterStringComparisonOutput {
	return pulumi.ToOutputWithContext(ctx, e).(FilterStringComparisonOutput)
}

func (e FilterStringComparison) ToFilterStringComparisonPtrOutput() FilterStringComparisonPtrOutput {
	return e.ToFilterStringComparisonPtrOutputWithContext(context.Background())
}

func (e FilterStringComparison) ToFilterStringComparisonPtrOutputWithContext(ctx context.Context) FilterStringComparisonPtrOutput {
	return FilterStringComparison(e).ToFilterStringComparisonOutputWithContext(ctx).ToFilterStringComparisonPtrOutputWithContext(ctx)
}

func (e FilterStringComparison) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e FilterStringComparison) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e FilterStringComparison) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e FilterStringComparison) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type FilterStringComparisonOutput struct{ *pulumi.OutputState }

func (FilterStringComparisonOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FilterStringComparison)(nil)).Elem()
}

func (o FilterStringComparisonOutput) ToFilterStringComparisonOutput() FilterStringComparisonOutput {
	return o
}

func (o FilterStringComparisonOutput) ToFilterStringComparisonOutputWithContext(ctx context.Context) FilterStringComparisonOutput {
	return o
}

func (o FilterStringComparisonOutput) ToFilterStringComparisonPtrOutput() FilterStringComparisonPtrOutput {
	return o.ToFilterStringComparisonPtrOutputWithContext(context.Background())
}

func (o FilterStringComparisonOutput) ToFilterStringComparisonPtrOutputWithContext(ctx context.Context) FilterStringComparisonPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FilterStringComparison) *FilterStringComparison {
		return &v
	}).(FilterStringComparisonPtrOutput)
}

func (o FilterStringComparisonOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o FilterStringComparisonOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e FilterStringComparison) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o FilterStringComparisonOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o FilterStringComparisonOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e FilterStringComparison) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type FilterStringComparisonPtrOutput struct{ *pulumi.OutputState }

func (FilterStringComparisonPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FilterStringComparison)(nil)).Elem()
}

func (o FilterStringComparisonPtrOutput) ToFilterStringComparisonPtrOutput() FilterStringComparisonPtrOutput {
	return o
}

func (o FilterStringComparisonPtrOutput) ToFilterStringComparisonPtrOutputWithContext(ctx context.Context) FilterStringComparisonPtrOutput {
	return o
}

func (o FilterStringComparisonPtrOutput) Elem() FilterStringComparisonOutput {
	return o.ApplyT(func(v *FilterStringComparison) FilterStringComparison {
		if v != nil {
			return *v
		}
		var ret FilterStringComparison
		return ret
	}).(FilterStringComparisonOutput)
}

func (o FilterStringComparisonPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o FilterStringComparisonPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *FilterStringComparison) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// FilterStringComparisonInput is an input type that accepts values of the FilterStringComparison enum
// A concrete instance of `FilterStringComparisonInput` can be one of the following:
//
//	FilterStringComparisonEquals
//	FilterStringComparisonPrefix
//	FilterStringComparisonNotEquals
type FilterStringComparisonInput interface {
	pulumi.Input

	ToFilterStringComparisonOutput() FilterStringComparisonOutput
	ToFilterStringComparisonOutputWithContext(context.Context) FilterStringComparisonOutput
}

var filterStringComparisonPtrType = reflect.TypeOf((**FilterStringComparison)(nil)).Elem()

type FilterStringComparisonPtrInput interface {
	pulumi.Input

	ToFilterStringComparisonPtrOutput() FilterStringComparisonPtrOutput
	ToFilterStringComparisonPtrOutputWithContext(context.Context) FilterStringComparisonPtrOutput
}

type filterStringComparisonPtr string

func FilterStringComparisonPtr(v string) FilterStringComparisonPtrInput {
	return (*filterStringComparisonPtr)(&v)
}

func (*filterStringComparisonPtr) ElementType() reflect.Type {
	return filterStringComparisonPtrType
}

func (in *filterStringComparisonPtr) ToFilterStringComparisonPtrOutput() FilterStringComparisonPtrOutput {
	return pulumi.ToOutput(in).(FilterStringComparisonPtrOutput)
}

func (in *filterStringComparisonPtr) ToFilterStringComparisonPtrOutputWithContext(ctx context.Context) FilterStringComparisonPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(FilterStringComparisonPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CisScanConfigurationCisSecurityLevelInput)(nil)).Elem(), CisScanConfigurationCisSecurityLevel("LEVEL_1"))
	pulumi.RegisterInputType(reflect.TypeOf((*CisScanConfigurationCisSecurityLevelPtrInput)(nil)).Elem(), CisScanConfigurationCisSecurityLevel("LEVEL_1"))
	pulumi.RegisterInputType(reflect.TypeOf((*CisScanConfigurationDayInput)(nil)).Elem(), CisScanConfigurationDay("MON"))
	pulumi.RegisterInputType(reflect.TypeOf((*CisScanConfigurationDayPtrInput)(nil)).Elem(), CisScanConfigurationDay("MON"))
	pulumi.RegisterInputType(reflect.TypeOf((*CisScanConfigurationDayArrayInput)(nil)).Elem(), CisScanConfigurationDayArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FilterActionInput)(nil)).Elem(), FilterAction("NONE"))
	pulumi.RegisterInputType(reflect.TypeOf((*FilterActionPtrInput)(nil)).Elem(), FilterAction("NONE"))
	pulumi.RegisterInputType(reflect.TypeOf((*FilterMapComparisonInput)(nil)).Elem(), FilterMapComparison("EQUALS"))
	pulumi.RegisterInputType(reflect.TypeOf((*FilterMapComparisonPtrInput)(nil)).Elem(), FilterMapComparison("EQUALS"))
	pulumi.RegisterInputType(reflect.TypeOf((*FilterStringComparisonInput)(nil)).Elem(), FilterStringComparison("EQUALS"))
	pulumi.RegisterInputType(reflect.TypeOf((*FilterStringComparisonPtrInput)(nil)).Elem(), FilterStringComparison("EQUALS"))
	pulumi.RegisterOutputType(CisScanConfigurationCisSecurityLevelOutput{})
	pulumi.RegisterOutputType(CisScanConfigurationCisSecurityLevelPtrOutput{})
	pulumi.RegisterOutputType(CisScanConfigurationDayOutput{})
	pulumi.RegisterOutputType(CisScanConfigurationDayPtrOutput{})
	pulumi.RegisterOutputType(CisScanConfigurationDayArrayOutput{})
	pulumi.RegisterOutputType(FilterActionOutput{})
	pulumi.RegisterOutputType(FilterActionPtrOutput{})
	pulumi.RegisterOutputType(FilterMapComparisonOutput{})
	pulumi.RegisterOutputType(FilterMapComparisonPtrOutput{})
	pulumi.RegisterOutputType(FilterStringComparisonOutput{})
	pulumi.RegisterOutputType(FilterStringComparisonPtrOutput{})
}

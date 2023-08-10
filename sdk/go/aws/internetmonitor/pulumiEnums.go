// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package internetmonitor

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type MonitorConfigState string

const (
	MonitorConfigStatePending  = MonitorConfigState("PENDING")
	MonitorConfigStateActive   = MonitorConfigState("ACTIVE")
	MonitorConfigStateInactive = MonitorConfigState("INACTIVE")
	MonitorConfigStateError    = MonitorConfigState("ERROR")
)

func (MonitorConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*MonitorConfigState)(nil)).Elem()
}

func (e MonitorConfigState) ToMonitorConfigStateOutput() MonitorConfigStateOutput {
	return pulumi.ToOutput(e).(MonitorConfigStateOutput)
}

func (e MonitorConfigState) ToMonitorConfigStateOutputWithContext(ctx context.Context) MonitorConfigStateOutput {
	return pulumi.ToOutputWithContext(ctx, e).(MonitorConfigStateOutput)
}

func (e MonitorConfigState) ToMonitorConfigStatePtrOutput() MonitorConfigStatePtrOutput {
	return e.ToMonitorConfigStatePtrOutputWithContext(context.Background())
}

func (e MonitorConfigState) ToMonitorConfigStatePtrOutputWithContext(ctx context.Context) MonitorConfigStatePtrOutput {
	return MonitorConfigState(e).ToMonitorConfigStateOutputWithContext(ctx).ToMonitorConfigStatePtrOutputWithContext(ctx)
}

func (e MonitorConfigState) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e MonitorConfigState) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e MonitorConfigState) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e MonitorConfigState) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type MonitorConfigStateOutput struct{ *pulumi.OutputState }

func (MonitorConfigStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MonitorConfigState)(nil)).Elem()
}

func (o MonitorConfigStateOutput) ToMonitorConfigStateOutput() MonitorConfigStateOutput {
	return o
}

func (o MonitorConfigStateOutput) ToMonitorConfigStateOutputWithContext(ctx context.Context) MonitorConfigStateOutput {
	return o
}

func (o MonitorConfigStateOutput) ToMonitorConfigStatePtrOutput() MonitorConfigStatePtrOutput {
	return o.ToMonitorConfigStatePtrOutputWithContext(context.Background())
}

func (o MonitorConfigStateOutput) ToMonitorConfigStatePtrOutputWithContext(ctx context.Context) MonitorConfigStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MonitorConfigState) *MonitorConfigState {
		return &v
	}).(MonitorConfigStatePtrOutput)
}

func (o MonitorConfigStateOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o MonitorConfigStateOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e MonitorConfigState) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o MonitorConfigStateOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o MonitorConfigStateOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e MonitorConfigState) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type MonitorConfigStatePtrOutput struct{ *pulumi.OutputState }

func (MonitorConfigStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MonitorConfigState)(nil)).Elem()
}

func (o MonitorConfigStatePtrOutput) ToMonitorConfigStatePtrOutput() MonitorConfigStatePtrOutput {
	return o
}

func (o MonitorConfigStatePtrOutput) ToMonitorConfigStatePtrOutputWithContext(ctx context.Context) MonitorConfigStatePtrOutput {
	return o
}

func (o MonitorConfigStatePtrOutput) Elem() MonitorConfigStateOutput {
	return o.ApplyT(func(v *MonitorConfigState) MonitorConfigState {
		if v != nil {
			return *v
		}
		var ret MonitorConfigState
		return ret
	}).(MonitorConfigStateOutput)
}

func (o MonitorConfigStatePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o MonitorConfigStatePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *MonitorConfigState) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// MonitorConfigStateInput is an input type that accepts MonitorConfigStateArgs and MonitorConfigStateOutput values.
// You can construct a concrete instance of `MonitorConfigStateInput` via:
//
//	MonitorConfigStateArgs{...}
type MonitorConfigStateInput interface {
	pulumi.Input

	ToMonitorConfigStateOutput() MonitorConfigStateOutput
	ToMonitorConfigStateOutputWithContext(context.Context) MonitorConfigStateOutput
}

var monitorConfigStatePtrType = reflect.TypeOf((**MonitorConfigState)(nil)).Elem()

type MonitorConfigStatePtrInput interface {
	pulumi.Input

	ToMonitorConfigStatePtrOutput() MonitorConfigStatePtrOutput
	ToMonitorConfigStatePtrOutputWithContext(context.Context) MonitorConfigStatePtrOutput
}

type monitorConfigStatePtr string

func MonitorConfigStatePtr(v string) MonitorConfigStatePtrInput {
	return (*monitorConfigStatePtr)(&v)
}

func (*monitorConfigStatePtr) ElementType() reflect.Type {
	return monitorConfigStatePtrType
}

func (in *monitorConfigStatePtr) ToMonitorConfigStatePtrOutput() MonitorConfigStatePtrOutput {
	return pulumi.ToOutput(in).(MonitorConfigStatePtrOutput)
}

func (in *monitorConfigStatePtr) ToMonitorConfigStatePtrOutputWithContext(ctx context.Context) MonitorConfigStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(MonitorConfigStatePtrOutput)
}

type MonitorLocalHealthEventsConfigStatus string

const (
	MonitorLocalHealthEventsConfigStatusEnabled  = MonitorLocalHealthEventsConfigStatus("ENABLED")
	MonitorLocalHealthEventsConfigStatusDisabled = MonitorLocalHealthEventsConfigStatus("DISABLED")
)

func (MonitorLocalHealthEventsConfigStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*MonitorLocalHealthEventsConfigStatus)(nil)).Elem()
}

func (e MonitorLocalHealthEventsConfigStatus) ToMonitorLocalHealthEventsConfigStatusOutput() MonitorLocalHealthEventsConfigStatusOutput {
	return pulumi.ToOutput(e).(MonitorLocalHealthEventsConfigStatusOutput)
}

func (e MonitorLocalHealthEventsConfigStatus) ToMonitorLocalHealthEventsConfigStatusOutputWithContext(ctx context.Context) MonitorLocalHealthEventsConfigStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(MonitorLocalHealthEventsConfigStatusOutput)
}

func (e MonitorLocalHealthEventsConfigStatus) ToMonitorLocalHealthEventsConfigStatusPtrOutput() MonitorLocalHealthEventsConfigStatusPtrOutput {
	return e.ToMonitorLocalHealthEventsConfigStatusPtrOutputWithContext(context.Background())
}

func (e MonitorLocalHealthEventsConfigStatus) ToMonitorLocalHealthEventsConfigStatusPtrOutputWithContext(ctx context.Context) MonitorLocalHealthEventsConfigStatusPtrOutput {
	return MonitorLocalHealthEventsConfigStatus(e).ToMonitorLocalHealthEventsConfigStatusOutputWithContext(ctx).ToMonitorLocalHealthEventsConfigStatusPtrOutputWithContext(ctx)
}

func (e MonitorLocalHealthEventsConfigStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e MonitorLocalHealthEventsConfigStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e MonitorLocalHealthEventsConfigStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e MonitorLocalHealthEventsConfigStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type MonitorLocalHealthEventsConfigStatusOutput struct{ *pulumi.OutputState }

func (MonitorLocalHealthEventsConfigStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MonitorLocalHealthEventsConfigStatus)(nil)).Elem()
}

func (o MonitorLocalHealthEventsConfigStatusOutput) ToMonitorLocalHealthEventsConfigStatusOutput() MonitorLocalHealthEventsConfigStatusOutput {
	return o
}

func (o MonitorLocalHealthEventsConfigStatusOutput) ToMonitorLocalHealthEventsConfigStatusOutputWithContext(ctx context.Context) MonitorLocalHealthEventsConfigStatusOutput {
	return o
}

func (o MonitorLocalHealthEventsConfigStatusOutput) ToMonitorLocalHealthEventsConfigStatusPtrOutput() MonitorLocalHealthEventsConfigStatusPtrOutput {
	return o.ToMonitorLocalHealthEventsConfigStatusPtrOutputWithContext(context.Background())
}

func (o MonitorLocalHealthEventsConfigStatusOutput) ToMonitorLocalHealthEventsConfigStatusPtrOutputWithContext(ctx context.Context) MonitorLocalHealthEventsConfigStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MonitorLocalHealthEventsConfigStatus) *MonitorLocalHealthEventsConfigStatus {
		return &v
	}).(MonitorLocalHealthEventsConfigStatusPtrOutput)
}

func (o MonitorLocalHealthEventsConfigStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o MonitorLocalHealthEventsConfigStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e MonitorLocalHealthEventsConfigStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o MonitorLocalHealthEventsConfigStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o MonitorLocalHealthEventsConfigStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e MonitorLocalHealthEventsConfigStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type MonitorLocalHealthEventsConfigStatusPtrOutput struct{ *pulumi.OutputState }

func (MonitorLocalHealthEventsConfigStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MonitorLocalHealthEventsConfigStatus)(nil)).Elem()
}

func (o MonitorLocalHealthEventsConfigStatusPtrOutput) ToMonitorLocalHealthEventsConfigStatusPtrOutput() MonitorLocalHealthEventsConfigStatusPtrOutput {
	return o
}

func (o MonitorLocalHealthEventsConfigStatusPtrOutput) ToMonitorLocalHealthEventsConfigStatusPtrOutputWithContext(ctx context.Context) MonitorLocalHealthEventsConfigStatusPtrOutput {
	return o
}

func (o MonitorLocalHealthEventsConfigStatusPtrOutput) Elem() MonitorLocalHealthEventsConfigStatusOutput {
	return o.ApplyT(func(v *MonitorLocalHealthEventsConfigStatus) MonitorLocalHealthEventsConfigStatus {
		if v != nil {
			return *v
		}
		var ret MonitorLocalHealthEventsConfigStatus
		return ret
	}).(MonitorLocalHealthEventsConfigStatusOutput)
}

func (o MonitorLocalHealthEventsConfigStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o MonitorLocalHealthEventsConfigStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *MonitorLocalHealthEventsConfigStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// MonitorLocalHealthEventsConfigStatusInput is an input type that accepts MonitorLocalHealthEventsConfigStatusArgs and MonitorLocalHealthEventsConfigStatusOutput values.
// You can construct a concrete instance of `MonitorLocalHealthEventsConfigStatusInput` via:
//
//	MonitorLocalHealthEventsConfigStatusArgs{...}
type MonitorLocalHealthEventsConfigStatusInput interface {
	pulumi.Input

	ToMonitorLocalHealthEventsConfigStatusOutput() MonitorLocalHealthEventsConfigStatusOutput
	ToMonitorLocalHealthEventsConfigStatusOutputWithContext(context.Context) MonitorLocalHealthEventsConfigStatusOutput
}

var monitorLocalHealthEventsConfigStatusPtrType = reflect.TypeOf((**MonitorLocalHealthEventsConfigStatus)(nil)).Elem()

type MonitorLocalHealthEventsConfigStatusPtrInput interface {
	pulumi.Input

	ToMonitorLocalHealthEventsConfigStatusPtrOutput() MonitorLocalHealthEventsConfigStatusPtrOutput
	ToMonitorLocalHealthEventsConfigStatusPtrOutputWithContext(context.Context) MonitorLocalHealthEventsConfigStatusPtrOutput
}

type monitorLocalHealthEventsConfigStatusPtr string

func MonitorLocalHealthEventsConfigStatusPtr(v string) MonitorLocalHealthEventsConfigStatusPtrInput {
	return (*monitorLocalHealthEventsConfigStatusPtr)(&v)
}

func (*monitorLocalHealthEventsConfigStatusPtr) ElementType() reflect.Type {
	return monitorLocalHealthEventsConfigStatusPtrType
}

func (in *monitorLocalHealthEventsConfigStatusPtr) ToMonitorLocalHealthEventsConfigStatusPtrOutput() MonitorLocalHealthEventsConfigStatusPtrOutput {
	return pulumi.ToOutput(in).(MonitorLocalHealthEventsConfigStatusPtrOutput)
}

func (in *monitorLocalHealthEventsConfigStatusPtr) ToMonitorLocalHealthEventsConfigStatusPtrOutputWithContext(ctx context.Context) MonitorLocalHealthEventsConfigStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(MonitorLocalHealthEventsConfigStatusPtrOutput)
}

type MonitorProcessingStatusCode string

const (
	MonitorProcessingStatusCodeOk                    = MonitorProcessingStatusCode("OK")
	MonitorProcessingStatusCodeInactive              = MonitorProcessingStatusCode("INACTIVE")
	MonitorProcessingStatusCodeCollectingData        = MonitorProcessingStatusCode("COLLECTING_DATA")
	MonitorProcessingStatusCodeInsufficientData      = MonitorProcessingStatusCode("INSUFFICIENT_DATA")
	MonitorProcessingStatusCodeFaultService          = MonitorProcessingStatusCode("FAULT_SERVICE")
	MonitorProcessingStatusCodeFaultAccessCloudwatch = MonitorProcessingStatusCode("FAULT_ACCESS_CLOUDWATCH")
)

type MonitorProcessingStatusCodeOutput struct{ *pulumi.OutputState }

func (MonitorProcessingStatusCodeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MonitorProcessingStatusCode)(nil)).Elem()
}

func (o MonitorProcessingStatusCodeOutput) ToMonitorProcessingStatusCodeOutput() MonitorProcessingStatusCodeOutput {
	return o
}

func (o MonitorProcessingStatusCodeOutput) ToMonitorProcessingStatusCodeOutputWithContext(ctx context.Context) MonitorProcessingStatusCodeOutput {
	return o
}

func (o MonitorProcessingStatusCodeOutput) ToMonitorProcessingStatusCodePtrOutput() MonitorProcessingStatusCodePtrOutput {
	return o.ToMonitorProcessingStatusCodePtrOutputWithContext(context.Background())
}

func (o MonitorProcessingStatusCodeOutput) ToMonitorProcessingStatusCodePtrOutputWithContext(ctx context.Context) MonitorProcessingStatusCodePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MonitorProcessingStatusCode) *MonitorProcessingStatusCode {
		return &v
	}).(MonitorProcessingStatusCodePtrOutput)
}

func (o MonitorProcessingStatusCodeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o MonitorProcessingStatusCodeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e MonitorProcessingStatusCode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o MonitorProcessingStatusCodeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o MonitorProcessingStatusCodeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e MonitorProcessingStatusCode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type MonitorProcessingStatusCodePtrOutput struct{ *pulumi.OutputState }

func (MonitorProcessingStatusCodePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MonitorProcessingStatusCode)(nil)).Elem()
}

func (o MonitorProcessingStatusCodePtrOutput) ToMonitorProcessingStatusCodePtrOutput() MonitorProcessingStatusCodePtrOutput {
	return o
}

func (o MonitorProcessingStatusCodePtrOutput) ToMonitorProcessingStatusCodePtrOutputWithContext(ctx context.Context) MonitorProcessingStatusCodePtrOutput {
	return o
}

func (o MonitorProcessingStatusCodePtrOutput) Elem() MonitorProcessingStatusCodeOutput {
	return o.ApplyT(func(v *MonitorProcessingStatusCode) MonitorProcessingStatusCode {
		if v != nil {
			return *v
		}
		var ret MonitorProcessingStatusCode
		return ret
	}).(MonitorProcessingStatusCodeOutput)
}

func (o MonitorProcessingStatusCodePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o MonitorProcessingStatusCodePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *MonitorProcessingStatusCode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type MonitorS3ConfigLogDeliveryStatus string

const (
	MonitorS3ConfigLogDeliveryStatusEnabled  = MonitorS3ConfigLogDeliveryStatus("ENABLED")
	MonitorS3ConfigLogDeliveryStatusDisabled = MonitorS3ConfigLogDeliveryStatus("DISABLED")
)

func (MonitorS3ConfigLogDeliveryStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*MonitorS3ConfigLogDeliveryStatus)(nil)).Elem()
}

func (e MonitorS3ConfigLogDeliveryStatus) ToMonitorS3ConfigLogDeliveryStatusOutput() MonitorS3ConfigLogDeliveryStatusOutput {
	return pulumi.ToOutput(e).(MonitorS3ConfigLogDeliveryStatusOutput)
}

func (e MonitorS3ConfigLogDeliveryStatus) ToMonitorS3ConfigLogDeliveryStatusOutputWithContext(ctx context.Context) MonitorS3ConfigLogDeliveryStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(MonitorS3ConfigLogDeliveryStatusOutput)
}

func (e MonitorS3ConfigLogDeliveryStatus) ToMonitorS3ConfigLogDeliveryStatusPtrOutput() MonitorS3ConfigLogDeliveryStatusPtrOutput {
	return e.ToMonitorS3ConfigLogDeliveryStatusPtrOutputWithContext(context.Background())
}

func (e MonitorS3ConfigLogDeliveryStatus) ToMonitorS3ConfigLogDeliveryStatusPtrOutputWithContext(ctx context.Context) MonitorS3ConfigLogDeliveryStatusPtrOutput {
	return MonitorS3ConfigLogDeliveryStatus(e).ToMonitorS3ConfigLogDeliveryStatusOutputWithContext(ctx).ToMonitorS3ConfigLogDeliveryStatusPtrOutputWithContext(ctx)
}

func (e MonitorS3ConfigLogDeliveryStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e MonitorS3ConfigLogDeliveryStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e MonitorS3ConfigLogDeliveryStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e MonitorS3ConfigLogDeliveryStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type MonitorS3ConfigLogDeliveryStatusOutput struct{ *pulumi.OutputState }

func (MonitorS3ConfigLogDeliveryStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MonitorS3ConfigLogDeliveryStatus)(nil)).Elem()
}

func (o MonitorS3ConfigLogDeliveryStatusOutput) ToMonitorS3ConfigLogDeliveryStatusOutput() MonitorS3ConfigLogDeliveryStatusOutput {
	return o
}

func (o MonitorS3ConfigLogDeliveryStatusOutput) ToMonitorS3ConfigLogDeliveryStatusOutputWithContext(ctx context.Context) MonitorS3ConfigLogDeliveryStatusOutput {
	return o
}

func (o MonitorS3ConfigLogDeliveryStatusOutput) ToMonitorS3ConfigLogDeliveryStatusPtrOutput() MonitorS3ConfigLogDeliveryStatusPtrOutput {
	return o.ToMonitorS3ConfigLogDeliveryStatusPtrOutputWithContext(context.Background())
}

func (o MonitorS3ConfigLogDeliveryStatusOutput) ToMonitorS3ConfigLogDeliveryStatusPtrOutputWithContext(ctx context.Context) MonitorS3ConfigLogDeliveryStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MonitorS3ConfigLogDeliveryStatus) *MonitorS3ConfigLogDeliveryStatus {
		return &v
	}).(MonitorS3ConfigLogDeliveryStatusPtrOutput)
}

func (o MonitorS3ConfigLogDeliveryStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o MonitorS3ConfigLogDeliveryStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e MonitorS3ConfigLogDeliveryStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o MonitorS3ConfigLogDeliveryStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o MonitorS3ConfigLogDeliveryStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e MonitorS3ConfigLogDeliveryStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type MonitorS3ConfigLogDeliveryStatusPtrOutput struct{ *pulumi.OutputState }

func (MonitorS3ConfigLogDeliveryStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MonitorS3ConfigLogDeliveryStatus)(nil)).Elem()
}

func (o MonitorS3ConfigLogDeliveryStatusPtrOutput) ToMonitorS3ConfigLogDeliveryStatusPtrOutput() MonitorS3ConfigLogDeliveryStatusPtrOutput {
	return o
}

func (o MonitorS3ConfigLogDeliveryStatusPtrOutput) ToMonitorS3ConfigLogDeliveryStatusPtrOutputWithContext(ctx context.Context) MonitorS3ConfigLogDeliveryStatusPtrOutput {
	return o
}

func (o MonitorS3ConfigLogDeliveryStatusPtrOutput) Elem() MonitorS3ConfigLogDeliveryStatusOutput {
	return o.ApplyT(func(v *MonitorS3ConfigLogDeliveryStatus) MonitorS3ConfigLogDeliveryStatus {
		if v != nil {
			return *v
		}
		var ret MonitorS3ConfigLogDeliveryStatus
		return ret
	}).(MonitorS3ConfigLogDeliveryStatusOutput)
}

func (o MonitorS3ConfigLogDeliveryStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o MonitorS3ConfigLogDeliveryStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *MonitorS3ConfigLogDeliveryStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// MonitorS3ConfigLogDeliveryStatusInput is an input type that accepts MonitorS3ConfigLogDeliveryStatusArgs and MonitorS3ConfigLogDeliveryStatusOutput values.
// You can construct a concrete instance of `MonitorS3ConfigLogDeliveryStatusInput` via:
//
//	MonitorS3ConfigLogDeliveryStatusArgs{...}
type MonitorS3ConfigLogDeliveryStatusInput interface {
	pulumi.Input

	ToMonitorS3ConfigLogDeliveryStatusOutput() MonitorS3ConfigLogDeliveryStatusOutput
	ToMonitorS3ConfigLogDeliveryStatusOutputWithContext(context.Context) MonitorS3ConfigLogDeliveryStatusOutput
}

var monitorS3ConfigLogDeliveryStatusPtrType = reflect.TypeOf((**MonitorS3ConfigLogDeliveryStatus)(nil)).Elem()

type MonitorS3ConfigLogDeliveryStatusPtrInput interface {
	pulumi.Input

	ToMonitorS3ConfigLogDeliveryStatusPtrOutput() MonitorS3ConfigLogDeliveryStatusPtrOutput
	ToMonitorS3ConfigLogDeliveryStatusPtrOutputWithContext(context.Context) MonitorS3ConfigLogDeliveryStatusPtrOutput
}

type monitorS3ConfigLogDeliveryStatusPtr string

func MonitorS3ConfigLogDeliveryStatusPtr(v string) MonitorS3ConfigLogDeliveryStatusPtrInput {
	return (*monitorS3ConfigLogDeliveryStatusPtr)(&v)
}

func (*monitorS3ConfigLogDeliveryStatusPtr) ElementType() reflect.Type {
	return monitorS3ConfigLogDeliveryStatusPtrType
}

func (in *monitorS3ConfigLogDeliveryStatusPtr) ToMonitorS3ConfigLogDeliveryStatusPtrOutput() MonitorS3ConfigLogDeliveryStatusPtrOutput {
	return pulumi.ToOutput(in).(MonitorS3ConfigLogDeliveryStatusPtrOutput)
}

func (in *monitorS3ConfigLogDeliveryStatusPtr) ToMonitorS3ConfigLogDeliveryStatusPtrOutputWithContext(ctx context.Context) MonitorS3ConfigLogDeliveryStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(MonitorS3ConfigLogDeliveryStatusPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MonitorConfigStateInput)(nil)).Elem(), MonitorConfigState("PENDING"))
	pulumi.RegisterInputType(reflect.TypeOf((*MonitorConfigStatePtrInput)(nil)).Elem(), MonitorConfigState("PENDING"))
	pulumi.RegisterInputType(reflect.TypeOf((*MonitorLocalHealthEventsConfigStatusInput)(nil)).Elem(), MonitorLocalHealthEventsConfigStatus("ENABLED"))
	pulumi.RegisterInputType(reflect.TypeOf((*MonitorLocalHealthEventsConfigStatusPtrInput)(nil)).Elem(), MonitorLocalHealthEventsConfigStatus("ENABLED"))
	pulumi.RegisterInputType(reflect.TypeOf((*MonitorS3ConfigLogDeliveryStatusInput)(nil)).Elem(), MonitorS3ConfigLogDeliveryStatus("ENABLED"))
	pulumi.RegisterInputType(reflect.TypeOf((*MonitorS3ConfigLogDeliveryStatusPtrInput)(nil)).Elem(), MonitorS3ConfigLogDeliveryStatus("ENABLED"))
	pulumi.RegisterOutputType(MonitorConfigStateOutput{})
	pulumi.RegisterOutputType(MonitorConfigStatePtrOutput{})
	pulumi.RegisterOutputType(MonitorLocalHealthEventsConfigStatusOutput{})
	pulumi.RegisterOutputType(MonitorLocalHealthEventsConfigStatusPtrOutput{})
	pulumi.RegisterOutputType(MonitorProcessingStatusCodeOutput{})
	pulumi.RegisterOutputType(MonitorProcessingStatusCodePtrOutput{})
	pulumi.RegisterOutputType(MonitorS3ConfigLogDeliveryStatusOutput{})
	pulumi.RegisterOutputType(MonitorS3ConfigLogDeliveryStatusPtrOutput{})
}

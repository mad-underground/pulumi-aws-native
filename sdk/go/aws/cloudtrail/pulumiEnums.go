// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudtrail

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The type of destination for events arriving from a channel.
type ChannelDestinationType string

const (
	ChannelDestinationTypeEventDataStore = ChannelDestinationType("EVENT_DATA_STORE")
)

func (ChannelDestinationType) ElementType() reflect.Type {
	return reflect.TypeOf((*ChannelDestinationType)(nil)).Elem()
}

func (e ChannelDestinationType) ToChannelDestinationTypeOutput() ChannelDestinationTypeOutput {
	return pulumi.ToOutput(e).(ChannelDestinationTypeOutput)
}

func (e ChannelDestinationType) ToChannelDestinationTypeOutputWithContext(ctx context.Context) ChannelDestinationTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ChannelDestinationTypeOutput)
}

func (e ChannelDestinationType) ToChannelDestinationTypePtrOutput() ChannelDestinationTypePtrOutput {
	return e.ToChannelDestinationTypePtrOutputWithContext(context.Background())
}

func (e ChannelDestinationType) ToChannelDestinationTypePtrOutputWithContext(ctx context.Context) ChannelDestinationTypePtrOutput {
	return ChannelDestinationType(e).ToChannelDestinationTypeOutputWithContext(ctx).ToChannelDestinationTypePtrOutputWithContext(ctx)
}

func (e ChannelDestinationType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ChannelDestinationType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ChannelDestinationType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ChannelDestinationType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ChannelDestinationTypeOutput struct{ *pulumi.OutputState }

func (ChannelDestinationTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ChannelDestinationType)(nil)).Elem()
}

func (o ChannelDestinationTypeOutput) ToChannelDestinationTypeOutput() ChannelDestinationTypeOutput {
	return o
}

func (o ChannelDestinationTypeOutput) ToChannelDestinationTypeOutputWithContext(ctx context.Context) ChannelDestinationTypeOutput {
	return o
}

func (o ChannelDestinationTypeOutput) ToChannelDestinationTypePtrOutput() ChannelDestinationTypePtrOutput {
	return o.ToChannelDestinationTypePtrOutputWithContext(context.Background())
}

func (o ChannelDestinationTypeOutput) ToChannelDestinationTypePtrOutputWithContext(ctx context.Context) ChannelDestinationTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ChannelDestinationType) *ChannelDestinationType {
		return &v
	}).(ChannelDestinationTypePtrOutput)
}

func (o ChannelDestinationTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ChannelDestinationTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ChannelDestinationType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ChannelDestinationTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ChannelDestinationTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ChannelDestinationType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ChannelDestinationTypePtrOutput struct{ *pulumi.OutputState }

func (ChannelDestinationTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ChannelDestinationType)(nil)).Elem()
}

func (o ChannelDestinationTypePtrOutput) ToChannelDestinationTypePtrOutput() ChannelDestinationTypePtrOutput {
	return o
}

func (o ChannelDestinationTypePtrOutput) ToChannelDestinationTypePtrOutputWithContext(ctx context.Context) ChannelDestinationTypePtrOutput {
	return o
}

func (o ChannelDestinationTypePtrOutput) Elem() ChannelDestinationTypeOutput {
	return o.ApplyT(func(v *ChannelDestinationType) ChannelDestinationType {
		if v != nil {
			return *v
		}
		var ret ChannelDestinationType
		return ret
	}).(ChannelDestinationTypeOutput)
}

func (o ChannelDestinationTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ChannelDestinationTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ChannelDestinationType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ChannelDestinationTypeInput is an input type that accepts values of the ChannelDestinationType enum
// A concrete instance of `ChannelDestinationTypeInput` can be one of the following:
//
//	ChannelDestinationTypeEventDataStore
type ChannelDestinationTypeInput interface {
	pulumi.Input

	ToChannelDestinationTypeOutput() ChannelDestinationTypeOutput
	ToChannelDestinationTypeOutputWithContext(context.Context) ChannelDestinationTypeOutput
}

var channelDestinationTypePtrType = reflect.TypeOf((**ChannelDestinationType)(nil)).Elem()

type ChannelDestinationTypePtrInput interface {
	pulumi.Input

	ToChannelDestinationTypePtrOutput() ChannelDestinationTypePtrOutput
	ToChannelDestinationTypePtrOutputWithContext(context.Context) ChannelDestinationTypePtrOutput
}

type channelDestinationTypePtr string

func ChannelDestinationTypePtr(v string) ChannelDestinationTypePtrInput {
	return (*channelDestinationTypePtr)(&v)
}

func (*channelDestinationTypePtr) ElementType() reflect.Type {
	return channelDestinationTypePtrType
}

func (in *channelDestinationTypePtr) ToChannelDestinationTypePtrOutput() ChannelDestinationTypePtrOutput {
	return pulumi.ToOutput(in).(ChannelDestinationTypePtrOutput)
}

func (in *channelDestinationTypePtr) ToChannelDestinationTypePtrOutputWithContext(ctx context.Context) ChannelDestinationTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ChannelDestinationTypePtrOutput)
}

// Specify if you want your trail to log read-only events, write-only events, or all. For example, the EC2 GetConsoleOutput is a read-only API operation and RunInstances is a write-only API operation.
type TrailEventSelectorReadWriteType string

const (
	TrailEventSelectorReadWriteTypeAll       = TrailEventSelectorReadWriteType("All")
	TrailEventSelectorReadWriteTypeReadOnly  = TrailEventSelectorReadWriteType("ReadOnly")
	TrailEventSelectorReadWriteTypeWriteOnly = TrailEventSelectorReadWriteType("WriteOnly")
)

func (TrailEventSelectorReadWriteType) ElementType() reflect.Type {
	return reflect.TypeOf((*TrailEventSelectorReadWriteType)(nil)).Elem()
}

func (e TrailEventSelectorReadWriteType) ToTrailEventSelectorReadWriteTypeOutput() TrailEventSelectorReadWriteTypeOutput {
	return pulumi.ToOutput(e).(TrailEventSelectorReadWriteTypeOutput)
}

func (e TrailEventSelectorReadWriteType) ToTrailEventSelectorReadWriteTypeOutputWithContext(ctx context.Context) TrailEventSelectorReadWriteTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(TrailEventSelectorReadWriteTypeOutput)
}

func (e TrailEventSelectorReadWriteType) ToTrailEventSelectorReadWriteTypePtrOutput() TrailEventSelectorReadWriteTypePtrOutput {
	return e.ToTrailEventSelectorReadWriteTypePtrOutputWithContext(context.Background())
}

func (e TrailEventSelectorReadWriteType) ToTrailEventSelectorReadWriteTypePtrOutputWithContext(ctx context.Context) TrailEventSelectorReadWriteTypePtrOutput {
	return TrailEventSelectorReadWriteType(e).ToTrailEventSelectorReadWriteTypeOutputWithContext(ctx).ToTrailEventSelectorReadWriteTypePtrOutputWithContext(ctx)
}

func (e TrailEventSelectorReadWriteType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e TrailEventSelectorReadWriteType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e TrailEventSelectorReadWriteType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e TrailEventSelectorReadWriteType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type TrailEventSelectorReadWriteTypeOutput struct{ *pulumi.OutputState }

func (TrailEventSelectorReadWriteTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TrailEventSelectorReadWriteType)(nil)).Elem()
}

func (o TrailEventSelectorReadWriteTypeOutput) ToTrailEventSelectorReadWriteTypeOutput() TrailEventSelectorReadWriteTypeOutput {
	return o
}

func (o TrailEventSelectorReadWriteTypeOutput) ToTrailEventSelectorReadWriteTypeOutputWithContext(ctx context.Context) TrailEventSelectorReadWriteTypeOutput {
	return o
}

func (o TrailEventSelectorReadWriteTypeOutput) ToTrailEventSelectorReadWriteTypePtrOutput() TrailEventSelectorReadWriteTypePtrOutput {
	return o.ToTrailEventSelectorReadWriteTypePtrOutputWithContext(context.Background())
}

func (o TrailEventSelectorReadWriteTypeOutput) ToTrailEventSelectorReadWriteTypePtrOutputWithContext(ctx context.Context) TrailEventSelectorReadWriteTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TrailEventSelectorReadWriteType) *TrailEventSelectorReadWriteType {
		return &v
	}).(TrailEventSelectorReadWriteTypePtrOutput)
}

func (o TrailEventSelectorReadWriteTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o TrailEventSelectorReadWriteTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TrailEventSelectorReadWriteType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o TrailEventSelectorReadWriteTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TrailEventSelectorReadWriteTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TrailEventSelectorReadWriteType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type TrailEventSelectorReadWriteTypePtrOutput struct{ *pulumi.OutputState }

func (TrailEventSelectorReadWriteTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TrailEventSelectorReadWriteType)(nil)).Elem()
}

func (o TrailEventSelectorReadWriteTypePtrOutput) ToTrailEventSelectorReadWriteTypePtrOutput() TrailEventSelectorReadWriteTypePtrOutput {
	return o
}

func (o TrailEventSelectorReadWriteTypePtrOutput) ToTrailEventSelectorReadWriteTypePtrOutputWithContext(ctx context.Context) TrailEventSelectorReadWriteTypePtrOutput {
	return o
}

func (o TrailEventSelectorReadWriteTypePtrOutput) Elem() TrailEventSelectorReadWriteTypeOutput {
	return o.ApplyT(func(v *TrailEventSelectorReadWriteType) TrailEventSelectorReadWriteType {
		if v != nil {
			return *v
		}
		var ret TrailEventSelectorReadWriteType
		return ret
	}).(TrailEventSelectorReadWriteTypeOutput)
}

func (o TrailEventSelectorReadWriteTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TrailEventSelectorReadWriteTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *TrailEventSelectorReadWriteType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// TrailEventSelectorReadWriteTypeInput is an input type that accepts values of the TrailEventSelectorReadWriteType enum
// A concrete instance of `TrailEventSelectorReadWriteTypeInput` can be one of the following:
//
//	TrailEventSelectorReadWriteTypeAll
//	TrailEventSelectorReadWriteTypeReadOnly
//	TrailEventSelectorReadWriteTypeWriteOnly
type TrailEventSelectorReadWriteTypeInput interface {
	pulumi.Input

	ToTrailEventSelectorReadWriteTypeOutput() TrailEventSelectorReadWriteTypeOutput
	ToTrailEventSelectorReadWriteTypeOutputWithContext(context.Context) TrailEventSelectorReadWriteTypeOutput
}

var trailEventSelectorReadWriteTypePtrType = reflect.TypeOf((**TrailEventSelectorReadWriteType)(nil)).Elem()

type TrailEventSelectorReadWriteTypePtrInput interface {
	pulumi.Input

	ToTrailEventSelectorReadWriteTypePtrOutput() TrailEventSelectorReadWriteTypePtrOutput
	ToTrailEventSelectorReadWriteTypePtrOutputWithContext(context.Context) TrailEventSelectorReadWriteTypePtrOutput
}

type trailEventSelectorReadWriteTypePtr string

func TrailEventSelectorReadWriteTypePtr(v string) TrailEventSelectorReadWriteTypePtrInput {
	return (*trailEventSelectorReadWriteTypePtr)(&v)
}

func (*trailEventSelectorReadWriteTypePtr) ElementType() reflect.Type {
	return trailEventSelectorReadWriteTypePtrType
}

func (in *trailEventSelectorReadWriteTypePtr) ToTrailEventSelectorReadWriteTypePtrOutput() TrailEventSelectorReadWriteTypePtrOutput {
	return pulumi.ToOutput(in).(TrailEventSelectorReadWriteTypePtrOutput)
}

func (in *trailEventSelectorReadWriteTypePtr) ToTrailEventSelectorReadWriteTypePtrOutputWithContext(ctx context.Context) TrailEventSelectorReadWriteTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(TrailEventSelectorReadWriteTypePtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ChannelDestinationTypeInput)(nil)).Elem(), ChannelDestinationType("EVENT_DATA_STORE"))
	pulumi.RegisterInputType(reflect.TypeOf((*ChannelDestinationTypePtrInput)(nil)).Elem(), ChannelDestinationType("EVENT_DATA_STORE"))
	pulumi.RegisterInputType(reflect.TypeOf((*TrailEventSelectorReadWriteTypeInput)(nil)).Elem(), TrailEventSelectorReadWriteType("All"))
	pulumi.RegisterInputType(reflect.TypeOf((*TrailEventSelectorReadWriteTypePtrInput)(nil)).Elem(), TrailEventSelectorReadWriteType("All"))
	pulumi.RegisterOutputType(ChannelDestinationTypeOutput{})
	pulumi.RegisterOutputType(ChannelDestinationTypePtrOutput{})
	pulumi.RegisterOutputType(TrailEventSelectorReadWriteTypeOutput{})
	pulumi.RegisterOutputType(TrailEventSelectorReadWriteTypePtrOutput{})
}

// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dynamodb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecision string

const (
	GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionMicrosecond = GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecision("MICROSECOND")
	GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionMillisecond = GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecision("MILLISECOND")
)

func (GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecision) ElementType() reflect.Type {
	return reflect.TypeOf((*GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecision)(nil)).Elem()
}

func (e GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecision) ToGlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionOutput() GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionOutput {
	return pulumi.ToOutput(e).(GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionOutput)
}

func (e GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecision) ToGlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionOutputWithContext(ctx context.Context) GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionOutput)
}

func (e GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecision) ToGlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtrOutput() GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtrOutput {
	return e.ToGlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtrOutputWithContext(context.Background())
}

func (e GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecision) ToGlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtrOutputWithContext(ctx context.Context) GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtrOutput {
	return GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecision(e).ToGlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionOutputWithContext(ctx).ToGlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtrOutputWithContext(ctx)
}

func (e GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecision) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecision) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecision) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecision) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionOutput struct{ *pulumi.OutputState }

func (GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecision)(nil)).Elem()
}

func (o GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionOutput) ToGlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionOutput() GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionOutput {
	return o
}

func (o GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionOutput) ToGlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionOutputWithContext(ctx context.Context) GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionOutput {
	return o
}

func (o GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionOutput) ToGlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtrOutput() GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtrOutput {
	return o.ToGlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtrOutputWithContext(context.Background())
}

func (o GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionOutput) ToGlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtrOutputWithContext(ctx context.Context) GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecision) *GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecision {
		return &v
	}).(GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtrOutput)
}

func (o GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecision) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecision) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtrOutput struct{ *pulumi.OutputState }

func (GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecision)(nil)).Elem()
}

func (o GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtrOutput) ToGlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtrOutput() GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtrOutput {
	return o
}

func (o GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtrOutput) ToGlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtrOutputWithContext(ctx context.Context) GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtrOutput {
	return o
}

func (o GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtrOutput) Elem() GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionOutput {
	return o.ApplyT(func(v *GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecision) GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecision {
		if v != nil {
			return *v
		}
		var ret GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecision
		return ret
	}).(GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionOutput)
}

func (o GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecision) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionInput is an input type that accepts values of the GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecision enum
// A concrete instance of `GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionInput` can be one of the following:
//
//	GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionMicrosecond
//	GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionMillisecond
type GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionInput interface {
	pulumi.Input

	ToGlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionOutput() GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionOutput
	ToGlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionOutputWithContext(context.Context) GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionOutput
}

var globalTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtrType = reflect.TypeOf((**GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecision)(nil)).Elem()

type GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtrInput interface {
	pulumi.Input

	ToGlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtrOutput() GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtrOutput
	ToGlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtrOutputWithContext(context.Context) GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtrOutput
}

type globalTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtr string

func GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtr(v string) GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtrInput {
	return (*globalTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtr)(&v)
}

func (*globalTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtr) ElementType() reflect.Type {
	return globalTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtrType
}

func (in *globalTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtr) ToGlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtrOutput() GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtrOutput {
	return pulumi.ToOutput(in).(GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtrOutput)
}

func (in *globalTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtr) ToGlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtrOutputWithContext(ctx context.Context) GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtrOutput)
}

type TableKinesisStreamSpecificationApproximateCreationDateTimePrecision string

const (
	TableKinesisStreamSpecificationApproximateCreationDateTimePrecisionMicrosecond = TableKinesisStreamSpecificationApproximateCreationDateTimePrecision("MICROSECOND")
	TableKinesisStreamSpecificationApproximateCreationDateTimePrecisionMillisecond = TableKinesisStreamSpecificationApproximateCreationDateTimePrecision("MILLISECOND")
)

func (TableKinesisStreamSpecificationApproximateCreationDateTimePrecision) ElementType() reflect.Type {
	return reflect.TypeOf((*TableKinesisStreamSpecificationApproximateCreationDateTimePrecision)(nil)).Elem()
}

func (e TableKinesisStreamSpecificationApproximateCreationDateTimePrecision) ToTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionOutput() TableKinesisStreamSpecificationApproximateCreationDateTimePrecisionOutput {
	return pulumi.ToOutput(e).(TableKinesisStreamSpecificationApproximateCreationDateTimePrecisionOutput)
}

func (e TableKinesisStreamSpecificationApproximateCreationDateTimePrecision) ToTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionOutputWithContext(ctx context.Context) TableKinesisStreamSpecificationApproximateCreationDateTimePrecisionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(TableKinesisStreamSpecificationApproximateCreationDateTimePrecisionOutput)
}

func (e TableKinesisStreamSpecificationApproximateCreationDateTimePrecision) ToTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtrOutput() TableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtrOutput {
	return e.ToTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtrOutputWithContext(context.Background())
}

func (e TableKinesisStreamSpecificationApproximateCreationDateTimePrecision) ToTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtrOutputWithContext(ctx context.Context) TableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtrOutput {
	return TableKinesisStreamSpecificationApproximateCreationDateTimePrecision(e).ToTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionOutputWithContext(ctx).ToTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtrOutputWithContext(ctx)
}

func (e TableKinesisStreamSpecificationApproximateCreationDateTimePrecision) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e TableKinesisStreamSpecificationApproximateCreationDateTimePrecision) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e TableKinesisStreamSpecificationApproximateCreationDateTimePrecision) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e TableKinesisStreamSpecificationApproximateCreationDateTimePrecision) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type TableKinesisStreamSpecificationApproximateCreationDateTimePrecisionOutput struct{ *pulumi.OutputState }

func (TableKinesisStreamSpecificationApproximateCreationDateTimePrecisionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TableKinesisStreamSpecificationApproximateCreationDateTimePrecision)(nil)).Elem()
}

func (o TableKinesisStreamSpecificationApproximateCreationDateTimePrecisionOutput) ToTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionOutput() TableKinesisStreamSpecificationApproximateCreationDateTimePrecisionOutput {
	return o
}

func (o TableKinesisStreamSpecificationApproximateCreationDateTimePrecisionOutput) ToTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionOutputWithContext(ctx context.Context) TableKinesisStreamSpecificationApproximateCreationDateTimePrecisionOutput {
	return o
}

func (o TableKinesisStreamSpecificationApproximateCreationDateTimePrecisionOutput) ToTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtrOutput() TableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtrOutput {
	return o.ToTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtrOutputWithContext(context.Background())
}

func (o TableKinesisStreamSpecificationApproximateCreationDateTimePrecisionOutput) ToTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtrOutputWithContext(ctx context.Context) TableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TableKinesisStreamSpecificationApproximateCreationDateTimePrecision) *TableKinesisStreamSpecificationApproximateCreationDateTimePrecision {
		return &v
	}).(TableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtrOutput)
}

func (o TableKinesisStreamSpecificationApproximateCreationDateTimePrecisionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o TableKinesisStreamSpecificationApproximateCreationDateTimePrecisionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TableKinesisStreamSpecificationApproximateCreationDateTimePrecision) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o TableKinesisStreamSpecificationApproximateCreationDateTimePrecisionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TableKinesisStreamSpecificationApproximateCreationDateTimePrecisionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TableKinesisStreamSpecificationApproximateCreationDateTimePrecision) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type TableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtrOutput struct{ *pulumi.OutputState }

func (TableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TableKinesisStreamSpecificationApproximateCreationDateTimePrecision)(nil)).Elem()
}

func (o TableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtrOutput) ToTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtrOutput() TableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtrOutput {
	return o
}

func (o TableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtrOutput) ToTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtrOutputWithContext(ctx context.Context) TableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtrOutput {
	return o
}

func (o TableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtrOutput) Elem() TableKinesisStreamSpecificationApproximateCreationDateTimePrecisionOutput {
	return o.ApplyT(func(v *TableKinesisStreamSpecificationApproximateCreationDateTimePrecision) TableKinesisStreamSpecificationApproximateCreationDateTimePrecision {
		if v != nil {
			return *v
		}
		var ret TableKinesisStreamSpecificationApproximateCreationDateTimePrecision
		return ret
	}).(TableKinesisStreamSpecificationApproximateCreationDateTimePrecisionOutput)
}

func (o TableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *TableKinesisStreamSpecificationApproximateCreationDateTimePrecision) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// TableKinesisStreamSpecificationApproximateCreationDateTimePrecisionInput is an input type that accepts values of the TableKinesisStreamSpecificationApproximateCreationDateTimePrecision enum
// A concrete instance of `TableKinesisStreamSpecificationApproximateCreationDateTimePrecisionInput` can be one of the following:
//
//	TableKinesisStreamSpecificationApproximateCreationDateTimePrecisionMicrosecond
//	TableKinesisStreamSpecificationApproximateCreationDateTimePrecisionMillisecond
type TableKinesisStreamSpecificationApproximateCreationDateTimePrecisionInput interface {
	pulumi.Input

	ToTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionOutput() TableKinesisStreamSpecificationApproximateCreationDateTimePrecisionOutput
	ToTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionOutputWithContext(context.Context) TableKinesisStreamSpecificationApproximateCreationDateTimePrecisionOutput
}

var tableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtrType = reflect.TypeOf((**TableKinesisStreamSpecificationApproximateCreationDateTimePrecision)(nil)).Elem()

type TableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtrInput interface {
	pulumi.Input

	ToTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtrOutput() TableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtrOutput
	ToTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtrOutputWithContext(context.Context) TableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtrOutput
}

type tableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtr string

func TableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtr(v string) TableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtrInput {
	return (*tableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtr)(&v)
}

func (*tableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtr) ElementType() reflect.Type {
	return tableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtrType
}

func (in *tableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtr) ToTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtrOutput() TableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtrOutput {
	return pulumi.ToOutput(in).(TableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtrOutput)
}

func (in *tableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtr) ToTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtrOutputWithContext(ctx context.Context) TableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(TableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionInput)(nil)).Elem(), GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecision("MICROSECOND"))
	pulumi.RegisterInputType(reflect.TypeOf((*GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtrInput)(nil)).Elem(), GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecision("MICROSECOND"))
	pulumi.RegisterInputType(reflect.TypeOf((*TableKinesisStreamSpecificationApproximateCreationDateTimePrecisionInput)(nil)).Elem(), TableKinesisStreamSpecificationApproximateCreationDateTimePrecision("MICROSECOND"))
	pulumi.RegisterInputType(reflect.TypeOf((*TableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtrInput)(nil)).Elem(), TableKinesisStreamSpecificationApproximateCreationDateTimePrecision("MICROSECOND"))
	pulumi.RegisterOutputType(GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionOutput{})
	pulumi.RegisterOutputType(GlobalTableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtrOutput{})
	pulumi.RegisterOutputType(TableKinesisStreamSpecificationApproximateCreationDateTimePrecisionOutput{})
	pulumi.RegisterOutputType(TableKinesisStreamSpecificationApproximateCreationDateTimePrecisionPtrOutput{})
}

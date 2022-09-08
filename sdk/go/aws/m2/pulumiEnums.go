// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package m2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The target platform for the environment.
type EnvironmentEngineType string

const (
	EnvironmentEngineTypeMicrofocus = EnvironmentEngineType("microfocus")
	EnvironmentEngineTypeBluage     = EnvironmentEngineType("bluage")
)

func (EnvironmentEngineType) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentEngineType)(nil)).Elem()
}

func (e EnvironmentEngineType) ToEnvironmentEngineTypeOutput() EnvironmentEngineTypeOutput {
	return pulumi.ToOutput(e).(EnvironmentEngineTypeOutput)
}

func (e EnvironmentEngineType) ToEnvironmentEngineTypeOutputWithContext(ctx context.Context) EnvironmentEngineTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(EnvironmentEngineTypeOutput)
}

func (e EnvironmentEngineType) ToEnvironmentEngineTypePtrOutput() EnvironmentEngineTypePtrOutput {
	return e.ToEnvironmentEngineTypePtrOutputWithContext(context.Background())
}

func (e EnvironmentEngineType) ToEnvironmentEngineTypePtrOutputWithContext(ctx context.Context) EnvironmentEngineTypePtrOutput {
	return EnvironmentEngineType(e).ToEnvironmentEngineTypeOutputWithContext(ctx).ToEnvironmentEngineTypePtrOutputWithContext(ctx)
}

func (e EnvironmentEngineType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e EnvironmentEngineType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e EnvironmentEngineType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e EnvironmentEngineType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type EnvironmentEngineTypeOutput struct{ *pulumi.OutputState }

func (EnvironmentEngineTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentEngineType)(nil)).Elem()
}

func (o EnvironmentEngineTypeOutput) ToEnvironmentEngineTypeOutput() EnvironmentEngineTypeOutput {
	return o
}

func (o EnvironmentEngineTypeOutput) ToEnvironmentEngineTypeOutputWithContext(ctx context.Context) EnvironmentEngineTypeOutput {
	return o
}

func (o EnvironmentEngineTypeOutput) ToEnvironmentEngineTypePtrOutput() EnvironmentEngineTypePtrOutput {
	return o.ToEnvironmentEngineTypePtrOutputWithContext(context.Background())
}

func (o EnvironmentEngineTypeOutput) ToEnvironmentEngineTypePtrOutputWithContext(ctx context.Context) EnvironmentEngineTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EnvironmentEngineType) *EnvironmentEngineType {
		return &v
	}).(EnvironmentEngineTypePtrOutput)
}

func (o EnvironmentEngineTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o EnvironmentEngineTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EnvironmentEngineType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o EnvironmentEngineTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EnvironmentEngineTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EnvironmentEngineType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type EnvironmentEngineTypePtrOutput struct{ *pulumi.OutputState }

func (EnvironmentEngineTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnvironmentEngineType)(nil)).Elem()
}

func (o EnvironmentEngineTypePtrOutput) ToEnvironmentEngineTypePtrOutput() EnvironmentEngineTypePtrOutput {
	return o
}

func (o EnvironmentEngineTypePtrOutput) ToEnvironmentEngineTypePtrOutputWithContext(ctx context.Context) EnvironmentEngineTypePtrOutput {
	return o
}

func (o EnvironmentEngineTypePtrOutput) Elem() EnvironmentEngineTypeOutput {
	return o.ApplyT(func(v *EnvironmentEngineType) EnvironmentEngineType {
		if v != nil {
			return *v
		}
		var ret EnvironmentEngineType
		return ret
	}).(EnvironmentEngineTypeOutput)
}

func (o EnvironmentEngineTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EnvironmentEngineTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *EnvironmentEngineType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// EnvironmentEngineTypeInput is an input type that accepts EnvironmentEngineTypeArgs and EnvironmentEngineTypeOutput values.
// You can construct a concrete instance of `EnvironmentEngineTypeInput` via:
//
//	EnvironmentEngineTypeArgs{...}
type EnvironmentEngineTypeInput interface {
	pulumi.Input

	ToEnvironmentEngineTypeOutput() EnvironmentEngineTypeOutput
	ToEnvironmentEngineTypeOutputWithContext(context.Context) EnvironmentEngineTypeOutput
}

var environmentEngineTypePtrType = reflect.TypeOf((**EnvironmentEngineType)(nil)).Elem()

type EnvironmentEngineTypePtrInput interface {
	pulumi.Input

	ToEnvironmentEngineTypePtrOutput() EnvironmentEngineTypePtrOutput
	ToEnvironmentEngineTypePtrOutputWithContext(context.Context) EnvironmentEngineTypePtrOutput
}

type environmentEngineTypePtr string

func EnvironmentEngineTypePtr(v string) EnvironmentEngineTypePtrInput {
	return (*environmentEngineTypePtr)(&v)
}

func (*environmentEngineTypePtr) ElementType() reflect.Type {
	return environmentEngineTypePtrType
}

func (in *environmentEngineTypePtr) ToEnvironmentEngineTypePtrOutput() EnvironmentEngineTypePtrOutput {
	return pulumi.ToOutput(in).(EnvironmentEngineTypePtrOutput)
}

func (in *environmentEngineTypePtr) ToEnvironmentEngineTypePtrOutputWithContext(ctx context.Context) EnvironmentEngineTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(EnvironmentEngineTypePtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EnvironmentEngineTypeInput)(nil)).Elem(), EnvironmentEngineType("microfocus"))
	pulumi.RegisterInputType(reflect.TypeOf((*EnvironmentEngineTypePtrInput)(nil)).Elem(), EnvironmentEngineType("microfocus"))
	pulumi.RegisterOutputType(EnvironmentEngineTypeOutput{})
	pulumi.RegisterOutputType(EnvironmentEngineTypePtrOutput{})
}

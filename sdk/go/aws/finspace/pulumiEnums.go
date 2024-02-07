// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package finspace

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Federation mode used with the Environment
type EnvironmentFederationMode string

const (
	EnvironmentFederationModeLocal     = EnvironmentFederationMode("LOCAL")
	EnvironmentFederationModeFederated = EnvironmentFederationMode("FEDERATED")
)

func (EnvironmentFederationMode) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentFederationMode)(nil)).Elem()
}

func (e EnvironmentFederationMode) ToEnvironmentFederationModeOutput() EnvironmentFederationModeOutput {
	return pulumi.ToOutput(e).(EnvironmentFederationModeOutput)
}

func (e EnvironmentFederationMode) ToEnvironmentFederationModeOutputWithContext(ctx context.Context) EnvironmentFederationModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(EnvironmentFederationModeOutput)
}

func (e EnvironmentFederationMode) ToEnvironmentFederationModePtrOutput() EnvironmentFederationModePtrOutput {
	return e.ToEnvironmentFederationModePtrOutputWithContext(context.Background())
}

func (e EnvironmentFederationMode) ToEnvironmentFederationModePtrOutputWithContext(ctx context.Context) EnvironmentFederationModePtrOutput {
	return EnvironmentFederationMode(e).ToEnvironmentFederationModeOutputWithContext(ctx).ToEnvironmentFederationModePtrOutputWithContext(ctx)
}

func (e EnvironmentFederationMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e EnvironmentFederationMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e EnvironmentFederationMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e EnvironmentFederationMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type EnvironmentFederationModeOutput struct{ *pulumi.OutputState }

func (EnvironmentFederationModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentFederationMode)(nil)).Elem()
}

func (o EnvironmentFederationModeOutput) ToEnvironmentFederationModeOutput() EnvironmentFederationModeOutput {
	return o
}

func (o EnvironmentFederationModeOutput) ToEnvironmentFederationModeOutputWithContext(ctx context.Context) EnvironmentFederationModeOutput {
	return o
}

func (o EnvironmentFederationModeOutput) ToEnvironmentFederationModePtrOutput() EnvironmentFederationModePtrOutput {
	return o.ToEnvironmentFederationModePtrOutputWithContext(context.Background())
}

func (o EnvironmentFederationModeOutput) ToEnvironmentFederationModePtrOutputWithContext(ctx context.Context) EnvironmentFederationModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EnvironmentFederationMode) *EnvironmentFederationMode {
		return &v
	}).(EnvironmentFederationModePtrOutput)
}

func (o EnvironmentFederationModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o EnvironmentFederationModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EnvironmentFederationMode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o EnvironmentFederationModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EnvironmentFederationModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EnvironmentFederationMode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type EnvironmentFederationModePtrOutput struct{ *pulumi.OutputState }

func (EnvironmentFederationModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnvironmentFederationMode)(nil)).Elem()
}

func (o EnvironmentFederationModePtrOutput) ToEnvironmentFederationModePtrOutput() EnvironmentFederationModePtrOutput {
	return o
}

func (o EnvironmentFederationModePtrOutput) ToEnvironmentFederationModePtrOutputWithContext(ctx context.Context) EnvironmentFederationModePtrOutput {
	return o
}

func (o EnvironmentFederationModePtrOutput) Elem() EnvironmentFederationModeOutput {
	return o.ApplyT(func(v *EnvironmentFederationMode) EnvironmentFederationMode {
		if v != nil {
			return *v
		}
		var ret EnvironmentFederationMode
		return ret
	}).(EnvironmentFederationModeOutput)
}

func (o EnvironmentFederationModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EnvironmentFederationModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *EnvironmentFederationMode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// EnvironmentFederationModeInput is an input type that accepts values of the EnvironmentFederationMode enum
// A concrete instance of `EnvironmentFederationModeInput` can be one of the following:
//
//	EnvironmentFederationModeLocal
//	EnvironmentFederationModeFederated
type EnvironmentFederationModeInput interface {
	pulumi.Input

	ToEnvironmentFederationModeOutput() EnvironmentFederationModeOutput
	ToEnvironmentFederationModeOutputWithContext(context.Context) EnvironmentFederationModeOutput
}

var environmentFederationModePtrType = reflect.TypeOf((**EnvironmentFederationMode)(nil)).Elem()

type EnvironmentFederationModePtrInput interface {
	pulumi.Input

	ToEnvironmentFederationModePtrOutput() EnvironmentFederationModePtrOutput
	ToEnvironmentFederationModePtrOutputWithContext(context.Context) EnvironmentFederationModePtrOutput
}

type environmentFederationModePtr string

func EnvironmentFederationModePtr(v string) EnvironmentFederationModePtrInput {
	return (*environmentFederationModePtr)(&v)
}

func (*environmentFederationModePtr) ElementType() reflect.Type {
	return environmentFederationModePtrType
}

func (in *environmentFederationModePtr) ToEnvironmentFederationModePtrOutput() EnvironmentFederationModePtrOutput {
	return pulumi.ToOutput(in).(EnvironmentFederationModePtrOutput)
}

func (in *environmentFederationModePtr) ToEnvironmentFederationModePtrOutputWithContext(ctx context.Context) EnvironmentFederationModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(EnvironmentFederationModePtrOutput)
}

// State of the Environment
type EnvironmentStatus string

const (
	EnvironmentStatusCreateRequested = EnvironmentStatus("CREATE_REQUESTED")
	EnvironmentStatusCreating        = EnvironmentStatus("CREATING")
	EnvironmentStatusCreated         = EnvironmentStatus("CREATED")
	EnvironmentStatusDeleteRequested = EnvironmentStatus("DELETE_REQUESTED")
	EnvironmentStatusDeleting        = EnvironmentStatus("DELETING")
	EnvironmentStatusDeleted         = EnvironmentStatus("DELETED")
	EnvironmentStatusFailedCreation  = EnvironmentStatus("FAILED_CREATION")
	EnvironmentStatusFailedDeletion  = EnvironmentStatus("FAILED_DELETION")
	EnvironmentStatusRetryDeletion   = EnvironmentStatus("RETRY_DELETION")
	EnvironmentStatusSuspended       = EnvironmentStatus("SUSPENDED")
)

type EnvironmentStatusOutput struct{ *pulumi.OutputState }

func (EnvironmentStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentStatus)(nil)).Elem()
}

func (o EnvironmentStatusOutput) ToEnvironmentStatusOutput() EnvironmentStatusOutput {
	return o
}

func (o EnvironmentStatusOutput) ToEnvironmentStatusOutputWithContext(ctx context.Context) EnvironmentStatusOutput {
	return o
}

func (o EnvironmentStatusOutput) ToEnvironmentStatusPtrOutput() EnvironmentStatusPtrOutput {
	return o.ToEnvironmentStatusPtrOutputWithContext(context.Background())
}

func (o EnvironmentStatusOutput) ToEnvironmentStatusPtrOutputWithContext(ctx context.Context) EnvironmentStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EnvironmentStatus) *EnvironmentStatus {
		return &v
	}).(EnvironmentStatusPtrOutput)
}

func (o EnvironmentStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o EnvironmentStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EnvironmentStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o EnvironmentStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EnvironmentStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EnvironmentStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type EnvironmentStatusPtrOutput struct{ *pulumi.OutputState }

func (EnvironmentStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnvironmentStatus)(nil)).Elem()
}

func (o EnvironmentStatusPtrOutput) ToEnvironmentStatusPtrOutput() EnvironmentStatusPtrOutput {
	return o
}

func (o EnvironmentStatusPtrOutput) ToEnvironmentStatusPtrOutputWithContext(ctx context.Context) EnvironmentStatusPtrOutput {
	return o
}

func (o EnvironmentStatusPtrOutput) Elem() EnvironmentStatusOutput {
	return o.ApplyT(func(v *EnvironmentStatus) EnvironmentStatus {
		if v != nil {
			return *v
		}
		var ret EnvironmentStatus
		return ret
	}).(EnvironmentStatusOutput)
}

func (o EnvironmentStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EnvironmentStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *EnvironmentStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EnvironmentFederationModeInput)(nil)).Elem(), EnvironmentFederationMode("LOCAL"))
	pulumi.RegisterInputType(reflect.TypeOf((*EnvironmentFederationModePtrInput)(nil)).Elem(), EnvironmentFederationMode("LOCAL"))
	pulumi.RegisterOutputType(EnvironmentFederationModeOutput{})
	pulumi.RegisterOutputType(EnvironmentFederationModePtrOutput{})
	pulumi.RegisterOutputType(EnvironmentStatusOutput{})
	pulumi.RegisterOutputType(EnvironmentStatusPtrOutput{})
}

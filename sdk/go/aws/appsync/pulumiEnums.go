// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package appsync

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configuration of the merged behavior for the association. For example when it could be auto or has to be manual.
type SourceApiAssociationConfigMergeType string

const (
	SourceApiAssociationConfigMergeTypeAutoMerge   = SourceApiAssociationConfigMergeType("AUTO_MERGE")
	SourceApiAssociationConfigMergeTypeManualMerge = SourceApiAssociationConfigMergeType("MANUAL_MERGE")
)

func (SourceApiAssociationConfigMergeType) ElementType() reflect.Type {
	return reflect.TypeOf((*SourceApiAssociationConfigMergeType)(nil)).Elem()
}

func (e SourceApiAssociationConfigMergeType) ToSourceApiAssociationConfigMergeTypeOutput() SourceApiAssociationConfigMergeTypeOutput {
	return pulumi.ToOutput(e).(SourceApiAssociationConfigMergeTypeOutput)
}

func (e SourceApiAssociationConfigMergeType) ToSourceApiAssociationConfigMergeTypeOutputWithContext(ctx context.Context) SourceApiAssociationConfigMergeTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SourceApiAssociationConfigMergeTypeOutput)
}

func (e SourceApiAssociationConfigMergeType) ToSourceApiAssociationConfigMergeTypePtrOutput() SourceApiAssociationConfigMergeTypePtrOutput {
	return e.ToSourceApiAssociationConfigMergeTypePtrOutputWithContext(context.Background())
}

func (e SourceApiAssociationConfigMergeType) ToSourceApiAssociationConfigMergeTypePtrOutputWithContext(ctx context.Context) SourceApiAssociationConfigMergeTypePtrOutput {
	return SourceApiAssociationConfigMergeType(e).ToSourceApiAssociationConfigMergeTypeOutputWithContext(ctx).ToSourceApiAssociationConfigMergeTypePtrOutputWithContext(ctx)
}

func (e SourceApiAssociationConfigMergeType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SourceApiAssociationConfigMergeType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SourceApiAssociationConfigMergeType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SourceApiAssociationConfigMergeType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SourceApiAssociationConfigMergeTypeOutput struct{ *pulumi.OutputState }

func (SourceApiAssociationConfigMergeTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SourceApiAssociationConfigMergeType)(nil)).Elem()
}

func (o SourceApiAssociationConfigMergeTypeOutput) ToSourceApiAssociationConfigMergeTypeOutput() SourceApiAssociationConfigMergeTypeOutput {
	return o
}

func (o SourceApiAssociationConfigMergeTypeOutput) ToSourceApiAssociationConfigMergeTypeOutputWithContext(ctx context.Context) SourceApiAssociationConfigMergeTypeOutput {
	return o
}

func (o SourceApiAssociationConfigMergeTypeOutput) ToSourceApiAssociationConfigMergeTypePtrOutput() SourceApiAssociationConfigMergeTypePtrOutput {
	return o.ToSourceApiAssociationConfigMergeTypePtrOutputWithContext(context.Background())
}

func (o SourceApiAssociationConfigMergeTypeOutput) ToSourceApiAssociationConfigMergeTypePtrOutputWithContext(ctx context.Context) SourceApiAssociationConfigMergeTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SourceApiAssociationConfigMergeType) *SourceApiAssociationConfigMergeType {
		return &v
	}).(SourceApiAssociationConfigMergeTypePtrOutput)
}

func (o SourceApiAssociationConfigMergeTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SourceApiAssociationConfigMergeTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SourceApiAssociationConfigMergeType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SourceApiAssociationConfigMergeTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SourceApiAssociationConfigMergeTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SourceApiAssociationConfigMergeType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SourceApiAssociationConfigMergeTypePtrOutput struct{ *pulumi.OutputState }

func (SourceApiAssociationConfigMergeTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceApiAssociationConfigMergeType)(nil)).Elem()
}

func (o SourceApiAssociationConfigMergeTypePtrOutput) ToSourceApiAssociationConfigMergeTypePtrOutput() SourceApiAssociationConfigMergeTypePtrOutput {
	return o
}

func (o SourceApiAssociationConfigMergeTypePtrOutput) ToSourceApiAssociationConfigMergeTypePtrOutputWithContext(ctx context.Context) SourceApiAssociationConfigMergeTypePtrOutput {
	return o
}

func (o SourceApiAssociationConfigMergeTypePtrOutput) Elem() SourceApiAssociationConfigMergeTypeOutput {
	return o.ApplyT(func(v *SourceApiAssociationConfigMergeType) SourceApiAssociationConfigMergeType {
		if v != nil {
			return *v
		}
		var ret SourceApiAssociationConfigMergeType
		return ret
	}).(SourceApiAssociationConfigMergeTypeOutput)
}

func (o SourceApiAssociationConfigMergeTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SourceApiAssociationConfigMergeTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SourceApiAssociationConfigMergeType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// SourceApiAssociationConfigMergeTypeInput is an input type that accepts values of the SourceApiAssociationConfigMergeType enum
// A concrete instance of `SourceApiAssociationConfigMergeTypeInput` can be one of the following:
//
//	SourceApiAssociationConfigMergeTypeAutoMerge
//	SourceApiAssociationConfigMergeTypeManualMerge
type SourceApiAssociationConfigMergeTypeInput interface {
	pulumi.Input

	ToSourceApiAssociationConfigMergeTypeOutput() SourceApiAssociationConfigMergeTypeOutput
	ToSourceApiAssociationConfigMergeTypeOutputWithContext(context.Context) SourceApiAssociationConfigMergeTypeOutput
}

var sourceApiAssociationConfigMergeTypePtrType = reflect.TypeOf((**SourceApiAssociationConfigMergeType)(nil)).Elem()

type SourceApiAssociationConfigMergeTypePtrInput interface {
	pulumi.Input

	ToSourceApiAssociationConfigMergeTypePtrOutput() SourceApiAssociationConfigMergeTypePtrOutput
	ToSourceApiAssociationConfigMergeTypePtrOutputWithContext(context.Context) SourceApiAssociationConfigMergeTypePtrOutput
}

type sourceApiAssociationConfigMergeTypePtr string

func SourceApiAssociationConfigMergeTypePtr(v string) SourceApiAssociationConfigMergeTypePtrInput {
	return (*sourceApiAssociationConfigMergeTypePtr)(&v)
}

func (*sourceApiAssociationConfigMergeTypePtr) ElementType() reflect.Type {
	return sourceApiAssociationConfigMergeTypePtrType
}

func (in *sourceApiAssociationConfigMergeTypePtr) ToSourceApiAssociationConfigMergeTypePtrOutput() SourceApiAssociationConfigMergeTypePtrOutput {
	return pulumi.ToOutput(in).(SourceApiAssociationConfigMergeTypePtrOutput)
}

func (in *sourceApiAssociationConfigMergeTypePtr) ToSourceApiAssociationConfigMergeTypePtrOutputWithContext(ctx context.Context) SourceApiAssociationConfigMergeTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SourceApiAssociationConfigMergeTypePtrOutput)
}

// Current status of SourceApiAssociation.
type SourceApiAssociationStatus string

const (
	SourceApiAssociationStatusMergeScheduled          = SourceApiAssociationStatus("MERGE_SCHEDULED")
	SourceApiAssociationStatusMergeFailed             = SourceApiAssociationStatus("MERGE_FAILED")
	SourceApiAssociationStatusMergeSuccess            = SourceApiAssociationStatus("MERGE_SUCCESS")
	SourceApiAssociationStatusMergeInProgress         = SourceApiAssociationStatus("MERGE_IN_PROGRESS")
	SourceApiAssociationStatusAutoMergeScheduleFailed = SourceApiAssociationStatus("AUTO_MERGE_SCHEDULE_FAILED")
	SourceApiAssociationStatusDeletionScheduled       = SourceApiAssociationStatus("DELETION_SCHEDULED")
	SourceApiAssociationStatusDeletionInProgress      = SourceApiAssociationStatus("DELETION_IN_PROGRESS")
	SourceApiAssociationStatusDeletionFailed          = SourceApiAssociationStatus("DELETION_FAILED")
)

type SourceApiAssociationStatusOutput struct{ *pulumi.OutputState }

func (SourceApiAssociationStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SourceApiAssociationStatus)(nil)).Elem()
}

func (o SourceApiAssociationStatusOutput) ToSourceApiAssociationStatusOutput() SourceApiAssociationStatusOutput {
	return o
}

func (o SourceApiAssociationStatusOutput) ToSourceApiAssociationStatusOutputWithContext(ctx context.Context) SourceApiAssociationStatusOutput {
	return o
}

func (o SourceApiAssociationStatusOutput) ToSourceApiAssociationStatusPtrOutput() SourceApiAssociationStatusPtrOutput {
	return o.ToSourceApiAssociationStatusPtrOutputWithContext(context.Background())
}

func (o SourceApiAssociationStatusOutput) ToSourceApiAssociationStatusPtrOutputWithContext(ctx context.Context) SourceApiAssociationStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SourceApiAssociationStatus) *SourceApiAssociationStatus {
		return &v
	}).(SourceApiAssociationStatusPtrOutput)
}

func (o SourceApiAssociationStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SourceApiAssociationStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SourceApiAssociationStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SourceApiAssociationStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SourceApiAssociationStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SourceApiAssociationStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SourceApiAssociationStatusPtrOutput struct{ *pulumi.OutputState }

func (SourceApiAssociationStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceApiAssociationStatus)(nil)).Elem()
}

func (o SourceApiAssociationStatusPtrOutput) ToSourceApiAssociationStatusPtrOutput() SourceApiAssociationStatusPtrOutput {
	return o
}

func (o SourceApiAssociationStatusPtrOutput) ToSourceApiAssociationStatusPtrOutputWithContext(ctx context.Context) SourceApiAssociationStatusPtrOutput {
	return o
}

func (o SourceApiAssociationStatusPtrOutput) Elem() SourceApiAssociationStatusOutput {
	return o.ApplyT(func(v *SourceApiAssociationStatus) SourceApiAssociationStatus {
		if v != nil {
			return *v
		}
		var ret SourceApiAssociationStatus
		return ret
	}).(SourceApiAssociationStatusOutput)
}

func (o SourceApiAssociationStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SourceApiAssociationStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SourceApiAssociationStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SourceApiAssociationConfigMergeTypeInput)(nil)).Elem(), SourceApiAssociationConfigMergeType("AUTO_MERGE"))
	pulumi.RegisterInputType(reflect.TypeOf((*SourceApiAssociationConfigMergeTypePtrInput)(nil)).Elem(), SourceApiAssociationConfigMergeType("AUTO_MERGE"))
	pulumi.RegisterOutputType(SourceApiAssociationConfigMergeTypeOutput{})
	pulumi.RegisterOutputType(SourceApiAssociationConfigMergeTypePtrOutput{})
	pulumi.RegisterOutputType(SourceApiAssociationStatusOutput{})
	pulumi.RegisterOutputType(SourceApiAssociationStatusPtrOutput{})
}
